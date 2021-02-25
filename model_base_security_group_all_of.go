/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// BaseSecurityGroupAllOf struct for BaseSecurityGroupAllOf
type BaseSecurityGroupAllOf struct {
	DirectSourceRules []RuleSet `json:"direct_source_rules,omitempty"`
	DirectDestinationRules []RuleSet `json:"direct_destination_rules,omitempty"`
	IndirectSourceRules []RuleSet `json:"indirect_source_rules,omitempty"`
	IndirectDestinationRules []RuleSet `json:"indirect_destination_rules,omitempty"`
	Parents []Reference `json:"parents,omitempty"`
	Members []Reference `json:"members,omitempty"`
	DirectMembers []Reference `json:"direct_members,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
	ExcludedMembers []Reference `json:"excluded_members,omitempty"`
}