package repository

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func newRepoTestResourceData(t *testing.T, values map[string]interface{}) *schema.ResourceData {
	t.Helper()
	r := ResourceGitopsRepositories()
	d := r.TestResourceData()
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("failed to set %q: %s", k, err)
		}
	}
	return d
}

func baseRepoBlock(overrides map[string]interface{}) []interface{} {
	base := map[string]interface{}{
		"repo":                          "https://github.com/example/repo",
		"username":                      "",
		"password":                      "",
		"ssh_private_key":               "",
		"insecure_ignore_host_key":      false,
		"insecure":                      false,
		"enable_lfs":                    false,
		"tls_client_cert_data":          "",
		"tls_client_cert_key":           "",
		"type_":                         "git",
		"name":                          "",
		"inherited_creds":               false,
		"enable_oci":                    false,
		"github_app_private_key":        "",
		"github_app_id":                 "",
		"github_app_installation_id":    "",
		"github_app_enterprise_base_url": "",
		"proxy":                         "",
		"project":                       "",
		"connection_type":               "HTTPS_ANONYMOUS",
		"password_wo_version":           0,
		"ssh_private_key_wo_version":    0,
		"github_app_private_key_wo_version": 0,
	}
	for k, v := range overrides {
		base[k] = v
	}
	return []interface{}{base}
}

// TestBuildRepo_LegacyPasswordUsed verifies that the legacy password field
// is forwarded to the API object when no _wo field is set.
func TestBuildRepo_LegacyPasswordUsed(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo":       baseRepoBlock(map[string]interface{}{"password": "ghp_legacy", "connection_type": "HTTPS"}),
	})

	repo := buildRepo(d)

	if repo.Password != "ghp_legacy" {
		t.Errorf("expected Password=ghp_legacy from legacy field, got %q", repo.Password)
	}
}

// TestBuildRepo_LegacySSHKeyUsed verifies that the legacy ssh_private_key field is forwarded.
func TestBuildRepo_LegacySSHKeyUsed(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo":       baseRepoBlock(map[string]interface{}{"ssh_private_key": "test-ssh-key-value", "connection_type": "SSH"}),
	})

	repo := buildRepo(d)

	if repo.SshPrivateKey == "" {
		t.Error("expected SshPrivateKey to be set from legacy field")
	}
}

// TestBuildRepo_LegacyTLSCertUsed verifies mTLS legacy fields are forwarded.
func TestBuildRepo_LegacyTLSCertUsed(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo": baseRepoBlock(map[string]interface{}{
			"tls_client_cert_data": "cert-pem-data",
			"tls_client_cert_key":  "key-pem-data",
			"connection_type":      "HTTPS",
		}),
	})

	repo := buildRepo(d)

	if repo.TlsClientCertData != "cert-pem-data" {
		t.Errorf("expected TlsClientCertData=cert-pem-data, got %q", repo.TlsClientCertData)
	}
	if repo.TlsClientCertKey != "key-pem-data" {
		t.Errorf("expected TlsClientCertKey=key-pem-data, got %q", repo.TlsClientCertKey)
	}
}

// TestBuildRepo_LegacyGithubAppFields verifies GitHub app legacy fields are forwarded.
func TestBuildRepo_LegacyGithubAppFields(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo": baseRepoBlock(map[string]interface{}{
			"github_app_private_key":     "app-key-pem",
			"github_app_id":              "app-123",
			"github_app_installation_id": "install-456",
			"connection_type":            "GITHUB",
		}),
	})

	repo := buildRepo(d)

	if repo.GithubAppPrivateKey != "app-key-pem" {
		t.Errorf("expected GithubAppPrivateKey=app-key-pem, got %q", repo.GithubAppPrivateKey)
	}
	if repo.GithubAppID != "app-123" {
		t.Errorf("expected GithubAppID=app-123, got %q", repo.GithubAppID)
	}
	if repo.GithubAppInstallationID != "install-456" {
		t.Errorf("expected GithubAppInstallationID=install-456, got %q", repo.GithubAppInstallationID)
	}
}

// TestRepoSchema_WoFieldsPresent verifies that all write-only fields exist in the schema.
func TestRepoSchema_WoFieldsPresent(t *testing.T) {
	r := ResourceGitopsRepositories()
	repoElem := r.Schema["repo"].Elem.(*schema.Resource)
	woFields := []string{
		"password_wo",
		"password_wo_version",
		"ssh_private_key_wo",
		"ssh_private_key_wo_version",
		"tls_client_cert_data_wo",
		"tls_client_cert_key_wo",
		"github_app_private_key_wo",
		"github_app_private_key_wo_version",
	}
	for _, field := range woFields {
		if _, ok := repoElem.Schema[field]; !ok {
			t.Errorf("expected field %q in repo schema, not found", field)
		}
	}
}

// TestRepoSchema_WoFieldsAreWriteOnly verifies WriteOnly is set on _wo credential fields.
func TestRepoSchema_WoFieldsAreWriteOnly(t *testing.T) {
	r := ResourceGitopsRepositories()
	repoElem := r.Schema["repo"].Elem.(*schema.Resource)
	writeOnlyFields := []string{
		"password_wo",
		"ssh_private_key_wo",
		"tls_client_cert_data_wo",
		"tls_client_cert_key_wo",
		"github_app_private_key_wo",
	}
	for _, field := range writeOnlyFields {
		s, ok := repoElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in repo schema", field)
			continue
		}
		if !s.WriteOnly {
			t.Errorf("expected WriteOnly=true on field %q", field)
		}
		if !s.Sensitive {
			t.Errorf("expected Sensitive=true on field %q", field)
		}
	}
}

