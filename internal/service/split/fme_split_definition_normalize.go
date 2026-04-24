package split

import (
	"encoding/json"
	"reflect"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
)

// splitDefinitionMergePresentationFromPrior copies title and comment from prior JSON onto apiStr when Split's
// GET-derived JSON omits them (SplitDefinition has no title/comment fields; see splitDefinitionRequestFromAPI).
func splitDefinitionMergePresentationFromPrior(apiStr, prior string) (string, error) {
	if prior == "" {
		return apiStr, nil
	}
	var apiReq splitsdk.SplitDefinitionRequest
	if err := json.Unmarshal([]byte(apiStr), &apiReq); err != nil {
		return "", err
	}
	var p splitsdk.SplitDefinitionRequest
	if err := json.Unmarshal([]byte(prior), &p); err != nil {
		return apiStr, nil
	}
	if apiReq.Title == "" && p.Title != "" {
		apiReq.Title = p.Title
	}
	if apiReq.Comment == "" && p.Comment != "" {
		apiReq.Comment = p.Comment
	}
	b, err := json.Marshal(apiReq)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// splitDefinitionRequestFromAPI maps a full Split API definition to the request shape used for create/update.
func splitDefinitionRequestFromAPI(def *splitsdk.SplitDefinition) splitsdk.SplitDefinitionRequest {
	if def == nil {
		return splitsdk.SplitDefinitionRequest{}
	}
	return splitsdk.SplitDefinitionRequest{
		Treatments:        def.Treatments,
		Rules:             def.Rules,
		DefaultRule:       def.DefaultRule,
		DefaultTreatment:  def.DefaultTreatment,
		BaselineTreatment: def.BaselineTreatment,
		TrafficAllocation: def.TrafficAllocation,
	}
}

// splitDefinitionRequestJSONForState returns stable JSON for the definition attribute in Terraform state.
func splitDefinitionRequestJSONForState(def *splitsdk.SplitDefinition) (string, error) {
	req := splitDefinitionRequestFromAPI(def)
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// splitDefinitionJSONSemanticallyEqual is true when both strings unmarshal to equivalent SplitDefinitionRequest values.
// Used so HCL jsonencode key order differs from json.Marshal field order without causing perpetual drift.
func splitDefinitionJSONSemanticallyEqual(a, b string) bool {
	var ar, br splitsdk.SplitDefinitionRequest
	if err := json.Unmarshal([]byte(a), &ar); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &br); err != nil {
		return false
	}
	normalizeSplitDefReq(&ar)
	normalizeSplitDefReq(&br)
	return reflect.DeepEqual(ar, br)
}

func normalizeSplitDefReq(r *splitsdk.SplitDefinitionRequest) {
	if r.Treatments == nil {
		r.Treatments = []splitsdk.Treatment{}
	}
	if r.Rules == nil {
		r.Rules = []splitsdk.Rule{}
	}
	if r.DefaultRule == nil {
		r.DefaultRule = []splitsdk.Bucket{}
	}
}
