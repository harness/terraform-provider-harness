package cd

import (
	"fmt"
	"log"
	"reflect"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
)

func (c *ConfigAsCodeClient) UpsertInfraDefinition(input *cac.InfrastructureDefinition) (*cac.InfrastructureDefinition, error) {
	log.Printf("[DEBUG] CAC: Upsert infra definition %s", input.Name)
	if ok, err := input.Validate(); !ok {
		return nil, err
	}

	app, err := c.ApiClient.ApplicationClient.GetApplicationById(input.ApplicationId)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, fmt.Errorf("application %s not found", input.ApplicationId)
	}

	env, err := c.ApiClient.ConfigAsCodeClient.GetEnvironmentById(input.ApplicationId, input.EnvironmentId)
	if err != nil {
		return nil, err
	}

	if env == nil {
		return nil, fmt.Errorf("environment %s not found", input.EnvironmentId)
	}

	yamlPath := cac.GetInfraDefinitionYamlPath(app.Name, env.Name, input.Name)
	infra := &cac.InfrastructureDefinition{}
	err = c.UpsertObject(input, yamlPath, infra)
	if err != nil {
		return nil, err
	}

	infra.EnvironmentId = env.Id

	return infra, nil
}

func (c *ConfigAsCodeClient) GetInfraDefinitionById(appId string, envId string, infraId string) (*cac.InfrastructureDefinition, error) {
	log.Printf("[DEBUG] CAC: Get infra definition by id %s", infraId)
	app, err := c.ApiClient.ApplicationClient.GetApplicationById(appId)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, fmt.Errorf("application %s not found", appId)
	}

	env, err := c.ApiClient.ConfigAsCodeClient.GetEnvironmentById(app.Id, envId)
	if err != nil {
		return nil, err
	}

	if env == nil {
		return nil, fmt.Errorf("environment %s not found", envId)
	}

	infra := &cac.InfrastructureDefinition{}
	err = c.FindObjectById(appId, infraId, infra)
	if err != nil {
		return nil, err
	}

	if infra.IsEmpty() {
		return nil, nil
	}

	infra.EnvironmentId = env.Id

	return infra, nil
}

func (c *ConfigAsCodeClient) GetInfraDefinitionByName(appId string, envId string, infraName string) (*cac.InfrastructureDefinition, error) {
	log.Printf("[DEBUG] CAC: Get infra definition by name %s", infraName)
	app, err := c.ApiClient.ApplicationClient.GetApplicationById(appId)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, fmt.Errorf("application %s not found", appId)
	}

	env, err := c.ApiClient.ConfigAsCodeClient.GetEnvironmentById(app.Id, envId)
	if err != nil {
		return nil, err
	}

	if env == nil {
		return nil, fmt.Errorf("environment %s not found", envId)
	}

	path := cac.GetInfraDefinitionYamlPath(app.Name, env.Name, infraName)
	infraDef := &cac.InfrastructureDefinition{}

	err = c.FindObjectByPath(app.Id, path, infraDef)
	if err != nil {
		return nil, err
	}

	return infraDef, nil
}

func (c *ConfigAsCodeClient) DeleteInfraDefinition(applicationId string, environmentId string, infraId string) error {
	log.Printf("[DEBUG] CAC: Delete infra definition %s", infraId)
	app, err := c.ApiClient.ApplicationClient.GetApplicationById(applicationId)
	if err != nil {
		return err
	}

	if app == nil {
		return fmt.Errorf("could not find application by id: '%s'", applicationId)
	}

	env, err := c.GetEnvironmentById(applicationId, environmentId)
	if err != nil {
		return err
	}

	if env == nil {
		return fmt.Errorf("could not find environment by id: '%s'", environmentId)
	}

	infraDef, err := c.GetInfraDefinitionById(applicationId, environmentId, infraId)
	if err != nil {
		return err
	}

	filePath := cac.GetInfraDefinitionYamlPath(app.Name, env.Name, infraDef.Name)

	return c.DeleteEntity(filePath)
}

var InfrastructureDefinitionTypeMap = map[cac.InfrastructureType]reflect.Type{
	cac.InfrastructureTypes.AwsAmi:           reflect.TypeOf(cac.InfrastructureAwsAmi{}),
	cac.InfrastructureTypes.AwsEcs:           reflect.TypeOf(cac.InfrastructureAwsEcs{}),
	cac.InfrastructureTypes.AwsLambda:        reflect.TypeOf(cac.InfrastructureAwsLambda{}),
	cac.InfrastructureTypes.AwsSSH:           reflect.TypeOf(cac.InfrastructureAwsSSH{}),
	cac.InfrastructureTypes.AzureVmss:        reflect.TypeOf(cac.InfrastructureAzureVmss{}),
	cac.InfrastructureTypes.Custom:           reflect.TypeOf(cac.InfrastructureCustom{}),
	cac.InfrastructureTypes.DataCenterSSH:    reflect.TypeOf(cac.InfrastructureDataCenterSSH{}),
	cac.InfrastructureTypes.DataCenterWinRM:  reflect.TypeOf(cac.InfrastructureDataCenterWinRM{}),
	cac.InfrastructureTypes.KubernetesDirect: reflect.TypeOf(cac.InfrastructureKubernetesDirect{}),
	cac.InfrastructureTypes.KubernetesGcp:    reflect.TypeOf(cac.InfrastructureKubernetesGcp{}),
	cac.InfrastructureTypes.Pcf:              reflect.TypeOf(cac.InfrastructureTanzu{}),
}
