package connector_test

/*
--------------------------------------------------------------------------
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

    go test -v -run "TestAccResourceConnectorAnthropicModel" -count=1 -timeout 120m \
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

    go test -v -run "TestOrgResourceConnectorAnthropicModel" -count=1 -timeout 120m \
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

    go test -v -run "TestProjectResourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...

---------------------------------------------------------------------------
# Run ALL unit tests
---------------------------------------------------------------------------

    go test -v -run "TestUnit.*ResourceConnectorAnthropicModel" -count=1 -timeout 120m \
      ./internal/service/platform/connector/...
*/

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// ---------------------------------------------------------------------------
// Helpers: read env vars, format delegate selectors for HCL
// ---------------------------------------------------------------------------

func formatDelegateSelectorsHCL(csv string) string {
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

func delegateSelectorCount(csv string) int {
	if csv == "" {
		return 0
	}
	return len(strings.Split(csv, ","))
}

// ============================================================================
// Acceptance Tests — Account-scope (env-var driven, skip if not set)
// ============================================================================

func TestAccResourceConnectorAnthropicModel_Token(t *testing.T) {
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF not set, skipping")
	}
	delegateCSV := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS")
	delegateHCL := formatDelegateSelectorsHCL(delegateCSV)

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_Token(id, name, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_Token(id, updatedName, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
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

func TestAccResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_BedrockApiKey(id, name, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_BedrockApiKey(id, updatedName, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
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

func TestAccResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_CloudProvider(id, name, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_CloudProvider(id, updatedName, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
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
// Acceptance Tests — Org-scope (env-var driven, skip if not set)
// ============================================================================

func TestOrgResourceConnectorAnthropicModel_Token(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_OrgToken(id, name, orgID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_OrgToken(id, updatedName, orgID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestOrgResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_OrgBedrockApiKey(id, name, orgID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_OrgBedrockApiKey(id, updatedName, orgID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestOrgResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_OrgCloudProvider(id, name, orgID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_OrgCloudProvider(id, updatedName, orgID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// ============================================================================
// Acceptance Tests — Project-scope (env-var driven, skip if not set)
// ============================================================================

func TestProjectResourceConnectorAnthropicModel_Token(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectToken(id, name, orgID, projectID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectToken(id, updatedName, orgID, projectID, delegateHCL, tokenRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
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

func TestProjectResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectBedrockApiKey(id, name, orgID, projectID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectBedrockApiKey(id, updatedName, orgID, projectID, delegateHCL, bedrockApiKeyRef, bedrockRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
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

func TestProjectResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
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

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectCloudProvider(id, name, orgID, projectID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAnthropicModel_ProjectCloudProvider(id, updatedName, orgID, projectID, delegateHCL, cloudProviderType, cloudConnectorRef),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
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
// Acceptance Tests — HCL configs (account-scope, env-var driven)
// ============================================================================

func testAccResourceConnectorAnthropicModel_Token(id, name, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

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
	`, id, name, delegateSelectorsHCL, tokenRef)
}

func testAccResourceConnectorAnthropicModel_BedrockApiKey(id, name, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

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
	`, id, name, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccResourceConnectorAnthropicModel_CloudProvider(id, name, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

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
	`, id, name, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Acceptance Tests — HCL configs (org-scope, env-var driven)
// ============================================================================

func testAccResourceConnectorAnthropicModel_OrgToken(id, name, orgID, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"

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
	`, id, name, orgID, delegateSelectorsHCL, tokenRef)
}

func testAccResourceConnectorAnthropicModel_OrgBedrockApiKey(id, name, orgID, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"

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
	`, id, name, orgID, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccResourceConnectorAnthropicModel_OrgCloudProvider(id, name, orgID, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"

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
	`, id, name, orgID, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Acceptance Tests — HCL configs (project-scope, env-var driven)
// ============================================================================

func testAccResourceConnectorAnthropicModel_ProjectToken(id, name, orgID, projectID, delegateSelectorsHCL, tokenRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"
		project_id = "%[4]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[5]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "Token"
			token {
				token_ref = "%[6]s"
			}
		}
	}
	`, id, name, orgID, projectID, delegateSelectorsHCL, tokenRef)
}

func testAccResourceConnectorAnthropicModel_ProjectBedrockApiKey(id, name, orgID, projectID, delegateSelectorsHCL, apiKeyRef, region string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"
		project_id = "%[4]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[5]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "BedrockApiKey"
			bedrock_api_key {
				api_key_ref = "%[6]s"
				region = "%[7]s"
			}
		}
	}
	`, id, name, orgID, projectID, delegateSelectorsHCL, apiKeyRef, region)
}

func testAccResourceConnectorAnthropicModel_ProjectCloudProvider(id, name, orgID, projectID, delegateSelectorsHCL, providerType, connectorRef string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = "%[3]s"
		project_id = "%[4]s"

		url = "https://api.anthropic.com"
		model = "claude-opus-4-6"
		delegate_selectors = [%[5]s]
		execute_on_delegate = true
		ignore_test_connection = true
		auth {
			auth_type = "CloudProvider"
			cloud_provider {
				type = "%[6]s"
				connector_ref = "%[7]s"
			}
		}
	}
	`, id, name, orgID, projectID, delegateSelectorsHCL, providerType, connectorRef)
}

// ============================================================================
// Following are Unit Tests
// ============================================================================

// ---------------------------------------------------------------------------
// Unit Tests — Account-scope
// ---------------------------------------------------------------------------

func TestUnitAccResourceConnectorAnthropicModel_Token(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_Token(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testUnitResourceConnectorAnthropicModel_Token(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
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

func TestUnitAccResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_BedrockApiKey(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testUnitResourceConnectorAnthropicModel_BedrockApiKey(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
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

func TestUnitAccResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_CloudProvider(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testUnitResourceConnectorAnthropicModel_CloudProvider(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
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

// ---------------------------------------------------------------------------
// Unit Tests — Org-scope
// ---------------------------------------------------------------------------

func TestUnitOrgResourceConnectorAnthropicModel_Token(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgToken(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgToken(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestUnitOrgResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgBedrockApiKey(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", "us-east-1"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgBedrockApiKey(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestUnitOrgResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgCloudProvider(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", "AWS"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_OrgCloudProvider(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateVerify:  true,
				ExpectNonEmptyPlan: true,
				ImportStateIdFunc:  acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unit Tests — Project-scope
// ---------------------------------------------------------------------------

func TestUnitProjectResourceConnectorAnthropicModel_Token(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectToken(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "Token"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectToken(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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

func TestUnitProjectResourceConnectorAnthropicModel_BedrockApiKey(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectBedrockApiKey(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "BedrockApiKey"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.bedrock_api_key.0.region", "us-east-1"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectBedrockApiKey(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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

func TestUnitProjectResourceConnectorAnthropicModel_CloudProvider(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_anthropic_model.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectCloudProvider(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "auth.0.auth_type", "CloudProvider"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.cloud_provider.0.type", "AWS"),
				),
			},
			{
				Config: testUnitResourceConnectorAnthropicModel_ProjectCloudProvider(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
// Unit Tests — HCL configs (account-scope, inline resources)
// ============================================================================

func testUnitResourceConnectorAnthropicModel_Token(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_BedrockApiKey(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_CloudProvider(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_aws" "aws_test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
		name = "%[2]s"
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
`, id, name)
}

// ============================================================================
// Unit Tests — HCL configs (org-scope, inline resources)
// ============================================================================

func testUnitResourceConnectorAnthropicModel_OrgToken(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_OrgBedrockApiKey(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_anthropic_model" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_OrgCloudProvider(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_aws" "aws_test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
		name = "%[2]s"
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
`, id, name)
}

// ============================================================================
// Unit Tests — HCL configs (project-scope, inline resources)
// ============================================================================

func testUnitResourceConnectorAnthropicModel_ProjectToken(id string, name string) string {
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

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_ProjectBedrockApiKey(id string, name string) string {
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

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
		name = "%[2]s"
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
`, id, name)
}

func testUnitResourceConnectorAnthropicModel_ProjectCloudProvider(id string, name string) string {
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

	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
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
		name = "%[2]s"
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
		name = "%[2]s"
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
`, id, name)
}
