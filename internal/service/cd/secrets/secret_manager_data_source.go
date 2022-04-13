package secrets

import (
	"context"
	"errors"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/cd/usagescope"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretManager() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness secret manager",

		ReadContext: dataSourceSecretManagerRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:   "Unique identifier of the secret manager",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "default"},
			},
			"name": {
				Description:   "The name of the secret manager",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"id", "default"},
			},
			"default": {
				Description:   "True to lookup the id of the default secret manager",
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"id", "name"},
			},
			"usage_scope": usagescope.Schema(),
		},
	}
}

func dataSourceSecretManagerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	// If we're looking up the default secret manager, then we'll just get the
	// ID of the default secret manager and not set any other fields.
	// This is done because the default secret manager details are not
	// available through the API
	if useDefault := d.Get("default").(bool); useDefault {
		id, err := c.CDClient.SecretClient.GetDefaultSecretManagerId()
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(id)
		return nil
	}

	var sm *graphql.SecretManager
	var err error

	if id := d.Get("id").(string); id != "" {
		sm, err = c.CDClient.SecretClient.GetSecretManagerById(id)
	} else if name := d.Get("name").(string); name != "" {
		sm, err = c.CDClient.SecretClient.GetSecretManagerByName(name)
	} else if err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if sm == nil {
		return diag.FromErr(errors.New("could not find secret manager"))
	}

	d.SetId(sm.Id)
	d.Set("name", sm.Name)
	d.Set("usage_scope", usagescope.FlattenUsageScope(sm.UsageScope))

	return nil
}
