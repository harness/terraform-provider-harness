package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetGcpCloudProviderById(id string) (*graphql.GcpCloudProvider, error) {
	cp := &graphql.GcpCloudProvider{}
	err := c.getCloudProviderById(id, getGcpCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetGcpCloudProviderByName(name string) (*graphql.GcpCloudProvider, error) {
	cp := &graphql.GcpCloudProvider{}
	err := c.getCloudProviderByName(name, getGcpCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateGcpCloudProvider(provider *graphql.GcpCloudProvider) (*graphql.GcpCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Gcp,
		GCPCloudProvider:  provider,
	}

	resp := &graphql.GcpCloudProvider{}
	err := c.createCloudProvider(input, getGcpCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getGcpCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on GcpCloudProvider {
			delegateSelectors
		}
	`, commonCloudProviderFields)
}
