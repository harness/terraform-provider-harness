package nextgen

type GithubApiAccessType string

var GithubApiAccessTypes = struct {
	Token     GithubApiAccessType
	GithubApp GithubApiAccessType
}{
	Token:     "Token",
	GithubApp: "GithubApp",
}

var GithubApiAccessTypeValues = []string{
	GithubApiAccessTypes.Token.String(),
	GithubApiAccessTypes.GithubApp.String(),
}

func (e GithubApiAccessType) String() string {
	return string(e)
}
