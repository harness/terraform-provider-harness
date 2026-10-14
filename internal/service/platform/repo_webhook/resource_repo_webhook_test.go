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
	"github.com/stretchr/testify/require"
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

func TestProjResourceRepoWebhookWithSecret(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testWebhookDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepoWebhookWithSecret(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				Config: testProjResourceRepoWebhookWithSecret(identifier, updatedDescription),
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

// TestProjResourceRepoWebhook_DeleteUnderlyingResource verifies that a webhook deleted
// outside Terraform is dropped from state on refresh instead of failing with a 404.
func TestProjResourceRepoWebhook_DeleteUnderlyingResource(t *testing.T) {
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
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.WebhookApi.DeleteWebhook(
						ctx, c.AccountId, "repo_"+identifier, identifier,
						&code.WebhookApiDeleteWebhookOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("proj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testProjResourceRepoWebhook(identifier, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestProjResourceRepoWebhook_DeleteParentRepo verifies that a webhook whose parent
// repo was deleted outside Terraform is dropped from state on refresh instead of
// failing with a 404. This is the destroy-retry scenario from CODE-6342.
func TestProjResourceRepoWebhook_DeleteParentRepo(t *testing.T) {
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
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, _, err := c.RepositoryApi.DeleteRepository(
						ctx, c.AccountId, "repo_"+identifier,
						&code.RepositoryApiDeleteRepositoryOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("proj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testProjResourceRepoWebhook(identifier, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
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

func testProjResourceRepoWebhookWithSecret(identifier string, description string) string {
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
			secret = "asd"
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
		// The resource may already be gone from state (e.g. deleted out-of-band and
		// dropped on refresh), in which case there is nothing left to look up.
		if _, ok := state.RootModule().Resources[resourceName]; !ok {
			return nil
		}
		webhook, _ := testFindRepoWebhook(resourceName, state)
		if webhook != nil {
			return fmt.Errorf("found webhook: %s", webhook.Identifier)
		}

		return nil
	}
}
