# \EaaSManualApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecryptWithManualKey**](EaaSManualApi.md#DecryptWithManualKey) | **Post** /crypto/manual/decrypt | Decrypt
[**DeleteManualKey**](EaaSManualApi.md#DeleteManualKey) | **Delete** /crypto/manual/key/{path} | Delete Key
[**EncryptWithManualKey**](EaaSManualApi.md#EncryptWithManualKey) | **Post** /crypto/manual/encrypt | Encrypt
[**ReadManualKey**](EaaSManualApi.md#ReadManualKey) | **Get** /crypto/manual/key/{path} | Read Key
[**RestoreManualKey**](EaaSManualApi.md#RestoreManualKey) | **Put** /crypto/manual/key/{path}/restore | Restore Key
[**UpdateKey**](EaaSManualApi.md#UpdateKey) | **Put** /crypto/manual/key/{path} | Update Key
[**UploadKey**](EaaSManualApi.md#UploadKey) | **Post** /crypto/manual/key/{path} | Upload Key


# **DecryptWithManualKey**
> DsvDecryptionResponse DecryptWithManualKey(ctx, path, ciphertext, optional)
Decrypt

Decrypt ciphertext with a key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path to data key | 
  **ciphertext** | [**[]int32**](int32.md)| A value to be decrypted | 
 **optional** | ***EaaSManualApiDecryptWithManualKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EaaSManualApiDecryptWithManualKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of the key with which to decrypt data | 

### Return type

[**DsvDecryptionResponse**](DecryptionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManualKey**
> DsvMessageResponse DeleteManualKey(ctx, path)
Delete Key

Delete an existing encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EncryptWithManualKey**
> DsvEncryptionResponse EncryptWithManualKey(ctx, path, plaintext, optional)
Encrypt

Encrypt plaintext with a key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path to data key | 
  **plaintext** | **string**| A value to be encrypted | 
 **optional** | ***EaaSManualApiEncryptWithManualKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EaaSManualApiEncryptWithManualKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of the key with which to encrypt data | 

### Return type

[**DsvEncryptionResponse**](EncryptionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadManualKey**
> DsvSecretResponse ReadManualKey(ctx, path)
Read Key

Read existing encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreManualKey**
> RestoreManualKey(ctx, path)
Restore Key

Restore a soft-deleted encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKey**
> DsvSecretResponse UpdateKey(ctx, body, path)
Update Key

Update an existing encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvUpdateKeyRequest**](DsvUpdateKeyRequest.md)|  | 
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadKey**
> DsvSecretResponse UploadKey(ctx, body, path)
Upload Key

Upload a new encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvManualKeyData**](DsvManualKeyData.md)|  | 
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**DsvSecretResponse**](SecretResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

