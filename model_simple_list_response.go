/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SimpleListResponse struct for SimpleListResponse
type SimpleListResponse struct {
	Results []EntityId `json:"results,omitempty"`
	TotalCount int32 `json:"total_count,omitempty"`
}
