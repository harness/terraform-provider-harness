package nextgen

type JDBCAuthType string

var JDBCAuthTypes = struct {
	UsernamePassword JDBCAuthType
}{
	UsernamePassword: "UsernamePassword",
}

func (e JDBCAuthType) String() string {
	return string(e)
}
