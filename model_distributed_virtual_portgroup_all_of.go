/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// DistributedVirtualPortgroupAllOf struct for DistributedVirtualPortgroupAllOf
type DistributedVirtualPortgroupAllOf struct {
	VendorId string `json:"vendor_id,omitempty"`
	VcenterManager Reference `json:"vcenter_manager,omitempty"`
	DistributedVirtualSwitch Reference `json:"distributed_virtual_switch,omitempty"`
}
