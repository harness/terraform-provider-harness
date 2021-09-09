package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestFindConfigAsCodeItemByUUID(t *testing.T) {
// 	c := getClient()

// 	root, err := c.ConfigAsCode().GetDirectoryTree("IDrB3WsxQWmvLeFGYstqLA")
// 	require.NoError(t, err)
// 	item := FindConfigAsCodeItemByUUID(root, "Fh3yqtNrQyilc3daEh_o1g")
// 	require.NoError(t, err)
// 	fmt.Println(item)
// }

// func TestFindObjectById(t *testing.T) {
// 	c := getClient()

// 	// root, err := c.ConfigAsCode().GetDirectoryTree("")
// 	// require.NoError(t, err)
// 	cp := &cac.AwsCloudProvider{}
// 	err := c.ConfigAsCode().FindObjectById("", "AlE6mvkqQ2m07n_Ofls0SQ", cp)
// 	require.NoError(t, err)
// 	// fmt.Println(item)
// }

func TestGetDirectoryTree(t *testing.T) {
	c := getClient()

	root, err := c.ConfigAsCode().GetDirectoryTree("")

	require.NoError(t, err)
	fmt.Println(root)
}

// func FindConfigAsCodeItemByUUID(rootItem *cac.ConfigAsCodeItem, uuid string) *cac.ConfigAsCodeItem {
// 	if rootItem.DirectoryPath != nil && rootItem.UUID == uuid {
// 		return rootItem
// 	}

// 	for _, item := range rootItem.Children {
// 		if matchingItem := FindConfigAsCodeItemByUUID(item, uuid); matchingItem != nil {
// 			return matchingItem
// 		}
// 	}

// 	return nil
// }
