package cd

import (
	"errors"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
)

func (c *ConfigAsCodeClient) UpsertEnvironment(input *cac.Environment) (*cac.Environment, error) {
	if input == nil {
		return nil, errors.New("cannot create environment. input is nil")
	}

	if ok, err := input.Validate(); !ok {
		return nil, err
	}

	app, err := c.ApiClient.ApplicationClient.GetApplicationById(input.ApplicationId)
	if err != nil {
		return nil, fmt.Errorf("could not find application '%s'", app.Id)
	}

	output := &cac.Environment{}
	filePath := cac.GetEnvironmentYamlPath(app.Name, input.Name)
	err = c.ApiClient.ConfigAsCodeClient.UpsertObject(input, filePath, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *ConfigAsCodeClient) GetEnvironmentByName(applicationId string, environmentName string) (*cac.Environment, error) {
	if ok, err := utils.CheckRequiredParameters(applicationId, ""); !ok {
		return nil, err
	}

	if ok, err := utils.CheckRequiredParameters(environmentName, ""); !ok {
		return nil, err
	}

	app, err := c.ApiClient.ApplicationClient.GetApplicationById(applicationId)
	if err != nil {
		return nil, err
	} else if app == nil {
		return nil, fmt.Errorf("could not find application '%s'", applicationId)
	}

	output := &cac.Environment{}
	yamlPath := cac.GetEnvironmentYamlPath(app.Name, environmentName)
	err = c.FindObjectByPath(applicationId, yamlPath, output)
	if err != nil {
		return nil, err
	}

	if output.IsEmpty() {
		return nil, nil
	}

	return output, nil
}

func (c *ConfigAsCodeClient) GetEnvironmentById(applicationId string, environmentId string) (*cac.Environment, error) {
	if ok, err := utils.CheckRequiredParameters(applicationId, ""); !ok {
		return nil, err
	}

	if ok, err := utils.CheckRequiredParameters(environmentId, ""); !ok {
		return nil, err
	}

	env := &cac.Environment{}
	err := c.FindObjectById(applicationId, environmentId, env)
	if err != nil {
		return nil, err
	}

	if env.IsEmpty() {
		return nil, nil
	}

	return env, nil
}

func (c *ConfigAsCodeClient) DeleteEnvironment(appName string, envName string) error {
	return c.DeleteEntity(cac.GetEnvironmentYamlPath(appName, envName))
}
