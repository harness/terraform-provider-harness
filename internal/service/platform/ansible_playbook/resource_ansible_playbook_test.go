package ansible_playbook_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAnsiblePlaybook(t *testing.T) {
	resourceName := "harness_platform_iacm_ansible_playbook.test"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceAnsiblePlaybookDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAnsiblePlaybook(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "repository_path", "ansible/site.yml"),
					resource.TestCheckResourceAttr(resourceName, "ansible_galaxy", "true"),
				),
			},
			{
				Config: testAccResourceAnsiblePlaybook(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceAnsiblePlaybookDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		pb, _ := testAccGetPlatformAnsiblePlaybook(resourceName, state)
		if pb != nil {
			return fmt.Errorf("Ansible playbook found: %s", pb.Identifier)
		}
		return nil
	}
}

func testAccGetPlatformAnsiblePlaybook(resourceName string, state *terraform.State) (*nextgen.ShowPlaybookResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	org := r.Primary.Attributes["org_id"]
	project := r.Primary.Attributes["project_id"]

	pb, resp, err := c.AnsibleApi.AnsibleShowPlaybook(ctx, org, project, id, c.AccountId)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return &pb, nil
}

func testAccResourceAnsiblePlaybook(id, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier                = "%[1]s"
			name                      = "%[2]s"
			secret_manager_identifier = "harnessSecretManager"
			value_type                = "Inline"
			value                     = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier       = "%[1]s"
			name             = "%[2]s"
			url              = "https://github.com/account"
			connection_type  = "Account"
			validation_repo  = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username  = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

		resource "harness_platform_iacm_ansible_playbook" "test" {
			identifier           = "%[1]s"
			name                 = "%[2]s"
			org_id               = harness_platform_organization.test.id
			project_id           = harness_platform_project.test.id
			repository           = "https://github.com/org/repo"
			repository_branch    = "main"
			repository_path      = "ansible/site.yml"
			repository_connector = "account.${harness_platform_connector_github.test.id}"
			ansible_galaxy       = true

			vars {
				key        = "env"
				value      = "prod"
				value_type = "string"
			}

			env_vars {
				key        = "ANSIBLE_CONFIG"
				value      = "ansible.cfg"
				value_type = "string"
			}
		}
	`, id, name)
}
