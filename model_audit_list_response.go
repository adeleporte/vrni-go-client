/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// AuditListResponse struct for AuditListResponse
type AuditListResponse struct {
	Results []AuditResponse `json:"results,omitempty"`
	// Cursor for the next page of logs
	Cursor string `json:"cursor,omitempty"`
	// Start timestamp of the window of the objects returned
	StartTime int64 `json:"start_time,omitempty"`
	// End timestamp of the window of the objects returned
	EndTime int64 `json:"end_time,omitempty"`
}
