package probe_template

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProbeTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Probe Template.",

		ReadContext: dataSourceProbeTemplateRead,

		Schema: map[string]*schema.Schema{
			"hub_identity": {
				Description: "Identity of the chaos hub.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identity": {
				Description:   "Unique identifier of the probe template.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Description:   "Name of the probe template.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"identity"},
			},
			"org_id": {
				Description: "Organization identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},

			// Computed fields - all fields from the resource
			"account_id": {
				Description: "Account identifier.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the probe template.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Type of the probe template.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infrastructure_type": {
				Description: "Infrastructure type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the probe template.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"revision": {
				Description: "Revision number.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"is_default": {
				Description: "Whether this is the default version.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"hub_ref": {
				Description: "Hub reference.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			// Probe properties - simplified for data source
			"http_probe": {
				Description: "HTTP probe configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cmd_probe": {
				Description: "Command probe configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"k8s_probe": {
				Description: "Kubernetes probe configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Run properties
			"run_properties": {
				Description: "Run properties.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"interval": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stop_on_failure": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},

			// Variables
			"variables": {
				Description: "Template variables.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProbeTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	// If identity is provided, fetch directly
	if identity, ok := d.GetOk("identity"); ok {
		identityStr := identity.(string)
		log.Printf("[DEBUG] Fetching probe template by identity: %s", identityStr)

		resp, httpResp, apiErr := c.DefaultApi.GetProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, identityStr, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil {
			return diag.Errorf("probe template not found: %s", identityStr)
		}

		// Set the ID
		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, resp.Data.Identity))

		// Set all the data
		return setProbeTemplateDataSimplified(d, resp.Data, accountID, orgID, projectID, hubIdentity)
	} else if name, ok := d.GetOk("name"); ok {
		// If name is provided, list and filter
		nameStr := name.(string)
		log.Printf("[DEBUG] Fetching probe template by name: %s", nameStr)

		resp, httpResp, apiErr := c.DefaultApi.ListProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, 0, 100, nameStr, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil || len(resp.Data) == 0 {
			return diag.Errorf("probe template not found with name: %s", nameStr)
		}

		// Find exact match
		var foundTemplate *chaos.GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosprobetemplateChaosProbeTemplate
		for i, t := range resp.Data {
			if t.Name == nameStr {
				foundTemplate = &resp.Data[i]
				break
			}
		}

		if foundTemplate == nil {
			return diag.Errorf("probe template not found with name: %s", nameStr)
		}

		// Set the ID
		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, foundTemplate.Identity))

		// Set all the data
		return setProbeTemplateDataSimplified(d, foundTemplate, accountID, orgID, projectID, hubIdentity)
	} else {
		return diag.Errorf("either 'identity' or 'name' must be specified")
	}
}
