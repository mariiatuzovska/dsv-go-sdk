# \ConfigApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigApi.md#GetConfig) | **Get** /config | Get Config
[**GetConfigByVersion**](ConfigApi.md#GetConfigByVersion) | **Get** /config/version/{version} | Get Config By Version
[**PostConfig**](ConfigApi.md#PostConfig) | **Post** /config | Create Config


# **GetConfig**
> DsvConfigResponse GetConfig(ctx, )
Get Config

Get config data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DsvConfigResponse**](ConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigByVersion**
> DsvConfigVersionResponse GetConfigByVersion(ctx, version)
Get Config By Version

Get zero or more representations of config specified by version number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **int64**| Versions to return | 

### Return type

[**DsvConfigVersionResponse**](ConfigVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostConfig**
> DsvConfigResponse PostConfig(ctx, body)
Create Config

Create or update config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvPostConfigModel**](DsvPostConfigModel.md)|  | 

### Return type

[**DsvConfigResponse**](ConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

