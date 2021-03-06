/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// MetricSchema struct for MetricSchema
type MetricSchema struct {
	Metric string `json:"metric,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Intervals []int32 `json:"intervals,omitempty"`
	Description string `json:"description,omitempty"`
	Unit string `json:"unit,omitempty"`
}
