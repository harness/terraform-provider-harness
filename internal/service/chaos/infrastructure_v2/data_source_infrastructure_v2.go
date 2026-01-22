package infrastructure_v2

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosInfrastructureV2() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Infrastructure V2.",

		ReadContext: dataSourceChaosInfrastructureV2Read,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "The ID of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "The ID of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment_id": {
				Description: "The ID of the environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"infra_id": {
				Description: "The ID of the infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},

			// All other fields should match the resource schema but be computed
			"identifier": {
				Description: "Identifier of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identity": {
				Description: "Identity of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"last_heartbeat": {
				Description: "Last heartbeat of the infrastructure.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"last_workflow_timestamp": {
				Description: "Last workflow timestamp of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"no_of_schedules": {
				Description: "Number of schedules for the infrastructure.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"no_of_workflows": {
				Description: "Number of workflows for the infrastructure.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"status": {
				Description: "Status of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags of the infrastructure.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update_status": {
				Description: "Update status of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Created at of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_by": {
				Description: "Created by of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Updated at of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_by": {
				Description: "Updated by of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infra_type": {
				Description: "Type of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infra_scope": {
				Description: "Scope of the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_account": {
				Description: "Service account used by the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"namespace": {
				Description: "Kubernetes namespace for the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"containers": {
				Description: "List of containers in the infrastructure.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"image_registry": imageRegistrySchema(),
			"mtls":           mtlsSchema(),
			"proxy":          proxySchema(),
			"volumes":        volumesSchema(),
			"volume_mounts":  volumeMountsSchema(),
			"tolerations":    tolerationsSchema(),
			"node_selector": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"label": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"annotation": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_ai_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_chaos_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"insecure_skip_verify": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"run_as_user": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"run_as_group": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceChaosInfrastructureV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Get required fields
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	environmentID := d.Get("environment_id").(string)
	infraID := d.Get("infra_id").(string)

	// Get the infrastructure
	infra, httpResp, err := c.ChaosSdkApi.GetInfraV2(
		ctx,
		infraID,
		c.AccountId,
		orgID,
		projectID,
		environmentID,
	)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	// Set the ID
	d.SetId(infra.Identity)

	// Set all the fields from the infrastructure
	if err := setInfrastructureFields(d, &infra); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// Helper function to set infrastructure fields in the resource data
func setInfrastructureFields(d *schema.ResourceData, infra *chaos.InfraV2KubernetesInfrastructureV2Details) error {
	// Set basic fields
	d.Set("name", infra.Name)
	d.Set("description", infra.Description)
	d.Set("infra_type", infra.InfraType)
	d.Set("infra_scope", infra.InfraScope)
	d.Set("service_account", infra.ServiceAccount)
	d.Set("namespace", infra.InfraNamespace)
	d.Set("identifier", infra.Identity)
	d.Set("infra_id", infra.InfraID)
	d.Set("environment_id", infra.EnvironmentID)
	d.Set("last_heartbeat", infra.LastHeartbeat)
	d.Set("last_workflow_timestamp", infra.LastWorkflowTimestamp)
	d.Set("no_of_schedules", infra.NoOfSchedules)
	d.Set("no_of_workflows", infra.NoOfWorkflows)
	d.Set("status", infra.Status)
	d.Set("tags", infra.Tags)
	d.Set("update_status", infra.UpdateStatus)
	d.Set("created_at", infra.CreatedAt)
	d.Set("created_by", infra.CreatedBy.Username)
	d.Set("updated_at", infra.UpdatedAt)
	d.Set("updated_by", infra.UpdatedBy.Username)
	d.Set("is_ai_enabled", infra.IsAIEnabled)
	d.Set("is_chaos_enabled", infra.IsChaosEnabled)
	d.Set("insecure_skip_verify", infra.InsecureSkipVerify)
	d.Set("run_as_user", infra.RunAsUser)
	d.Set("run_as_group", infra.RunAsGroup)
	d.Set("containers", infra.Containers)
	d.Set("identity", infra.Identity)

	// Set maps
	if err := d.Set("node_selector", infra.NodeSelector); err != nil {
		return fmt.Errorf("failed to set node_selector: %v", err)
	}
	if err := d.Set("label", infra.Label); err != nil {
		return fmt.Errorf("failed to set labels: %v", err)
	}
	if err := d.Set("annotation", infra.Annotation); err != nil {
		return fmt.Errorf("failed to set annotations: %v", err)
	}

	// Set nested objects
	if infra.ImageRegistry != nil {
		if err := setImageRegistry(d, infra.ImageRegistry); err != nil {
			return fmt.Errorf("failed to set image_registry: %v", err)
		}
	}
	if infra.Mtls != nil {
		if err := setMtls(d, infra.Mtls); err != nil {
			return fmt.Errorf("failed to set mtls: %v", err)
		}
	}
	if infra.Proxy != nil {
		if err := setProxy(d, infra.Proxy); err != nil {
			return fmt.Errorf("failed to set proxy: %v", err)
		}
	}
	if len(infra.Volumes) > 0 {
		if err := setVolumes(d, infra.Volumes); err != nil {
			return fmt.Errorf("failed to set volumes: %v", err)
		}
	}
	if len(infra.VolumeMounts) > 0 {
		if err := setVolumeMounts(d, infra.VolumeMounts); err != nil {
			return fmt.Errorf("failed to set volume_mounts: %v", err)
		}
	}
	if len(infra.Tolerations) > 0 {
		if err := setTolerations(d, infra.Tolerations); err != nil {
			return fmt.Errorf("failed to set tolerations: %v", err)
		}
	}

	return nil
}
