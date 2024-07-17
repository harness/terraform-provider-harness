package notification_rule

import (
	"encoding/json"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func getNotificationRuleConditionByType(nrc map[string]interface{}) (nextgen.NotificationRuleCondition, error) {
	notifRuleConditionType := nrc["type"].(string)
	notifRuleCondition := nrc["spec"].(string)
	var err error
	if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.ErrorBudgetRemainingPercentage.String() {
		data := &nextgen.ErrorBudgetRemainingPercentageConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			ErrorBudgetRemainingPercentage: data,
			Type_:                          nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.ErrorBudgetRemainingMinutes.String() {
		data := nextgen.ErrorBudgetRemainingMinutesConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			ErrorBudgetRemainingMinutes: &data,
			Type_:                       nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.ErrorBudgetBurnRate.String() {
		data := nextgen.ErrorBudgetBurnRateConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			ErrorBudgetBurnRate: &data,
			Type_:               nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.ChangeImpact.String() {
		data := nextgen.ChangeImpactConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			ChangeImpact: &data,
			Type_:        nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.ChangeObserved.String() {
		data := nextgen.ChangeObservedConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			ChangeObserved: &data,
			Type_:          nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.DeploymentImpactReport.String() {
		data := nextgen.DeploymentImpactReportConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			DeploymentImpactReport: &data,
			Type_:                  nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	} else if notifRuleConditionType == nextgen.NotificationRuleConditionTypes.HealthScore.String() {
		data := nextgen.HealthScoreConditionSpec{}
		err = json.Unmarshal([]byte(notifRuleCondition), &data)
		if err != nil {
			return nextgen.NotificationRuleCondition{}, err
		}

		return nextgen.NotificationRuleCondition{
			HealthScore: &data,
			Type_:       nextgen.NotificationRuleConditionType(notifRuleConditionType),
		}, nil
	}

	return nextgen.NotificationRuleCondition{}, fmt.Errorf("Notification rule condition not yet supported/invalid")
}

func getNotificationChannelByType(nc map[string]interface{}) (nextgen.CvngNotificationChannel, error) {
	notifChannelType := nc["type"].(string)
	notifChannel := nc["spec"].(string)
	var err error

	if notifChannelType == nextgen.CVNGNotificationChannelTypes.Email.String() {
		data := nextgen.CvngEmailChannelSpec{}
		err = json.Unmarshal([]byte(notifChannel), &data)
		if err != nil {
			return nextgen.CvngNotificationChannel{}, err
		}

		return nextgen.CvngNotificationChannel{
			Type_: nextgen.CVNGNotificationChannelType(notifChannelType),
			Email: &data,
		}, nil
	} else if notifChannelType == nextgen.CVNGNotificationChannelTypes.Slack.String() {
		data := nextgen.CvngSlackChannelSpec{}
		err = json.Unmarshal([]byte(notifChannel), &data)
		if err != nil {
			return nextgen.CvngNotificationChannel{}, err
		}

		return nextgen.CvngNotificationChannel{
			Type_: nextgen.CVNGNotificationChannelType(notifChannelType),
			Slack: &data,
		}, nil
	} else if notifChannelType == nextgen.CVNGNotificationChannelTypes.PagerDuty.String() {
		data := nextgen.CvngPagerDutyChannelSpec{}
		err = json.Unmarshal([]byte(notifChannel), &data)
		if err != nil {
			return nextgen.CvngNotificationChannel{}, err
		}

		return nextgen.CvngNotificationChannel{
			Type_:     nextgen.CVNGNotificationChannelType(notifChannelType),
			PagerDuty: &data,
		}, nil
	} else if notifChannelType == nextgen.CVNGNotificationChannelTypes.MsTeams.String() {
		data := nextgen.CvngMsTeamsChannelSpec{}
		err = json.Unmarshal([]byte(notifChannel), &data)
		if err != nil {
			return nextgen.CvngNotificationChannel{}, err
		}

		return nextgen.CvngNotificationChannel{
			Type_:   nextgen.CVNGNotificationChannelType(notifChannelType),
			MsTeams: &data,
		}, nil
	}

	return nextgen.CvngNotificationChannel{}, fmt.Errorf("Notification channel not yet supported/invalid")
}
