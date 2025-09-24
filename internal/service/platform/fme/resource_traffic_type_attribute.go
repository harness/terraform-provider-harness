package fme

import (
	"context"
	"strings"
	"regexp"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMETrafficTypeAttribute() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME traffic type attribute.",
		ReadContext:   resourceFMETrafficTypeAttributeRead,
		CreateContext: resourceFMETrafficTypeAttributeCreate,
		UpdateContext: resourceFMETrafficTypeAttributeUpdate,
		DeleteContext: resourceFMETrafficTypeAttributeDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the traffic type attribute.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"workspace_id": {
				Description:  "ID of the workspace this traffic type attribute belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"traffic_type_id": {
				Description:  "ID of the traffic type this attribute belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"attribute_id": {
				Description: "ID of the attribute (derived from display_name).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_name": {
				Description:  "Display name of the traffic type attribute.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"description": {
				Description: "Description of the traffic type attribute.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"data_type": {
				Description:  "Data type of the attribute (STRING, NUMBER, BOOLEAN, etc.).",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"STRING", "NUMBER", "BOOLEAN", "DATETIME", "SET"}, false),
			},
			"is_searchable": {
				Description: "Whether the attribute is searchable.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"suggested_values": {
				Description: "List of suggested values for this attribute.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"organization_id": {
				Description: "Organization ID of the traffic type attribute.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMETrafficTypeAttributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	displayName := d.Get("display_name").(string)

	// Generate attribute ID from display name (convert to valid identifier)
	attributeID := generateAttributeID(displayName)

	req := &api.TrafficTypeAttributeCreateRequest{
		ID:          attributeID,
		DisplayName: displayName,
		DataType:    d.Get("data_type").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		desc := description.(string)
		req.Description = &desc
	}

	if isSearchable, ok := d.GetOk("is_searchable"); ok {
		searchable := isSearchable.(bool)
		req.IsSearchable = &searchable
	}

	if suggestedValues, ok := d.GetOk("suggested_values"); ok {
		values := suggestedValues.([]interface{})
		req.SuggestedValues = make([]string, len(values))
		for i, v := range values {
			req.SuggestedValues[i] = v.(string)
		}
	}

	trafficTypeID := d.Get("traffic_type_id").(string)
	attribute, err := c.APIClient.TrafficTypeAttributes.Create(workspaceID, trafficTypeID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*attribute.ID)
	d.Set("attribute_id", attributeID)
	return resourceFMETrafficTypeAttributeRead(ctx, d, meta)
}

func resourceFMETrafficTypeAttributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Get("traffic_type_id").(string)
	attributeID := d.Id()

	attribute, err := c.APIClient.TrafficTypeAttributes.Get(workspaceID, trafficTypeID, attributeID)
	if err != nil {
		return diag.FromErr(err)
	}

	if attribute == nil {
		d.SetId("")
		return nil
	}

	d.Set("traffic_type_id", attribute.TrafficTypeID)
	d.Set("display_name", attribute.DisplayName)
	d.Set("description", attribute.Description)
	d.Set("data_type", attribute.DataType)
	d.Set("is_searchable", attribute.IsSearchable)
	d.Set("suggested_values_for", attribute.SuggestedValuesFor)
	d.Set("suggested_values_json", attribute.SuggestedValuesJSON)
	d.Set("organization_id", attribute.OrganizationID)

	return nil
}

func resourceFMETrafficTypeAttributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	attributeID := d.Id()

	req := &api.TrafficTypeAttributeUpdateRequest{
		DisplayName: d.Get("display_name").(string),
		DataType:    d.Get("data_type").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		desc := description.(string)
		req.Description = &desc
	}

	if isSearchable, ok := d.GetOk("is_searchable"); ok {
		searchable := isSearchable.(bool)
		req.IsSearchable = &searchable
	}

	if suggestedValuesFor, ok := d.GetOk("suggested_values_for"); ok {
		values := suggestedValuesFor.(string)
		req.SuggestedValuesFor = &values
	}

	if suggestedValuesJSON, ok := d.GetOk("suggested_values_json"); ok {
		valuesJSON := suggestedValuesJSON.(string)
		req.SuggestedValuesJSON = &valuesJSON
	}

	_, err := c.APIClient.TrafficTypeAttributes.Update(workspaceID, attributeID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceFMETrafficTypeAttributeRead(ctx, d, meta)
}

func resourceFMETrafficTypeAttributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Get("traffic_type_id").(string)
	attributeID := d.Id()

	err := c.APIClient.TrafficTypeAttributes.Delete(workspaceID, trafficTypeID, attributeID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// generateAttributeID converts a display name to a valid attribute ID
func generateAttributeID(displayName string) string {
	// Convert to lowercase and replace spaces/special chars with underscores
	id := strings.ToLower(displayName)
	// Replace non-alphanumeric characters with underscores
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	id = reg.ReplaceAllString(id, "_")
	// Remove leading/trailing underscores
	id = strings.Trim(id, "_")
	return id
}