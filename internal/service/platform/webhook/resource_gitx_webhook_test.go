package webhook_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceGitxWebhookProjectLevel(t *testing.T) {
	resourceName := "harness_platform_gitx_webhook.test"
	accountId := "rXUXvbFqRr2XwcjBu3Oq-Q"
	webhook_identifier := "WebhookTest"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGitXProjectProjectLevel(webhook_identifier, accountId, webhook_identifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", webhook_identifier),
				),
			},
			{
				Config: testGitXProjectProjectLevel(webhook_identifier, accountId, "WebhookNew2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "WebhookNew2"),
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

func TestResourceGitxWebhookOrgLevel(t *testing.T) {
	resourceName := "harness_platform_gitx_webhook.test2"
	accountId := "rXUXvbFqRr2XwcjBu3Oq-Q"
	webhook_identifier := "WebhookTestOrg"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGitXProjectOrgLevel(webhook_identifier, accountId, webhook_identifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", webhook_identifier),
				),
			},
			{
				Config: testGitXProjectOrgLevel(webhook_identifier, accountId, "WebhookNewOrg2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "WebhookNewOrg2"),
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

func testGitXProjectOrgLevel(webhook_identifier string, accountId string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test2" {
			identifier= "%[2]s"
			name = "%[3]s"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, accountId, webhook_identifier, webhook_name)
}

func testGitXProjectProjectLevel(webhook_identifier string, accountId string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[2]s"
			name = "%[3]s"
			project_id = "shivam"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, accountId, webhook_identifier, webhook_name)
}
