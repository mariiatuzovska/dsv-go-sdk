# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngine**](EnginesApi.md#CreateEngine) | **Post** /engines | Create Engine
[**DeleteEngine**](EnginesApi.md#DeleteEngine) | **Delete** /engines/{name} | Delete Engine
[**GetEngine**](EnginesApi.md#GetEngine) | **Get** /engines/{name} | Get Engine
[**ListEngines**](EnginesApi.md#ListEngines) | **Get** /engines | List Engines
[**PingEngine**](EnginesApi.md#PingEngine) | **Post** /engines/{name}/ping | Ping Engine

# **CreateEngine**
> EngineCreateResponse CreateEngine(ctx, body)
Create Engine

Registers a new engine and returns its key pair.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EngineCreate**](EngineCreate.md)|  | 

### Return type

[**EngineCreateResponse**](EngineCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEngine**
> MessageResponse DeleteEngine(ctx, path, optional)
Delete Engine

Delete an engine  by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Name of engine | 
 **optional** | ***EnginesApiDeleteEngineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnginesApiDeleteEngineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEngine**
> EngineGetResponse GetEngine(ctx, path)
Get Engine

Retrieve engine by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Name of engine | 

### Return type

[**EngineGetResponse**](EngineGetResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEngines**
> EngineListResult ListEngines(ctx, )
List Engines

List all existing engines.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EngineListResult**](EngineListResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingEngine**
> EnginePingResponse PingEngine(ctx, path)
Ping Engine

Sends a message to the engine to validate connectivity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Name of engine | 

### Return type

[**EnginePingResponse**](EnginePingResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

