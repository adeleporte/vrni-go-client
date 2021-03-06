/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VidmUserRequest struct for VidmUserRequest
type VidmUserRequest struct {
	// Provide username (domain should not be part of username).
	Username string `json:"username,omitempty"`
	// Provide domain name to which user belongs to.
	Domain string `json:"domain,omitempty"`
	// Provide user's display name (could be \"givenName familyName\")
	DisplayName string `json:"display_name,omitempty"`
	Role Role `json:"role,omitempty"`
}
