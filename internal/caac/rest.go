package caac

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/httphelpers"
	"gopkg.in/yaml.v3"
)

func GetServiceById(applicationId string, serviceId string) *Service {

	return nil
}

func FindConfigAsCodeItemByPath(rootItem *ConfigAsCodeItem, path string) *ConfigAsCodeItem {

	if rootItem.DirectoryPath != nil && rootItem.DirectoryPath.Path == path {
		return rootItem
	}

	for _, item := range rootItem.Children {
		if matchingItem := FindConfigAsCodeItemByPath(item, path); matchingItem != nil {
			return matchingItem
		}
	}

	return nil
}

func (i *ConfigAsCodeItem) ParseYamlContent() (interface{}, error) {
	if i.Yaml == "" {
		return nil, nil
	}

	fmt.Println(i.Yaml)

	tmp := map[string]interface{}{}
	data := []byte(i.Yaml)

	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	val, ok := tmp["type"]
	if !ok {
		return nil, errors.New("could not find field 'type' in yaml object")
	}

	switch val {
	case ObjectTypes.Service:
		obj := &Service{}
		if err := yaml.Unmarshal(data, &obj); err != nil {
			return nil, err
		}
		obj.Name = TrimFileExtension(i.Name)
		return obj, err
	default:
		return nil, fmt.Errorf("could not parse object type of '%s'", val)
	}

}

func (c *ConfigAsCodeClient) GetDirectoryItemContent(restName string, uuid string, applicationId string) (*ConfigAsCodeItem, error) {
	path := fmt.Sprintf("/gateway/api/setup-as-code/yaml/%s/%s", restName, uuid)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)
	req.Header.Set(httphelpers.HeaderAuthorization, fmt.Sprintf("Bearer %s", c.ApiClient.BearerToken))

	// Set query parameters
	q := req.URL.Query()
	q.Add(client.QueryParamAccountId, c.ApiClient.AccountId)
	q.Add(client.QueryParamApplicationId, applicationId)
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) GetDirectoryTree(applicationId string) (*ConfigAsCodeItem, error) {
	path := "/gateway/api/setup-as-code/yaml/directory"

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, c.ApiClient.APIKey)
	req.Header.Set(httphelpers.HeaderContentType, httphelpers.HeaderApplicationJson)
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	// Set query parameters
	q := req.URL.Query()
	q.Add(client.QueryParamAccountId, c.ApiClient.AccountId)
	q.Add(client.QueryParamApplicationId, applicationId)
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) UpsertService(applicationName string, serviceName string, service interface{}) (*ConfigAsCodeItem, error) {
	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", applicationName, serviceName)
	return c.UpsertEntity(filePath, service)
}

func (c *ConfigAsCodeClient) UpsertEntity(filePath string, entity interface{}) (*ConfigAsCodeItem, error) {

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

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, "/gateway/api/setup-as-code/yaml/upsert-entity"), &b)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, c.ApiClient.APIKey)
	req.Header.Set(httphelpers.HeaderContentType, w.FormDataContentType())
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(client.QueryParamAccountId, c.ApiClient.AccountId)
	q.Add(client.QueryParamYamlFilePath, filePath)
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) ExecuteRequest(request *http.Request) (*ConfigAsCodeItem, error) {

	log.Printf("[DEBUG] Request url: %s", request.URL)

	res, err := c.ApiClient.HTTPClient.Do(request)
	if err != nil {
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
	log.Printf("[DEBUG] HTTP response: %s", responseString)

	if throttledRegex.MatchString(responseString) {
		return nil, errors.New(responseString)
	}

	responseObj := &Response{}

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

func (i *ConfigAsCodeItem) IsEmpty() bool {
	return i == &ConfigAsCodeItem{}
}

// Indicates an error condition
func (r *Response) IsEmpty() bool {
	// return true
	return r.Metadata == ResponseMetadata{} && r.Resource.IsEmpty() && len(r.ResponseMessages) == 0
}

func (m *ResponseMessage) ToError() error {
	return fmt.Errorf("%s: %s", m.Code, m.Message)
}

func TrimFileExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

const (
	statusSuccess = "SUCCESS"
)

var (
	throttledRegex = regexp.MustCompile("(?i).*throttled.*qps.*")
)
