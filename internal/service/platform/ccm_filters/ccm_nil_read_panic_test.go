package ccm_filters

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func capturePanic(fn func()) (panicked bool, value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			value = r
		}
	}()
	fn()
	return panicked, value
}

func TestCCM32488Class_ReadCCMFilterNilFilterProperties(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceCCMFilters().Schema, map[string]interface{}{
		"identifier": "id",
		"name":       "n",
		"type":       "CCMRecommendation",
	})
	panicked, val := capturePanic(func() {
		readCCMFilter(d, &nextgen.Filter{
			Identifier:       "id",
			Name:             "n",
			FilterProperties: nil,
		})
	})
	if panicked {
		t.Errorf("readCCMFilter panicked (bug confirmed): %v", val)
	} else {
		t.Log("RESULT: PASS (no panic)")
	}
}
