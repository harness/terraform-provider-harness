package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMEEnvironment_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME environment acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	res := "harness_fme_environment.test"
	var envID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironment(id, envName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", envName),
					resource.TestCheckResourceAttr(res, "production", "false"),
					resource.TestCheckResourceAttrSet(res, "environment_id"),
					resource.TestCheckResourceAttrPair(res, "id", res, "environment_id"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				Config: testAccResourceFMEEnvironment(id, envName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "production", "true"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				ResourceName:            res,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       fmeImportStateIDOrgProjectThird(res, "environment_id"),
				ImportStateVerifyIgnore: []string{"bootstrap_api_token_ids"},
				Check:                   testAccFMECaptureAttr(res, "environment_id", &envID),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyEnvironmentGone(id, id, envID),
			},
		},
	})
}

func TestAccResourceFMEEnvironment_changePermissions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME environment change_permissions acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	res := "harness_fme_environment.test"
	var envID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironmentWithPermissions(id, envName, false, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", envName),
					resource.TestCheckResourceAttr(res, "production", "false"),
					resource.TestCheckResourceAttrSet(res, "environment_id"),
					resource.TestCheckResourceAttr(res, "change_permissions.#", "1"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.allow_kills", "true"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.are_approvals_required", "true"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.are_approvers_restricted", "true"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approvers.#", "1"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approvers.0.type", "group"),
					resource.TestCheckResourceAttrSet(res, "change_permissions.0.approvers.0.id"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approval_skippable_by.#", "1"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approval_skippable_by.0.type", "group"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				Config: testAccResourceFMEEnvironmentWithPermissions(id, envName, false, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "change_permissions.0.allow_kills", "false"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approvers.#", "1"),
					resource.TestCheckResourceAttr(res, "change_permissions.0.approval_skippable_by.#", "1"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				ResourceName:            res,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       fmeImportStateIDOrgProjectThird(res, "environment_id"),
				ImportStateVerifyIgnore: []string{"bootstrap_api_token_ids", "change_permissions"},
				Check:                   testAccFMECaptureAttr(res, "environment_id", &envID),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyEnvironmentGone(id, id, envID),
			},
		},
	})
}

func testAccResourceFMEEnvironmentWithPermissions(id, envName string, production, allowKills bool) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	resource "harness_platform_usergroup" "approvers" {
		identifier = "%[1]s_approvers"
		name       = "%[1]s_approvers"
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}

	resource "harness_fme_environment" "test" {
		org_id       = harness_platform_organization.test.id
		project_id   = harness_platform_project.test.id
		name         = "%[2]s"
		production   = %[3]t

		change_permissions {
			allow_kills              = %[4]t
			are_approvals_required   = true
			are_approvers_restricted = true

			approvers {
				id   = harness_platform_usergroup.approvers.id
				name = harness_platform_usergroup.approvers.name
				type = "group"
			}

			approval_skippable_by {
				id   = harness_platform_usergroup.approvers.id
				name = harness_platform_usergroup.approvers.name
				type = "group"
			}
		}
	}
	`, id, envName, production, allowKills)
}

func testAccResourceFMEEnvironment(id, envName string, production bool) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	resource "harness_fme_environment" "test" {
		org_id       = harness_platform_organization.test.id
		project_id   = harness_platform_project.test.id
		name         = "%[2]s"
		production   = %[3]t
	}
	`, id, envName, production)
}
