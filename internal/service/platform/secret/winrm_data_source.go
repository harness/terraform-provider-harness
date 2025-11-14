package secret

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretWinRM() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for looking up a WinRM credential secret.",
		ReadContext: resourceSecretWinRMRead,

		Schema: map[string]*schema.Schema{
			"port": {
				Description: "WinRM port. Default is 5986 for HTTPS, 5985 for HTTP.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"ntlm": {
				Description: "NTLM authentication scheme",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Description: "Domain name for NTLM authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_ssl": {
							Description: "Use SSL/TLS for WinRM communication.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"skip_cert_check": {
							Description: "Skip certificate verification.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"use_no_profile": {
							Description: "Use no profile.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
			"kerberos": {
				Description: "Kerberos authentication scheme",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"principal": {
							Description: "Kerberos principal.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"realm": {
							Description: "Kerberos realm.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tgt_generation_method": {
							Description: "Method to generate TGT (Ticket Granting Ticket).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_ssl": {
							Description: "Use SSL/TLS for WinRM communication.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"skip_cert_check": {
							Description: "Skip certificate verification.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"use_no_profile": {
							Description: "Use no profile.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"tgt_key_tab_file_path_spec": {
							Description: "TGT generation using key tab file.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_path": {
										Description: "Path to the key tab file.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"tgt_password_spec": {
							Description: "TGT generation using password.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password_ref": {
										Description: "Reference to a secret containing the password.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
