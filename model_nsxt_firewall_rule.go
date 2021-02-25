/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxtFirewallRule struct for NsxtFirewallRule
type NsxtFirewallRule struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	RuleId string `json:"rule_id,omitempty"`
	SectionId string `json:"section_id,omitempty"`
	SectionName string `json:"section_name,omitempty"`
	SequenceNumber int32 `json:"sequence_number,omitempty"`
	SourceAny bool `json:"source_any,omitempty"`
	DestinationAny bool `json:"destination_any,omitempty"`
	ServiceAny bool `json:"service_any,omitempty"`
	Sources []Reference `json:"sources,omitempty"`
	Destinations []Reference `json:"destinations,omitempty"`
	Services []Reference `json:"services,omitempty"`
	Action FirewallAction `json:"action,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	SourceInversion bool `json:"source_inversion,omitempty"`
	DestinationInversion bool `json:"destination_inversion,omitempty"`
	PortRanges []PortRange `json:"port_ranges,omitempty"`
	LoggingEnabled bool `json:"logging_enabled,omitempty"`
	Direction FirewallDirection `json:"direction,omitempty"`
	Scope ScopeEnum `json:"scope,omitempty"`
	NsxManagers []Reference `json:"nsx_managers,omitempty"`
	AppliedTos []AppliedTo `json:"applied_tos,omitempty"`
}
