/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// BaseVnic struct for BaseVnic
type BaseVnic struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	IpAddresses []IpV4Address `json:"ip_addresses,omitempty"`
	Layer2Network Reference `json:"layer2_network,omitempty"`
	Vlan Vlan `json:"vlan,omitempty"`
	Vm Reference `json:"vm,omitempty"`
}
