package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorAzureCloudProvider() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure Cloud Provider in Harness.",
		ReadContext:   resourceConnectorAzureCloudProviderRead,
		CreateContext: resourceConnectorAzureCloudProviderCreateOrUpdate,
		UpdateContext: resourceConnectorAzureCloudProviderCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"credentials": {
				Description: "Contains Azure connector credentials.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description:  "Type can either be InheritFromDelegate or ManualConfig.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"InheritFromDelegate", "ManualConfig"}, false),
						},
						"azure_manual_details": {
							Description:   "Authenticate to Azure Cloud Provider using manual details.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"credentials.0.azure_inherit_from_delegate_details"},
							AtLeastOneOf:  []string{"credentials.0.azure_manual_details", "credentials.0.azure_inherit_from_delegate_details"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_id": {
										Description: "Application ID of the Azure App.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"tenant_id": {
										Description: "The Azure Active Directory (AAD) directory ID where you created your application.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"auth": {
										Description: "Contains Azure auth details.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Description:  "Type can either be Certificate or Secret.",
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice([]string{"Certificate", "Secret"}, false),
												},
												"azure_client_key_cert": {
													Description: "Azure client key certificate details.",
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate_ref": {
																Description: "Reference of the secret for the certificate." + secret_ref_text,
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"azure_client_secret_key": {
													Description: "Azure Client Secret Key details.",
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"secret_ref": {
																Description: "Reference of the secret for the secret key." + secret_ref_text,
																Type:        schema.TypeString,
																Optional:    true,
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
							Description:   "Authenticate to Azure Cloud Provider using details inheriting from delegate.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"credentials.0.azure_manual_details"},
							AtLeastOneOf:  []string{"credentials.0.azure_manual_details", "credentials.0.azure_inherit_from_delegate_details"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": {
										Description: "Auth to authenticate to Azure Cloud Provider using details inheriting from delegate.",
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Description:  "Type can either be SystemAssignedManagedIdentity or UserAssignedManagedIdentity.",
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringInSlice([]string{"SystemAssignedManagedIdentity", "UserAssignedManagedIdentity"}, false),
												},
												"azure_msi_auth_ua": {
													Description: "Azure UserAssigned MSI auth details.",
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"client_id": {
																Description: "Client Id of the ManagedIdentity resource.",
																Type:        schema.TypeString,
																Optional:    true,
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
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"azure_environment_type": {
				Description:  "Specifies the Azure Environment type, which is AZURE by default. Can either be AZURE or AZURE_US_GOVERNMENT",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"AZURE", "AZURE_US_GOVERNMENT"}, false),
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"force_delete": {
				Description: "Enable this flag for force deletion of connector",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAzureCloudProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Azure)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAzureCloudProvider(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAzureCloudProviderCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAzureCloudProvider(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAzureCloudProvider(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAzureCloudProvider(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Azure,
		Azure: &nextgen.AzureConnector{
			ConnectorType: nextgen.ConnectorTypes.Azure.String(),
			Credential:    &nextgen.AzureCredential{},
		},
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		if attr, ok := config["type"]; ok {
			connector.Azure.Credential.Type_ = nextgen.AzureCredentialType(attr.(string))

			if attr.(string) == nextgen.AzureCredentialTypes.ManualConfig.String() {
				if attr, ok := config["azure_manual_details"]; ok {
					configCredentials := attr.([]interface{})[0].(map[string]interface{})

					connector.Azure.Credential.AzureManualDetails = &nextgen.AzureManualDetails{}
					if attr, ok := configCredentials["application_id"]; ok {
						connector.Azure.Credential.AzureManualDetails.ApplicationId = attr.(string)
					}

					if attr, ok := configCredentials["tenant_id"]; ok {
						connector.Azure.Credential.AzureManualDetails.TenantId = attr.(string)
					}

					if attr, ok := configCredentials["auth"]; ok {
						configCredentialsManualAuth := attr.([]interface{})[0].(map[string]interface{})

						connector.Azure.Credential.AzureManualDetails.Auth = &nextgen.AzureAuth{}
						if attr, ok := configCredentialsManualAuth["type"]; ok {
							connector.Azure.Credential.AzureManualDetails.Auth.Type_ = attr.(string)

							if attr.(string) == nextgen.AzureAuthTypes.Certificate.String() {
								if attr, ok := configCredentialsManualAuth["azure_client_key_cert"]; ok {
									configCredentialsManualAuthSpec := attr.([]interface{})[0].(map[string]interface{})

									connector.Azure.Credential.AzureManualDetails.Auth.AzureClientKeyCert = &nextgen.AzureClientKeyCert{}
									if attr, ok := configCredentialsManualAuthSpec["certificate_ref"]; ok {
										connector.Azure.Credential.AzureManualDetails.Auth.AzureClientKeyCert.CertificateRef = attr.(string)
									}

								}
							}

							if attr.(string) == nextgen.AzureAuthTypes.SecretKey.String() {
								if attr, ok := configCredentialsManualAuth["azure_client_secret_key"]; ok {
									configCredentialsManualAuthSpec := attr.([]interface{})[0].(map[string]interface{})

									connector.Azure.Credential.AzureManualDetails.Auth.AzureClientSecretKey = &nextgen.AzureClientSecretKey{}
									if attr, ok := configCredentialsManualAuthSpec["secret_ref"]; ok {
										connector.Azure.Credential.AzureManualDetails.Auth.AzureClientSecretKey.SecretRef = attr.(string)
									}
								}
							}
						}
					}

				}
			}
			if attr.(string) == nextgen.AzureCredentialTypes.InheritFromDelegate.String() {

				if attr, ok := config["azure_inherit_from_delegate_details"]; ok {

					if attr, ok := attr.([]interface{})[0].(map[string]interface{})["auth"]; ok {
						connector.Azure.Credential.AzureInheritFromDelegateDetails = &nextgen.AzureInheritFromDelegateDetails{}
						configCredentials := attr.([]interface{})[0].(map[string]interface{})

						connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth = &nextgen.AzureMsiAuth{}
						if attr, ok := configCredentials["type"]; ok {
							connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.Type_ = attr.(string)

							if attr.(string) == nextgen.AzureMsiAuthTypes.UserAssignedManagedIdentity.String() {
								if attr, ok := configCredentials["azure_msi_auth_ua"]; ok {
									configCredentialsInheritFromDelegate := attr.([]interface{})[0].(map[string]interface{})

									connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.AzureMSIAuthUA = &nextgen.AzureUserAssignedMsiAuth{}
									if attr, ok := configCredentialsInheritFromDelegate["client_id"]; ok {
										connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.AzureMSIAuthUA.ClientId = attr.(string)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		delegate_selectors := attr.(*schema.Set).List()
		if len(delegate_selectors) > 0 {
			connector.Azure.DelegateSelectors = utils.InterfaceSliceToStringSlice(delegate_selectors)
		}
	}

	if attr, ok := d.GetOk("azure_environment_type"); ok {
		connector.Azure.AzureEnvironmentType = attr.(string)
	}

	return connector
}

func readConnectorAzureCloudProvider(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("credentials", []interface{}{
		map[string]interface{}{
			"type":                                connector.Azure.Credential.Type_,
			"azure_manual_details":                readAzureManualDetails(connector),
			"azure_inherit_from_delegate_details": readAzureInheritFromDelegateDetails(connector),
		},
	})
	d.Set("delegate_selectors", connector.Azure.DelegateSelectors)
	d.Set("azure_environment_type", connector.Azure.AzureEnvironmentType)

	return nil
}

func readAzureInheritFromDelegateDetails(connector *nextgen.ConnectorInfo) []map[string]interface{} {
	var spec []map[string]interface{}

	switch connector.Azure.Credential.Type_ {
	case nextgen.AzureCredentialTypes.ManualConfig:
		// noop
	case nextgen.AzureCredentialTypes.InheritFromDelegate:
		spec = []map[string]interface{}{
			{
				"auth": []map[string]interface{}{
					{
						"type":              connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.Type_,
						"azure_msi_auth_ua": readAzureInheritFromDelegateAuth(connector),
					},
				},
			},
		}
	}
	return spec
}

func readAzureManualDetails(connector *nextgen.ConnectorInfo) []map[string]interface{} {
	var spec []map[string]interface{}
	switch connector.Azure.Credential.Type_ {
	case nextgen.AzureCredentialTypes.ManualConfig:
		spec = []map[string]interface{}{
			{
				"application_id": connector.Azure.Credential.AzureManualDetails.ApplicationId,
				"tenant_id":      connector.Azure.Credential.AzureManualDetails.TenantId,
				"auth":           readAzureManualConfigAuth(connector),
			},
		}
	case nextgen.AzureCredentialTypes.InheritFromDelegate:
		//noop
	}

	return spec
}

func readAzureInheritFromDelegateAuth(connector *nextgen.ConnectorInfo) []map[string]interface{} {
	var spec []map[string]interface{}

	switch connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.Type_ {
	case nextgen.AzureMsiAuthTypes.SystemAssignedManagedIdentity.String():
		//noop
	case nextgen.AzureMsiAuthTypes.UserAssignedManagedIdentity.String():
		spec = []map[string]interface{}{
			{
				"client_id": connector.Azure.Credential.AzureInheritFromDelegateDetails.Auth.AzureMSIAuthUA.ClientId,
			},
		}
	}

	return spec
}

func readAzureManualConfigAuth(connector *nextgen.ConnectorInfo) []map[string]interface{} {
	var spec []map[string]interface{}

	switch connector.Azure.Credential.AzureManualDetails.Auth.Type_ {
	case string(nextgen.AzureAuthTypes.Certificate):
		spec = []map[string]interface{}{
			{
				"type": connector.Azure.Credential.AzureManualDetails.Auth.Type_,
				"azure_client_key_cert": []map[string]interface{}{
					{
						"certificate_ref": connector.Azure.Credential.AzureManualDetails.Auth.AzureClientKeyCert.CertificateRef,
					},
				},
			},
		}
	case nextgen.AzureAuthTypes.SecretKey.String():
		spec = []map[string]interface{}{
			{
				"type": connector.Azure.Credential.AzureManualDetails.Auth.Type_,
				"azure_client_secret_key": []map[string]interface{}{
					{
						"secret_ref": connector.Azure.Credential.AzureManualDetails.Auth.AzureClientSecretKey.SecretRef,
					},
				},
			},
		}
	}
	return spec
}
