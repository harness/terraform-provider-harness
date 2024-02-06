package repo_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProject(t *testing.T) {
	size := int64(1024)
	sizeUpdated := time.Now().Unix()
	updated := sizeUpdated
	resourceName := "data.code_repo.test"
	path := t.Name()

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepo(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "path", path),
					resource.TestCheckResourceAttr(resourceName, "size", strconv.FormatInt(size, 10)),
					resource.TestCheckResourceAttr(resourceName, "size_updated", strconv.FormatInt(sizeUpdated, 10)),
					resource.TestCheckResourceAttr(resourceName, "updated", strconv.FormatInt(updated, 10)),
				),
			},
		},
	})
}
