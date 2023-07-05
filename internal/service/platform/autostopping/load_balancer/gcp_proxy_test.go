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
	hostName := fmt.Sprintf("ab%s.com", name)
	resourceName := "harness_autostopping_gcp_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testGCPProxy(name, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
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

func testGCPProxy(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_gcp_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "developerxgcpfm"
			host_name = "%[2]s"
            region = "us-central1"
			vpc = "https://www.googleapis.com/compute/v1/projects/gcpccmdemo/global/networks/default"
			zone = "us-central1-a"
			security_groups =["http-server"]
			machine_type = "e2-micro"
			subnet_id = "https://www.googleapis.com/compute/v1/projects/gcpccmdemo/regions/us-central1/subnetworks/default"
            api_key = ""
			allocate_static_ip = false
			certificates {
				key_secret_id = "projects/643752427051/secrets/secret-ccmpriv/versions/1"
				cert_secret_id = "projects/643752427051/secrets/secret-ccmpub/versions/1"
			  }  
		}
`, name, hostName)
}

func testGCPProxyUpdate(name string, hostName string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_gcp_proxy" "test" {
		name = "%[1]s"
		cloud_connector_id = "developerxgcpfm"
		host_name = "%[2]s"
		region = "us-central1"
		vpc = "https://www.googleapis.com/compute/v1/projects/gcpccmdemo/global/networks/default"
		zone = "us-central1-a"
		security_groups =["http-server","https-server"]
		machine_type = "e2-micro"
		subnet_id = "https://www.googleapis.com/compute/v1/projects/gcpccmdemo/regions/us-central1/subnetworks/default"
		api_key = ""
		allocate_static_ip = false
		certificates {
			key_secret_id = "projects/643752427051/secrets/secret-ccmpriv/versions/1"
			cert_secret_id = "projects/643752427051/secrets/secret-ccmpub/versions/1"
		  }  
	}
`, name, hostName)
}