// TestRepoSchema_WoVersionFieldsAreInt verifies _version fields are TypeInt.
func TestRepoSchema_WoVersionFieldsAreInt(t *testing.T) {
	r := ResourceGitopsRepositories()
	repoElem := r.Schema["repo"].Elem.(*schema.Resource)
	versionFields := []string{
		"password_wo_version",
		"ssh_private_key_wo_version",
		"github_app_private_key_wo_version",
	}
	for _, field := range versionFields {
		s, ok := repoElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in repo schema", field)
			continue
		}
		if s.Type != schema.TypeInt {
			t.Errorf("expected TypeInt for field %q, got %v", field, s.Type)
		}
	}
}

// TestRepoSchema_WoConflictsWithLegacy verifies ConflictsWith is wired between legacy and _wo fields.
func TestRepoSchema_WoConflictsWithLegacy(t *testing.T) {
	r := ResourceGitopsRepositories()
	repoElem := r.Schema["repo"].Elem.(*schema.Resource)

	pairs := map[string]string{
		"password_wo":            "repo.0.password",
		"ssh_private_key_wo":     "repo.0.ssh_private_key",
		"tls_client_cert_data_wo": "repo.0.tls_client_cert_data",
		"tls_client_cert_key_wo": "repo.0.tls_client_cert_key",
		"github_app_private_key_wo": "repo.0.github_app_private_key",
	}
	for woField, legacyPath := range pairs {
		s, ok := repoElem.Schema[woField]
		if !ok {
			t.Errorf("field %q not found in repo schema", woField)
			continue
		}
		found := false
		for _, c := range s.ConflictsWith {
			if c == legacyPath {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected %q in ConflictsWith of %q, got %v", legacyPath, woField, s.ConflictsWith)
		}
	}
}

// TestRepoSchema_LegacyFieldsStillComputed verifies old fields retain Computed=true (backward compat).
func TestRepoSchema_LegacyFieldsStillComputed(t *testing.T) {
	r := ResourceGitopsRepositories()
	repoElem := r.Schema["repo"].Elem.(*schema.Resource)
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
		s, ok := repoElem.Schema[field]
		if !ok {
			t.Errorf("legacy field %q not found in repo schema", field)
			continue
		}
		if !s.Computed {
			t.Errorf("expected Computed=true on legacy field %q (backward compat)", field)
		}
	}
}

// fakeServicev1Repository builds a minimal API response for unit tests.
func fakeServicev1Repository(identifier string, repo nextgen.RepositoriesRepository) nextgen.Servicev1Repository {
	return nextgen.Servicev1Repository{
		Identifier:        identifier,
		AccountIdentifier: "acc",
		AgentIdentifier:   "agent",
		Repository:        &repo,
	}
}

// TestSetRepositoryDetails_AllWoVersionsPreserved is the regression test for the bug where
// setRepositoryDetails only injected 3 of 5 _wo_version fields into the map before d.Set,
// causing tls_client_cert_data_wo_version and tls_client_cert_key_wo_version to be zeroed.
func TestSetRepositoryDetails_AllWoVersionsPreserved(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo": baseRepoBlock(map[string]interface{}{
			"connection_type":                  "HTTPS",
			"password_wo_version":              1,
			"ssh_private_key_wo_version":       2,
			"tls_client_cert_data_wo_version":  3,
			"tls_client_cert_key_wo_version":   4,
			"github_app_private_key_wo_version": 5,
		}),
	})

	resp := fakeServicev1Repository("test-repo", nextgen.RepositoriesRepository{
		Repo:           "https://github.com/example/repo",
		ConnectionType: "HTTPS",
	})
	setRepositoryDetails(d, &resp)

	cases := map[string]int{
		"repo.0.password_wo_version":               1,
		"repo.0.ssh_private_key_wo_version":        2,
		"repo.0.tls_client_cert_data_wo_version":   3,
		"repo.0.tls_client_cert_key_wo_version":    4,
		"repo.0.github_app_private_key_wo_version": 5,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after setRepositoryDetails, got %v (ok=%v)", field, want, got, ok)
		}
	}
}

// TestPreserveWoVersions_AllFieldsRoundtrip verifies that preserveWoVersions re-writes all 5
// _wo_version fields, including the two tls fields that were missing before the fix.
func TestPreserveWoVersions_AllFieldsRoundtrip(t *testing.T) {
	d := newRepoTestResourceData(t, map[string]interface{}{
		"identifier": "test-repo",
		"agent_id":   "test-agent",
		"repo": baseRepoBlock(map[string]interface{}{
			"connection_type":                  "HTTPS",
			"password_wo_version":              6,
			"ssh_private_key_wo_version":       7,
			"tls_client_cert_data_wo_version":  8,
			"tls_client_cert_key_wo_version":   9,
			"github_app_private_key_wo_version": 10,
		}),
	})

	preserveWoVersions(d)

	cases := map[string]int{
		"repo.0.password_wo_version":               6,
		"repo.0.ssh_private_key_wo_version":        7,
		"repo.0.tls_client_cert_data_wo_version":   8,
		"repo.0.tls_client_cert_key_wo_version":    9,
		"repo.0.github_app_private_key_wo_version": 10,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after preserveWoVersions, got %v (ok=%v)", field, want, got, ok)
		}
	}
}
