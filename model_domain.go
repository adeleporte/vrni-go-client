/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// Domain struct for Domain
type Domain struct {
	DomainType string `json:"domain_type,omitempty"`
	// domain value, not required for LOCAL domain
	Value string `json:"value,omitempty"`
}
