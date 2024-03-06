package secret

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretWinRM() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for looking up an SSH Key type secret.",
		ReadContext: resourceSecretSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"port": {
				Description: "WinRM port",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"kerberos": {
				Description: "Kerberos authentication scheme",
				Type:        schema.TypeList,
				Computed:    true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"principal": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"realm": {
							Description: "Reference to a secret containing the password to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tgt_generation_method": {
							Description: "Method to generate tgt",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tgt_key_tab_file_path_spec": {
							Description: "Authenticate to App Dynamics using username and password.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_path": {
										Description: "key path",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"tgt_password_spec": {
							Description: "Authenticate to App Dynamics using username and password.",
							Type:        schema.TypeList,
							Computed:    true,

							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Description: "password",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"ssh": { // TODO: This is a copy of the SSH schema. It should be updated to reflect WinRM
				Description: "Kerberos authentication scheme",
				Type:        schema.TypeList,
				Computed:    true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credential_type": {
							Description: "This specifies SSH credential type as Password, KeyPath or KeyReference",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sshkey_path_credential": {
							Description: "SSH credential of type keyPath",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Description: "SSH Username.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"key_path": {
										Description: "Path of the key file.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"encrypted_passphrase": {
										Description: "Encrypted Passphrase",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"sshkey_reference_credential": {
							Description: "SSH credential of type keyReference",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Description: "SSH Username.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"key": {
										Description: "SSH key.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"encrypted_assphrase": {
										Description: "Encrypted Passphrase",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"ssh_password_credential": {
							Description: "SSH credential of type keyReference",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Description: "SSH Username.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password": {
										Description: "SSH Password.",
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
