# \AuditApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadAudit**](AuditApi.md#DownloadAudit) | **Get** /download/audit | Download Audit Records
[**FindAudit**](AuditApi.md#FindAudit) | **Get** /audit | Find Audit Records


# **DownloadAudit**
> string DownloadAudit(ctx, startDate, endDate)
Download Audit Records

Download a zip of audit records for a time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **string**| The start date to find audits from | 
  **endDate** | **string**| The end date to find audits to | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAudit**
> DsvLogSearchResponse FindAudit(ctx, optional)
Find Audit Records

Find audit records based on search critera.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditApiFindAuditOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditApiFindAuditOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principal** | **optional.String**| Security principal name | 
 **startDate** | **optional.Time**| The start date to find audits from | 
 **endDate** | **optional.Time**| The end date to find audits to | 
 **action** | **optional.String**| The audit action | 
 **path** | **optional.String**| The secret path | 
 **cursor** | **optional.String**| The cursor for pagination | 
 **limit** | **optional.Int64**| The number of results to return | 
 **sort** | **optional.String**| Sort results ascending (asc) or descending (desc) order. Default is desc | 

### Return type

[**DsvLogSearchResponse**](LogSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

