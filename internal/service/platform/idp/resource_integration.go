package idp

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// integrationTypeBlocks are the type-specific spec blocks; at most one is set, matching type.
var integrationTypeBlocks = map[string]string{
	"github":          "GitHub",
	"bitbucket_cloud": "BitbucketCloud",
	"pagerduty":       "PagerDuty",
	"datadog":         "DataDog",
	"dynatrace":       "DynaTrace",
	"gcp":             "GCP",
	"catalog_info":    "CatalogInfo",
	"sonarqube":       "SonarQube",
	"harness_cd":      "HarnessCD",
	"harness_code":    "HarnessCode",
	"harness_k8s":     "HarnessK8s",
	"custom":          "Custom",
}

func ResourceIntegration() *schema.Resource {
	s := integrationSpecSchema()
	s["kind"].Elem.(*schema.Resource).Schema["mapping"] = kindMappingSchema()
	for k, v := range integrationStatusSchema() {
		s[k] = v
	}
	helpers.SetMultiLevelResourceSchema(s)
	s["identifier"].Description = "Identifier of the integration, stored as sent."
	s["identifier"].ValidateFunc = validation.All(
		validation.StringMatch(identifierPattern, "must start with a letter or _ and contain only letters, digits, "+
			"_ and $, up to 128 characters"),
		validation.StringNotInSlice(reservedIdentifiers, false),
	)
	s["type"] = &schema.Schema{
		Description: "Integration type, e.g. ServiceNow, GitHub, HarnessCD or Custom.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	}
	s["enabled"] = &schema.Schema{
		Description: "Whether the integration syncs.",
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
	}

	return &schema.Resource{
		Description:   "Resource for an IDP integration, which imports entities from an external system or Harness.",
		CreateContext: resourceIntegrationCreate,
		ReadContext:   resourceIntegrationRead,
		UpdateContext: resourceIntegrationUpdate,
		DeleteContext: resourceIntegrationDelete,
		CustomizeDiff: integrationCustomizeDiff,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema:        s,
	}
}

func optionalString(desc string) *schema.Schema {
	return &schema.Schema{Description: desc, Type: schema.TypeString, Optional: true}
}

// defaultedString is set by the server when omitted, so the server's value doesn't show as a diff.
func defaultedString(desc string) *schema.Schema {
	return &schema.Schema{Description: desc, Type: schema.TypeString, Optional: true, Computed: true}
}

func requiredString(desc string) *schema.Schema {
	return &schema.Schema{Description: desc, Type: schema.TypeString, Required: true}
}

func stringList(desc string) *schema.Schema {
	return &schema.Schema{Description: desc, Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}}
}

func stringSet(desc string) *schema.Schema {
	return &schema.Schema{Description: desc, Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}}
}

func block(desc string, required bool, fields map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Description: desc,
		Type:        schema.TypeList,
		MaxItems:    1,
		Required:    required,
		Optional:    !required,
		Elem:        &schema.Resource{Schema: fields},
	}
}

// defaultedBlock is Optional and Computed so a server default doesn't show as a diff.
func defaultedBlock(desc string, fields map[string]*schema.Schema) *schema.Schema {
	b := block(desc, false, fields)
	b.Computed = true
	return b
}

func typeBlock(typ string, fields map[string]*schema.Schema) *schema.Schema {
	return defaultedBlock(fmt.Sprintf("Settings for %s integrations.", typ), fields)
}

