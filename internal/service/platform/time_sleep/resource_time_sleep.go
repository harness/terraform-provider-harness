package time_sleep

import (
	"context"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validDuration = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?(ms|s|m|h)$`)

func ResourceTimeSleep() *schema.Resource {
	return &schema.Resource{
		Description: "Resource that introduces a time delay during create and/or destroy lifecycle phases. " +
			"Useful for waiting on eventually-consistent upstream resources before proceeding.",

		CreateContext: resourceTimeSleepCreate,
		ReadContext:   resourceTimeSleepRead,
		UpdateContext: resourceTimeSleepUpdate,
		DeleteContext: resourceTimeSleepDelete,

		Schema: map[string]*schema.Schema{
			"create_duration": {
				Description: "Duration to sleep when the resource is created. " +
					"Accepts Go duration strings: ms (milliseconds), s (seconds), m (minutes), h (hours). " +
					"Example: \"30s\", \"5m\". At least one of create_duration or destroy_duration must be set.",
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					validDuration,
					"must be a positive number followed by ms, s, m, or h — e.g. \"30s\" or \"5m\"",
				),
			},
			"destroy_duration": {
				Description: "Duration to sleep when the resource is destroyed. " +
					"Accepts Go duration strings: ms (milliseconds), s (seconds), m (minutes), h (hours). " +
					"Example: \"30s\", \"5m\". At least one of create_duration or destroy_duration must be set.",
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					validDuration,
					"must be a positive number followed by ms, s, m, or h — e.g. \"30s\" or \"5m\"",
				),
			},
			"triggers": {
				Description: "Map of arbitrary string values. Any change to this map forces replacement of the resource, " +
					"causing the create and destroy delays to fire again.",
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceTimeSleepCreate(ctx context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	createDur := d.Get("create_duration").(string)
	destroyDur := d.Get("destroy_duration").(string)

	if createDur == "" && destroyDur == "" {
		return diag.Errorf("at least one of create_duration or destroy_duration must be set")
	}

	if createDur != "" {
		dur, err := time.ParseDuration(createDur)
		if err != nil {
			return diag.Errorf("invalid create_duration %q: %s", createDur, err)
		}
		select {
		case <-ctx.Done():
			return diag.Errorf("interrupted during create sleep: %s", ctx.Err())
		case <-time.After(dur):
		}
	}

	d.SetId(time.Now().UTC().Format(time.RFC3339Nano))
	return nil
}

func resourceTimeSleepRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourceTimeSleepUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourceTimeSleepDelete(ctx context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	destroyDur := d.Get("destroy_duration").(string)

	if destroyDur != "" {
		dur, err := time.ParseDuration(destroyDur)
		if err != nil {
			return diag.Errorf("invalid destroy_duration %q: %s", destroyDur, err)
		}
		select {
		case <-ctx.Done():
			return diag.Errorf("interrupted during destroy sleep: %s", ctx.Err())
		case <-time.After(dur):
		}
	}

	return nil
}

