/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ApplicationFlowSummary struct for ApplicationFlowSummary
type ApplicationFlowSummary struct {
	ServiceEndpointCount int32 `json:"service_endpoint_count,omitempty"`
	Groups []GroupEntry `json:"groups,omitempty"`
	// Start timestamp of the window of the objects returned
	StartTime int64 `json:"start_time,omitempty"`
	// End timestamp of the window of the objects returned
	EndTime int64 `json:"end_time,omitempty"`
}