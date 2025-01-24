package nextgen

type JDBCAuthType string

var JDBCAuthTypes = struct {
	UsernamePassword JDBCAuthType
	ServiceAccount   JDBCAuthType
}{
	UsernamePassword: "UsernamePassword",
	ServiceAccount:   "ServiceAccount",
}

func (e JDBCAuthType) String() string {
	return string(e)
}
