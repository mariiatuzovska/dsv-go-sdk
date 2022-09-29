# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSiems**](SIEMApi.md#SearchSiems) | **Get** /config/siem | Search SIEM Endpoints
[**SiemCreate**](SIEMApi.md#SiemCreate) | **Post** /config/siem | Create SIEM Endpoint
[**SiemDelete**](SIEMApi.md#SiemDelete) | **Delete** /config/siem/{name} | Delete SIEM Endpoint
[**SiemGet**](SIEMApi.md#SiemGet) | **Get** /config/siem/{name} | Get SIEM Endpoint
[**SiemUpdate**](SIEMApi.md#SiemUpdate) | **Put** /config/siem/{name} | Update SIEM Endpoint

# **SearchSiems**
> SiemSearchResponse SearchSiems(ctx, optional)
Search SIEM Endpoints

Search SIEM Endpoints by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SIEMApiSearchSiemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIEMApiSearchSiemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| Partial search term for search by name | 
 **limit** | **optional.Int64**| Limit for the number of results per page (cursor) | 
 **cursor** | **optional.String**| Cursor to next batch of results | 
 **sort** | **optional.String**| Sort results ascending (asc) or descending (desc) order by name. Default is asc | 

### Return type

[**SiemSearchResponse**](SiemSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiemCreate**
> SiemResponse SiemCreate(ctx, body)
Create SIEM Endpoint

Creates a new SIEM endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiemCreateUpdateRequestModel**](SiemCreateUpdateRequestModel.md)|  | 

### Return type

[**SiemResponse**](SiemResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiemDelete**
> SiemDelete(ctx, name)
Delete SIEM Endpoint

Delete an existing SIEM endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiemGet**
> SiemResponse SiemGet(ctx, name)
Get SIEM Endpoint

Retrieve an existing SIEM endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**SiemResponse**](SiemResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiemUpdate**
> SiemResponse SiemUpdate(ctx, body, name)
Update SIEM Endpoint

Update an existing SIEM endpoint. Update sets \"failed\" to \"false\" and \"failedEvents\" to \"0\" automatically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiemCreateUpdateRequestModel**](SiemCreateUpdateRequestModel.md)|  | 
  **name** | **string**|  | 

### Return type

[**SiemResponse**](SiemResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

