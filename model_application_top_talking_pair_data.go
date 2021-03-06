/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// ApplicationTopTalkingPairData struct for ApplicationTopTalkingPairData
type ApplicationTopTalkingPairData struct {
	// The entity ID of the application this flow is attached
	ApplicationId string `json:"application_id,omitempty"`
	// List of objects that are the top talkers
	TopTalkerMember []TopTalkingPairData `json:"top_talker_member,omitempty"`
}
