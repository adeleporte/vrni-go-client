/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VendorInfo struct for VendorInfo
type VendorInfo struct {
	VendorIds []VendorId `json:"vendor_ids,omitempty"`
	Manager Reference `json:"manager,omitempty"`
}