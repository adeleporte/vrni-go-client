# AzureDataSource

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
**Credentials** | [**AzureCredentials**](AzureCredentials.md) |  | [optional] 
**FlowsEnabled** | **bool** | Whether or not to collect flow data from the VNET Flow Logs | [optional] [default to false]
**WebProxyId** | **string** | Identifier of web proxy to be used for connection. If not given or empty, will clear any web proxy used earlier | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


