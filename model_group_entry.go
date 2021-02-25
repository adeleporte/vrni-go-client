/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// GroupEntry struct for GroupEntry
type GroupEntry struct {
	GroupId string `json:"group_id,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	FlowSummary FlowSummary `json:"flow_summary,omitempty"`
}
