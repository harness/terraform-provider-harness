package secretManagers

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildConnectorAwsSM_OidcSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAwsSM()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"credentials": []interface{}{
			map[string]interface{}{
				"oidc_authentication": []interface{}{
					map[string]interface{}{
						"iam_role_arn": "arn:aws:iam::123456789012:role/example",
						"oidc_session_tag_keys": []interface{}{
							"account_id",
							"organization_id",
						},
					},
				},
			},
		},
	})

	connector := buildConnectorAwsSM(d)

	if connector.AwsSecretManager.Credential.OidcConfig == nil {
		t.Fatal("expected oidc config to be populated")
	}

	if len(connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys) != 2 {
		t.Fatalf("expected 2 oidc session tag keys, got %d", len(connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys))
	}

	if connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys[0] != "account_id" {
		t.Fatalf("expected first tag key account_id, got %q", connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys[0])
	}
}

func TestReadConnectorAwsSM_OidcSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAwsSM()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		AwsSecretManager: &nextgen.AwsSecretManager{
			Credential: &nextgen.AwsSecretManagerCredential{
				Type_: nextgen.AwsSecretManagerAuthTypes.OidcAuthentication,
				OidcConfig: &nextgen.AwsSmCredentialSpecOidcConfig{
					IamRoleArn: "arn:aws:iam::123456789012:role/example",
					OidcSessionTagKeys: []string{
						"account_id",
						"project_id",
					},
				},
			},
		},
	}

	if err := readConnectorAwsSM(d, connector); err != nil {
		t.Fatalf("readConnectorAwsSM failed: %v", err)
	}

	tagKeys := d.Get("credentials.0.oidc_authentication.0.oidc_session_tag_keys").([]interface{})
	if len(tagKeys) != 2 {
		t.Fatalf("expected 2 oidc session tag keys in state, got %d", len(tagKeys))
	}

	if tagKeys[0].(string) != "account_id" {
		t.Fatalf("expected first tag key account_id, got %q", tagKeys[0].(string))
	}
}

func TestBuildConnectorAwsSM_OidcWithoutSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAwsSM()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"credentials": []interface{}{
			map[string]interface{}{
				"oidc_authentication": []interface{}{
					map[string]interface{}{
						"iam_role_arn": "arn:aws:iam::123456789012:role/example",
					},
				},
			},
		},
	})

	connector := buildConnectorAwsSM(d)

	if connector.AwsSecretManager.Credential.OidcConfig == nil {
		t.Fatal("expected oidc config to be populated")
	}

	if len(connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys) != 0 {
		t.Fatalf("expected no oidc session tag keys, got %v", connector.AwsSecretManager.Credential.OidcConfig.OidcSessionTagKeys)
	}
}

func TestReadConnectorAwsSM_OidcWithoutSessionTagKeys(t *testing.T) {
	resource := ResourceConnectorAwsSM()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

	connector := &nextgen.ConnectorInfo{
		AwsSecretManager: &nextgen.AwsSecretManager{
			Credential: &nextgen.AwsSecretManagerCredential{
				Type_: nextgen.AwsSecretManagerAuthTypes.OidcAuthentication,
				OidcConfig: &nextgen.AwsSmCredentialSpecOidcConfig{
					IamRoleArn: "arn:aws:iam::123456789012:role/example",
				},
			},
		},
	}

	if err := readConnectorAwsSM(d, connector); err != nil {
		t.Fatalf("readConnectorAwsSM failed: %v", err)
	}

	tagKeys, ok := d.GetOk("credentials.0.oidc_authentication.0.oidc_session_tag_keys")
	if ok {
		t.Fatalf("expected oidc session tag keys to be unset, got %v", tagKeys)
	}
}
