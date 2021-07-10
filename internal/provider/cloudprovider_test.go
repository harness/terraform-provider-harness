package provider

import "strings"

func testSweepCloudProviders(r string) error {
	c := testAccGetApiClientFromProvider()

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		cloudProviders, pagination, err := c.CloudProviders().ListCloudProviders(limit, offset)
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

		hasMore = pagination.HasMore
		offset += 1
	}

	return nil
}
