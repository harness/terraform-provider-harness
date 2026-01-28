package delegatetoken_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"net/http"
)

func TestAccResourceDelegateToken(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testDelegateTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: tesAccResourceDelegateToken(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccResourceDelegateTokenOrgLevel(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
	org_id := utils.RandStringBytes(5)

	resourceName := "harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testDelegateTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: tesAccResourceDelegateTokenOrgLevel(name, account_id, org_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccResourceDelegateTokenProjectLevel(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
	org_id := utils.RandStringBytes(5)
	project_id := utils.RandStringBytes(5)

	resourceName := "harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testDelegateTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: tesAccResourceDelegateTokenProjectLevel(name, account_id, org_id, project_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccResourceDelegateTokenWithRevokeAfter(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testDelegateTokenDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDelegateTokenWithRevokeAfter(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "revoke_after", "1769689600000"),
				),
			},
		},
	})
}

//will add these tests when change DELETE CONTEXT to call delete delegate token api

// func TestAccResourceDelegateTokenUpdate(t *testing.T) {
// 	name := utils.RandStringBytes(5)
// 	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

// 	resourceName := "harness_platform_delegatetoken.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: tesAccResourceDelegateToken(name, account_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
// 				),
// 			},
// 			{
// 				Config: tesAccResourceDelegateTokenUpdateAccountLevelUpdated(name, account_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "REVOKED"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccResourceDelegateTokenUpdateOrgLevel(t *testing.T) {
// 	name := utils.RandStringBytes(5)
// 	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
// 	org_id := utils.RandStringBytes(5)

// 	resourceName := "harness_platform_delegatetoken.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: tesAccResourceDelegateTokenOrgLevel(name, account_id, org_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
// 				),
// 			},
// 			{
// 				Config: tesAccResourceDelegateTokenUpdateOrgLevelUpdated(name, account_id, org_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "REVOKED"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccResourceDelegateTokenUpdateProjectLevel(t *testing.T) {
// 	name := utils.RandStringBytes(5)
// 	account_id := os.Getenv("HARNESS_ACCOUNT_ID")
// 	org_id := utils.RandStringBytes(5)
// 	project_id := utils.RandStringBytes(5)

// 	resourceName := "harness_platform_delegatetoken.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: tesAccResourceDelegateTokenProjectLevel(name, account_id, org_id, project_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
// 				),
// 			},
// 			{
// 				Config: tesAccResourceDelegateTokenUpdateProjectLevelUpdated(name, account_id, org_id, project_id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "token_status", "REVOKED"),
// 				),
// 			},
// 		},
// 	})
// }

func tesAccResourceDelegateToken(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_delegatetoken" "test" {			
			name = "%[1]s"
			account_id = "%[2]s"
		}
		`, name, accountId)
}

func tesAccResourceDelegateTokenOrgLevel(name string, accountId string, org_id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_delegatetoken" "test" {			
			name = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
		}
		`, name, accountId, org_id)
}

func tesAccResourceDelegateTokenProjectLevel(name string, accountId string, org_id string, project_id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[3]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[4]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_delegatetoken" "test" {			
			name = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
		`, name, accountId, org_id, project_id)
}

//will add these tests when change DELETE CONTEXT to call delete delegate token api

// func tesAccResourceDelegateTokenUpdateAccountLevelUpdated(name string, accountId string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_platform_delegatetoken" "test" {
// 			name = "%[1]s"
// 			account_id = "%[2]s"
// 			token_status = "REVOKED"
// 		}
// 		`, name, accountId)
// }

// func tesAccResourceDelegateTokenUpdateOrgLevelUpdated(name string, accountId string, org_id string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_platform_organization" "test" {
// 			identifier = "%[1]s"
// 			name = "%[1]s"
// 		}

// 		resource "harness_platform_delegatetoken" "test" {
// 			name = "%[1]s"
// 			account_id = "%[2]s"
// 			org_id = harness_platform_organization.test.id
// 			token_status = "REVOKED"
// 		}
// 		`, name, accountId, org_id)
// }

// func tesAccResourceDelegateTokenUpdateProjectLevelUpdated(name string, accountId string, org_id string, project_id string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_platform_organization" "test" {
// 			identifier = "%[3]s"
// 			name = "%[1]s"
// 		}

// 		resource "harness_platform_project" "test" {
// 			identifier = "%[4]s"
// 			name = "%[1]s"
// 			org_id = harness_platform_organization.test.id
// 			color = "#472848"
// 		}

