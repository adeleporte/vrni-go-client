# \MicrosegmentationApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportNsxRecommendedRules**](MicrosegmentationApi.md#ExportNsxRecommendedRules) | **Post** /micro-seg/recommended-rules/nsx | Export recommended firewall rules for NSX-v
[**ListRecommendedRules**](MicrosegmentationApi.md#ListRecommendedRules) | **Post** /micro-seg/recommended-rules | Get logical recommended firewall rules



## ExportNsxRecommendedRules

> *os.File ExportNsxRecommendedRules(ctx, optional)

Export recommended firewall rules for NSX-v

Export recommended firewall rules based on the flow data gathered by vRealize Network Insight in NSX-V compatible format. The output will be a .zip file download.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportNsxRecommendedRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportNsxRecommendedRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RecommendedRulesRequest**](RecommendedRulesRequest.md)| NSX Recommended Firewall Rules Request | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecommendedRules

> RecommendedRules ListRecommendedRules(ctx, optional)

Get logical recommended firewall rules

Get recommended firewall rules based on the flow data gathered by vRealize Network Insight. This API provides service to retrieve recommended rules based on flow traffic that is observed between two groups OR for a single group based on all the inbound and outboud traffic for that group. In case two groups are provided, both the groups should be of same type. Currently supported groups are Application, Tier, NSXSecurityGroup, EC2SecurityGroup.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRecommendedRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecommendedRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RecommendedRulesRequest**](RecommendedRulesRequest.md)| Recommended Rules Request | 

### Return type

[**RecommendedRules**](RecommendedRules.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, results, time_range

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

