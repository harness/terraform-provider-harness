package cd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"gopkg.in/yaml.v3"
)

func FindConfigAsCodeItemByPath(rootItem *cac.ConfigAsCodeItem, path cac.YamlPath) *cac.ConfigAsCodeItem {

	if rootItem.DirectoryPath != nil && rootItem.DirectoryPath.Path == string(path) {
		return rootItem
	}

	for _, item := range rootItem.Children {
		if item == nil {
			continue
		}

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
		// There's a strange edgecase where child items are nil
		if item == nil {
			continue
		}

		if matchingItem := FindConfigAsCodeItemByUUID(item, uuid); matchingItem != nil {
			return matchingItem
		}
	}

	return nil
}

func (c *ConfigAsCodeClient) GetDirectoryItemContent(restName string, uuid string, applicationId string) (*cac.ConfigAsCodeItem, error) {
	path := fmt.Sprintf("/setup-as-code/yaml/%s/%s", restName, uuid)
	log.Printf("[DEBUG] CAC: Getting directory item content at %s", path)

	req, err := c.ApiClient.NewAuthorizedGetRequest(path)

	if err != nil {
		return nil, err
	}

	// Set query parameters
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)

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
	path := "/setup-as-code/yaml/directory"
	log.Printf("[DEBUG] CAC: Getting directory tree for app '%s'", applicationId)

	req, err := c.ApiClient.NewAuthorizedGetRequest(path)

	if err != nil {
		return nil, err
	}

	// Set query parameters
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)

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

	return c.UpsertRawYaml(filePath, payload)
}

func (c *ConfigAsCodeClient) UpsertRawYaml(filePath cac.YamlPath, yaml []byte) (*cac.ConfigAsCodeItem, error) {
	log.Printf("[DEBUG] CAC: Upserting yaml at %s", filePath)

	// Setup form fields
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("yamlContent")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, strings.NewReader(string(yaml))); err != nil {
		return nil, err
	}

	w.Close()

	log.Printf("[TRACE] CAC: HTTP Request Body %s", string(yaml))

	req, err := c.ApiClient.NewAuthorizedPostRequest("/setup-as-code/yaml/upsert-entity", &b)

	// Set proper content header
	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), w.FormDataContentType())

	if err != nil {
		return nil, err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)
	q.Add(helpers.QueryParameters.YamlFilePath.String(), string(filePath))
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) ExecuteRequest(request *retryablehttp.Request) (*cac.ConfigAsCodeItem, error) {

	log.Printf("[TRACE] CAC: Request url %s", request.URL)

	res, err := c.ApiClient.Configuration.HTTPClient.Do(request)
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
	log.Printf("[TRACE] CAC: HTTP response %d - %s", res.StatusCode, responseString)

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
	ApiClient *ApiClient
}

func (c *ConfigAsCodeClient) DeleteEntity(filePath cac.YamlPath) error {
	log.Printf("[DEBUG] CAC: Deleting entity at %s", filePath)
	req, err := c.ApiClient.NewAuthorizedDeleteRequest("/setup-as-code/yaml/delete-entities")

	if err != nil {
		return err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)
	q.Add(helpers.QueryParameters.FilePaths.String(), string(filePath))
	req.URL.RawQuery = q.Encode()

	log.Printf("[DEBUG] Url: %s", req.URL)

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
	if v, ok := input.(cac.Entity); ok {
		if ok, err := v.Validate(); !ok {
			return err
		}
	}

	// Upsert the yaml document
	resp, err := c.UpsertYamlEntity(filePath, input)
	if err != nil {
		return err
	}

	log.Printf("[TRACE] UUID: %s", resp.UUID)
	log.Printf("[TRACE] EntityId: %s", resp.EntityId)

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
	log.Printf("[DEBUG] CAC: Finding object by path %s", filePath)
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return err
	}

	item := FindConfigAsCodeItemByPath(rootItem, filePath)
	if item == nil {
		log.Printf("unable to find item at `%s`", filePath)
		return nil
	}

	return c.ParseObject(item, filePath, applicationId, obj)
}

func (c *ConfigAsCodeClient) FindYamlByPath(applicationId string, filePath cac.YamlPath) (*cac.YamlEntity, error) {
	log.Printf("[DEBUG] CAC: Find yaml by path %s", filePath)
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return nil, err
	}

	item := FindConfigAsCodeItemByPath(rootItem, filePath)
	if item == nil {
		log.Printf("unable to find item at `%s`", filePath)
		return nil, nil
	}

	return c.GetYamlDetails(item, filePath, applicationId)
}

func (c *ConfigAsCodeClient) FindObjectById(applicationId string, objectId string, out interface{}) error {
	log.Printf("[DEBUG] CAC: Find object by id %s", objectId)
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return err
	}

	i := FindConfigAsCodeItemByUUID(rootItem, objectId)
	if i == nil {
		log.Printf("[DEBUG] cannot find obj with id: " + objectId)
		return nil
	}

	return c.ParseObject(i, cac.YamlPath(i.DirectoryPath.Path), applicationId, out)
}

func (c *ConfigAsCodeClient) FindRootAccountObjectByName(name string) (*cac.ConfigAsCodeItem, error) {
	log.Printf("[DEBUG] CAC: Finding account by name %s", name)
	root, err := c.GetDirectoryTree("")
	if err != nil || root == nil {
		return root, err
	}

	for _, item := range root.Children {
		if item.Name == name {
			return item, nil
		}
	}

	return nil, nil
}

func (c *ConfigAsCodeClient) GetTemplateLibraryRootPathName() (cac.YamlPath, error) {
	libraryItem, err := c.FindRootAccountObjectByName("Template Library")
	if err != nil || libraryItem == nil {
		return cac.YamlPath(""), err
	}

	return cac.YamlPath(libraryItem.Children[0].DirectoryPath.Path), nil
}

func (c *ConfigAsCodeClient) GetYamlDetails(item *cac.ConfigAsCodeItem, filePath cac.YamlPath, applicationId string) (*cac.YamlEntity, error) {
	log.Printf("[DEBUG] CAC: Get yaml details %s", filePath)
	itemContent, err := c.GetDirectoryItemContent(item.RestName, item.UUID, applicationId)
	if err != nil {
		return nil, err
	}

	results := &cac.YamlEntity{
		Content:       itemContent.Yaml,
		Name:          cac.GetEntityNameFromPath(filePath),
		Id:            item.UUID,
		ApplicationId: applicationId,
		Path:          filePath,
	}

	return results, nil
}

func (c *ConfigAsCodeClient) ParseObject(item *cac.ConfigAsCodeItem, filePath cac.YamlPath, applicationId string, obj interface{}) error {
	log.Printf("[DEBUG] CAC: Prase yaml entity %s", filePath)
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
