package application

import (
	"log"
	"strings"

	"github.com/harness-io/terraform-provider-harness/internal/sweep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("harness_application", &resource.Sweeper{
		Name: "harness_application",
		F:    testSweepApplications,
	})
}

func testSweepApplications(r string) error {
	c := sweep.SweeperClient

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		apps, _, err := c.CDClient.ApplicationClient.ListApplications(limit, offset)
		if err != nil {
			return err
		}

		log.Printf("[INFO] Deleting %d applications", len(apps))

		for _, app := range apps {
			// Only delete applications that start with 'Test'
			if strings.HasPrefix(app.Name, "Test") {
				if err = c.CDClient.ApplicationClient.DeleteApplication(app.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(apps) == limit
	}

	return nil
}
