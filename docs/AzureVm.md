# AzureVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID that can be references in detail API calls | [optional] 
**Name** | **string** | Name of the object | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | [optional] 
**IpAddresses** | [**[]IpV4Address**](IpV4Address.md) |  | [optional] 
**DefaultGateway** | **string** | Default gateway IP of the VM | [optional] 
**Vnics** | [**[]Reference**](Reference.md) |  | [optional] 
**SecurityGroups** | [**[]Reference**](Reference.md) |  | [optional] 
**SourceFirewallRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**DestinationFirewallRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**IpSets** | [**[]Reference**](Reference.md) |  | [optional] 
**Vnet** | [**Reference**](Reference.md) |  | [optional] 
**Region** | **string** | The Azure region this VM is hosted | [optional] 
**Subscription** | [**Reference**](Reference.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


