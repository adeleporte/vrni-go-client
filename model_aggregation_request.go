/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// AggregationRequest struct for AggregationRequest
type AggregationRequest struct {
	EntityType AllEntityType `json:"entity_type,omitempty"`
	// query filter
	Filter string `json:"filter,omitempty"`
	Aggregations []Aggregation `json:"aggregations,omitempty"`
	TimeRange TimeRange `json:"time_range,omitempty"`
}
