package graphql

type DeploymentTagType string

var DeploymentTagTypes = struct {
	Application, Service, Environment, Deployment DeploymentTagType
}{
	Application: "APPLICATION",
	Service:     "SERVICE",
	Environment: "ENVIRONMENT",
	Deployment:  "DEPLOYMENT",
}
