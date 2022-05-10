package cac

import (
	"reflect"

	"github.com/harness/harness-go-sdk/harness/utils"
)

type Service struct {
	HarnessApiVersion         HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type                      ObjectType         `yaml:"type" json:"type"`
	Id                        string             `yaml:"-"`
	Name                      string             `yaml:"-"`
	ArtifactType              ArtifactType       `yaml:"artifactType,omitempty"`
	DeploymentType            DeploymentType     `yaml:"deploymentType,omitempty"`
	Description               string             `yaml:"description,omitempty"`
	Tags                      map[string]string  `yaml:"tags,omitempty"`
	HelmVersion               HelmVersion        `yaml:"helmVersion,omitempty"`
	ApplicationId             string             `yaml:"-"`
	DeploymentTypeTemplateUri string             `yaml:"deploymentTypeTemplateUri,omitempty"`
	ConfigVariables           []*ServiceVariable `yaml:"configVariables,omitempty"`
}

func (a *Service) IsEmpty() bool {
	return reflect.DeepEqual(a, &Service{})
}

func (s *Service) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(s, []string{"ApplicationId"})
}
