package internal

import (
	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

type Session struct {
	AccountId string
	Endpoint  string
	CDClient  *cd.ApiClient
	PLClient  *nextgen.APIClient
}
