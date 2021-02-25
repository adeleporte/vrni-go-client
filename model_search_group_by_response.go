/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SearchGroupByResponse struct for SearchGroupByResponse
type SearchGroupByResponse struct {
	Results []GroupWithValue `json:"results,omitempty"`
	Size int32 `json:"size,omitempty"`
	TotalBucket Cardinality `json:"total_bucket,omitempty"`
	// Total count of objects returned
	TotalCount int32 `json:"total_count,omitempty"`
	TimeRange TimeRange `json:"time_range,omitempty"`
	// Cursor for the next page
	Cursor string `json:"cursor,omitempty"`
}
