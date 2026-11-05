package idp_test

// Acceptance tests for harness_platform_idp_integration, one per integration type.
//
// Each test creates an integration, fails if the plan after the apply isn't empty, then imports it and fails if the
// imported state differs. Both catch fields the server fills in that the schema doesn't expect.
//
// They require TF_ACC=1, HARNESS_ENDPOINT, HARNESS_ACCOUNT_ID and HARNESS_PLATFORM_API_KEY. Types that read with a
// connector or secret also need the HARNESS_TEST_IDP_* variables named in their test, and are skipped without them:
//
//   HARNESS_TEST_IDP_SECRET_REF              - any secret, e.g. account.my_secret (Custom webhook mode)
//   HARNESS_TEST_IDP_K8S_AGENT_REF           - service discovery agent, e.g. account.my_agent
//   HARNESS_TEST_IDP_K8S_ENVIRONMENT         - environment the agent runs in
//   HARNESS_TEST_IDP_SERVICENOW_CONNECTOR    - ServiceNow connector
//   HARNESS_TEST_IDP_PAGERDUTY_CONNECTOR     - PagerDuty connector
//   HARNESS_TEST_IDP_DATADOG_CONNECTOR       - DataDog connector
//   HARNESS_TEST_IDP_GITHUB_CONNECTOR        - GitHub connector
//   HARNESS_TEST_IDP_GITHUB_ORG              - GitHub organization the connector can read
//   HARNESS_TEST_IDP_JIRA_CONNECTOR          - Jira connector
//   HARNESS_TEST_IDP_DYNATRACE_CONNECTOR     - DynaTrace connector
//   HARNESS_TEST_IDP_DYNATRACE_TOKEN_REF     - secret with a DynaTrace platform token
//   HARNESS_TEST_IDP_GCP_CONNECTOR           - GCP connector
//   HARNESS_TEST_IDP_GCP_FOLDER_ID           - GCP folder to sync
//   HARNESS_TEST_IDP_BITBUCKET_CONNECTOR     - Bitbucket Cloud connector
//   HARNESS_TEST_IDP_CATALOG_INFO_CONNECTOR  - GitHub connector for CatalogInfo
//   HARNESS_TEST_IDP_CATALOG_INFO_REPO       - repository with catalog-info.yaml files
//   HARNESS_TEST_IDP_SONARQUBE_TOKEN_REF     - secret with a SonarQube Cloud token
//   HARNESS_TEST_IDP_SONARQUBE_ORG           - SonarQube Cloud organization key
//
// Singleton types, such as HarnessCD, aren't tested, as an account can hold only one of each.
//
// Run them:
//   TF_ACC=1 go test -v ./internal/service/platform/idp/... -run TestAccResourceIdpIntegration -timeout=60m

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const integrationResourceName = "harness_platform_idp_integration.test"

func TestAccResourceIdpIntegrationCustomAPI(t *testing.T) {
	testAccIntegration(t, func(id string, _ map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "Custom"
  tags       = ["terraform"]

  custom {
    mode         = "api"
    source_label = "orders"
  }

  kind {
    kind   = "custom"
    action = "merge"

    mapping {
      ingestion {
        identifier = ".service.id"
        name       = ".service.name"
        properties = {
          team = ".service.team"
        }
      }
      correlation {
        source_path = "service.id"
        target_path = "metadata.identifier"
      }
    }
  }
}
`, id, id)
	})
}

func TestAccResourceIdpIntegrationCustomWebhook(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "Custom"

  custom {
    mode         = "webhook"
    source_label = "builds"

    webhook_auth {
      auth_type   = "hmac_auth"
      hash_algo   = "hmac_sha256"
      auth_header = "X-Signature"
      secret_ref  = "%s"
    }
  }

  kind {
    kind = "build"
  }
}
`, id, id, env["HARNESS_TEST_IDP_SECRET_REF"])
	}, "HARNESS_TEST_IDP_SECRET_REF")
}

