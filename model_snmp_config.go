/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SnmpConfig struct for SnmpConfig
type SnmpConfig struct {
	// Enable SNMP stats collection?
	SnmpEnabled bool `json:"snmp_enabled,omitempty"`
	// SNMP version to use
	SnmpVersion string `json:"snmp_version,omitempty"`
	ConfigSnmp2c Snmp2cConfig `json:"config_snmp_2c,omitempty"`
	ConfigSnmp3 Snmp3Config `json:"config_snmp_3,omitempty"`
}
