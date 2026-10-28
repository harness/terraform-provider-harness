package connector_test

/*
OpenAI Model Connector — Terraform Data Source Acceptance Tests
Tests the data.harness_platform_connector_openai_model data source.

See openai_model_test.go for env var documentation per scope.

---------------------------------------------------------------------------
# Run ALL data source acceptance tests
---------------------------------------------------------------------------

    go test -v -run "Test.*DataSourceConnectorOpenAIModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...
*/

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// ============================================================================
// Acceptance Tests — Account-scope
// ============================================================================

func TestAccDataSourceConnectorOpenAIModel_Token(t *testing.T) {
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_token(id, name, tokenRef, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorOpenAIModel_Vertex(t *testing.T) {
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	if saKeyRef == "" || projectID == "" || region == "" {
		t.Skip("OPENAI_MODEL_VERTEX_* not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_vertex(id, name, saKeyRef, projectID, region, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
				),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Org-scope
// ============================================================================

func TestOrgDataSourceConnectorOpenAIModel_Token(t *testing.T) {
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	orgID := os.Getenv("OPENAI_MODEL_ORG_SCOPE_ORG_ID")
	if tokenRef == "" || orgID == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF and OPENAI_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	delegateCSV := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS")
	delegateHCL := openaiFormatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_tokenOrg(id, name, tokenRef, orgID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorOpenAIModel_Vertex(t *testing.T) {
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
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_vertexOrg(id, name, saKeyRef, projectID, region, orgID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
				),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Project-scope
// ============================================================================

func TestProjectDataSourceConnectorOpenAIModel_Token(t *testing.T) {
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
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_tokenProject(id, name, tokenRef, orgID, projectID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorOpenAIModel_Vertex(t *testing.T) {
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
	resourceName := "data.harness_platform_connector_openai_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorOpenAIModel_vertexProject(id, name, saKeyRef, gcpProjectID, region, orgID, projectID, delegateHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Vertex"),
				),
			},
		},
	})
}

// ============================================================================
// HCL config generators — Data Source (resource + data source)
// ============================================================================

func testAccDataSourceConnectorOpenAIModel_token(id, name, tokenRef, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_token(id, name, tokenRef, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
}`)
}

func testAccDataSourceConnectorOpenAIModel_vertex(id, name, saKeyRef, gcpProjectID, region, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_vertex(id, name, saKeyRef, gcpProjectID, region, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
}`)
}

func testAccDataSourceConnectorOpenAIModel_tokenOrg(id, name, tokenRef, orgID, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_tokenOrg(id, name, tokenRef, orgID, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
  org_id     = "%s"
}`, orgID)
}

func testAccDataSourceConnectorOpenAIModel_vertexOrg(id, name, saKeyRef, gcpProjectID, region, orgID, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_vertexOrg(id, name, saKeyRef, gcpProjectID, region, orgID, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
  org_id     = "%s"
}`, orgID)
}

func testAccDataSourceConnectorOpenAIModel_tokenProject(id, name, tokenRef, orgID, projectID, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_tokenProject(id, name, tokenRef, orgID, projectID, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
  org_id     = "%s"
  project_id = "%s"
}`, orgID, projectID)
}

func testAccDataSourceConnectorOpenAIModel_vertexProject(id, name, saKeyRef, gcpProjectID, region, orgID, projectID, delegateHCL string) string {
	return testAccResourceConnectorOpenAIModel_vertexProject(id, name, saKeyRef, gcpProjectID, region, orgID, projectID, delegateHCL) + fmt.Sprintf(`
data "harness_platform_connector_openai_model" "test" {
  identifier = harness_platform_connector_openai_model.test.identifier
  org_id     = "%s"
  project_id = "%s"
}`, orgID, projectID)
}

