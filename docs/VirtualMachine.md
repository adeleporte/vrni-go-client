# VirtualMachine

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
**Cluster** | [**Reference**](Reference.md) |  | [optional] 
**ResourcePool** | [**Reference**](Reference.md) |  | [optional] 
**SecurityTags** | [**[]Reference**](Reference.md) |  | [optional] 
**Layer2Networks** | [**[]Reference**](Reference.md) |  | [optional] 
**Host** | [**Reference**](Reference.md) |  | [optional] 
**Vlans** | [**[]Vlan**](Vlan.md) |  | [optional] 
**VendorId** | **string** |  | [optional] 
**VcenterManager** | [**Reference**](Reference.md) |  | [optional] 
**Folders** | [**[]Reference**](Reference.md) |  | [optional] 
**Datastores** | [**[]Reference**](Reference.md) |  | [optional] 
**Datacenter** | [**Reference**](Reference.md) |  | [optional] 
**NsxManager** | [**Reference**](Reference.md) |  | [optional] 
**SourceInversionRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**DestinationInversionRules** | [**[]RuleSet**](RuleSet.md) |  | [optional] 
**CpuCount** | **int32** | Number of vCPUs this VM has | [optional] 
**Memory** | **int32** | Amount of memory this VM has | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


