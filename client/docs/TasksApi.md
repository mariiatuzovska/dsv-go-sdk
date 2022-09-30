# \TasksApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaskStatus**](TasksApi.md#GetTaskStatus) | **Get** /task/status/{id} | Get background task status


# **GetTaskStatus**
> DsvTaskState GetTaskStatus(ctx, id)
Get background task status

Get background task status by task uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Task status uuid | 

### Return type

[**DsvTaskState**](TaskState.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

