# \SearchApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateSearchResults**](SearchApi.md#AggregateSearchResults) | **Post** /search/aggregation | Search Aggregation API
[**GroupSearchResults**](SearchApi.md#GroupSearchResults) | **Post** /search/groupby | Search Group By API
[**SearchEntities**](SearchApi.md#SearchEntities) | **Post** /search | Search entities



## AggregateSearchResults

> AggregationResponse AggregateSearchResults(ctx, optional)

Search Aggregation API

Using aggregate API you can aggregate search results for vRealize Network Insight entities by specifying entity type, filter expression and aggregate clause. Please refer to API Guide on details of how to construct filter expression and aggregate clause. A successful search request will return a list of aggregations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AggregateSearchResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AggregateSearchResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AggregationRequest**](AggregationRequest.md)| Aggregation Request | 

### Return type

[**AggregationResponse**](AggregationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSearchResults

> SearchGroupByResponse GroupSearchResults(ctx, optional)

Search Group By API

Using groupby search API you can group search results for vRealize Network Insight entities by specifying entity type, filter expression, aggregate clause and groupby clause. Please refer to API Guide on details of how to construct filter expression, aggregate clause and groupby clause. A successful search request will return a list of groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupSearchResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSearchResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SearchGroupByRequest**](SearchGroupByRequest.md)| GroupBy Request | 

### Return type

[**SearchGroupByResponse**](SearchGroupByResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEntities

> PagedListResponseWithTime SearchEntities(ctx, optional)

Search entities

Using search API you can search vRealize Network Insight entities by specifying entity type and filter expression. A filter expression is a predicate expression (similar to SQL where clause) used to define the search criteria. Please refer to API Guide on details of how to construct filter expression. A successful search request will return a list of entity ids that matches the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchEntitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SearchRequest**](SearchRequest.md)| Search Request | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

