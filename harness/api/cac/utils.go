package cac

import (
	"errors"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"gopkg.in/yaml.v3"
)

func (i *ConfigAsCodeItem) ParseYamlContent() (interface{}, error) {
	if i.Yaml == "" {
		return nil, nil
	}

	tmp := map[string]interface{}{}
	data := []byte(i.Yaml)

	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	val, ok := tmp["type"]
	if !ok {
		return nil, errors.New("could not find field 'type' in yaml object")
	}

	switch val {
	case ObjectTypes.Service:
		obj := &Service{}
		if err := yaml.Unmarshal(data, &obj); err != nil {
			return nil, err
		}
		obj.Name = utils.TrimFileExtension(i.Name)
		return obj, err
	default:
		return nil, fmt.Errorf("could not parse object type of '%s'", val)
	}

}

func (s *Service) Validate() (bool, error) {
	if s.ApplicationId == "" {
		return false, errors.New("service is invalid. missing field `ApplicationId`")
	}

	return true, nil
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

func GetDefaultArtifactType(deploymentType string, fallbackArtifactType string) (string, error) {

	var artifactType string

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
