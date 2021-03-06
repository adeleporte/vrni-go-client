/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// EntityWithTime struct for EntityWithTime
type EntityWithTime struct {
	// Entity Identifier
	EntityId string `json:"entity_id,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Entity BaseEntity `json:"entity,omitempty"`
	// The timestamp of this object
	Time int64 `json:"time,omitempty"`
}
