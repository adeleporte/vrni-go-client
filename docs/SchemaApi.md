# \SchemaApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkFetchEventMetaInfo**](SchemaApi.md#BulkFetchEventMetaInfo) | **Get** /schema/problems | Get Event meta Information
[**GetMetricsSchema**](SchemaApi.md#GetMetricsSchema) | **Get** /schema/{entity-type}/metrics | Get metrics schema for an entity type



## BulkFetchEventMetaInfo

> EventMetaInfoResponse BulkFetchEventMetaInfo(ctx, )

Get Event meta Information

Bulk fetch of event meta info. Max batch size is 1000.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventMetaInfoResponse**](EventMetaInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsSchema

> EntityMetricsSchema GetMetricsSchema(ctx, entityType)

Get metrics schema for an entity type

Get details of metrics available for entity type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string**| entity type | 

### Return type

[**EntityMetricsSchema**](EntityMetricsSchema.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

