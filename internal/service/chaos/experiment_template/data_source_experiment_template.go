package experiment_template

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceExperimentTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Experiment Template by identity or name.",

		ReadContext: dataSourceExperimentTemplateRead,

		Schema: map[string]*schema.Schema{
			"identity": {
				Description:  "Unique identifier of the experiment template. Either identity or name must be specified.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},
			"name": {
				Description:  "Name of the experiment template. Either identity or name must be specified.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},
			"org_id": {
				Description: "Organization identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hub_identity": {
				Description: "Hub identifier where the template is stored",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the experiment template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the experiment template",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_default": {
				Description: "Whether this is a default template",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"revision": {
				Description: "Revision of the experiment template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"api_version": {
				Description: "API version of the experiment template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"kind": {
				Description: "Kind of the experiment template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"spec": {
				Description: "Specification of the experiment template",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"infra_id": {
							Description: "Infrastructure identifier",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"infra_type": {
							Description: "Infrastructure type",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"actions": {
							Description: "List of actions in the experiment",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identity": {
										Description: "Action template identity",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": {
										Description: "Action name",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"infra_id": {
										Description: "Infrastructure identifier for this action",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"revision": {
										Description: "Action template revision",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"is_enterprise": {
										Description: "Whether this is an enterprise action",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"continue_on_completion": {
										Description: "Whether to continue on completion",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"values": {
										Description: "Variable values for the action",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "Variable name",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": {
													Description: "Variable value",
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"faults": {
							Description: "List of faults in the experiment",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identity": {
										Description: "Fault template identity",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": {
										Description: "Fault name",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"infra_id": {
										Description: "Infrastructure identifier for this fault",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"revision": {
										Description: "Fault template revision",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"is_enterprise": {
										Description: "Whether this is an enterprise fault",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"auth_enabled": {
										Description: "Whether authentication is enabled",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"values": {
										Description: "Variable values for the fault",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "Variable name",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": {
													Description: "Variable value",
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"probes": {
							Description: "List of probes in the experiment",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identity": {
										Description: "Probe template identity",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": {
										Description: "Probe name",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"infra_id": {
										Description: "Infrastructure identifier for this probe",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"revision": {
										Description: "Probe template revision",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"is_enterprise": {
										Description: "Whether this is an enterprise probe",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"duration": {
										Description: "Probe duration",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"weightage": {
										Description: "Probe weightage for resilience score calculation",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"enable_data_collection": {
										Description: "Whether to enable data collection",
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"conditions": {
										Description: "Probe execution conditions",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"execute_upon": {
													Description: "When to execute the probe",
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"values": {
										Description: "Variable values for the probe",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "Variable name",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": {
													Description: "Variable value",
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"vertices": {
							Description: "Workflow graph vertices defining execution order",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Vertex name",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"start": {
										Description: "Start configuration for the vertex",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"actions": {
													Description: "Actions to execute at start",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Action name",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"faults": {
													Description: "Faults to execute at start",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Fault name",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"probes": {
													Description: "Probes to execute at start",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Probe name",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"end": {
										Description: "End configuration for the vertex",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"actions": {
													Description: "Actions to execute at end",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Action name",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"faults": {
													Description: "Faults to execute at end",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Fault name",
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"probes": {
													Description: "Probes to execute at end",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Probe name",
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
						"cleanup_policy": {
							Description: "Cleanup policy for experiment resources",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_check_timeouts": {
							Description: "Status check timeout configuration",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delay": {
										Description: "Delay before status check (in seconds)",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"timeout": {
										Description: "Timeout for status check (in seconds)",
										Type:        schema.TypeInt,
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
}

func dataSourceExperimentTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	var template *chaos.ChaosexperimenttemplateGetExperimentTemplateResponse
	var err error

	// Check if identity is provided
	if identity, ok := d.GetOk("identity"); ok {
		// Lookup by identity (direct API call)
		log.Printf("[DEBUG] Looking up experiment template by identity: %s", identity.(string))

		opts := &chaos.ExperimenttemplateApiGetExperimentTemplateOpts{}
		opts.HubIdentity = optional.NewString(hubIdentity)
		opts.OrganizationIdentifier = optional.NewString(orgID)
		opts.ProjectIdentifier = optional.NewString(projectID)

		resp, httpResp, apiErr := c.ExperimenttemplateApi.GetExperimentTemplate(ctx, accountID, identity.(string), opts)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}
		template = &resp
	} else if name, ok := d.GetOk("name"); ok {
		// Lookup by name (list and filter)
		log.Printf("[DEBUG] Looking up experiment template by name: %s", name.(string))

		opts := &chaos.ExperimenttemplateApiListExperimentTemplateOpts{}
		opts.HubIdentity = optional.NewString(hubIdentity)
		opts.OrganizationIdentifier = optional.NewString(orgID)
		opts.ProjectIdentifier = optional.NewString(projectID)

		listResp, httpResp, apiErr := c.ExperimenttemplateApi.ListExperimentTemplate(ctx, accountID, opts)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		// Find template by name
		var found *chaos.ChaosexperimenttemplateChaosExperimentTemplate
		for _, t := range listResp.Data {
			if t.Name == name.(string) {
				found = &t
				break
			}
		}

		if found == nil {
			return diag.Errorf("experiment template with name '%s' not found in hub '%s'", name.(string), hubIdentity)
		}

		// Get full template details
		opts2 := &chaos.ExperimenttemplateApiGetExperimentTemplateOpts{}
		opts2.HubIdentity = optional.NewString(hubIdentity)
		opts2.OrganizationIdentifier = optional.NewString(orgID)
		opts2.ProjectIdentifier = optional.NewString(projectID)

		resp, httpResp, apiErr := c.ExperimenttemplateApi.GetExperimentTemplate(ctx, accountID, found.Identity, opts2)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}
		template = &resp
	} else {
		return diag.Errorf("either identity or name must be specified")
	}

	// Set ID
	d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, template.Identity))

	// Set data
	if err = setExperimentTemplateData(d, template, orgID, projectID, hubIdentity); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
