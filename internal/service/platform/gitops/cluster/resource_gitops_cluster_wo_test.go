package cluster

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func baseClusterConfigBlock(overrides map[string]interface{}) []interface{} {
	base := map[string]interface{}{
		"username":                "",
		"password":                "",
		"bearer_token":            "",
		"role_a_r_n":              "",
		"aws_cluster_name":        "",
		"cluster_connection_type": "",
		"disable_compression":     false,
		"proxy_url":               "",
		"password_wo_version":     0,
		"bearer_token_wo_version": 0,
		"tls_client_config": []interface{}{
			map[string]interface{}{
				"insecure":    false,
				"server_name": "",
				"cert_data":   "",
				"key_data":    "",
				"ca_data":     "",
			},
		},
		"exec_provider_config": []interface{}{},
	}
	for k, v := range overrides {
		base[k] = v
	}
	return []interface{}{base}
}

func baseClusterRequestBlock(configOverrides map[string]interface{}) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"upsert":         false,
			"updated_fields": []interface{}{},
			"tags":           schema.NewSet(schema.HashString, []interface{}{}),
			"cluster": []interface{}{
				map[string]interface{}{
					"server":               "https://k8s.example.com",
					"name":                 "my-cluster",
					"config":               baseClusterConfigBlock(configOverrides),
					"namespaces":           []interface{}{},
					"shard":                "",
					"cluster_resources":    false,
					"project":              "",
					"labels":               map[string]interface{}{},
					"annotations":          map[string]interface{}{},
					"refresh_requested_at": []interface{}{},
					"info":                 []interface{}{},
				},
			},
		},
	}
}

// TestBuildClusterDetails_LegacyBearerTokenUsed verifies the legacy bearer_token field
// is forwarded to the API object when no _wo field is present.
func TestBuildClusterDetails_LegacyBearerTokenUsed(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request":    baseClusterRequestBlock(map[string]interface{}{"bearer_token": "legacy-token", "cluster_connection_type": "SERVICE_ACCOUNT"}),
	})

	cluster := buildClusterDetails(d)

	if cluster.Config == nil {
		t.Fatal("expected Config to be non-nil")
	}
	if cluster.Config.BearerToken != "legacy-token" {
		t.Errorf("expected BearerToken=legacy-token from legacy field, got %q", cluster.Config.BearerToken)
	}
}

// TestBuildClusterDetails_LegacyPasswordUsed verifies the legacy password field is forwarded.
func TestBuildClusterDetails_LegacyPasswordUsed(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request":    baseClusterRequestBlock(map[string]interface{}{"password": "legacy-pass", "username": "admin"}),
	})

	cluster := buildClusterDetails(d)

	if cluster.Config == nil {
		t.Fatal("expected Config to be non-nil")
	}
	if cluster.Config.Password != "legacy-pass" {
		t.Errorf("expected Password=legacy-pass from legacy field, got %q", cluster.Config.Password)
	}
}

// TestBuildClusterDetails_LegacyTLSCertUsed verifies mTLS cert fields are forwarded.
func TestBuildClusterDetails_LegacyTLSCertUsed(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request": []interface{}{
			map[string]interface{}{
				"upsert":         false,
				"updated_fields": []interface{}{},
				"tags":           schema.NewSet(schema.HashString, []interface{}{}),
				"cluster": []interface{}{
					map[string]interface{}{
						"server": "https://k8s.example.com",
						"name":   "my-cluster",
						"config": []interface{}{
							map[string]interface{}{
								"username":                "",
								"password":                "",
								"bearer_token":            "",
								"role_a_r_n":              "",
								"aws_cluster_name":        "",
								"cluster_connection_type": "",
								"disable_compression":     false,
								"proxy_url":               "",
								"password_wo_version":     0,
								"bearer_token_wo_version": 0,
								"tls_client_config": []interface{}{
									map[string]interface{}{
										"insecure":    false,
										"server_name": "",
										"cert_data":   "cert-pem-data",
										"key_data":    "key-pem-data",
										"ca_data":     "",
									},
								},
								"exec_provider_config": []interface{}{},
							},
						},
						"namespaces":           []interface{}{},
						"shard":                "",
						"cluster_resources":    false,
						"project":              "",
						"labels":               map[string]interface{}{},
						"annotations":          map[string]interface{}{},
						"refresh_requested_at": []interface{}{},
						"info":                 []interface{}{},
					},
				},
			},
		},
	})

	cluster := buildClusterDetails(d)

	if cluster.Config == nil {
		t.Fatal("expected Config to be non-nil")
	}
	if cluster.Config.TlsClientConfig == nil {
		t.Fatal("expected TlsClientConfig to be non-nil")
	}
	if cluster.Config.TlsClientConfig.CertData != "cert-pem-data" {
		t.Errorf("expected CertData=cert-pem-data, got %q", cluster.Config.TlsClientConfig.CertData)
	}
	if cluster.Config.TlsClientConfig.KeyData != "key-pem-data" {
		t.Errorf("expected KeyData=key-pem-data, got %q", cluster.Config.TlsClientConfig.KeyData)
	}
}

