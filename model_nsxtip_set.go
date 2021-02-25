/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxtipSet struct for NsxtipSet
type NsxtipSet struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	IpAddresses []IpV4Address `json:"ip_addresses,omitempty"`
	IpRanges []IpAddressRange `json:"ip_ranges,omitempty"`
	IpNumericRanges []IpNumericRange `json:"ip_numeric_ranges,omitempty"`
	ParentSecurityGroups []Reference `json:"parent_security_groups,omitempty"`
	DirectSourceRules []RuleSet `json:"direct_source_rules,omitempty"`
	DirectDestinationRules []RuleSet `json:"direct_destination_rules,omitempty"`
	IndirectSourceRules []RuleSet `json:"indirect_source_rules,omitempty"`
	IndirectDestinationRules []RuleSet `json:"indirect_destination_rules,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	NsxtManagers []Reference `json:"nsxt_managers,omitempty"`
}
