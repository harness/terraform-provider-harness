package cd

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
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

func (c *CloudProviderClient) UpdateAzureCloudProvider(id string, cp *graphql.UpdateAzureCloudProviderInput) (*graphql.AzureCloudProvider, error) {
	input := &graphql.UpdateCloudProvider{
		AzureCloudProvider: cp,
		CloudProviderType:  &graphql.CloudProviderTypes.Azure,
		CloudProviderId:    id,
	}

	resp := &graphql.AzureCloudProvider{}
	err := c.updateCloudProvider(input, getAzureCloudProviderFields(), resp)
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
