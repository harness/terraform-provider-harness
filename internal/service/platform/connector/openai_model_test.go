package connector_test

/*
--------------------------------------------------------------------------
# Account-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export HARNESS_ENDPOINT='https://app.harness.io/gateway'
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="account.<account-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="account.<account-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    # Optional
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccResourceConnectorOpenAIModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Org-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export OPENAI_MODEL_ORG_SCOPE_ORG_ID="<org-id>"
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="org.<org-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="org.<org-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    # Optional
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestOrgResourceConnectorOpenAIModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Project-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export OPENAI_MODEL_PROJECT_SCOPE_ORG_ID="<org-id>"
    export OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID="<project-id>"
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="<project-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="<project-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    # Optional
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestProjectResourceConnectorOpenAIModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

*/

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// ---------------------------------------------------------------------------
// Helpers: read env vars, format delegate selectors for HCL
// ---------------------------------------------------------------------------

func openaiFormatDelegateSelectorsHCL(csv string) string {
	if csv == "" {
		return ""
	}
	parts := strings.Split(csv, ",")
	quoted := make([]string, len(parts))
	for i, p := range parts {
		quoted[i] = fmt.Sprintf(`"%s"`, strings.TrimSpace(p))
	}
	return strings.Join(quoted, ", ")
}

func openaiDelegateSelectorCount(csv string) int {
	if csv == "" {
		return 0
	}
	return len(strings.Split(csv, ","))
}

// ============================================================================
// Unit Tests — CustomizeDiff auth_type / block mismatch validation
// ============================================================================

func TestResourceConnectorOpenAIModel_TokenAuthWithVertexBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "harness_platform_connector_openai_model" "test" {
  identifier = "test_mismatch_token"
  name       = "test_mismatch_token"
  url        = "https://api.openai.com"
  model      = "gpt-4o"

  auth {
    auth_type = "Token"
    vertex {
      service_account_key_ref = "account.gcp_sa_key"
      project_id              = "my-gcp-project"
      region                  = "us-central1"
    }
  }
}`,
				ExpectError: regexp.MustCompile(`auth_type is "Token" but no token block provided`),
			},
		},
	})
}

func TestResourceConnectorOpenAIModel_VertexAuthWithTokenBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "harness_platform_connector_openai_model" "test" {
  identifier = "test_mismatch_vertex"
  name       = "test_mismatch_vertex"
  url        = "https://api.openai.com"
  model      = "gpt-4o"

  auth {
    auth_type = "Vertex"
    token {
      token_ref = "account.openai_api_key"
    }
  }
}`,
				ExpectError: regexp.MustCompile(`auth_type is "Vertex" but no vertex block provided`),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Account-scope (env-var driven, skip if not set)
// ============================================================================

