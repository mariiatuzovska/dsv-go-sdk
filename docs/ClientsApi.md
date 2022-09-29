# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientCredential**](ClientsApi.md#CreateClientCredential) | **Post** /clients | Create a Client Credential
[**DeleteClientCredential**](ClientsApi.md#DeleteClientCredential) | **Delete** /clients/{clientId} | Delete a Client Credential
[**GetBootstrapClientCredential**](ClientsApi.md#GetBootstrapClientCredential) | **Get** /clients/bootstrap/{clientId} | Get a Client Bootstrap Credential including secret
[**GetClientCredential**](ClientsApi.md#GetClientCredential) | **Get** /clients/{clientId} | Get a Client Credential
[**RestoreClient**](ClientsApi.md#RestoreClient) | **Get** /clients/{clientId}/restore | Restore a Client
[**SearchClients**](ClientsApi.md#SearchClients) | **Get** /clients | Search for Client Credentials

# **CreateClientCredential**
> ClientCredentialsResponse CreateClientCredential(ctx, body)
Create a Client Credential

Request a new client credential for a role and get back the client id and secret key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCreate**](ClientCreate.md)|  | 

### Return type

[**ClientCredentialsResponse**](ClientCredentialsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClientCredential**
> MessageResponse DeleteClientCredential(ctx, clientId, optional)
Delete a Client Credential

Delete a client credential by its unique client id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| ClientId property of the client credentials | 
 **optional** | ***ClientsApiDeleteClientCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiDeleteClientCredentialOpts struct
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

# **GetBootstrapClientCredential**
> ClientCredentialsResponse GetBootstrapClientCredential(ctx, )
Get a Client Bootstrap Credential including secret

Get a client credential by url.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClientCredentialsResponse**](ClientCredentialsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCredential**
> ClientCredentialsResponse GetClientCredential(ctx, clientId)
Get a Client Credential

Get a client credential by client id. The secret key is omitted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| ClientId property of the client credentials | 

### Return type

[**ClientCredentialsResponse**](ClientCredentialsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreClient**
> RestoreClient(ctx, clientId)
Restore a Client

Restore a client by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| ClientId property of the client credentials | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchClients**
> ClientSearchModel SearchClients(ctx, role, optional)
Search for Client Credentials

Search for one or more client credentials associated with a particular role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| Role name associated with client credentials | 
 **optional** | ***ClientsApiSearchClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiSearchClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| The maximum number of results per cursor | 
 **cursor** | **optional.String**| Cursor to next batch of results | 

### Return type

[**ClientSearchModel**](ClientSearchModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

