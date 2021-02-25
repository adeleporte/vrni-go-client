/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxtFirewall struct for NsxtFirewall
type NsxtFirewall struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	FirewallRules []RuleSet `json:"firewall_rules,omitempty"`
	Exclusions []Reference `json:"exclusions,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
}