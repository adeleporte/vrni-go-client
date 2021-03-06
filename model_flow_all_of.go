/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// FlowAllOf struct for FlowAllOf
type FlowAllOf struct {
	SourceVm Reference `json:"source_vm,omitempty"`
	DestinationVm Reference `json:"destination_vm,omitempty"`
	SourceVnic Reference `json:"source_vnic,omitempty"`
	DestinationVnic Reference `json:"destination_vnic,omitempty"`
	SourceVpc Reference `json:"source_vpc,omitempty"`
	DestinationVpc Reference `json:"destination_vpc,omitempty"`
	SourceCloudNetwork Reference `json:"source_cloud_network,omitempty"`
	DestinationCloudNetwork Reference `json:"destination_cloud_network,omitempty"`
	SourceAzureNsg Reference `json:"source_azure_nsg,omitempty"`
	DestinationAzureNsg Reference `json:"destination_azure_nsg,omitempty"`
	SourceDatacenter Reference `json:"source_datacenter,omitempty"`
	DestinationDatacenter Reference `json:"destination_datacenter,omitempty"`
	SourceIp IpV4Address `json:"source_ip,omitempty"`
	DestinationIp IpV4Address `json:"destination_ip,omitempty"`
	SourceL2Network Reference `json:"source_l2_network,omitempty"`
	DestinationL2Network Reference `json:"destination_l2_network,omitempty"`
	Port PortRange `json:"port,omitempty"`
	SourceFolders []Reference `json:"source_folders,omitempty"`
	DestinationFolders []Reference `json:"destination_folders,omitempty"`
	SourceResourcePool Reference `json:"source_resource_pool,omitempty"`
	DestinationResourcePool Reference `json:"destination_resource_pool,omitempty"`
	SourceCluster Reference `json:"source_cluster,omitempty"`
	DestinationCluster Reference `json:"destination_cluster,omitempty"`
	Protocol Protocol `json:"protocol,omitempty"`
	SourceIpSets []Reference `json:"source_ip_sets,omitempty"`
	DestinationIpSets []Reference `json:"destination_ip_sets,omitempty"`
	SourceSecurityGroups []Reference `json:"source_security_groups,omitempty"`
	DestinationSecurityGroups []Reference `json:"destination_security_groups,omitempty"`
	TrafficType FlowTrafficType `json:"traffic_type,omitempty"`
	SourceSecurityTags []Reference `json:"source_security_tags,omitempty"`
	DestinationSecurityTags []Reference `json:"destination_security_tags,omitempty"`
	SourceHost Reference `json:"source_host,omitempty"`
	DestinationHost Reference `json:"destination_host,omitempty"`
	SourceVmTags []string `json:"source_vm_tags,omitempty"`
	DestinationVmTags []string `json:"destination_vm_tags,omitempty"`
	WithinHost bool `json:"within_host,omitempty"`
	FirewallAction FirewallAction `json:"firewall_action,omitempty"`
	FirewallRuleId string `json:"firewall_rule_id,omitempty"`
	FlowTag []FlowTag `json:"flow_tag,omitempty"`
}
