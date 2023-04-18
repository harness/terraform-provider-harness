package nextgen

type TerraformCloudAuthType string

var TerraformCloudAuthTypes = struct {
	ApiToken TerraformCloudAuthType
}{
	ApiToken: "ApiToken",
}

func (e TerraformCloudAuthType) String() string {
	return string(e)
}
