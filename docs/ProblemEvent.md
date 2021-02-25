# ProblemEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID that can be references in detail API calls | [optional] 
**Name** | **string** | Name of the object | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | [optional] 
**AnchorEntities** | [**[]Reference**](Reference.md) |  | [optional] 
**RelatedEntities** | [**[]Reference**](Reference.md) | The entity IDs of all related objects | [optional] 
**Message** | **string** | Event message | [optional] 
**EventTags** | **[]string** | Event tags | [optional] 
**AdminState** | **string** | Administrative state of the event | [optional] 
**Archived** | **bool** | Whether of not the event is archived | [optional] 
**EventTimeEpochMs** | **int64** | Epoc timestamp of when the event was triggered | [optional] 
**EventType** | **string** | The type of event | [optional] 
**Recommendations** | **[]string** | A list of recommended remedies | [optional] 
**Severity** | **string** |  | [optional] 
**EventCloseTimeEpochMs** | **int64** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


