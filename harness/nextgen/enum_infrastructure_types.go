package nextgen

type InfrastructureType string

var InfrastructureTypes = struct {
	KUBERNETES_DIRECT     InfrastructureType
	KUBERNETES_GCP        InfrastructureType
	SERVERLESS_AWS_LAMBDA InfrastructureType
	PDC                   InfrastructureType
	KUBERNETES_AZURE      InfrastructureType
	SSH_WINRM_AZURE       InfrastructureType
	SSH_WINRM_AWS         InfrastructureType
	AZURE_WEB_APP         InfrastructureType
	ECS                   InfrastructureType
	GITOPS                InfrastructureType
	CUSTOM_DEPLOYMENT     InfrastructureType
}{
	KUBERNETES_DIRECT:     "KubernetesDirect",
	KUBERNETES_GCP:        "KubernetesGcp",
	SERVERLESS_AWS_LAMBDA: "ServerlessAwsLambda",
	PDC:                   "Pdc",
	KUBERNETES_AZURE:      "KubernetesAzure",
	SSH_WINRM_AZURE:       "SshWinRmAzure",
	SSH_WINRM_AWS:         "SshWinRmAws",
	AZURE_WEB_APP:         "AzureWebApp",
	ECS:                   "ECS",
	GITOPS:                "GitOps",
	CUSTOM_DEPLOYMENT:     "CustomDeployment",
}

var InfrastructureTypeValues = []string{
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

func (e InfrastructureType) String() string {
	return string(e)
}
