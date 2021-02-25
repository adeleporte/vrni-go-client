/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxvManagerDataSourceRequestAllOf struct for NsxvManagerDataSourceRequestAllOf
type NsxvManagerDataSourceRequestAllOf struct {
	Credentials PasswordCredentials `json:"credentials"`
	// Associated vcenter data source entity ID
	VcenterId string `json:"vcenter_id"`
	// Whether or not to configure NSX-v to send IPFIX to vRNI
	IpfixEnabled bool `json:"ipfix_enabled,omitempty"`
	// Whether or not collection via the Central CLI for ESG data is enabled
	CentralCliEnabled bool `json:"central_cli_enabled,omitempty"`
	// Whether or not to configure NSX-v to send virtual infrastructure latency metrics to vRNI
	LatencyEnabled bool `json:"latency_enabled,omitempty"`
}
