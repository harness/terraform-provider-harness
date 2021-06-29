package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetAzureCloudProviderById(Id string) (*graphql.AzureCloudProvider, error) {
	cp := &graphql.AzureCloudProvider{}
	err := c.getCloudProviderById(Id, getAzureCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetAzureCloudProviderByName(name string) (*graphql.AzureCloudProvider, error) {
	cp := &graphql.AzureCloudProvider{}
	err := c.getCloudProviderByName(name, getAzureCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateAzureCloudProvider(provider *graphql.AzureCloudProvider) (*graphql.AzureCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:  graphql.CloudProviderTypes.Azure,
		AzureCloudProvider: provider,
	}

	resp := &graphql.AzureCloudProvider{}
	err := c.createCloudProvider(input, getAzureCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getAzureCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
	`, commonCloudProviderFields)
}