// integrationSpecSchema is the API's spec, flattened into the resource.
func integrationSpecSchema() map[string]*schema.Schema {
	startDate := func() *schema.Schema { return optionalString("Lower bound for the first sync, RFC3339.") }
	return map[string]*schema.Schema{
		"connector": block("Connector the integration reads with. Required or not allowed depending on type.", false,
			map[string]*schema.Schema{
				"inline": block("Credentials stored on the integration.", false, map[string]*schema.Schema{
					"credentials_ref": {
						Description: "Harness connector reference, e.g. account.snow. Can't change.",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
					},
				}),
				"ref": optionalString("Reference to a standalone IDP connector."),
			}),
		"update_frequency": defaultedString("How often an Airbyte integration syncs, e.g. 24h."),
		"secondary_kinds": {
			Description: "Secondary kinds to sync, for Airbyte types.",
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"scopes": defaultedBlock("Harness scopes a platform integration reads from. Set exactly one of all "+
			"or selected.", map[string]*schema.Schema{
			"all":      {Description: "Read from every scope.", Type: schema.TypeBool, Optional: true},
			"selected": stringList("Scopes to read from: \"\" for the account, \"org\" or \"org/project\"."),
		}),
		"custom": typeBlock("Custom", map[string]*schema.Schema{
			"mode": {
				Description:  "webhook or api.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"webhook", "api"}, false),
			},
			"source_label": {
				Description: "Label for the sending system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"webhook_auth": block("Webhook authentication. Required in webhook mode.", false, map[string]*schema.Schema{
				"auth_type": {
					Description:  "hmac_auth or jwt_auth.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"hmac_auth", "jwt_auth"}, false),
				},
				"hash_algo": {
					Description: "HMAC hash: hmac_sha224, hmac_sha256, hmac_sha384 or hmac_sha512.",
					Type:        schema.TypeString,
					Optional:    true,
					ValidateFunc: validation.StringInSlice(
						[]string{"hmac_sha224", "hmac_sha256", "hmac_sha384", "hmac_sha512"}, false),
				},
				"auth_header": optionalString("Header carrying the signature or token."),
				"secret_ref":  optionalString("Secret reference holding the shared secret."),
			}),
			"layout": {
				Description: "Columns shown for the integration's entities.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"column": {
						Description: "A column.",
						Type:        schema.TypeList,
						Required:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"title": requiredString("Column title."),
							"path":  requiredString("JQ path of the value."),
						}},
					},
				}},
			},
		}),
		"kind": {
			Description: "A kind to sync. Import reads kinds sorted by name, so list them that way to import without a diff.",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"kind":        requiredString("Kind name."),
				"auto_import": {Description: "Import new entities automatically.", Type: schema.TypeBool, Optional: true},
				"action": {
					Description:  "register, merge or merge_register; the type's default when omitted.",
					Type:         schema.TypeString,
					Optional:     true,
					Computed:     true,
					ValidateFunc: validation.StringInSlice([]string{"register", "merge", "merge_register"}, false),
				},
				"replication_mode": {
					Description:  "Airbyte types: full_refresh or incremental; the stream's default when omitted.",
					Type:         schema.TypeString,
					Optional:     true,
					Computed:     true,
					ValidateFunc: validation.StringInSlice([]string{"full_refresh", "incremental"}, false),
				},
				"filter": optionalString("JQ that returns true for records to keep."),
				"idp_kind": defaultedString("Airbyte types: entity kind, e.g. component; the type's default when " +
					"omitted. Set at create only; required for kinds with no default, such as ServiceNow's."),
				"idp_type": defaultedString("Airbyte types: entity type, e.g. service; the type's default when " +
					"omitted. Set at create only."),
			}},
		},
		"github": typeBlock("GitHub", map[string]*schema.Schema{
			"orgs":            {Description: "Organizations to sync.", Type: schema.TypeSet, Required: true, Elem: &schema.Schema{Type: schema.TypeString}},
			"start_date":      startDate(),
			"agent_file_name": defaultedString("Agents file to detect in repositories."),
		}),
		"bitbucket_cloud": typeBlock("BitbucketCloud", map[string]*schema.Schema{"start_date": startDate()}),
		"pagerduty": typeBlock("PagerDuty", map[string]*schema.Schema{
			"start_date": startDate(),
			"analytics_duration_months": {
				Description: "Months of analytics to aggregate.", Type: schema.TypeInt, Optional: true, Computed: true,
			},
		}),
		"datadog": typeBlock("DataDog", map[string]*schema.Schema{
			"service_dependencies_env": optionalString("Environment used to add service dependencies."),
		}),
		"dynatrace": typeBlock("DynaTrace", map[string]*schema.Schema{
			"platform_token_ref": optionalString("Secret reference to a Platform API token."),
			"start_date":         defaultedString("Lower bound for incremental streams, RFC3339."),
		}),
		"gcp": typeBlock("GCP", map[string]*schema.Schema{
			"parent": block("The organization or folder whose assets are synced. Set exactly one field.", true,
				map[string]*schema.Schema{
					"org_id":    optionalString("Organization ID."),
					"folder_id": optionalString("Folder ID."),
				}),
		}),
		"catalog_info": typeBlock("CatalogInfo", map[string]*schema.Schema{
			"source_type":  defaultedString("github or bitbucket."),
			"repositories": stringSet("Repositories to scan; omit to scan all."),
			"catalog_path": defaultedString("Catalog file glob, relative to the repository root."),
			"branch":       optionalString("Branch to read; omit for each repository's default branch."),
			"git_sync": block("Writes converted Harness YAML to git.", false, map[string]*schema.Schema{
				"connector_ref": requiredString("Connector used to push."),
				"base_path":     defaultedString("Directory the files are written to."),
				"target": block("Where to push. Set exactly one of source_repo or repo.", true, map[string]*schema.Schema{
					"source_repo": {Description: "Push to the repository each file was read from.", Type: schema.TypeBool, Optional: true},
					"repo": block("Push to one repository.", false, map[string]*schema.Schema{
						"name":   requiredString("Repository."),
						"branch": requiredString("Branch."),
					}),
				}),
			}),
		}),
		"sonarqube": typeBlock("SonarQube", map[string]*schema.Schema{
			"token_ref": requiredString("Secret reference to a SonarQube user token."),
			"instance": block("Set exactly one of server or cloud.", true, map[string]*schema.Schema{
				"server": block("Self-hosted SonarQube Server.", false, map[string]*schema.Schema{
					"host_url": requiredString("Server URL."),
				}),
				"cloud": block("SonarQube Cloud.", false, map[string]*schema.Schema{
					"organization": requiredString("Organization key."),
					"host_url":     defaultedString("Cloud URL."),
				}),
			}),
			"projects": block("Projects to sync; omit to sync all. Set exactly one of include or exclude.", false,
				map[string]*schema.Schema{
					"include": stringSet("Only these project keys."),
					"exclude": stringSet("Every project except these keys."),
				}),
			"issues": block("Filters for the issues stream.", false, map[string]*schema.Schema{
				"impact_severities": stringSet("INFO, LOW, MEDIUM, HIGH or BLOCKER."),
				"issue_statuses":    stringSet("OPEN, CONFIRMED, FALSE_POSITIVE, ACCEPTED, FIXED or IN_SANDBOX."),
				"severities":        stringSet("INFO, MINOR, MAJOR, CRITICAL or BLOCKER."),
				"types":             stringSet("CODE_SMELL, BUG or VULNERABILITY."),
				"resolved": {
					Description:  "\"true\" for resolved issues only, \"false\" for unresolved only; omit for both.",
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
				},
			}),
			"start_date": startDate(),
		}),
		"harness_cd": typeBlock("HarnessCD", map[string]*schema.Schema{
			"deployment_cycle": {Description: "Days of deployments to read.", Type: schema.TypeInt, Optional: true, Computed: true},
			"mapping_fields": {
				Description: "Fields to map per kind; omit to map all.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"kind":   requiredString("Kind."),
					"fields": {Description: "Fields.", Type: schema.TypeList, Required: true, Elem: &schema.Schema{Type: schema.TypeString}},
				}},
			},
		}),
		"harness_code": typeBlock("HarnessCode", map[string]*schema.Schema{
			"agent_file_name": defaultedString("Agents file to detect in repositories."),
		}),
		"harness_k8s": typeBlock("HarnessK8s", map[string]*schema.Schema{
			"service_discovery_agent_identifier": requiredString("Service discovery agent."),
			"environment_identifier":             requiredString("Environment."),
		}),
	}
}

