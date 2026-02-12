package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceGCPProxy(t *testing.T) {
	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_gcp_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGCPProxy(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testGCPProxyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		proxy, _ := testGetLoadBalancer(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found aws proxy: %s", proxy.Id)
		}
		return nil
	}
}

func testGCPProxy(name string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_gcp_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "developerxgcpfm"
            region             = "region"
			vpc                = "https://www.googleapis.com/compute/v1/projects/project_id/global/networks/netwok_id"
			zone               = "zone"
			security_groups    = ["http-server"]
			machine_type       = "e2-micro"
			subnet_id          = "https://www.googleapis.com/compute/v1/projects/project_id/regions/region/subnetworks/subnet_name"
			api_key            = ""
			allocate_static_ip = false
			certificates {
				key_secret_id  = "projects/project_id/secrets/secret_id/versions/1"
				cert_secret_id = "projects/project_id/secrets/secret_id/versions/1"
			}
		}
`, name)
}

func testGCPProxyUpdate(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_gcp_proxy" "test" {
		name = "%[1]s"
		cloud_connector_id = "developerxgcpfm"
		region             = "region"
		vpc                = "https://www.googleapis.com/compute/v1/projects/project_id/global/networks/netwok_id"
		zone               = "zone"
		security_groups    = ["http-server","https-server"]
		machine_type       = "e2-micro"
		subnet_id          = "https://www.googleapis.com/compute/v1/projects/project_id/regions/region/subnetworks/subnet_name"
		api_key            = ""
		allocate_static_ip = false
		certificates {
			key_secret_id  = "projects/project_id/secrets/secret_id/versions/1"
			cert_secret_id = "projects/project_id/secrets/secret_id/versions/1"
		}
		delete_cloud_resources_on_destroy = true
	}
`, name)
}
