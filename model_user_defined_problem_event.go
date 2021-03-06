/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// UserDefinedProblemEvent struct for UserDefinedProblemEvent
type UserDefinedProblemEvent struct {
	// Entity ID that can be references in detail API calls
	EntityId string `json:"entity_id,omitempty"`
	// Name of the object
	Name string `json:"name,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	AnchorEntities []Reference `json:"anchor_entities,omitempty"`
	// The entity IDs of all related objects
	RelatedEntities []Reference `json:"related_entities,omitempty"`
	// Event message
	Message string `json:"message,omitempty"`
	// Event tags
	EventTags []string `json:"event_tags,omitempty"`
	// Administrative state of the event
	AdminState string `json:"admin_state,omitempty"`
	// Whether of not the event is archived
	Archived bool `json:"archived,omitempty"`
	// Epoc timestamp of when the event was triggered
	EventTimeEpochMs int64 `json:"event_time_epoch_ms,omitempty"`
	// The type of event
	EventType string `json:"event_type,omitempty"`
	// A list of recommended remedies
	Recommendations []string `json:"recommendations,omitempty"`
	Severity string `json:"severity,omitempty"`
	EventCloseTimeEpochMs int64 `json:"event_close_time_epoch_ms,omitempty"`
	TotalCount int64 `json:"total_count,omitempty"`
	RefsAddedCount int64 `json:"refs_added_count,omitempty"`
	RefsDeletedCount int64 `json:"refs_deleted_count,omitempty"`
	Subscriptions []Reference `json:"subscriptions,omitempty"`
}