func integrationStatusSchema() map[string]*schema.Schema {
	computed := func(t schema.ValueType, desc string) *schema.Schema {
		return &schema.Schema{Description: desc, Type: t, Computed: true}
	}
	return map[string]*schema.Schema{
		"uid":               computed(schema.TypeString, "Internal key; new on every create."),
		"revision":          computed(schema.TypeInt, "Revision, sent as If-Match on writes."),
		"health":            computed(schema.TypeString, "Health of the integration."),
		"mode_effective":    computed(schema.TypeString, "How the integration gets data: airbyte, platform, webhook or api."),
		"connector_version": computed(schema.TypeString, "Connector image tag recorded at create."),
		"effective_kinds": {
			Description: "Kinds that sync, including secondary kinds.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"endpoints": {
			Description: "URLs that receive data, per kind.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"kind":   computed(schema.TypeString, "Kind."),
				"method": computed(schema.TypeString, "HTTP method."),
				"url":    computed(schema.TypeString, "Path relative to the integration manager, e.g. https://app.harness.io/gateway/integration."),
			}},
		},
		"created": computed(schema.TypeInt, "Creation time, in milliseconds since the epoch."),
		"updated": computed(schema.TypeInt, "Last update time, in milliseconds since the epoch."),
	}
}

// integrationBody is the create and replace body.
func integrationBody(d *schema.ResourceData) map[string]any {
	specSchema := integrationSpecSchema()
	values := make(map[string]any, len(specSchema))
	for k := range specSchema {
		values[k] = d.Get(k)
	}
	typ := d.Get("type").(string)
	spec := expandObject(specSchema, values)
	for blockName, blockType := range integrationTypeBlocks {
		if blockType != typ {
			delete(spec, blockName)
		}
	}
	spec["type"] = typ
	if _, ok := spec["kinds"]; !ok {
		spec["kinds"] = []any{}
	}
	return map[string]any{
		"identifier":  d.Get("identifier").(string),
		"name":        d.Get("name").(string),
		"description": d.Get("description").(string),
		"tags":        expandTags(d.Get("tags").(*schema.Set).List()),
		"type":        typ,
		"enabled":     d.Get("enabled").(bool),
		"spec":        spec,
	}
}

