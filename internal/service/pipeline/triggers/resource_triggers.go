package triggers

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Regex patterns for validation
const (
	triggerIdentifierPattern = "^[a-zA-Z_][0-9a-zA-Z_]{0,127}$"
	triggerNamePattern       = "^[a-zA-Z_0-9-.][-0-9a-zA-Z_\\s.]{0,127}$"
)

func ResourceTriggers() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating triggers in Harness.",
		ReadContext:   resourceTriggersRead,
		UpdateContext: resourceTriggersCreateOrUpdate,
		CreateContext: resourceTriggersCreateOrUpdate,
		DeleteContext: resourceTriggersDelete,
		Importer:      helpers.TriggerResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the trigger",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(triggerIdentifierPattern),
					"identifier must start with a letter or underscore and can contain only alphanumeric characters and underscores (max 128 chars)",
				),
			},
			"name": {
				Description: "Name of the trigger",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(triggerNamePattern),
					"name must start with a letter, number, underscore, hyphen, or dot and can contain alphanumeric characters, hyphens, underscores, spaces, and dots (max 128 chars)",
				),
			},
			"target_id": {
				Description: "Identifier of the target pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ignore_error": {
				Description: "ignore error default false",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"yaml": {
				Description:      "trigger yaml." + helpers.Descriptions.YamlText.String(),
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
				ValidateFunc:     validateTriggerYaml,
			},
			"if_match": {
				Description: "if-Match",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

// validateTriggerYaml validates that the YAML contains valid identifier and name
func validateTriggerYaml(v interface{}, k string) (warns []string, errs []error) {
	yamlStr := v.(string)
	
	// Extract identifier and name from YAML
	identifierRegex := regexp.MustCompile(`identifier:\s*([^\s]+)`)
	nameRegex := regexp.MustCompile(`name:\s*([^\n]+)`)
	
	identifierMatches := identifierRegex.FindStringSubmatch(yamlStr)
	nameMatches := nameRegex.FindStringSubmatch(yamlStr)
	
	// Validate identifier if found
	if len(identifierMatches) > 1 {
		identifier := identifierMatches[1]
		if !regexp.MustCompile(triggerIdentifierPattern).MatchString(identifier) {
			errs = append(errs, fmt.Errorf("invalid identifier in YAML: %s. Identifier must start with a letter or underscore and can contain only alphanumeric characters and underscores (max 128 chars)", identifier))
		}
	}
	
	// Validate name if found
	if len(nameMatches) > 1 {
		name := nameMatches[1]
		name = regexp.MustCompile(`^[\s"']*|[\s"']*$`).ReplaceAllString(name, "") // Trim quotes and spaces
		if !regexp.MustCompile(triggerNamePattern).MatchString(name) {
			errs = append(errs, fmt.Errorf("invalid name in YAML: %s. Name must start with a letter, number, underscore, hyphen, or dot and can contain alphanumeric characters, hyphens, underscores, spaces, and dots (max 128 chars)", name))
		}
	}
	
	return warns, errs
}

func resourceTriggersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	resp, httpResp, err := c.TriggersApi.GetTrigger(ctx, c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string), d.Get("target_id").(string), id)

	if httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTriggers(d, resp.Data)

	return nil
}

func resourceTriggersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtongTriggerResponse
	var httpResp *http.Response
	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.TriggersApi.CreateTrigger(ctx, d.Get("yaml").(string), c.AccountId,
			d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), &nextgen.TriggersApiCreateTriggerOpts{
				WithServiceV2: optional.NewBool(true),
			})
		// A FORCED PAUSE TO PREVENT DUPLICATE WEBHOOK CREATION.
		time.Sleep(5 * time.Second)
	} else {
		resp, httpResp, err = c.TriggersApi.UpdateTrigger(ctx, d.Get("yaml").(string), c.AccountId, d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), id, &nextgen.TriggersApiUpdateTriggerOpts{
				IfMatch: helpers.BuildField(d, "if_match"),
			})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTriggers(d, resp.Data)

	return nil
}

func resourceTriggersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.TriggersApi.DeleteTrigger(ctx, c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("target_id").(string), d.Id(), &nextgen.TriggersApiDeleteTriggerOpts{
		IfMatch: helpers.BuildField(d, "if_match"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readTriggers(d *schema.ResourceData, trigger *nextgen.NgTriggerResponse) {
	d.SetId(trigger.Identifier)
	d.Set("identifier", trigger.Identifier)
	d.Set("name", trigger.Name)
	d.Set("org_id", trigger.OrgIdentifier)
	d.Set("project_id", trigger.ProjectIdentifier)
	d.Set("target_id", trigger.TargetIdentifier)
	d.Set("yaml", trigger.Yaml)
}
