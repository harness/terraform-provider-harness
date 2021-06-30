package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetSpotInstCloudProviderById(id string) (*graphql.SpotInstCloudProvider, error) {
	cp := &graphql.SpotInstCloudProvider{}
	err := c.getCloudProviderById(id, getSpotInstCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetSpotInstCloudProviderByName(name string) (*graphql.SpotInstCloudProvider, error) {
	cp := &graphql.SpotInstCloudProvider{}
	err := c.getCloudProviderByName(name, getSpotInstCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateSpotInstCloudProvider(provider *graphql.SpotInstCloudProvider) (*graphql.SpotInstCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:     graphql.CloudProviderTypes.SpotInst,
		SpotInstCloudProvider: provider,
	}

	resp := &graphql.SpotInstCloudProvider{}
	err := c.createCloudProvider(input, getSpotInstCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudProviderClient) UpdateSpotInstCloudProvider(id string, cp *graphql.UpdateSpotInstCloudProviderInst) (*graphql.SpotInstCloudProvider, error) {
	input := &graphql.UpdateCloudProvider{
		SpotInstCloudProvider: cp,
		CloudProviderType:     &graphql.CloudProviderTypes.SpotInst,
		CloudProviderId:       id,
	}

	resp := &graphql.SpotInstCloudProvider{}
	err := c.updateCloudProvider(input, getSpotInstCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getSpotInstCloudProviderFields() string {
	return fmt.Sprintf(`
	... on SpotInstCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}
