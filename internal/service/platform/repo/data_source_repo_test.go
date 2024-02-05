package repo_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProject(t *testing.T) {

	id := int64(123)
	name := t.Name()
	uid := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	size := int64(1024)
	sizeUpdated := time.Now().Unix()
	updated := sizeUpdated
	resourceName := "code_repo.test"

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
		},
	})
}
