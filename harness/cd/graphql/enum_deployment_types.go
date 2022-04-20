package graphql

type DeploymentType string

var DeploymentTypes = struct {
	AMI           DeploymentType
	AWSCodeDeploy DeploymentType
	AWSLambda     DeploymentType
	AzureVMSS     DeploymentType
	AzureWebApp   DeploymentType
	Custom        DeploymentType
	ECS           DeploymentType
	Helm          DeploymentType
	Kubernetes    DeploymentType
	PCF           DeploymentType
	SSH           DeploymentType
	WinRM         DeploymentType
}{
	AMI:           "AMI",
	AWSCodeDeploy: "AWS_CODEDEPLOY",
	AWSLambda:     "AWS_LAMBDA",
	AzureVMSS:     "AZURE_VMSS",
	AzureWebApp:   "AZURE_WEBAPP",
	Custom:        "Custom",
	ECS:           "ECS",
	Helm:          "HELM",
	Kubernetes:    "KUBERNETES",
	PCF:           "PCF",
	SSH:           "SSH",
	WinRM:         "WINRM",
}

var DeploymentTypeValues = []string{
	DeploymentTypes.AMI.String(),
	DeploymentTypes.AWSCodeDeploy.String(),
	DeploymentTypes.AWSLambda.String(),
	DeploymentTypes.AzureVMSS.String(),
	DeploymentTypes.AzureWebApp.String(),
	DeploymentTypes.Custom.String(),
	DeploymentTypes.ECS.String(),
	DeploymentTypes.Helm.String(),
	DeploymentTypes.Kubernetes.String(),
	DeploymentTypes.PCF.String(),
	DeploymentTypes.SSH.String(),
	DeploymentTypes.WinRM.String(),
}

func (e DeploymentType) String() string {
	return string(e)
}
