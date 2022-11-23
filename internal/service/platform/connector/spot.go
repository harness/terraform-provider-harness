package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorSpot() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Spot connector.",
		ReadContext:   resourceConnectorSpotRead,
		CreateContext: resourceConnectorSpotCreateOrUpdate,
		UpdateContext: resourceConnectorSpotCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"permanent_token": {
				Description: "Authenticate to Spot using account id and permanent token.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spot_account_id": {
							Description:   "Spot account id.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"permanent_token.0.spot_account_id_ref"},
							AtLeastOneOf:  []string{"permanent_token.0.spot_account_id", "permanent_token.0.spot_account_id_ref"},
						},
						"spot_account_id_ref": {
							Description:   "Reference to the Harness secret containing the spot account id." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"permanent_token.0.spot_account_id"},
							AtLeastOneOf:  []string{"permanent_token.0.spot_account_id", "permanent_token.0.spot_account_id_ref"},
						},
						"api_token_ref": {
							Description: "Reference to the Harness secret containing the permanent api token." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"delegate_selectors": {
							Description: "Connect only using delegates with these tags.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"execute_on_delegate": {
							Description: "Execute on delegate or not.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorSpotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Spot)
	if err != nil {
		return err
	}

	if err := readConnectorSpot(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorSpotCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorSpot(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorSpot(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorSpot(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Spot,
		Spot: &nextgen.SpotConnector{
			Credential: &nextgen.SpotCredential{},
		},
	}

	if attr, ok := d.GetOk("permanent_token"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Spot.Credential.Type_ = nextgen.SpotAuthTypes.PermanentTokenConfig
		connector.Spot.Credential.PermanentTokenConfig = &nextgen.SpotPermanentTokenConfigSpec{}

		if attr := config["spot_account_id"].(string); attr != "" {
			connector.Spot.Credential.PermanentTokenConfig.SpotAccountId = attr
		}

		if attr := config["spot_account_id_ref"].(string); attr != "" {
			connector.Spot.Credential.PermanentTokenConfig.SpotAccountIdRef = attr
		}

		if attr := config["api_token_ref"].(string); attr != "" {
			connector.Spot.Credential.PermanentTokenConfig.ApiTokenRef = attr
		}

		if attr := config["execute_on_delegate"].(bool); attr {
			connector.Spot.ExecuteOnDelegate = attr
		}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Spot.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	return connector
}

func readConnectorSpot(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	switch connector.Spot.Credential.Type_ {
	case nextgen.SpotAuthTypes.PermanentTokenConfig:
		d.Set("permanent_token", []map[string]interface{}{
			{
				"spot_account_id":     connector.Spot.Credential.PermanentTokenConfig.SpotAccountId,
				"spot_account_id_ref": connector.Spot.Credential.PermanentTokenConfig.SpotAccountIdRef,
				"api_token_ref":       connector.Spot.Credential.PermanentTokenConfig.ApiTokenRef,
				"execute_on_delegate": connector.Spot.ExecuteOnDelegate,
				"delegate_selectors":  connector.Spot.DelegateSelectors,
			},
		})
	default:
		return fmt.Errorf("unsupported spot credential type: %s", connector.Spot.Credential.Type_)
	}

	return nil
}
