package split

import (
	"context"
	"strings"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMETrafficTypeAttribute manages a schema attribute for a traffic type.
func ResourceFMETrafficTypeAttribute() *schema.Resource {
	return &schema.Resource{
		Description: "Create, update, and delete a Harness FME (Split) traffic type attribute. Import id format: `org_id/project_id/traffic_type_id/attribute_id`.",

		CreateContext: resourceFMETrafficTypeAttributeCreate,
		ReadContext:   resourceFMETrafficTypeAttributeRead,
		UpdateContext: resourceFMETrafficTypeAttributeUpdate,
		DeleteContext: resourceFMETrafficTypeAttributeDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMETrafficTypeAttributeImport,
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Harness organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Harness project identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"traffic_type_id": {
				Description: "Split traffic type ID.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"identifier": {
				Description: "Attribute identifier (id) in Split.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"display_name": {
				Description: "Display name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"data_type": {
				Description: "Data type: e.g. `string`, `number`, `datetime`, `set` (normalized to lowercase in state; Split may return uppercase). Config casing is ignored for diffs.",
				Type:        schema.TypeString,
				Required:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(strings.TrimSpace(old), strings.TrimSpace(new))
				},
			},
			"is_searchable": {
				Description: "Whether the attribute is searchable.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"suggested_values": {
				Description: "Suggested values for set types.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"attribute_id": {
				Description: "The attribute ID returned by Split (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMETrafficTypeAttributeImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, ttID, attrID, err := ParseImportID4(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return nil, err
	}
	if err := d.Set("identifier", attrID); err != nil {
		return nil, err
	}
	d.SetId(attrID)
	return []*schema.ResourceData{d}, nil
}

func attrRequestFromResourceData(d *schema.ResourceData) splitsdk.AttributeRequest {
	req := splitsdk.AttributeRequest{
		Identifier:    d.Get("identifier").(string),
		DisplayName:   d.Get("display_name").(string),
		Description:   d.Get("description").(string),
		DataType:      d.Get("data_type").(string),
		TrafficTypeID: d.Get("traffic_type_id").(string),
	}
	if v, ok := d.GetOk("suggested_values"); ok {
		for _, x := range v.([]interface{}) {
			req.SuggestedValues = append(req.SuggestedValues, x.(string))
		}
	}
	s := d.Get("is_searchable").(bool)
	req.IsSearchable = &s
	return req
}

func resourceFMETrafficTypeAttributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	attr, err := client.Attributes.Create(wsID, ttID, attrRequestFromResourceData(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(attr.ID)
	return resourceFMETrafficTypeAttributeRead(ctx, d, meta)
}

func resourceFMETrafficTypeAttributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	attr, err := client.Attributes.FindByID(wsID, ttID, d.Id(), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	if attr == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("display_name", attr.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", attr.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_type", strings.ToLower(strings.TrimSpace(attr.DataType))); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_searchable", attr.IsSearchable); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("suggested_values", attr.SuggestedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("attribute_id", attr.ID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMETrafficTypeAttributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	_, err := client.Attributes.Update(wsID, ttID, d.Id(), attrRequestFromResourceData(d))
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceFMETrafficTypeAttributeRead(ctx, d, meta)
}

func resourceFMETrafficTypeAttributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	if err := client.Attributes.Delete(wsID, ttID, d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
