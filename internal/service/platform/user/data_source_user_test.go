package user_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func fetchEnvironmentFromEndpoint(endpoint string) (string, error) {

	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	hostParts := strings.Split(parsedURL.Hostname(), ".")
	if len(hostParts) < 2 {
		return "", fmt.Errorf("unexpected URL format: %s", parsedURL.Hostname())
	}

	env := hostParts[0]

	log.Printf("Fetched environment: %s", env)

	return env, nil
}

func TestAccDataSourceUserProjectLevel(t *testing.T) {
	endpoint := os.Getenv("HARNESS_ENDPOINT")

	if endpoint == "" {
		t.Fatal("HARNESS_ENDPOINT environment variable is not set")
	}

	env, err := fetchEnvironmentFromEndpoint(endpoint)
	if err != nil {
		t.Fatalf("Error fetching environment: %v", err)
	}

	log.Printf("Environment fetched from endpoint: %s", env)

	if !strings.EqualFold(env, "QA") {
		log.Printf("Skipping test as the environment is not QA (found: %s)", env)
		t.Skip("Skipping test because environment is not QA")
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "data.harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserProjectLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", email),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "externally_managed", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceUserAccountLevel(t *testing.T) {
	endpoint := os.Getenv("HARNESS_ENDPOINT")

	if endpoint == "" {
		t.Fatal("HARNESS_ENDPOINT environment variable is not set")
	}

	env, err := fetchEnvironmentFromEndpoint(endpoint)
	if err != nil {
		t.Fatalf("Error fetching environment: %v", err)
	}

	log.Printf("Environment fetched from endpoint: %s", env)

	if !strings.EqualFold(env, "QA") {
		log.Printf("Skipping test as the environment is not QA (found: %s)", env)
		t.Skip("Skipping test because environment is not QA")
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "data.harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserAccountLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", email),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "externally_managed", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceUserOrgLevel(t *testing.T) {
	endpoint := os.Getenv("HARNESS_ENDPOINT")

	if endpoint == "" {
		t.Fatal("HARNESS_ENDPOINT environment variable is not set")
	}

	env, err := fetchEnvironmentFromEndpoint(endpoint)
	if err != nil {
		t.Fatalf("Error fetching environment: %v", err)
	}

	log.Printf("Environment fetched from endpoint: %s", env)

	if !strings.EqualFold(env, "QA") {
		log.Printf("Skipping test as the environment is not QA (found: %s)", env)
		t.Skip("Skipping test because environment is not QA")
	}

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "data.harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserOrgLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", email),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "externally_managed", "false"),
				),
			},
		},
	})
}

func testAccDataSourceUserProjectLevel(id string, name string, email string) string {
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

	resource "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		email = "%[3]s"
		user_groups = ["_project_all_users"]
		role_bindings {
			resource_group_identifier = "_all_project_level_resources"
			role_identifier = "_project_viewer"
			role_name = "Project Viewer"
			resource_group_name = "All Project Level Resources"
			managed_role = true
		}
	}

	data "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		email = "%[3]s"
		depends_on = [harness_platform_user.test]
	}
	`, id, name, email)
}

func testAccDataSourceUserAccountLevel(id string, name string, email string) string {
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

	resource "harness_platform_user" "test" {
		email = "%[3]s"
		user_groups = ["_project_all_users"]
		role_bindings {
			resource_group_identifier = "_all_project_level_resources"
			role_identifier = "_project_viewer"
			role_name = "Project Viewer"
			resource_group_name = "All Project Level Resources"
			managed_role = true
		}
	}

	data "harness_platform_user" "test" {
		email = "%[3]s"
		depends_on = [harness_platform_user.test]
	}
	`, id, name, email)
}

func testAccDataSourceUserOrgLevel(id string, name string, email string) string {
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

	resource "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		email = "%[3]s"
		user_groups = ["_project_all_users"]
		role_bindings {
			resource_group_identifier = "_all_project_level_resources"
			role_identifier = "_project_viewer"
			role_name = "Project Viewer"
			resource_group_name = "All Project Level Resources"
			managed_role = true
		}
	}

	data "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		email = "%[3]s"
		depends_on = [harness_platform_user.test]
	}
	`, id, name, email)
}
