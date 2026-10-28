package token_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceToken(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
	parent_id := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")

	tokenName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := tokenName + "updated"

	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceToken(id, tokenName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", tokenName),
				),
			},
			{
				Config: testAccResourceToken(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       AccountResourceImportTokenIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})

}

func TestAccResourceTokenOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
	parent_id := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")

	tokenName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := tokenName + "updated"

	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceToken(id, tokenName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", tokenName),
				),
			},
			{
				Config: testOrgResourceToken(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       OrgResourceImportTokenIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})

}

func TestAccResourceTokenProjectLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
	parent_id := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")

	tokenName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := tokenName + "updated"

	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceToken(id, tokenName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", tokenName),
				),
			},
			{
				Config: testProjectResourceToken(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       ProjectResourceImportTokenIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})

}

func testAccGetResourceToken(resourceName string, state *terraform.State) (*nextgen.Token, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.TokenApi.ListAggregatedTokens(ctx, c.AccountId, r.Primary.Attributes["apikey_type"], r.Primary.Attributes["parent_id"], r.Primary.Attributes["apikey_id"], &nextgen.TokenApiListAggregatedTokensOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
		Identifiers:       optional.NewInterface(id),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil || resp.Data.Content == nil || len(resp.Data.Content) == 0 {
		return nil, nil
	}

	return resp.Data.Content[0].Token, nil
}

func testAccResourceToken(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_apikey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			apikey_type = "USER"
			default_time_to_expire_token = 1000000
		}

		resource "harness_platform_token" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			apikey_type = "USER"
			apikey_id = harness_platform_apikey.test.id
		}
	`, id, name, parentId, accountId)
}

func testOrgResourceToken(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_apikey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000000
		}

		resource "harness_platform_token" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			apikey_type = "USER"
			apikey_id = harness_platform_apikey.test.id
		}
	`, id, name, parentId, accountId)
}

