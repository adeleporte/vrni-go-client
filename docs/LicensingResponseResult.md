# LicensingResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]VrniLicense**](VRNILicense.md) |  | [optional] 
**LicenseUsage** | [**[][]map[string]interface{}**](array.md) |  | [optional] 
**CustomerId** | **int32** | customer id | [optional] 
**CpuCapacity** | **int32** | cpu capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period) | [optional] 
**CcuCapacity** | **int32** | ccu capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period) | [optional] 
**VmCapacity** | **int32** | VM capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period) | [optional] 
**CoreCapacity** | **int32** | core capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period) | [optional] 
**EdgeCapacity** | **int32** | edge capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period) | [optional] 
**ServiceTag** | **string** | service tag | [optional] 
**ExpireDate** | **int32** | Maximum of expiry date for current licenses, 0 if any non-expiry license is present. Invalid/already expired license is not considered here. | [optional] 
**Enterprise** | **bool** | true if non-expired ENTERPRISE license is available | [optional] 
**Evaluation** | **bool** | true if all licenses are evaluation license | [optional] 
**NumOfEnterpriseLicense** | **int32** | number of total enterprise licenses available | [optional] 
**NumOfValidLicense** | **int32** | Total Number of valid licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfValidEnterpriseLicense** | **int32** | Number of valid Enterprise licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfValidSDWANLicense** | **int32** | Number of valid SDWAN licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfValidNAVLicense** | **int32** | Number of valid NAV licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfValidAdvancedLicense** | **int32** | Number of valid Advanced licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfValidROBOLicense** | **int32** | Number of valid ROBO licenses available. Active licenses from last major version which are only valid for grace period are not counted | [optional] 
**NumOfActiveLicense** | **int32** | Total Number of active licenses (non-expired and licenses from last major version during grace period) | [optional] 
**NumOfActiveEnterpriseLicense** | **int32** | Total Number of active Enterprise licenses (non-expired and licenses from last major version during grace period) | [optional] 
**NumOfActiveSDWANLicense** | **int32** | Total Number of active SDWAN licenses (non-expired and licenses from last major version during grace period) | [optional] 
**NumOfActiveNAVLicense** | **int32** | Total Number of active NAV licenses (non-expired and licenses from last major version during grace period) | [optional] 
**NumOfActiveAdvancedLicense** | **int32** | Total Number of active Advanced licenses (non-expired and licenses from last major version during grace period) | [optional] 
**NumOfActiveROBOLicense** | **int32** | Total Number of active ROBO licenses (non-expired and licenses from last major version during grace period) | [optional] 
**UsedSockets** | **int32** | Count of used sockets | [optional] 
**UsedCores** | **int32** | Count of used cores | [optional] 
**UsedEdges** | **int32** | Count of used edges | [optional] 
**UsedDevices** | **int32** | Count of used devices | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


