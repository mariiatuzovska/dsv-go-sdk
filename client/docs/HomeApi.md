# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHomeSecret**](HomeApi.md#CreateHomeSecret) | **Post** /home/{principalName}/{path} | Create home secrets
[**DeleteHomeSecret**](HomeApi.md#DeleteHomeSecret) | **Delete** /home/{principalName}/{path} | Delete Home Secret
[**GetHomeSecret**](HomeApi.md#GetHomeSecret) | **Get** /home/{principalName}/{path} | Get Home
[**GetHomeSecretByVersion**](HomeApi.md#GetHomeSecretByVersion) | **Get** /home/{principalName}/{path}/version/{version} | Get Home Secret By Version
[**GetHomeSecretDescription**](HomeApi.md#GetHomeSecretDescription) | **Get** /home/{principalName}/{path}::description | Get home Secret Description
[**RestoreHomeSecret**](HomeApi.md#RestoreHomeSecret) | **Get** /home/{principalName}/{path}/restore | Restore home Secret
[**RollbackHomeSecret**](HomeApi.md#RollbackHomeSecret) | **Put** /home/{principalName}/{path}/rollback/{version} | Rollback a Home Secret
[**SearchHomeSecrets**](HomeApi.md#SearchHomeSecrets) | **Get** /home/{principalName} | Search for Home Secrets
[**UpdateHomeSecret**](HomeApi.md#UpdateHomeSecret) | **Put** /home/{principalName}/{path} | Update home Secret

# **CreateHomeSecret**
> DsvSecretResponse CreateHomeSecret(ctx, body, path)
Create home secrets

Creates a new home secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvSecretCreate**](DsvSecretCreate.md)|  | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHomeSecret**
> DsvMessageResponse DeleteHomeSecret(ctx, path, optional)
Delete Home Secret

Deletes a home secret by path or by id in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***HomeApiDeleteHomeSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HomeApiDeleteHomeSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHomeSecret**
> DsvSecretResponse GetHomeSecret(ctx, principalName, path, optional)
Get Home

Gets a full home secret by path or by id in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalName** | **string**| Principal Name | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***HomeApiGetHomeSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HomeApiGetHomeSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHomeSecretByVersion**
> DsvSecretVersionResponse GetHomeSecretByVersion(ctx, path, principalName, version)
Get Home Secret By Version

Gets a full home secret by path and version in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
  **principalName** | **string**| Principal Name | 
  **version** | **int64**| Versions to return | 

### Return type

[**DsvSecretVersionResponse**](SecretVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHomeSecretDescription**
> DsvSecretDescription GetHomeSecretDescription(ctx, principalName, path, optional)
Get home Secret Description

Gets a home secret's metadata without returning the sensitive data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalName** | **string**| Principal Name | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***HomeApiGetHomeSecretDescriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HomeApiGetHomeSecretDescriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**DsvSecretDescription**](SecretDescription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreHomeSecret**
> RestoreHomeSecret(ctx, principalName, path, optional)
Restore home Secret

Restores a home secret by path or by id in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalName** | **string**| Principal Name | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***HomeApiRestoreHomeSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HomeApiRestoreHomeSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackHomeSecret**
> DsvSecretVersionResponse RollbackHomeSecret(ctx, path, principalName, version)
Rollback a Home Secret

Rollbacks a Home Secret to a previous version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
  **principalName** | **string**| Principal Name | 
  **version** | **int64**| Versions to return | 

### Return type

[**DsvSecretVersionResponse**](SecretVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchHomeSecrets**
> DsvSecretSearch SearchHomeSecrets(ctx, principalName, optional)
Search for Home Secrets

Lists home secret paths that start with the path parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalName** | **string**| Principal Name | 
 **optional** | ***HomeApiSearchHomeSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HomeApiSearchHomeSecretsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchTerm** | **optional.String**| Partial search term for search by path | 
 **searchLinks** | **optional.Bool**| Whether to search for secrets that link to the path in the search term | 
 **searchField** | **optional.String**| Secret field for advanced searching | 
 **searchComparison** | **optional.String**| Comparison type (equal, contains, begins_with) for advanced searching | 
 **searchType** | **optional.String**| Attribute type (string, number) to search on | 
 **limit** | **optional.Int64**| Limit for the number of results per page (cursor) | 
 **cursor** | **optional.String**| Cursor to next batch of results | 
 **sort** | **optional.String**| Sort results ascending (asc) or descending (desc) order by lastModified attribute on field search. Default is desc | 

### Return type

[**DsvSecretSearch**](SecretSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHomeSecret**
> DsvSecretResponse UpdateHomeSecret(ctx, body, path)
Update home Secret

Updates an existing home secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvSecretUpdate**](DsvSecretUpdate.md)|  | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

