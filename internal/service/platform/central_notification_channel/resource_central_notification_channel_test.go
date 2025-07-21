package central_notification_channel_test

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

func TestAccResourceCentralNotificationChannel_basic(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationChannelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCentralNotificationChannelBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.email_ids.0", "notify@harness.io"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.execute_on_delegate", "true"),
				),
			},
		},
	})
}
func testAccCentralNotificationChannelBasic(id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
 identifier                = "%[1]s"
 name                      = "%[2]s"
 notification_channel_type = "EMAIL"
 status                    = "ENABLED"

 channel {
   email_ids            = ["notify@harness.io"]
   execute_on_delegate  = true
   user_groups {
     identifier = "account.test"
   }
 }
}
`, id, name)
}

func testAccGetCentralNotificationChannel(resourceName string, state *terraform.State) (*nextgen.NotificationChannelDto, error) {
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

func testAccCheckCentralNotificationChannelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		notificationRule, _ := testAccGetCentralNotificationChannel(resourceName, state)
		if notificationRule != nil {
			return fmt.Errorf("Found notification channel: %s", notificationRule.Identifier)
		}

		return nil
	}
}
