/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// BaseVirtualMachine struct for BaseVirtualMachine
type BaseVirtualMachine struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	IpAddresses []IpV4Address `json:"ip_addresses,omitempty"`
	// Default gateway IP of the VM
	DefaultGateway string `json:"default_gateway,omitempty"`
	Vnics []Reference `json:"vnics,omitempty"`
	SecurityGroups []Reference `json:"security_groups,omitempty"`
	SourceFirewallRules []RuleSet `json:"source_firewall_rules,omitempty"`
	DestinationFirewallRules []RuleSet `json:"destination_firewall_rules,omitempty"`
	IpSets []Reference `json:"ip_sets,omitempty"`
}
