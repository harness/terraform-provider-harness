package graphql

type ArtifactType string

var ArtifactTypes = struct {
	AMI                 ArtifactType
	AWSCodeDeploy       ArtifactType
	AWSLambda           ArtifactType
	Docker              ArtifactType
	Jar                 ArtifactType
	Other               ArtifactType
	PCF                 ArtifactType
	RPM                 ArtifactType
	Tar                 ArtifactType
	War                 ArtifactType
	IISVirtualDirectory ArtifactType
	IISApp              ArtifactType
	IISWebsite          ArtifactType
	Zip                 ArtifactType
}{
	AMI:                 "AMI",
	AWSCodeDeploy:       "AWS_CODEDEPLOY",
	AWSLambda:           "AWS_LAMBDA",
	Docker:              "DOCKER",
	Jar:                 "JAR",
	Other:               "OTHER",
	PCF:                 "PCF",
	RPM:                 "RPM",
	Tar:                 "TAR",
	War:                 "WAR",
	IISVirtualDirectory: "IIS_VirtualDirectory",
	IISApp:              "IIS_APP",
	IISWebsite:          "IIS",
	Zip:                 "ZIP",
}

var ArtifactTypeValues = []string{
	ArtifactTypes.AMI.String(),
	ArtifactTypes.AWSCodeDeploy.String(),
	ArtifactTypes.AWSLambda.String(),
	ArtifactTypes.Docker.String(),
	ArtifactTypes.Jar.String(),
	ArtifactTypes.Other.String(),
	ArtifactTypes.PCF.String(),
	ArtifactTypes.RPM.String(),
	ArtifactTypes.Tar.String(),
	ArtifactTypes.War.String(),
	ArtifactTypes.IISVirtualDirectory.String(),
	ArtifactTypes.IISApp.String(),
	ArtifactTypes.IISWebsite.String(),
	ArtifactTypes.Zip.String(),
}

func (e ArtifactType) String() string {
	return string(e)
}
