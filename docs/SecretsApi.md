# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](SecretsApi.md#CreateSecret) | **Post** /secrets/{path} | Create Secret
[**DeleteSecret**](SecretsApi.md#DeleteSecret) | **Delete** /secrets/{path} | Delete Secret
[**GetSecret**](SecretsApi.md#GetSecret) | **Get** /secrets/{path} | Get Secret
[**GetSecretByVersion**](SecretsApi.md#GetSecretByVersion) | **Get** /secrets/{path}/version/{version} | Get Secret By Version
[**GetSecretDescription**](SecretsApi.md#GetSecretDescription) | **Get** /secrets/{path}::description | Get Secret Description
[**ListSecretPaths**](SecretsApi.md#ListSecretPaths) | **Get** /secrets/{path}::listpaths | List Secret Paths
[**RollbackSecret**](SecretsApi.md#RollbackSecret) | **Put** /secrets/{path}/rollback/{version} | Rollback a Secret
[**SearchSecrets**](SecretsApi.md#SearchSecrets) | **Get** /secrets | Search for Secrets
[**UpdateSecret**](SecretsApi.md#UpdateSecret) | **Put** /secrets/{path} | Update Secret

# **CreateSecret**
> SecretResponse CreateSecret(ctx, body, path)
Create Secret

Creates a new secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretCreate**](SecretCreate.md)|  | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecret**
> MessageResponse DeleteSecret(ctx, path, optional)
Delete Secret

Deletes a secret by path or by id in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***SecretsApiDeleteSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiDeleteSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 
 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecret**
> SecretResponse GetSecret(ctx, path, optional)
Get Secret

Gets a full secret by path or by id in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***SecretsApiGetSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiGetSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretByVersion**
> SecretVersionResponse GetSecretByVersion(ctx, path, version)
Get Secret By Version

Gets a full secret by path and version in the query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
  **version** | **int64**| Versions to return | 

### Return type

[**SecretVersionResponse**](SecretVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretDescription**
> SecretDescription GetSecretDescription(ctx, path, optional)
Get Secret Description

Gets a secret's metadata without returning the sensitive data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
 **optional** | ***SecretsApiGetSecretDescriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiGetSecretDescriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| Unique uuid identifying a secret | 

### Return type

[**SecretDescription**](SecretDescription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecretPaths**
> SecretListPathsResponse ListSecretPaths(ctx, path, optional)
List Secret Paths

Lists secret paths that start with the path parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The secret path to match on | 
 **optional** | ***SecretsApiListSecretPathsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiListSecretPathsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| The maximum number of path matches to return | 

### Return type

[**SecretListPathsResponse**](SecretListPathsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackSecret**
> SecretVersionResponse RollbackSecret(ctx, path, version)
Rollback a Secret

Rollbacks a Secret to a previous version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 
  **version** | **int64**| Versions to return | 

### Return type

[**SecretVersionResponse**](SecretVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSecrets**
> SecretSearch SearchSecrets(ctx, optional)
Search for Secrets

Lists secret paths that start with the path parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecretsApiSearchSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiSearchSecretsOpts struct
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

[**SecretSearch**](SecretSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecret**
> SecretResponse UpdateSecret(ctx, body, path)
Update Secret

Updates an existing secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretUpdate**](SecretUpdate.md)|  | 
  **path** | **string**| The full secret path i.e. servers/prod/webserver-01 | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

