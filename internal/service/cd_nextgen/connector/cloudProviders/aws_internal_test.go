package cloudProviders

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildConnectorAws_OidcSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAws()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"oidc_authentication": []interface{}{
			map[string]interface{}{
				"iam_role_arn": "arn:aws:iam::123456789012:role/example",
				"oidc_session_tag_keys": []interface{}{
					"account_id",
					"organization_id",
				},
			},
		},
	})

	connector := buildConnectorAws(d)

	if connector.Aws.Credential.OidcConfig == nil {
		t.Fatal("expected oidc config to be populated")
	}

	if len(connector.Aws.Credential.OidcConfig.OidcSessionTagKeys) != 2 {
		t.Fatalf("expected 2 oidc session tag keys, got %d", len(connector.Aws.Credential.OidcConfig.OidcSessionTagKeys))
	}

	if connector.Aws.Credential.OidcConfig.OidcSessionTagKeys[0] != "account_id" {
		t.Fatalf("expected first tag key account_id, got %q", connector.Aws.Credential.OidcConfig.OidcSessionTagKeys[0])
	}
}

func TestReadConnectorAws_OidcSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAws()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		Aws: &nextgen.AwsConnector{
			Credential: &nextgen.AwsCredential{
				Type_: nextgen.AwsAuthTypes.OidcAuthentication,
				OidcConfig: &nextgen.AwsOidcConfigSpec{
					IamRoleArn: "arn:aws:iam::123456789012:role/example",
					OidcSessionTagKeys: []string{
						"account_id",
						"project_id",
					},
				},
			},
		},
	}

	if err := readConnectorAws(d, connector); err != nil {
		t.Fatalf("readConnectorAws failed: %v", err)
	}

	tagKeys := d.Get("oidc_authentication.0.oidc_session_tag_keys").([]interface{})
	if len(tagKeys) != 2 {
		t.Fatalf("expected 2 oidc session tag keys in state, got %d", len(tagKeys))
	}

	if tagKeys[0].(string) != "account_id" {
		t.Fatalf("expected first tag key account_id, got %q", tagKeys[0].(string))
	}
}

func TestBuildConnectorAws_OidcWithoutSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAws()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"oidc_authentication": []interface{}{
			map[string]interface{}{
				"iam_role_arn": "arn:aws:iam::123456789012:role/example",
				"region":       "me-central-1",
			},
		},
	})

	connector := buildConnectorAws(d)

	if connector.Aws.Credential.OidcConfig == nil {
		t.Fatal("expected oidc config to be populated")
	}

	if len(connector.Aws.Credential.OidcConfig.OidcSessionTagKeys) != 0 {
		t.Fatalf("expected no oidc session tag keys, got %v", connector.Aws.Credential.OidcConfig.OidcSessionTagKeys)
	}
}

func TestReadConnectorAws_OidcWithoutSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAws()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		Aws: &nextgen.AwsConnector{
			Credential: &nextgen.AwsCredential{
				Type_: nextgen.AwsAuthTypes.OidcAuthentication,
				Region: "me-central-1",
				OidcConfig: &nextgen.AwsOidcConfigSpec{
					IamRoleArn: "arn:aws:iam::123456789012:role/example",
				},
			},
		},
	}

	if err := readConnectorAws(d, connector); err != nil {
		t.Fatalf("readConnectorAws failed: %v", err)
	}

	tagKeys, ok := d.GetOk("oidc_authentication.0.oidc_session_tag_keys")
	if ok {
		t.Fatalf("expected oidc session tag keys to be unset, got %v", tagKeys)
	}
}
