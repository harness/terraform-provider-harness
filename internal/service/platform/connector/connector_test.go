package connector_test

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccConnectorDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		connector, _ := testAccGetConnector(resourceName, state)
		if connector != nil {
			return fmt.Errorf("Found connector: %s", connector.Identifier)
		}

		return nil
	}
}

func testAccGetConnector(resourceName string, state *terraform.State) (*nextgen.ConnectorInfo, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	options := &nextgen.ConnectorsApiGetConnectorOpts{}

	if attr := r.Primary.Attributes["org_id"]; attr != "" {
		options.OrgIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["project_id"]; attr != "" {
		options.ProjectIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["branch"]; attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["repo_id"]; attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["repo_id"]; attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.ConnectorsApi.GetConnector(ctx, c.AccountId, id, options)
	if err != nil {
		return nil, err
	}

	return resp.Data.Connector, nil
}
