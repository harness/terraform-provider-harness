package connector_test

/*
Anthropic Model Connector — Terraform Data Source Acceptance & Unit Tests
Tests the data.harness_platform_connector_anthropic_model data source. Each test
first creates the connector resource, then reads it back via the data source and
verifies all attributes match.

The file is organized as:
  1. Acceptance Tests — run against a live Harness account, env-var driven, auto-skip
     if required vars are not set. 9 tests total.
  2. Unit Tests — use hardcoded inline HCL with fake values, test provider schema logic
     and data source read wiring without requiring pre-existing Harness resources. 9 tests total.

Each test has one step:
  1. Create resource + read via data source — verifies all attributes on the data source

Tests skip automatically if the required env vars are not set.

Data source fields verified:
  - identifier, name:        match the resource
  - description:             "test"
  - tags:                    ["foo:bar"]
  - url:                     https://api.anthropic.com
  - model:                   claude-opus-4-6
  - execute_on_delegate:     true
  - ignore_test_connection:  true
  - delegate_selectors:      from ANTHROPIC_MODEL_DELEGATE_SELECTORS (optional, comma-separated)
  - org_id:                  set for org/project scope
  - project_id:              set for project scope
  - auth.0.auth_type:        Token | BedrockApiKey | CloudProvider

Optional env vars (apply to all scopes):

    # Comma-separated delegate selectors (optional, omit to use empty list)
    export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

---------------------------------------------------------------------------
# Account-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export HARNESS_ENDPOINT='https://app.harness.io/gateway'
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="account.<account-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="account.<account-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="account.<account-level-cloud-connector>"
    # Optional
    export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccDataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Org-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID="<org-id>"
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="org.<org-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="org.<org-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="org.<org-level-cloud-connector>"
    # Optional
    export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestOrgDataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Project-Scope Acceptance Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
    export ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID="<org-id>"
    export ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID="<project-id>"
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="<project-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="<project-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="<project-level-cloud-connector>"
    # Optional
    export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestProjectDataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Run ALL acceptance data source tests
---------------------------------------------------------------------------

    go test -v -run "Test(Acc|Org|Project)DataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Run ALL unit tests
---------------------------------------------------------------------------

    go test -v -run "TestUnit.*DataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
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
// Acceptance Tests — Account-scope data source (env-var driven, skip if not set)
// ============================================================================

func TestAccDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_Token(name, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	bedrockApiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	if bedrockApiKeyRef == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF not set, skipping")
	}
	bedrockRegion := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if bedrockRegion == "" {
		bedrockRegion = "us-east-1"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_BedrockApiKey(name, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", bedrockRegion),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	cloudConnectorRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cloudConnectorRef == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF not set, skipping")
	}
	cloudProviderType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	if cloudProviderType == "" {
		cloudProviderType = "AWS"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_CloudProvider(name, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", cloudProviderType),
				),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Org-scope data source (env-var driven, skip if not set)
// ============================================================================

func TestOrgDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_OrgToken(name, orgID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	bedrockApiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	if bedrockApiKeyRef == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF not set, skipping")
	}
	bedrockRegion := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if bedrockRegion == "" {
		bedrockRegion = "us-east-1"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_OrgBedrockApiKey(name, orgID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", bedrockRegion),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID not set, skipping")
	}
	cloudConnectorRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cloudConnectorRef == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF not set, skipping")
	}
	cloudProviderType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	if cloudProviderType == "" {
		cloudProviderType = "AWS"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_OrgCloudProvider(name, orgID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", cloudProviderType),
				),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Project-scope data source (env-var driven, skip if not set)
// ============================================================================

func TestProjectDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_ProjectToken(name, orgID, projectID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	bedrockApiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	if bedrockApiKeyRef == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF not set, skipping")
	}
	bedrockRegion := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if bedrockRegion == "" {
		bedrockRegion = "us-east-1"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_ProjectBedrockApiKey(name, orgID, projectID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", bedrockRegion),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	orgID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	cloudConnectorRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cloudConnectorRef == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF not set, skipping")
	}
	cloudProviderType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	if cloudProviderType == "" {
		cloudProviderType = "AWS"
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAnthropicModel_ProjectCloudProvider(name, orgID, projectID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", fmt.Sprintf("%d", delegateSelectorCount(delegateCSV))),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", cloudProviderType),
				),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — HCL configs (account-scope, env-var driven)
// ============================================================================

func testAccDataSourceConnectorAnthropicModel_Token(name, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[2]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "%[3]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name, delegateSelectorsHCL, tokenRef)
}

func testAccDataSourceConnectorAnthropicModel_BedrockApiKey(name, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[2]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "%[3]s"
				region = "%[4]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccDataSourceConnectorAnthropicModel_CloudProvider(name, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[2]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "%[3]s"
				connector_ref = "%[4]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Acceptance Tests — HCL configs (org-scope, env-var driven)
// ============================================================================

func testAccDataSourceConnectorAnthropicModel_OrgToken(name, orgID, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[3]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "%[4]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
	}
	`, name, orgID, delegateSelectorsHCL, tokenRef)
}

func testAccDataSourceConnectorAnthropicModel_OrgBedrockApiKey(name, orgID, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[3]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "%[4]s"
				region = "%[5]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
	}
	`, name, orgID, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccDataSourceConnectorAnthropicModel_OrgCloudProvider(name, orgID, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[3]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "%[4]s"
				connector_ref = "%[5]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
	}
	`, name, orgID, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Acceptance Tests — HCL configs (project-scope, env-var driven)
// ============================================================================

func testAccDataSourceConnectorAnthropicModel_ProjectToken(name, orgID, projectID, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"
		project_id = "%[3]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[4]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "%[5]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
		project_id = "%[3]s"
	}
	`, name, orgID, projectID, delegateSelectorsHCL, tokenRef)
}

func testAccDataSourceConnectorAnthropicModel_ProjectBedrockApiKey(name, orgID, projectID, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"
		project_id = "%[3]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[4]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "%[5]s"
				region = "%[6]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
		project_id = "%[3]s"
	}
	`, name, orgID, projectID, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccDataSourceConnectorAnthropicModel_ProjectCloudProvider(name, orgID, projectID, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[2]s"
		project_id = "%[3]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[4]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "%[5]s"
				connector_ref = "%[6]s"
			}
		}
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = "%[2]s"
		project_id = "%[3]s"
	}
	`, name, orgID, projectID, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Following are Unit Tests — use hardcoded inline HCL with test resources
// (secrets, orgs, projects, AWS connectors) created within the test config.
// These test the provider schema and data source read wiring without requiring
// pre-existing Harness resources via env vars.
//
// Run:
//   go test -v -run "TestUnit.*DataSourceConnectorAnthropicModel" -count=1 -timeout 120m \
//     ./internal/service/platform/connector/...
// ============================================================================

// ---------------------------------------------------------------------------
// Unit Tests — Account-scope data source
// ---------------------------------------------------------------------------

func TestUnitAccDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_Token(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestUnitAccDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_BedrockApiKey(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", "us-east-1"),
				),
			},
		},
	})
}

func TestUnitAccDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_CloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "identifier", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.anthropic.com"),
					resource.TestCheckResourceAttr(resourceName, "model", "claude-opus-4-6"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "ignore_test_connection", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", "AWS"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unit Tests — Org-scope data source
// ---------------------------------------------------------------------------

func TestUnitOrgDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_OrgToken(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestUnitOrgDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_OrgBedrockApiKey(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", "us-east-1"),
				),
			},
		},
	})
}

func TestUnitOrgDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_OrgCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "identifier", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", "AWS"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unit Tests — Project-scope data source
// ---------------------------------------------------------------------------

func TestUnitProjectDataSourceConnectorAnthropicModel_Token(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_ProjectToken(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
		},
	})
}

func TestUnitProjectDataSourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_ProjectBedrockApiKey(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", "us-east-1"),
				),
			},
		},
	})
}

func TestUnitProjectDataSourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testUnitDataSourceConnectorAnthropicModel_ProjectCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "identifier", name+"_anthropic"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", "AWS"),
				),
			},
		},
	})
}

// ============================================================================
// Unit Tests — HCL configs (account-scope, inline resources)
// ============================================================================

func testUnitDataSourceConnectorAnthropicModel_Token(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_BedrockApiKey(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "account.${harness_platform_secret_text.test.id}"
				region = "us-east-1"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_CloudProvider(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_aws" "aws_test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test aws connector for anthropic cloud provider"
		tags = ["foo:bar"]

		manual {
			secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			access_key_ref = "account.${harness_platform_secret_text.test.id}"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s_anthropic"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "AWS"
				connector_ref = "account.${harness_platform_connector_aws.aws_test.id}"
			}
		}
		depends_on = [harness_platform_connector_aws.aws_test]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
	}
	`, name)
}

// ============================================================================
// Unit Tests — HCL configs (org-scope, inline resources)
// ============================================================================

func testUnitDataSourceConnectorAnthropicModel_OrgToken(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "org.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_OrgBedrockApiKey(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "org.${harness_platform_secret_text.test.id}"
				region = "us-east-1"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_OrgCloudProvider(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_aws" "aws_test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test aws connector for anthropic cloud provider"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		manual {
			secret_key_ref = "org.${harness_platform_secret_text.test.id}"
			access_key_ref = "org.${harness_platform_secret_text.test.id}"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s_anthropic"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "AWS"
				connector_ref = "org.${harness_platform_connector_aws.aws_test.id}"
			}
		}
		depends_on = [harness_platform_connector_aws.aws_test]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
	}
	`, name)
}

// ============================================================================
// Unit Tests — HCL configs (project-scope, inline resources)
// ============================================================================

func testUnitDataSourceConnectorAnthropicModel_ProjectToken(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_ProjectBedrockApiKey(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "${harness_platform_secret_text.test.id}"
				region = "us-east-1"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name)
}

func testUnitDataSourceConnectorAnthropicModel_ProjectCloudProvider(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_aws" "aws_test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test aws connector for anthropic cloud provider"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		manual {
			secret_key_ref = "${harness_platform_secret_text.test.id}"
			access_key_ref = "${harness_platform_secret_text.test.id}"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s_anthropic"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = ["harness-delegate"]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "AWS"
				connector_ref = "${harness_platform_connector_aws.aws_test.id}"
			}
		}
		depends_on = [harness_platform_connector_aws.aws_test]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	data "harness_platform_connector_anthropic_model" "test" {
		identifier = harness_platform_connector_anthropic_model.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name)
}
