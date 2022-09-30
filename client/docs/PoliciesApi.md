# \PoliciesApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](PoliciesApi.md#CreatePolicy) | **Post** /config/policies/ | Create Policy
[**DeletePolicy**](PoliciesApi.md#DeletePolicy) | **Delete** /config/policies/{path} | Delete Policy
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /config/policies/{path} | Get Policy
[**GetPolicyByVersion**](PoliciesApi.md#GetPolicyByVersion) | **Get** /config/policies/{path}/version/{version} | Get a list of policies by version
[**RestorePolicy**](PoliciesApi.md#RestorePolicy) | **Get** /config/policies/{path}/restore | Restore Policy
[**RollbackPolicy**](PoliciesApi.md#RollbackPolicy) | **Put** /config/policies/{path}/rollback/{version} | Rollback Policy
[**SearchFilter**](PoliciesApi.md#SearchFilter) | **Get** /config/policies | Search Policies
[**UpdatePolicy**](PoliciesApi.md#UpdatePolicy) | **Put** /config/policies/{path} | Update Policy


# **CreatePolicy**
> DsvPolicyResponse CreatePolicy(ctx, body)
Create Policy

Creates new policy with permission document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvPolicyCreate**](DsvPolicyCreate.md)|  | 

### Return type

[**DsvPolicyResponse**](PolicyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DsvMessageResponse DeletePolicy(ctx, path, optional)
Delete Policy

Remove an existing policy by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path to lookup policy | 
 **optional** | ***PoliciesApiDeletePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiDeletePolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> DsvPolicyResponse GetPolicy(ctx, path)
Get Policy

Retrieve policy by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path to lookup policy | 

### Return type

[**DsvPolicyResponse**](PolicyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyByVersion**
> DsvPolicyVersionResponse GetPolicyByVersion(ctx, path, version)
Get a list of policies by version

Get a full policies by path and version in the query.Returns a list of zero up to n versions of an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path to lookup policy | 
  **version** | **int64**| Versions to return | 

### Return type

[**DsvPolicyVersionResponse**](PolicyVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestorePolicy**
> RestorePolicy(ctx, path)
Restore Policy

Restore an existing policy by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path to lookup policy | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackPolicy**
> DsvPolicyResponse RollbackPolicy(ctx, path, version)
Rollback Policy

Overwrites an existing policy with its previous version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path to lookup policy | 
  **version** | **int64**| Versions to return | 

### Return type

[**DsvPolicyResponse**](PolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchFilter**
> DsvPolicySearchResponse SearchFilter(ctx, optional)
Search Policies

Search permission policies by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesApiSearchFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiSearchFilterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| Search term | 
 **limit** | **optional.Int64**| The maximum number of results per cursor | 
 **cursor** | **optional.String**| Cursor to next batch of results | 

### Return type

[**DsvPolicySearchResponse**](PolicySearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> DsvPolicyResponse UpdatePolicy(ctx, path, body)
Update Policy

Overwrites an existing policy with an updated permissions document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Full path for policy | 
  **body** | [**DsvPolicyUpdate**](DsvPolicyUpdate.md)|  | 

### Return type

[**DsvPolicyResponse**](PolicyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

