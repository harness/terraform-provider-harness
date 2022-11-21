package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type TriggerClient struct {
	ApiClient *ApiClient
}

func (c *TriggerClient) GetTriggerById(triggerId string) (*graphql.Trigger, error) {
	queryToGetTriggerConditionType := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			trigger(triggerId: "%[1]s"){
				description
				id
				name
				
				condition{
					triggerConditionType
				}
			}
		}`, triggerId),
	}

	resConditionType := &struct {
		Trigger graphql.Trigger
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(queryToGetTriggerConditionType, &resConditionType)

	if err != nil {
		return nil, err
	}

	if resConditionType.Trigger.Condition.TriggerConditionType == "WEBHOOK" {
		query := &GraphQLQuery{
			Query: fmt.Sprintf(`{
			trigger(triggerId: "%[1]s"){
				description
				id
				name
				
				condition{
					%[2]s
				}
			}
		}`, triggerId, getCondition("WEBHOOK")),
		}

		res := &struct {
			Trigger graphql.Trigger
		}{}

		err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

		if err != nil {
			return nil, err
		}

		return &res.Trigger, nil
	}
	return nil, nil
}

func (c *TriggerClient) GetTriggerByName(triggerName string, appId string) (*graphql.Trigger, error) {
	queryToGetTriggerConditionType := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			triggerByName(triggerName: "%[1]s",applicationId: "%[2]s"){
				description
				id
				name
				
				condition{
					triggerConditionType
				}
			}
		}`, triggerName, appId),
	}

	resConditionType := &struct {
		TriggerByName struct {
			Condition struct {
				TriggerConditionType string
			}
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(queryToGetTriggerConditionType, &resConditionType)

	if err != nil {
		return nil, err
	}

	if resConditionType.TriggerByName.Condition.TriggerConditionType == "WEBHOOK" {
		query := &GraphQLQuery{
			Query: fmt.Sprintf(`{
			triggerByName(triggerName: "%[1]s",applicationId: "%[2]s"){
				description
				id
				name
				
				condition{
					%[3]s
				}
			}
		}`, triggerName, appId, getCondition("WEBHOOK")),
		}

		res := &struct {
			TriggerByName graphql.Trigger
		}{}

		err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

		if err != nil {
			return nil, err
		}

		return &res.TriggerByName, nil
	}
	return nil, nil
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

func getCondition(triggerConditionType string) string {
	if triggerConditionType == "WEBHOOK" {
		return `
		... on OnWebhook{
			triggerConditionType
			webhookDetails{
				payload
				webhookURL
				header
				method
			}
	}`
	}
	return ""
}
