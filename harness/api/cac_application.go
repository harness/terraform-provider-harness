package api

import (
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
)

func (c *ConfigAsCodeClient) GetApplicationById(applicationId string) (*cac.Application, error) {
	app := &cac.Application{}
	err := c.FindObjectById(applicationId, applicationId, app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (c *ConfigAsCodeClient) GetApplicationByName(name string) (*cac.Application, error) {
	appItems, err := c.FindRootAccountObjectByName("Applications")
	if err != nil || appItems == nil {
		return nil, err
	}

	appItem := findApplicationItemById(appItems, name)
	if appItem == nil {
		return nil, nil
	}

	return c.GetApplicationById(appItem.AppId)
}

func findApplicationItemById(rootItem *cac.ConfigAsCodeItem, appName string) *cac.ConfigAsCodeItem {
	for _, i := range rootItem.Children {
		if i.Name == appName {
			return i
		}
	}

	return nil
}
