/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// BaseFirewallAllOf struct for BaseFirewallAllOf
type BaseFirewallAllOf struct {
	FirewallRules []RuleSet `json:"firewall_rules,omitempty"`
	Exclusions []Reference `json:"exclusions,omitempty"`
}
