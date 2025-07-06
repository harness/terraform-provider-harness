package notification_channel_test

import "fmt"

import (
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceNotificationChannel_basic(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationChannelBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.email_ids.0", "notify@harness.io"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.api_key", "dummy-api-key"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.execute_on_delegate", "true"),
				),
			},
		},
	})
}
func testAccNotificationChannelBasic(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_notification_channel" "test" {
 identifier                = "%[1]s"
 name                      = "%[2]s"
 notification_channel_type = "EMAIL"
 status                    = "ENABLED"

 channel {
   email_ids            = ["notify@harness.io"]
   api_key              = "dummy-api-key"
   execute_on_delegate  = true
   user_groups {
     identifier = "account.test"
   }
   headers {
     key   = "X-Custom-Header"
     value = "HeaderValue"
   }
 }
}
`, id, name)
}

func testAccGetNotificationChannel(resourceName string, state *terraform.State) (*nextgen.NotificationChannelDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	resp, _, err := c.NotificationChannelsApi.GetNotificationChannel(ctx, id, buildField(r, "org_id").Value(), buildField(r, "project_id").Value(),
		&nextgen.NotificationChannelsApiGetNotificationChannelOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
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

func testAccCheckNotificationChannelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		notificationRule, _ := testAccGetNotificationChannel(resourceName, state)
		if notificationRule != nil {
			return fmt.Errorf("Found notification channel: %s", notificationRule.Identifier)
		}

		return nil
	}
}
