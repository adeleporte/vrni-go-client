# \SettingsApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSerialNumber**](SettingsApi.md#ActivateSerialNumber) | **Post** /settings/licensing/activate | Activates a valid license key. This API is not applicable for on-boarding.
[**AddBackupConfig**](SettingsApi.md#AddBackupConfig) | **Post** /settings/backup | Configure backup of configuration
[**AddIpTag**](SettingsApi.md#AddIpTag) | **Post** /settings/ip-tags/{tag-id}/add | Tag IP addresses with tag-id
[**AddRestoreConfig**](SettingsApi.md#AddRestoreConfig) | **Post** /settings/restore | Configure restore of configuration and triggers restore operation
[**AddVidmUser**](SettingsApi.md#AddVidmUser) | **Post** /settings/users/vidm | Add a VMware Identity manager user to vRealize Network Insight
[**AddVidmUserGroup**](SettingsApi.md#AddVidmUserGroup) | **Post** /settings/user-groups/vidm | Add a VMware Identity Manager user-group to vRealize Network Insight
[**AddWebProxy**](SettingsApi.md#AddWebProxy) | **Post** /settings/proxy-servers | Add new Web Proxy in the system
[**CreateSubnetMapping**](SettingsApi.md#CreateSubnetMapping) | **Post** /settings/subnet-mappings | Create subnet mapping
[**CreateUserDefinedEvent**](SettingsApi.md#CreateUserDefinedEvent) | **Post** /settings/events/user-defined-events | Add new User-Defined event
[**DeactivateSerialNumber**](SettingsApi.md#DeactivateSerialNumber) | **Delete** /settings/licensing/deactivate | Deactivates an existing license key
[**DeleteBackupConfig**](SettingsApi.md#DeleteBackupConfig) | **Delete** /settings/backup | Delete existing Backup configuration
[**DeleteRestoreConfig**](SettingsApi.md#DeleteRestoreConfig) | **Delete** /settings/restore | Delete existing restore configuration
[**DeleteSubnetMapping**](SettingsApi.md#DeleteSubnetMapping) | **Delete** /settings/subnet-mappings/{id} | Delete subnet mapping
[**DeleteUser**](SettingsApi.md#DeleteUser) | **Delete** /settings/users/{id} | Delete an existing user.
[**DeleteUserDefinedEvent**](SettingsApi.md#DeleteUserDefinedEvent) | **Delete** /settings/events/user-defined-events/{id} | Delete an existing User-Defined event
[**DeleteUserGroup**](SettingsApi.md#DeleteUserGroup) | **Delete** /settings/user-groups/{id} | Delete an existing user-group
[**DeleteVidmConfiguration**](SettingsApi.md#DeleteVidmConfiguration) | **Delete** /settings/vidm | Delete VMware Identity Manager configuration
[**DeleteWebProxy**](SettingsApi.md#DeleteWebProxy) | **Delete** /settings/proxy-servers/{id} | Delete an existing Web Proxy server
[**DisableUserDefinedEvent**](SettingsApi.md#DisableUserDefinedEvent) | **Post** /settings/events/user-defined-events/{id}/disable | Disable an existing User-Defined event
[**DisableVidm**](SettingsApi.md#DisableVidm) | **Post** /settings/vidm/disable | Disable VMware Identity Manager integration
[**EnableUserDefinedEvent**](SettingsApi.md#EnableUserDefinedEvent) | **Post** /settings/events/user-defined-events/{id}/enable | Enable an existing User-Defined event
[**EnableVidm**](SettingsApi.md#EnableVidm) | **Post** /settings/vidm/enable | Enable VMware Identity Manager integration
[**GetAllUserDefinedEvents**](SettingsApi.md#GetAllUserDefinedEvents) | **Get** /settings/events/user-defined-events | List the created User Defined Event defintions.
[**GetBackupConfig**](SettingsApi.md#GetBackupConfig) | **Get** /settings/backup | Get Backup configuration
[**GetBackupStatusReport**](SettingsApi.md#GetBackupStatusReport) | **Get** /settings/backup/status | Get currently running or last Backup job status
[**GetCertificate**](SettingsApi.md#GetCertificate) | **Get** /settings/vidm/certificate | Get certificate for given url
[**GetConnectedClientsToWebProxy**](SettingsApi.md#GetConnectedClientsToWebProxy) | **Get** /settings/proxy-servers/{id}/connected-clients | Get details of connected clients to Web Proxy Server
[**GetInfraNodesWebProxy**](SettingsApi.md#GetInfraNodesWebProxy) | **Get** /settings/proxy-servers/infra | Get details of web proxy config associated with infra nodes
[**GetIpTag**](SettingsApi.md#GetIpTag) | **Get** /settings/ip-tags/{tag-id} | Show IP tag details
[**GetIpTagIds**](SettingsApi.md#GetIpTagIds) | **Get** /settings/ip-tags/tag-ids | Show IP tag IDs
[**GetLicenses**](SettingsApi.md#GetLicenses) | **Get** /settings/licensing/ | Get current licensing and license usage information
[**GetRestoreConfig**](SettingsApi.md#GetRestoreConfig) | **Get** /settings/restore | Get Restore configuration
[**GetRestoreStatusReport**](SettingsApi.md#GetRestoreStatusReport) | **Get** /settings/restore/status | Get currently running or last Restore job status
[**GetSubnetMappings**](SettingsApi.md#GetSubnetMappings) | **Get** /settings/subnet-mappings | Get all subnet mappings
[**GetUser**](SettingsApi.md#GetUser) | **Get** /settings/users/{id} | Get details of a user
[**GetUserDefinedEvent**](SettingsApi.md#GetUserDefinedEvent) | **Get** /settings/events/user-defined-events/{id} | Get details of an existing User-Defined event.
[**GetUserGroup**](SettingsApi.md#GetUserGroup) | **Get** /settings/user-groups/{id} | Get details of a user-group
[**GetUserGroups**](SettingsApi.md#GetUserGroups) | **Get** /settings/user-groups | List user-groups
[**GetUsers**](SettingsApi.md#GetUsers) | **Get** /settings/users | List the users
[**GetVidmConfiguration**](SettingsApi.md#GetVidmConfiguration) | **Get** /settings/vidm | Get configuration details of VMware Identity Manager
[**GetWebProxies**](SettingsApi.md#GetWebProxies) | **Get** /settings/proxy-servers | List of configured web proxy servers
[**GetWebProxy**](SettingsApi.md#GetWebProxy) | **Get** /settings/proxy-servers/{id} | Get details of an existing Web Proxy Server
[**RemoveIpTag**](SettingsApi.md#RemoveIpTag) | **Post** /settings/ip-tags/{tag-id}/remove | Remove tag from IP addresses
[**SaveVidmConfiguration**](SettingsApi.md#SaveVidmConfiguration) | **Post** /settings/vidm | Configure VMware Identity Manager
[**SettingsSnmpProfilesGet**](SettingsApi.md#SettingsSnmpProfilesGet) | **Get** /settings/snmp/profiles | List the configured SNMP Trap destination profiles
[**SettingsSnmpProfilesIdDelete**](SettingsApi.md#SettingsSnmpProfilesIdDelete) | **Delete** /settings/snmp/profiles/{id} | Delete an existing SNMP Trap destination profile
[**SettingsSnmpProfilesIdGet**](SettingsApi.md#SettingsSnmpProfilesIdGet) | **Get** /settings/snmp/profiles/{id} | Get details of an existing SNMP destination profile
[**SettingsSnmpProfilesIdMigrateEventsPost**](SettingsApi.md#SettingsSnmpProfilesIdMigrateEventsPost) | **Post** /settings/snmp/profiles/{id}/migrate-events | Migrate event subscriptions to other SNMP Trap destination profiles
[**SettingsSnmpProfilesIdPut**](SettingsApi.md#SettingsSnmpProfilesIdPut) | **Put** /settings/snmp/profiles/{id} | Update an existing SNMP destination profile
[**SettingsSnmpProfilesPost**](SettingsApi.md#SettingsSnmpProfilesPost) | **Post** /settings/snmp/profiles | Add new SNMP Trap destination profile
[**SettingsSnmpProfilesSendTestTrapPost**](SettingsApi.md#SettingsSnmpProfilesSendTestTrapPost) | **Post** /settings/snmp/profiles/send-test-trap | Send Test trap to SNMP destination profile
[**SettingsUsersPasswordPut**](SettingsApi.md#SettingsUsersPasswordPut) | **Put** /settings/users/password | Update user password
[**UpdateBackupConfig**](SettingsApi.md#UpdateBackupConfig) | **Put** /settings/backup | Update Backup configuration
[**UpdateSubnetMapping**](SettingsApi.md#UpdateSubnetMapping) | **Put** /settings/subnet-mappings/{id} | Update subnet mapping
[**UpdateUserDefinedEvent**](SettingsApi.md#UpdateUserDefinedEvent) | **Put** /settings/events/user-defined-events/{id} | Update an existing User-Defined event.
[**UpdateVidmConfiguration**](SettingsApi.md#UpdateVidmConfiguration) | **Put** /settings/vidm | Update VMware Identity Manager configuration
[**UpdateVidmUserGroupRole**](SettingsApi.md#UpdateVidmUserGroupRole) | **Put** /settings/user-groups/vidm | Update role for user-group mapped through VMware Identity Manager
[**UpdateVidmUserRole**](SettingsApi.md#UpdateVidmUserRole) | **Put** /settings/users/vidm | Update role for user mapped through VMware Identity Manager
[**UpdateWebProxy**](SettingsApi.md#UpdateWebProxy) | **Put** /settings/proxy-servers/{id} | Update the existing web proxy server
[**ValidateConnectionsViaWebProxy**](SettingsApi.md#ValidateConnectionsViaWebProxy) | **Post** /settings/proxy-servers/{id}/validate-entities | Validate the connections via web proxy
[**ValidateSerialNumber**](SettingsApi.md#ValidateSerialNumber) | **Post** /settings/licensing/validate | Validates license key
[**ValidateWebProxyMigration**](SettingsApi.md#ValidateWebProxyMigration) | **Post** /settings/proxy-servers/validate-migration | Validate the connections when migrating from one web proxy to another



## ActivateSerialNumber

> LicensingActivate ActivateSerialNumber(ctx, body)

Activates a valid license key. This API is not applicable for on-boarding.

Activates a valid license key to an existing deployment. For example, adding a SD-WAN license key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LicensingKeyRequest**](LicensingKeyRequest.md)| Licensing Request | 

### Return type

[**LicensingActivate**](LicensingActivate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddBackupConfig

> BackupRestoreRequest AddBackupConfig(ctx, body)

Configure backup of configuration

Configure backup of platform configuration; system settings, data sources, SMTP and SNMP settings, and more. This can be sent to a local filesystem, or a remove SSH or FTP server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BackupRestoreRequest**](BackupRestoreRequest.md)| Backup configuration | 

### Return type

[**BackupRestoreRequest**](BackupRestoreRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIpTag

> AddIpTag(ctx, tagId, body)

Tag IP addresses with tag-id

Tag IP addresses with tag-id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string**|  | 
**body** | [**IpTag**](IpTag.md)| Ip Tag | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRestoreConfig

> BackupRestoreRequest AddRestoreConfig(ctx, body)

Configure restore of configuration and triggers restore operation

Configure restore of platform configuration and triggers restore operation. Must be done on a clean platform deployment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BackupRestoreRequest**](BackupRestoreRequest.md)| Restore configuration | 

### Return type

[**BackupRestoreRequest**](BackupRestoreRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVidmUser

> UserResponse AddVidmUser(ctx, body)

Add a VMware Identity manager user to vRealize Network Insight

Add a new VMware Identity manager user to vRealize Network Insight. Only admin users can perform this action. VMware Identity Manager must be configured prior to this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VidmUserRequest**](VidmUserRequest.md)| User details | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVidmUserGroup

> UserGroupResponse AddVidmUserGroup(ctx, body)

Add a VMware Identity Manager user-group to vRealize Network Insight

Add a new VMware Identity Manager user-group to vRealize Network Insight. Only admin users can perform this action. VMware Identity Manager must be configured prior to this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VidmUserGroupRequest**](VidmUserGroupRequest.md)| User-group details | 

### Return type

[**UserGroupResponse**](UserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebProxy

> ProxyProfileIdResponse AddWebProxy(ctx, body)

Add new Web Proxy in the system

Add new Web Proxy server. By Default, maximum of 10 web proxy servers can be configured. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**WebProxyRequest**](WebProxyRequest.md)| Web Proxy Details | 

### Return type

[**ProxyProfileIdResponse**](ProxyProfileIdResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubnetMapping

> SubnetMapping CreateSubnetMapping(ctx, body)

Create subnet mapping

Create a new subnet mapping (CIDR to VLAN ID). Physical to physical flows (without the switches added as a data source) can be linked to a VLAN ID with to these mappings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SubnetMappingRequest**](SubnetMappingRequest.md)|  | 

### Return type

[**SubnetMapping**](SubnetMapping.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserDefinedEvent

> SubscriptionResponse CreateUserDefinedEvent(ctx, body)

Add new User-Defined event

Add new User-Defined event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SubscriptionRequest**](SubscriptionRequest.md)| Add new User-Defined events | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSerialNumber

> DeactivateSerialNumber(ctx, body)

Deactivates an existing license key

Deactivates an existing license key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LicensingKeyRequest**](LicensingKeyRequest.md)| Licensing Request | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupConfig

> DeleteBackupConfig(ctx, )

Delete existing Backup configuration

Delete existing Backup configuration.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRestoreConfig

> DeleteRestoreConfig(ctx, )

Delete existing restore configuration

Delete existing restore configuration.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnetMapping

> DeleteSubnetMapping(ctx, id)

Delete subnet mapping

Delete an existing subnet mapping (CIDR to VLAN ID).

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id)

Delete an existing user.

Delete an existing user. Only admin users can perform this action. This action is currently restricted to user mapped through VMware Identity Manager

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDefinedEvent

> DeleteUserDefinedEvent(ctx, id)

Delete an existing User-Defined event

Delete an existing User-Defined event.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroup

> DeleteUserGroup(ctx, id)

Delete an existing user-group

Delete an existing user-group. Only admin users can perform this action. This action is currently restricted to user-group mapped through VMware Identity Manager

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVidmConfiguration

> DeleteVidmConfiguration(ctx, )

Delete VMware Identity Manager configuration

Delete the existing VMware Identity Manager configuration. Only admin users can perform this action.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebProxy

> DeleteWebProxy(ctx, id, optional)

Delete an existing Web Proxy server

Delete an existing existing Web Proxy Server. Only admin user can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***DeleteWebProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWebProxyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migrationWebProxyId** | **optional.String**| Identifier of web proxy to which all the existing connections will be migrated to | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUserDefinedEvent

> SubscriptionResponse DisableUserDefinedEvent(ctx, id)

Disable an existing User-Defined event

Disable an existing User-Defined event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableVidm

> DisableVidm(ctx, )

Disable VMware Identity Manager integration

Disable VMware Identity Manager integration, although it will not delete the configuration. Only admin users can perform this action.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableUserDefinedEvent

> SubscriptionResponse EnableUserDefinedEvent(ctx, id)

Enable an existing User-Defined event

Enable an existing User-Defined event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableVidm

> EnableVidm(ctx, )

Enable VMware Identity Manager integration

Enable VMware Identity Manager integration. Only admin users can perform this action.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUserDefinedEvents

> SubscriptionListResponse GetAllUserDefinedEvents(ctx, )

List the created User Defined Event defintions.

List the created User Defined Event defintions.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupConfig

> BackupRestoreRequest GetBackupConfig(ctx, )

Get Backup configuration

Get Backup configuration

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BackupRestoreRequest**](BackupRestoreRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupStatusReport

> BackupRestoreStatusResponse GetBackupStatusReport(ctx, )

Get currently running or last Backup job status

Get currently running or last Backup job status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BackupRestoreStatusResponse**](BackupRestoreStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> VidmCertificate GetCertificate(ctx, url)

Get certificate for given url

Get certificate for given url

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**| Complete URL path to fetch certificate | 

### Return type

[**VidmCertificate**](VidmCertificate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedClientsToWebProxy

> EntityUsingOrBehindWebProxyList GetConnectedClientsToWebProxy(ctx, id)

Get details of connected clients to Web Proxy Server

Get details of connected clients to Web Proxy Server. Only admin and audit users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**EntityUsingOrBehindWebProxyList**](EntityUsingOrBehindWebProxyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfraNodesWebProxy

> WebProxyConfigListForNodes GetInfraNodesWebProxy(ctx, )

Get details of web proxy config associated with infra nodes

Get web proxy for infra nodes. Only admin and audit users can perform this action.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebProxyConfigListForNodes**](WebProxyConfigListForNodes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpTag

> IpTag GetIpTag(ctx, tagId)

Show IP tag details

Show IP tag details with member IP addresses and subnets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string**|  | 

### Return type

[**IpTag**](IpTag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpTagIds

> IpTagIdList GetIpTagIds(ctx, )

Show IP tag IDs

Get all tag IDs e.g. EAST_WEST, INTERNET etc. These are used in the flow tags.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**IpTagIdList**](IpTagIdList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenses

> LicensingResponse GetLicenses(ctx, )

Get current licensing and license usage information

Get information for current licenses and license usage

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LicensingResponse**](LicensingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreConfig

> BackupRestoreRequest GetRestoreConfig(ctx, )

Get Restore configuration

Get Restore configuration

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BackupRestoreRequest**](BackupRestoreRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreStatusReport

> BackupRestoreStatusResponse GetRestoreStatusReport(ctx, )

Get currently running or last Restore job status

Get currently running or last Restore job status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BackupRestoreStatusResponse**](BackupRestoreStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnetMappings

> SubnetMappingList GetSubnetMappings(ctx, )

Get all subnet mappings

Get all subnet mappings (CIDR to VLAN ID).

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SubnetMappingList**](SubnetMappingList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, id)

Get details of a user

Get details of a user. Only admin users can perform this action. This action is currently restricted to user mapped through VMware Identity Manager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDefinedEvent

> SubscriptionResponse GetUserDefinedEvent(ctx, id)

Get details of an existing User-Defined event.

Get details of an existing User-Defined event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroup

> UserGroupResponse GetUserGroup(ctx, id)

Get details of a user-group

Get details of a user-group. Only admin users can perform this action. This action is currently restricted to user-group mapped through VMware Identity Manager.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**UserGroupResponse**](UserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> PagedUserGroupListResponse GetUserGroups(ctx, type_, optional)

List user-groups

Retrieve the list of existing user-groups. Only admin users can retrieve this information. This action is currently restricted to user-groups mapped through VMware Identity Manager.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Type of user-group | 
 **optional** | ***GetUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 

### Return type

[**PagedUserGroupListResponse**](PagedUserGroupListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> PagedUserListResponse GetUsers(ctx, type_, optional)

List the users

List the existing users. Only admin users can retrieve this information. This action is currently restricted to users mapped through VMware Identity Manager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Type of user | 
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 

### Return type

[**PagedUserListResponse**](PagedUserListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVidmConfiguration

> VidmConfigResponse GetVidmConfiguration(ctx, )

Get configuration details of VMware Identity Manager

Retrieve the saved configuration of VMware Identity Manager in Network Insight. Only admin users can retrieve this information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VidmConfigResponse**](VidmConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebProxies

> ProxyListResponse GetWebProxies(ctx, )

List of configured web proxy servers

List of configured web proxy servers. Only admin and audit users can retrieve this information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ProxyListResponse**](ProxyListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebProxy

> ProxyResponse GetWebProxy(ctx, id)

Get details of an existing Web Proxy Server

Get details of an existing Web Proxy Server. Only admin and audit users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIpTag

> RemoveIpTag(ctx, tagId, body)

Remove tag from IP addresses

Remove tag from IP addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string**|  | 
**body** | [**IpTag**](IpTag.md)| Ip Tag | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveVidmConfiguration

> VidmConfigResponse SaveVidmConfiguration(ctx, body)

Configure VMware Identity Manager

vRealize Network Insight supports SSO authentication through VMware Identity Manager. To authenticate against a particular VMware Identity Manager appliance, it must be configured and enabled in vRealize Network Insight. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VidmConfiguration**](VidmConfiguration.md)| VMware Identity Manager configuration details | 

### Return type

[**VidmConfigResponse**](VidmConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesGet

> SnmpProfileListResponse SettingsSnmpProfilesGet(ctx, )

List the configured SNMP Trap destination profiles

List the configured SNMP Trap destination profiles. Only admin users can retrieve this information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SnmpProfileListResponse**](SnmpProfileListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesIdDelete

> SettingsSnmpProfilesIdDelete(ctx, id)

Delete an existing SNMP Trap destination profile

Delete an existing SNMP destination profile. Only admin users can perform this action.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesIdGet

> SnmpProfileResponse SettingsSnmpProfilesIdGet(ctx, id)

Get details of an existing SNMP destination profile

Get details of an existing SNMP Trap destination profile. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**SnmpProfileResponse**](SnmpProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesIdMigrateEventsPost

> SettingsSnmpProfilesIdMigrateEventsPost(ctx, id, body)

Migrate event subscriptions to other SNMP Trap destination profiles

Migrate event subscriptions from given SNMP profile to multiple other SNMP Trap destination profiles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**SnmpProfileIds**](SnmpProfileIds.md)| SNMP Trap destination profiles list | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesIdPut

> SnmpProfileResponse SettingsSnmpProfilesIdPut(ctx, id, body)

Update an existing SNMP destination profile

Update an existing SNMP Trap destination profile. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**SnmpProfileRequest**](SnmpProfileRequest.md)| SNMP destination profile | 

### Return type

[**SnmpProfileResponse**](SnmpProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesPost

> SnmpProfileResponse SettingsSnmpProfilesPost(ctx, body)

Add new SNMP Trap destination profile

Add a new SNMP destination profile. By Default, maximum of 4 SNMP Trap profiles can be configured. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SnmpProfileRequest**](SnmpProfileRequest.md)| SNMP Trap destination profile | 

### Return type

[**SnmpProfileResponse**](SnmpProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsSnmpProfilesSendTestTrapPost

> TestSnmpProfileResponse SettingsSnmpProfilesSendTestTrapPost(ctx, body)

Send Test trap to SNMP destination profile

Send Test trap to SNMP destination profile. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SnmpProfileRequest**](SnmpProfileRequest.md)| SNMP destination profile | 

### Return type

[**TestSnmpProfileResponse**](TestSnmpProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsUsersPasswordPut

> SettingsUsersPasswordPut(ctx, body)

Update user password

Update password of a local user. Any user can change his own password. Admin has permission to change any user password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserUpdateRequest**](UserUpdateRequest.md)| User details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupConfig

> BackupRestoreRequest UpdateBackupConfig(ctx, body)

Update Backup configuration

Update Backup configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BackupRestoreRequest**](BackupRestoreRequest.md)| Updated Backup configuration | 

### Return type

[**BackupRestoreRequest**](BackupRestoreRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnetMapping

> SubnetMapping UpdateSubnetMapping(ctx, id, body)

Update subnet mapping

Update an existing subnet mapping (VLAN ID for the given CIDR).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**SubnetMappingRequest**](SubnetMappingRequest.md)|  | 

### Return type

[**SubnetMapping**](SubnetMapping.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserDefinedEvent

> SubscriptionResponse UpdateUserDefinedEvent(ctx, id, body)

Update an existing User-Defined event.

Update an existing User-Defined event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**SubscriptionRequest**](SubscriptionRequest.md)| Update an existing User-Defined event. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVidmConfiguration

> VidmConfigResponse UpdateVidmConfiguration(ctx, body)

Update VMware Identity Manager configuration

Update the existing VMware Identity Manager configuration. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VidmConfiguration**](VidmConfiguration.md)| VMware Identity Manager configuration details | 

### Return type

[**VidmConfigResponse**](VidmConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVidmUserGroupRole

> UserGroupResponse UpdateVidmUserGroupRole(ctx, id, body)

Update role for user-group mapped through VMware Identity Manager

Update role for user-group mapped through VMware Identity Manager. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**VidmUserGroupRequest**](VidmUserGroupRequest.md)| User-group details | 

### Return type

[**UserGroupResponse**](UserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVidmUserRole

> UserResponse UpdateVidmUserRole(ctx, id, body)

Update role for user mapped through VMware Identity Manager

Update role for user mapped through VMware Identity Manager. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**VidmUserRequest**](VidmUserRequest.md)| User details | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebProxy

> UpdateWebProxy(ctx, id, body)

Update the existing web proxy server

Update the details of an existing Web Proxy Server. Only admin can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**WebProxyRequest**](WebProxyRequest.md)| Updated Web Proxy Details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnectionsViaWebProxy

> ValidateConnectionViaWebProxyResponseList ValidateConnectionsViaWebProxy(ctx, id, body)

Validate the connections via web proxy

Validate the connections with the updated web proxy settings. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
**body** | [**WebProxyRequest**](WebProxyRequest.md)| Web Proxy Details to be updated | 

### Return type

[**ValidateConnectionViaWebProxyResponseList**](ValidateConnectionViaWebProxyResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSerialNumber

> LicensingActivate ValidateSerialNumber(ctx, body)

Validates license key

Validates license key, makes sure it can be applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LicensingKeyRequest**](LicensingKeyRequest.md)| Licensing Request | 

### Return type

[**LicensingActivate**](LicensingActivate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateWebProxyMigration

> ValidateConnectionViaWebProxyResponseList ValidateWebProxyMigration(ctx, body)

Validate the connections when migrating from one web proxy to another

Validate the connections when migrating from one web proxy to another. Only admin users can perform this action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**WebProxyMigrationRequest**](WebProxyMigrationRequest.md)| Web Proxy Details for migration | 

### Return type

[**ValidateConnectionViaWebProxyResponseList**](ValidateConnectionViaWebProxyResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

