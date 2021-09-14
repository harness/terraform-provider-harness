# Importing a global config only using the yaml path
terraform import harness_yaml_config.k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"

# Importing a service which requires both the application id and the yaml path.
terraform import harness_yaml_config.k8s_cloudprovider "Setup/Applications/MyApp/Services/MyService/Index.yaml:<APPLICATION_ID>"
