/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// VidmUserGroupRequest struct for VidmUserGroupRequest
type VidmUserGroupRequest struct {
	// Specify group name (domain should not be part of group name).
	GroupName string `json:"group_name,omitempty"`
	// Provide domain name to which user-group belongs to.
	Domain string `json:"domain,omitempty"`
	Role Role `json:"role,omitempty"`
}
