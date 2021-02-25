/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SearchRequest struct for SearchRequest
type SearchRequest struct {
	EntityType AllEntityType `json:"entity_type,omitempty"`
	// query filter
	Filter string `json:"filter,omitempty"`
	SortBy SortByClause `json:"sort_by,omitempty"`
	Size int32 `json:"size,omitempty"`
	Cursor string `json:"cursor,omitempty"`
	TimeRange TimeRange `json:"time_range,omitempty"`
}
