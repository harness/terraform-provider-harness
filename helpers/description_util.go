package helpers

type Desc string

var Descriptions = struct {
	ConnectorRefText Desc
	YamlText         Desc
}{
	ConnectorRefText: " To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.",
	YamlText:         " In YAML, to reference an entity at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference an entity at the account scope, prefix 'account` to the expression: account.{identifier}. For eg, to reference a connector with identifier 'connectorId' at the organization scope in a stage mention it as connectorRef: org.connectorId.",
}

var DescriptionValues = []string{
	Descriptions.ConnectorRefText.String(),
	Descriptions.YamlText.String(),
}

func (e Desc) String() string {
	return string(e)
}
