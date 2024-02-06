package repo_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

const createdBy = 1

func TestAccResourceRepo(t *testing.T) {
	size := int64(1024)
	updatedSize := size * 2
	sizeUpdated := time.Now().Unix()
	updatedSizeUpdated := sizeUpdated + 3
	updated := sizeUpdated
	updatedUpdated := updatedSizeUpdated
	resourceName := "harness_platform_repo.test"
	path := t.Name()

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(sizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updated, 10)),
				),
			},
			{
				Config: testAccResourceRepo(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(updatedSize, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(updatedSizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updatedUpdated, 10)),
				),
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
	t.Skip()
	size := int64(1024)
	sizeUpdated := time.Now().Unix()
	updated := sizeUpdated
	resourceName := "harness_platform_repo.test"
	path := t.Name()

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(sizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updated, 10)),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.DeleteRepository(
						ctx, strconv.Itoa(createdBy), path, &code.RepositoryApiDeleteRepositoryOpts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceRepo(path),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceRepo(path string) string {
	return fmt.Sprintf(`
		resource "harness_platform_repo" "test" {
			identifier = "example_identifier"
			name       = "example_name"
			path = "%[1]s"
			created_by = %[2]d
		}
	`, path, createdBy,
	)
}

func testAccFindRepo(resourceName string, state *terraform.State) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	path := r.Primary.Attributes["path"]

	repo, _, err := c.RepositoryApi.FindRepository(
		ctx, strconv.Itoa(createdBy), path, &code.RepositoryApiFindRepositoryOpts{})
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
