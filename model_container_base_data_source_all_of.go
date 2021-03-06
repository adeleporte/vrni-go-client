/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ContainerBaseDataSourceAllOf struct for ContainerBaseDataSourceAllOf
type ContainerBaseDataSourceAllOf struct {
	// Associated NSX-T data source entity ID
	ManagerId string `json:"manager_id,omitempty"`
	Credentials K8SCredentials `json:"credentials,omitempty"`
}
