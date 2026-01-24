package experiment_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceExperimentTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identity": {
			Description: "Unique identifier for the experiment template",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the experiment template",
			Type:        schema.TypeString,
			Required:    true,
		},
		"org_id": {
			Description: "Organization identifier",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"hub_identity": {
			Description: "Hub identifier where the template is stored",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"description": {
			Description: "Description of the experiment template",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags associated with the experiment template",
			Type:        schema.TypeList,
			Optional:    true,
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

		// Spec block
		"spec": {
			Description: "Specification of the experiment template",
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"infra_id": {
						Description: "Infrastructure identifier (supports runtime input: <+input>)",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"infra_type": {
						Description: "Infrastructure type (Windows, Linux, CloudFoundry, Container, Kubernetes, KubernetesV2)",
						Type:        schema.TypeString,
						Required:    true,
						ValidateFunc: validation.StringInSlice([]string{
							"Windows",
							"Linux",
							"CloudFoundry",
							"Container",
							"Kubernetes",
							"KubernetesV2",
						}, false),
					},

					// Actions
					"actions": {
						Description: "List of actions in the experiment",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"identity": {
									Description: "Action template identity",
									Type:        schema.TypeString,
									Required:    true,
								},
								"name": {
									Description: "Action name",
									Type:        schema.TypeString,
									Required:    true,
								},
								"infra_id": {
									Description: "Infrastructure identifier for this action",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"revision": {
									Description: "Action template revision",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"is_enterprise": {
									Description: "Whether this is an enterprise action",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"continue_on_completion": {
									Description: "Whether to continue on completion",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"values": {
									Description: "Variable values for the action",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Description: "Variable name",
												Type:        schema.TypeString,
												Required:    true,
											},
											"value": {
												Description: "Variable value (supports runtime input: <+input>)",
												Type:        schema.TypeString,
												Required:    true,
											},
										},
									},
								},
							},
						},
					},

					// Faults
					"faults": {
						Description: "List of faults in the experiment",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"identity": {
									Description: "Fault template identity",
									Type:        schema.TypeString,
									Required:    true,
								},
								"name": {
									Description: "Fault name",
									Type:        schema.TypeString,
									Required:    true,
								},
								"infra_id": {
									Description: "Infrastructure identifier for this fault",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"revision": {
									Description: "Fault template revision",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"is_enterprise": {
									Description: "Whether this is an enterprise fault",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"auth_enabled": {
									Description: "Whether authentication is enabled",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"values": {
									Description: "Variable values for the fault",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Description: "Variable name",
												Type:        schema.TypeString,
												Required:    true,
											},
											"value": {
												Description: "Variable value (supports runtime input: <+input>)",
												Type:        schema.TypeString,
												Required:    true,
											},
										},
									},
								},
							},
						},
					},

					// Probes
					"probes": {
						Description: "List of probes in the experiment",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"identity": {
									Description: "Probe template identity",
									Type:        schema.TypeString,
									Required:    true,
								},
								"name": {
									Description: "Probe name",
									Type:        schema.TypeString,
									Required:    true,
								},
								"infra_id": {
									Description: "Infrastructure identifier for this probe",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"revision": {
									Description: "Probe template revision",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"is_enterprise": {
									Description: "Whether this is an enterprise probe",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"duration": {
									Description: "Probe duration",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"weightage": {
									Description: "Probe weightage for resilience score calculation",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"enable_data_collection": {
									Description: "Whether to enable data collection",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"conditions": {
									Description: "Probe execution conditions",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"execute_upon": {
												Description: "When to execute the probe (onChaosStart, duringChaos, afterChaos)",
												Type:        schema.TypeString,
												Required:    true,
											},
										},
									},
								},
								"values": {
									Description: "Variable values for the probe",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Description: "Variable name",
												Type:        schema.TypeString,
												Required:    true,
											},
											"value": {
												Description: "Variable value (supports runtime input: <+input>)",
												Type:        schema.TypeString,
												Required:    true,
											},
										},
									},
								},
							},
						},
					},

					// Vertices (workflow graph)
					"vertices": {
						Description: "Workflow graph vertices defining execution order",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Description: "Vertex name",
									Type:        schema.TypeString,
									Required:    true,
								},
								"start": {
									Description: "Start configuration for the vertex",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"actions": {
												Description: "Actions to execute at start",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Action name",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},
											"faults": {
												Description: "Faults to execute at start",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Fault name",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},
											"probes": {
												Description: "Probes to execute at start",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Probe name",
															Type:        schema.TypeString,
															Required:    true,
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
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"actions": {
												Description: "Actions to execute at end",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Action name",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},
											"faults": {
												Description: "Faults to execute at end",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Fault name",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},
											"probes": {
												Description: "Probes to execute at end",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Probe name",
															Type:        schema.TypeString,
															Required:    true,
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

					// Cleanup Policy
					"cleanup_policy": {
						Description: "Cleanup policy for experiment resources (retain, delete)",
						Type:        schema.TypeString,
						Optional:    true,
						ValidateFunc: validation.StringInSlice([]string{
							"retain",
							"delete",
						}, false),
					},

					// Status Check Timeouts
					"status_check_timeouts": {
						Description: "Status check timeout configuration",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"delay": {
									Description: "Delay before status check (in seconds)",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"timeout": {
									Description: "Timeout for status check (in seconds)",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}