func TestAccResourceIdpIntegrationHarnessK8s(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "HarnessK8s"

  harness_k8s {
    service_discovery_agent_identifier = "%s"
    environment_identifier             = "%s"
  }

  kind {
    kind = "service"
  }
}
`, id, id, env["HARNESS_TEST_IDP_K8S_AGENT_REF"], env["HARNESS_TEST_IDP_K8S_ENVIRONMENT"])
	}, "HARNESS_TEST_IDP_K8S_AGENT_REF", "HARNESS_TEST_IDP_K8S_ENVIRONMENT")
}

func TestAccResourceIdpIntegrationServiceNow(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier       = "%s"
  name             = "%s"
  type             = "ServiceNow"
  update_frequency = "12h"
%s
  kind {
    kind     = "cmdb_ci_win_server"
    idp_kind = "component"
    idp_type = "service"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_SERVICENOW_CONNECTOR"]))
	}, "HARNESS_TEST_IDP_SERVICENOW_CONNECTOR")
}

func TestAccResourceIdpIntegrationPagerDuty(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "PagerDuty"
%s
  kind {
    kind = "services"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_PAGERDUTY_CONNECTOR"]))
	}, "HARNESS_TEST_IDP_PAGERDUTY_CONNECTOR")
}

func TestAccResourceIdpIntegrationDataDog(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "DataDog"
%s
  kind {
    kind = "services"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_DATADOG_CONNECTOR"]))
	}, "HARNESS_TEST_IDP_DATADOG_CONNECTOR")
}

func TestAccResourceIdpIntegrationGitHub(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "GitHub"
%s
  github {
    orgs = ["%s"]
  }

  kind {
    kind = "repo_details"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_GITHUB_CONNECTOR"]), env["HARNESS_TEST_IDP_GITHUB_ORG"])
	}, "HARNESS_TEST_IDP_GITHUB_CONNECTOR", "HARNESS_TEST_IDP_GITHUB_ORG")
}

func TestAccResourceIdpIntegrationJira(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "Jira"
%s
  kind {
    kind = "components"
  }
  kind {
    kind = "projects"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_JIRA_CONNECTOR"]))
	}, "HARNESS_TEST_IDP_JIRA_CONNECTOR")
}

func TestAccResourceIdpIntegrationDynaTrace(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "DynaTrace"
%s
  dynatrace {
    platform_token_ref = "%s"
  }

  kind {
    kind = "services"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_DYNATRACE_CONNECTOR"]),
			env["HARNESS_TEST_IDP_DYNATRACE_TOKEN_REF"])
	}, "HARNESS_TEST_IDP_DYNATRACE_CONNECTOR", "HARNESS_TEST_IDP_DYNATRACE_TOKEN_REF")
}

func TestAccResourceIdpIntegrationGCP(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "GCP"
%s
  gcp {
    parent {
      folder_id = "%s"
    }
  }

  kind {
    kind = "compute_instances"
  }
  kind {
    kind = "projects"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_GCP_CONNECTOR"]), env["HARNESS_TEST_IDP_GCP_FOLDER_ID"])
	}, "HARNESS_TEST_IDP_GCP_CONNECTOR", "HARNESS_TEST_IDP_GCP_FOLDER_ID")
}

func TestAccResourceIdpIntegrationBitbucketCloud(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "BitbucketCloud"
%s
  kind {
    kind = "repo_details"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_BITBUCKET_CONNECTOR"]))
	}, "HARNESS_TEST_IDP_BITBUCKET_CONNECTOR")
}

func TestAccResourceIdpIntegrationCatalogInfo(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "CatalogInfo"
%s
  catalog_info {
    source_type  = "github"
    repositories = ["%s"]
  }

  kind {
    kind = "catalog_info_file"
  }
}
`, id, id, testAccIntegrationConnector(env["HARNESS_TEST_IDP_CATALOG_INFO_CONNECTOR"]),
			env["HARNESS_TEST_IDP_CATALOG_INFO_REPO"])
	}, "HARNESS_TEST_IDP_CATALOG_INFO_CONNECTOR", "HARNESS_TEST_IDP_CATALOG_INFO_REPO")
}

