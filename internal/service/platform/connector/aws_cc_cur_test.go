package connector

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildConnectorAwsCC_CUR2(t *testing.T) {
	resource := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"account_id":  "123456789012",
		"report_name": "test_report",
		"report_type": "CUR2.0",
		"region":      "us-east-1",
		"s3_bucket":   "s3bucket",
		"s3_prefix":   "cur",
		"features_enabled": []interface{}{
			"BILLING",
		},
		"cross_account_access": []interface{}{
			map[string]interface{}{
				"role_arn":    "arn:aws:iam::123456789012:role/S3Access",
				"external_id": "harness:999999999999",
			},
		},
	})

	connector := buildConnectorAwsCC(d)
	attrs := connector.AwsCC.CurAttributes
	if attrs.ReportType != "CUR2.0" || attrs.Region != "us-east-1" || attrs.S3Prefix != "cur" {
		t.Fatalf("cur attributes = %+v", attrs)
	}

	raw, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"reportType":"CUR2.0"`, `"region":"us-east-1"`, `"s3Prefix":"cur"`} {
		if !strings.Contains(string(raw), field) {
			t.Fatalf("payload %s missing %s", raw, field)
		}
	}
}

func TestBuildConnectorAwsCC_LegacyCUROmitsReportType(t *testing.T) {
	resource := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"account_id":  "123456789012",
		"report_name": "test_report",
		"s3_bucket":   "s3bucket",
		"features_enabled": []interface{}{
			"BILLING",
		},
		"cross_account_access": []interface{}{
			map[string]interface{}{
				"role_arn":    "arn:aws:iam::123456789012:role/S3Access",
				"external_id": "harness:999999999999",
			},
		},
	})

	raw, err := json.Marshal(buildConnectorAwsCC(d).AwsCC.CurAttributes)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"reportType", "region", "s3Prefix"} {
		if strings.Contains(string(raw), field) {
			t.Fatalf("legacy CUR payload should omit %s, got %s", field, raw)
		}
	}
}

func TestBuildConnectorAwsCC_PreservesServerCURWhenConfigOmitsFields(t *testing.T) {
	resourceSchema := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resourceSchema.Schema, map[string]interface{}{
		"account_id":  "123456789012",
		"report_name": "test_report",
		"s3_bucket":   "s3bucket",
		"features_enabled": []interface{}{
			"BILLING",
		},
		"cross_account_access": []interface{}{
			map[string]interface{}{
				"role_arn":    "arn:aws:iam::123456789012:role/S3Access",
				"external_id": "harness:999999999999",
			},
		},
	})

	server := &nextgen.ConnectorInfo{
		AwsCC: &nextgen.CeAwsConnector{
			AwsAccountId:    "123456789012",
			FeaturesEnabled: []string{"BILLING"},
			CurAttributes: &nextgen.AwsCurAttributes{
				ReportName:   "test_report",
				S3BucketName: "s3bucket",
				Region:       "us-east-1",
				S3Prefix:     "cur",
				ReportType:   "CUR2.0",
			},
			CrossAccountAccess: &nextgen.CrossAccountAccess{
				CrossAccountRoleArn: "arn:aws:iam::123456789012:role/S3Access",
				ExternalId:          "harness:999999999999",
			},
		},
	}
	if err := readConnectorAwsCC(d, server); err != nil {
		t.Fatal(err)
	}

	attrs := buildConnectorAwsCC(d).AwsCC.CurAttributes
	if attrs.ReportType != "CUR2.0" || attrs.Region != "us-east-1" || attrs.S3Prefix != "cur" {
		t.Fatalf("update payload dropped server CUR settings: %+v", attrs)
	}
}

func TestReadConnectorAwsCC_CUR2(t *testing.T) {
	resource := ResourceConnectorAwsCC()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		AwsCC: &nextgen.CeAwsConnector{
			AwsAccountId:    "123456789012",
			FeaturesEnabled: []string{"BILLING"},
			CurAttributes: &nextgen.AwsCurAttributes{
				ReportName:   "test_report",
				S3BucketName: "s3bucket",
				Region:       "us-east-1",
				S3Prefix:     "cur",
				ReportType:   "CUR2.0",
			},
			CrossAccountAccess: &nextgen.CrossAccountAccess{
				CrossAccountRoleArn: "arn:aws:iam::123456789012:role/S3Access",
				ExternalId:          "harness:999999999999",
			},
		},
	}

	if err := readConnectorAwsCC(d, connector); err != nil {
		t.Fatal(err)
	}
	if got := d.Get("report_type"); got != "CUR2.0" {
		t.Fatalf("report_type = %#v", got)
	}
	if got := d.Get("region"); got != "us-east-1" {
		t.Fatalf("region = %#v", got)
	}
	if got := d.Get("s3_prefix"); got != "cur" {
		t.Fatalf("s3_prefix = %#v", got)
	}
}
