/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SubnetMappingRequest struct for SubnetMappingRequest
type SubnetMappingRequest struct {
	Cidr string `json:"cidr,omitempty"`
	VlanId int32 `json:"vlan_id,omitempty"`
}