// 		resource "harness_platform_delegatetoken" "test" {
// 			name = "%[1]s"
// 			account_id = "%[2]s"
// 			org_id = harness_platform_organization.test.id
// 			project_id = harness_platform_project.test.id
// 			token_status = "REVOKED"
// 		}
// 		`, name, accountId, org_id, project_id)
// }

func testAccResourceDelegateTokenWithRevokeAfter(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_delegatetoken" "test" {			
			name = "%[1]s"
			account_id = "%[2]s"
			revoke_after = 1769689600000 
		}
		`, name, accountId)
}

func testAccGetResourceDelegateToken(resourceName string, state *terraform.State) (*nextgen.DelegateTokenDetails, error) {
	d := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	resp, _, err := c.DelegateTokenResourceApi.GetCgDelegateTokens(ctx, c.AccountId, &nextgen.DelegateTokenResourceApiGetCgDelegateTokensOpts{
		OrgIdentifier:     buildField(d, "org_id"),
		ProjectIdentifier: buildField(d, "project_id"),
		Name:              buildField(d, "name"),
		Status:            buildField(d, "token_status"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Resource == nil {
		return nil, nil
	}

	return &resp.Resource[0], nil
}

func testDelegateTokenDestroy(resourceName string) resource.TestCheckFunc {
	var token *nextgen.DelegateTokenDetails
	return func(state *terraform.State) error {
		token, _ = testAccGetResourceDelegateToken(resourceName, state)
		if token != nil && token.Status != "REVOKED" {
			return fmt.Errorf("Token is not revoked : %s", token.Name)
		}

		return nil
	}
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

// test for purge_on_delete
func TestAccResourceDelegateTokenAccountAndProjectLevel(t *testing.T) {
	orgID := fmt.Sprintf("o%s", utils.RandStringBytes(5))
	projectID := fmt.Sprintf("p%s", utils.RandStringBytes(5))
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")

	accountLevelTokenName := fmt.Sprintf("acctok-%s", utils.RandStringBytes(5))
	projectLevelTokenName := fmt.Sprintf("projtok-%s", utils.RandStringBytes(5))

	accountLevelResourceName := "harness_platform_delegatetoken.account_level"
	projectLevelResourceName := "harness_platform_delegatetoken.project_level"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: testDelegateTokensDestroy(
			accountLevelResourceName,
			projectLevelResourceName,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDelegateTokenAccountAndProjectLevel(orgID, projectID, accountID, accountLevelTokenName, projectLevelTokenName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(accountLevelResourceName, "name", accountLevelTokenName),
					resource.TestCheckResourceAttr(accountLevelResourceName, "token_status", "ACTIVE"),
					resource.TestCheckResourceAttr(accountLevelResourceName, "purge_on_delete", "true"),
					resource.TestCheckResourceAttr(projectLevelResourceName, "name", projectLevelTokenName),
					resource.TestCheckResourceAttr(projectLevelResourceName, "token_status", "ACTIVE"),
				),
			},
		},
	})
}

func testAccResourceDelegateTokenAccountAndProjectLevel(orgID string, projectID string, accountID string, accountLevelTokenName string, projectLevelTokenName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[2]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			color      = "#472848"
		}

		resource "harness_platform_delegatetoken" "account_level" {
			name             = "%[4]s"
			account_id        = "%[3]s"
			purge_on_delete  = true
		}

		resource "harness_platform_delegatetoken" "project_level" {
			name       = "%[5]s"
			account_id  = "%[3]s"
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
		}
	`, orgID, projectID, accountID, accountLevelTokenName, projectLevelTokenName)
}

func testAccGetResourceDelegateTokenForDestroy(resourceName string, state *terraform.State) (*nextgen.DelegateTokenDetails, error) {
	rm := state.RootModule()
	if rm == nil {
		return nil, nil
	}
	r, ok := rm.Resources[resourceName]
	if !ok || r == nil {
		return nil, nil
	}

	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	resp, httpResp, err := c.DelegateTokenResourceApi.GetCgDelegateTokens(ctx, c.AccountId, &nextgen.DelegateTokenResourceApiGetCgDelegateTokensOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
		Name:              buildField(r, "name"),
		Status:            optional.EmptyString(),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, err
	}
	if resp.Resource == nil || len(resp.Resource) == 0 {
		return nil, nil
	}

	return &resp.Resource[0], nil
}

func testDelegateTokensDestroy(resourceNames ...string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for _, resourceName := range resourceNames {
			token, err := testAccGetResourceDelegateTokenForDestroy(resourceName, state)
			if err != nil {
				return err
			}
			if token != nil && token.Status != "REVOKED" {
				return fmt.Errorf("Token is not revoked : %s", token.Name)
			}
		}

		return nil
	}
}
