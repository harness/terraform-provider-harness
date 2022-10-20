package repository

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsRepository(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	repo := id
	repoName := id
	agentId := "account.terraformagent1"
	resourceName := "harness_platform_gitops_repository.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepository(id, name, repo, repoName, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func testAccGetRepository(resourceName string, state *terraform.State) (*nextgen.Servicev1Repository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_id"]
	identifier := r.Primary.Attributes["identifier"]
	resp, _, err := c.RepositoriesApiService.AgentRepositoryServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoriesApiAgentRepositoryServiceGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
		QueryRepo:         optional.NewString(r.Primary.Attributes["query_repo"]),
		QueryForceRefresh: optional.NewBool(r.Primary.Attributes["query_force_refresh"] == "True"),
		QueryProject:      optional.NewString(r.Primary.Attributes["query_project"]),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsRepositoryDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repo, _ := testAccGetRepository(resourceName, state)
		if repo != nil {
			return fmt.Errorf("Found Repository: %s", repo.Identifier)
		}
		return nil
	}
}

func testAccResourceGitopsRepository(id string, name string, repo string, repoName string, agentId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[5]s"
			"repo": {
					"repo": "%[3]s",
        			"name": ""%[4]s",
        			"insecure": true,
        			"connectionType": "HTTPS"
			},
    		"upsert": true
		}
	`, id, name, repo, repoName, agentId)
}
