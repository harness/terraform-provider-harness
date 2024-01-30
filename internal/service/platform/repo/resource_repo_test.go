package repo_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRepo(t *testing.T) {
	id := int32(123)
	name := t.Name()
	uid := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	size := int64(1024)
	updatedSize := size * 2
	sizeUpdated := time.Now().Unix()
	updatedsizeUpdated := sizeUpdated + 3
	updated := sizeUpdated
	updatedUpdated := updatedsizeUpdated
	resourceName := "harness_code_repo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRepoDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(id, uid, size, sizeUpdated, updated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", string(id)),
					resource.TestCheckResourceAttr(resourceName, "uid", uid),
					resource.TestCheckResourceAttr(resourceName, "size", string(size)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", string(sizeUpdated)),
					resource.TestCheckResourceAttr(resourceName, "updated", string(updated)),
				),
			},
			{
				Config: testAccResourceRepo(id, uid, updatedSize, updatedsizeUpdated, updatedUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", string(id)),
					resource.TestCheckResourceAttr(resourceName, "uid", uid),
					resource.TestCheckResourceAttr(resourceName, "size", string(updatedSize)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", string(updatedsizeUpdated)),
					resource.TestCheckResourceAttr(resourceName, "updated", string(updatedUpdated)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceRepo(id int32, uid string, size, sizeUpdated, updated int64) string {
	return fmt.Sprintf(`
		resource "harness_code_repo" "test" {
			id = "%[1]s"
			uid = "%[2]s"
			path = "example/path"
			git_url = "https://github.com/example/repo.git"
			is_public = true
			importing = false
			created_by = %[3]d
			created = %[4]d
			updated = %[5]d
			size = %[6]d
			size_updated = %[7]d
			parent_id = %[8]d
			fork_id = %[9]d
			pull_req_seq = %[10]d
			num_forks = %[11]d
			num_pulls = %[12]d
			num_closed_pulls = %[13]d
			num_open_pulls = %[14]d
			num_merged_pulls = %[15]d
		}
	`, id, uid, rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), size, sizeUpdated,
		rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
		rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
	)
}

func testAccFindRepo(resourceName string, state *terraform.State) (*code.TypesRepository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	id := r.Primary.Attributes["id"]

	repo, _, err := c.RepositoryApi.FindRepository(ctx, id)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

func testAccRepoDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repo, _ := testAccFindRepo(resourceName, state)
		if repo != nil {
			return fmt.Errorf("Found repo: %s", repo.Id)
		}

		return nil
	}
}
