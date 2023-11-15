package agent

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsAgent() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness Gitops Agent.",

		CreateContext: resourceGitopsAgentCreate,
		ReadContext:   resourceGitopsAgentRead,
		UpdateContext: resourceGitopsAgentUpdate,
		DeleteContext: resourceGitopsAgentDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Default: \"AGENT_TYPE_UNSET\"\nEnum: \"AGENT_TYPE_UNSET\" \"CONNECTED_ARGO_PROVIDER\" \"MANAGED_ARGO_PROVIDER\"",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tags": {
				Description: "Tags for the GitOps agents. These can be used to search or filter the GitOps agents.",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"metadata": {
				Description: "Metadata of the agent.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Description: "The kubernetes namespace where the agent is installed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"high_availability": {
							Description: "Indicates if the agent is installed in HA mode.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_namespaced": {
							Description: "Indicates if the agent is namespaced.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					}},
			},
			"agent_token": {
				Description: "Agent token to be used for authentication of the agent with Harness.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"operator": {
				Description: "The Operator to use for the Harness GitOps agent. Enum: \"ARGO\" \"FLAMINGO\"",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	return resource
}

func resourceGitopsAgentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	createAgentRequest := buildCreateAgentRequest(d)
	createAgentRequest.AccountIdentifier = c.AccountId
	resp, httpResp, err := c.AgentApi.AgentServiceForServerCreate(ctx, *createAgentRequest)

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
	readAgent(d, &resp)
	return nil
}

func resourceGitopsAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
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
	readAgent(d, &resp)
	return nil
}

func resourceGitopsAgentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)
	updateAgentRequest := buildCreateUpdateAgentRequest(d)
	updateAgentRequest.AccountIdentifier = c.AccountId
	resp, httpResp, err := c.AgentApi.AgentServiceForServerUpdate(ctx, *updateAgentRequest, agentIdentifier)

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
	readAgent(d, &resp)
	return nil
}

func resourceGitopsAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)

	_, httpResp, err := c.AgentApi.AgentServiceForServerDelete(ctx, agentIdentifier, &nextgen.AgentsApiAgentServiceForServerDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildCreateUpdateAgentRequest(d *schema.ResourceData) *nextgen.V1Agent {
	var v1Agent nextgen.V1Agent
	if attr, ok := d.GetOk("account_id"); ok {
		v1Agent.AccountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		v1Agent.ProjectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		v1Agent.OrgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		v1Agent.Identifier = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		v1Agent.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		v1Agent.Description = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		agentType := nextgen.V1AgentType(attr.(string))
		v1Agent.Type_ = &agentType
	}
	if attr, ok := d.GetOk("tags"); ok {
		v1Agent.Tags = attr.(map[string]string)
	}
	if attr, ok := d.GetOk("metadata"); ok {
		metadata := attr.([]interface{})
		if attr != nil && len(metadata) > 0 {
			meta := metadata[0].(map[string]interface{})
			fmt.Println("META: ", meta)
			var v1MetaData nextgen.V1AgentMetadata

			if meta["high_availability"] != nil {
				v1MetaData.HighAvailability = meta["high_availability"].(bool)
			}
			if meta["namespace"] != nil {
				v1MetaData.Namespace = meta["namespace"].(string)
			}
			if meta["is_namespaced"] != nil {
				v1MetaData.IsNamespaced = meta["is_namespaced"].(bool)
			}

			v1Agent.Metadata = &v1MetaData
		}
	}
	return &v1Agent
}

func buildCreateAgentRequest(d *schema.ResourceData) *nextgen.V1Agent {
	v1Agent := buildCreateUpdateAgentRequest(d)
	if attr, ok := d.GetOk("operator"); ok {
		agentOperator := nextgen.V1AgentOperator(attr.(string))
		v1Agent.Operator = &agentOperator
	}
	return v1Agent
}

func readAgent(d *schema.ResourceData, agent *nextgen.V1Agent) {
	d.SetId(agent.Identifier)
	d.Set("identifier", agent.Identifier)
	d.Set("name", agent.Name)
	d.Set("description", agent.Description)
	d.Set("tags", agent.Tags)
	d.Set("org_id", agent.OrgIdentifier)
	d.Set("type", agent.Type_)
	d.Set("project_id", agent.ProjectIdentifier)
	d.Set("operator", agent.Operator)
	metadata := []interface{}{}
	metaDataMap := map[string]interface{}{}
	metaDataMap["namespace"] = agent.Metadata.Namespace
	metaDataMap["high_availability"] = agent.Metadata.HighAvailability
	metaDataMap["is_namespaced"] = agent.Metadata.IsNamespaced
	metadata = append(metadata, metaDataMap)
	d.Set("metadata", metadata)
	if agent.Credentials != nil && agent.Credentials.PrivateKey != "" {
		d.Set("agent_token", agent.Credentials.PrivateKey)
	}
}
