/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// PagedEntityPairListResponse struct for PagedEntityPairListResponse
type PagedEntityPairListResponse struct {
	// Array of a source and destination pair
	Results []EntityPair `json:"results,omitempty"`
	// Cursor for the next page
	Cursor string `json:"cursor,omitempty"`
	// Total number of objects in the system, despite the page limit
	TotalCount int32 `json:"total_count,omitempty"`
	// Start timestamp of the window of the objects returned
	StartTime int64 `json:"start_time,omitempty"`
	// End timestamp of the window of the objects returned
	EndTime int64 `json:"end_time,omitempty"`
}