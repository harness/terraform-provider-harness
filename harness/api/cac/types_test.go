package cac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestEnvironmentSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.Environment).(*Environment)
	testObj.EnvironmentType = EnvironmentTypes.Prod
	testObj.VariableOverrides = []*VariableOverride{{
		ServiceName: "servicename",
		Value:       "gcpkms:secret_name",
		ValueType:   VariableOverrideValueTypes.EncryptedText,
	}}

	expectedObjYaml := `
harnessApiVersion: "1.0"
type: ENVIRONMENT
environmentType: PROD
variableOverrides:
  - serviceName: servicename
    value: gcpkms:secret_name
    valueType: ENCRYPTED_TEXT
`

	expectedObj := &Environment{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestSpotInstCloudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.SpotInstCloudProvider).(*SpotInstCloudProvider)
	testObj.AccountId = "accountId"
	testObj.Token = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "secret_name",
	}

	expectedObjYaml := `
harnessApiVersion: "1.0"
type: SPOT_INST
spotInstAccountId: accountId
spotInstToken: gcpkms:secret_name
`

	expectedObj := &SpotInstCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestKubernetesCLoudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.KubernetesCloudProvider).(*KubernetesCloudProvider)
	testObj.AuthType = KubernetesAuthTypes.UsernameAndPassword
	testObj.CACert = &SecretRef{
		Name: "cert",
	}
	testObj.ClientCert = &SecretRef{
		Name: "clientCert",
	}
	testObj.ClientKey = &SecretRef{
		Name: "clientKey",
	}
	testObj.ClientKeyAlgorithm = "algorithm"
	testObj.ClientKeyPassPhrase = &SecretRef{
		Name: "passphrase",
	}
	testObj.ContinuousEfficiencyConfig = &ContinuousEfficiencyConfig{
		ContinuousEfficiencyEnabled: true,
	}
	testObj.MasterUrl = "masterurl"
	testObj.OIDCClientId = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "secret_name",
	}
	testObj.OIDCIdentityProviderUrl = "providerUrl"
	testObj.OIDCPassword = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "secret_name",
	}
	testObj.OIDCScopes = "scope1 scope2"
	testObj.OIDCUsername = "username"
	testObj.ServiceAccountToken = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "token",
	}
	testObj.SkipValidation = true
	testObj.UseEncryptedUsername = true
	testObj.UseKubernetesDelegate = true
	testObj.DelegateSelectors = []string{"test"}

	data, _ := yaml.Marshal(testObj)
	fmt.Println(string(data))
	expectedObjYaml := `
harnessApiVersion: "1.0"
type: KUBERNETES_CLUSTER
authType: USER_PASSWORD
cacert: cert
clientCert: clientCert
clientKey: clientKey
clientKeyAlgorithm: algorithm
clientKeyPassPhrase: passphrase
delegateSelectors:
  - test
continuousEfficiencyConfig:
  continuousefficiencyenabled: true
masterUrl: masterurl
serviceAccountToken: gcpkms:token
skipValidation: true
useKubernetesDelegate: true
useEncryptedUsername: true
oidcClientId: gcpkms:secret_name
oidcIdentityProviderUrl: providerUrl
oidcPassword: gcpkms:secret_name
oidcScopes: scope1 scope2
oidcUsername: username
`

	expectedObj := &KubernetesCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestAwsCloudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.AwsCloudProvider).(*AwsCloudProvider)
	testObj.AccessKey = "accessKey"
	testObj.AssumeCrossAccountRole = true
	testObj.CrossAccountAttributes = &AwsCrossAccountAttributes{
		CrossAccountRoleArn: "roleArn",
		ExternalId:          "externalId",
	}
	testObj.SecretKey = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "secret_name",
	}
	testObj.UseEc2IamCredentials = true
	testObj.UseIRSA = true
	testObj.DelegateSelector = "selector"

	expectedObjYaml := `
harnessApiVersion: "1.0"
type: AWS
accessKey: accessKey
assumeCrossAccountRole: true
secretKey: gcpkms:secret_name
useEc2IamCredentials: true
useIRSA: true
tag: selector
crossAccountAttributes:
  crossAccountRoleArn: roleArn
  externalId: externalId
`

	expectedObj := &AwsCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestAzureCloudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.AzureCloudProvider).(*AzureCloudProvider)
	testObj.ClientId = "clientId"
	testObj.Key = &SecretRef{
		SecretManagerType: SecretManagerTypes.GcpKMS,
		Name:              "secret_name",
	}
	testObj.TenantId = "tenantId"
	testObj.AzureEnvironmentType = AzureEnvironmentTypes.AzureGlobal
	testObj.ClientId = "clientId"

	expectedObjYaml := `
harnessApiVersion: "1.0"
type: AZURE
azureEnvironmentType: AZURE
clientId: clientId
key: gcpkms:secret_name
tenantId: tenantId
`

	expectedObj := &AzureCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestPcfCloudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.PcfCloudProvider).(*PcfCloudProvider)
	testObj.EndpointUrl = "http://endpoint.com"
	testObj.Password = &SecretRef{
		SecretManagerType: SecretManagerTypes.AwsKMS,
		Name:              "secret_name",
	}
	testObj.SkipValidation = true
	testObj.Username = "username"

	expectedObjYaml := `
harnessApiVersion: "1.0"
type: PCF
endpointUrl: http://endpoint.com
password: amazonkms:secret_name
skipValidation: true
username: username
`

	expectedObj := &PcfCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestPhysicalDataCenterCloudProviderSerialization(t *testing.T) {
	testObj := NewEntity(ObjectTypes.PhysicalDataCenterCloudProvider)
	expectedObjYaml := `
harnessApiVersion: "1.0"
type: PHYSICAL_DATA_CENTER
`

	expectedObj := &PhysicalDatacenterCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

func TestGcpCloudProviderSerialization(t *testing.T) {
	testObj := &GcpCloudProvider{
		HarnessApiVersion: HarnessApiVersions.V1,
		Type:              ObjectTypes.GcpCloudProvider,
		DelegateSelectors: []string{"primary"},
		SkipValidation:    true,
		ServiceAccountKeyFileContent: &SecretRef{
			SecretManagerType: SecretManagerTypes.AwsKMS,
			Name:              "abc123",
		},
		UsageRestrictions: &UsageRestrictions{
			AppEnvRestrictions: []*AppEnvRestriction{
				{
					AppFilter: &AppFilter{
						EntityNames: []string{"TestAccDataSourceService_bJmj"},
						FilterType:  ApplicationFilterTypes.Selected,
					},
					EnvFilter: &EnvFilter{
						FilterTypes: []EnvironmentFilterType{EnvironmentFilterTypes.Prod},
					},
				},
				{
					AppFilter: &AppFilter{
						FilterType: ApplicationFilterTypes.All,
					},
					EnvFilter: &EnvFilter{
						FilterTypes: []EnvironmentFilterType{EnvironmentFilterTypes.NonProd},
					},
				},
			},
		},
		UseDelegate:          false,
		UseDelegateSelectors: true,
	}

	expectedObjYaml := `
harnessApiVersion: '1.0'
type: GCP
delegateSelectors:
- primary
skipValidation: true
serviceAccountKeyFileContent: amazonkms:abc123
usageRestrictions:
  appEnvRestrictions:
  - appFilter:
      entityNames:
      - TestAccDataSourceService_bJmj
      filterType: SELECTED
    envFilter:
      filterTypes:
      - PROD
  - appFilter:
      filterType: ALL
    envFilter:
      filterTypes:
      - NON_PROD
useDelegate: false
useDelegateSelectors: true
`

	expectedObj := &GcpCloudProvider{}
	err := yaml.Unmarshal([]byte(expectedObjYaml), expectedObj)
	require.NoError(t, err)
	require.Equal(t, expectedObj, testObj)
}

type TestSecretRefMarshal struct {
	SecretKeyId *SecretRef `yaml:"secretKeyId"`
}

func TestSecretRefMarshalYaml(t *testing.T) {

	testStruct := &TestSecretRefMarshal{
		SecretKeyId: &SecretRef{
			SecretManagerType: SecretManagerTypes.AwsKMS,
			Name:              "abc123",
		},
	}

	bytes, err := yaml.Marshal(&testStruct)
	require.NoError(t, err)
	fmt.Println(string(bytes))
	require.Equal(t, "secretKeyId: amazonkms:abc123\n", string(bytes))

	newStruct := &TestSecretRefMarshal{}
	err = yaml.Unmarshal(bytes, newStruct)
	require.NoError(t, err)
	require.Equal(t, testStruct.SecretKeyId.SecretManagerType, newStruct.SecretKeyId.SecretManagerType)
	require.Equal(t, testStruct.SecretKeyId.Name, newStruct.SecretKeyId.Name)

}
