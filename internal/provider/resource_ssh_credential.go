package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/harness/graphql"
)

func resourceSSHCredential() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an encrypted text secret",
		CreateContext: resourceSSHCredentialCreate,
		ReadContext:   resourceSSHCredentialRead,
		UpdateContext: resourceSSHCredentialUpdate,
		DeleteContext: resourceSSHCredentialDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the encrypted text secret",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the encrypted text secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ssh_authentication": {
				Description: "Authentication method for SSH. Cannot be used if kerberos_authentication is specified. Only one of `inline_ssh`, `server_password`, or `ssh_key_file` should be set",
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Description: "The port to connect to",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"username": {
							Description: "The username to use when connecting to ssh",
							Type:        schema.TypeString,
							Required:    true,
						},
						"inline_ssh": {
							Description: "Inline SSH authentication configuration. Only ond of `passphrase_secret_id` or `ssh_key_file_id` should be used",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"passphrase_secret_id": {
										Description: "The id of the encrypted secret to use",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"ssh_key_file_id": {
										Description: "The id of the secret containing the SSH key",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"ssh_key_file": {
							Description: "Use ssh key file for authentication",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Description: "The path to the key file on the delegate",
										Type:        schema.TypeString,
										Required:    true,
									},
									"passphrase_secret_id": {
										Description: "The id of the secret containing the password to use for the ssh key",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"server_password": {
							Description: "Server password authentication configuration",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password_secret_id": {
										Description: "The id of the encrypted secret",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"kerberos_authentication": {
				Description: "Kerberos authentication for SSH. Cannot be used if ssh_authentication is specified",
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Description: "Port to use for Kerberos authentication",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"principal": {
							Description: "Name of the principal for authentication",
							Type:        schema.TypeString,
							Required:    true,
						},
						"realm": {
							Description: "Realm associated with the Kerberos authentication",
							Type:        schema.TypeString,
							Required:    true,
						},
						"tgt_generation_method": {
							Description: "TGT generation method",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kerberos_password_id": {
										Description: "The id of the encrypted text secret",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"key_tab_file_path": {
										Description: "The path to the key tab file",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"usage_scope": usageScopeSchema(),
		},
	}
}

func resourceSSHCredentialCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	input := &graphql.SSHCredential{
		Name: d.Get("name").(string),
	}

	if err := expandAuthenticationScheme(d, input); err != nil {
		return diag.FromErr(err)
	}

	usageScope, err := expandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.UsageScope = usageScope

	cred, err := c.Secrets().CreateSSHCredential(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cred.Id)
	d.Set("name", cred.Name)

	return nil
}

func resourceSSHCredentialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	credId := d.Get("id").(string)

	cred, err := c.Secrets().GetSSHCredentialById(credId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", cred.Name)
	d.Set("ssh_authentication", flattenSSHAuthentication(cred.SSHAuthentication))
	d.Set("kerberos_authentication", flattenKerberosAuthentication(cred.KerberosAuthentication))
	d.Set("usage_scope", flattenUsageScope(cred.UsageScope))

	return nil
}

func resourceSSHCredentialUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	input := &graphql.SSHCredential{
		Name: d.Get("name").(string),
	}

	if err := expandAuthenticationScheme(d, input); err != nil {
		return diag.FromErr(err)
	}

	usageScope, err := expandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.UsageScope = usageScope

	_, err = c.Secrets().UpdateSSHCredential(d.Get("id").(string), input)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSSHCredentialDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	err := c.Secrets().DeleteSecret(d.Get("id").(string), graphql.SecretTypes.SSHCredential)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenKerberosAuthentication(cred *graphql.KerberosAuthentication) []interface{} {
	response := make([]interface{}, 0)

	if cred == nil {
		return response
	}

	data := map[string]interface{}{}
	data["port"] = cred.Port
	data["principal"] = cred.Principal
	data["realm"] = cred.Realm
	data["tgt_generation_method"] = flattenTGTGenerationMethod(cred.TGTGenerationMethod)

	return response
}

func flattenTGTGenerationMethod(tgt *graphql.TGTGenerationMethod) []interface{} {
	response := make([]interface{}, 0)

	if tgt == nil {
		return response
	}

	data := map[string]interface{}{}
	if tgt.KeyTabFile != nil {
		data["key_tab_file_path"] = tgt.KeyTabFile.FilePath
	}

	if tgt.KerberosPassword != nil {
		data["kerberos_password_id"] = tgt.KerberosPassword.PasswordSecretId
	}

	return response
}

func flattenSSHAuthentication(cred *graphql.SSHAuthentication) []interface{} {
	response := make([]interface{}, 0)

	if cred == nil {
		return response
	}

	data := map[string]interface{}{}
	data["port"] = cred.Port
	data["username"] = cred.Username

	if cred.SSHAuthenticationMethod != nil {
		switch cred.SSHAuthenticationMethod.SSHCredentialType {
		case graphql.SSHCredentialTypes.Password:
			data["server_password"] = flattenSSHServerPasswordConfig(cred.SSHAuthenticationMethod.ServerPassword)
		case graphql.SSHCredentialTypes.SSHKeyFilePath:
			data["ssh_key_file"] = flattenSSHKeyFileConfig(cred.SSHAuthenticationMethod.SSHKeyFile)
		case graphql.SSHCredentialTypes.SSHKey:
			data["inline_ssh"] = flattenInlineSSHConfig(cred.SSHAuthenticationMethod.InlineSSHKey)
		}
	}

	return append(response, data)
}

func flattenInlineSSHConfig(cred *graphql.InlineSSHKey) []interface{} {
	response := make([]interface{}, 0)
	if cred == nil {
		return response
	}

	data := map[string]interface{}{}
	data["passphrase_secret_id"] = cred.PassphraseSecretId
	data["ssh_key_file_id"] = cred.SSHKeySecretFileId

	return append(response, data)
}

func flattenSSHKeyFileConfig(cred *graphql.SSHKeyFile) []interface{} {
	response := make([]interface{}, 0)
	if cred == nil {
		return response
	}

	data := map[string]interface{}{}
	data["path"] = cred.Path
	data["passphrase_secret_id"] = cred.PassphraseSecretId

	return append(response, data)
}

func flattenSSHServerPasswordConfig(cred *graphql.SSHPassword) []interface{} {

	response := make([]interface{}, 0)
	if cred == nil {
		return response
	}

	data := map[string]interface{}{}
	data["password_secret_id"] = cred.PasswordSecretId

	return append(response, data)
}

func expandAuthenticationScheme(d *schema.ResourceData, cred *graphql.SSHCredential) error {
	k := d.Get("kerberos_authentication").(*schema.Set).List()
	s := d.Get("ssh_authentication").(*schema.Set).List()

	hasKerberos := len(k) > 0
	hasSSHAuth := len(s) > 0

	if (hasKerberos && hasSSHAuth) || (!hasKerberos && !hasSSHAuth) {
		return errors.New("must specify only one of either `kerberos_authentication` or `ssh_authentication`")
	}

	if err := expandSSHAuthentication(s, cred); err != nil {
		return err
	}

	if err := expandKerberosAuthentication(k, cred); err != nil {
		return err
	}

	return nil
}

func expandSSHAuthentication(d []interface{}, cred *graphql.SSHCredential) error {

	if len(d) <= 0 {
		cred.SSHAuthentication = nil
		return nil
	}

	data := d[0].(map[string]interface{})

	inlineSSHConfig := data["inline_ssh"].(*schema.Set).List()
	keyFileConfig := data["ssh_key_file"].(*schema.Set).List()
	serverPassConfig := data["server_password"].(*schema.Set).List()

	totalConfigs := len(inlineSSHConfig) + len(keyFileConfig) + len(serverPassConfig)
	if totalConfigs > 1 {
		return errors.New("only one of `inline_ssh`, `ssh_key_file`, or `server_password` must be specified")
	} else if totalConfigs == 0 {
		// BUG: https://harness.atlassian.net/browse/SWAT-4653
		// lifecycle { ignore_changes = [ssh_authentication] } must be set
		// GraphQL doesn't send proper response
		return nil
	}

	auth := &graphql.SSHAuthentication{
		SSHAuthenticationMethod: &graphql.SSHAuthenticationMethod{},
	}

	if attr, ok := data["port"]; ok {
		auth.Port = attr.(int)
	}

	if attr, ok := data["username"]; ok && attr != "" {
		auth.Username = attr.(string)
	}

	expandInlineSSHConfig(inlineSSHConfig, auth)
	expandSSHKeyFileConfig(keyFileConfig, auth)
	expandServerPasswordConfig(serverPassConfig, auth)

	cred.AuthenticationScheme = graphql.SSHAuthenticationSchemes.SSH
	cred.SSHAuthentication = auth

	return nil
}

func expandServerPasswordConfig(d []interface{}, auth *graphql.SSHAuthentication) {
	if len(d) <= 0 {
		auth.SSHAuthenticationMethod.ServerPassword = nil
		return
	}

	serverPasswordConfig := &graphql.SSHPassword{}
	data := d[0].(map[string]interface{})

	if attr, ok := data["password_secret_id"]; ok && attr != "" {
		serverPasswordConfig.PasswordSecretId = attr.(string)
	}

	auth.SSHAuthenticationMethod.SSHCredentialType = graphql.SSHCredentialTypes.Password
	auth.SSHAuthenticationMethod.ServerPassword = serverPasswordConfig
}

func expandSSHKeyFileConfig(d []interface{}, auth *graphql.SSHAuthentication) {
	if len(d) <= 0 {
		auth.SSHAuthenticationMethod.SSHKeyFile = nil
		return
	}

	sshKeyConfig := &graphql.SSHKeyFile{}
	data := d[0].(map[string]interface{})

	if attr, ok := data["passphrase_secret_id"]; ok && attr != "" {
		sshKeyConfig.PassphraseSecretId = attr.(string)
	}

	if attr, ok := data["path"]; ok && attr != "" {
		sshKeyConfig.Path = attr.(string)
	}

	auth.SSHAuthenticationMethod.SSHCredentialType = graphql.SSHCredentialTypes.SSHKeyFilePath
	auth.SSHAuthenticationMethod.SSHKeyFile = sshKeyConfig
}

func expandInlineSSHConfig(d []interface{}, auth *graphql.SSHAuthentication) {
	if len(d) <= 0 {
		auth.SSHAuthenticationMethod.InlineSSHKey = nil
		return
	}

	inlineConfig := &graphql.InlineSSHKey{}
	data := d[0].(map[string]interface{})

	if attr, ok := data["passphrase_secret_id"]; ok && attr != "" {
		inlineConfig.PassphraseSecretId = attr.(string)
	}

	if attr, ok := data["ssh_key_file_id"]; ok && attr != "" {
		inlineConfig.SSHKeySecretFileId = attr.(string)
	}

	auth.SSHAuthenticationMethod.SSHCredentialType = graphql.SSHCredentialTypes.SSHKey
	auth.SSHAuthenticationMethod.InlineSSHKey = inlineConfig
}

func expandKerberosAuthentication(d []interface{}, cred *graphql.SSHCredential) error {

	if len(d) <= 0 {
		cred.KerberosAuthentication = nil
		return nil
	}

	data := d[0].(map[string]interface{})
	auth := &graphql.KerberosAuthentication{}

	if attr, ok := data["port"]; ok {
		auth.Port = attr.(int)
	}

	if attr, ok := data["principal"]; ok && attr != "" {
		auth.Principal = attr.(string)
	}

	if attr, ok := data["realm"]; ok && attr != "" {
		auth.Principal = attr.(string)
	}

	expandTGTGenerationMethod(data["tgt_generation_method"].(*schema.Set).List(), auth)
	if auth.TGTGenerationMethod.KerberosPassword != nil && auth.TGTGenerationMethod.KeyTabFile != nil {
		return errors.New("must set only one of `kerberos_password_id` or `key_tab_file_path` for tgt_generation_method")
	}

	cred.AuthenticationScheme = graphql.SSHAuthenticationSchemes.Kerberos
	cred.KerberosAuthentication = auth

	return nil
}

func expandTGTGenerationMethod(d []interface{}, auth *graphql.KerberosAuthentication) {
	if len(d) <= 0 {
		auth.TGTGenerationMethod = nil
		return
	}

	details := &graphql.TGTGenerationMethod{}
	data := d[0].(map[string]interface{})

	if attr, ok := data["kerberos_password_id"]; ok && attr != "" {
		details.TGTGenerationUsing = graphql.TGTGenerationUsingOptions.Password
		details.KerberosPassword = &graphql.KerberosPassword{
			PasswordSecretId: attr.(string),
		}
	}

	if attr, ok := data["key_tab_file_path"]; ok && attr != "" {
		details.TGTGenerationUsing = graphql.TGTGenerationUsingOptions.KeyTabFile
		details.KeyTabFile = &graphql.KeyTabFile{
			FilePath: attr.(string),
		}
	}
}
