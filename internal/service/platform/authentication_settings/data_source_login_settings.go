package authentication_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceLoginSettings() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving the login settings of a Harness account.",

		ReadContext: resourceLoginSettingsRead,

		Schema: map[string]*schema.Schema{
			"password_strength_policy": {
				Description: "Password strength policy of the account.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Description: "Whether the password strength policy is enabled.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"min_number_of_characters": {
							Description: "Minimum number of characters required in a password.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"min_number_of_uppercase_characters": {
							Description: "Minimum number of uppercase characters required in a password.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"min_number_of_lowercase_characters": {
							Description: "Minimum number of lowercase characters required in a password.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"min_number_of_special_characters": {
							Description: "Minimum number of special characters required in a password.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"min_number_of_digits": {
							Description: "Minimum number of digits required in a password.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"user_lockout_policy": {
				Description: "User lockout policy of the account.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_lockout_policy": {
							Description: "Whether the lockout policy is enabled.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"number_of_failed_attempts_before_lockout": {
							Description: "Number of failed login attempts before a user is locked out.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"lockout_period": {
							Description: "Lockout period in hours.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"notify_user": {
							Description: "Whether the user is notified about the lockout.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
			"password_expiration_policy": {
				Description: "Password expiration policy of the account.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Description: "Whether the password expiration policy is enabled.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"days_before_password_expires": {
							Description: "Number of days before a password expires.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"days_before_user_notified_of_password_expiration": {
							Description: "Number of days before expiration the user is notified.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}
