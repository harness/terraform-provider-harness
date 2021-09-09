package cac

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v3"
)

func (i *ConfigAsCodeItem) ParseYamlContent(respObj interface{}) error {
	if i.Yaml == "" {
		return nil
	}

	if err := yaml.Unmarshal([]byte(i.Yaml), respObj); err != nil {
		return err
	}

	return nil
}

func (i *ConfigAsCodeItem) IsEmpty() bool {
	return i == &ConfigAsCodeItem{}
}

// Indicates an error condition
func (r *Response) IsEmpty() bool {
	// return true
	return r.Metadata == ResponseMetadata{} && r.Resource.IsEmpty() && len(r.ResponseMessages) == 0
}

func (m *ResponseMessage) ToError() error {
	return fmt.Errorf("%s: %s", m.Code, m.Message)
}

func GetDefaultArtifactType(deploymentType DeploymentType, fallbackArtifactType ArtifactType) (ArtifactType, error) {

	var artifactType ArtifactType

	switch deploymentType {
	case DeploymentTypes.Kubernetes:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.SSH:
		artifactType = fallbackArtifactType
	case DeploymentTypes.AMI:
		artifactType = ArtifactTypes.AMI
	case DeploymentTypes.AWSCodeDeploy:
		artifactType = ArtifactTypes.AWSCodeDeploy
	case DeploymentTypes.AWSLambda:
		artifactType = ArtifactTypes.AWSLambda
	case DeploymentTypes.ECS:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.PCF:
		artifactType = ArtifactTypes.PCF
	case DeploymentTypes.Helm:
		artifactType = ArtifactTypes.Docker
	case DeploymentTypes.WinRM:
		artifactType = fallbackArtifactType
	default:
		return "", fmt.Errorf("no default artifact type for '%s' deployments", deploymentType)
	}

	return artifactType, nil
}

func NewEntity(objectType ObjectType) interface{} {
	t, ok := objectTypeMap[objectType]
	if !ok {
		panic(fmt.Errorf("could not create entity of type `%s`", objectType))
	}

	i := reflect.New(t).Interface()
	utils.MustSetField(i, "HarnessApiVersion", HarnessApiVersions.V1)
	utils.MustSetField(i, "Type", objectType)
	return i
}

var objectTypeMap = map[ObjectType]reflect.Type{
	ObjectTypes.Application:                     reflect.TypeOf(Application{}),
	ObjectTypes.Application:                     reflect.TypeOf(Application{}),
	ObjectTypes.AwsCloudProvider:                reflect.TypeOf(AwsCloudProvider{}),
	ObjectTypes.AzureCloudProvider:              reflect.TypeOf(AzureCloudProvider{}),
	ObjectTypes.Environment:                     reflect.TypeOf(Environment{}),
	ObjectTypes.GcpCloudProvider:                reflect.TypeOf(GcpCloudProvider{}),
	ObjectTypes.InfrastructureDefinition:        reflect.TypeOf(InfrastructureDefinition{}),
	ObjectTypes.KubernetesCloudProvider:         reflect.TypeOf(KubernetesCloudProvider{}),
	ObjectTypes.PcfCloudProvider:                reflect.TypeOf(PcfCloudProvider{}),
	ObjectTypes.PhysicalDataCenterCloudProvider: reflect.TypeOf(PhysicalDatacenterCloudProvider{}),
	ObjectTypes.Service:                         reflect.TypeOf(Service{}),
	ObjectTypes.SpotInstCloudProvider:           reflect.TypeOf(SpotInstCloudProvider{}),
}

func (r *SecretRef) MarshalYAML() (interface{}, error) {
	if (r == &SecretRef{}) {
		return []byte{}, nil
	}

	if r.Name == "" {
		return nil, errors.New("name must be set")
	}

	if r.SecretManagerType == "" {
		return r.Name, nil
	}

	return fmt.Sprintf("%s:%s", r.SecretManagerType, r.Name), nil
}

func (r *SecretRef) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var val interface{}
	err := unmarshal(&val)
	if err != nil {
		return err
	}

	value := val.(string)

	parts := strings.Split(value, ":")

	if len(parts) == 1 {
		r.Name = parts[0]
	} else if len(parts) == 2 {
		r.SecretManagerType = SecretManagerType(parts[0])
		r.Name = parts[1]
	}

	return nil
}

