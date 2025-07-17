package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestExpandAgentConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    *svcdiscovery.DatabaseAgentConfiguration
		wantErr bool
	}{
		{
			name:    "Empty config",
			input:   []interface{}{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Full config with all fields",
			input: []interface{}{
				map[string]interface{}{
					"collector_image":    "harness/collector:latest",
					"log_watcher_image":  "harness/watcher:latest",
					"skip_secure_verify": true,
					"image_pull_secrets": []interface{}{"secret1", "secret2"},
					"kubernetes": []interface{}{
						map[string]interface{}{
							"namespace":                  "test-ns",
							"service_account":            "test-sa",
							"disable_namespace_creation": true,
							"image_pull_policy":          "Always",
							"namespaced":                 true,
							"run_as_user":                1000,
							"run_as_group":               2000,
							"labels": map[string]interface{}{
								"app": "test",
								"env": "dev",
							},
							"annotations": map[string]interface{}{
								"test.com/annotation": "value",
							},
							"node_selector": map[string]interface{}{
								"kubernetes.io/os": "linux",
							},
							"resources": []interface{}{
								map[string]interface{}{
									"limits": []interface{}{
										map[string]interface{}{
											"cpu":    "500m",
											"memory": "512Mi",
										},
									},
									"requests": []interface{}{
										map[string]interface{}{
											"cpu":    "250m",
											"memory": "256Mi",
										},
									},
								},
							},
							"tolerations": []interface{}{
								map[string]interface{}{
									"key":                "key1",
									"operator":           "Equal",
									"value":              "value1",
									"effect":             "NoSchedule",
									"toleration_seconds": 3600,
								},
							},
						},
					},
					"data": []interface{}{
						map[string]interface{}{
							"enable_node_agent":        true,
							"node_agent_selector":      "app=test",
							"enable_batch_resources":   true,
							"enable_orphaned_pod":      true,
							"namespace_selector":       "environment=dev",
							"collection_window_in_min": 15,
							"blacklisted_namespaces":   []interface{}{"kube-system", "kube-public"},
							"observed_namespaces":      []interface{}{"default", "harness"},
							"cron": []interface{}{
								map[string]interface{}{
									"expression": "0/10 * * * *",
								},
							},
						},
					},
					"mtls": []interface{}{
						map[string]interface{}{
							"cert_path":   "/etc/certs/tls.crt",
							"key_path":    "/etc/certs/tls.key",
							"secret_name": "mtls-secret",
							"url":         "https://mtls.example.com:8443",
						},
					},
					"proxy": []interface{}{
						map[string]interface{}{
							"http_proxy":  "http://proxy.example.com:8080",
							"https_proxy": "https://proxy.example.com:8080",
							"no_proxy":    "localhost,127.0.0.1,.svc,.cluster.local",
							"url":         "https://proxy.example.com",
						},
					},
				},
			},
			want: &svcdiscovery.DatabaseAgentConfiguration{
				CollectorImage:   "harness/collector:latest",
				LogWatcherImage:  "harness/watcher:latest",
				SkipSecureVerify: true,
				ImagePullSecrets: []string{"secret1", "secret2"},
				Kubernetes: &svcdiscovery.DatabaseKubernetesAgentConfiguration{
					Namespace:                "test-ns",
					ServiceAccount:           "test-sa",
					DisableNamespaceCreation: true,
					ImagePullPolicy:          "Always",
					Namespaced:               true,
					RunAsUser:                1000,
					RunAsGroup:               2000,
					Labels: map[string]string{
						"app": "test",
						"env": "dev",
					},
					Annotations: map[string]string{
						"test.com/annotation": "value",
					},
					NodeSelector: map[string]string{
						"kubernetes.io/os": "linux",
					},
					Resources: &svcdiscovery.DatabaseResourceRequirements{
						Limits: &svcdiscovery.DatabaseResourceList{
							Cpu:    "500m",
							Memory: "512Mi",
						},
						Requests: &svcdiscovery.DatabaseResourceList{
							Cpu:    "250m",
							Memory: "256Mi",
						},
					},
					Tolerations: []svcdiscovery.V1Toleration{
						{
							Key:               "key1",
							Operator:          "Equal",
							Value:             "value1",
							Effect:            "NoSchedule",
							TolerationSeconds: int32(3600),
						},
					},
				},
				Data: &svcdiscovery.DatabaseDataCollectionConfiguration{
					EnableNodeAgent:       true,
					NodeAgentSelector:     "app=test",
					EnableBatchResources:  true,
					EnableOrphanedPod:     true,
					NamespaceSelector:     "environment=dev",
					CollectionWindowInMin: int32(15),
					BlacklistedNamespaces: []string{"kube-system", "kube-public"},
					ObservedNamespaces:    []string{"default", "harness"},
					Cron: &svcdiscovery.DatabaseCronConfig{
						Expression: "0/10 * * * *",
					},
				},
				Mtls: &svcdiscovery.DatabaseMtlsConfiguration{
					CertPath:   "/etc/certs/tls.crt",
					KeyPath:    "/etc/certs/tls.key",
					SecretName: "mtls-secret",
					Url:        "https://mtls.example.com:8443",
				},
				Proxy: &svcdiscovery.DatabaseProxyConfiguration{
					HttpProxy:  "http://proxy.example.com:8080",
					HttpsProxy: "https://proxy.example.com:8080",
					NoProxy:    "localhost,127.0.0.1,.svc,.cluster.local",
					Url:        "https://proxy.example.com",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandAgentConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandAgentConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFlattenAgentConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   *svcdiscovery.DatabaseAgentConfiguration
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			want:    nil,
			wantErr: false,
		},
		{
			name: "Full config",
			input: &svcdiscovery.DatabaseAgentConfiguration{
				CollectorImage:   "test-image",
				LogWatcherImage:  "log-watcher",
				SkipSecureVerify: true,
				ImagePullSecrets: []string{"secret1", "secret2"},
			},
			want: map[string]interface{}{
				"collector_image":    "test-image",
				"log_watcher_image":  "log-watcher",
				"skip_secure_verify": true,
				"image_pull_secrets": []string{"secret1", "secret2"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlattenAgentConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlattenAgentConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