// setIntegration sets the integration and, when mappings is non-nil, each kind's mapping.
func setIntegration(d *schema.ResourceData, in map[string]any, mappings map[string]map[string]any) error {
	d.SetId(str(in["identifier"]))
	set := map[string]any{
		"identifier":  in["identifier"],
		"name":        in["name"],
		"description": in["description"],
		"type":        in["type"],
		"enabled":     in["enabled"],
		"uid":         in["uid"],
		"revision":    num(in["revision"]),
		"created":     num(in["created"]),
		"updated":     num(in["updated"]),
		"org_id":      in["org"],
		"project_id":  in["project"],
	}
	tags, _ := in["tags"].(map[string]any)
	set["tags"] = flattenTags(tags)

	spec, _ := in["spec"].(map[string]any)
	for k, v := range flattenObject(integrationSpecSchema(), spec) {
		set[k] = v
	}
	if prior, ok := d.Get("kind").([]any); ok {
		set["kind"] = orderLike(prior, set["kind"])
	}
	if mappings != nil {
		kinds, _ := set["kind"].([]any)
		for _, k := range kinds {
			kind := k.(map[string]any)
			kind["mapping"] = []any{}
			if m, ok := mappings[str(kind["kind"])]; ok {
				kind["mapping"] = []any{kindMappingState(m)}
			}
		}
	}
	status, _ := in["status"].(map[string]any)
	set["health"] = status["health"]
	set["mode_effective"] = status["mode"]
	set["effective_kinds"] = status["effective_kinds"]
	set["endpoints"] = flattenObject(integrationStatusSchema(), status)["endpoints"]
	connector, _ := status["connector"].(map[string]any)
	set["connector_version"] = connector["version"]

	for k, v := range set {
		if err := d.Set(k, v); err != nil {
			return fmt.Errorf("setting %s: %w", k, err)
		}
	}
	return nil
}

