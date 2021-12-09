package cd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/cd/unpublished"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
)

func (c *SecretClient) GetWinRMCredentialById(id string) (*graphql.WinRMCredential, error) {

	resp := &graphql.WinRMCredential{}
	err := c.getSecretById(id, graphql.SecretTypes.WinRMCredential, getWinRMCredentialFields(), resp)
	if err != nil {
		return nil, err
	}

	if resp.IsEmpty() {
		return nil, nil
	}

	return resp, nil
}

func (c *SecretClient) GetWinRMCredentialByName(name string) (*graphql.WinRMCredential, error) {

	resp := &graphql.WinRMCredential{}
	err := c.getSecretByName(name, graphql.SecretTypes.WinRMCredential, getWinRMCredentialFields(), resp)
	if err != nil {
		return nil, err
	}

	if resp.IsEmpty() {
		return nil, nil
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

func (c *SecretClient) ListWinRMCredentials() ([]*unpublished.Credential, error) {
	req, err := c.ApiClient.NewAuthorizedGetRequest("/secrets/list-values")
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)
	req.URL.RawQuery = query.Encode()

	resp, err := c.ApiClient.Configuration.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	log.Printf("[DEBUG] GraphQL response: %s", buf.String())

	responsePackage := &unpublished.Package{}
	err = json.NewDecoder(resp.Body).Decode(responsePackage)

	if err != nil {
		return nil, err
	}

	credentials := []*unpublished.Credential{}
	err = json.Unmarshal(*responsePackage.Resource, &credentials)
	if err != nil {
		return nil, err
	}

	winrmCreds := []*unpublished.Credential{}
	for _, cred := range credentials {

		if cred.Value.Type == "WINRM_CONNECTION_ATTRIBUTES" {
			winrmCreds = append(winrmCreds, cred)
		}
	}

	return winrmCreds, nil
}
