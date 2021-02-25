/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// F5BigipDataSourceRequest struct for F5BigipDataSourceRequest
type F5BigipDataSourceRequest struct {
	// IP address of data source (use either IP or FQDN, not both)
	Ip string `json:"ip,omitempty"`
	// Hostname of data source (use either IP or FQDN, not both)
	Fqdn string `json:"fqdn,omitempty"`
	// ID of Collector VM which should register this vcenter
	ProxyId string `json:"proxy_id"`
	// A friendly nickname for the data source
	Nickname string `json:"nickname"`
	// Whether or not data collection is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Room for notes or comments about the data source
	Notes string `json:"notes,omitempty"`
	Credentials PasswordCredentials `json:"credentials,omitempty"`
}
