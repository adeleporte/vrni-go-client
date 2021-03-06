/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// NsxvManagerAllOf struct for NsxvManagerAllOf
type NsxvManagerAllOf struct {
	Fqdn string `json:"fqdn,omitempty"`
	IpAddress IpV4Address `json:"ip_address,omitempty"`
	Version string `json:"version,omitempty"`
	PrimaryNsxManager Reference `json:"primary_nsx_manager,omitempty"`
	Vm Reference `json:"vm,omitempty"`
	Role string `json:"role,omitempty"`
}
