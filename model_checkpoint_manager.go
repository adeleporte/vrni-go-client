/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// CheckpointManager struct for CheckpointManager
type CheckpointManager struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	NsxManager Reference `json:"nsx_manager,omitempty"`
	Version string `json:"version,omitempty"`
	IpAddress IpV4Address `json:"ip_address,omitempty"`
}
