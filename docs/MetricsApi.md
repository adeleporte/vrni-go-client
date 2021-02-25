# \MetricsApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /metrics | Get metric points for an entity



## GetMetrics

> MetricResponse GetMetrics(ctx, entityId, metric, interval, start, end)

Get metric points for an entity

Get metric points for an entity for an entity id and metric for a given time interval. Maximum number of metrics point returned by API is 300. In case the interval and time period combination have more than 300 metrics points, client should break the time period to multiple batches to get all the metrics points. These metric points are the points inside the metric charts (CPU, network rate, etc.)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string**| entity type | 
**metric** | **string**| metric name | 
**interval** | **int32**| metric points interval | 
**start** | **int64**| start time for query in epoch seconds | 
**end** | **int64**| end time for query in epoch seconds | 

### Return type

[**MetricResponse**](MetricResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

