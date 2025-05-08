# ApplicationsScmProviderGenerator

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Github** | [***ApplicationsScmProviderGeneratorGithub**](applicationsSCMProviderGeneratorGithub.md) |  | [optional] [default to null]
**Gitlab** | [***ApplicationsScmProviderGeneratorGitlab**](applicationsSCMProviderGeneratorGitlab.md) |  | [optional] [default to null]
**Bitbucket** | [***ApplicationsScmProviderGeneratorBitbucket**](applicationsSCMProviderGeneratorBitbucket.md) |  | [optional] [default to null]
**BitbucketServer** | [***ApplicationsScmProviderGeneratorBitbucketServer**](applicationsSCMProviderGeneratorBitbucketServer.md) |  | [optional] [default to null]
**Gitea** | [***ApplicationsScmProviderGeneratorGitea**](applicationsSCMProviderGeneratorGitea.md) |  | [optional] [default to null]
**AzureDevOps** | [***ApplicationsScmProviderGeneratorAzureDevOps**](applicationsSCMProviderGeneratorAzureDevOps.md) |  | [optional] [default to null]
**Filters** | [**[]ApplicationsScmProviderGeneratorFilter**](applicationsSCMProviderGeneratorFilter.md) | Filters for which repos should be considered. | [optional] [default to null]
**CloneProtocol** | **string** | Which protocol to use for the SCM URL. Default is provider-specific but ssh if possible. Not all providers necessarily support all protocols. | [optional] [default to null]
**RequeueAfterSeconds** | **string** | Standard parameters. | [optional] [default to null]
**Template** | [***ApplicationsApplicationSetTemplate**](applicationsApplicationSetTemplate.md) |  | [optional] [default to null]
**Values** | **map[string]string** |  | [optional] [default to null]
**AwsCodeCommit** | [***ApplicationsScmProviderGeneratorAwsCodeCommit**](applicationsSCMProviderGeneratorAWSCodeCommit.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

