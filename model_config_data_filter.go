/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ConfigDataFilter struct for ConfigDataFilter
type ConfigDataFilter struct {
	Snmp bool `json:"snmp,omitempty"`
	Smtp bool `json:"smtp,omitempty"`
	WebProxy bool `json:"web_proxy,omitempty"`
	DataSources bool `json:"data_sources,omitempty"`
	Events bool `json:"events,omitempty"`
	Syslog bool `json:"syslog,omitempty"`
	Ldap bool `json:"ldap,omitempty"`
	Vidm bool `json:"vidm,omitempty"`
	UserData bool `json:"user_data,omitempty"`
	PhysicalSubnetVlan bool `json:"physical_subnet_vlan,omitempty"`
	PhysicalIpDnsMapping bool `json:"physical_ip_dns_mapping,omitempty"`
	SystemConfiguration bool `json:"system_configuration,omitempty"`
	EastWestIp bool `json:"east_west_ip,omitempty"`
	NorthSouthIp bool `json:"north_south_ip,omitempty"`
	DataManagement bool `json:"data_management,omitempty"`
	OnlineUpdateStatus bool `json:"online_update_status,omitempty"`
	CeipStatus bool `json:"ceip_status,omitempty"`
	AuditLogsPiiStatus bool `json:"audit_logs_pii_status,omitempty"`
}
