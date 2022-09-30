# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthSettings**](SettingsApi.md#CreateAuthSettings) | **Post** /config/auth/ | Create Authentication Provider
[**DeleteAuthSettings**](SettingsApi.md#DeleteAuthSettings) | **Delete** /config/auth/{name} | Delete Authentication Provider
[**GetAuthSettings**](SettingsApi.md#GetAuthSettings) | **Get** /config/auth/{name} | Get Authentication Provider
[**GetAuthSettingsByVersion**](SettingsApi.md#GetAuthSettingsByVersion) | **Get** /config/auth/{name}/version/{version} | Get a list of Authentication Settings by version
[**RestoreAuthSettings**](SettingsApi.md#RestoreAuthSettings) | **Get** /config/auth/{name}/restore | Restore Authentication Provider
[**RollbackAuthSettings**](SettingsApi.md#RollbackAuthSettings) | **Put** /config/auth/{name}/rollback/{version} | Rollback Authentication Provider
[**SearchSettings**](SettingsApi.md#SearchSettings) | **Get** /config/auth | Search Authentication Providers
[**UpdateAuthSettings**](SettingsApi.md#UpdateAuthSettings) | **Put** /config/auth/{name} | Update Authentication Provider

# **CreateAuthSettings**
> DsvAuthenticationSettingsResponse CreateAuthSettings(ctx, body)
Create Authentication Provider

Creates new authentication provider settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvAuthenticationSettingsCreateModel**](DsvAuthenticationSettingsCreateModel.md)|  | 

### Return type

[**DsvAuthenticationSettingsResponse**](AuthenticationSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthSettings**
> DsvMessageResponse DeleteAuthSettings(ctx, name, optional)
Delete Authentication Provider

Remove authentication provider settings from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup authentication settings by. | 
 **optional** | ***SettingsApiDeleteAuthSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiDeleteAuthSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthSettings**
> DsvAuthenticationSettingsResponse GetAuthSettings(ctx, name)
Get Authentication Provider

Retrieve authentication providers by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup authentication settings by. | 

### Return type

[**DsvAuthenticationSettingsResponse**](AuthenticationSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthSettingsByVersion**
> DsvAuthenticationSettingsVersionResponse GetAuthSettingsByVersion(ctx, name, version)
Get a list of Authentication Settings by version

Get a full Settings by name and version in the query. Returns a list of zero up to n versions of an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup authentication settings by. | 
  **version** | **int64**| Versions to return. | 

### Return type

[**DsvAuthenticationSettingsVersionResponse**](AuthenticationSettingsVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreAuthSettings**
> RestoreAuthSettings(ctx, name)
Restore Authentication Provider

Restore authentication provider settings if it had been marked for deletion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup authentication settings by. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackAuthSettings**
> DsvAuthenticationSettingsResponse RollbackAuthSettings(ctx, name, version)
Rollback Authentication Provider

Rollback authentication provider to a previous version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup authentication settings by. | 
  **version** | **int64**| Versions to return. | 

### Return type

[**DsvAuthenticationSettingsResponse**](AuthenticationSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSettings**
> DsvAuthenticationSettingsSearchResponse SearchSettings(ctx, optional)
Search Authentication Providers

Search authentication providers by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiSearchSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSearchSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| Search pattern for setting name. | 
 **limit** | **optional.Int64**| The maximum number of results per cursor. | 
 **cursor** | **optional.String**| Cursor to next batch of results. | 

### Return type

[**DsvAuthenticationSettingsSearchResponse**](AuthenticationSettingsSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthSettings**
> DsvAuthenticationSettingsResponse UpdateAuthSettings(ctx, body, name)
Update Authentication Provider

Updates all fields on an existing authentication provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvAuthenticationDetailsModel**](DsvAuthenticationDetailsModel.md)|  | 
  **name** | **string**| Full name to lookup authentication settings by. | 

### Return type

[**DsvAuthenticationSettingsResponse**](AuthenticationSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

