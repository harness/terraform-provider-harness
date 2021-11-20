package nextgen

type BitBucketApiAccessType string

var BitBucketApiAccessTypes = struct {
	UsernameToken BitBucketApiAccessType
}{
	UsernameToken: "UsernameToken",
}

var BitBucketApiAccessTypeValues = []string{
	BitBucketApiAccessTypes.UsernameToken.String(),
}

func (e BitBucketApiAccessType) String() string {
	return string(e)
}
