package nextgen

type SpotAuthType string

var SpotAuthTypes = struct {
	PermanentTokenConfig SpotAuthType
}{
	PermanentTokenConfig: "PermanentTokenConfig",
}

func (e SpotAuthType) String() string {
	return string(e)
}
