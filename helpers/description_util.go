package helpers

type Desc string

var Descriptions = struct {
	ConnectorRefText Desc
}{
	ConnectorRefText: " To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.",
}

var DescriptionValues = []string{
	Descriptions.ConnectorRefText.String(),
}

func (e Desc) String() string {
	return string(e)
}
