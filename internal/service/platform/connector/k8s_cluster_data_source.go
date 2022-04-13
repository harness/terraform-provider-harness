package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/gitsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorKubernetes() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Kubernetes connector.",
		ReadContext: resourceConnectorK8sRead,

		Schema: map[string]*schema.Schema{
			"inherit_from_delegate": {
				Description: "Credentials are inherited from the delegate.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "Selectors to use for the delegate.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"username_password": {
				Description: "Username and password for the connector.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username": {
							Description: "Username for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_ref": {
							Description: "Reference to the secret containing the username for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "Reference to the secret containing the password for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"service_account": {
				Description: "Service account for the connector.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"service_account_token_ref": {
							Description: "Reference to the secret containing the service account token for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"openid_connect": {
				Description: "OpenID configuration for the connector.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"issuer_url": {
							Description: "The URL of the OpenID Connect issuer.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username": {
							Description: "Username for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_ref": {
							Description: "Reference to the secret containing the username for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_id_ref": {
							Description: "Reference to the secret containing the client ID for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "Reference to the secret containing the password for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"secret_ref": {
							Description: "Reference to the secret containing the client secret for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"scopes": {
							Description: "Scopes to request for the connector.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"client_key_cert": {
				Description: "Client key and certificate config for the connector.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ca_cert_ref": {
							Description: "Reference to the secret containing the CA certificate for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_cert_ref": {
							Description: "Reference to the secret containing the client certificate for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_key_ref": {
							Description: "Reference to the secret containing the client key for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_key_passphrase_ref": {
							Description: "Reference to the secret containing the client key passphrase for the connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_key_algorithm": {
							Description: fmt.Sprintf("The algorithm used to generate the client key for the connector. Valid values are %s", strings.Join(nextgen.ClientKeyAlgorithmsSlice, ", ")),
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)
	gitsync.SetGitSyncSchema(resource.Schema, true)

	return resource
}