func testProjectResourceToken(id string, name string, parentId string, accountId string) string {
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

		resource "harness_platform_apikey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000000
		}

		resource "harness_platform_token" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			apikey_type = "USER"
			apikey_id = harness_platform_apikey.test.id
		}
	`, id, name, parentId, accountId)
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccTokenDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		token, _ := testAccGetResourceToken(resourceName, state)
		if token != nil {
			return fmt.Errorf("Found token: %s", token.Identifier)
		}

		return nil
	}
}

func AccountResourceImportTokenIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		apikey_id := primary.Attributes["apikey_id"]
		parent_id := primary.Attributes["parent_id"]
		apikey_type := primary.Attributes["apikey_type"]
		return fmt.Sprintf("%s/%s/%s/%s", parent_id, apikey_id, apikey_type, id), nil
	}
}

func OrgResourceImportTokenIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		orgId := primary.Attributes["org_id"]
		apikey_id := primary.Attributes["apikey_id"]
		parent_id := primary.Attributes["parent_id"]
		apikey_type := primary.Attributes["apikey_type"]
		return fmt.Sprintf("%s/%s/%s/%s/%s", orgId, parent_id, apikey_id, apikey_type, id), nil
	}
}

func ProjectResourceImportTokenIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		orgId := primary.Attributes["org_id"]
		projId := primary.Attributes["project_id"]
		apikey_id := primary.Attributes["apikey_id"]
		parent_id := primary.Attributes["parent_id"]
		apikey_type := primary.Attributes["apikey_type"]
		return fmt.Sprintf("%s/%s/%s/%s/%s/%s", orgId, projId, parent_id, apikey_id, apikey_type, id), nil
	}
}

/*
---------------------------------------------------------------------------
# HSF-99 Project-Scope Lifecycle Acceptance Tests
---------------------------------------------------------------------------
    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'

    export TOKEN_PROJECT_SCOPE_ORG_ID='<org-id>'
    export TOKEN_PROJECT_SCOPE_PROJECT_ID='<project-id>'
    export TOKEN_APIKEY_TYPE='USER'   # USER or SERVICE_ACCOUNT (defaults to USER)
    export HARNESS_PAT_KEY_PARENT_IDENTIFIER='<user-or-sa-id>'

    go test -v -run "TestProjectResourceTokenLifecycle" -count=1 -timeout 120m \
      ./internal/service/platform/token/...
*/

func TestProjectResourceTokenLifecycle_IgnoreChanges(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_IgnoreChanges(id, "initial name", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", "initial name"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
				),
			},
			{
				// name changed in config but ignore_changes = [name] must suppress the diff
				Config:             testAccProjectTokenLifecycle_IgnoreChanges(id, "changed name", orgID, projectID, parentID, accountID, apikeyType),
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
				),
			},
		},
	})
}

func TestProjectResourceTokenLifecycle_CreateBeforeDestroy(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_CreateBeforeDestroy(id, "token name", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
				),
			},
			{
				// identifier change forces replace; create_before_destroy ensures new resource
				// is provisioned before the old one is removed
				Config: testAccProjectTokenLifecycle_CreateBeforeDestroy(id+"_new", "token name", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id+"_new"),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
				),
			},
		},
	})
}

func TestProjectResourceTokenLifecycle_PreventDestroy(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_PreventDestroy(id, "token name", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
				),
			},
			{
				// destroy while prevent_destroy = true must be rejected by Terraform
				Config:      testAccProjectTokenLifecycle_PreventDestroy(id, "token name", orgID, projectID, parentID, accountID, apikeyType),
				Destroy:     true,
				ExpectError: regexp.MustCompile(`prevent_destroy`),
			},
			{
				// drop prevent_destroy so the test framework can clean up the resource
				Config:  testAccProjectTokenLifecycle_Base(id, "token name", orgID, projectID, parentID, accountID, apikeyType),
				Destroy: true,
			},
		},
	})
}

func TestProjectResourceTokenLifecycle_ReplaceTriggeredBy(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	apikeyID := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	tokenID := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_ReplaceTriggeredBy(apikeyID, tokenID, "apikey name", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", tokenID),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
				),
			},
			{
				// changing the apikey name triggers token replacement via replace_triggered_by
				Config: testAccProjectTokenLifecycle_ReplaceTriggeredBy(apikeyID, tokenID, "apikey name updated", orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", tokenID),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
				),
			},
		},
	})
}

func TestProjectResourceTokenLifecycle_Precondition(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_Precondition(id, apikeyType, apikeyType, orgID, projectID, parentID, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
				),
			},
			{
				// apikey stays valid; only the token variable is set to an invalid type so the
				// precondition fires before the token apply rather than apikey schema validation
				Config:      testAccProjectTokenLifecycle_Precondition(id, "INVALID_TYPE", apikeyType, orgID, projectID, parentID, accountID),
				ExpectError: regexp.MustCompile(`apikey_type must be USER or SERVICE_ACCOUNT`),
			},
		},
	})
}

func TestProjectResourceTokenLifecycle_Postcondition(t *testing.T) {
	orgID := os.Getenv("TOKEN_PROJECT_SCOPE_ORG_ID")
	if orgID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_ORG_ID not set, skipping")
	}
	projectID := os.Getenv("TOKEN_PROJECT_SCOPE_PROJECT_ID")
	if projectID == "" {
		t.Skip("TOKEN_PROJECT_SCOPE_PROJECT_ID not set, skipping")
	}
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID not set, skipping")
	}
	parentID := os.Getenv("HARNESS_PAT_KEY_PARENT_IDENTIFIER")
	if parentID == "" {
		t.Skip("HARNESS_PAT_KEY_PARENT_IDENTIFIER not set, skipping")
	}
	apikeyType := os.Getenv("TOKEN_APIKEY_TYPE")
	if apikeyType == "" {
		apikeyType = "USER"
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_token.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTokenLifecycle_Postcondition(id, orgID, projectID, parentID, accountID, apikeyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "apikey_type", apikeyType),
					resource.TestCheckResourceAttrSet(resourceName, "value"),
				),
			},
		},
	})
}

func testAccProjectTokenLifecycle_Base(id, name, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[5]s"
  account_id                   = "%[6]s"
  org_id                       = "%[3]s"
  project_id                   = "%[4]s"
  apikey_type                  = "%[7]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[2]s"
  parent_id   = "%[5]s"
  account_id  = "%[6]s"
  org_id      = "%[3]s"
  project_id  = "%[4]s"
  apikey_type = "%[7]s"
  apikey_id   = harness_platform_apikey.test.id
}
`, id, name, orgID, projectID, parentID, accountID, apikeyType)
}

