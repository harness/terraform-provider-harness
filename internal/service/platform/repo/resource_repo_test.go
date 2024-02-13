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
	orgId        = "example_org_123"
	prjId        = "example_project_123"
	repoId       = "example_identifier"
	repoDesc     = "example_description"
	providerRepo = "octocat/hello-worId"
)

var accountId = utils.GetEnv("HARNESS_ACCOUNT_ID", "")

func TestAccResourceRepo(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(repoDesc, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", repoDesc),
				),
			},
			{
				Config: testAccResourceRepo(repoDesc, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", repoDesc),
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
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(repoDesc, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", repoDesc),
				),
			},
			{
				Config: testAccResourceRepo(repoDesc, providerRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", repoDesc),
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

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(repoDesc, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", repoDesc),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.DeleteRepository(
						ctx, accountId, repoId,
						&code.RepositoryApiDeleteRepositoryOpts{
							OrgIdentifier:     optional.NewString("default"),
							ProjectIdentifier: optional.NewString(prjId),
						})
					require.NoError(t, err)
				},
				Config:             testAccResourceRepo(repoDesc, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceRepo(description, providerRepo string) string {
	accountId := utils.GetEnv("HARNESS_ACCOUNT_ID", "")
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}	

		resource "harness_platform_project" "test" {
			identifier = "%[2]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}
		
		resource "harness_platform_repo" "test" {
			identifier  = "%[3]s"
			name       	= "%[3]s"
			org_identifier = harness_platform_organization.test.identifier
			project_identifier = harness_platform_project.test.identifier
			default_branch = "master"
			description = "%[4]s"
			account_id 	= "%[5]s"
			provider_repo = "%[6]s"
			is_public = true
			type = "github"
		}
	`, orgId, prjId, repoId, description, accountId, providerRepo,
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
			return fmt.Errorf("Found repo: %s", repo.Path)
		}

		return nil
	}
}