func GetEntityNameFromPath(yamlPath YamlPath) string {
	dir, file := path.Split(string(yamlPath))

	if ok, _ := regexp.MatchString("Index.yaml", file); ok {
		parts := strings.Split(strings.TrimSpace(dir), "/")
		last := parts[len(parts)-2]
		return last
	}

	return utils.TrimFileExtension(file)
}

func GetServiceYamlPath(applicationName string, serviceName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", applicationName, serviceName))
}

func GetCloudProviderYamlPath(cloudProviderName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Cloud Providers/%s.yaml", cloudProviderName))
}

func GetApplicationYamlPath(applicationName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/index.yaml", applicationName))
}

func GetEnvironmentYamlPath(applicationName string, environmentName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/Environments/%s/Index.yaml", applicationName, environmentName))
}

func GetInfraDefinitionYamlPath(applicationName string, environmentName string, infraName string) YamlPath {
	return YamlPath(fmt.Sprintf("Setup/Applications/%s/Environments/%s/Infrastructure Definitions/%s.yaml", applicationName, environmentName, infraName))
}

func (i *InfrastructureDetail) ToAwsAmi() *InfrastructureAwsAmi {
	if i.Type != InfrastructureTypes.AwsAmi {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsAmi))
	}

	d := &InfrastructureAwsAmi{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAwsAmi) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsAmi,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAwsEcs() *InfrastructureAwsEcs {
	if i.Type != InfrastructureTypes.AwsEcs {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsEcs))
	}

	d := &InfrastructureAwsEcs{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAwsEcs) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsEcs,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAwsLambda() *InfrastructureAwsLambda {
	if i.Type != InfrastructureTypes.AwsLambda {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsLambda))
	}

	d := &InfrastructureAwsLambda{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAwsLambda) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsLambda,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAwsWinRm() *InfrastructureAwsWinRM {
	if i.Type != InfrastructureTypes.AwsSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsSSH))
	}

	d := &InfrastructureAwsWinRM{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAwsWinRM) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsSSH,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAwsSSH() *InfrastructureAwsSSH {
	if i.Type != InfrastructureTypes.AwsSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsSSH))
	}

	d := &InfrastructureAwsSSH{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAwsSSH) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsSSH,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAzureVmss() *InfrastructureAzureVmss {
	if i.Type != InfrastructureTypes.AzureVmss {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AzureVmss))
	}

	d := &InfrastructureAzureVmss{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAzureVmss) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AzureVmss,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToAzureWebApp() *InfrastructureAzureWebApp {
	if i.Type != InfrastructureTypes.AzureWebApp {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AzureWebApp))
	}

	d := &InfrastructureAzureWebApp{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureAzureWebApp) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AzureWebApp,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToCustom() *InfrastructureCustom {
	if i.Type != InfrastructureTypes.Custom {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.Custom))
	}

	d := &InfrastructureCustom{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureCustom) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.Custom,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToDataCenterSSH() *InfrastructureDataCenterSSH {
	if i.Type != InfrastructureTypes.DataCenterSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.DataCenterSSH))
	}

	d := &InfrastructureDataCenterSSH{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureDataCenterSSH) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.DataCenterSSH,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToDataCenterWinRM() *InfrastructureDataCenterWinRM {
	if i.Type != InfrastructureTypes.DataCenterWinRM {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.DataCenterWinRM))
	}

	d := &InfrastructureDataCenterWinRM{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureDataCenterWinRM) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.DataCenterWinRM,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToKubernetesDirect() *InfrastructureKubernetesDirect {
	if i.Type != InfrastructureTypes.KubernetesDirect {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.KubernetesDirect))
	}

	d := &InfrastructureKubernetesDirect{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureKubernetesDirect) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.KubernetesDirect,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToKubernetesGcp() *InfrastructureKubernetesGcp {
	if i.Type != InfrastructureTypes.KubernetesGcp {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.KubernetesGcp))
	}

	d := &InfrastructureKubernetesGcp{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureKubernetesGcp) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.KubernetesGcp,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}

func (i *InfrastructureDetail) ToPcf() *InfrastructureTanzu {
	if i.Type != InfrastructureTypes.Pcf {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.Pcf))
	}

	d := &InfrastructureTanzu{}
	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return d
}

func (i *InfrastructureTanzu) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.Pcf,
	}

	if err := copier.Copy(d, i); err != nil {
		panic(err)
	}

	return []*InfrastructureDetail{d}
}
