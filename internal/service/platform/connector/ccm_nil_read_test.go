package connector

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestReadConnectorGCPCloudCost_NilBillingExportSpec(t *testing.T) {
	resource := ResourceConnectorGCPCloudCost()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		GcpCloudCost: &nextgen.GcpCloudCostConnectorDto{
			FeaturesEnabled:     []string{"BILLING", "VISIBILITY"},
			ProjectId:           "gcp-project",
			ServiceAccountEmail: "sa@gcp-project.iam.gserviceaccount.com",
		},
	}

	if err := readConnectorGCPCloudCost(d, connector); err != nil {
		t.Fatal(err)
	}
	if got := d.Get("billing_export_spec").([]interface{}); len(got) != 0 {
		t.Fatalf("billing_export_spec = %#v, want empty", got)
	}
	if got := d.Get("gcp_project_id"); got != "gcp-project" {
		t.Fatalf("gcp_project_id = %#v", got)
	}
}

func TestReadConnectorAwsCC_NilCurAttributes(t *testing.T) {
	resource := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		AwsCC: &nextgen.CeAwsConnector{
			AwsAccountId:    "123456789012",
			FeaturesEnabled: []string{"BILLING"},
			CrossAccountAccess: &nextgen.CrossAccountAccess{
				CrossAccountRoleArn: "arn:aws:iam::123456789012:role/S3Access",
				ExternalId:          "harness:999999999999",
			},
		},
	}

	if err := readConnectorAwsCC(d, connector); err != nil {
		t.Fatal(err)
	}
	if got := d.Get("report_name"); got != "" {
		t.Fatalf("report_name = %#v, want empty", got)
	}
	if got := d.Get("cross_account_access").([]interface{}); len(got) != 1 {
		t.Fatalf("cross_account_access = %#v, want one entry", got)
	}
}

func TestReadConnectorAwsCC_NilCrossAccountAccess(t *testing.T) {
	resource := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		AwsCC: &nextgen.CeAwsConnector{
			AwsAccountId:    "123456789012",
			FeaturesEnabled: []string{"VISIBILITY"},
		},
	}

	if err := readConnectorAwsCC(d, connector); err != nil {
		t.Fatal(err)
	}
	if got := d.Get("cross_account_access").([]interface{}); len(got) != 0 {
		t.Fatalf("cross_account_access = %#v, want empty", got)
	}
	if got := d.Get("account_id"); got != "123456789012" {
		t.Fatalf("account_id = %#v", got)
	}
}
