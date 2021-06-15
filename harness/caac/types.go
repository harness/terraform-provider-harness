package caac

import (
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/graphql"
)

type ConfigAsCodeClient struct {
	ApiClient *graphql.ApiClient
}

type Response struct {
	Metadata         ResponseMetadata  `json:"metaData"`
	Resource         ConfigAsCodeItem  `json:"resource"`
	ResponseMessages []ResponseMessage `json:"responseMessages"`
}

type ConfigAsCodeItem struct {
	AccountId       string              `json:"accountId,omitempty"`
	Type            string              `json:"type,omitempty"`
	Name            string              `json:"name,omitempty"`
	ClassName       string              `json:"className,omitempty"`
	ShortClassName  string              `json:"shortClassName,omitempty"`
	RestName        string              `json:"restName,omitempty"`
	DirectoryPath   *DirectoryPath      `json:"directoryPath,omitempty"`
	DefaultToClosed bool                `json:"defaultToClosed,omitempty"`
	Children        []*ConfigAsCodeItem `json:"children,omitempty"`
	AppId           string              `json:"appId,omitempty"`
	YamlGitConfig   interface{}         `json:"yamlGitConfig,omitempty"`
	UUID            string              `json:"uuid,omitempty"`
	YamlVersionType string              `json:"yamlVersionType,omitempty"`
	YamlFilePath    string              `json:"yamlFilePath,omitempty"`
	Status          string              `json:"status,omitempty"`
	ErrorMessage    string              `json:"errorMssg,omitempty"`
	Yaml            string              `json:"yaml"`
}

type classType struct {
	Account                   string
	Application               string
	ArtifactStream            string
	ConfigFile                string
	Defaults                  string
	Environment               string
	InfrastructureProvisioner string
	NotificationGroup         string
	Pipeline                  string
	Service                   string
	SettingAttribute          string
	Tags                      string
	Template                  string
	Workflow                  string
}

var ClassTypes classType = classType{
	Account:                   "Account",
	Application:               "Application",
	ArtifactStream:            "ArtifactStream",
	ConfigFile:                "ConfigFile",
	Defaults:                  "Defaults",
	Environment:               "Environment",
	InfrastructureProvisioner: "InfrastructureProvisioner",
	NotificationGroup:         "NotificationGroup",
	Pipeline:                  "Pipeline",
	Service:                   "Service",
	SettingAttribute:          "SettingAttribute",
	Tags:                      "HarnessTag",
	Template:                  "Template",
	Workflow:                  "Workflow",
}

type DirectoryPath struct {
	Path string `json:"path,omitempty"`
}

type ResponseMetadata struct{}

