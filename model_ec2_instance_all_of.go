/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// Ec2InstanceAllOf struct for Ec2InstanceAllOf
type Ec2InstanceAllOf struct {
	Vpc Reference `json:"vpc,omitempty"`
	Region string `json:"region,omitempty"`
	Account Reference `json:"account,omitempty"`
}
