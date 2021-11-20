package connector_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceConnector_WithOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_WithOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_WithOrg(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.Attributes["id"]
					org_id := primary.Attributes["org_id"]
					return fmt.Sprintf("%s/%s", org_id, id), nil
				},
			},
		},
	})
}

func TestAccResourceConnector_WithProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_WithProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_WithProject(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.Attributes["id"]
					org_id := primary.Attributes["org_id"]
					project_id := primary.Attributes["project_id"]
					return fmt.Sprintf("%s/%s/%s", org_id, project_id, id), nil
				},
			},
		},
	})
}

// func TestAccResourceConnector_ChangeType(t *testing.T) {

// 	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
// 	name := id
// 	resourceName := "harness_connector.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		CheckDestroy:      testAccConnectorDestroy(resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceConnector_k8s_ServiceAccount(id, name),
// 				// Check: resource.ComposeTestCheckFunc(
// 				// 	resource.TestCheckResourceAttr(resourceName, "id", id),
// 				// 	resource.TestCheckResourceAttr(resourceName, "identifier", id),
// 				// 	resource.TestCheckResourceAttr(resourceName, "org_id", id),
// 				// 	resource.TestCheckResourceAttr(resourceName, "project_id", id),
// 				// 	resource.TestCheckResourceAttr(resourceName, "name", name),
// 				// 	resource.TestCheckResourceAttr(resourceName, "description", "test"),
// 				// 	resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
// 				// ),
// 			},
// 			{
// 				Config: testAccResourceConnector_docker_DockerHub(id, name),
// 				// PlanOnly:           true,
// 				// ExpectNonEmptyPlan: true,
// 				Check: func(*terraform.State) error {
// 					return nil
// 				},
// 				// 	Check: resource.ComposeTestCheckFunc(
// 				// 		resource.TestCheckResourceAttr(resourceName, "id", id),
// 				// 		resource.TestCheckResourceAttr(resourceName, "identifier", id),
// 				// 		resource.TestCheckResourceAttr(resourceName, "org_id", id),
// 				// 		resource.TestCheckResourceAttr(resourceName, "project_id", id),
// 				// 		resource.TestCheckResourceAttr(resourceName, "name", updatedName),
// 				// 		resource.TestCheckResourceAttr(resourceName, "description", "test"),
// 				// 		resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
// 				// 	),
// 			},
// 		},
// 	})
// }

func TestAccResourceConnector_DeleteUnderlyingResource(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_k8s_InheritFromDelegate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "k8s_cluster.0.inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*api.Client)
					_, _, err := c.NGClient.ConnectorsApi.DeleteConnector(context.Background(), c.AccountId, id, nil)
					require.NoError(t, err)
				},
				Config:             testAccResourceConnector_k8s_InheritFromDelegate(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccConnectorDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		connector, _ := testAccGetConnector(resourceName, state)
		if connector != nil {
			return fmt.Errorf("Found connector: %s", connector.Identifier)
		}

		return nil
	}
}

func testAccGetConnector(resourceName string, state *terraform.State) (*nextgen.ConnectorInfo, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID

	options := &nextgen.ConnectorsApiGetConnectorOpts{AccountIdentifier: optional.NewString(c.AccountId)}

	if attr := r.Primary.Attributes["org_id"]; attr != "" {
		options.OrgIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["project_id"]; attr != "" {
		options.ProjectIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["branch"]; attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["repo_id"]; attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	if attr := r.Primary.Attributes["repo_id"]; attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.NGClient.ConnectorsApi.GetConnector(context.Background(), id, options)
	if err != nil {
		return nil, err
	}

	return resp.Data.Connector, nil
}

func testAccResourceConnector_WithProject(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			color = "#FFFFFF"
			org_id = harness_organization.test.id
		}

		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_organization.test.id
			project_id = harness_project.test.id

			k8s_cluster {
				inherit_from_delegate {
					delegate_selectors = ["harness-delegate"]
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_WithOrg(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_organization.test.id

			k8s_cluster {
				inherit_from_delegate {
					delegate_selectors = ["harness-delegate"]
				}
			}
		}
`, id, name)
}
