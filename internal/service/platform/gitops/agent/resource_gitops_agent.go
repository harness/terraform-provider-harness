package agent

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceGitopsAgent() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness GitOps Agent.",

		CreateContext: resourceGitopsAgentCreate,
		ReadContext:   resourceGitopsAgentRead,
		UpdateContext: resourceGitopsAgentUpdate,
		DeleteContext: resourceGitopsAgentDelete,
		Importer:      resourceGitopsAgentImporter,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
			var e error
			if diff.HasChange("project_id") && diff.Id() != "" {
				e = fmt.Errorf("field 'project_id' cannot be changed after the resource is created")
			}
			if diff.HasChange("org_id") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'org_id' cannot be changed after the resource is created:%w", e)
				} else {
					e = fmt.Errorf("field 'org_id' cannot be changed after the resource is created")
				}
			}
			if diff.HasChange("identifier") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'identifier' cannot be changed after the resource is created:%v", e)
				} else {
					e = fmt.Errorf("field 'identifier' cannot be changed after the resource is created")
				}
			}
			if diff.HasChange("metadata.0.existing_installation") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'metadata.existing_installation' cannot be changed after the agent is created:%w", e)
				} else {
					e = fmt.Errorf("field 'metadata.existing_installation' cannot be changed after the agent is created")
				}
			}

			oldDrIdentifier, newDrIdentifier := diff.GetChange("disaster_recovery_identifier")
			oldDrType, newDrType := diff.GetChange("disaster_recovery_type")
			forceNewIdentifier, forceNewType, drErr := validateDisasterRecoveryDiff(
				oldDrIdentifier.(string), newDrIdentifier.(string),
				oldDrType.(string), newDrType.(string),
				diff.Id() != "",
			)
			if drErr != nil {
				if e != nil {
					e = fmt.Errorf("%v: %w", e, drErr)
				} else {
					e = drErr
				}
			}
			if forceNewIdentifier {
				if fnErr := diff.ForceNew("disaster_recovery_identifier"); fnErr != nil {
					return fnErr
				}
			}
			if forceNewType {
				if fnErr := diff.ForceNew("disaster_recovery_type"); fnErr != nil {
					return fnErr
				}
			}

			// TODO(CDS-135089): if/when token_rotation lands (PR #1556), add a CustomizeDiff guard:
			// diff.Id() != "" && diff.HasChange("token_rotation") && disaster_recovery_identifier != "" -> plan-time error (RegenerateCredentials isn't DR-aware yet)

			return e
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release.",
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
				Description: "Use `MANAGED_ARGO_PROVIDER` for agents in your cluster (recommended). `HOSTED_ARGO_PROVIDER` is for Harness-hosted agents and is not customer-provisioned.\nEnum: \"AGENT_TYPE_UNSET\" \"MANAGED_ARGO_PROVIDER\" \"HOSTED_ARGO_PROVIDER\"",
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
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Description: "The kubernetes namespace where the agent should be installed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"high_availability": {
							Description: "Indicates if the agent is deployed in HA mode.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_namespaced": {
							Description: "Indicates if the agent is namespaced.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"existing_installation": {
							Description: "Indicates if the agent connects to an existing Argo CD installation (BYOA). When true, Harness skips CRD and Argo CD component installation. This field cannot be changed after the agent is created.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
					}},
			},
			"agent_token": {
				Description: "Agent token to be used for authentication of the agent with Harness.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"operator": {
				Description: "The Operator to use for the Harness GitOps agent. Enum: \"ARGO\" \"FLAMINGO\"",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"prefixed_identifier": {
				Description: "Prefixed identifier of the GitOps agent. Agent identifier prefixed with scope of the agent",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"disaster_recovery_identifier": {
				Description:  "Disaster Recovery identifier shared by all nodes (primary and secondary) in the same DR group. Required together with disaster_recovery_type.",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"disaster_recovery_type"},
				ValidateFunc: func(val interface{}, key string) ([]string, []error) {
					v := val.(string)
					if strings.Contains(v, ":") || strings.Contains(v, "/") {
						return nil, []error{fmt.Errorf("%q must not contain ':' or '/', got: %s", key, v)}
					}
					return nil, nil
				},
			},
			"disaster_recovery_type": {
				Description:  "Disaster Recovery role of this agent node. Required together with disaster_recovery_identifier. Enum: \"PRIMARY\" \"SECONDARY\"",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"disaster_recovery_identifier"},
				ValidateFunc: validation.StringInSlice([]string{"PRIMARY", "SECONDARY"}, false),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old != ""
				},
			},
			"disaster_recovery_live_type": {
				Description: "The live Disaster Recovery role of this agent node as reported by the server. Populated by Read/Import only.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"uuid": {
				Description: "UUID of the GitOps agent. Shared by every Disaster Recovery sibling under the same disaster_recovery_identifier.",
				Type:        schema.TypeString,
				Computed:    true,
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
		return helpers.HandleReadApiError(err, d, httpResp)
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

	getOpts := &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	}
	if drIdentifier := d.Get("disaster_recovery_identifier").(string); drIdentifier != "" {
		getOpts.DrIdentifier = optional.NewString(drIdentifier)
	}
	resp, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, getOpts)

	if err != nil && httpResp != nil && httpResp.StatusCode != 404 {
		return helpers.HandleApiError(err, d, httpResp)
	}
	if err != nil && httpResp == nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil || (httpResp != nil && httpResp.StatusCode == 404) {
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

	var e diag.Diagnostics
	if d.HasChange("identifier") {
		oldValue, newValue := d.GetChange("identifier")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'identifier' cannot be updated after creation.")[0])
		}
		if err := d.Set("identifier", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("org_id") {
		oldValue, newValue := d.GetChange("org_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'org_id' cannot be updated after creation.")[0])
		}
		if err := d.Set("org_id", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("project_id") {
		oldValue, newValue := d.GetChange("project_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'project_id' cannot be updated after creation.")[0])
		}
		if err := d.Set("project_id", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if len(e) > 0 {
		return e
	}

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

	drIdentifier := d.Get("disaster_recovery_identifier").(string)
	if drIdentifier == "" {
		readAgent(d, &resp)
		return nil
	}

	// The server's DR update path always returns the group's PRIMARY record regardless of which
	// node was targeted, so the Update response cannot be used to populate a SECONDARY's state.
	// Discard it and re-fetch this node specifically via Get.
	getResp, getHttpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		DrIdentifier:      optional.NewString(drIdentifier),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, getHttpResp)
	}
	readAgent(d, &getResp)
	return nil
}

func resourceGitopsAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	drIdentifier := d.Get("disaster_recovery_identifier").(string)

	if drIdentifier != "" && d.Get("disaster_recovery_live_type").(string) == "PRIMARY" {
		// Check-then-act: a SECONDARY could be promoted or created between this check and the
		// Delete call below. This is an accepted limitation, not an atomic guarantee.
		getResp, getHttpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
			OrgIdentifier:     optional.NewString(orgId),
			ProjectIdentifier: optional.NewString(projectId),
			DrIdentifier:      optional.NewString(drIdentifier),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, getHttpResp)
		}
		if getResp.DisasterRecoveryNode != nil && getResp.DisasterRecoveryNode.Type_ != nil &&
			*getResp.DisasterRecoveryNode.Type_ == nextgen.PRIMARY_V1DisasterRecoveryType {
			listResp, listHttpResp, err := c.AgentApi.AgentServiceForServerList(ctx, c.AccountId, string(*getResp.Type_), &nextgen.AgentsApiAgentServiceForServerListOpts{
				OrgIdentifier:     optional.NewString(orgId),
				ProjectIdentifier: optional.NewString(projectId),
				Identifier:        optional.NewString(agentIdentifier),
				DrIdentifier:      optional.NewString(drIdentifier),
				IncludeSecondary:  optional.NewBool(true),
			})
			if err != nil {
				return helpers.HandleApiError(err, d, listHttpResp)
			}
			for _, sibling := range listResp.Content {
				if sibling.DisasterRecoveryNode != nil && sibling.DisasterRecoveryNode.Type_ != nil &&
					*sibling.DisasterRecoveryNode.Type_ == nextgen.SECONDARY_V1DisasterRecoveryType {
					return diag.Errorf("cannot delete live PRIMARY node '%s' while SECONDARY sibling(s) still exist for disaster_recovery_identifier '%s' — promote a SECONDARY to PRIMARY first, or destroy the SECONDARY resource(s) first", agentIdentifier, drIdentifier)
				}
			}
		}
	}

	deleteOpts := &nextgen.AgentsApiAgentServiceForServerDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projectId),
	}
	if drIdentifier != "" {
		deleteOpts.DrIdentifier = optional.NewString(drIdentifier)
	}
	_, httpResp, err := c.AgentApi.AgentServiceForServerDelete(ctx, agentIdentifier, deleteOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

// validateDisasterRecoveryDiff implements the CustomizeDiff rules for the disaster_recovery_identifier
// and disaster_recovery_type fields. It reports whether either field must be forced to replacement,
// and an error when a transition is disallowed at plan time.
func validateDisasterRecoveryDiff(oldIdentifier, newIdentifier, oldType, newType string, hasId bool) (forceNewIdentifier, forceNewType bool, err error) {
	if oldIdentifier != "" && newIdentifier != "" && oldIdentifier != newIdentifier {
		forceNewIdentifier = true
	}

	if oldIdentifier != "" && oldType != "" && newIdentifier == "" && newType == "" {
		forceNewIdentifier = true
		forceNewType = true
	}

	if oldIdentifier == "" && oldType == "" && newIdentifier != "" && newType != "" {
		if hasId && newType == "SECONDARY" {
			err = fmt.Errorf("a regular agent can only be converted to a disaster recovery primary; converting an existing non-DR agent directly to 'SECONDARY' is not supported")
		}
	}

	return forceNewIdentifier, forceNewType, err
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
		tags := make(map[string]string)
		for k, v := range attr.(map[string]interface{}) {
			tags[k] = v.(string)
		}
		v1Agent.Tags = tags
	}
	if attr, ok := d.GetOk("metadata"); ok {
		metadata := attr.([]interface{})
		if attr != nil && len(metadata) > 0 {
			meta := metadata[0].(map[string]interface{})
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
			if meta["existing_installation"] != nil {
				v1MetaData.ExistingInstallation = meta["existing_installation"].(bool)
			}

			v1Agent.Metadata = &v1MetaData
		}
	}
	drIdentifier, hasDrIdentifier := d.GetOk("disaster_recovery_identifier")
	drType, hasDrType := d.GetOk("disaster_recovery_type")
	if hasDrIdentifier && hasDrType {
		drTypeVal := nextgen.V1DisasterRecoveryType(drType.(string))
		v1Agent.DisasterRecoveryNode = &nextgen.V1DisasterRecoveryNode{
			Identifier: drIdentifier.(string),
			Type_:      &drTypeVal,
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
	d.Set("account_id", agent.AccountIdentifier)
	d.Set("identifier", agent.Identifier)
	d.Set("name", agent.Name)
	d.Set("description", agent.Description)
	d.Set("tags", agent.Tags)
	d.Set("org_id", agent.OrgIdentifier)
	d.Set("type", agent.Type_)
	d.Set("project_id", agent.ProjectIdentifier)
	d.Set("operator", agent.Operator)
	d.Set("prefixed_identifier", agent.PrefixedIdentifier)
	d.Set("uuid", agent.Uuid)
	if agent.DisasterRecoveryNode != nil {
		d.Set("disaster_recovery_live_type", agent.DisasterRecoveryNode.Type_)
	} else {
		d.Set("disaster_recovery_live_type", "")
	}
	metadata := []interface{}{}
	metaDataMap := map[string]interface{}{}
	metaDataMap["namespace"] = agent.Metadata.Namespace
	metaDataMap["high_availability"] = agent.Metadata.HighAvailability
	metaDataMap["is_namespaced"] = agent.Metadata.IsNamespaced
	metaDataMap["existing_installation"] = agent.Metadata.ExistingInstallation
	metadata = append(metadata, metaDataMap)
	d.Set("metadata", metadata)
	if agent.Credentials != nil && agent.Credentials.PrivateKey != "" {
		d.Set("agent_token", agent.Credentials.PrivateKey)
	}
}

// resourceGitopsAgentImporter is a dedicated importer for this resource because
// helpers.MultiLevelResourceImporter dispatches purely on path-segment count and cannot
// disambiguate a DR node's import ID (which carries an extra ":<dr_identifier>" suffix)
// from an ordinary org/project-scoped one.
var resourceGitopsAgentImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		id := d.Id()
		rest := id
		var drIdentifier string
		if idx := strings.LastIndex(id, ":"); idx != -1 {
			drIdentifier = id[idx+1:]
			rest = id[:idx]
		}

		parts := strings.Split(rest, "/")
		switch len(parts) {
		case 1:
			d.SetId(parts[0])
			d.Set("identifier", parts[0])
		case 2:
			d.SetId(parts[1])
			d.Set("identifier", parts[1])
			d.Set("org_id", parts[0])
		case 3:
			d.SetId(parts[2])
			d.Set("identifier", parts[2])
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
		default:
			return nil, fmt.Errorf("invalid identifier: %s, expected format: [org_id/[project_id/]]identifier[:disaster_recovery_identifier]", id)
		}

		if drIdentifier == "" {
			return []*schema.ResourceData{d}, nil
		}
		d.Set("disaster_recovery_identifier", drIdentifier)

		// disaster_recovery_type has no config to source from during import, so it is populated
		// here from the server's reported live role instead — the one exception to it otherwise
		// only ever coming from user config.
		c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(context.Background())
		ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
		resp, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
			DrIdentifier:      optional.NewString(drIdentifier),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to fetch disaster recovery role for agent '%s': %w (http status: %v)", d.Get("identifier").(string), err, httpResp)
		}
		if resp.DisasterRecoveryNode != nil && resp.DisasterRecoveryNode.Type_ != nil {
			d.Set("disaster_recovery_type", string(*resp.DisasterRecoveryNode.Type_))
		}
		return []*schema.ResourceData{d}, nil
	},
}
