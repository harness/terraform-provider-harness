package cac

import (
	"fmt"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
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
