# Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID that can be references in detail API calls | [optional] 
**Name** | **string** | Name of the object | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | [optional] 
**SourceVm** | [**Reference**](Reference.md) |  | [optional] 
**DestinationVm** | [**Reference**](Reference.md) |  | [optional] 
**SourceVnic** | [**Reference**](Reference.md) |  | [optional] 
**DestinationVnic** | [**Reference**](Reference.md) |  | [optional] 
**SourceVpc** | [**Reference**](Reference.md) |  | [optional] 
**DestinationVpc** | [**Reference**](Reference.md) |  | [optional] 
**SourceCloudNetwork** | [**Reference**](Reference.md) |  | [optional] 
**DestinationCloudNetwork** | [**Reference**](Reference.md) |  | [optional] 
**SourceAzureNsg** | [**Reference**](Reference.md) |  | [optional] 
**DestinationAzureNsg** | [**Reference**](Reference.md) |  | [optional] 
**SourceDatacenter** | [**Reference**](Reference.md) |  | [optional] 
**DestinationDatacenter** | [**Reference**](Reference.md) |  | [optional] 
**SourceIp** | [**IpV4Address**](IpV4Address.md) |  | [optional] 
**DestinationIp** | [**IpV4Address**](IpV4Address.md) |  | [optional] 
**SourceL2Network** | [**Reference**](Reference.md) |  | [optional] 
**DestinationL2Network** | [**Reference**](Reference.md) |  | [optional] 
**Port** | [**PortRange**](PortRange.md) |  | [optional] 
**SourceFolders** | [**[]Reference**](Reference.md) |  | [optional] 
**DestinationFolders** | [**[]Reference**](Reference.md) |  | [optional] 
**SourceResourcePool** | [**Reference**](Reference.md) |  | [optional] 
**DestinationResourcePool** | [**Reference**](Reference.md) |  | [optional] 
**SourceCluster** | [**Reference**](Reference.md) |  | [optional] 
**DestinationCluster** | [**Reference**](Reference.md) |  | [optional] 
**Protocol** | [**Protocol**](Protocol.md) |  | [optional] 
**SourceIpSets** | [**[]Reference**](Reference.md) |  | [optional] 
**DestinationIpSets** | [**[]Reference**](Reference.md) |  | [optional] 
**SourceSecurityGroups** | [**[]Reference**](Reference.md) |  | [optional] 
**DestinationSecurityGroups** | [**[]Reference**](Reference.md) |  | [optional] 
**TrafficType** | [**FlowTrafficType**](FlowTrafficType.md) |  | [optional] 
**SourceSecurityTags** | [**[]Reference**](Reference.md) |  | [optional] 
**DestinationSecurityTags** | [**[]Reference**](Reference.md) |  | [optional] 
**SourceHost** | [**Reference**](Reference.md) |  | [optional] 
**DestinationHost** | [**Reference**](Reference.md) |  | [optional] 
**SourceVmTags** | **[]string** |  | [optional] 
**DestinationVmTags** | **[]string** |  | [optional] 
**WithinHost** | **bool** |  | [optional] 
**FirewallAction** | [**FirewallAction**](FirewallAction.md) |  | [optional] 
**FirewallRuleId** | **string** |  | [optional] 
**FlowTag** | [**[]FlowTag**](FlowTag.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


