package secret_test

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccSecretDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		secret, _ := testAccGetSecret(resourceName, state)
		if secret != nil {
			return fmt.Errorf("Found secret: %s", secret.Identifier)
		}

		return nil
	}
}

func testAccGetSecret(resourceName string, state *terraform.State) (*nextgen.Secret, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	options := &nextgen.SecretsApiGetSecretV2Opts{}

	if attr := r.Primary.Attributes["org_id"]; attr != "" {
		options.OrgIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["project_id"]; attr != "" {
		options.ProjectIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.SecretsApi.GetSecretV2(ctx, id, c.AccountId, options)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.Secret, nil
}
