package nextgen

type ConnectorFilterType string

var ConnectorFilterTypes = struct {
	Connector ConnectorFilterType
}{
	Connector: "Connector",
}

var ConnectorFilterTypeValues = []string{
	ConnectorFilterTypes.Connector.String(),
}

func (e ConnectorFilterType) String() string {
	return string(e)
}
