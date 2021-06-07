package client

import "fmt"

const encryptedSecretFields = `
				id
				name
				secretManagerId
				secretType
				usageScope {
					appEnvScopes {
						application {
							filterType
							appId
						}
						environment {
							filterType
							envId
						}
					}
				}
				inheritScopesFromSM
				scopedToAccount
`

func (c *SecretClient) GetEncryptedSecretByName(name string) (*EncryptedText, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on EncryptedText {
					%s
				}
			}
		}`, SecretTypes.EncryptedText, encryptedSecretFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName EncryptedText
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}
