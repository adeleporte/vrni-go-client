/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// FetchRequestEntityIds struct for FetchRequestEntityIds
type FetchRequestEntityIds struct {
	// Entity Identifier
	EntityId string `json:"entity_id,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Time int64 `json:"time,omitempty"`
}
