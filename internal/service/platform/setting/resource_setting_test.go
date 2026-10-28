package setting_test

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

const (
	testGitRepoSetting  = "default_repo_for_git_experience"
	testGitRepoCategory = "GIT_EXPERIENCE"
	// A boolean setting that only changes what the UI displays, so setting it briefly has no
	// side effects on the test account.
	testBoolSetting  = "display_raw_mode_setting"
	testBoolCategory = "PMS"
)

func TestAccResourceSetting_ProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingProjectLevel(id, "repo-one", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "default_repo_for_git_experience"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "value", "repo-one"),
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "true"),
					resource.TestCheckResourceAttr(resourceName, "category", "GIT_EXPERIENCE"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "String"),
				),
			},
			{
				Config: testAccResourceSettingProjectLevel(id, "repo-two", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", "repo-two"),
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

func TestAccResourceSetting_OrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingOrgLevel(id, "org-repo", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "value", "org-repo"),
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceSettingProjectLevel(id string, value string, allowOverrides bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_setting" "test" {
			identifier      = "default_repo_for_git_experience"
			org_id          = harness_platform_organization.test.id
			project_id      = harness_platform_project.test.id
			value           = "%[2]s"
			allow_overrides = %[3]t
		}
	`, id, value, allowOverrides)
}

func testAccResourceSettingOrgLevel(id string, value string, allowOverrides bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_setting" "test" {
			identifier      = "default_repo_for_git_experience"
			org_id          = harness_platform_organization.test.id
			value           = "%[2]s"
			allow_overrides = %[3]t
		}
	`, id, value, allowOverrides)
}

func TestAccResourceSetting_AccountLevel(t *testing.T) {
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			// Destroy resets the setting, so don't run if the account already has its own value.
			if setting := testAccGetSetting(t, testBoolSetting, testBoolCategory, "", ""); setting.SettingSource == "ACCOUNT" {
				t.Skipf("account already sets %s; skipping so the test doesn't reset it", testBoolSetting)
			}
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingAccountLevel("true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", testBoolSetting),
					resource.TestCheckNoResourceAttr(resourceName, "org_id"),
					resource.TestCheckNoResourceAttr(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "value", "true"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Boolean"),
				),
			},
			{
				Config: testAccResourceSettingAccountLevel("false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if setting := testAccGetSetting(t, testBoolSetting, testBoolCategory, "", ""); setting.SettingSource == "ACCOUNT" {
				return fmt.Errorf("setting %s is still set at account level after destroy", testBoolSetting)
			}
			return nil
		},
	})
}

func TestAccResourceSetting_DisallowOverridesRemovesChildOverride(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingOrgWithProject(id, "org-repo", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "true"),
				),
			},
			{
				// Override the setting in the project outside Terraform, then disallow overrides at org level.
				PreConfig: func() {
					testAccSetSetting(t, testGitRepoSetting, "project-repo", id, id)
					if setting := testAccGetSetting(t, testGitRepoSetting, testGitRepoCategory, id, id); setting.SettingSource != "PROJECT" {
						t.Fatalf("expected a project override before disallowing overrides, got source %s", setting.SettingSource)
					}
				},
				Config: testAccResourceSettingOrgWithProject(id, "org-repo", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "false"),
					func(_ *terraform.State) error {
						setting := testAccGetSetting(t, testGitRepoSetting, testGitRepoCategory, id, id)
						if setting.SettingSource != "ORG" || setting.Value != "org-repo" {
							return fmt.Errorf("expected project to inherit org value org-repo, got %q from %s", setting.Value, setting.SettingSource)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccResourceSetting_BooleanValue(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			// The API rejects project values when the account (or the setting's default) disallows overrides.
			if setting := testAccGetSetting(t, testBoolSetting, testBoolCategory, "", ""); !setting.AllowOverrides {
				t.Skipf("%s doesn't allow overrides on this account; enable \"Allow overrides\" for it in Account Settings > Default Settings to run this test", testBoolSetting)
			}
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingBooleanProjectLevel(id, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", testBoolSetting),
					resource.TestCheckResourceAttr(resourceName, "value", "true"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Boolean"),
					resource.TestCheckResourceAttr(resourceName, "category", testBoolCategory),
				),
			},
			{
				Config: testAccResourceSettingBooleanProjectLevel(id, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", "false"),
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

// testAccGetSetting reads a setting at the given scope through the settings API.
func testAccGetSetting(t *testing.T, identifier, category, org, project string) nextgen.SettingDto {
	t.Helper()
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	opts := &nextgen.SettingApiGetSettingsListOpts{}
	if org != "" {
		opts.OrgIdentifier = optional.NewString(org)
	}
	if project != "" {
		opts.ProjectIdentifier = optional.NewString(project)
	}
	resp, _, err := c.SettingApi.GetSettingsList(ctx, c.AccountId, category, opts)
	if err != nil {
		t.Fatalf("listing %s settings: %v", category, err)
	}
	for _, item := range resp.Data {
		if item.Setting != nil && item.Setting.Identifier == identifier {
			return *item.Setting
		}
	}
	t.Fatalf("setting %s not found in category %s", identifier, category)
	return nextgen.SettingDto{}
}

// testAccSetSetting sets a value at the given scope through the settings API, outside Terraform.
func testAccSetSetting(t *testing.T, identifier, value, org, project string) {
	t.Helper()
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	resp, _, err := c.SettingApi.UpdateSettingValue(ctx, []nextgen.SettingRequestDto{{
		Identifier:     identifier,
		Value:          value,
		AllowOverrides: true,
		UpdateType:     "UPDATE",
	}}, c.AccountId, &nextgen.SettingApiUpdateSettingValueOpts{
		OrgIdentifier:     optional.NewString(org),
		ProjectIdentifier: optional.NewString(project),
	})
	if err != nil {
		t.Fatalf("setting %s outside Terraform: %v", identifier, err)
	}
	for _, item := range resp.Data {
		if item.Identifier == identifier && !item.UpdateStatus {
			t.Fatalf("setting %s outside Terraform: %s", identifier, item.ErrorMessage)
		}
	}
}

func testAccResourceSettingAccountLevel(value string) string {
	return fmt.Sprintf(`
		resource "harness_platform_setting" "test" {
			identifier = "%[1]s"
			value      = "%[2]s"
		}
	`, testBoolSetting, value)
}

func testAccResourceSettingOrgWithProject(id string, value string, allowOverrides bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_setting" "test" {
			identifier      = "%[4]s"
			org_id          = harness_platform_organization.test.id
			value           = "%[2]s"
			allow_overrides = %[3]t

			depends_on = [harness_platform_project.test]
		}
	`, id, value, allowOverrides, testGitRepoSetting)
}

func testAccResourceSettingBooleanProjectLevel(id string, value string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_setting" "test" {
			identifier = "%[3]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			value      = "%[2]s"
		}
	`, id, value, testBoolSetting)
}
