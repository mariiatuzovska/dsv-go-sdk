# \UsageApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsage**](UsageApi.md#GetUsage) | **Get** /usage | Get Usage


# **GetUsage**
> DsvUsageResponseGeneral GetUsage(ctx, startDate, optional)
Get Usage

Get usage statistics for API calls and total number of secrets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **string**| Start date to get usage statistics | 
 **optional** | ***UsageApiGetUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiGetUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.String**| End date to get usage statistics, defaults to current date if not set | 

### Return type

[**DsvUsageResponseGeneral**](UsageResponseGeneral.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

