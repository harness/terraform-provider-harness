package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *SecretClient) GetWinRMCredentialById(id string) (*graphql.WinRMCredential, error) {

	resp := &graphql.WinRMCredential{}
	err := c.getSecretById(id, graphql.SecretTypes.WinRMCredential, getWinRMCredentialFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SecretClient) GetWinRMCredentialByName(name string) (*graphql.WinRMCredential, error) {

	resp := &graphql.WinRMCredential{}
	err := c.getSecretByName(name, graphql.SecretTypes.WinRMCredential, getWinRMCredentialFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getWinRMCredentialFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on WinRMCredential {
			authenticationScheme
			domain
			port
			skipCertCheck
			useSSL
			userName
		}
	`, commonSecretFields)
}
