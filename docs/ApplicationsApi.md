# \ApplicationsApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](ApplicationsApi.md#AddApplication) | **Post** /groups/applications | Create an application
[**AddTier**](ApplicationsApi.md#AddTier) | **Post** /groups/applications/{id}/tiers | Create tier in application
[**DeleteApplication**](ApplicationsApi.md#DeleteApplication) | **Delete** /groups/applications/{id} | Delete an application
[**DeleteTier**](ApplicationsApi.md#DeleteTier) | **Delete** /groups/applications/{id}/tiers/{tier-id} | Delete application tier
[**EditApplicationTier**](ApplicationsApi.md#EditApplicationTier) | **Put** /groups/applications/{id}/tiers/{tier-id} | Edit tier in application
[**GetAppFlowMetrics**](ApplicationsApi.md#GetAppFlowMetrics) | **Get** /groups/applications/{id}/flow-metrics | Get application flow metrics
[**GetAppFlowProperties**](ApplicationsApi.md#GetAppFlowProperties) | **Get** /groups/applications/{id}/flow-props | Get application flow properties
[**GetAppProblems**](ApplicationsApi.md#GetAppProblems) | **Get** /groups/applications/{id}/problems | List Application problems
[**GetAppTopTalkingMembers**](ApplicationsApi.md#GetAppTopTalkingMembers) | **Get** /groups/applications/{id}/top-talking-member | List Application top talking members
[**GetAppTopTalkingPairs**](ApplicationsApi.md#GetAppTopTalkingPairs) | **Get** /groups/applications/{id}/top-talking-pair | List Application top talking pairs
[**GetApplication**](ApplicationsApi.md#GetApplication) | **Get** /groups/applications/{id} | Show application details
[**GetApplicationFlowSummary**](ApplicationsApi.md#GetApplicationFlowSummary) | **Get** /groups/applications/{id}/flow-summary | Show application flow summary
[**GetApplicationTier**](ApplicationsApi.md#GetApplicationTier) | **Get** /groups/applications/{id}/tiers/{tier-id} | Show application tier details
[**GetApplicationVms**](ApplicationsApi.md#GetApplicationVms) | **Get** /groups/applications/{id}/members/vms | Show application members
[**GetApplicationsMembers**](ApplicationsApi.md#GetApplicationsMembers) | **Post** /groups/applications/members | Get member details of applications
[**GetTier**](ApplicationsApi.md#GetTier) | **Get** /groups/tiers/{tier-id} | Show application tier details
[**GetTiersMembers**](ApplicationsApi.md#GetTiersMembers) | **Post** /groups/tiers/members | Get member details of tiers
[**ListApplicationTiers**](ApplicationsApi.md#ListApplicationTiers) | **Get** /groups/applications/{id}/tiers | List tiers of an application
[**ListApplications**](ApplicationsApi.md#ListApplications) | **Get** /groups/applications | List applications
[**ListApplicationsDetails**](ApplicationsApi.md#ListApplicationsDetails) | **Get** /groups/applications/fetch | Get application details in bulk



## AddApplication

> Application AddApplication(ctx, body)

Create an application

Application is a group of tiers. A tier is a group of virtual machines or IP addresses based on membership criteria. Tiers are bound to a single application. An application name is unique and should not conflict with another application name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ApplicationRequest**](ApplicationRequest.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, entity_id, name, entity_type, create_time, created_by, last_modified_time, last_modified_by

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTier

> Tier AddTier(ctx, id, body)

Create tier in application

Create a tier of an application by with specified membership criteria. The membership criteria id defined in terms of virtual machines or IP addresses/subnet. Please refer to API Guide on how to construct membership criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**TierRequest**](TierRequest.md)|  | 

### Return type

[**Tier**](Tier.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, entity_id, name, entity_type, group_membership_criteria, application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, id)

Delete an application

Deleting an application deletes all the tiers of the application along with the application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTier

> DeleteTier(ctx, id, tierId)

Delete application tier

Delete application tier of an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**tierId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditApplicationTier

> Tier EditApplicationTier(ctx, id, tierId, body)

Edit tier in application

Edit a tier of an application with specified membership criteria. The membership criteria id defined in terms of virtual machines or IP addresses/subnet. Please refer to API Guide on how to construct membership criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**tierId** | **string**|  | 
**body** | [**TierRequest**](TierRequest.md)|  | 

### Return type

[**Tier**](Tier.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, entity_id, name, entity_type, group_membership_criteria, application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppFlowMetrics

> ApplicationFlowData GetAppFlowMetrics(ctx, id, metrics, optional)

Get application flow metrics

Get application flow properties sum of bytes of incoming and outgoing flows

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**metrics** | [**[]AppFlowMetricEnum**](AppFlowMetricEnum.md)| List of metrics to include | 
 **optional** | ***GetAppFlowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppFlowMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**ApplicationFlowData**](ApplicationFlowData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppFlowProperties

> ApplicationFlowData GetAppFlowProperties(ctx, id, props, optional)

Get application flow properties

Get application flow properties e.g flow count, count of interent flows etc

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**props** | [**[]AppFlowPropEnum**](AppFlowPropEnum.md)| List of properties to include | 
 **optional** | ***GetAppFlowPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppFlowPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**ApplicationFlowData**](ApplicationFlowData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppProblems

> PagedListResponseWithTime GetAppProblems(ctx, id, optional)

List Application problems

List Application problem events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetAppProblemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppProblemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 
 **eventType** | **optional.String**| Type of event, e.g UserDefinedProblemEvent | 
 **eventTags** | [**optional.Interface of []string**](string.md)| Event tags | 
 **eventStatus** | **optional.String**| Status of event open or closed | [default to all]
 **updateTimeFrom** | **optional.Float32**| Events that were created or updated or closed between this time and update_time_to, in seconds | 
 **updateTimeTo** | **optional.Float32**| Events that were created or updated or closed between update_time_from and this time, in seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppTopTalkingMembers

> ApplicationTopTalkingMemberData GetAppTopTalkingMembers(ctx, id, sortCriteria, optional)

List Application top talking members

List Application top talking members based on provided criteria

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**sortCriteria** | [**[]TopTalkerSortEnum**](TopTalkerSortEnum.md)| Sorting criteria | 
 **optional** | ***GetAppTopTalkingMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppTopTalkingMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**ApplicationTopTalkingMemberData**](ApplicationTopTalkingMemberData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppTopTalkingPairs

> ApplicationTopTalkingPairData GetAppTopTalkingPairs(ctx, id, sortCriteria, optional)

List Application top talking pairs

List Application top talking pairs based on provided criteria

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**sortCriteria** | [**[]TopTalkerSortEnum**](TopTalkerSortEnum.md)| Sorting criteria | 
 **optional** | ***GetAppTopTalkingPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppTopTalkingPairsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**ApplicationTopTalkingPairData**](ApplicationTopTalkingPairData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, id)

Show application details

Show application details for an entity ID. This returns the application name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**Application**](Application.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, create_time, created_by, last_modified_time, last_modified_by

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationFlowSummary

> ApplicationFlowSummary GetApplicationFlowSummary(ctx, id, optional)

Show application flow summary

Show application details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetApplicationFlowSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationFlowSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**ApplicationFlowSummary**](ApplicationFlowSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationTier

> Tier GetApplicationTier(ctx, id, tierId)

Show application tier details

Show application tier details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**tierId** | **string**|  | 

### Return type

[**Tier**](Tier.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, group_membership_criteria, application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVms

> ApplicationMembers GetApplicationVms(ctx, id)

Show application members

Show application members

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**ApplicationMembers**](ApplicationMembers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsMembers

> ApplicationsMembersResponse GetApplicationsMembers(ctx, body)

Get member details of applications

Get member details of applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GroupsMembersRequest**](GroupsMembersRequest.md)|  | 

### Return type

[**ApplicationsMembersResponse**](ApplicationsMembersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTier

> Tier GetTier(ctx, tierId, authorization)

Show application tier details

Show application tier details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tierId** | **string**|  | 
**authorization** | **string**| Authorization Header | 

### Return type

[**Tier**](Tier.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, group_membership_criteria, application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTiersMembers

> TiersMembersResponse GetTiersMembers(ctx, body)

Get member details of tiers

Get member details of tiers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GroupsMembersRequest**](GroupsMembersRequest.md)|  | 

### Return type

[**TiersMembersResponse**](TiersMembersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTiers

> TierListResponse ListApplicationTiers(ctx, id)

List tiers of an application

List tiers of an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**TierListResponse**](TierListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> PagedListResponse ListApplications(ctx, optional)

List applications

List applications in vRealize Network Insight

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **modifiedAfter** | **optional.Float32**| timestamp after which apps has been modified | 

### Return type

[**PagedListResponse**](PagedListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, total_count

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationsDetails

> PagedApplicationListResponse ListApplicationsDetails(ctx, optional)

Get application details in bulk

Get the details of applications in bulk, by providing a list of entity IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListApplicationsDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationsDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **modifiedAfter** | **optional.Float32**| timestamp after which apps has been modified | 

### Return type

[**PagedApplicationListResponse**](PagedApplicationListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

