resource "harness_yaml_config" "test" {
  path    = "Setup/Cloud Providers/Kubernetes.yaml"
  content = <<EOF
harnessApiVersion: '1.0'
type: KUBERNETES_CLUSTER
delegateSelectors:
- k8s
skipValidation: true
useKubernetesDelegate: true
EOF
}
