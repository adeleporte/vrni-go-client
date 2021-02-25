/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ServiceNowSourceRequest struct for ServiceNowSourceRequest
type ServiceNowSourceRequest struct {
	// ID of Collector VM which should register this vcenter
	ProxyId string `json:"proxy_id"`
	// A friendly nickname for the data source
	Nickname string `json:"nickname"`
	// Whether or not data collection is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Room for notes or comments about the data source
	Notes string `json:"notes,omitempty"`
}
