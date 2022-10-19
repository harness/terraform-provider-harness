package nextgen

type InfrastructureDeploymenType string

var InfrastructureDeploymenTypes = struct {
	Kubernetes          InfrastructureDeploymenType
	NativeHelm          InfrastructureDeploymenType
	Ssh                 InfrastructureDeploymenType
	WinRm               InfrastructureDeploymenType
	ServerlessAwsLambda InfrastructureDeploymenType
	AzureWebApp         InfrastructureDeploymenType
	CustomDeployment    InfrastructureDeploymenType
	ECS                 InfrastructureDeploymenType
}{
	Kubernetes:          "Kubernetes",
	NativeHelm:          "NativeHelm",
	Ssh:                 "Ssh",
	WinRm:               "WinRm",
	ServerlessAwsLambda: "ServerlessAwsLambda",
	AzureWebApp:         "AzureWebApp",
	CustomDeployment:    "CustomDeployment",
	ECS:                 "ECS",
}

var InfrastructureDeploymentypeValues = []string{
	InfrastructureTypes.KUBERNETES_DIRECT.String(),
	InfrastructureTypes.KUBERNETES_GCP.String(),
	InfrastructureTypes.SERVERLESS_AWS_LAMBDA.String(),
	InfrastructureTypes.PDC.String(),
	InfrastructureTypes.KUBERNETES_AZURE.String(),
	InfrastructureTypes.SSH_WINRM_AZURE.String(),
	InfrastructureTypes.SSH_WINRM_AWS.String(),
	InfrastructureTypes.AZURE_WEB_APP.String(),
	InfrastructureTypes.ECS.String(),
	InfrastructureTypes.GITOPS.String(),
	InfrastructureTypes.CUSTOM_DEPLOYMENT.String(),
}

func (e InfrastructureDeploymenType) String() string {
	return string(e)
}
