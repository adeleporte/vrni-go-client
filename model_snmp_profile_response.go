/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SnmpProfileResponse struct for SnmpProfileResponse
type SnmpProfileResponse struct {
	// Entity Identifier for SNMP profile.
	EntityId string `json:"entity_id,omitempty"`
	// User defined descriptor or identifier for particular SNMP profile.
	NickName string `json:"nick_name,omitempty"`
	// IP address of SNMP target destination
	TargetIp string `json:"target_ip,omitempty"`
	// Receiving port number of SNMP target destination
	TargetPort int32 `json:"target_port,omitempty"`
	SnmpVersion string `json:"snmp_version,omitempty"`
	SnmpV2c Snmp2cConfig `json:"snmp_v2c,omitempty"`
	SnmpV3 Snmp3Config `json:"snmp_v3,omitempty"`
}
