# \PathApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PathFirewallRules**](PathApi.md#PathFirewallRules) | **Post** /path/firewall-rules | Get firewall rules for specified client server ips and port/protocol



## PathFirewallRules

> PathFirewallRules PathFirewallRules(ctx, body)

Get firewall rules for specified client server ips and port/protocol

Get firewall rules applicable in path for a client IP and server IP for specified port/protocol.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PathFirewallRulesRequest**](PathFirewallRulesRequest.md)| VMware Identity Manager configuration details | 

### Return type

[**PathFirewallRules**](PathFirewallRules.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

