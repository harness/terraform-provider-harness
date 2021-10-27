package cd_test

import (
	"strings"

	"github.com/harness-io/terraform-provider-harness/internal/acctest"
)

func testSweepCloudProviders(r string) error {
	c := acctest.TestAccGetApiClientFromProvider()

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		cloudProviders, _, err := c.CloudProviders().ListCloudProviders(limit, offset)
		if err != nil {
			return err
		}

		for _, cp := range cloudProviders {
			if strings.HasPrefix(cp.Name, "Test") {
				if err = c.CloudProviders().DeleteCloudProvider(cp.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(cloudProviders) == limit
	}

	return nil
}
