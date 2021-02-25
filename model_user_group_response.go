/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// UserGroupResponse struct for UserGroupResponse
type UserGroupResponse struct {
	GroupType UserGroupType `json:"group_type,omitempty"`
	// Unique identifier assigned to user-group.
	Id string `json:"id,omitempty"`
	// Group-name of the user-group.
	GroupName string `json:"group_name,omitempty"`
	// Domain name to which user-group belongs to.
	Domain string `json:"domain,omitempty"`
	Role Role `json:"role,omitempty"`
}