func testAccProjectTokenLifecycle_IgnoreChanges(id, name, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[5]s"
  account_id                   = "%[6]s"
  org_id                       = "%[3]s"
  project_id                   = "%[4]s"
  apikey_type                  = "%[7]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[2]s"
  parent_id   = "%[5]s"
  account_id  = "%[6]s"
  org_id      = "%[3]s"
  project_id  = "%[4]s"
  apikey_type = "%[7]s"
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    ignore_changes = [name]
  }
}
`, id, name, orgID, projectID, parentID, accountID, apikeyType)
}

func testAccProjectTokenLifecycle_CreateBeforeDestroy(id, name, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[5]s"
  account_id                   = "%[6]s"
  org_id                       = "%[3]s"
  project_id                   = "%[4]s"
  apikey_type                  = "%[7]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[2]s"
  parent_id   = "%[5]s"
  account_id  = "%[6]s"
  org_id      = "%[3]s"
  project_id  = "%[4]s"
  apikey_type = "%[7]s"
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    create_before_destroy = true
  }
}
`, id, name, orgID, projectID, parentID, accountID, apikeyType)
}

func testAccProjectTokenLifecycle_PreventDestroy(id, name, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[5]s"
  account_id                   = "%[6]s"
  org_id                       = "%[3]s"
  project_id                   = "%[4]s"
  apikey_type                  = "%[7]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[2]s"
  parent_id   = "%[5]s"
  account_id  = "%[6]s"
  org_id      = "%[3]s"
  project_id  = "%[4]s"
  apikey_type = "%[7]s"
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    prevent_destroy = true
  }
}
`, id, name, orgID, projectID, parentID, accountID, apikeyType)
}

func testAccProjectTokenLifecycle_ReplaceTriggeredBy(apikeyID, tokenID, apikeyName, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[3]s"
  parent_id                    = "%[6]s"
  account_id                   = "%[7]s"
  org_id                       = "%[4]s"
  project_id                   = "%[5]s"
  apikey_type                  = "%[8]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[2]s"
  name        = "%[2]s"
  parent_id   = "%[6]s"
  account_id  = "%[7]s"
  org_id      = "%[4]s"
  project_id  = "%[5]s"
  apikey_type = "%[8]s"
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    replace_triggered_by = [harness_platform_apikey.test.name]
  }
}
`, apikeyID, tokenID, apikeyName, orgID, projectID, parentID, accountID, apikeyType)
}

// tokenApikeyType is injected into the variable block (token resource); apikeyType is used
// directly on harness_platform_apikey so its ValidateFunc is always satisfied.
func testAccProjectTokenLifecycle_Precondition(id, tokenApikeyType, apikeyType, orgID, projectID, parentID, accountID string) string {
	return fmt.Sprintf(`
variable "token_apikey_type" {
  default = "%[2]s"
}

resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[6]s"
  account_id                   = "%[7]s"
  org_id                       = "%[4]s"
  project_id                   = "%[5]s"
  apikey_type                  = "%[3]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[1]s"
  parent_id   = "%[6]s"
  account_id  = "%[7]s"
  org_id      = "%[4]s"
  project_id  = "%[5]s"
  apikey_type = var.token_apikey_type
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    precondition {
      condition     = contains(["USER", "SERVICE_ACCOUNT"], var.token_apikey_type)
      error_message = "apikey_type must be USER or SERVICE_ACCOUNT"
    }
  }
}
`, id, tokenApikeyType, apikeyType, orgID, projectID, parentID, accountID)
}

func testAccProjectTokenLifecycle_Postcondition(id, orgID, projectID, parentID, accountID, apikeyType string) string {
	return fmt.Sprintf(`
resource "harness_platform_apikey" "test" {
  identifier                   = "%[1]s"
  name                         = "%[1]s"
  parent_id                    = "%[4]s"
  account_id                   = "%[5]s"
  org_id                       = "%[2]s"
  project_id                   = "%[3]s"
  apikey_type                  = "%[6]s"
  default_time_to_expire_token = 1000000
}

resource "harness_platform_token" "test" {
  identifier  = "%[1]s"
  name        = "%[1]s"
  parent_id   = "%[4]s"
  account_id  = "%[5]s"
  org_id      = "%[2]s"
  project_id  = "%[3]s"
  apikey_type = "%[6]s"
  apikey_id   = harness_platform_apikey.test.id

  lifecycle {
    postcondition {
      condition     = self.value != ""
      error_message = "token value must be non-empty after creation"
    }
  }
}
`, id, orgID, projectID, parentID, accountID, apikeyType)
}
