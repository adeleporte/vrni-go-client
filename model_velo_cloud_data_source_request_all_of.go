/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VeloCloudDataSourceRequestAllOf struct for VeloCloudDataSourceRequestAllOf
type VeloCloudDataSourceRequestAllOf struct {
	Credentials PasswordCredentials `json:"credentials"`
	// Identifier of web proxy to be used for connection. If not given or empty, any web proxy used earlier will be cleared
	WebProxyId string `json:"web_proxy_id,omitempty"`
}
