package connector

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

func TestCCM32488Class_ReadNilNestedFields(t *testing.T) {
	cases := []struct {
		name        string
		run         func(d *schema.ResourceData)
		resourceSch map[string]*schema.Schema
	}{
		{
			name: "gcp_cloud_cost: BILLING enabled, billingExportSpec null",
			resourceSch: func() map[string]*schema.Schema {
				r := ResourceConnectorGCPCloudCost()
				return r.Schema
			}(),
			run: func(d *schema.ResourceData) {
				conn := &nextgen.ConnectorInfo{
					GcpCloudCost: &nextgen.GcpCloudCostConnectorDto{
						FeaturesEnabled:   []string{"BILLING"},
						ProjectId:         "p",
						ServiceAccountEmail: "sa@test",
						BillingExportSpec: nil,
					},
				}
				_ = readConnectorGCPCloudCost(d, conn)
			},
		},
		{
			name: "aws_cc: BILLING enabled, curAttributes null",
			resourceSch: func() map[string]*schema.Schema {
				return ResourceConnectorAwsCC().Schema
			}(),
			run: func(d *schema.ResourceData) {
				conn := &nextgen.ConnectorInfo{
					AwsCC: &nextgen.CeAwsConnector{
						AwsAccountId:    "123456789012",
						FeaturesEnabled: []string{"BILLING"},
						CurAttributes:   nil,
						CrossAccountAccess: &nextgen.CrossAccountAccess{
							CrossAccountRoleArn: "arn:aws:iam::123456789012:role/x",
							ExternalId:          "ext",
						},
					},
				}
				_ = readConnectorAwsCC(d, conn)
			},
		},
		{
			name: "aws_cc: crossAccountAccess null",
			resourceSch: func() map[string]*schema.Schema {
				return ResourceConnectorAwsCC().Schema
			}(),
			run: func(d *schema.ResourceData) {
				conn := &nextgen.ConnectorInfo{
					AwsCC: &nextgen.CeAwsConnector{
						AwsAccountId:       "123456789012",
						FeaturesEnabled:    []string{"VISIBILITY"},
						CurAttributes:      nil,
						CrossAccountAccess: nil,
					},
				}
				_ = readConnectorAwsCC(d, conn)
			},
		},
		{
			name: "azure_cloud_cost: BILLING enabled, billingExportSpec2 null (CCM-32488 scenario)",
			resourceSch: func() map[string]*schema.Schema {
				return ResourceConnectorAzureCloudCost().Schema
			}(),
			run: func(d *schema.ResourceData) {
				conn := &nextgen.ConnectorInfo{
					AzureCloudCost: &nextgen.CeAzureConnector{
						TenantId:       "tenant",
						SubscriptionId: "sub",
						FeaturesEnabled: []string{"BILLING"},
						BillingExportSpec: &nextgen.BillingExportSpec{
							StorageAccountName: "sa",
							ContainerName:      "c",
							ReportName:           "r",
						},
						BillingExportSpec2: nil,
					},
				}
				_ = readConnectorAzureCloudCost(d, conn)
			},
		},
		{
			name: "azure_cloud_cost: BILLING enabled, billingExportSpec null",
			resourceSch: func() map[string]*schema.Schema {
				return ResourceConnectorAzureCloudCost().Schema
			}(),
			run: func(d *schema.ResourceData) {
				conn := &nextgen.ConnectorInfo{
					AzureCloudCost: &nextgen.CeAzureConnector{
						TenantId:          "tenant",
						SubscriptionId:    "sub",
						FeaturesEnabled:   []string{"BILLING"},
						BillingExportSpec: nil,
						BillingExportSpec2: &nextgen.BillingExportSpec{
							StorageAccountName: "sa2",
						},
					},
				}
				_ = readConnectorAzureCloudCost(d, conn)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, tc.resourceSch, map[string]interface{}{})
			panicked, val := capturePanic(func() { tc.run(d) })
			if panicked {
				t.Logf("RESULT: FAIL (panic) — %v", val)
			} else {
				t.Logf("RESULT: PASS (no panic)")
			}
			// Store outcome in test name for -v summary; mark as failure only when we document expected bugs
			if panicked {
				t.Errorf("read panicked (CCM-32488-class bug confirmed): %v", val)
			}
		})
	}
}
