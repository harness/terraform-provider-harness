package caac

import (
	"errors"
	"fmt"
)

func (c *ConfigAsCodeClient) UpsertService(serviceInput *Service) (*Service, error) {
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

	svc := svcObj.(*Service)
	svc.Id = svcItem.UUID
	svc.ApplicationId = serviceInput.ApplicationId

	return svc, nil
}

func (s *Service) Validate() (bool, error) {
	if s.ApplicationId == "" {
		return false, errors.New("service is invalid. missing field `ApplicationId`")
	}

	return true, nil
}

// func (c *ConfigAsCodeClient) UpsertService(applicationName string, serviceName string, service interface{}) (*ConfigAsCodeItem, error) {
// 	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", applicationName, serviceName)
// 	return c.UpsertEntity(filePath, service)
// }
