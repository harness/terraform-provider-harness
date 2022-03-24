package cd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestUpsertRawYaml_Service(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))
	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	defer func() {
		err := c.ApplicationClient.DeleteApplication(app.Id)
		require.NoError(t, err)
	}()

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: SERVICE
artifactType: DOCKER
deploymentType: KUBERNETES
helmVersion: V2
`)

	yamlBytes := []byte(rawYaml)
	path := cac.GetServiceYamlPath(app.Name, name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(path, yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath(app.Id, path)
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, app.Id, yamlItem.ApplicationId)
	require.Equal(t, path, yamlItem.Path)
	require.NotEmpty(t, yamlItem.Id)

}

func TestUpsertRawYaml_ArtifactSource(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))
	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	defer func() {
		err := c.ApplicationClient.DeleteApplication(app.Id)
		require.NoError(t, err)
	}()

	svc, err := createService(app.Id, app.Name, name, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
	require.NoError(t, err)

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: DOCKER
imageName: waynz0r/echo-server
serverName: Harness Docker Hub
`)

	yamlBytes := []byte(rawYaml)
	path := fmt.Sprintf("Setup/Applications/%s/Services/%s/Artifact Servers/%s.yaml", app.Name, svc.Name, name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(cac.YamlPath(path), yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath(app.Id, cac.YamlPath(path))
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, app.Id, yamlItem.ApplicationId)
	require.Equal(t, cac.YamlPath(path), yamlItem.Path)
	require.NotEmpty(t, yamlItem.Id)

}

func TestUpsertRawYaml_CloudProvider(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: KUBERNETES_CLUSTER
authType: NONE
caCert: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_Nkmm_password
clientCert: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_BS2C_token
clientKey: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_BS2C_token
masterUrl: https://asdfasdfasd.com
password: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_Nkmm_password
serviceAccountToken: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_Nkmm_password
skipValidation: true
useKubernetesDelegate: false
usernameSecretId: gcpkms:TestAccResourceK8sCloudProviderConnector_service_account_Nkmm_token
`)

	yamlBytes := []byte(rawYaml)
	path := cac.GetCloudProviderYamlPath(name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(path, yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath("", path)
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, "", yamlItem.ApplicationId)
	require.Equal(t, path, yamlItem.Path)
	require.NotEmpty(t, yamlItem.Id)

	err = c.ConfigAsCodeClient.DeleteEntityV2(path, rawYaml)
	require.NoError(t, err)
}

func TestUpsertRawYaml_ArtifactServer(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: DOCKER
url: https://registry.hub.docker.com/v2/
usageRestrictions:
  appEnvRestrictions:
  - appFilter:
      filterType: ALL
    envFilter:
      filterTypes:
      - PROD
  - appFilter:
      filterType: ALL
    envFilter:
      filterTypes:
      - NON_PROD
`)

	yamlBytes := []byte(rawYaml)
	path := cac.GetArtifactServerYamlPath(name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(path, yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath("", path)
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, "", yamlItem.ApplicationId)
	require.Equal(t, path, yamlItem.Path)
	require.NotEmpty(t, yamlItem.Id)

	err = c.ConfigAsCodeClient.DeleteEntityV2(path, rawYaml)
	require.NoError(t, err)
}

func TestUpsertRawYaml_SourceRepoProvider(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: GIT
branch: master
keyAuth: false
password: gcpkms:TestAccResourceGitConnector_5JcVVWnvl83g_updated
providerType: GIT
url: https://github.com/micahlmartin/harness-demo
urlType: REPO
username: someuser
`)

	yamlBytes := []byte(rawYaml)
	path := cac.GetSourceRepoProviderYamlPath(name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(path, yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath("", path)
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, "", yamlItem.ApplicationId)
	require.Equal(t, path, yamlItem.Path)
	require.Equal(t, rawYaml, strings.TrimSpace(yamlItem.Content))
	require.NotEmpty(t, yamlItem.Id)

	err = c.ConfigAsCodeClient.DeleteEntityV2(path, rawYaml)
	require.NoError(t, err)
}

func TestUpsertRawYaml_GlobalTemplateLibrary(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	rawYaml := strings.TrimSpace(`
harnessApiVersion: '1.0'
type: SSH
commandUnitType: OTHER
commandUnits:
- command: some script
  commandUnitType: EXEC
  deploymentType: SSH
  name: Exec
  scriptType: BASH
- commandUnitType: COMMAND
  name: Global Command
  templateUri: Micah Testing Acount/Global Command:2
  variables:
  - name: appname
    value: foo
variables:
- name: appname
  value: foo
`)

	yamlBytes := []byte(rawYaml)
	rootPath, err := c.ConfigAsCodeClient.GetTemplateLibraryRootPathName()
	require.NoError(t, err)

	path := cac.GetTemplateLibraryYamlPath(rootPath, "subdirectory/nested/path", name)

	item, err := c.ConfigAsCodeClient.UpsertRawYaml(path, yamlBytes)
	require.NoError(t, err)
	require.NotNil(t, item)

	yamlItem, err := c.ConfigAsCodeClient.FindYamlByPath("", path)
	require.NoError(t, err)
	require.NotNil(t, yamlItem)
	require.Equal(t, "", yamlItem.ApplicationId)
	require.Equal(t, path, yamlItem.Path)
	require.Equal(t, rawYaml, strings.TrimSpace(yamlItem.Content))
	require.NotEmpty(t, yamlItem.Id)

	err = c.ConfigAsCodeClient.DeleteEntityV2(path, rawYaml)
	require.NoError(t, err)
}

func TestGetDirectoryTree(t *testing.T) {
	c := getClient()

	root, err := c.ConfigAsCodeClient.GetDirectoryTree("")

	require.NoError(t, err)
	fmt.Println(root)
}

func TestGetTemplateLibraryRootPath(t *testing.T) {
	c := getClient()
	path, err := c.ConfigAsCodeClient.GetTemplateLibraryRootPathName()
	require.NoError(t, err)
	require.NotEmpty(t, path)
}