// TestClusterSchema_WoFieldsPresent verifies all write-only fields exist in the config schema.
func TestClusterSchema_WoFieldsPresent(t *testing.T) {
	r := ResourceGitopsCluster()
	configElem := r.Schema["request"].Elem.(*schema.Resource).
		Schema["cluster"].Elem.(*schema.Resource).
		Schema["config"].Elem.(*schema.Resource)

	woFields := []string{
		"password_wo",
		"password_wo_version",
		"bearer_token_wo",
		"bearer_token_wo_version",
	}
	for _, field := range woFields {
		if _, ok := configElem.Schema[field]; !ok {
			t.Errorf("expected field %q in config schema, not found", field)
		}
	}

	tlsElem := configElem.Schema["tls_client_config"].Elem.(*schema.Resource)
	tlsWoFields := []string{"cert_data_wo", "cert_data_wo_version", "key_data_wo", "key_data_wo_version", "ca_data_wo", "ca_data_wo_version"}
	for _, field := range tlsWoFields {
		if _, ok := tlsElem.Schema[field]; !ok {
			t.Errorf("expected field %q in tls_client_config schema, not found", field)
		}
	}
}

// TestClusterSchema_WoFieldsAreWriteOnly verifies WriteOnly and Sensitive are set on _wo fields.
func TestClusterSchema_WoFieldsAreWriteOnly(t *testing.T) {
	r := ResourceGitopsCluster()
	configElem := r.Schema["request"].Elem.(*schema.Resource).
		Schema["cluster"].Elem.(*schema.Resource).
		Schema["config"].Elem.(*schema.Resource)

	writeOnlyFields := []string{"password_wo", "bearer_token_wo"}
	for _, field := range writeOnlyFields {
		s, ok := configElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in config schema", field)
			continue
		}
		if !s.WriteOnly {
			t.Errorf("expected WriteOnly=true on field %q", field)
		}
		if !s.Sensitive {
			t.Errorf("expected Sensitive=true on field %q", field)
		}
	}

	tlsElem := configElem.Schema["tls_client_config"].Elem.(*schema.Resource)
	for _, field := range []string{"cert_data_wo", "key_data_wo", "ca_data_wo"} {
		s, ok := tlsElem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in tls_client_config schema", field)
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

// TestClusterSchema_WoConflictsWithLegacy verifies ConflictsWith between legacy and _wo fields.
func TestClusterSchema_WoConflictsWithLegacy(t *testing.T) {
	r := ResourceGitopsCluster()
	configElem := r.Schema["request"].Elem.(*schema.Resource).
		Schema["cluster"].Elem.(*schema.Resource).
		Schema["config"].Elem.(*schema.Resource)

	pairs := map[string]string{
		"password_wo":     "request.0.cluster.0.config.0.password",
		"bearer_token_wo": "request.0.cluster.0.config.0.bearer_token",
	}
	for woField, legacyPath := range pairs {
		s, ok := configElem.Schema[woField]
		if !ok {
			t.Errorf("field %q not found in config schema", woField)
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

// TestClusterSchema_BearerTokenStillComputed verifies bearer_token retains Computed=true.
func TestClusterSchema_BearerTokenStillComputed(t *testing.T) {
	r := ResourceGitopsCluster()
	configElem := r.Schema["request"].Elem.(*schema.Resource).
		Schema["cluster"].Elem.(*schema.Resource).
		Schema["config"].Elem.(*schema.Resource)

	s, ok := configElem.Schema["bearer_token"]
	if !ok {
		t.Fatal("bearer_token field not found in config schema")
	}
	if !s.Computed {
		t.Error("expected Computed=true on bearer_token for backward compatibility")
	}
}

// fakeServicev1Cluster builds a minimal API response for unit tests.
func fakeServicev1Cluster(identifier string) nextgen.Servicev1Cluster {
	return nextgen.Servicev1Cluster{
		Identifier:        identifier,
		AccountIdentifier: "acc",
		AgentIdentifier:   "agent",
		Cluster: &nextgen.ClustersCluster{
			Server: "https://k8s.example.com",
			Name:   "my-cluster",
			Config: &nextgen.ClustersClusterConfig{
				ClusterConnectionType: "SERVICE_ACCOUNT",
				TlsClientConfig:       &nextgen.ClustersTlsClientConfig{},
			},
		},
	}
}

// TestSetClusterDetails_AllWoVersionsPreserved is the regression test for the bug where
// setClusterDetails called d.Set("request", list) which zeroed _wo_version integers.
func TestSetClusterDetails_AllWoVersionsPreserved(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request": baseClusterRequestBlock(map[string]interface{}{
			"cluster_connection_type": "SERVICE_ACCOUNT",
			"password_wo_version":     1,
			"bearer_token_wo_version": 2,
			"tls_client_config": []interface{}{
				map[string]interface{}{
					"insecure":             false,
					"server_name":          "",
					"cert_data":            "",
					"key_data":             "",
					"ca_data":              "",
					"cert_data_wo_version": 3,
					"key_data_wo_version":  4,
					"ca_data_wo_version":   5,
				},
			},
		}),
	})

	resp := fakeServicev1Cluster("test-cluster")
	setClusterDetails(d, &resp)

	cases := map[string]int{
		"request.0.cluster.0.config.0.password_wo_version":                      1,
		"request.0.cluster.0.config.0.bearer_token_wo_version":                  2,
		"request.0.cluster.0.config.0.tls_client_config.0.cert_data_wo_version": 3,
		"request.0.cluster.0.config.0.tls_client_config.0.key_data_wo_version":  4,
		"request.0.cluster.0.config.0.tls_client_config.0.ca_data_wo_version":   5,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after setClusterDetails, got %v (ok=%v)", field, want, got, ok)
		}
	}
}

