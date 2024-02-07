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

var accountId = utils.GetEnv("HARNESS_ACCOUNT_ID", "")

func TestAccResourceRepo(t *testing.T) {
	resourceName := "harness_platform_repo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(),
				// Check:  resource.ComposeTestCheckFunc(
				// // resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
				// ),
			},
			{
				Config: testAccResourceRepo(),
				// Check:  resource.ComposeTestCheckFunc(
				// // resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(updatedSize, 10)),
				// ),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.UserResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceRepo_DeleteUnderlyingResource(t *testing.T) {
	path := t.Name()

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(),
				Check:  resource.ComposeTestCheckFunc(
				// resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.DeleteRepository(
						ctx, accountId, path, &code.RepositoryApiDeleteRepositoryOpts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceRepo(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// resource "harness_platform_organization" "test" {
// 	identifier = "default"
// 	name = "default"
// 	description = "test"
// 	tags = ["foo:bar", "baz:qux"]
// }

// resource "harness_platform_project" "test" {
// 	identifier = "default_project"
// 	name = "default_project"
// 	org_id = harness_platform_organization.test.id
// }

func testAccResourceRepo() string {
	accountId := utils.GetEnv("HARNESS_ACCOUNT_ID", "")
	return fmt.Sprintf(`

		resource "harness_platform_repo" "test" {
			identifier  = "example_identifier"
			name       	= "example_name"
			description = "example_description"
			account_id 	= "%[1]s"
			org_identifier = "default"
			project_identifier = "default_project"
		}
	`, accountId,
	)
}

func testAccFindRepo(
	resourceName string,
	state *terraform.State,
) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	path := r.Primary.Attributes["path"]

	repo, _, err := c.RepositoryApi.FindRepository(
		ctx, accountId, path,
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
