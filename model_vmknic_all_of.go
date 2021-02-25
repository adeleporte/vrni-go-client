/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VmknicAllOf struct for VmknicAllOf
type VmknicAllOf struct {
	IpAddresses []IpV4Address `json:"ip_addresses,omitempty"`
	Vlan Vlan `json:"vlan,omitempty"`
	Host Reference `json:"host,omitempty"`
	Layer2Network Reference `json:"layer2_network,omitempty"`
}
