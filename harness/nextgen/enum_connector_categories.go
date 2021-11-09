package nextgen

type ConnectorCategory string

var ConnectorCategories = struct {
	CloudProvider ConnectorCategory
	SecretManager ConnectorCategory
	CloudCost     ConnectorCategory
	Artifactory   ConnectorCategory
	CodeRepo      ConnectorCategory
	Monitoring    ConnectorCategory
	Ticketing     ConnectorCategory
}{
	CloudProvider: "CLOUD_PROVIDER",
	SecretManager: "SECRET_MANAGER",
	CloudCost:     "CLOUD_COST",
	Artifactory:   "ARTIFACTORY",
	CodeRepo:      "CODE_REPO",
	Monitoring:    "MONITORING",
	Ticketing:     "TICKETING",
}

func (c ConnectorCategory) String() string {
	return string(c)
}
