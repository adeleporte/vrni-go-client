# PolicyManagerDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address of data source (use either IP or FQDN, not both) | [optional] 
**Fqdn** | **string** | Hostname of data source (use either IP or FQDN, not both) | [optional] 
**ProxyId** | **string** | ID of Collector VM which should register this vcenter | 
**Nickname** | **string** | A friendly nickname for the data source | 
**Enabled** | **bool** | Whether or not data collection is enabled | [optional] [default to true]
**Notes** | **string** | Room for notes or comments about the data source | [optional] 
**CspRefreshToken** | **string** | Cloud Services Portal API Refresh Token. You can generate this via the CSP UI under My Account and API Tokens. | [optional] 
**VcenterId** | **string** | Associated vcenter data source entity ID | 
**IpfixEnabled** | **bool** | Whether or not to configure NSX-T to send IPFIX to vRNI | [optional] [default to false]
**WebProxyId** | **string** | Identifier of web proxy to be used for connection. If not given or empty, any web proxy used earlier will be cleared | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


