package webhook_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceGitxWebhookProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	updatedName := fmt.Sprintf("%s_updated", id)
	resourceName := "harness_platform_gitx_webhook.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGitXProjectProjectLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testGitXProjectProjectLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsWebhookImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestResourceGitxWebhookOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	updatedName := fmt.Sprintf("%s_updated", id)
	resourceName := "harness_platform_gitx_webhook.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGitXProjectOrgLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testGitXProjectOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsWebhookImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestResourceGitxWebhookAccountLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	updatedName := fmt.Sprintf("%s_updated", id)
	resourceName := "harness_platform_gitx_webhook.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGitXAccountLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testGitXAccountLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsWebhookImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccResourceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		variable, _ := testAccGetResourceWebhook(resourceName, state)
		if variable != nil {
			return fmt.Errorf("Found variable: %s", variable.WebhookIdentifier)
		}
		return nil
	}
}

func testAccGetResourceWebhook(resourceName string, state *terraform.State) (*nextgen.GitXWebhookResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	orgIdentifier := r.Primary.Attributes["org_id"]
	projectIdentifier := r.Primary.Attributes["project_id"]
	id := r.Primary.ID

	if len(orgIdentifier) > 0 && len(projectIdentifier) > 0 {
		resp, _, err := c.ProjectGitxWebhooksApiService.GetProjectGitxWebhook(ctx, orgIdentifier, projectIdentifier, id, &nextgen.ProjectGitxWebhooksApiGetProjectGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, err
		}
		return &resp, nil
	} else if len(orgIdentifier) > 0 {
		resp, _, err := c.OrgGitxWebhooksApiService.GetOrgGitxWebhook(ctx, orgIdentifier, id, &nextgen.OrgGitxWebhooksApiGetOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, err
		}
		return &resp, nil
	} else {
		resp, _, err := c.GitXWebhooksApiService.GetGitxWebhook(ctx, id, &nextgen.GitXWebhooksApiGetGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, err
		}
		return &resp, nil
	}
}

func testGitXProjectOrgLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, webhook_identifier, webhook_name)
}

func testGitXAccountLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, webhook_identifier, webhook_name)
}

func testGitXProjectProjectLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
			project_id = "shivam"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, webhook_identifier, webhook_name)
}
