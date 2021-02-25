/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VxlanLayer2NetworkAllOf struct for VxlanLayer2NetworkAllOf
type VxlanLayer2NetworkAllOf struct {
	SegmentId int32 `json:"segment_id,omitempty"`
	Vteps []Reference `json:"vteps,omitempty"`
	Scope ScopeEnum `json:"scope,omitempty"`
	NsxManagers []Reference `json:"nsx_managers,omitempty"`
}