// orderLike sorts the server's kinds into the order of prior; the server returns them sorted by name.
func orderLike(prior []any, got any) any {
	kinds, ok := got.([]any)
	if !ok || len(prior) == 0 {
		return got
	}
	rank := map[string]int{}
	for i, p := range prior {
		if m, ok := p.(map[string]any); ok {
			rank[str(m["kind"])] = i
		}
	}
	sorted := slices.Clone(kinds)
	slices.SortStableFunc(sorted, func(a, b any) int {
		ra, okA := rank[str(a.(map[string]any)["kind"])]
		rb, okB := rank[str(b.(map[string]any)["kind"])]
		switch {
		case okA && okB:
			return ra - rb
		case okA:
			return -1
		case okB:
			return 1
		}
		return 0
	})
	return sorted
}

func str(v any) string {
	s, _ := v.(string)
	return s
}

func num(v any) int {
	f, _ := v.(float64)
	return int(f)
}

func scope(d *schema.ResourceData) (string, string) {
	return d.Get("org_id").(string), d.Get("project_id").(string)
}

func resourceIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project := scope(d)
	body := integrationBody(d)
	if m := createMappings(d.Get("kind").([]any)); len(m) > 0 {
		body["mappings"] = m
	}
	out, err := c.createIntegration(ctx, org, project, body)
	if isIntegrationsStatus(err, http.StatusConflict) {
		return diag.Errorf("%s; if a timed-out apply created it, import it with terraform import", err)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(str(out["identifier"]))
	return refreshIntegration(ctx, c, d)
}

func resourceIntegrationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return refreshIntegration(ctx, newIntegrationsClient(meta.(*internal.Session)), d)
}

// refreshIntegration reads the integration and the mappings of kinds that have a mapping block.
func refreshIntegration(ctx context.Context, c *integrationsClient, d *schema.ResourceData) diag.Diagnostics {
	managed := kindMappings(d.Get("kind").([]any))
	org, project := scope(d)
	out, err := c.getIntegration(ctx, org, project, d.Id())
	if isIntegrationsStatus(err, http.StatusNotFound) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	mappings := map[string]map[string]any{}
	spec, _ := out["spec"].(map[string]any)
	kinds, _ := spec["kinds"].([]any)
	for _, k := range kinds {
		kind := str(k.(map[string]any)["kind"])
		if _, ok := managed[kind]; !ok {
			continue
		}
		m, err := c.getMapping(ctx, org, project, d.Id(), kind)
		if err != nil {
			return diag.FromErr(err)
		}
		mappings[kind] = m
	}
	return diag.FromErr(setIntegration(d, out, mappings))
}

// currentRevision reads the revision just before a write; mapping writes also change it.
func currentRevision(ctx context.Context, c *integrationsClient, org, project, id string) (int64, error) {
	cur, err := c.getIntegration(ctx, org, project, id)
	if err != nil {
		return 0, err
	}
	return int64(num(cur["revision"])), nil
}

func resourceIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project := scope(d)
	body := integrationBody(d)
	oldKinds, newKinds := d.GetChange("kind")
	prior, desired := kindMappings(oldKinds.([]any)), kindMappings(newKinds.([]any))
	// A kind switching to merge needs its correlation stored first; other mappings follow the integration.
	toMerge := switchingToMerge(oldKinds.([]any), newKinds.([]any))
	writeMappings := func(first bool) error {
		for kind, m := range desired {
			if p, ok := prior[kind]; (ok && reflect.DeepEqual(p, m)) || toMerge[kind] != first {
				continue
			}
			if _, err := writeMapping(ctx, c, org, project, d.Id(), kind, mappingBody(m, false)); err != nil {
				return fmt.Errorf("writing the mapping of kind %s: %w", kind, err)
			}
		}
		return nil
	}
	if err := writeMappings(true); err != nil {
		return diag.FromErr(err)
	}
	for attempt := 0; attempt < 2; attempt++ {
		revision, err := currentRevision(ctx, c, org, project, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = c.replaceIntegration(ctx, org, project, d.Id(), revision, body)
		if attempt == 0 && isIntegrationsStatus(err, http.StatusPreconditionFailed) {
			continue
		}
		if err != nil {
			return diag.FromErr(err)
		}
		break
	}
	if err := writeMappings(false); err != nil {
		return diag.FromErr(err)
	}
	// A removed mapping block resets the kind's mapping.
	var diags diag.Diagnostics
	for kind := range prior {
		if _, ok := desired[kind]; ok {
			continue
		}
		err := resetMapping(ctx, c, org, project, d.Id(), kind)
		switch {
		case err == nil, isIntegrationsStatus(err, http.StatusNotFound):
		case isIntegrationsStatus(err, http.StatusUnprocessableEntity):
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("The mapping of kind %s was left as it is", kind),
				Detail:   err.Error(),
			})
		default:
			return diag.FromErr(fmt.Errorf("resetting the mapping of kind %s: %w", kind, err))
		}
	}
	return append(diags, refreshIntegration(ctx, c, d)...)
}

