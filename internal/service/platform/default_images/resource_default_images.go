package default_images

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceDefaultImages() *schema.Resource {

	return &schema.Resource{
		Description:   "Resource for managing a Harness execution image override for CI, IACM, or IDP.",
		CreateContext: resourceDefaultImagesCreate,
		ReadContext:   resourceDefaultImagesRead,
		UpdateContext: resourceDefaultImagesUpdate,
		DeleteContext: resourceDefaultImagesDelete,
		Schema: map[string]*schema.Schema{
			"kind": {
				Description:  "The service kind. Supported values: `ci`, `iacm`, `idp`.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ci", "iacm", "idp"}, false),
			},
			"infra_type": {
				Description: "The infrastructure type passed to the execution config API (e.g. `K8`, `VM`). Defaults to `K8`.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "K8",
				ForceNew:    true,
			},
			"field": {
				Description: "The image field name to override (e.g. `addonTag`, `liteEngineTag`).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"value": {
				Description: "The image tag value to set. When omitted or set to `null`, the field override is reset to the Harness default.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceDefaultImagesCreate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	return resourceDefaultImagesUpdate(ctx, d, meta)
}

func resourceDefaultImagesUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {

	kind := d.Get("kind").(string)
	infraType := d.Get("infra_type").(string)
	field := d.Get("field").(string)
	value := d.Get("value").(string)

	s := meta.(*internal.Session)
	var err error
	if value == "" {
		err = execResetConfig(ctx, s, kind, infraType, field)
	} else {
		err = execUpdateConfig(ctx, s, kind, infraType, field, value)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", kind, infraType, field))
	return resourceDefaultImagesRead(ctx, d, meta)
}

func resourceDefaultImagesRead(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {

	kind := d.Get("kind").(string)
	infraType := d.Get("infra_type").(string)
	field := d.Get("field").(string)

	s := meta.(*internal.Session)
	data, err := execGetCustomerConfig(ctx, s, kind, infraType)
	if err != nil {
		return diag.FromErr(err)
	}

	if v, ok := data[field]; ok && v != "" {
		if err := d.Set("value", v); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("value", ""); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceDefaultImagesDelete(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {

	kind := d.Get("kind").(string)
	infraType := d.Get("infra_type").(string)
	field := d.Get("field").(string)

	if err := execResetConfig(ctx, meta.(*internal.Session), kind, infraType, field); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func newIdpExecCfgClient(s *internal.Session) *idp.APIClient {
	cfg := idp.NewConfiguration()
	if s.IDPClient != nil {
		cfg.AccountId = s.IDPClient.AccountId
		cfg.ApiKey = s.IDPClient.ApiKey
		if s.IDPClient.Endpoint != "" {
			cfg.BasePath = s.IDPClient.Endpoint
		}
	}
	return idp.NewAPIClient(cfg)
}

func execUpdateConfig(ctx context.Context, s *internal.Session, kind,
	infraType, field, value string) error {

	switch kind {
	case "ci":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		_, err := c.CiExecutionConfigApi.UpdateConfig(authCtx, infraType,
			[]nextgen.ExecutionConfigUpdate{{Field: field, Value: value}})
		return err
	case "iacm":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		_, err := c.IacmExecutionConfigApi.UpdateConfig(authCtx, infraType,
			[]nextgen.ExecutionConfigUpdate{{Field: field, Value: value}})
		return err
	case "idp":
		c := newIdpExecCfgClient(s)
		c, authCtx := c.WithAuthContext(ctx)
		_, err := c.ExecutionConfigApi.UpdateConfig(authCtx, infraType,
			[]idp.ExecutionConfigUpdate{{Field: field, Value: value}})
		return err
	default:
		return fmt.Errorf("unsupported kind %q", kind)
	}
}

func execResetConfig(ctx context.Context, s *internal.Session, kind, infraType, field string) error {

	switch kind {
	case "ci":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		_, err := c.CiExecutionConfigApi.ResetConfig(authCtx, infraType,
			[]nextgen.ExecutionConfigUpdate{{Field: field}})
		return err
	case "iacm":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		_, err := c.IacmExecutionConfigApi.ResetConfig(authCtx, infraType,
			[]nextgen.ExecutionConfigUpdate{{Field: field}})
		return err
	case "idp":
		c := newIdpExecCfgClient(s)
		c, authCtx := c.WithAuthContext(ctx)
		_, err := c.ExecutionConfigApi.ResetConfig(authCtx, infraType,
			[]idp.ExecutionConfigUpdate{{Field: field}})
		return err
	default:
		return fmt.Errorf("unsupported kind %q", kind)
	}
}

func execGetCustomerConfig(ctx context.Context, s *internal.Session, kind,
	infraType string) (map[string]string, error) {

	switch kind {
	case "ci":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		resp, err := c.CiExecutionConfigApi.GetCustomerConfig(authCtx, infraType, true)
		if err != nil {
			return nil, err
		}
		return map[string]string(resp.Data), nil
	case "iacm":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		resp, err := c.IacmExecutionConfigApi.GetCustomerConfig(authCtx, infraType, true)
		if err != nil {
			return nil, err
		}
		return map[string]string(resp.Data), nil
	case "idp":
		c := newIdpExecCfgClient(s)
		c, authCtx := c.WithAuthContext(ctx)
		resp, err := c.ExecutionConfigApi.GetCustomerConfig(authCtx, infraType, true)
		if err != nil {
			return nil, err
		}
		return map[string]string(resp.Data), nil
	default:
		return nil, fmt.Errorf("unsupported kind %q", kind)
	}
}
