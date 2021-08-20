package api

import (
	"errors"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
)

func (c *ConfigAsCodeClient) UpsertSpotInstCloudProvider(input *cac.SpotInstCloudProvider) (*cac.SpotInstCloudProvider, error) {
	out := &cac.SpotInstCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertPcfCloudProvider(input *cac.PcfCloudProvider) (*cac.PcfCloudProvider, error) {
	out := &cac.PcfCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertKubernetesCloudProvider(input *cac.KubernetesCloudProvider) (*cac.KubernetesCloudProvider, error) {
	out := &cac.KubernetesCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertAzureCloudProvider(input *cac.AzureCloudProvider) (*cac.AzureCloudProvider, error) {
	out := &cac.AzureCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertGcpCloudProvider(input *cac.GcpCloudProvider) (*cac.GcpCloudProvider, error) {
	out := &cac.GcpCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertPhysicalDataCenterCloudProvider(input *cac.PhysicalDatacenterCloudProvider) (*cac.PhysicalDatacenterCloudProvider, error) {
	out := &cac.PhysicalDatacenterCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertAwsCloudProvider(input *cac.AwsCloudProvider) (*cac.AwsCloudProvider, error) {
	out := &cac.AwsCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertCloudProvider(input interface{}, output interface{}) error {
	if input == nil {
		return errors.New("cannot create cloud provider. input is nil")
	}

	name, ok := utils.TryGetFieldValue(input, "Name")
	if !ok || name == "" {
		return errors.New("expected cloud provider to have Name field set")
	}

	filePath := cac.GetCloudProviderYamlPath(name.(string))
	err := c.ApiClient.ConfigAsCode().UpsertObject(input, filePath, output)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigAsCodeClient) GetCloudProviderById(providerId string, out interface{}) error {
	rootItem, err := c.GetDirectoryTree("")
	if err != nil {
		return err
	}

	i := FindConfigAsCodeItemByUUID(rootItem, providerId)
	if i == nil {
		return errors.New("cannot find cloud provider with id: " + providerId)
	}

	return c.ParseObject(i, cac.YamlPath(i.DirectoryPath.Path), "", out)
}

func (c *ConfigAsCodeClient) GetCloudProviderByName(name string, obj interface{}) error {
	filePath := cac.GetCloudProviderYamlPath(name)
	return c.FindObjectByPath("", filePath, obj)
}

// func (c *ConfigAsCodeClient) DeleteCloudProvider(name string) error {
// 	filePath := cac.GetCloudProviderYamlPath(name)
// 	return c.DeleteEntity(filePath)
// }