// switchingToMerge returns the kinds whose action changes to merge.
func switchingToMerge(oldKinds, newKinds []any) map[string]bool {
	before := map[string]string{}
	for _, k := range oldKinds {
		kind, _ := k.(map[string]any)
		before[str(kind["kind"])] = str(kind["action"])
	}
	out := map[string]bool{}
	for _, k := range newKinds {
		kind, _ := k.(map[string]any)
		if a := str(kind["action"]); a == "merge" && before[str(kind["kind"])] != "merge" {
			out[str(kind["kind"])] = true
		}
	}
	return out
}

func resourceIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project := scope(d)
	revision, err := currentRevision(ctx, c, org, project, d.Id())
	if err == nil {
		err = c.deleteIntegration(ctx, org, project, d.Id(), revision)
	}
	if err != nil && !isIntegrationsStatus(err, http.StatusNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

// integrationCustomizeDiff fails the plan for changes the API rejects, rather than at apply.
func integrationCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ any) error {
	typ := d.Get("type").(string)
	for blockName, blockType := range integrationTypeBlocks {
		if blockType == typ {
			continue
		}
		if raw := d.GetRawConfig().GetAttr(blockName); raw.IsKnown() && !raw.IsNull() && raw.LengthInt() > 0 {
			return fmt.Errorf("%s is only allowed when type is %s", blockName, blockType)
		}
	}
	if d.Id() == "" || !d.HasChange("kind") {
		return nil
	}
	oldKinds, newKinds := d.GetChange("kind")
	if !sameKinds(oldKinds.([]any), newKinds.([]any)) {
		return fmt.Errorf("kinds can't be added or removed on an existing integration yet; recreate it instead")
	}
	if err := payloadSchemasUnchanged(oldKinds.([]any), newKinds.([]any)); err != nil {
		return err
	}
	airbyte := d.Get("mode_effective") == "airbyte"
	if err := mappingFieldsUnchanged(oldKinds.([]any), newKinds.([]any), airbyte); err != nil {
		return err
	}
	return entityTypesUnchanged(oldKinds.([]any), newKinds.([]any))
}

// entityTypesUnchanged rejects a changed idp_kind or idp_type, which are fixed at create.
func entityTypesUnchanged(oldKinds, newKinds []any) error {
	old := make(map[string]map[string]any, len(oldKinds))
	for _, k := range oldKinds {
		if m, ok := k.(map[string]any); ok {
			old[str(m["kind"])] = m
		}
	}
	for _, k := range newKinds {
		m, _ := k.(map[string]any)
		for _, field := range []string{"idp_kind", "idp_type"} {
			if v, prev := str(m[field]), str(old[str(m["kind"])][field]); v != "" && prev != "" && v != prev {
				return fmt.Errorf("%s of kind %s can't change from %q to %q; recreate the integration instead",
					field, str(m["kind"]), prev, v)
			}
		}
	}
	return nil
}

func sameKinds(a, b []any) bool {
	names := func(l []any) map[string]bool {
		out := make(map[string]bool, len(l))
		for _, k := range l {
			if m, ok := k.(map[string]any); ok {
				out[str(m["kind"])] = true
			}
		}
		return out
	}
	x, y := names(a), names(b)
	if len(x) != len(y) {
		return false
	}
	for k := range x {
		if !y[k] {
			return false
		}
	}
	return true
}
