package app_project_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsAppProjectMapping(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_app_project_mapping.test"
	argoProject := "test123"
	argoProjectUpdated := "test123Updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		// CheckDestroy:      testAccResourceGitopsAppProjectMappingDestroy(resourceName, agentId), //commenting this since app project mapping cannot exist without an agent.
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAppProjectMapping(id, accountId, argoProject),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "argo_project_name", argoProject),
				),
			},
			{
				Config: testAccResourceGitopsAppProjectMapping(id, accountId, argoProjectUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "argo_project_name", argoProjectUpdated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetAppProjectMapping(resourceName string, state *terraform.State, agentId string) (*nextgen.V1AppProjectMappingV2, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	projectMappingId := r.Primary.Attributes["identifier"]

	resp, _, err := c.ProjectMappingsApi.AppProjectMappingServiceGetAppProjectMappingV2(ctx, agentId, projectMappingId, &nextgen.ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingV2Opts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
func testAccResourceGitopsAppProjectMappingDestroy(resourceName string, agentId string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		appProjectMapping, _ := testAccGetAppProjectMapping(resourceName, state, agentId)
		if appProjectMapping != nil {
			return fmt.Errorf("found app project mapping: %s", appProjectMapping.Identifier)
		}
		return nil
	}
}

func testAccResourceGitopsAppProjectMapping(agentId string, accountId string, argoProject string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			operator = "ARGO"
			metadata {
				namespace = "%[1]s"
        		high_availability = false
    		}
		}
		resource "harness_platform_gitops_app_project_mapping" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			agent_id = "%[1]s"
			argo_project_name = "%[3]s"
		}
		`, agentId, accountId, argoProject)
}
