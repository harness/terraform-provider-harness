package cac

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"gopkg.in/yaml.v3"
)

func (i *ConfigAsCodeItem) ParseYamlContent(respObj interface{}) error {
	if i.Yaml == "" {
		return nil
	}

	if err := yaml.Unmarshal([]byte(i.Yaml), respObj); err != nil {
		return err
	}

	return nil
}

func (s *Service) Validate() (bool, error) {
	if s.ApplicationId == "" {
		return false, errors.New("service is invalid. missing field `ApplicationId`")
	}

	return true, nil
}

func (cp *GcpCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

func (cp *PhysicalDatacenterCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

func (i *ConfigAsCodeItem) IsEmpty() bool {
	return i == &ConfigAsCodeItem{}
}

// Indicates an error condition
func (r *Response) IsEmpty() bool {
	// return true
	return r.Metadata == ResponseMetadata{} && r.Resource.IsEmpty() && len(r.ResponseMessages) == 0
}

func (m *ResponseMessage) ToError() error {
	return fmt.Errorf("%s: %s", m.Code, m.Message)
}

func GetDefaultArtifactType(deploymentType DeploymentType, fallbackArtifactType ArtifactType) (ArtifactType, error) {

	var artifactType ArtifactType

	switch deploymentType {
	case DeploymentTypes.Kubernetes:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.SSH:
		artifactType = fallbackArtifactType
	case DeploymentTypes.AMI:
		artifactType = ArtifactTypes.AMI
	case DeploymentTypes.AWSCodeDeploy:
		artifactType = ArtifactTypes.AWSCodeDeploy
	case DeploymentTypes.AWSLambda:
		artifactType = ArtifactTypes.AWSLambda
	case DeploymentTypes.ECS:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.PCF:
		artifactType = ArtifactTypes.PCF
	case DeploymentTypes.Helm:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.WinRM:
		artifactType = fallbackArtifactType
	default:
		return "", fmt.Errorf("no default artifact type for '%s' deployments", deploymentType)
	}

	return artifactType, nil
}

func NewEntity(objectType ObjectType) interface{} {
	t, ok := objectTypeMap[objectType]
	if !ok {
		panic(fmt.Errorf("could not create entity of type `%s`", objectType))
	}

	i := reflect.New(t).Interface()
	utils.MustSetField(i, "HarnessApiVersion", HarnessApiVersions.V1)
	utils.MustSetField(i, "Type", objectType)
	return i
}

var objectTypeMap = map[ObjectType]reflect.Type{
	ObjectTypes.Application:                     reflect.TypeOf(Application{}),
	ObjectTypes.AwsCloudProvider:                reflect.TypeOf(AwsCloudProvider{}),
	ObjectTypes.AzureCloudProvider:              reflect.TypeOf(AzureCloudProvider{}),
	ObjectTypes.GcpCloudProvider:                reflect.TypeOf(GcpCloudProvider{}),
	ObjectTypes.KubernetesCloudProvider:         reflect.TypeOf(KubernetesCloudProvider{}),
	ObjectTypes.PcfCloudProvider:                reflect.TypeOf(PcfCloudProvider{}),
	ObjectTypes.PhysicalDataCenterCloudProvider: reflect.TypeOf(PhysicalDatacenterCloudProvider{}),
	ObjectTypes.Service:                         reflect.TypeOf(Service{}),
	ObjectTypes.SpotInstCloudProvider:           reflect.TypeOf(SpotInstCloudProvider{}),
	ObjectTypes.Application:                     reflect.TypeOf(Application{}),
}

func (r *SecretRef) MarshalYAML() (interface{}, error) {
	if (r == &SecretRef{}) {
		return []byte{}, nil
	}

	if r.SecretId == "" {
		return nil, errors.New("SecretId must be set")
	}

	// if r.SecretManagerType == "" {
	// 	return nil, errors.New("SecretManagerType must be set")
	// }

	if r.SecretManagerType == "" {
		return r.SecretId, nil
	}

	return fmt.Sprintf("%s:%s", r.SecretManagerType, r.SecretId), nil
}

func (r *SecretRef) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var val interface{}
	err := unmarshal(&val)
	if err != nil {
		return err
	}

	value := val.(string)

	parts := strings.Split(value, ":")
	r.SecretManagerType = SecretManagerType(parts[0])
	r.SecretId = parts[1]

	return nil
}

func GetEntityNameFromPath(yamlPath YamlPath) string {
	dir, file := path.Split(string(yamlPath))

	if ok, _ := regexp.MatchString("Index.yaml", file); ok {
		parts := strings.Split(strings.TrimSpace(dir), "/")
		last := parts[len(parts)-2]
		return last
	}

	return utils.TrimFileExtension(file)
}

func GetServiceYamlPath(applicationName string, serviceName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", applicationName, serviceName))
}

func GetCloudProviderYamlPath(cloudProviderName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Cloud Providers/%s.yaml", cloudProviderName))
}

func GetApplicationYamlPath(applicationName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/index.yaml", applicationName))
}
