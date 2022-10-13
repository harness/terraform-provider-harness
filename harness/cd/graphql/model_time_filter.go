package graphql

import "github.com/harness/harness-go-sdk/harness/time"

type TimeFilter struct {
	Operator TimeOperatorType
	Value    *time.Time
}
