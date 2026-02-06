package idp

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/po"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"gopkg.in/yaml.v3"
)

const environmentKind = "environment"

type environmentInfo struct {
	Identifier       string
	Scope            string
	BlueprintID      string
	BlueprintVersion string
	OrgID            optional.String
	ProjectID        optional.String
}

func ResourceEnvironment() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating IDP environments.",
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdateOrCreate,
		CreateContext: resourceEnvironmentUpdateOrCreate,
		DeleteContext: resourceEnvironmentDelete,
		Importer:      environmentImporter,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
			"org_id":     helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Required),
			"project_id": helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Required),
			"name":       helpers.GetNameSchema(helpers.SchemaFlagTypes.Optional),
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Owner of the environment",
			},
			"blueprint_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Blueprint to base the environment on",
			},
			"blueprint_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Version of the blueprint to base the environment on",
			},
			"based_on": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Based on environment reference. This should be passed as <orgIdentifier>.<projectIdentifier>/<environmentIdentifier>",
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+/[A-Za-z0-9_-]+$`),
					"Based on environment reference. This should be passed as <orgIdentifier>.<projectIdentifier>/<environmentIdentifier>",
				),
			},
			"target_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "target state of the environment. If different from the current, a pipeline will be triggered to update the environment",
				Default:     "inactive",
				ValidateFunc: validation.StringInSlice([]string{
					"running", "inactive", "paused",
				}, false),
			},
			"overrides": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Overrides for environment blueprint inputs in YAML format",
			},
			"inputs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Additional inputs for controlling the environment in YAML format",
			},
		},
	}

	return resource
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	environmentInfo := getEnvironmentInfo(d)

	id := d.Id()
	if id == "" {
		id = environmentInfo.Identifier
	}

	resp, httpResp, err := c.EntitiesApi.GetEntity(ctx, environmentInfo.Scope, environmentKind, id, &idp.EntitiesApiGetEntityOpts{
		OrgIdentifier:     environmentInfo.OrgID,
		ProjectIdentifier: environmentInfo.ProjectID,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if err := readEnvironment(d, resp); err != nil {
		return diag.Errorf("failed to read environment: %v", err)
	}

	return nil
}

func resourceEnvironmentUpdateOrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	var httpResp *http.Response
	var err error

	if id == "" {
		httpResp, err = createEnvironment(ctx, d, meta)
	} else {
		httpResp, err = updateEnvironment(ctx, d, meta)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	resourceEnvironmentRead(ctx, d, meta)

	return nil
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	id := d.Id()
	environmentInfo := getEnvironmentInfo(d)
	d.Set("target_state", "inactive")

	if _, err := updateEnvironment(ctx, d, meta); err != nil {
		return diag.Errorf("failed to delete environment %s: %v", fmt.Sprintf("%s/%s", environmentInfo.Scope, d.Id()), err)
	}

	httpResp, err := c.EnvironmentProxyApi.DeleteEnvironment(ctx, id, &idp.EnvironmentProxyApiDeleteEnvironmentOpts{
		HarnessAccount:    optional.NewString(c.AccountId),
		OrgIdentifier:     environmentInfo.OrgID,
		ProjectIdentifier: environmentInfo.ProjectID,
	})

	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}

		if isNotFoundError(err) {
			d.SetId("")
			return nil
		}

		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func createEnvironment(ctx context.Context, d *schema.ResourceData, meta interface{}) (*http.Response, error) {
	environmentInfo := getEnvironmentInfo(d)

	createRequest := idp.EnvironmentProxyCreateRequest{
		EnvironmentIdentifier:          environmentInfo.Identifier,
		EnvironmentName:                d.Get("name").(string),
		Owner:                          d.Get("owner").(string),
		EnvironmentBlueprintIdentifier: d.Get("blueprint_identifier").(string),
		EnvironmentBlueprintVersion:    d.Get("blueprint_version").(string),
		Overrides:                      d.Get("overrides").(string),
	}

	if v, ok := d.GetOk("inputs"); ok {
		createRequest.Inputs = v.(string)
	}

	if v, ok := d.GetOk("target_state"); ok {
		createRequest.TargetState = v.(string)
	}

	if v, ok := d.GetOk("based_on"); ok {
		createRequest.BasedOnIdentifier = fmt.Sprintf("environment:account.%s", v.(string))
	}

	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	resp, httpResp, err := c.EnvironmentProxyApi.CreateCompileAndExecuteEnvironment(ctx, createRequest, &idp.EnvironmentProxyApiCreateCompileAndExecuteEnvironmentOpts{
		OrgIdentifier:     environmentInfo.OrgID,
		ProjectIdentifier: environmentInfo.ProjectID,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
	if err != nil {
		return httpResp, fmt.Errorf("failed to create environment %s: err %w", fmt.Sprintf("%s/%s", environmentInfo.Scope, d.Id()), err)
	}

	if createRequest.TargetState != "inactive" && len(resp.Change.Instances) > 0 {
		if err := waitForExecutionTerminated(ctx, meta, environmentInfo, d.Timeout(schema.TimeoutCreate)); err != nil {
			return httpResp, fmt.Errorf("failed to observe environment execution %s: err %w", fmt.Sprintf("%s/%s", environmentInfo.Scope, d.Id()), err)
		}
	}

	return httpResp, nil
}

func updateEnvironment(ctx context.Context, d *schema.ResourceData, meta interface{}) (*http.Response, error) {
	environmentInfo := getEnvironmentInfo(d)

	updateRequest := idp.EnvironmentProxyUpdateRequest{
		EnvironmentBlueprintVersion: d.Get("blueprint_version").(string),
		Overrides:                   d.Get("overrides").(string),
	}

	if v, ok := d.GetOk("based_on"); ok {
		updateRequest.BasedOnIdentifier = fmt.Sprintf("environment:account.%s", v.(string))
	}

	if v, ok := d.GetOk("target_state"); ok {
		updateRequest.TargetState = v.(string)
	}

	if v, ok := d.GetOk("inputs"); ok {
		updateRequest.Inputs = v.(string)
	}

	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	resp, httpResp, err := c.EnvironmentProxyApi.UpdateCompileAndExecuteEnvironment(ctx, updateRequest, d.Id(), &idp.EnvironmentProxyApiUpdateCompileAndExecuteEnvironmentOpts{
		HarnessAccount:    optional.NewString(c.AccountId),
		OrgIdentifier:     environmentInfo.OrgID,
		ProjectIdentifier: environmentInfo.ProjectID,
	})
	if err != nil {
		return httpResp, fmt.Errorf("failed to update environment %s: err %w", fmt.Sprintf("%s/%s", environmentInfo.Scope, d.Id()), err)
	}

	if len(resp.Change.Instances) > 0 {
		err := waitForExecutionTerminated(ctx, meta, environmentInfo, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return httpResp, fmt.Errorf("failed to observe environment execution %s: err %w", fmt.Sprintf("%s/%s", environmentInfo.Scope, d.Id()), err)
		}
	}

	return httpResp, nil
}

var environmentImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		// Expected format: <scope>/<identifier>
		// Scope examples: "org.project"
		id := d.Id()
		parts := strings.Split(id, "/")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid import ID format: %s. Expected: <scope>/<identifier>", id)
		}

		scope := parts[0]
		identifier := parts[1]

		// Extract org and project from scope if present
		var orgId, projectId optional.String
		scopeParts := strings.Split(scope, ".")
		if len(scopeParts) != 2 {
			return nil, fmt.Errorf("invalid scope format: %s. Expected: <orgIdentifier>.<projectIdentifier>", scope)
		}
		orgId = optional.NewString(scopeParts[0])
		projectId = optional.NewString(scopeParts[1])

		c, ctx := meta.(*internal.Session).GetIDPClientWithContext(context.Background())

		resp, _, err := c.EntitiesApi.GetEntity(ctx, fmt.Sprintf("account.%s", scope), environmentKind, identifier, &idp.EntitiesApiGetEntityOpts{
			OrgIdentifier:     orgId,
			ProjectIdentifier: projectId,
			HarnessAccount:    optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to fetch entity for import: %w", err)
		}

		if err := readEnvironment(d, resp); err != nil {
			return nil, fmt.Errorf("failed to read environment for import: %w", err)
		}

		return []*schema.ResourceData{d}, nil
	},
}

func readEnvironment(d *schema.ResourceData, entity idp.EntityResponse) error {
	d.SetId(entity.Identifier)
	d.Set("identifier", entity.Identifier)
	d.Set("name", entity.Name)
	d.Set("org_id", entity.OrgIdentifier)
	d.Set("project_id", entity.ProjectIdentifier)
	d.Set("owner", entity.Owner)
	overridesSet := false
	inputsSet := false

	if v, ok := d.GetOk("overrides"); ok && v.(string) != "" {
		d.Set("overrides", v.(string))
		overridesSet = true
	}

	if v, ok := d.GetOk("inputs"); ok && v.(string) != "" {
		d.Set("inputs", v.(string))
		inputsSet = true
	}

	if entity.Spec != nil && *entity.Spec != nil {
		raw := *entity.Spec
		entitySpec := raw.(map[string]any)

		if !overridesSet {
			if v, ok := entitySpec["overrides"]; ok {
				b, err := yaml.Marshal(v)
				if err != nil {
					return fmt.Errorf("failed to marshal overrides as yaml. Err: %w", err)
				}

				d.Set("overrides", string(b))
			}
		}

		if !inputsSet {
			if v, ok := entitySpec["inputs"]; ok {
				b, err := yaml.Marshal(v)
				if err != nil {
					return fmt.Errorf("failed to marshal inputs as yaml. Err: %w", err)
				}

				d.Set("inputs", string(b))
			}
		}

		if v, ok := entitySpec["environmentBlueprint"]; ok {
			blueprint := v.(map[string]any)
			d.Set("blueprint_identifier", blueprint["identifier"])
			d.Set("blueprint_version", blueprint["version"])
		}

		if v, ok := entitySpec["targetState"]; ok {
			targetState := v.(map[string]any)
			d.Set("target_state", targetState["state"])
		}

		if v, ok := entitySpec["basedOn"]; ok {
			basedOn := v.(map[string]any)
			d.Set("based_on", fmt.Sprintf("%s.%s/%s", basedOn["orgIdentifier"], basedOn["projectIdentifier"], basedOn["identifier"]))
		}
	}

	return nil
}

func getEnvironmentInfo(d *schema.ResourceData) environmentInfo {
	identifier := d.Get("identifier").(string)
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	blueprintId := d.Get("blueprint_identifier").(string)
	blueprintVersion := d.Get("blueprint_version").(string)

	scope := fmt.Sprintf("account.%s.%s", orgID, projectID)

	return environmentInfo{
		Identifier:       identifier,
		BlueprintID:      blueprintId,
		BlueprintVersion: blueprintVersion,
		Scope:            scope,
		OrgID:            optional.NewString(orgID),
		ProjectID:        optional.NewString(projectID),
	}
}

type PoEnvInfo struct {
	Progress string
}

var progressTerminal = []string{"done", "failed", "aborted", "deleted", "replaced"}
var progressRunning = []string{"pending", "processing"}

func waitForExecutionTerminated(ctx context.Context, meta any, info environmentInfo, timeout time.Duration) error {
	c, ctx := meta.(*internal.Session).GetPOClientWithContext(ctx)

	stateConf := &retry.StateChangeConf{
		Pending: progressRunning,
		Target:  progressTerminal,
		Refresh: func() (any, string, error) {
			resp, _, err := c.InfrastructureApi.InfrastructureGet(ctx, info.Identifier, c.AccountId, &po.InfrastructureApiInfrastructureGetOpts{
				OrgIdentifier:     info.OrgID,
				ProjectIdentifier: info.ProjectID,
			})
			if err != nil {
				return nil, "", err
			}

			if len(resp.ActiveExecutions) == 0 {
				return resp, "done", nil
			}

			var progress string
			for _, exec := range resp.ActiveExecutions {
				// If an execution is still running, we can break and wait poll interval
				progress = exec.Progress.Progress
				if slices.Index(progressRunning, progress) != -1 {
					break
				}
			}

			return resp, progress, nil
		},
		Timeout:      timeout,
		Delay:        10 * time.Second,
		PollInterval: 15 * time.Second,
		MinTimeout:   5 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
