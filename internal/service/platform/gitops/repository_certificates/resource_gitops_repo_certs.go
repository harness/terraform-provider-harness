package repository_certificates

import (
	"context"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsRepoCerts() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Gitops Repositories Certificates.",

		CreateContext: resourceGitopsRepoCertsCreateOrUpdate,
		ReadContext:   resourceGitopsRepoCertsRead,
		UpdateContext: resourceGitopsRepoCertsCreateOrUpdate,
		DeleteContext: resourceGitopsRepoCertsDelete,
		Importer:      helpers.GitopsRepoCertResourceImporter,

		Schema: map[string]*schema.Schema{
			"agent_id": {
				Description: "agent identifier of the Repository Certificates.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_id": {
				Description: "account identifier of the Repository Certificates.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "organization identifier of the Repository Certificates.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project identifier of the Repository Certificates.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request": {
				Description: "Repository Certificates create/Update request.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "if the Repository Certificates should be upserted.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"certificates": {
							Description: "certificates details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metadata": {
										Description: "metadata details",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"self_link": {
													Description: "selfLink is a URL representing this object.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"resource_version": {
													Description: "dentifies the server's internal version.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"continue": {
													Description: "continue may be set if the user set a limit on the number of items returned.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"remaining_item_count": {
													Description: "subsequent items in the list.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"items": {
										Description: "List of certificates to be processed.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"server_name": {
													Description: "ServerName specifies the DNS name of the server this certificate is intended.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cert_type": {
													Description: "CertType specifies the type of the certificate - currently one of https or ssh.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cert_sub_type": {
													Description: "CertSubType specifies the sub type of the cert, i.e. ssh-rsa.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cert_data": {
													Description: "CertData contains the actual certificate data, dependent on the certificate type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cert_info": {
													Description: "CertInfo will hold additional certificate info, depdendent on the certificate type .",
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
	}

	return resource
}

func resourceGitopsRepoCertsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	repoCertRequest := buildRepoCertRequest(d)

	resp, httpResp, err := c.RepositoryCertificatesApi.AgentCertificateServiceCreate(ctx, *repoCertRequest, agentIdentifier,
		&nextgen.RepositoryCertificatesApiAgentCertificateServiceCreateOpts{
			AccountIdentifier: optional.NewString(accountIdentifier),
			OrgIdentifier:     optional.NewString(orgIdentifier),
			ProjectIdentifier: optional.NewString(projectIdentifier),
		})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setGitopsRepositoriesCertificates(d, &resp)
	return nil
}

func resourceGitopsRepoCertsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	agentIdentifier := d.Get("agent_id").(string)

	resp, httpResp, err := c.RepositoryCertificatesApi.AgentCertificateServiceList(ctx, agentIdentifier, c.AccountId, &nextgen.RepositoryCertificatesApiAgentCertificateServiceListOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setGitopsRepositoriesCertificates(d, &resp)
	return nil
}

func resourceGitopsRepoCertsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)

	_, httpResp, err := c.RepositoryCertificatesApi.AgentCertificateServiceDelete(ctx, agentIdentifier, &nextgen.RepositoryCertificatesApiAgentCertificateServiceDeleteOpts{
		AccountIdentifier:    optional.NewString(c.AccountId),
		OrgIdentifier:        optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier:    optional.NewString(d.Get("project_id").(string)),
		QueryHostNamePattern: optional.NewString(d.Get("request.0.certificates.0.items.0.server_name").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func setGitopsRepositoriesCertificates(d *schema.ResourceData, repoCert *nextgen.CertificatesRepositoryCertificateList) {
	d.SetId("1234") // assigning temp id.
	var repoCertDetails nextgen.Applicationv1alpha1RepositoryCertificateList
	metadata := &nextgen.V1ListMeta{}

	if repoCert.Metadata != nil {
		metadata.SelfLink = repoCert.Metadata.SelfLink
		metadata.ResourceVersion = repoCert.Metadata.ResourceVersion
		metadata.Continue_ = repoCert.Metadata.Continue_
		metadata.RemainingItemCount = repoCert.Metadata.RemainingItemCount
		repoCertDetails.Metadata = metadata
	}

	if repoCert.Items != nil {
		var Items []nextgen.Applicationv1alpha1RepositoryCertificate
		for _, value := range repoCert.Items {
			var item nextgen.Applicationv1alpha1RepositoryCertificate
			item.ServerName = value.ServerName
			item.CertType = value.CertType
			item.CertSubType = value.CertSubType
			item.CertData = value.CertData
			item.CertInfo = value.CertInfo
			Items = append(Items, item)
		}
		repoCertDetails.Items = Items
	}
}

func buildRepoCertRequest(d *schema.ResourceData) *nextgen.CertificateRepositoryCertificateCreateRequest {
	var upsert bool
	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})
		upsert = request["upsert"].(bool)
	}

	return &nextgen.CertificateRepositoryCertificateCreateRequest{
		Upsert:       upsert,
		Certificates: buildRepoCertDetails(d),
	}
}

func buildRepoCertDetails(d *schema.ResourceData) *nextgen.Applicationv1alpha1RepositoryCertificateList {
	var repoCertDetails nextgen.Applicationv1alpha1RepositoryCertificateList

	var request map[string]interface{}

	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})

		if request["certificates"] != nil && len(request["certificates"].([]interface{})) > 0 {
			for _, v := range request["certificates"].([]interface{}) {
				requestCertificate := v.(map[string]interface{})

				if requestCertificate["metadata"] != nil {
					if requestCertificate["metadata"] != nil && len(requestCertificate["metadata"].([]interface{})) > 0 &&
						requestCertificate["metadata"].([]interface{})[0] != nil {
						certMetadata := requestCertificate["metadata"].([]interface{})[0].(map[string]interface{})
						repoCertDetails.Metadata = &nextgen.V1ListMeta{}

						if certMetadata["self_link"] != nil {
							repoCertDetails.Metadata.SelfLink = certMetadata["self_link"].(string)
						}

						if certMetadata["resource_version"] != nil {
							repoCertDetails.Metadata.ResourceVersion = certMetadata["resource_version"].(string)
						}

						if certMetadata["continue"] != nil {
							repoCertDetails.Metadata.Continue_ = certMetadata["continue"].(string)
						}

						if certMetadata["remaining_item_count"] != nil {
							repoCertDetails.Metadata.RemainingItemCount = certMetadata["remaining_item_count"].(string)
						}

					}
				}

				if requestCertificate["items"] != nil && len(requestCertificate["items"].([]interface{})) > 0 {
					var items []nextgen.Applicationv1alpha1RepositoryCertificate
					for _, v := range requestCertificate["items"].([]interface{}) {
						if v != nil {
							var vMap = v.(map[string]interface{})
							var item nextgen.Applicationv1alpha1RepositoryCertificate

							if vMap["server_name"] != nil && len(vMap["server_name"].(string)) > 0 {
								item.ServerName = vMap["server_name"].(string)
							}

							if vMap["cert_type"] != nil && len(vMap["cert_type"].(string)) > 0 {
								item.CertType = vMap["cert_type"].(string)
							}

							if vMap["cert_sub_type"] != nil && len(vMap["cert_sub_type"].(string)) > 0 {
								item.CertSubType = vMap["cert_sub_type"].(string)
							}

							if vMap["cert_data"] != nil && len(vMap["cert_data"].(string)) > 0 {
								item.CertData = vMap["cert_data"].(string)
							}

							if vMap["cert_info"] != nil && len(vMap["cert_info"].(string)) > 0 {
								item.CertInfo = vMap["cert_info"].(string)
							}
							items = append(items, item)
						}
					}
					repoCertDetails.Items = items
				}
			}
		}
	}
	return &repoCertDetails
}
