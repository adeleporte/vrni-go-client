/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsService struct for NsService
type NsService struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	PortRanges []PortRange `json:"port_ranges,omitempty"`
	NsxtManagers []Reference `json:"nsxt_managers,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
}