func TestAccResourceIdpIntegrationSonarQube(t *testing.T) {
	testAccIntegration(t, func(id string, env map[string]string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "SonarQube"

  sonarqube {
    token_ref = "%s"
    instance {
      cloud {
        organization = "%s"
      }
    }
  }

  kind {
    kind = "projects"
  }
}
`, id, id, env["HARNESS_TEST_IDP_SONARQUBE_TOKEN_REF"], env["HARNESS_TEST_IDP_SONARQUBE_ORG"])
	}, "HARNESS_TEST_IDP_SONARQUBE_TOKEN_REF", "HARNESS_TEST_IDP_SONARQUBE_ORG")
}

// TestAccResourceIdpIntegrationKindMapping adopts the mapping of a kind created without one, changes it, then imports it.
func TestAccResourceIdpIntegrationKindMapping(t *testing.T) {
	id := fmt.Sprintf("tf_acc_%s", strings.ToLower(utils.RandStringBytes(8)))
	const mapping = "harness_platform_idp_integration_kind_mapping.test"
	config := func(team string) string {
		return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%s"
  name       = "%s"
  type       = "Custom"

  custom {
    mode         = "api"
    source_label = "orders"
  }

  kind {
    kind = "custom"
  }
}

resource "harness_platform_idp_integration_kind_mapping" "test" {
  integration_id = harness_platform_idp_integration.test.identifier
  kind           = "custom"

  ingestion {
    identifier = ".service.id"
    name       = ".service.name"
    properties = {
      team = "%s"
    }
  }
}
`, id, id, team)
	}
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config(".service.team"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(mapping, "id", id+"/custom"),
					resource.TestCheckResourceAttr(mapping, "ingestion.0.properties.team", ".service.team"),
				),
			},
			{
				Config: config(".service.owner"),
				Check:  resource.TestCheckResourceAttr(mapping, "ingestion.0.properties.team", ".service.owner"),
			},
			{
				ResourceName:            mapping,
				ImportState:             true,
				ImportStateId:           id + "/custom",
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated"},
			},
		},
	})
}

// testAccIntegration applies config, which fails on a non-empty plan after the apply, renames the integration,
// then verifies an import.
func testAccIntegration(t *testing.T, config func(id string, env map[string]string) string, envVars ...string) {
	env := map[string]string{}
	for _, name := range envVars {
		if env[name] = os.Getenv(name); env[name] == "" {
			t.Skipf("%s must be set for this acceptance test", name)
		}
	}
	id := fmt.Sprintf("tf_acc_%s", strings.ToLower(utils.RandStringBytes(8)))
	// Every config sets the name to the identifier.
	namePattern := regexp.MustCompile(`(?m)^(\s*name\s*=\s*)"` + regexp.QuoteMeta(id) + `"$`)
	renamed := namePattern.ReplaceAllString(config(id, env), `${1}"`+id+`_updated"`)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config(id, env),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(integrationResourceName, "id", id),
					resource.TestCheckResourceAttr(integrationResourceName, "revision", "1"),
				),
			},
			{
				Config: renamed,
				Check:  resource.TestCheckResourceAttr(integrationResourceName, "name", id+"_updated"),
			},
			{
				ResourceName:      integrationResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// Platform types are touched by the server right after the create. Import doesn't read the
				// create-only mapping blocks.
				ImportStateVerifyIgnore: []string{"updated", "kind.0.mapping", "kind.1.mapping"},
			},
		},
	})
}

func testAccIntegrationConnector(ref string) string {
	return fmt.Sprintf(`
  connector {
    inline {
      credentials_ref = "%s"
    }
  }
`, ref)
}

// testAccIntegrationDestroy checks the integrations API no longer returns the integration.
func testAccIntegrationDestroy(state *terraform.State) error {
	s := acctest.TestAccGetApiClientFromProvider()
	for _, r := range state.RootModule().Resources {
		if r.Type != "harness_platform_idp_integration" {
			continue
		}
		url := fmt.Sprintf("%s/integration/api/v1/accounts/%s/integrations/%s",
			strings.TrimRight(s.Endpoint, "/"), s.AccountId, r.Primary.ID)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		req.Header.Set("x-api-key", s.IDPClient.ApiKey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("integration %s still exists: GET returned %d", r.Primary.ID, resp.StatusCode)
		}
	}
	return nil
}
