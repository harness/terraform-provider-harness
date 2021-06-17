package api

import (
	"errors"
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/api/caac"
)

func (c *Client) Services() *ConfigAsCodeClient {
	return &ConfigAsCodeClient{
		ApiClient: c,
	}
}

func (c *ConfigAsCodeClient) UpsertService(serviceInput *caac.Service) (*caac.Service, error) {
	if serviceInput == nil {
		return nil, errors.New("service is nil")
	}

	if ok, err := serviceInput.Validate(); !ok {
		return nil, err
	}

	app, err := c.ApiClient.Applications().GetApplicationById(serviceInput.ApplicationId)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, fmt.Errorf("could not find application by id: '%s'", serviceInput.ApplicationId)
	}

	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", app.Name, serviceInput.Name)

	item, err := c.UpsertEntity(filePath, serviceInput)
	if err != nil {
		return nil, err
	}

	rootItem, err := c.GetDirectoryTree(app.Id)
	if err != nil {
		return nil, err
	}

	svcItem := FindConfigAsCodeItemByPath(rootItem, item.YamlFilePath)
	if svcItem == nil {
		return nil, fmt.Errorf("unable to find service '%s'", serviceInput.Name)
	}

	itemContent, err := c.GetDirectoryItemContent(svcItem.RestName, svcItem.UUID, serviceInput.ApplicationId)
	if err != nil {
		return nil, err
	}

	svcObj, err := itemContent.ParseYamlContent()
	if err != nil {
		return nil, err
	}

	svc := svcObj.(*caac.Service)
	svc.Id = svcItem.UUID
	svc.ApplicationId = serviceInput.ApplicationId

	return svc, nil
}

func (c *ConfigAsCodeClient) GetServiceById(applicationId string, serviceId string) (*caac.Service, error) {
	item, err := c.GetDirectoryItemContent("services", serviceId, applicationId)
	if err != nil {
		return nil, err
	}

	obj, err := item.ParseYamlContent()
	if err != nil {
		return nil, err
	}

	return obj.(*caac.Service), nil
}
