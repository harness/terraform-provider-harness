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
	root, err := c.GetDirectoryTree("")
	if err != nil {
		return nil, err
	}

	appItems := findApplicationItems(root)
	if appItems == nil {
		return nil, nil
	}

	appItem := findApplicationItemById(appItems, name)
	if appItem == nil {
		return nil, nil
	}

	return c.GetApplicationById(appItem.AppId)
}

func findApplicationItems(rootItem *cac.ConfigAsCodeItem) *cac.ConfigAsCodeItem {

	for _, i := range rootItem.Children {
		if i.Name == "Applications" {
			return i
		}
	}

	return nil
}

func findApplicationItemById(rootItem *cac.ConfigAsCodeItem, appName string) *cac.ConfigAsCodeItem {
	for _, i := range rootItem.Children {
		if i.Name == appName {
			return i
		}
	}

	return nil
}
