package apiKey_test

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

func TestAccResourceApiKey(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := apiKeyName + "updated"

	resourceName := "harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApiKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApiKey(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
				),
			},
			{
				Config: testAccResourceApiKey(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})

}

func TestAccResourceApiKeyOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := apiKeyName + "updated"

	resourceName := "harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApiKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceApiKey(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
				),
			},
			{
				Config: testOrgResourceApiKey(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})

}

func TestAccResourceApiKeyProjectLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := apiKeyName + "updated"

	resourceName := "harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApiKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceApiKey(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
				),
			},
			{
				Config: testProjectResourceApiKey(id, updatedName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})

}

func testAccGetResourceApiKey(resourceName string, state *terraform.State) (*nextgen.ApiKey, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.ApiKeyApi.GetAggregatedApiKey(ctx, c.AccountId, r.Primary.Attributes["apikey_type"], r.Primary.Attributes["parent_id"], id, &nextgen.ApiKeyApiGetAggregatedApiKeyOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data.ApiKey, nil
}

func testAccResourceApiKey(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}
	`, id, name, parentId, accountId)
}

func testOrgResourceApiKey(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}
	`, id, name, parentId, accountId)
}

func testProjectResourceApiKey(id string, name string, parentId string, accountId string) string {
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

		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}
	`, id, name, parentId, accountId)
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccApiKeyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		apiKey, _ := testAccGetResourceApiKey(resourceName, state)
		if apiKey != nil {
			return fmt.Errorf("Found apiKey: %s", apiKey.Identifier)
		}

		return nil
	}
}
