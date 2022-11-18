package nextgen

type InfrastructureDeploymenType string

var InfrastructureDeploymenTypes = struct {
	KUBERNETES            InfrastructureDeploymenType
	NATIVE_HELM           InfrastructureDeploymenType
	SSH                   InfrastructureDeploymenType
	WINRM                 InfrastructureDeploymenType
	SERVERLESS_AWS_LAMBDA InfrastructureDeploymenType
	AZURE_WEBAPP          InfrastructureDeploymenType
	CUSTOM                InfrastructureDeploymenType
	ECS                   InfrastructureDeploymenType
}{
	KUBERNETES:            "Kubernetes",
	NATIVE_HELM:           "NativeHelm",
	SSH:                   "Ssh",
	WINRM:                 "WinRm",
	SERVERLESS_AWS_LAMBDA: "ServerlessAwsLambda",
	AZURE_WEBAPP:          "AzureWebApp",
	CUSTOM:                "Custom",
	ECS:                   "ECS",
}

var InfrastructureDeploymentypeValues = []string{
	InfrastructureDeploymenTypes.KUBERNETES.String(),
	InfrastructureDeploymenTypes.NATIVE_HELM.String(),
	InfrastructureDeploymenTypes.SSH.String(),
	InfrastructureDeploymenTypes.WINRM.String(),
	InfrastructureDeploymenTypes.SERVERLESS_AWS_LAMBDA.String(),
	InfrastructureDeploymenTypes.AZURE_WEBAPP.String(),
	InfrastructureDeploymenTypes.CUSTOM.String(),
	InfrastructureDeploymenTypes.ECS.String(),
}

func (e InfrastructureDeploymenType) String() string {
	return string(e)
}
