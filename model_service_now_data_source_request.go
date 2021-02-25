/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ServiceNowDataSourceRequest struct for ServiceNowDataSourceRequest
type ServiceNowDataSourceRequest struct {
	// ID of Collector VM which should register this vcenter
	ProxyId string `json:"proxy_id"`
	// A friendly nickname for the data source
	Nickname string `json:"nickname"`
	// Whether or not data collection is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Room for notes or comments about the data source
	Notes string `json:"notes,omitempty"`
	Credentials PasswordCredentials `json:"credentials"`
	// Associated ServiceNow instance
	InstanceId string `json:"instance_id"`
	// CMDB configuration for CIs, relationships and graph traversal rules
	GraphConfiguration string `json:"graph_configuration,omitempty"`
	// Has graph configuration been modified from the default configuration
	IsGraphConfigCustomized bool `json:"is_graph_config_customized"`
	// Identifier of web proxy to be used for connection. If not given or empty, any web proxy used earlier will be cleared
	WebProxyId string `json:"web_proxy_id,omitempty"`
}
