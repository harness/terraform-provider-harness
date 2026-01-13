package idp

import (
	"encoding/json"
	"strings"
)

type IDPErrorBody struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func toIDPErrBody(err error) (*IDPErrorBody, error) {
	var body IDPErrorBody
	apiErr, ok := err.(interface{ Body() []byte })
	if !ok {
		return nil, err
	}

	newErr := json.Unmarshal(apiErr.Body(), &body)
	if newErr != nil {
		return nil, newErr
	}

	return &body, nil
}

func isOnlyVersionDeleteError(err error) bool {
	body, newErr := toIDPErrBody(err)
	if newErr != nil {
		return false
	}

	msg := strings.ToLower(body.Message)
	return strings.Contains(msg, "cannot delete version") &&
		strings.Contains(msg, "at least one version must exist")
}

func isNotFoundError(err error) bool {
	body, newErr := toIDPErrBody(err)
	if newErr != nil {
		return false
	}

	msg := strings.ToLower(body.Message)
	return strings.Contains(msg, "not found")
}
