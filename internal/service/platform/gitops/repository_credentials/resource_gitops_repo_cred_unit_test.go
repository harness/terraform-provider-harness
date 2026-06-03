package repository_credentials

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func newRepoCredTestResourceData(t *testing.T, values map[string]interface{}) *schema.ResourceData {
	t.Helper()
	r := ResourceGitopsRepoCred()
	d := r.TestResourceData()
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("failed to set %q: %s", k, err)
		}
	}
	return d
}

func baseCredsBlock(overrides map[string]interface{}) []interface{} {
	base := map[string]interface{}{
		"url":                            "https://github.com/example",
		"username":                       "",
		"password":                       "",
		"ssh_private_key":                "",
		"tls_client_cert_data":           "",
		"tls_client_cert_key":            "",
		"github_app_private_key":         "",
		"github_app_id":                  "",
		"github_app_installation_id":     "",
		"github_app_enterprise_base_url": "",
		"enable_oci":                     false,
		"type":                           "git",
	}
	for k, v := range overrides {
		base[k] = v
	}
	return []interface{}{base}
}

// TestBuildRepoCred_PasswordForwarded verifies that the password field is forwarded to the API object.
func TestBuildRepoCred_PasswordForwarded(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds":      baseCredsBlock(map[string]interface{}{"password": "ghp_secret", "username": "user"}),
	})

	cred := buildRepoCred(d)

	if cred.Password != "ghp_secret" {
		t.Errorf("expected Password=ghp_secret, got %q", cred.Password)
	}
	if cred.Username != "user" {
		t.Errorf("expected Username=user, got %q", cred.Username)
	}
}

// TestBuildRepoCred_SSHKeyForwarded verifies that ssh_private_key is forwarded to the API object.
func TestBuildRepoCred_SSHKeyForwarded(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds":      baseCredsBlock(map[string]interface{}{"ssh_private_key": "test-ssh-key-value"}),
	})

	cred := buildRepoCred(d)

	if cred.SshPrivateKey == "" {
		t.Error("expected SshPrivateKey to be set")
	}
}

// TestBuildRepoCred_TLSCertForwarded verifies that mTLS cert fields are forwarded.
func TestBuildRepoCred_TLSCertForwarded(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds": baseCredsBlock(map[string]interface{}{
			"tls_client_cert_data": "cert-pem",
			"tls_client_cert_key":  "key-pem",
		}),
	})

	cred := buildRepoCred(d)

	if cred.TlsClientCertData != "cert-pem" {
		t.Errorf("expected TlsClientCertData=cert-pem, got %q", cred.TlsClientCertData)
	}
	if cred.TlsClientCertKey != "key-pem" {
		t.Errorf("expected TlsClientCertKey=key-pem, got %q", cred.TlsClientCertKey)
	}
}

// TestBuildRepoCred_GithubAppForwarded verifies GitHub app fields are forwarded to the API object.
func TestBuildRepoCred_GithubAppForwarded(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds": baseCredsBlock(map[string]interface{}{
			"github_app_private_key":     "app-key-pem",
			"github_app_id":              "app-123",
			"github_app_installation_id": "install-456",
		}),
	})

	cred := buildRepoCred(d)

	if cred.GithubAppPrivateKey != "app-key-pem" {
		t.Errorf("expected GithubAppPrivateKey=app-key-pem, got %q", cred.GithubAppPrivateKey)
	}
	if cred.GithubAppID != "app-123" {
		t.Errorf("expected GithubAppID=app-123, got %q", cred.GithubAppID)
	}
	if cred.GithubAppInstallationID != "install-456" {
		t.Errorf("expected GithubAppInstallationID=install-456, got %q", cred.GithubAppInstallationID)
	}
}

// TestBuildRepoCred_URLAndTypeForwarded verifies URL and type fields are forwarded.
func TestBuildRepoCred_URLAndTypeForwarded(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds": baseCredsBlock(map[string]interface{}{
			"url":  "https://github.com/myorg",
			"type": "helm",
		}),
	})

	cred := buildRepoCred(d)

	if cred.Url != "https://github.com/myorg" {
		t.Errorf("expected Url=https://github.com/myorg, got %q", cred.Url)
	}
	if cred.Type_ != "helm" {
		t.Errorf("expected Type_=helm, got %q", cred.Type_)
	}
}

// TestRepoCredSchema_SensitiveFieldsAreSensitive verifies Sensitive=true on credential fields.
func TestRepoCredSchema_SensitiveFieldsAreSensitive(t *testing.T) {
	r := ResourceGitopsRepoCred()
	credsElem := r.Schema["creds"].Elem.(*schema.Resource)
	sensitiveFields := []string{
		"password",
		"ssh_private_key",
		"tls_client_cert_data",
		"tls_client_cert_key",
		"github_app_private_key",
		"github_app_id",
		"github_app_installation_id",
	}
	for _, field := range sensitiveFields {
		s, ok := credsElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in creds schema", field)
			continue
		}
		if !s.Sensitive {
			t.Errorf("expected Sensitive=true on field %q", field)
		}
	}
}

// TestRepoCredSchema_CredentialFieldsAreComputed verifies Computed=true on credential fields
// so that existing configs aren't broken.
func TestRepoCredSchema_CredentialFieldsAreComputed(t *testing.T) {
	r := ResourceGitopsRepoCred()
	credsElem := r.Schema["creds"].Elem.(*schema.Resource)
	computedFields := []string{
		"password",
		"ssh_private_key",
		"tls_client_cert_data",
		"tls_client_cert_key",
		"github_app_private_key",
		"github_app_id",
		"github_app_installation_id",
	}
	for _, field := range computedFields {
		s, ok := credsElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in creds schema", field)
			continue
		}
		if !s.Computed {
			t.Errorf("expected Computed=true on field %q", field)
		}
	}
}

// TestRepoCredSchema_BlockIsComputed verifies the creds block-level Computed=true is preserved.
// NOTE: This is why write-only fields cannot be added to this resource — the SDK prohibits
// WriteOnly attributes inside a Computed=true block. Tracked separately.
func TestRepoCredSchema_BlockIsComputed(t *testing.T) {
	r := ResourceGitopsRepoCred()
	if !r.Schema["creds"].Computed {
		t.Error("expected creds block to have Computed=true")
	}
}