func TestAccResourceConnectorOpenAIModel_Token(t *testing.T) {
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_token(id, name, tokenRef, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.openai.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "gpt-4o"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.token.0.token_ref", tokenRef),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", openaiDelegateSelectorCount(delegateCSV))),
				),
			},
			{
				Config: testAccResourceConnectorOpenAIModel_token(id, updatedName, tokenRef, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceConnectorOpenAIModel_Vertex(t *testing.T) {
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	if saKeyRef == "" || projectID == "" || region == "" {
		t.Skip("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF, OPENAI_MODEL_VERTEX_PROJECT_ID, and OPENAI_MODEL_VERTEX_REGION not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_vertex(id, name, saKeyRef, projectID, region, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.vertex.0.service_account_key_ref", saKeyRef),
					resource.TestCheckResourceAttr(resourceName, "auth.0.vertex.0.project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.vertex.0.region", region),
				),
			},
			{
				Config: testAccResourceConnectorOpenAIModel_vertex(id, updatedName, saKeyRef, projectID, region, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Org-scope
// ============================================================================

func TestOrgResourceConnectorOpenAIModel_Token(t *testing.T) {
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	orgID := os.Getenv("OPENAI_MODEL_ORG_SCOPE_ORG_ID")
	if tokenRef == "" || orgID == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF and OPENAI_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_tokenOrg(id, name, tokenRef, orgID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestOrgResourceConnectorOpenAIModel_Vertex(t *testing.T) {
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	orgID := os.Getenv("OPENAI_MODEL_ORG_SCOPE_ORG_ID")
	if saKeyRef == "" || projectID == "" || region == "" || orgID == "" {
		t.Skip("OPENAI_MODEL_VERTEX_* and OPENAI_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_vertexOrg(id, name, saKeyRef, projectID, region, orgID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Project-scope
// ============================================================================

func TestProjectResourceConnectorOpenAIModel_Token(t *testing.T) {
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	orgID := os.Getenv("OPENAI_MODEL_PROJECT_SCOPE_ORG_ID")
	projectID := os.Getenv("OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID")
	if tokenRef == "" || orgID == "" || projectID == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF, OPENAI_MODEL_PROJECT_SCOPE_ORG_ID, and OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_tokenProject(id, name, tokenRef, orgID, projectID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
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

func TestProjectResourceConnectorOpenAIModel_Vertex(t *testing.T) {
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	gcpProjectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	orgID := os.Getenv("OPENAI_MODEL_PROJECT_SCOPE_ORG_ID")
	projectID := os.Getenv("OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID")
	if saKeyRef == "" || gcpProjectID == "" || region == "" || orgID == "" || projectID == "" {
		t.Skip("OPENAI_MODEL_VERTEX_* and OPENAI_MODEL_PROJECT_SCOPE_* not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorOpenAIModel_vertexProject(id, name, saKeyRef, gcpProjectID, region, orgID, projectID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
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

// ============================================================================
// HCL config generators — Acceptance Tests
// ============================================================================

func testAccResourceConnectorOpenAIModel_token(id, name, tokenRef, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Token"
    token {
      token_ref = "%s"
    }
  }
}`, id, name, delegateBlock, tokenRef)
}

func testAccResourceConnectorOpenAIModel_vertex(id, name, saKeyRef, gcpProjectID, region, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Vertex"
    vertex {
      service_account_key_ref = "%s"
      project_id              = "%s"
      region                  = "%s"
    }
  }
}`, id, name, delegateBlock, saKeyRef, gcpProjectID, region)
}

func testAccResourceConnectorOpenAIModel_tokenOrg(id, name, tokenRef, orgID, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  org_id                 = "%s"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Token"
    token {
      token_ref = "%s"
    }
  }
}`, id, name, orgID, delegateBlock, tokenRef)
}

func testAccResourceConnectorOpenAIModel_vertexOrg(id, name, saKeyRef, gcpProjectID, region, orgID, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  org_id                 = "%s"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Vertex"
    vertex {
      service_account_key_ref = "%s"
      project_id              = "%s"
      region                  = "%s"
    }
  }
}`, id, name, orgID, delegateBlock, saKeyRef, gcpProjectID, region)
}

func testAccResourceConnectorOpenAIModel_tokenProject(id, name, tokenRef, orgID, projectID, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  org_id                 = "%s"
  project_id             = "%s"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Token"
    token {
      token_ref = "%s"
    }
  }
}`, id, name, orgID, projectID, delegateBlock, tokenRef)
}

func testAccResourceConnectorOpenAIModel_vertexProject(id, name, saKeyRef, gcpProjectID, region, orgID, projectID, delegateHCL string) string {
	delegateBlock := ""
	if delegateHCL != "" {
		delegateBlock = fmt.Sprintf(`delegate_selectors  = [%s]`, delegateHCL)
	}
	return fmt.Sprintf(`
resource "harness_platform_connector_openai_model" "test" {
  identifier             = "%s"
  name                   = "%s"
  description            = "test"
  tags                   = ["foo:bar"]
  org_id                 = "%s"
  project_id             = "%s"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  %s

  auth {
    auth_type = "Vertex"
    vertex {
      service_account_key_ref = "%s"
      project_id              = "%s"
      region                  = "%s"
    }
  }
}`, id, name, orgID, projectID, delegateBlock, saKeyRef, gcpProjectID, region)
}

