package manual_freeze_test

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
	"github.com/stretchr/testify/require"
)

func TestAccResourceManualFreezeWithRecurrence(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreeze(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Daily"),
				),
			},
			{
				Config: testAccResourceManualFreeze(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Daily"),
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

func TestAccResourceManualFreezeWithQuarterlyRecurrence(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreezeQuarterly(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Monthly"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.value", "3"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.until", "2023-12-31 11:59 PM"),
				),
			},
			{
				Config: testAccResourceManualFreezeQuarterly(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Monthly"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.value", "3"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.until", "2023-12-31 11:59 PM"),
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

func TestAccResourceManualFreeze_WithoutRecurrence(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreezeWithoutRecurrence(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
				),
			},
			{
				Config: testAccResourceManualFreezeWithoutRecurrence(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
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

func TestAccResourceManualFreeze_WithoutRecurrenceForNonExpiredWindows(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreezeWithoutRecurrenceNonExpired(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
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

func TestAccResourceManualFreeze_DeleteUnderlyingResource(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreeze(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Daily"),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, err := c.FreezeCRUDApi.DeleteFreeze(ctx, c.AccountId, id, &nextgen.FreezeCRUDApiDeleteFreezeOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testAccResourceManualFreeze(id, name, accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceQuarterlyManualFreeze_DeleteUnderlyingResource(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_manual_freeze.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManualFreezeQuarterly(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Monthly"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.value", "3"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.until", "2023-12-31 11:59 PM"),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, err := c.FreezeCRUDApi.DeleteFreeze(ctx, c.AccountId, id, &nextgen.FreezeCRUDApiDeleteFreezeOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testAccResourceManualFreeze(id, name, accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		manualFreeze, _ := testAccGetManualFreeze(resourceName, state)
		if manualFreeze != nil {
			return fmt.Errorf("Found freeze: %s", manualFreeze.Identifier)
		}
		return nil
	}
}

func testAccGetManualFreeze(resourceName string, state *terraform.State) (*nextgen.FreezeDetailedResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.FreezeCRUDApi.GetFreeze(ctx, c.AccountId, id, &nextgen.FreezeCRUDApiGetFreezeOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceManualFreeze(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[2]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2023-05-03 04:16 PM
          duration: 30m
          recurrence:
            type: Daily
        notificationRules: []
        tags: {}
      EOT
		}
	`, id, name, accountId)
}

func testAccResourceManualFreezeQuarterly(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[2]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2023-05-03 04:16 PM
          duration: 30m
          recurrence:
            type: Monthly
            spec:
              value: 3
              until: 2023-12-31 11:59 PM
        notificationRules: []
        tags: {}
      EOT
		}
	`, id, name, accountId)
}

func testAccResourceManualFreezeWithoutRecurrence(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[2]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2025-05-07 04:16 PM
          duration: 30m
        notificationRules: []
        tags: {}
      EOT
		}
	`, id, name, accountId)
}

func testAccResourceManualFreezeWithoutRecurrenceNonExpired(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[2]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2026-05-03 04:16 PM
          duration: 30m
        notificationRules: []
        tags: {}
      EOT
		}
	`, id, name, accountId)
}
