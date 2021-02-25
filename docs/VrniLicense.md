# VrniLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** |  | [optional] 
**Evaluation** | **bool** | true if license is evaluation license | [optional] 
**LicenseKey** | **string** | license serial key | [optional] 
**NumberOfSockets** | **int32** | sockets entitlement for this license | [optional] 
**ServiceTag** | **string** | service tag | [optional] 
**DeploymentType** | **string** | environment deployment type | [optional] 
**CapacityType** | [**LicenseCapacityType**](LicenseCapacityType.md) |  | [optional] 
**Edition** | [**LicenseEdition**](LicenseEdition.md) |  | [optional] 
**Suspended** | **bool** | true if customer state is suspended | [optional] 
**Deleted** | **bool** | true if customer state is deleted | [optional] 
**Invalid** | **bool** | true for license keys from last version, which are active only for grace period | [optional] 
**Entitlements** | [**[]VrniLicenseEntitlements**](VRNILicense_entitlements.md) | list of entitlements supported by the license | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


