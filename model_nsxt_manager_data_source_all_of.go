/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxtManagerDataSourceAllOf struct for NsxtManagerDataSourceAllOf
type NsxtManagerDataSourceAllOf struct {
	// Whether or not to configure NSX-T to send IPFIX to vRNI
	IpfixEnabled bool `json:"ipfix_enabled,omitempty"`
	// Whether or not to configure NSX-T to send virtual infrastructure latency metrics to vRNI
	LatencyEnabled bool `json:"latency_enabled,omitempty"`
	// Whether or not to configure NSX Intelligence to send additional traffic information to vRNI
	NsxiEnabled bool `json:"nsxi_enabled,omitempty"`
}
