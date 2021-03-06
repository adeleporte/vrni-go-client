/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// LicensingResponseResult struct for LicensingResponseResult
type LicensingResponseResult struct {
	Licenses []VrniLicense `json:"licenses,omitempty"`
	LicenseUsage [][]map[string]interface{} `json:"licenseUsage,omitempty"`
	// customer id
	CustomerId int32 `json:"customerId,omitempty"`
	// cpu capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period)
	CpuCapacity int32 `json:"cpuCapacity,omitempty"`
	// ccu capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period)
	CcuCapacity int32 `json:"ccuCapacity,omitempty"`
	// VM capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period)
	VmCapacity int32 `json:"vmCapacity,omitempty"`
	// core capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period)
	CoreCapacity int32 `json:"coreCapacity,omitempty"`
	// edge capacity of current non-expired (valid licenses) and older (invalid) licenses (during grace period)
	EdgeCapacity int32 `json:"edgeCapacity,omitempty"`
	// service tag
	ServiceTag string `json:"serviceTag,omitempty"`
	// Maximum of expiry date for current licenses, 0 if any non-expiry license is present. Invalid/already expired license is not considered here.
	ExpireDate int32 `json:"expireDate,omitempty"`
	// true if non-expired ENTERPRISE license is available
	Enterprise bool `json:"enterprise,omitempty"`
	// true if all licenses are evaluation license
	Evaluation bool `json:"evaluation,omitempty"`
	// number of total enterprise licenses available
	NumOfEnterpriseLicense int32 `json:"numOfEnterpriseLicense,omitempty"`
	// Total Number of valid licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidLicense int32 `json:"numOfValidLicense,omitempty"`
	// Number of valid Enterprise licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidEnterpriseLicense int32 `json:"numOfValidEnterpriseLicense,omitempty"`
	// Number of valid SDWAN licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidSDWANLicense int32 `json:"numOfValidSDWANLicense,omitempty"`
	// Number of valid NAV licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidNAVLicense int32 `json:"numOfValidNAVLicense,omitempty"`
	// Number of valid Advanced licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidAdvancedLicense int32 `json:"numOfValidAdvancedLicense,omitempty"`
	// Number of valid ROBO licenses available. Active licenses from last major version which are only valid for grace period are not counted
	NumOfValidROBOLicense int32 `json:"numOfValidROBOLicense,omitempty"`
	// Total Number of active licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveLicense int32 `json:"numOfActiveLicense,omitempty"`
	// Total Number of active Enterprise licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveEnterpriseLicense int32 `json:"numOfActiveEnterpriseLicense,omitempty"`
	// Total Number of active SDWAN licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveSDWANLicense int32 `json:"numOfActiveSDWANLicense,omitempty"`
	// Total Number of active NAV licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveNAVLicense int32 `json:"numOfActiveNAVLicense,omitempty"`
	// Total Number of active Advanced licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveAdvancedLicense int32 `json:"numOfActiveAdvancedLicense,omitempty"`
	// Total Number of active ROBO licenses (non-expired and licenses from last major version during grace period)
	NumOfActiveROBOLicense int32 `json:"numOfActiveROBOLicense,omitempty"`
	// Count of used sockets
	UsedSockets int32 `json:"usedSockets,omitempty"`
	// Count of used cores
	UsedCores int32 `json:"usedCores,omitempty"`
	// Count of used edges
	UsedEdges int32 `json:"usedEdges,omitempty"`
	// Count of used devices
	UsedDevices int32 `json:"usedDevices,omitempty"`
}
