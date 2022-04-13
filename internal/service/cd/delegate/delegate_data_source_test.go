package delegate_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/delegate"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDelegate_hostname(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	delegateName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_delegate.test"

	defer deleteDelegate(t, delegateName)

	acctest.TestAccPreCheck(t)
	pullDelegateImage(context.Background(), &delegate.DockerDelegateConfig{})
	delegate := createDelegateContainer(t, delegateName, false)

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegate_hostname(delegate.HostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", delegateName),
					resource.TestCheckResourceAttr(resourceName, "hostname", delegate.HostName),
				),
			},
		},
	})
}

func TestAccDataSourceDelegate_name(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	delegateName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_delegate.test"

	defer deleteDelegate(t, delegateName)

	acctest.TestAccPreCheck(t)
	pullDelegateImage(context.Background(), &delegate.DockerDelegateConfig{})
	delegate := createDelegateContainer(t, delegateName, false)

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegate_name(delegate.DelegateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", delegateName),
				),
			},
		},
	})
}

func testAccDataSourceDelegate_hostname(hostname string) string {
	return fmt.Sprintf(`
		data "harness_delegate" "test" {
			hostname = "%[1]s"
		}
	`, hostname)
}

func testAccDataSourceDelegate_name(name string) string {
	return fmt.Sprintf(`
		data "harness_delegate" "test" {
			name = "%[1]s"
		}
	`, name)
}
