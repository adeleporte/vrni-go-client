/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VCenterDataSourceAllOf struct for VCenterDataSourceAllOf
type VCenterDataSourceAllOf struct {
	Credentials PasswordCredentials `json:"credentials,omitempty"`
	// Whether vCenter is a VMware Cloud operated vCenter (true or false).
	IsVmc bool `json:"is_vmc,omitempty"`
}
