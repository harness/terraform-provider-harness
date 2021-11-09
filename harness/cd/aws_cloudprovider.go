package cd

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
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

func (c *CloudProviderClient) UpdateAwsCloudProvider(id string, cp *graphql.UpdateAwsCloudProviderInput) (*graphql.AwsCloudProvider, error) {
	input := &graphql.UpdateCloudProvider{
		AwsCloudProvider:  cp,
		CloudProviderType: &graphql.CloudProviderTypes.Aws,
		CloudProviderId:   id,
	}

	resp := &graphql.AwsCloudProvider{}
	err := c.updateCloudProvider(input, getAwsCloudProviderFields(), resp)
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
