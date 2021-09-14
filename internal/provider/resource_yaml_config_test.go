package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceYamlConfig_CloudProvider(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "harness_yaml_config.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccYamlConfigCloudProvider(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
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
delegateSelectors:
- k8s
skipValidation: true
useKubernetesDelegate: true
EOF
}
`, name)
}
