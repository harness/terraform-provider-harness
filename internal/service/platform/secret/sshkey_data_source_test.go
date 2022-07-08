package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSSHKey(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_secret_sshkey.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_sshkey(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "principal"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "key_path"))},
		},
	})
}

func testAccDataSourceSecret_sshkey(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		port = 22
		kerberos {
			tgt_key_tab_file_path_spec {
				key_path = "key_path"
			}
			principal = "principal"
			realm = "realm"
			tgt_generation_method = "KeyTabFilePath"
		}
	}
	data "harness_platform_secret_sshkey" "test"{
		identifier = harness_platform_secret_sshkey.test.identifier
	}
	`, name)
}
