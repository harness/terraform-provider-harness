package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

func init() {
	resource.AddTestSweepers("harness_cloudprovider", &resource.Sweeper{
		Name: "harness_cloudprovider",
		F:    testSweepCloudProviders,
	})
}
