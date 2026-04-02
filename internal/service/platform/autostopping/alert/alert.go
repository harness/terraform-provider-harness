package alert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	entityTypeAutostopRule = "autostopping_rule"
	relationTypeAll        = "all"
	relationTypeSpecific   = "specific"
	recipientTypeEmail     = "email"
	recipientTypeSlack     = "slack"
)

// validateRecipientsBlock ensures at least one of email or slack is set with at least one value.
// Used by CustomizeDiff (plan-time) and buildAlertRequest (apply-time).
func validateRecipientsBlock(recList interface{}) error {
	list, ok := recList.([]interface{})
	if !ok || len(list) == 0 {
		return fmt.Errorf("recipients block requires at least one of email or slack with at least one value")
	}
	block, ok := list[0].(map[string]interface{})
	if !ok {
		return fmt.Errorf("recipients block requires at least one of email or slack with at least one value")
	}
	hasEmail := false
	if e, ok := block["email"]; ok && e != nil {
		if arr, ok := e.([]interface{}); ok && len(arr) > 0 {
			for _, v := range arr {
				if s, ok := v.(string); ok && s != "" {
					hasEmail = true
					break
				}
			}
		}
	}
	hasSlack := false
	if s, ok := block["slack"]; ok && s != nil {
		if arr, ok := s.([]interface{}); ok && len(arr) > 0 {
			for _, v := range arr {
				if str, ok := v.(string); ok && str != "" {
					hasSlack = true
					break
				}
			}
		}
	}
	if !hasEmail && !hasSlack {
		return fmt.Errorf("recipients block requires at least one of email or slack with at least one value")
	}
	return nil
}

// validateApplicableToAllAndRuleIDList ensures applicable_to_all_rules and rule_id_list are mutually exclusive
// and that rule_id_list has at least one value when applicable_to_all_rules is false.
// Used by CustomizeDiff (plan-time) and buildAlertRequest (apply-time).
func validateApplicableToAllAndRuleIDList(applicableToAll bool, ruleIDList interface{}) error {
	var list []interface{}
	if v, ok := ruleIDList.([]interface{}); ok && v != nil {
		list = v
	}
	if applicableToAll && len(list) > 0 {
		return fmt.Errorf("applicable_to_all_rules and rule_id_list are mutually exclusive: set applicable_to_all_rules to true (and leave rule_id_list empty) or set applicable_to_all_rules to false and provide rule_id_list")
	}
	if !applicableToAll && len(list) == 0 {
		return fmt.Errorf("rule_id_list is required when applicable_to_all_rules is false and must have at least one AutoStopping rule ID")
	}
	return nil
}

// ruleIDFromSchema converts a schema value (int, int64, or float64 from Terraform) to int.
func ruleIDFromSchema(v interface{}) (int, error) {
	switch n := v.(type) {
	case int:
		return n, nil
	case int64:
		return int(n), nil
	case float64:
		return int(n), nil
	default:
		return 0, fmt.Errorf("rule_id_list must contain integers, got %T", v)
	}
}

func ResourceAlert() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating and managing Harness AutoStopping alerts. Alerts notify users via email or Slack when events such as warmup failures, cooldown failures, or rule lifecycle changes occur.",
		ReadContext:   resourceAlertRead,
		CreateContext: resourceAlertCreate,
		UpdateContext: resourceAlertUpdate,
		DeleteContext: resourceAlertDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
			return validateSchema(d)
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the alert.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"enabled": {
				Description: "Whether the alert is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"recipients": {
				Description: "Notification recipients. At least one of `email` or `slack` is required (with at least one value).",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "List of email addresses to notify. Required if `slack` is not set.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"slack": {
							Description: "List of Slack webhook URLs or channel identifiers to notify. Required if `email` is not set.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"events": {
				Description: "List of event types that trigger the alert (e.g. autostopping_rule_created, autostopping_warmup_failed, autostopping_cooldown_failed).",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"applicable_to_all_rules": {
				Description: "When true, the alert applies to all AutoStopping rules in the account (leave `rule_id_list` empty). Mutually exclusive with `rule_id_list`.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"rule_id_list": {
				Description: "List of AutoStopping rule IDs to apply the alert to. Required when `applicable_to_all_rules` is false. Mutually exclusive with `applicable_to_all_rules` = true.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func validateSchema(d *schema.ResourceDiff) error {
	if err := validateRecipientsBlock(d.Get("recipients")); err != nil {
		return err
	}
	applicableToAll := false
	if v, ok := d.GetOk("applicable_to_all_rules"); ok {
		applicableToAll = v.(bool)
	}
	var ruleIDList interface{}
	if v, ok := d.GetOk("rule_id_list"); ok {
		ruleIDList = v
	}
	return validateApplicableToAllAndRuleIDList(applicableToAll, ruleIDList)
}

func resourceAlertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	return alertReadByID(ctx, d, meta, id)
}

