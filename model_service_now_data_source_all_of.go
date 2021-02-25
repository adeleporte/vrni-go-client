/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ServiceNowDataSourceAllOf struct for ServiceNowDataSourceAllOf
type ServiceNowDataSourceAllOf struct {
	// Associated nsxt data source entity Id
	InstanceId string `json:"instance_id,omitempty"`
	// CMDB configuration for CIs, relationships and graph traversal rules
	GraphConfiguration string `json:"graph_configuration,omitempty"`
	// Has graph configuration been modified from the default configuration
	IsGraphConfigCustomized bool `json:"is_graph_config_customized,omitempty"`
	// Identifier of web proxy to be used for connection
	WebProxyId string `json:"web_proxy_id,omitempty"`
}