package repo_test

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
	resourceName = "harness_platform_repo.test"
	description  = "example_description"
	providerRepo = "octocat/hello-worId"
)

var accountId = utils.GetEnv("HARNESS_ACCOUNT_ID", "")

func TestAccResourceRepo(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(identifier, description, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testAccResourceRepo(identifier, description, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceRepo_Import(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(identifier, description, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "provider_repo", providerRepo),
				),
			},
			{
				Config: testAccResourceRepo(identifier, description, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "provider_repo", providerRepo),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceRepo_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()

	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(identifier, description, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.DeleteRepository(
						ctx, accountId, identifier,
						&code.RepositoryApiDeleteRepositoryOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("prj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testAccResourceRepo(identifier, description, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func identifier(testName string) string {
	return fmt.Sprintf("%s_%s", testName, utils.RandStringBytes(5))
}

func testAccResourceRepo(identifier, description, providerRepo string) string {
	accountId := utils.GetEnv("HARNESS_ACCOUNT_ID", "")
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "org_%[1]s"
			name = "org_%[1]s"
		}	

		resource "harness_platform_project" "test" {
			identifier = "prj_%[1]s"
			name = "prj_%[1]s"
			org_id = harness_platform_organization.test.id
		}
		
		resource "harness_platform_repo" "test" {
			identifier  = "%[1]s"
			name       	= "%[1]s"
			org_identifier = harness_platform_organization.test.id
			project_identifier = harness_platform_project.test.id
			default_branch = "master"
			description = "%[2]s"
			account_id 	= "%[3]s"
			provider_repo = "%[4]s"
			is_public = true
			type = "github"
		}
	`, identifier, description, accountId, providerRepo,
	)
}

func testAccFindRepo(
	resourceName string,
	state *terraform.State,
) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	identifier := r.Primary.Attributes["identifier"]

	repo, _, err := c.RepositoryApi.FindRepository(
		ctx, accountId, identifier,
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
			ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
		})
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

func testAccRepoDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repo, _ := testAccFindRepo(resourceName, state)
		if repo != nil {
			return fmt.Errorf("Found repo: %s", repo.Identifier)
		}

		return nil
	}
}
