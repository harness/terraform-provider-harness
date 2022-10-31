package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorAzureCloudProvider() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Azure Cloud Provider Connector.",
		ReadContext: resourceConnectorAzureCloudProviderRead,

		Schema: map[string]*schema.Schema{
			"credentials": {
				Description: "Contains Azure connector credentials.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Type can either be InheritFromDelegate or ManualConfig.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"azure_manual_details": {
							Description: "Authenticate to Azure Cloud Provider using manual details.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_id": {
										Description: "Application ID of the Azure App.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"tenant_id": {
										Description: "The Azure Active Directory (AAD) directory ID where you created your application.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"auth": {
										Description: "Contains Azure auth details.",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Description: "Type can either be Certificate or Secret.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"azure_client_key_cert": {
													Description: "Azure client key certificate details.",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate_ref": {
																Description: "Reference of the secret for the certificate.",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"azure_client_secret_key": {
													Description: "Azure Client Secret Key details.",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"secret_ref": {
																Description: "Reference of the secret for the secret key.",
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
							},
						},
						"azure_inherit_from_delegate_details": {
							Description: "Authenticate to Azure Cloud Provider using details inheriting from delegate.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": {
										Description: "Auth to authenticate to Azure Cloud Provider using details inheriting from delegate.",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Description: "Type can either be SystemAssignedManagedIdentity or UserAssignedManagedIdentity.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"azure_msi_auth_ua": {
													Description: "Azure UserAssigned MSI auth details.",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"client_id": {
																Description: "Client Id of the ManagedIdentity resource.",
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
							},
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"azure_environment_type": {
				Description: "Specifies the Azure Environment type, which is AZURE by default. Can either be AZURE or AZURE_US_GOVERNMENT",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
