package repository_credentials

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
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
// TestRepoCredSchema_BlockIsNotComputed verifies the creds block does NOT have Computed=true.
// The SDK prohibits WriteOnly attributes inside a Computed=true block, so Computed was removed
// from the creds block when write-only fields were introduced (CDS-123457).
func TestRepoCredSchema_BlockIsNotComputed(t *testing.T) {
	r := ResourceGitopsRepoCred()
	if r.Schema["creds"].Computed {
		t.Error("creds block must not have Computed=true — WriteOnly fields are incompatible with Computed blocks")
	}
}

// fakeRepoCred builds a minimal Servicev1RepositoryCredentials API response with the given cred fields.
func fakeRepoCred(identifier string, creds nextgen.HrepocredsRepoCreds) nextgen.Servicev1RepositoryCredentials {
	return nextgen.Servicev1RepositoryCredentials{
		Identifier:        identifier,
		AccountIdentifier: "acc",
		AgentIdentifier:   "agent",
		RepoCreds:         &creds,
	}
}

// TestSetGitopsRepositoriesCredential_WoVersionsPreservedAfterSet is the regression test for
// PIPE-34920: setGitopsRepositoriesCredential called d.Set("creds", list) which zeroed all
// _wo_version integers because the API never returns write-only values.
func TestSetGitopsRepositoriesCredential_WoVersionsPreservedAfterSet(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds": baseCredsBlock(map[string]interface{}{
			"ssh_private_key_wo_version":        2,
			"password_wo_version":               1,
			"tls_client_cert_data_wo_version":   3,
			"tls_client_cert_key_wo_version":    4,
			"github_app_private_key_wo_version": 5,
		}),
	})

	resp := fakeRepoCred("test-cred", nextgen.HrepocredsRepoCreds{
		Url:  "https://github.com/example",
		Type_: "git",
	})
	setGitopsRepositoriesCredential(d, &resp)

	cases := map[string]int{
		"creds.0.password_wo_version":               1,
		"creds.0.ssh_private_key_wo_version":        2,
		"creds.0.tls_client_cert_data_wo_version":   3,
		"creds.0.tls_client_cert_key_wo_version":    4,
		"creds.0.github_app_private_key_wo_version": 5,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after setGitopsRepositoriesCredential, got %v (ok=%v)", field, want, got, ok)
		}
	}
}

// TestPreserveRepoCredWoVersions_AllFieldsRoundtrip verifies that preserveRepoCredWoVersions
// re-writes each _wo_version field back to state (belt-and-suspenders after the in-map fix).
func TestPreserveRepoCredWoVersions_AllFieldsRoundtrip(t *testing.T) {
	d := newRepoCredTestResourceData(t, map[string]interface{}{
		"identifier": "test-cred",
		"agent_id":   "test-agent",
		"creds": baseCredsBlock(map[string]interface{}{
			"password_wo_version":               7,
			"ssh_private_key_wo_version":        8,
			"tls_client_cert_data_wo_version":   9,
			"tls_client_cert_key_wo_version":    10,
			"github_app_private_key_wo_version": 11,
		}),
	})

	preserveRepoCredWoVersions(d)

	cases := map[string]int{
		"creds.0.password_wo_version":               7,
		"creds.0.ssh_private_key_wo_version":        8,
		"creds.0.tls_client_cert_data_wo_version":   9,
		"creds.0.tls_client_cert_key_wo_version":    10,
		"creds.0.github_app_private_key_wo_version": 11,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after preserveRepoCredWoVersions, got %v (ok=%v)", field, want, got, ok)
		}
	}
}
