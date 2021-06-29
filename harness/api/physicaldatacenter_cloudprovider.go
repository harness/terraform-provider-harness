package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetPhysicalDataCEnterCloudProviderById(id string) (*graphql.PhysicalDataCenterCloudProvider, error) {
	cp := &graphql.PhysicalDataCenterCloudProvider{}
	err := c.getCloudProviderById(id, getPhyisicalDataCenterCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetPhysicalDatacenterCloudProviderByName(name string) (*graphql.PhysicalDataCenterCloudProvider, error) {
	cp := &graphql.PhysicalDataCenterCloudProvider{}
	err := c.getCloudProviderByName(name, getPhyisicalDataCenterCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreatePhysicalDataCenterCloudProvider(provider *graphql.PhysicalDataCenterCloudProvider) (*graphql.PhysicalDataCenterCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:               graphql.CloudProviderTypes.PhysicalDataCenter,
		PhysicalDataCenterCloudProvider: provider,
	}

	resp := &graphql.PhysicalDataCenterCloudProvider{}
	err := c.createCloudProvider(input, getPhyisicalDataCenterCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getPhyisicalDataCenterCloudProviderFields() string {
	return fmt.Sprintf(`
	... on PhysicalDataCenterCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}
