/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// PortRange struct for PortRange
type PortRange struct {
	Start int32 `json:"start,omitempty"`
	End int32 `json:"end,omitempty"`
	Display string `json:"display,omitempty"`
	IanaName string `json:"iana_name,omitempty"`
	IanaPortDisplay string `json:"iana_port_display,omitempty"`
}