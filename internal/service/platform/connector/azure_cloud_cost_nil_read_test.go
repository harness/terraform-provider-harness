package connector

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestConnectorAzureCloudCost_CCM32488_ReadNilBillingExportSpec guards the
// CCM-32488 fix: read must not panic when BILLING is enabled and only one of
// billing_export_spec / billing_export_spec2 is present in the API response.
func TestConnectorAzureCloudCost_CCM32488_ReadNilBillingExportSpec(t *testing.T) {
	resourceSchema := ResourceConnectorAzureCloudCost().Schema
	cases := []struct {
		name string
		conn *nextgen.ConnectorInfo
	}{
		{
			name: "billingExportSpec2 null (CCM-32488 scenario)",
			conn: &nextgen.ConnectorInfo{
				AzureCloudCost: &nextgen.CeAzureConnector{
					TenantId:        "tenant",
					SubscriptionId:  "sub",
					FeaturesEnabled: []string{"BILLING"},
					BillingExportSpec: &nextgen.BillingExportSpec{
						StorageAccountName: "sa",
						ContainerName:      "c",
						ReportName:         "r",
					},
					BillingExportSpec2: nil,
				},
			},
		},
		{
			name: "billingExportSpec null",
			conn: &nextgen.ConnectorInfo{
				AzureCloudCost: &nextgen.CeAzureConnector{
					TenantId:          "tenant",
					SubscriptionId:    "sub",
					FeaturesEnabled:   []string{"BILLING"},
					BillingExportSpec: nil,
					BillingExportSpec2: &nextgen.BillingExportSpec{
						StorageAccountName: "sa2",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			panicked, val := captureAzureCloudCostReadPanic(func() {
				_ = readConnectorAzureCloudCost(d, tc.conn)
			})
			if panicked {
				t.Fatalf("read panicked (CCM-32488 regression): %v", val)
			}
		})
	}
}

func captureAzureCloudCostReadPanic(fn func()) (panicked bool, value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			value = r
		}
	}()
	fn()
	return panicked, value
}
