package graphql

type ArtifactInputType string

var ArtifactInputTypes = struct {
	ArtifactId                  ArtifactInputType
	BuildNumber                 ArtifactInputType
	ParameterizedArtifactSource ArtifactInputType
}{
	ArtifactId:                  "ARTIFACT_ID",
	BuildNumber:                 "BUILD_NUMBER",
	ParameterizedArtifactSource: "PARAMETERIZED_ARTIFACT_SOURCE",
}

var ArtifactInputTypeList = []string{
	ArtifactInputTypes.ArtifactId.String(),
	ArtifactInputTypes.BuildNumber.String(),
	ArtifactInputTypes.ParameterizedArtifactSource.String(),
}

func (d ArtifactInputType) String() string {
	return string(d)
}
