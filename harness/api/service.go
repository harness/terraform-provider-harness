package api

import (
	"errors"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
)

type ServiceClient struct {
	ConfigAsCodeClient *ConfigAsCodeClient
}

func (c *Client) Services() *ServiceClient {
	return &ServiceClient{
		ConfigAsCodeClient: &ConfigAsCodeClient{
			ApiClient: c,
		},
	}
}

func (c *ServiceClient) UpsertService(serviceInput *cac.Service) (*cac.Service, error) {
	if serviceInput == nil {
		return nil, errors.New("service is nil")
	}

	if ok, err := serviceInput.Validate(); !ok {
		return nil, err
	}

	app, err := c.ConfigAsCodeClient.ApiClient.Applications().GetApplicationById(serviceInput.ApplicationId)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, fmt.Errorf("could not find application by id: '%s'", serviceInput.ApplicationId)
	}

	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", app.Name, serviceInput.Name)

	item, err := c.ConfigAsCodeClient.UpsertEntity(filePath, serviceInput)
	if err != nil {
		return nil, err
	}

	rootItem, err := c.ConfigAsCodeClient.GetDirectoryTree(app.Id)
	if err != nil {
		return nil, err
	}

	svcItem := FindConfigAsCodeItemByPath(rootItem, item.YamlFilePath)
	if svcItem == nil {
		return nil, fmt.Errorf("unable to find service '%s'", serviceInput.Name)
	}

	itemContent, err := c.ConfigAsCodeClient.GetDirectoryItemContent(svcItem.RestName, svcItem.UUID, serviceInput.ApplicationId)
	if err != nil {
		return nil, err
	}

	svcObj, err := itemContent.ParseYamlContent()
	if err != nil {
		return nil, err
	}

	svc := svcObj.(*cac.Service)
	svc.Id = svcItem.UUID
	svc.ApplicationId = serviceInput.ApplicationId

	return svc, nil
}

func (c *ServiceClient) GetServiceById(applicationId string, serviceId string) (*cac.Service, error) {
	item, err := c.ConfigAsCodeClient.GetDirectoryItemContent("services", serviceId, applicationId)
	if err != nil {
		return nil, err
	}

	// Item not found
	if item == nil {
		return nil, nil
	}

	obj, err := item.ParseYamlContent()
	if err != nil {
		return nil, err
	}

	svc := obj.(*cac.Service)
	svc.ApplicationId = applicationId

	return svc, nil
}

func (c *ServiceClient) DeleteService(applicationId string, serviceId string) error {

	app, err := c.ConfigAsCodeClient.ApiClient.Applications().GetApplicationById(applicationId)
	if err != nil {
		return err
	}

	if app == nil {
		return fmt.Errorf("could not find application by id: '%s'", applicationId)
	}

	svc, err := c.GetServiceById(applicationId, serviceId)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", app.Name, svc.Name)

	return c.ConfigAsCodeClient.DeleteEntities([]string{filePath})
}
