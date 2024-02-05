package repo_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceRepo(t *testing.T) {
	id := int64(123)
	name := t.Name()
	uid := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	size := int64(1024)
	updatedSize := size * 2
	sizeUpdated := time.Now().Unix()
	updatedSizeUpdated := sizeUpdated + 3
	updated := sizeUpdated
	updatedUpdated := updatedSizeUpdated
	resourceName := "harness_platform_repo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(id, uid, size, sizeUpdated, updated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", strconv.FormatInt(id, 10)),
					resource.TestCheckResourceAttr(resourceName, "uid", uid),
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(sizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updated, 10)),
				),
			},
			{
				Config: testAccResourceRepo(id, uid, updatedSize, updatedSizeUpdated, updatedUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", strconv.FormatInt(id, 10)),
					resource.TestCheckResourceAttr(resourceName, "uid", uid),
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
	id := int64(123)
	name := t.Name()
	uid := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	size := int64(1024)
	sizeUpdated := time.Now().Unix()
	updated := sizeUpdated
	resourceName := "harness_platform_repo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(id, uid, size, sizeUpdated, updated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", strconv.FormatInt(id, 10)),
					resource.TestCheckResourceAttr(resourceName, "uid", uid),
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(sizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updated, 10)),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.DeleteRepository(ctx, uid)
					require.NoError(t, err)
				},
				Config:             testAccResourceRepo(id, uid, size, sizeUpdated, updated),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceRepo(id int64, uid string, size, sizeUpdated, updated int64) string {
	return fmt.Sprintf(`
		resource "harness_platform_repo" "test" {
			id = %[1]d
			uid = "%[2]s"
			path = "example/path"
			git_url = "https://github.com/example/repo.git"
			is_public = true
			updated = %[3]d
			size = %[4]d
			size_updated = %[5]d
		}
	`, id, uid, updated, size, sizeUpdated,
	)
}

func testAccFindRepo(resourceName string, state *terraform.State) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	path := r.Primary.Attributes["path"]

	repo, _, err := c.RepositoryApi.FindRepository(ctx, path)
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
