# \LogsApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](LogsApi.md#GetAuditLogs) | **Post** /logs/audit | Get Audit logs



## GetAuditLogs

> AuditListResponse GetAuditLogs(ctx, body)

Get Audit logs

get audit logs based on filters specified in query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AuditRequest**](AuditRequest.md)| Query Filters | 

### Return type

[**AuditListResponse**](AuditListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

