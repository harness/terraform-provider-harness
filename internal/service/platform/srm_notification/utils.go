package srm_notification

import (
	"encoding/json"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func getNotificationRuleConditionByType(hs map[string]interface{}) nextgen.NotificationRuleCondition {
	notificationRuleConditionType := hs["type"].(string)
	notificationRuleConditionSpec := hs["spec"].(string)

	if notificationRuleConditionType == "ErrorBudgetRemainingPercentage" {
		data := nextgen.ErrorBudgetRemainingPercentageConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:                          nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			ErrorBudgetRemainingPercentage: &data,
		}
	}
	if notificationRuleConditionType == "ErrorBudgetRemainingMinutes" {
		data := nextgen.ErrorBudgetRemainingMinutesConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:                       nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			ErrorBudgetRemainingMinutes: &data,
		}
	}
	if notificationRuleConditionType == "ErrorBudgetBurnRate" {
		data := nextgen.ErrorBudgetBurnRateConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:               nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			ErrorBudgetBurnRate: &data,
		}
	}
	if notificationRuleConditionType == "ChangeImpact" {
		data := nextgen.ChangeImpactConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:        nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			ChangeImpact: &data,
		}
	}
	if notificationRuleConditionType == "HealthScore" {
		data := nextgen.HealthScoreConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:       nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			HealthScore: &data,
		}
	}
	if notificationRuleConditionType == "ChangeObserved" {
		data := nextgen.ChangeObservedConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:          nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			ChangeObserved: &data,
		}
	}
	if notificationRuleConditionType == "CodeErrors" {
		data := nextgen.ErrorTrackingConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:      nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			CodeErrors: &data,
		}
	}
	if notificationRuleConditionType == "DeploymentImpactReport" {
		data := nextgen.DeploymentImpactReportConditionSpec{}
		json.Unmarshal([]byte(notificationRuleConditionSpec), &data)

		return nextgen.NotificationRuleCondition{
			Type_:                  nextgen.NotificationRuleConditionType(notificationRuleConditionType),
			DeploymentImpactReport: &data,
		}
	}
	panic(fmt.Sprintf("Invalid notification rule conditions for srm notification"))
}

func getNotificationChannelByType(hs map[string]interface{}) nextgen.CvngNotificationChannel {
	notificationChannelType := hs["type"].(string)
	notificationChannelSpec := hs["spec"].(string)

	if notificationChannelType == "Email" {
		data := nextgen.CvngEmailChannelSpec{}
		json.Unmarshal([]byte(notificationChannelSpec), &data)

		return nextgen.CvngNotificationChannel{
			Type_: nextgen.CVNGNotificationChannelType(notificationChannelType),
			Email: &data,
		}
	}
	if notificationChannelType == "Slack" {
		data := nextgen.CvngSlackChannelSpec{}
		json.Unmarshal([]byte(notificationChannelSpec), &data)

		return nextgen.CvngNotificationChannel{
			Type_: nextgen.CVNGNotificationChannelType(notificationChannelType),
			Slack: &data,
		}
	}
	if notificationChannelType == "PagerDuty" {
		data := nextgen.CvngPagerDutyChannelSpec{}
		json.Unmarshal([]byte(notificationChannelSpec), &data)

		return nextgen.CvngNotificationChannel{
			Type_:     nextgen.CVNGNotificationChannelType(notificationChannelType),
			PagerDuty: &data,
		}
	}
	if notificationChannelType == "MsTeams" {
		data := nextgen.CvngMsTeamsChannelSpec{}
		json.Unmarshal([]byte(notificationChannelSpec), &data)

		return nextgen.CvngNotificationChannel{
			Type_:   nextgen.CVNGNotificationChannelType(notificationChannelType),
			MsTeams: &data,
		}
	}
	panic(fmt.Sprintf("Invalid notification channel for srm notification"))
}
