/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// Ec2SecurityGroup struct for Ec2SecurityGroup
type Ec2SecurityGroup struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Members []Reference `json:"members,omitempty"`
	DirectSourceRules []RuleSet `json:"direct_source_rules,omitempty"`
	DirectDestinationRules []RuleSet `json:"direct_destination_rules,omitempty"`
	IndirectSourceRules []RuleSet `json:"indirect_source_rules,omitempty"`
	IndirectDestinationRules []RuleSet `json:"indirect_destination_rules,omitempty"`
	Parents []Reference `json:"parents,omitempty"`
	DirectMembers []Reference `json:"direct_members,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
	ExcludedMembers []Reference `json:"excluded_members,omitempty"`
	Vpc Reference `json:"vpc,omitempty"`
	Region string `json:"region,omitempty"`
}
