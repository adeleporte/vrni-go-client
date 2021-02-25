/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// Host struct for Host
type Host struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	Vmknics []Reference `json:"vmknics,omitempty"`
	Cluster Reference `json:"cluster,omitempty"`
	VcenterManager Reference `json:"vcenter_manager,omitempty"`
	// Number of VMs running on the host
	VmCount int32 `json:"vm_count,omitempty"`
	Datastores []Reference `json:"datastores,omitempty"`
	// Hardware service tag of the host
	ServiceTag string `json:"service_tag,omitempty"`
	// Internal vCenter ID of the host
	VendorId string `json:"vendor_id,omitempty"`
	NsxManager Reference `json:"nsx_manager,omitempty"`
	// Is the host in maintenance mode?
	MaintenanceMode string `json:"maintenance_mode,omitempty"`
	// Connection state of the host
	ConnectionState string `json:"connection_state,omitempty"`
}
