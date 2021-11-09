package cac

import (
	"fmt"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestGetEntityNameFromPath_Index(t *testing.T) {
	yamlPath := "Setup/Applications/myapplication/services/myservice/Index.yaml"

	fmt.Println(yamlPath)

	fmt.Println(path.Dir(yamlPath))
	fmt.Println(path.Split(yamlPath))

	name := GetEntityNameFromPath(YamlPath(yamlPath))
	require.Equal(t, "myservice", name)
}

func TestGetEntityNameFromPath_NameInFile(t *testing.T) {
	yamlPath := "Setup/Cloud Providers/google-cloud-provider.yaml"

	fmt.Println(yamlPath)

	fmt.Println(path.Dir(yamlPath))
	fmt.Println(path.Split(yamlPath))

	name := GetEntityNameFromPath(YamlPath(yamlPath))
	require.Equal(t, "google-cloud-provider", name)
}

func TestSecretRefUnmarshalYaml(t *testing.T) {
	yamlString := `gcpkms:secretname`

	secretRef := SecretRef{}
	err := yaml.Unmarshal([]byte(yamlString), &secretRef)

	require.NoError(t, err)
	require.Equal(t, "secretname", secretRef.Name)
	require.Equal(t, SecretManagerTypes.GcpKMS, secretRef.SecretManagerType)

	yamlString = `secretname`

	secretRef = SecretRef{}
	err = yaml.Unmarshal([]byte(yamlString), &secretRef)

	require.NoError(t, err)
	require.Equal(t, "secretname", secretRef.Name)
	require.Equal(t, SecretManagerType(""), secretRef.SecretManagerType)
}
