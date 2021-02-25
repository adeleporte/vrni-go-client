/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ValidateConnectionViaWebProxyResponse struct for ValidateConnectionViaWebProxyResponse
type ValidateConnectionViaWebProxyResponse struct {
	// Entity ID for a connected client
	EntityId string `json:"entity_id,omitempty"`
	// Type of Entity
	EntityType string `json:"entity_type,omitempty"`
	// Name of connected client
	Name string `json:"name,omitempty"`
	// Identifier/AccessKey/IP Adress of connected client
	Identifier string `json:"identifier,omitempty"`
	// Validation status of a connected entity with the updated web proxy details
	Status string `json:"status,omitempty"`
	// Error code in case validation failed
	ErrorCode int32 `json:"error_code,omitempty"`
	// Error message in case validation failed
	Message string `json:"message,omitempty"`
}
