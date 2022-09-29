# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePool**](PoolsApi.md#CreatePool) | **Post** /pools | Create Pool
[**DeletePool**](PoolsApi.md#DeletePool) | **Delete** /pools/{name} | Delete Pool
[**GetPool**](PoolsApi.md#GetPool) | **Get** /pools/{name} | Get Pool
[**ListPools**](PoolsApi.md#ListPools) | **Get** /pools | List Pools

# **CreatePool**
> DsvPool CreatePool(ctx, body)
Create Pool

Create an engine pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvPoolCreate**](DsvPoolCreate.md)|  | 

### Return type

[**DsvPool**](Pool.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePool**
> DsvMessageResponse DeletePool(ctx, name)
Delete Pool

Delete an engine pool by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name to lookup path by | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPool**
> DsvPool GetPool(ctx, name)
Get Pool

Retrieve pool by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name to lookup path by | 

### Return type

[**DsvPool**](Pool.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPools**
> DsvPoolListResult ListPools(ctx, )
List Pools

List all existing pools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DsvPoolListResult**](PoolListResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

