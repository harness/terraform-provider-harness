package yamlconfig_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceYamlConfig_CloudProvider(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", t.Name())
	resourceName := "harness_yaml_config.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccYamlConfigCloudProvider(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
				),
			},
			{
				Config: testAccYamlConfigCloudProvider(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["path"], nil
				},
			},
		},
	})
}

func testAccYamlConfigCloudProvider(name string) string {
	return fmt.Sprintf(`
resource "harness_yaml_config" "test" {
	path = "Setup/Cloud Providers/%[1]s.yaml"
	content = <<EOF
harnessApiVersion: '1.0'
type: KUBERNETES_CLUSTER
continuousEfficiencyConfig:
  continuousEfficiencyEnabled: false
delegateSelectors:
- k8s
skipValidation: true
useKubernetesDelegate: true
EOF
}
`, name)
}
