package nextgen

type BitBucketHttpCredentialType string

var BitBucketHttpCredentialTypes = struct {
	UsernamePassword BitBucketHttpCredentialType
}{
	UsernamePassword: "UsernamePassword",
}

var BitBucketHttpCredentialTypeValues = []string{
	BitBucketHttpCredentialTypes.UsernamePassword.String(),
}

func (e BitBucketHttpCredentialType) String() string {
	return string(e)
}
