package agent_yaml

import (
	"context"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsAgentDeployYaml() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for fetching a Harness Gitops Agent deployment manifest YAML.",

		ReadContext: dataSourceGitopsAgentDeployYamlRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description: "The kubernetes namespace where the agent is installed.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "The deployment manifest YAML of the GitOps agent.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ca_data": {
				Description: "CA data of the GitOps agent, base64 encoded content of ca chain.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"proxy": {
				Description: "Proxy settings for the GitOps agent.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http": {
							Description: "HTTP proxy settings for the GitOps agent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"https": {
							Description: "HTTPS proxy settings for the GitOps agent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username for the proxy.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": {
							Description: "Password for the proxy.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourceGitopsAgentDeployYamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	agentIdentifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AgentApi.AgentServiceForServerPostDeployYaml(ctx, func() nextgen.V1AgentYamlQuery {
		var yamlQuery nextgen.V1AgentYamlQuery
		if attr, ok := d.GetOk("account_id"); ok {
			yamlQuery.AccountIdentifier = attr.(string)
		}
		if attr, ok := d.GetOk("project_id"); ok {
			yamlQuery.ProjectIdentifier = attr.(string)
		}
		if attr, ok := d.GetOk("org_id"); ok {
			yamlQuery.OrgIdentifier = attr.(string)
		}
		if attr, ok := d.GetOk("namespace"); ok {
			yamlQuery.Namespace = attr.(string)
		}
		if attr, ok := d.GetOk("ca_data"); ok {
			yamlQuery.CaData = attr.(string)
		}

		if attr, ok := d.GetOk("proxy"); ok {
			proxy := attr.([]interface{})
			if attr != nil && len(proxy) > 0 {
				p := proxy[0].(map[string]interface{})
				var v1Proxy nextgen.V1Proxy
				if p["http"] != nil {
					v1Proxy.Http = p["http"].(string)
				}
				if p["https"] != nil {
					v1Proxy.Https = p["https"].(string)
				}
				if p["username"] != nil {
					v1Proxy.Username = p["username"].(string)
				}
				if p["password"] != nil {
					v1Proxy.Password = p["password"].(string)
				}
				yamlQuery.Proxy = &v1Proxy
			}
		}
		return yamlQuery
	}(), agentIdentifier)

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
	readAgentYaml(agentIdentifier, d, resp)
	return nil
}

func readAgentYaml(agentIdentifier string, d *schema.ResourceData, yaml string) {
	d.SetId(agentIdentifier)
	d.Set("yaml", yaml)
}
