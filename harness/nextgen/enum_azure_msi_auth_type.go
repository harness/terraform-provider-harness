package nextgen

type AzureMsiAuthType string

var AzureMsiAuthTypes = struct {
	SystemAssignedManagedIdentity AzureMsiAuthType
	UserAssignedManagedIdentity   AzureMsiAuthType
}{
	SystemAssignedManagedIdentity: "SystemAssignedManagedIdentity",
	UserAssignedManagedIdentity:   "UserAssignedManagedIdentity",
}

var AzureMsiAuthTypeValues = []string{
	AzureMsiAuthTypes.SystemAssignedManagedIdentity.String(),
	AzureMsiAuthTypes.SystemAssignedManagedIdentity.String(),
}

func (e AzureMsiAuthType) String() string {
	return string(e)
}
