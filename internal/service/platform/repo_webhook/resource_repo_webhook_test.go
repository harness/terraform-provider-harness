package repo_webhook_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	resourceName       = "harness_platform_repo_webhook.test"
	description        = "example_description"
	updatedDescription = "example_description_updated"
)

func TestProjResourceRepoWebhook(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testWebhookDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepoWebhook(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				Config: testProjResourceRepoWebhook(identifier, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.RepoRuleProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func identifier(testName string) string {
	return fmt.Sprintf("%s_%s", testName, utils.RandStringBytes(5))
}

func testProjResourceRepoWebhook(identifier string, description string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "org_%[1]s"
			name = "org_%[1]s"
		}	

		resource "harness_platform_project" "test" {
			identifier = "proj_%[1]s"
			name = "proj_%[1]s"
			org_id = harness_platform_organization.test.id
		}
		
		resource "harness_platform_repo" "test" {
			identifier  = "repo_%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			default_branch = "master"
			readme = true
		}

		resource "harness_platform_repo_webhook" "test" {
			identifier  = "%[1]s"
			repo_identifier = harness_platform_repo.test.identifier
			description = "%[2]s"
			url = "http://harness.io"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			enabled = true
			insecure = true
			triggers = ["branch_deleted"]
		}
	`, identifier, description,
	)
}

func testFindRepoWebhook(
	resourceName string,
	state *terraform.State,
) (*code.OpenapiWebhookType, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	repoIdentifier := r.Primary.Attributes["repo_identifier"]
	webhookId := r.Primary.Attributes["identifier"]

	webhook, _, err := c.WebhookApi.GetWebhook(
		ctx, c.AccountId, repoIdentifier, webhookId,
		&code.WebhookApiGetWebhookOpts{
			OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
			ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
		})
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

func testWebhookDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		webhook, _ := testFindRepoWebhook(resourceName, state)
		if webhook != nil {
			return fmt.Errorf("found webhook: %s", webhook.Identifier)
		}

		return nil
	}
}
