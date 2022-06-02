package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type TriggerClient struct {
	ApiClient *ApiClient
}

func (c *TriggerClient) GetWebhookUrl(appId string, triggerName string) (*graphql.WebhookUrl, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			triggerByName(triggerName: "%[1]s", applicationId: "%[2]s")
			{
					condition{
						... on OnWebhook{
							webhookDetails{
								webhookURL
							}
						}
					}
			}
		}`, triggerName, appId),
	}

	res := &struct {
		TriggerByName struct {
			Condition struct {
				WebhookDetails struct {
					WebhookURL *graphql.WebhookUrl
				}
			}
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return res.TriggerByName.Condition.WebhookDetails.WebhookURL, nil
}
