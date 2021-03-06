/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VlanL2NetworkAllOf struct for VlanL2NetworkAllOf
type VlanL2NetworkAllOf struct {
	VlanId string `json:"vlan_id,omitempty"`
	DistributedVirtualSwitches []Reference `json:"distributed_virtual_switches,omitempty"`
	DistributedVirtualPortgroups []Reference `json:"distributed_virtual_portgroups,omitempty"`
}
