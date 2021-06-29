package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetAwsCloudProviderById(Id string) (*graphql.AwsCloudProvider, error) {
	cp := &graphql.AwsCloudProvider{}
	err := c.getCloudProviderById(Id, getAwsCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetAwsCloudProviderByName(name string) (*graphql.AwsCloudProvider, error) {
	cp := &graphql.AwsCloudProvider{}
	err := c.getCloudProviderByName(name, getAwsCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateAwsCloudProvider(provider *graphql.AwsCloudProvider) (*graphql.AwsCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Aws,
		AwsCloudProvider:  provider,
	}

	resp := &graphql.AwsCloudProvider{}
	err := c.createCloudProvider(input, getAwsCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getAwsCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on AwsCloudProvider {
			%[2]s
		}
	`, commonCloudProviderFields, ceHealthStatusFields)
}
