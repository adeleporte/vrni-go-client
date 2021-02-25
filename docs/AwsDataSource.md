# AwsDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Internal ID of data source, to be used in subsequent API calls | [optional] 
**EntityType** | [**DataSourceType**](DataSourceType.md) |  | [optional] 
**Ip** | **string** | IP address of data source (use either IP or FQDN, not both) | [optional] 
**Fqdn** | **string** | Hostname of data source (use either IP or FQDN, not both) | [optional] 
**ProxyId** | **string** | ID of Collector VM which should register this vcenter | [optional] 
**Nickname** | **string** | A friendly nickname for the data source | [optional] 
**Enabled** | **bool** | Whether or not data collection is enabled | [optional] [default to true]
**Notes** | **string** | Room for notes or comments about the data source | [optional] 
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


