# NsxtManagerDataSource

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
**Credentials** | [**PasswordCredentials**](PasswordCredentials.md) |  | [optional] 
**IpfixEnabled** | **bool** | Whether or not to configure NSX-T to send IPFIX to vRNI | [optional] [default to false]
**LatencyEnabled** | **bool** | Whether or not to configure NSX-T to send virtual infrastructure latency metrics to vRNI | [optional] [default to false]
**NsxiEnabled** | **bool** | Whether or not to configure NSX Intelligence to send additional traffic information to vRNI | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


