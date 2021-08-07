package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"gopkg.in/yaml.v3"
)

func (c *Client) ConfigAsCode() *ConfigAsCodeClient {
	return &ConfigAsCodeClient{
		ApiClient: c,
	}
}

func FindConfigAsCodeItemByPath(rootItem *cac.ConfigAsCodeItem, path cac.YamlPath) *cac.ConfigAsCodeItem {

	if rootItem.DirectoryPath != nil && rootItem.DirectoryPath.Path == string(path) {
		return rootItem
	}

	for _, item := range rootItem.Children {
		if matchingItem := FindConfigAsCodeItemByPath(item, path); matchingItem != nil {
			return matchingItem
		}
	}

	return nil
}

func FindConfigAsCodeItemByUUID(rootItem *cac.ConfigAsCodeItem, uuid string) *cac.ConfigAsCodeItem {
	if rootItem.DirectoryPath != nil && rootItem.UUID == uuid {
		return rootItem
	}

	for _, item := range rootItem.Children {
		if matchingItem := FindConfigAsCodeItemByUUID(item, uuid); matchingItem != nil {
			return matchingItem
		}
	}

	return nil
}

func (c *ConfigAsCodeClient) GetDirectoryItemContent(restName string, uuid string, applicationId string) (*cac.ConfigAsCodeItem, error) {
	path := fmt.Sprintf("/gateway/api/setup-as-code/yaml/%s/%s", restName, uuid)

	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.Authorization.String(), fmt.Sprintf("Bearer %s", c.ApiClient.BearerToken))

	// Set query parameters
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.AccountId)

	if applicationId != "" {
		q.Add(helpers.QueryParameters.ApplicationId.String(), applicationId)
	}
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) GetDirectoryTree(applicationId string) (*cac.ConfigAsCodeItem, error) {
	path := "/gateway/api/setup-as-code/yaml/directory"

	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), c.ApiClient.APIKey)
	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())

	// Set query parameters
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.AccountId)

	if applicationId != "" {
		q.Add(helpers.QueryParameters.ApplicationId.String(), applicationId)
	}

	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) UpsertYamlEntity(filePath cac.YamlPath, entity interface{}) (*cac.ConfigAsCodeItem, error) {

	payload, err := yaml.Marshal(&entity)
	if err != nil {
		return nil, err
	}

	// Setup form fields
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("yamlContent")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, strings.NewReader(string(payload))); err != nil {
		return nil, err
	}

	w.Close()

	log.Printf("[DEBUG] HTTP Request Body: %s", string(payload))

	req, err := retryablehttp.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, "/gateway/api/setup-as-code/yaml/upsert-entity"), &b)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), c.ApiClient.APIKey)
	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), w.FormDataContentType())
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.AccountId)
	q.Add(helpers.QueryParameters.YamlFilePath.String(), string(filePath))
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) ExecuteRequest(request *retryablehttp.Request) (*cac.ConfigAsCodeItem, error) {

	log.Printf("[DEBUG] Request url: %s", request.URL)

	res, err := c.ApiClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	if ok, err := checkStatusCode(res.StatusCode); !ok {
		return nil, err
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	// Check for request throttling
	responseString := buf.String()
	log.Printf("[DEBUG] HTTP response: %d - %s", res.StatusCode, responseString)

	responseObj := &cac.Response{}

	// Unmarshal into our response object
	if err := json.NewDecoder(&buf).Decode(&responseObj); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	if responseObj.IsEmpty() {
		return nil, errors.New("received an empty response")
	}

	if len(responseObj.ResponseMessages) > 0 {
		return nil, responseObj.ResponseMessages[0].ToError()
	}

	if responseObj.Resource.Status != "" && responseObj.Resource.Status != statusSuccess {
		return nil, fmt.Errorf("%s: %s", responseObj.Resource.Status, responseObj.Resource.ErrorMessage)
	}

	return &responseObj.Resource, nil
}

const (
	statusSuccess = "SUCCESS"
)

func checkStatusCode(statusCode int) (bool, error) {
	if statusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("received http status code '%d'", statusCode)
}

type ConfigAsCodeClient struct {
	ApiClient *Client
}

func (c *ConfigAsCodeClient) DeleteEntity(filePath cac.YamlPath) error {

	req, err := retryablehttp.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, "/gateway/api/setup-as-code/yaml/delete-entities"), nil)
	if err != nil {
		return err
	}

	// Configure additional headers
	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), c.ApiClient.APIKey)
	req.Header.Set(helpers.HTTPHeaders.Authorization.String(), fmt.Sprintf("Bearer %s", c.ApiClient.BearerToken))
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.AccountId)
	q.Add(helpers.QueryParameters.FilePaths.String(), string(filePath))
	req.URL.RawQuery = q.Encode()

	log.Printf("[DEBUG] Url: %s", req.URL)
	log.Printf("[DEBUG] Headers: %s", req.Header)

	resp, err := c.ExecuteRequest(req)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func (c *ConfigAsCodeClient) UpsertObject(input interface{}, filePath cac.YamlPath, responseObj interface{}) error {
	if input == nil {
		return errors.New("object to upsert is nil")
	}

	if ok, err := utils.RequiredFieldsCheck(input, []string{"Name", "Id"}); !ok {
		return err
	}

	// If the object implements the Validation interface then check it
	if v, ok := input.(cac.Validation); ok {
		if ok, err := v.Validate(); !ok {
			return err
		}
	}

	// Upsert the yaml document
	_, err := c.UpsertYamlEntity(filePath, input)
	if err != nil {
		return err
	}

	appId, ok := utils.TryGetFieldValue(input, "ApplicationId")
	if !ok {
		appId = ""
	}

	err = c.FindObjectByPath(appId.(string), filePath, responseObj)
	if err != nil {
		return err
	}

	return nil
}

// This function is used when the Id of the object is unknown and we need to look it up.
// Typically this is needed just after an Upsert command. The Upsert API unfortunately does not
// return the Id of the newly created object.
func (c *ConfigAsCodeClient) FindObjectByPath(applicationId string, filePath cac.YamlPath, obj interface{}) error {
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return err
	}

	item := FindConfigAsCodeItemByPath(rootItem, filePath)
	if item == nil {
		return fmt.Errorf("unable to item at `%s`", filePath)
	}

	return c.ParseObject(item, filePath, applicationId, obj)
}

func (c *ConfigAsCodeClient) FindObjectById(applicationId string, objectId string, out interface{}) error {
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return err
	}

	i := FindConfigAsCodeItemByUUID(rootItem, objectId)
	if i == nil {
		return errors.New("cannot find obj with id: " + objectId)
	}

	return c.ParseObject(i, cac.YamlPath(i.DirectoryPath.Path), applicationId, out)
}

func (c *ConfigAsCodeClient) ParseObject(item *cac.ConfigAsCodeItem, filePath cac.YamlPath, applicationId string, obj interface{}) error {
	itemContent, err := c.GetDirectoryItemContent(item.RestName, item.UUID, applicationId)
	if err != nil {
		return err
	}

	// Parse the new entity
	err = itemContent.ParseYamlContent(obj)
	if err != nil {
		return err
	}

	// Set the required fields for all entities
	utils.MustSetField(obj, "Name", cac.GetEntityNameFromPath(filePath))
	utils.MustSetField(obj, "Id", item.UUID)

	if applicationId != "" && utils.HasField(obj, "ApplicationId") {
		utils.MustSetField(obj, "ApplicationId", applicationId)
	}

	return nil
}
