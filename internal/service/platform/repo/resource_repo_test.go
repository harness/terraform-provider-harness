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
	resourceName        = "harness_platform_repo.test"
	description         = "example_description"
	description_updated = "example_description_updated"
	providerRepo        = "octocat/hello-worId"
)

func TestProjResourceRepo(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepo(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testProjResourceRepo(identifier, description_updated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description_updated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.RepoResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceRepo(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testAccResourceRepo(identifier, description_updated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description_updated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.RepoResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceRepo_Import(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepoImport(identifier, description, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testAccResourceRepoImport(identifier, description_updated, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "name", identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description_updated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.RepoResourceImportStateIdFunc(resourceName),
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
		CheckDestroy:      testRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepo(identifier, description),
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
						ctx, c.AccountId, identifier,
						&code.RepositoryApiDeleteRepositoryOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("proj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testProjResourceRepo(identifier, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func identifier(testName string) string {
	return fmt.Sprintf("%s_%s", testName, utils.RandStringBytes(5))
}

func testProjResourceRepo(identifier, description string) string {
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
			identifier  = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			default_branch = "master"
			description = "%[2]s"
			readme   = true
		}
	`, identifier, description,
	)
}

func testAccResourceRepo(identifier, description string) string {
	return fmt.Sprintf(`
		resource "harness_platform_repo" "test" {
			identifier  = "%[1]s"
			default_branch = "master"
			description = "%[2]s"
		}
	`, identifier, description,
	)
}

func testAccResourceRepoImport(identifier, description, providerRepo string) string {
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
			identifier  = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			default_branch = "master"
			description = "%[2]s"
			
			source {
				type = "github"
				repo = "%[3]s"
			}
		}
	`, identifier, description, providerRepo,
	)
}

func testFindRepo(
	resourceName string,
	state *terraform.State,
) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	identifier := r.Primary.Attributes["identifier"]

	repo, _, err := c.RepositoryApi.FindRepository(
		ctx, c.AccountId, identifier,
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
			ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
		})
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

func testRepoDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repo, _ := testFindRepo(resourceName, state)
		if repo != nil {
			return fmt.Errorf("Found repo: %s", repo.Identifier)
		}

		return nil
	}
}
