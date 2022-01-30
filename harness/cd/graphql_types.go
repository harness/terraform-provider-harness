package cd

import (
	"encoding/json"
)

type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQLStandardResponse struct {
	Data             *json.RawMessage         `json:"data"`
	Metadata         interface{}              `json:"metadata"`
	Resource         string                   `json:"resource"`
	ResponseMessages []GraphQLResponseMessage `json:"responseMessages"`
	Errors           []GraphQLError           `json:"errors"`
}

type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int      `json:"line"`
		Column int      `json:"column"`
		Path   []string `json:"path"`
	} `json:"column"`
}

type GraphQLResponseMessage struct {
	Code         string   `json:"code"`
	Level        string   `json:"level"`
	Message      string   `json:"message"`
	Exception    string   `json:"exception"`
	FailureTypes []string `json:"failureTypes"`
}
