package nextgen

type GitConnectorType string

var GitConnectorTypes = struct {
	Account GitConnectorType
	Repo    GitConnectorType
}{
	Account: "Account",
	Repo:    "Repo",
}

var GitConnectorTypeValues = []string{
	GitConnectorTypes.Account.String(),
	GitConnectorTypes.Repo.String(),
}

func (e GitConnectorType) String() string {
	return string(e)
}
