/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// DataSourceListResponse struct for DataSourceListResponse
type DataSourceListResponse struct {
	Results []DataSourceEntityId `json:"results,omitempty"`
	TotalCount int32 `json:"total_count,omitempty"`
}
