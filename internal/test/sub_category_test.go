package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubCategoryField(t *testing.T) {
	resource_files, _ := ioutil.ReadDir(getAbsFilePath("../../docs/resources/"))
	data_source_files, _ := ioutil.ReadDir(getAbsFilePath("../../docs/data-sources/"))

	for _, er := range resource_files {
		file_name := er.Name()
		file_path := getAbsFilePath("../../docs/resources/" + file_name)
		read_file, err := ioutil.ReadFile(file_path)

		require.NoError(t, err)
		if !(strings.Contains(string(read_file), fmt.Sprint(`subcategory: "Next Gen"`)) || strings.Contains(string(read_file), fmt.Sprint(`subcategory: "First Gen"`))) {
			println("Sub category field is incorrectly populated in file " + file_name + " with file path " + file_path)
		}
		require.Equal(t, !(strings.Contains(string(read_file), fmt.Sprint(`subcategory: "Next Gen"`)) || strings.Contains(string(read_file), fmt.Sprint(`subcategory: "First Gen"`))), false)
	}

	for _, er := range data_source_files {
		file_name := er.Name()
		file_path := getAbsFilePath("../../docs/data-sources/" + file_name)
		read_file, err := ioutil.ReadFile(file_path)

		require.NoError(t, err)
		if !(strings.Contains(string(read_file), fmt.Sprint(`subcategory: "Next Gen"`)) || strings.Contains(string(read_file), fmt.Sprint(`subcategory: "First Gen"`))) {
			println("Sub category field is incorrectly populated in file " + file_name + " with file path " + file_path)
		}
		require.Equal(t, !(strings.Contains(string(read_file), fmt.Sprint(`subcategory: "Next Gen"`)) || strings.Contains(string(read_file), fmt.Sprint(`subcategory: "First Gen"`))), false)
	}
}

func getAbsFilePath(file_path string) string {
	absPath, _ := filepath.Abs(file_path)
	return absPath
}
