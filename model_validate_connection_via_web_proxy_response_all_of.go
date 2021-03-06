/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ValidateConnectionViaWebProxyResponseAllOf struct for ValidateConnectionViaWebProxyResponseAllOf
type ValidateConnectionViaWebProxyResponseAllOf struct {
	// Validation status of a connected entity with the updated web proxy details
	Status string `json:"status,omitempty"`
	// Error code in case validation failed
	ErrorCode int32 `json:"error_code,omitempty"`
	// Error message in case validation failed
	Message string `json:"message,omitempty"`
}
