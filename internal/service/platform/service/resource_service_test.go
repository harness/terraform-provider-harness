package service_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceService(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceServiceWithYaml(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	varValue := t.Name()
	updatedVarValue := fmt.Sprintf("%s_updated", varValue)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceServiceWithYaml(id, name, varValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceServiceWithYaml(id, updatedName, updatedVarValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceService_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.ServicesApi.DeleteServiceV2(ctx, id, c.AccountId, &nextgen.ServicesApiDeleteServiceV2Opts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testAccResourceService(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetService(resourceName string, state *terraform.State) (*nextgen.ServiceResponseDetails, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.ServicesApi.GetServiceV2(ctx, id, c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil || resp.Data.Service == nil {
		return nil, nil
	}

	return resp.Data.Service, nil
}

func testAccServiceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetService(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found service: %s", env.Identifier)
		}

		return nil
	}
}

func testAccResourceService(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_organization" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
    }

    resource "harness_platform_project" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_organization.test.id
      color = "#472848"
  }

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
    }
`, id, name)
}

func testAccResourceServiceWithYaml(id string, name string, varValue string) string {
	return fmt.Sprintf(`
    resource "harness_platform_organization" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
    }

    resource "harness_platform_project" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_organization.test.id
      color = "#472848"
    }

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
              configFiles:
                - configFile:
                    identifier: configFile1
                    spec:
                      store:
                        type: Harness
                        spec:
                          files:
                            - <+org.description>
              variables:
                - name: var1
                  type: String
                  value: %[3]s
                - name: var2
                  type: String
                  value: val2
            type: Kubernetes
          gitOpsEnabled: false
      EOT
    }
`, id, name, varValue)
}
