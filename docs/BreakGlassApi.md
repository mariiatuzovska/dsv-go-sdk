# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyRequest**](BreakGlassApi.md#ApplyRequest) | **Post** /breakglass/apply | Apply
[**GenerateRequest**](BreakGlassApi.md#GenerateRequest) | **Post** /breakglass/generate | Generate
[**GetStatus**](BreakGlassApi.md#GetStatus) | **Get** /breakglass | Get Status

# **ApplyRequest**
> ApplyResponse ApplyRequest(ctx, shares)
Apply

Apply secret shares to break glass and give users admin rights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shares** | [**[]string**](string.md)|  | 

### Return type

[**ApplyResponse**](ApplyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateRequest**
> GenerateResponse GenerateRequest(ctx, newAdmins, minNumberOfShares)
Generate

Generate a new break glass secret and shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newAdmins** | [**[]string**](string.md)|  | 
  **minNumberOfShares** | **int64**|  | 

### Return type

[**GenerateResponse**](GenerateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> StatusResponse GetStatus(ctx, )
Get Status

Get break glass status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

