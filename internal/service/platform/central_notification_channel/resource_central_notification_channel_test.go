package central_notification_channel_test

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

// EMAIL Channel Tests

func TestAccCentralNotificationChannel_Email(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelEmailAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.email_ids.0", "test@harness.io"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_Email(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelEmailOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_Email(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelEmailProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_Slack(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelSlackOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_Slack(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelSlackProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_MSTeams(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelMSTeamsOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "MSTEAMS"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_MSTeams(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelMSTeamsProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "MSTEAMS"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_PagerDuty(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelPagerDutyOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "PAGERDUTY"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_PagerDuty(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelPagerDutyProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "PAGERDUTY"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_Webhook(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelWebhookOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "WEBHOOK"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_Webhook(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelWebhookProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "WEBHOOK"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

// DATADOG Channel Tests

func TestAccCentralNotificationChannel_Datadog(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelDatadogAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "DATADOG"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.datadog_urls.0", "https://api.datadoghq.com/api/v1/events"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				),
			},
		},
	})
}

func TestAccCentralNotificationChannel_Slack(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelSlackAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.slack_webhook_urls.0", "https://hooks.slack.com/services/test"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccCentralNotificationChannel_MSTeams(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelMSTeamsAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "MSTEAMS"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.ms_team_keys.0", "https://outlook.office.com/webhook/test"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccCentralNotificationChannel_PagerDuty(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelPagerDutyAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "PAGERDUTY"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.pager_duty_integration_keys.0", "test-integration-key"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccCentralNotificationChannel_Webhook(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelWebhookAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "WEBHOOK"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.webhook_urls.0", "https://webhook.example.com/test"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestOrgCentralNotificationChannel_Datadog(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelDatadogOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "DATADOG"),
					resource.TestCheckResourceAttr(resourceName, "org", "default"),
				),
			},
		},
	})
}

func TestProjectCentralNotificationChannel_Datadog(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelDatadogProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "DATADOG"),
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
				),
			},
		},
	})
}

// Test Configuration Functions

// EMAIL Configurations
func testAccCentralNotificationChannelEmailAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "EMAIL"
	status                    = "ENABLED"

	channel {
		email_ids = ["test@harness.io"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelEmailOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "EMAIL"
	status                    = "ENABLED"

	channel {
		email_ids = ["test@harness.io"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelEmailProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "EMAIL"
	status                    = "ENABLED"

	channel {
		email_ids = ["test@harness.io"]
	}
}
`, id, name)
}

// SLACK Configurations
func testAccCentralNotificationChannelSlackAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "SLACK"
	status                    = "ENABLED"

	channel {
		slack_webhook_urls = ["https://hooks.slack.com/services/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelSlackOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "SLACK"
	status                    = "ENABLED"

	channel {
		slack_webhook_urls = ["https://hooks.slack.com/services/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelSlackProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "SLACK"
	status                    = "ENABLED"

	channel {
		slack_webhook_urls = ["https://hooks.slack.com/services/test"]
	}
}
`, id, name)
}

// MSTEAMS Configurations
func testAccCentralNotificationChannelMSTeamsAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "MSTEAMS"
	status                    = "ENABLED"

	channel {
		ms_team_keys = ["https://outlook.office.com/webhook/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelMSTeamsOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "MSTEAMS"
	status                    = "ENABLED"

	channel {
		ms_team_keys = ["https://outlook.office.com/webhook/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelMSTeamsProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "MSTEAMS"
	status                    = "ENABLED"

	channel {
		ms_team_keys = ["https://outlook.office.com/webhook/test"]
	}
}
`, id, name)
}

// PAGERDUTY Configurations
func testAccCentralNotificationChannelPagerDutyAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "PAGERDUTY"
	status                    = "ENABLED"

	channel {
		pager_duty_integration_keys = ["test-integration-key"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelPagerDutyOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "PAGERDUTY"
	status                    = "ENABLED"

	channel {
		pager_duty_integration_keys = ["test-integration-key"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelPagerDutyProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "PAGERDUTY"
	status                    = "ENABLED"

	channel {
		pager_duty_integration_keys = ["test-integration-key"]
	}
}
`, id, name)
}

// WEBHOOK Configurations
func testAccCentralNotificationChannelWebhookAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "WEBHOOK"
	status                    = "ENABLED"

	channel {
		webhook_urls = ["https://webhook.example.com/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelWebhookOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "WEBHOOK"
	status                    = "ENABLED"

	channel {
		webhook_urls = ["https://webhook.example.com/test"]
	}
}
`, id, name)
}

func testAccCentralNotificationChannelWebhookProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "WEBHOOK"
	status                    = "ENABLED"

	channel {
		webhook_urls = ["https://webhook.example.com/test"]
	}
}
`, id, name)
}

// DATADOG Configurations
func testAccCentralNotificationChannelDatadogAccount(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	name                      = "%[2]s"
	notification_channel_type = "DATADOG"
	status                    = "ENABLED"

	channel {
		datadog_urls = ["https://api.datadoghq.com/api/v1/events"]
		api_key      = "test-api-key"
	}
}
`, id, name)
}

func testAccCentralNotificationChannelDatadogOrg(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
	identifier                = "%[1]s"
	org                       = "default"
	name                      = "%[2]s"
	notification_channel_type = "DATADOG"
	status                    = "ENABLED"

	channel {
		datadog_urls = ["https://api.datadoghq.com/api/v1/events"]
		api_key      = "test-api-key"
	}
}
`, id, name)
}

func testAccCentralNotificationChannelDatadogProject(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
}

resource "harness_platform_project" "test" {
	identifier = "%[1]s"
	name       = "%[2]s"
	org_id     = harness_platform_organization.test.id
	color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
	depends_on = [
		harness_platform_organization.test,
		harness_platform_project.test,
	]
	identifier                = "%[1]s"
	org                       = harness_platform_organization.test.id
	project                   = harness_platform_project.test.id
	name                      = "%[2]s"
	notification_channel_type = "DATADOG"
	status                    = "ENABLED"

	channel {
		datadog_urls = ["https://api.datadoghq.com/api/v1/events"]
		api_key      = "test-api-key"
	}
}
`, id, name)
}

// Helper Functions
func testAccGetCentralNotificationChannel(resourceName string, state *terraform.State) (*nextgen.NotificationChannelDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	org := buildField(r, "org").Value()
	project := buildField(r, "project").Value()

	var resp nextgen.NotificationChannelDto
	var err error

	if org != "" && project != "" {
		// Project level
		resp, _, err = c.NotificationChannelsApi.GetNotificationChannel(ctx, id, org, project,
			&nextgen.NotificationChannelsApiGetNotificationChannelOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	} else if org != "" {
		// Org level
		resp, _, err = c.NotificationChannelsApi.GetNotificationChannelOrg(ctx, id, org,
			&nextgen.NotificationChannelsApiGetNotificationChannelOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	} else {
		// Account level
		resp, _, err = c.NotificationChannelsApi.GetNotificationChannelAccount(ctx, id,
			&nextgen.NotificationChannelsApiGetNotificationChannelAccountOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	}

	if err != nil {
		return nil, err
	}
	if resp.Channel == nil {
		return nil, fmt.Errorf("empty resource received in response")
	}

	return &resp, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccCheckCentralNotificationChannelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		notificationChannel, err := testAccGetCentralNotificationChannel(resourceName, state)
		if err != nil {
			// If we get any error (including 404 not found), that means resource is deleted
			// This is expected behavior after destroy
			return nil
		}
		if notificationChannel != nil {
			return fmt.Errorf("found notification channel: %s", notificationChannel.Identifier)
		}

		return nil
	}
}
