package nextgen

type AwsSdkClientBackOffStrategyType string

var AwsSdkClientBackOffStrategyTypes = struct {
	FixedDelayBackoffStrategy  AwsSdkClientBackOffStrategyType
	EqualJitterBackoffStrategy AwsSdkClientBackOffStrategyType
	FullJitterBackoffStrategy  AwsSdkClientBackOffStrategyType
}{
	FixedDelayBackoffStrategy:  "FixedDelayBackoffStrategy",
	EqualJitterBackoffStrategy: "EqualJitterBackoffStrategy",
	FullJitterBackoffStrategy:  "FullJitterBackoffStrategy",
}

func (e AwsSdkClientBackOffStrategyType) String() string {
	return string(e)
}
