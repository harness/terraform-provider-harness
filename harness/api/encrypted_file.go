package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *SecretClient) GetEncryptedFileByName(name string) (*graphql.EncryptedFile, error) {
	resp := &graphql.EncryptedFile{}
	err := c.getSecretByName(name, graphql.SecretTypes.EncryptedFile, getEncryptedFileFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SecretClient) GetEncryptedFileById(id string) (*graphql.EncryptedFile, error) {
	resp := &graphql.EncryptedFile{}
	err := c.getSecretById(id, graphql.SecretTypes.EncryptedFile, getEncryptedFileFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getEncryptedFileFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on EncryptedFile {
			inheritScopesFromSM
			scopedToAccount
			secretManagerId
	}`, commonSecretFields)
}
