package token_test

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
