/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SecurityTag struct for SecurityTag
type SecurityTag struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Description string `json:"description,omitempty"`
	DirectSecurityGroups []Reference `json:"direct_security_groups,omitempty"`
	SecurityGroups []Reference `json:"security_groups,omitempty"`
	VendorId string `json:"vendor_id,omitempty"`
	NsxManager Reference `json:"nsx_manager,omitempty"`
}
