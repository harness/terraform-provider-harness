package nextgen

type NotificationRuleConditionType string

var NotificationRuleConditionTypes = struct {
	ErrorBudgetRemainingPercentage NotificationRuleConditionType
	ErrorBudgetRemainingMinutes    NotificationRuleConditionType
	ErrorBudgetBurnRate            NotificationRuleConditionType
	ChangeImpact                   NotificationRuleConditionType
	HealthScore                    NotificationRuleConditionType
	ChangeObserved                 NotificationRuleConditionType
	CodeErrors                     NotificationRuleConditionType
	DeploymentImpactReport         NotificationRuleConditionType
}{
	ErrorBudgetRemainingPercentage: "ErrorBudgetRemainingPercentage",
	ErrorBudgetRemainingMinutes:    "ErrorBudgetRemainingMinutes",
	ErrorBudgetBurnRate:            "ErrorBudgetBurnRate",
	ChangeImpact:                   "ChangeImpact",
	HealthScore:                    "HealthScore",
	ChangeObserved:                 "ChangeObserved",
	CodeErrors:                     "CodeErrors",
	DeploymentImpactReport:         "DeploymentImpactReport",
}

var NotificationRuleConditionTypesSlice = []string{
	NotificationRuleConditionTypes.ErrorBudgetRemainingPercentage.String(),
	NotificationRuleConditionTypes.ErrorBudgetRemainingMinutes.String(),
	NotificationRuleConditionTypes.ErrorBudgetBurnRate.String(),
	NotificationRuleConditionTypes.ChangeImpact.String(),
	NotificationRuleConditionTypes.HealthScore.String(),
	NotificationRuleConditionTypes.ChangeObserved.String(),
	NotificationRuleConditionTypes.CodeErrors.String(),
	NotificationRuleConditionTypes.DeploymentImpactReport.String(),
}

func (c NotificationRuleConditionType) String() string {
	return string(c)
}
