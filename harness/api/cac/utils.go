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

func (s *Service) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(s, []string{"ApplicationId"})
}

func (e *Environment) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(e, []string{"Name", "EnvironmentType", "ApplicationId"})
}

func (cp *GcpCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

func (cp *PhysicalDatacenterCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

func (i *InfrastructureDefinition) Validate() (bool, error) {
	if _, err := utils.RequiredStringFieldsSet(i, []string{"ApplicationId", "EnvironmentId"}); err != nil {
		return false, err
	}

	if len(i.InfrastructureDetail) != 1 {
		return false, errors.New("expect one infrastructure detail to be set")
	}

	// detail := i.InfrastructureDetail[0]

	// switch i.CloudProviderType {
	// case CloudProviderTypes.DataCenter:
	// 	switch i.DeploymentType {
	// 	case DeploymentTypes.SSH:
	// 		return detail.ToDataCenterSSH().Validate()
	// 	case DeploymentTypes.WinRM:
	// 		return detail.ToDataCenterWinRM().Validate()
	// 	default:
	// 		return false, fmt.Errorf("unsupported deployment type '%s' for '%s' cloud provider", i.DeploymentType, i.CloudProviderType)
	// 	}
	// case CloudProviderTypes.KubernetesCluster:
	// 	switch i.DeploymentType {
	// 	case DeploymentTypes.Kubernetes:
	// 		return detail.ToKubernetesDirect().Validate()
	// 	case DeploymentTypes.Helm:
	// 		return detail.ToKubernetesDirect().Validate()
	// 	default:
	// 		return false, fmt.Errorf("unsupported deployment type '%s' for '%s' cloud provider", i.DeploymentType, i.CloudProviderType)
	// 	}
	// case CloudProviderTypes.Aws:
	// 	switch i.DeploymentType {
	// 	case DeploymentTypes.AMI:
	// 		return detail.ToAwsAmi().Validate()
	// 	}
	// default:
	// 	return false, fmt.Errorf("unknown cloud provider type '%s'", i.CloudProviderType)
	// }

	return true, nil
}

// func (i *InfrastructureAwsAmi) Validate() (bool, error) {
// 	if _, err := utils.RequiredValueOptionsSet(i, map[string][]interface{}{
// 		"AmiDeploymentType": {AmiDeploymentTypes.ASG},
// 	}); err != nil {
// 		return false, err
// 	}

// 	if len(i.HostNames) == 0 {
// 		return false, errors.New("host names must be set")
// 	}

// 	return true, nil
// }

func (i *InfrastructureDataCenterSSH) Validate() (bool, error) {
	if _, err := utils.RequiredStringFieldsSet(i, []string{"CloudProviderName", "HostConnectionAttrsName"}); err != nil {
		return false, err
	}

	if len(i.HostNames) == 0 {
		return false, errors.New("host names must be set")
	}

	return true, nil
}

func (i *InfrastructureDataCenterWinRM) Validate() (bool, error) {
	if _, err := utils.RequiredStringFieldsSet(i, []string{"CloudProviderName", "WinRmConnectionAttributesName"}); err != nil {
		return false, err
	}

	if len(i.HostNames) == 0 {
		return false, errors.New("host names must be set")
	}

	return true, nil
}

func (i *InfrastructureKubernetesDirect) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(i, []string{"CloudProviderName", "Namespace", "ReleaseName"})
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

func (d *InfrastructureDetail) ToAwsAmi() *InfrastructureAwsAmi {
	if d.Type != InfrastructureTypes.AwsAmi {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsAmi))
	}

	i := &InfrastructureAwsAmi{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAwsAmi) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsAmi,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAwsEcs() *InfrastructureAwsEcs {
	if d.Type != InfrastructureTypes.AwsEcs {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsEcs))
	}

	i := &InfrastructureAwsEcs{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAwsEcs) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsEcs,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAwsLambda() *InfrastructureAwsLambda {
	if d.Type != InfrastructureTypes.AwsLambda {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsLambda))
	}

	i := &InfrastructureAwsLambda{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAwsLambda) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsLambda,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAwsWinRm() *InfrastructureAwsWinRM {
	if d.Type != InfrastructureTypes.AwsSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsSSH))
	}

	i := &InfrastructureAwsWinRM{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAwsWinRM) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsSSH,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAwsSSH() *InfrastructureAwsSSH {
	if d.Type != InfrastructureTypes.AwsSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AwsSSH))
	}

	i := &InfrastructureAwsSSH{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAwsSSH) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AwsSSH,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAzureVmss() *InfrastructureAzureVmss {
	if d.Type != InfrastructureTypes.AzureVmss {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AzureVmss))
	}

	i := &InfrastructureAzureVmss{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAzureVmss) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AzureVmss,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToAzureWebApp() *InfrastructureAzureWebApp {
	if d.Type != InfrastructureTypes.AzureWebApp {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.AzureWebApp))
	}

	i := &InfrastructureAzureWebApp{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureAzureWebApp) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.AzureWebApp,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToCustom() *InfrastructureCustom {
	if d.Type != InfrastructureTypes.Custom {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.Custom))
	}

	i := &InfrastructureCustom{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureCustom) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.Custom,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToDataCenterSSH() *InfrastructureDataCenterSSH {
	if d.Type != InfrastructureTypes.DataCenterSSH {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.DataCenterSSH))
	}

	i := &InfrastructureDataCenterSSH{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureDataCenterSSH) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.DataCenterSSH,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToDataCenterWinRM() *InfrastructureDataCenterWinRM {
	if d.Type != InfrastructureTypes.DataCenterWinRM {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.DataCenterWinRM))
	}

	i := &InfrastructureDataCenterWinRM{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureDataCenterWinRM) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.DataCenterWinRM,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToKubernetesDirect() *InfrastructureKubernetesDirect {
	if d.Type != InfrastructureTypes.KubernetesDirect {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.KubernetesDirect))
	}

	i := &InfrastructureKubernetesDirect{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureKubernetesDirect) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.KubernetesDirect,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToKubernetesGcp() *InfrastructureKubernetesGcp {
	if d.Type != InfrastructureTypes.KubernetesGcp {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.KubernetesGcp))
	}

	i := &InfrastructureKubernetesGcp{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureKubernetesGcp) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.KubernetesGcp,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}

func (d *InfrastructureDetail) ToPcf() *InfrastructureTanzu {
	if d.Type != InfrastructureTypes.Pcf {
		panic(fmt.Errorf("expected Type of %s", InfrastructureTypes.Pcf))
	}

	i := &InfrastructureTanzu{}
	copier.Copy(i, d)
	return i
}

func (i *InfrastructureTanzu) ToInfrastructureDetail() []*InfrastructureDetail {
	d := &InfrastructureDetail{
		Type: InfrastructureTypes.Pcf,
	}
	copier.Copy(d, i)
	return []*InfrastructureDetail{d}
}
