package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
)

func (c *ConfigAsCodeClient) GetServiceById(applicationId string, serviceId string) (*cac.Service, error) {
	item, err := c.GetDirectoryItemContent("services", serviceId, applicationId)
	if err != nil {
		return nil, err
	}

	// Item not found
	if item == nil {
		return nil, nil
	}

	svc := &cac.Service{}
	err = item.ParseYamlContent(svc)
	if err != nil {
		return nil, err
	}

	// // Set the required fields for all entities
	svc.Name = utils.TrimFileExtension(item.Name)
	svc.Id = serviceId
	svc.ApplicationId = applicationId

	return svc, nil
}

func (c *ConfigAsCodeClient) UpsertService(input *cac.Service) (*cac.Service, error) {
	app, err := c.ApiClient.Applications().GetApplicationById(input.ApplicationId)
	if err != nil {
		return nil, err
	}

	if ok, err := input.Validate(); !ok {
		return nil, err
	}

	yamlPath := cac.GetServiceYamlPath(app.Name, input.Name)
	svc := &cac.Service{}
	err = c.UpsertObject(input, yamlPath, svc)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func (c *ConfigAsCodeClient) DeleteService(applicationId string, serviceId string) error {
	app, err := c.ApiClient.Applications().GetApplicationById(applicationId)
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

	filePath := cac.GetServiceYamlPath(app.Name, svc.Name)

	return c.DeleteEntity(filePath)
}
