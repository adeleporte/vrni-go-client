# AwsDataSourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**AwsCredentials**](AWSCredentials.md) |  | [optional] 
**FlowsEnabled** | **bool** | Whether or not to collect flows from the VPC Flow Logs | [optional] [default to false]
**AddLinkedAccounts** | **bool** | Automatically add linked accounts | [optional] [default to false]
**RoleArnSuffix** | **string** |  | [optional] 
**LinkedRoleArn** | **string** | Should not be populated. Will be ignored if populated. | [optional] 
**DataSourceHierarchy** | **string** | Should not be populated. Will be ignored if populated. | [optional] 
**EnableAwsGeoRestrictions** | **bool** | Limit collection to only regions specified in selected_regions | [optional] [default to false]
**SelectedRegions** | **[]string** |  | [optional] 
**ChildDatasourceCount** | **int32** | Should not be populated. Will be ignored if populated. | [optional] 
**ChildDatasources** | [**[]DataSourceEntityId**](DataSourceEntityId.md) | Should not be populated. Will be ignored if populated. | [optional] 
**WebProxyId** | **string** | Identifier of web proxy to be used for connection. If not given or empty, any web proxy used earlier will be cleared | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