type ResponseMessage struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type Service struct {
	HarnessApiVersion string       `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              string       `yaml:"type" json:"type"`
	ArtifactType      string       `yaml:"artifactType,omitempty"`
	CreatedAt         graphql.Time `yaml:"createdAt,omitempty"`
	CreatedBy         graphql.User `yaml:"createdBy,omitempty"`
	DeploymentType    string       `yaml:"deploymentType,omitempty"`
	Description       string       `yaml:"description,omitempty"`
	Id                string       `yaml:"id,omitempty"`
	Name              string       `yaml:"-"`
	Tags              []*Tag       `yaml:"tags,omitempty"`
	HelmVersion       string       `yaml:"helmVersion,omitempty"`
}

var APIV1 = "1.0"

// func NewKubernetesService(name string) *Service {
// 	return &Service{
// 		HarnessApiVersion: APIV1,
// 		Type:              ObjectTypes.Service,
// 		Name:              name,
// 		ArtifactType:      ArtifactTypes.Docker,
// 		DeploymentType:    DeploymentTypes.Kubernetes,
// 		HelmVersion:       HelmVersions.V2,
// 	}
// }

// func NewSSHService(name string, artifactType string) *Service {
// 	return &Service{
// 		HarnessApiVersion: APIV1,
// 		Type:              ObjectTypes.Service,
// 		Name:              name,
// 		ArtifactType:      artifactType,
// 		DeploymentType:    DeploymentTypes.SSH,
// 	}
// }

// func NewAMIService(name string) *Service {
// 	return &Service{
// 		HarnessApiVersion: APIV1,
// 	}
// }

func ServiceFactory(name string, deploymentType string, artifactType string) (*Service, error) {
	svc := &Service{
		HarnessApiVersion: APIV1,
		Type:              ObjectTypes.Service,
		Name:              name,
		DeploymentType:    deploymentType,
	}

	switch deploymentType {
	case DeploymentTypes.Kubernetes:
		svc.HelmVersion = HelmVersions.V2
		svc.ArtifactType = ArtifactTypes.Docker
	case DeploymentTypes.SSH:
		svc.ArtifactType = artifactType
	case DeploymentTypes.AMI:
		svc.ArtifactType = ArtifactTypes.AMI
	case DeploymentTypes.AWSCodeDeploy:
		svc.ArtifactType = ArtifactTypes.AWSCodeDeploy
	case DeploymentTypes.AWSLambda:
		svc.ArtifactType = ArtifactTypes.AWSLambda
	case DeploymentTypes.ECS:
		svc.ArtifactType = ArtifactTypes.Docker
	case DeploymentTypes.PCF:
		svc.ArtifactType = ArtifactTypes.PCF
	case DeploymentTypes.Helm:
		svc.ArtifactType = ArtifactTypes.Docker
	case DeploymentTypes.WinRM:
		svc.ArtifactType = artifactType

	default:
		return nil, fmt.Errorf("could not create service of type '%s'", deploymentType)
	}

	return svc, nil
}

type objectType struct {
	Service string
}

var ObjectTypes objectType = objectType{
	Service: "SERVICE",
}

type helmVersion struct {
	V2 string
	V3 string
}

var HelmVersions helmVersion = helmVersion{
	V2: "V2",
	V3: "V3",
}

// type KubernetesService struct {
// 	HarnessApiVersion string      `yaml:"harnessApiVersion" json:"harnessApiVersion"`
// 	Type              string      `yaml:"type" json:"type"`
// 	ArtifactType      string      `yaml:"artifactType,omitempty"`
// 	CreatedAt         graphql.Time `yaml:"createdAt,omitempty"`
// 	CreatedBy         graphql.User `yaml:"createdBy,omitempty"`
// 	DeploymentType    string      `yaml:"deploymentType,omitempty"`
// 	Description       string      `yaml:"description,omitempty"`
// 	Id                string      `yaml:"id,omitempty"`
// 	Name              string      `yaml:"name,omitempty"`
// 	Tags              []*Tag      `yaml:"tags,omitempty"`
// }

type deploymentType struct {
	AMI           string
	AWSCodeDeploy string
	AWSLambda     string
	ECS           string
	Custom        string
	SSH           string
	Kubernetes    string
	Helm          string
	PCF           string
	WinRM         string
}

var DeploymentTypes deploymentType = deploymentType{
	AMI:           "AMI",
	ECS:           "ECS",
	AWSCodeDeploy: "AWS_CODEDEPLOY",
	AWSLambda:     "AWS_LAMBDA",
	Custom:        "Custom",
	Kubernetes:    "KUBERNETES",
	Helm:          "HELM",
	PCF:           "PCF",
	SSH:           "SSH",
	WinRM:         "WINRM",
}

type artifactType struct {
	AMI                 string
	AWSCodeDeploy       string
	AWSLambda           string
	Docker              string
	Jar                 string
	Other               string
	PCF                 string
	Tar                 string
	War                 string
	RPM                 string
	IISVirtualDirectory string
	IISApp              string
	IISWebsite          string
	Zip                 string
}

var ArtifactTypes artifactType = artifactType{
	AMI:                 "AMI",
	AWSCodeDeploy:       "AWS_CODEDEPLOY",
	AWSLambda:           "AWS_LAMBDA",
	Docker:              "DOCKER",
	Jar:                 "JAR",
	Other:               "OTHER",
	PCF:                 "PCF",
	RPM:                 "RPM",
	Tar:                 "TAR",
	War:                 "WAR",
	IISVirtualDirectory: "IIS_VirtualDirectory",
	IISApp:              "IIS_APP",
	IISWebsite:          "IIS",
	Zip:                 "ZIP",
}

type Tag struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}
