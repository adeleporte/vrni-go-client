/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// FlowTag the model 'FlowTag'
type FlowTag string

// List of FlowTag
const (
	FLOWTAG_TRAFFIC_TYPE_UNKNOWN FlowTag = "TAG_TRAFFIC_TYPE_UNKNOWN"
	FLOWTAG_INTERNET_TRAFFIC FlowTag = "TAG_INTERNET_TRAFFIC"
	FLOWTAG_EAST_WEST_TRAFFIC FlowTag = "TAG_EAST_WEST_TRAFFIC"
	FLOWTAG_VM_VM_TRAFFIC FlowTag = "TAG_VM_VM_TRAFFIC"
	FLOWTAG_VM_PHY_TRAFFIC FlowTag = "TAG_VM_PHY_TRAFFIC"
	FLOWTAG_PHY_PHY_TRAFFIC FlowTag = "TAG_PHY_PHY_TRAFFIC"
	FLOWTAG_SRC_IP_VMKNIC FlowTag = "TAG_SRC_IP_VMKNIC"
	FLOWTAG_DST_IP_VMKNIC FlowTag = "TAG_DST_IP_VMKNIC"
	FLOWTAG_SRC_IP_VM FlowTag = "TAG_SRC_IP_VM"
	FLOWTAG_DST_IP_VM FlowTag = "TAG_DST_IP_VM"
	FLOWTAG_SRC_IP_INTERNET FlowTag = "TAG_SRC_IP_INTERNET"
	FLOWTAG_DST_IP_INTERNET FlowTag = "TAG_DST_IP_INTERNET"
	FLOWTAG_SRC_IP_PHYSICAL FlowTag = "TAG_SRC_IP_PHYSICAL"
	FLOWTAG_DST_IP_PHYSICAL FlowTag = "TAG_DST_IP_PHYSICAL"
	FLOWTAG_SAME_HOST FlowTag = "TAG_SAME_HOST"
	FLOWTAG_DIFF_HOST FlowTag = "TAG_DIFF_HOST"
	FLOWTAG_SHARED_SERVICE FlowTag = "TAG_SHARED_SERVICE"
	FLOWTAG_NOT_SHARED_SERVICE FlowTag = "TAG_NOT_SHARED_SERVICE"
	FLOWTAG_NETWORK_SWITCHED FlowTag = "TAG_NETWORK_SWITCHED"
	FLOWTAG_NETWORK_ROUTED FlowTag = "TAG_NETWORK_ROUTED"
	FLOWTAG_NETWORK_UNKNOWN FlowTag = "TAG_NETWORK_UNKNOWN"
	FLOWTAG_SRC_IP_VTEP FlowTag = "TAG_SRC_IP_VTEP"
	FLOWTAG_DST_IP_VTEP FlowTag = "TAG_DST_IP_VTEP"
	FLOWTAG_UNICAST FlowTag = "TAG_UNICAST"
	FLOWTAG_BROADCAST FlowTag = "TAG_BROADCAST"
	FLOWTAG_MULTICAST FlowTag = "TAG_MULTICAST"
	FLOWTAG_SRC_IP_LINK_LOCAL FlowTag = "TAG_SRC_IP_LINK_LOCAL"
	FLOWTAG_DST_IP_LINK_LOCAL FlowTag = "TAG_DST_IP_LINK_LOCAL"
	FLOWTAG_SRC_IP_CLASS_E FlowTag = "TAG_SRC_IP_CLASS_E"
	FLOWTAG_DST_IP_CLASS_E FlowTag = "TAG_DST_IP_CLASS_E"
	FLOWTAG_SRC_IP_CLASS_A_RESERVED FlowTag = "TAG_SRC_IP_CLASS_A_RESERVED"
	FLOWTAG_DST_IP_CLASS_A_RESERVED FlowTag = "TAG_DST_IP_CLASS_A_RESERVED"
	FLOWTAG_INVALID_IP_PACKETS FlowTag = "TAG_INVALID_IP_PACKETS"
	FLOWTAG_NOT_ANALYZED FlowTag = "TAG_NOT_ANALYZED"
	FLOWTAG_GENERIC_INTERNET_SRC_IP FlowTag = "TAG_GENERIC_INTERNET_SRC_IP"
	FLOWTAG_SNAT_DNAT_FLOW FlowTag = "TAG_SNAT_DNAT_FLOW"
	FLOWTAG_MULTINICS FlowTag = "TAG_MULTINICS"
	FLOWTAG_SRC_VC FlowTag = "TAG_SRC_VC"
	FLOWTAG_DST_VC FlowTag = "TAG_DST_VC"
	FLOWTAG_SRC_AWS FlowTag = "TAG_SRC_AWS"
	FLOWTAG_DST_AWS FlowTag = "TAG_DST_AWS"
	FLOWTAG_WITHIN_DC FlowTag = "TAG_WITHIN_DC"
	FLOWTAG_DIFF_DC FlowTag = "TAG_DIFF_DC"
	FLOWTAG_WITHIN_VPC FlowTag = "TAG_WITHIN_VPC"
	FLOWTAG_DIFF_VPC FlowTag = "TAG_DIFF_VPC"
	FLOWTAG_WITHIN_VNET FlowTag = "TAG_WITHIN_VNET"
	FLOWTAG_DIFF_VNET FlowTag = "TAG_DIFF_VNET"
	FLOWTAG_SRC_AZURE FlowTag = "TAG_SRC_AZURE"
	FLOWTAG_DST_AZURE FlowTag = "TAG_DST_AZURE"
)