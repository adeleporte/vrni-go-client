# NsxtFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID that can be references in detail API calls | [optional] 
**Name** | **string** | Name of the object | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | [optional] 
**RuleId** | **string** |  | [optional] 
**SectionId** | **string** |  | [optional] 
**SectionName** | **string** |  | [optional] 
**SequenceNumber** | **int32** |  | [optional] 
**SourceAny** | **bool** |  | [optional] 
**DestinationAny** | **bool** |  | [optional] 
**ServiceAny** | **bool** |  | [optional] 
**Sources** | [**[]Reference**](Reference.md) |  | [optional] 
**Destinations** | [**[]Reference**](Reference.md) |  | [optional] 
**Services** | [**[]Reference**](Reference.md) |  | [optional] 
**Action** | [**FirewallAction**](FirewallAction.md) |  | [optional] 
**Disabled** | **bool** |  | [optional] 
**SourceInversion** | **bool** |  | [optional] 
**DestinationInversion** | **bool** |  | [optional] 
**PortRanges** | [**[]PortRange**](PortRange.md) |  | [optional] 
**LoggingEnabled** | **bool** |  | [optional] 
**Direction** | [**FirewallDirection**](FirewallDirection.md) |  | [optional] 
**Scope** | [**ScopeEnum**](ScopeEnum.md) |  | [optional] 
**NsxManagers** | [**[]Reference**](Reference.md) |  | [optional] 
**AppliedTos** | [**[]AppliedTo**](AppliedTo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