// alertReadByID fetches an alert by id and flattens it into the resource state.
// Used by the resource only; the data source looks up by name and populates id from the server.
func alertReadByID(ctx context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AutoStoppingAlertsApi.GetAlert(ctx, c.AccountId, id, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Response == nil {
		d.SetId("")
		return nil
	}

	d.SetId(id)
	return flattenAlert(d, resp.Response)
}

func resourceAlertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	req, diags := buildAlertRequest(d, "")
	if diags.HasError() {
		return diags
	}

	createResp, httpResp, err := c.AutoStoppingAlertsApi.CreateAlert(ctx, *req, c.AccountId, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(createResp.Response)
	return resourceAlertRead(ctx, d, meta)
}

func resourceAlertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	updateReq, diags := buildAlertUpdateRequest(d)
	if diags.HasError() {
		return diags
	}

	_, httpResp, err := c.AutoStoppingAlertsApi.UpdateAlert(ctx, *updateReq, c.AccountId, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return resourceAlertRead(ctx, d, meta)
}

func resourceAlertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	_, httpResp, err := c.AutoStoppingAlertsApi.DeleteAlert(ctx, c.AccountId, id, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

// buildAlertRequest builds AlertRequest from the resource data. id is empty for create.
func buildAlertRequest(d *schema.ResourceData, id string) (*nextgen.AlertRequest, diag.Diagnostics) {
	req := &nextgen.AlertRequest{
		Id:                 id,
		Name:               d.Get("name").(string),
		Enabled:            d.Get("enabled").(bool),
		EntityType:         entityTypeAutostopRule,
		Recipients:         nil,
		Events:             nil,
		AssociatedEntities: nil,
	}

	// Recipients: at least one of email or slack (validated at plan time via CustomizeDiff; re-check at apply)
	recList := d.Get("recipients").([]interface{})
	if err := validateRecipientsBlock(recList); err != nil {
		return nil, diag.FromErr(err)
	}
	rec := recList[0].(map[string]interface{})
	emailRaw, hasEmail := rec["email"]
	slackRaw, hasSlack := rec["slack"]

	var emails, slacks []string
	if hasEmail && emailRaw != nil {
		for _, v := range emailRaw.([]interface{}) {
			if s, ok := v.(string); ok && s != "" {
				emails = append(emails, s)
			}
		}
	}
	if hasSlack && slackRaw != nil {
		for _, v := range slackRaw.([]interface{}) {
			if s, ok := v.(string); ok && s != "" {
				slacks = append(slacks, s)
			}
		}
	}
	if len(emails) > 0 {
		req.Recipients = append(req.Recipients, nextgen.AlertRecipient{Type_: recipientTypeEmail, Ids: emails})
	}
	if len(slacks) > 0 {
		req.Recipients = append(req.Recipients, nextgen.AlertRecipient{Type_: recipientTypeSlack, Ids: slacks})
	}

	// Events
	eventsRaw := d.Get("events").([]interface{})
	for _, e := range eventsRaw {
		ev, ok := e.(string)
		if !ok || ev == "" {
			continue
		}
		req.Events = append(req.Events, nextgen.AlertEvent{
			Id:    ev,
			Event: ev,
		})
	}
	if len(req.Events) == 0 {
		return nil, diag.Errorf("at least one event is required")
	}

	// Associated entities: either all rules or specific rule IDs (validated at plan time via CustomizeDiff; re-check at apply)
	applicableToAll := d.Get("applicable_to_all_rules").(bool)
	var ruleIDList []interface{}
	if v, ok := d.GetOk("rule_id_list"); ok && v != nil {
		ruleIDList = v.([]interface{})
	}
	if err := validateApplicableToAllAndRuleIDList(applicableToAll, ruleIDList); err != nil {
		return nil, diag.FromErr(err)
	}

	if applicableToAll {
		req.AssociatedEntities = []nextgen.AlertEntity{{
			RelationType: relationTypeAll,
		}}
	} else {
		for _, r := range ruleIDList {
			ruleID, err := ruleIDFromSchema(r)
			if err != nil {
				return nil, diag.FromErr(err)
			}
			req.AssociatedEntities = append(req.AssociatedEntities, nextgen.AlertEntity{
				RelationType: relationTypeSpecific,
				EntityId:     strconv.Itoa(ruleID),
			})
		}
	}

	return req, nil
}

func buildAlertUpdateRequest(d *schema.ResourceData) (*nextgen.AlertUpdateRequest, diag.Diagnostics) {
	req, diags := buildAlertRequest(d, d.Id())
	if diags != nil && diags.HasError() {
		return nil, diags
	}
	return &nextgen.AlertUpdateRequest{
		Id:                 req.Id,
		Name:               req.Name,
		Enabled:            req.Enabled,
		EntityType:         req.EntityType,
		Recipients:         req.Recipients,
		Events:             req.Events,
		AssociatedEntities: req.AssociatedEntities,
	}, nil
}

func flattenAlert(d *schema.ResourceData, a *nextgen.Alert) diag.Diagnostics {
	if a == nil {
		return nil
	}

	d.SetId(a.Id)
	d.Set("identifier", a.Id)
	if err := d.Set("name", a.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", a.Enabled); err != nil {
		return diag.FromErr(err)
	}

	// Recipients: single block with email and slack lists
	recMap := map[string]interface{}{
		"email": []string{},
		"slack": []string{},
	}
	for _, r := range a.Recipients {
		if r.Type_ == recipientTypeEmail {
			recMap["email"] = r.Ids
		}
		if r.Type_ == recipientTypeSlack {
			recMap["slack"] = r.Ids
		}
	}
	if err := d.Set("recipients", []interface{}{recMap}); err != nil {
		return diag.FromErr(err)
	}

	// Events: list of event strings
	events := make([]interface{}, 0, len(a.Events))
	for _, e := range a.Events {
		events = append(events, e.Event)
	}
	if err := d.Set("events", events); err != nil {
		return diag.FromErr(err)
	}

	// Associated entities -> applicable_to_all_rules and rule_id_list
	var ruleIDs []interface{}
	applicableToAll := false
	for _, ae := range a.AssociatedEntities {
		if ae.RelationType == relationTypeAll {
			applicableToAll = true
			break
		}
		if ae.RelationType == relationTypeSpecific && ae.EntityId != "" {
			id, err := strconv.Atoi(ae.EntityId)
			if err != nil {
				return diag.FromErr(fmt.Errorf("invalid rule entity_id: %w", err))
			}
			ruleIDs = append(ruleIDs, id)
		}
	}
	if err := d.Set("applicable_to_all_rules", applicableToAll); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_id_list", ruleIDs); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
