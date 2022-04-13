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

func DataSourceSshCredential() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving an SSH credential.",

		ReadContext: dataSourceSshCredentialRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:   "Unique identifier of the secret manager",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name"},
				AtLeastOneOf:  []string{"name", "id"},
			},
			"name": {
				Description:   "The name of the secret manager",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"id"},
				AtLeastOneOf:  []string{"name", "id"},
			},
			"usage_scope": usagescope.Schema(),
		},
	}
}

func dataSourceSshCredentialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	var sshCred *graphql.SSHCredential
	var err error

	if id := d.Get("id").(string); id != "" {
		sshCred, err = c.CDClient.SecretClient.GetSSHCredentialById(id)
	} else if name := d.Get("name").(string); name != "" {
		sshCred, err = c.CDClient.SecretClient.GetSSHCredentialByName(name)
	} else if err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if sshCred == nil {
		return diag.FromErr(errors.New("could not find ssh credential"))
	}

	d.SetId(sshCred.Id)
	d.Set("name", sshCred.Name)
	d.Set("usage_scope", usagescope.FlattenUsageScope(sshCred.UsageScope))

	return nil
}
