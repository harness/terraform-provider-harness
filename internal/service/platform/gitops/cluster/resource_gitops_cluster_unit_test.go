package cluster

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func newTestResourceData(t *testing.T, values map[string]interface{}) *schema.ResourceData {
	t.Helper()
	r := ResourceGitopsCluster()
	d := r.TestResourceData()
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("failed to set %q: %s", k, err)
		}
	}
	return d
}

func TestBuildClusterDetails_IRSAExcludesBearerToken(t *testing.T) {
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
						"server": "https://eks.amazonaws.com",
						"name":   "my-cluster",
						"config": []interface{}{
							map[string]interface{}{
								"bearer_token":            "stale-token-from-state",
								"username":                "old-user",
								"password":                "old-pass",
								"role_a_r_n":              "arn:aws:iam::123456:role/my-role",
								"aws_cluster_name":        "my-eks-cluster",
								"cluster_connection_type": "IRSA",
								"disable_compression":     false,
								"proxy_url":               "",
								"tls_client_config": []interface{}{
									map[string]interface{}{
										"insecure":    true,
										"server_name": "",
										"cert_data":   "",
										"key_data":    "",
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
	if cluster.Config.ClusterConnectionType != "IRSA" {
		t.Errorf("expected ClusterConnectionType=IRSA, got %q", cluster.Config.ClusterConnectionType)
	}
	if cluster.Config.BearerToken != "" {
		t.Errorf("expected BearerToken to be empty when connection type is IRSA, got %q", cluster.Config.BearerToken)
	}
	if cluster.Config.Username != "" {
		t.Errorf("expected Username to be empty when connection type is IRSA, got %q", cluster.Config.Username)
	}
	if cluster.Config.Password != "" {
		t.Errorf("expected Password to be empty when connection type is IRSA, got %q", cluster.Config.Password)
	}
	if cluster.Config.RoleARN != "arn:aws:iam::123456:role/my-role" {
		t.Errorf("expected RoleARN to be set for IRSA, got %q", cluster.Config.RoleARN)
	}
	if cluster.Config.AwsClusterName != "my-eks-cluster" {
		t.Errorf("expected AwsClusterName to be set for IRSA, got %q", cluster.Config.AwsClusterName)
	}
}

func TestBuildClusterDetails_ServiceAccountExcludesIRSAFields(t *testing.T) {
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
								"bearer_token":            "valid-token",
								"username":                "",
								"password":                "",
								"role_a_r_n":              "stale-arn-from-state",
								"aws_cluster_name":        "stale-cluster-name",
								"cluster_connection_type": "SERVICE_ACCOUNT",
								"disable_compression":     false,
								"proxy_url":               "",
								"tls_client_config": []interface{}{
									map[string]interface{}{
										"insecure":    true,
										"server_name": "",
										"cert_data":   "",
										"key_data":    "",
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
	if cluster.Config.ClusterConnectionType != "SERVICE_ACCOUNT" {
		t.Errorf("expected ClusterConnectionType=SERVICE_ACCOUNT, got %q", cluster.Config.ClusterConnectionType)
	}
	if cluster.Config.BearerToken != "valid-token" {
		t.Errorf("expected BearerToken to be set for SERVICE_ACCOUNT, got %q", cluster.Config.BearerToken)
	}
	if cluster.Config.RoleARN != "" {
		t.Errorf("expected RoleARN to be empty when connection type is SERVICE_ACCOUNT, got %q", cluster.Config.RoleARN)
	}
	if cluster.Config.AwsClusterName != "" {
		t.Errorf("expected AwsClusterName to be empty when connection type is SERVICE_ACCOUNT, got %q", cluster.Config.AwsClusterName)
	}
}

func TestBuildClusterDetails_NoConnectionTypePreservesAllFields(t *testing.T) {
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
								"bearer_token":            "my-token",
								"username":                "",
								"password":                "",
								"role_a_r_n":              "some-arn",
								"aws_cluster_name":        "some-cluster",
								"cluster_connection_type": "",
								"disable_compression":     false,
								"proxy_url":               "",
								"tls_client_config": []interface{}{
									map[string]interface{}{
										"insecure":    true,
										"server_name": "",
										"cert_data":   "",
										"key_data":    "",
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
	if cluster.Config.BearerToken != "my-token" {
		t.Errorf("expected BearerToken preserved when no connection type set, got %q", cluster.Config.BearerToken)
	}
	if cluster.Config.RoleARN != "some-arn" {
		t.Errorf("expected RoleARN preserved when no connection type set, got %q", cluster.Config.RoleARN)
	}
	if cluster.Config.AwsClusterName != "some-cluster" {
		t.Errorf("expected AwsClusterName preserved when no connection type set, got %q", cluster.Config.AwsClusterName)
	}
}
