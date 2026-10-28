package ccm_filters

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestReadCCMFilter_NilFilterProperties(t *testing.T) {
	resource := ResourceCCMFilters()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	readCCMFilter(d, &nextgen.Filter{
		Identifier:       "filter_id",
		Name:             "filter name",
		FilterVisibility: "EveryOne",
	})

	if got := d.Id(); got != "filter_id" {
		t.Fatalf("id = %#v", got)
	}
	if got := d.Get("name"); got != "filter name" {
		t.Fatalf("name = %#v", got)
	}
	if got := d.Get("filter_properties").([]interface{}); len(got) != 0 {
		t.Fatalf("filter_properties = %#v, want empty", got)
	}
}

func TestReadCCMFilter_NilFilter(t *testing.T) {
	resource := ResourceCCMFilters()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	readCCMFilter(d, nil)

	if got := d.Id(); got != "" {
		t.Fatalf("id = %#v, want empty", got)
	}
}

func TestReadCCMFilter_PopulatesFilterProperties(t *testing.T) {
	resource := ResourceCCMFilters()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	readCCMFilter(d, &nextgen.Filter{
		Identifier: "filter_id",
		Name:       "filter name",
		FilterProperties: &nextgen.FilterProperties{
			FilterType: "CCMRecommendation",
		},
	})

	if got := d.Get("type"); got != "CCMRecommendation" {
		t.Fatalf("type = %#v", got)
	}
	if got := d.Get("filter_properties").([]interface{}); len(got) != 1 {
		t.Fatalf("filter_properties = %#v, want one entry", got)
	}
}
