package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *NotificationRuleCondition) UnmarshalJSON(data []byte) error {

	type Alias NotificationRuleCondition

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case NotificationRuleConditionTypes.ErrorBudgetRemainingPercentage:
		err = json.Unmarshal(aux.Spec, &a.ErrorBudgetRemainingPercentage)
	case NotificationRuleConditionTypes.ErrorBudgetRemainingMinutes:
		err = json.Unmarshal(aux.Spec, &a.ErrorBudgetRemainingMinutes)
	case NotificationRuleConditionTypes.ErrorBudgetBurnRate:
		err = json.Unmarshal(aux.Spec, &a.ErrorBudgetBurnRate)
	case NotificationRuleConditionTypes.ChangeImpact:
		err = json.Unmarshal(aux.Spec, &a.ChangeImpact)
	case NotificationRuleConditionTypes.HealthScore:
		err = json.Unmarshal(aux.Spec, &a.HealthScore)
	case NotificationRuleConditionTypes.ChangeObserved:
		err = json.Unmarshal(aux.Spec, &a.ChangeObserved)
	case NotificationRuleConditionTypes.CodeErrors:
		err = json.Unmarshal(aux.Spec, &a.CodeErrors)
	case NotificationRuleConditionTypes.DeploymentImpactReport:
		err = json.Unmarshal(aux.Spec, &a.DeploymentImpactReport)
	default:
		panic(fmt.Sprintf("unknown notification rule condition type type %s", a.Type_))
	}

	return err
}

func (a *NotificationRuleCondition) MarshalJSON() ([]byte, error) {
	type Alias NotificationRuleCondition

	var spec []byte
	var err error

	switch a.Type_ {
	case NotificationRuleConditionTypes.ErrorBudgetRemainingPercentage:
		spec, err = json.Marshal(a.ErrorBudgetRemainingPercentage)
	case NotificationRuleConditionTypes.ErrorBudgetRemainingMinutes:
		spec, err = json.Marshal(a.ErrorBudgetRemainingMinutes)
	case NotificationRuleConditionTypes.ErrorBudgetBurnRate:
		spec, err = json.Marshal(a.ErrorBudgetBurnRate)
	case NotificationRuleConditionTypes.ChangeImpact:
		spec, err = json.Marshal(a.ChangeImpact)
	case NotificationRuleConditionTypes.HealthScore:
		spec, err = json.Marshal(a.HealthScore)
	case NotificationRuleConditionTypes.ChangeObserved:
		spec, err = json.Marshal(a.ChangeObserved)
	case NotificationRuleConditionTypes.CodeErrors:
		spec, err = json.Marshal(a.CodeErrors)
	case NotificationRuleConditionTypes.DeploymentImpactReport:
		spec, err = json.Marshal(a.DeploymentImpactReport)
	default:
		panic(fmt.Sprintf("unknown notification rule condition type type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
