/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// DataSourceEntityId struct for DataSourceEntityId
type DataSourceEntityId struct {
	// Entity Identifier
	EntityId string `json:"entity_id,omitempty"`
	EntityType DataSourceType `json:"entity_type,omitempty"`
}