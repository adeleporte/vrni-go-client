/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// Token struct for Token
type Token struct {
	Token string `json:"token,omitempty"`
	// expiry epoch time in secs.
	Expiry int64 `json:"expiry,omitempty"`
}