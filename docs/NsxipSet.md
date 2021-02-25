# NsxipSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID that can be references in detail API calls | [optional] 
**Name** | **string** | Name of the object | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | [optional] 
**IpAddresses** | [**[]IpV4Address**](IpV4Address.md) |  | [optional] 
**IpRanges** | [**[]IpAddressRange**](IpAddressRange.md) |  | [optional] 
**IpNumericRanges** | [**[]IpNumericRange**](IpNumericRange.md) |  | [optional] 
**ParentSecurityGroups** | [**[]Reference**](Reference.md) |  | [optional] 
**DirectSourceRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**DirectDestinationRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**IndirectSourceRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**IndirectDestinationRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**VendorId** | **string** |  | [optional] 
**Vendor** | **string** |  | [optional] 
**NsxManagers** | [**[]Reference**](Reference.md) |  | [optional] 
**Scope** | [**ScopeEnum**](ScopeEnum.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