// TestSetClusterDetails_BearerTokenMaskedWhenWoVersion ensures bearer_token is not
// rehydrated into state when bearer_token_wo(_version) mode is active.
func TestSetClusterDetails_BearerTokenMaskedWhenWoVersion(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request": baseClusterRequestBlock(map[string]interface{}{
			"cluster_connection_type": "SERVICE_ACCOUNT",
			"bearer_token_wo_version": 1,
			"bearer_token":            "",
		}),
	})

	resp := fakeServicev1Cluster("test-cluster")
	resp.Cluster.Config.BearerToken = "server-bearer-token"

	setClusterDetails(d, &resp)

	got := d.Get("request.0.cluster.0.config.0.bearer_token").(string)
	if got != "" {
		t.Fatalf("expected bearer_token to remain masked when bearer_token_wo_version is set, got %q", got)
	}
}

// TestPreserveClusterWoVersions_AllFieldsRoundtrip verifies all 5 cluster _wo_version fields
// survive a preserveClusterWoVersions call.
func TestPreserveClusterWoVersions_AllFieldsRoundtrip(t *testing.T) {
	d := newTestResourceData(t, map[string]interface{}{
		"identifier": "test-cluster",
		"agent_id":   "test-agent",
		"request": baseClusterRequestBlock(map[string]interface{}{
			"cluster_connection_type": "SERVICE_ACCOUNT",
			"password_wo_version":     6,
			"bearer_token_wo_version": 7,
			"tls_client_config": []interface{}{
				map[string]interface{}{
					"insecure":             false,
					"server_name":          "",
					"cert_data":            "",
					"key_data":             "",
					"ca_data":              "",
					"cert_data_wo_version": 8,
					"key_data_wo_version":  9,
					"ca_data_wo_version":   10,
				},
			},
		}),
	})

	preserveClusterWoVersions(d)

	cases := map[string]int{
		"request.0.cluster.0.config.0.password_wo_version":                      6,
		"request.0.cluster.0.config.0.bearer_token_wo_version":                  7,
		"request.0.cluster.0.config.0.tls_client_config.0.cert_data_wo_version": 8,
		"request.0.cluster.0.config.0.tls_client_config.0.key_data_wo_version":  9,
		"request.0.cluster.0.config.0.tls_client_config.0.ca_data_wo_version":   10,
	}
	for field, want := range cases {
		got, ok := d.GetOk(field)
		if !ok || got.(int) != want {
			t.Errorf("field %q: want %d after preserveClusterWoVersions, got %v (ok=%v)", field, want, got, ok)
		}
	}
}

// TestClusterSchema_PasswordIsComputed verifies password has Computed=true so that legacy
// customers whose state has password="mypass" do not get a perpetual plan diff when the
// API redacts the field on read (returns ""). Without Computed=true, the SDK zero-fills the
// absent key to "" and the plan shows "" → "mypass" every run.
func TestClusterSchema_PasswordIsComputed(t *testing.T) {
	r := ResourceGitopsCluster()
	configElem := r.Schema["request"].Elem.(*schema.Resource).
		Schema["cluster"].Elem.(*schema.Resource).
		Schema["config"].Elem.(*schema.Resource)

	s, ok := configElem.Schema["password"]
	if !ok {
		t.Fatal("password field not found in config schema")
	}
	if !s.Computed {
		t.Error("expected Computed=true on password for backward compatibility (SDK preserves prior state when API redacts)")
	}
}
