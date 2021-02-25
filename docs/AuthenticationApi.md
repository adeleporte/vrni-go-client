# \AuthenticationApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](AuthenticationApi.md#Create) | **Post** /auth/token | Create an auth token
[**CreateVidmUserToken**](AuthenticationApi.md#CreateVidmUserToken) | **Post** /auth/token/vidm | Create an auth token for user mapped through VMware Identity Manager
[**Delete**](AuthenticationApi.md#Delete) | **Delete** /auth/token | Delete an auth token.
[**GetVidmOauthClienId**](AuthenticationApi.md#GetVidmOauthClienId) | **Get** /auth/vidm/client-id | Get the client-id for making user access-token request to VMware Identity Manager



## Create

> Token Create(ctx, body)

Create an auth token

<html><body> vRealize Network Insight supports token based authentication. Tokens are non-modifiable identifiers returned by the system when the user has successfully authenticated using valid credentials. Token expires after expiry time returned in the response. All API requests must provide the auth token in Authorization header in following format:<br> Authorization &#58; NetworkInsight {token} <br> If a token is invalid or expired, 401-Unauthorized error gets returned in the response of the API request. <br> There is limit of 100 valid tokens per user and further requests will return 401-Unauthorized. So, users are advised to delete the tokens after use <br> Expired tokens are cleaned periodically by the system. </body></html>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserCredential**](UserCredential.md)| User Credentials | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, token, expiry

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVidmUserToken

> Token CreateVidmUserToken(ctx, body)

Create an auth token for user mapped through VMware Identity Manager

<html><body> vRealize Network Insight supports token based authentication. Tokens are non-modifiable identifiers returned by the system when the user has been successfully authenticated using valid access token provided by configured VMware Identity Manager appliance. Token expires after expiry time (returned in the response). All API requests must provide the auth token in Authorization header in following format:<br> Authorization &#58; NetworkInsight {token} <br> If a token is invalid or expired, 401-Unauthorized error gets returned in the response of the API request. </body></html>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VidmToken**](VidmToken.md)| User access token provided by VMware Identity Manager. | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, )

Delete an auth token.

Deletes the auth token provided in Authorization header. Deleting an expired or invalid token will result in 401 Unauthorized error.

### Required Parameters

This endpoint does not need any parameter.

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


## GetVidmOauthClienId

> VidmOauthClientResponse GetVidmOauthClienId(ctx, )

Get the client-id for making user access-token request to VMware Identity Manager

Get client-id of password grant OAuth client registered at VMware Identity Manager. This client-id along with user credentials is required while making an access token request to VMware Identity Manager.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VidmOauthClientResponse**](VidmOauthClientResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

