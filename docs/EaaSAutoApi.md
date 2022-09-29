# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](EaaSAutoApi.md#CreateKey) | **Post** /crypto/key/{path} | Create Key
[**Decrypt**](EaaSAutoApi.md#Decrypt) | **Post** /crypto/decrypt | Decrypt
[**DeleteKey**](EaaSAutoApi.md#DeleteKey) | **Delete** /crypto/key/{path} | Delete Key
[**Encrypt**](EaaSAutoApi.md#Encrypt) | **Post** /crypto/encrypt | Encrypt
[**GetKeyMetadata**](EaaSAutoApi.md#GetKeyMetadata) | **Get** /crypto/key/{path} | Get Key Metadata
[**RestoreKey**](EaaSAutoApi.md#RestoreKey) | **Put** /crypto/key/{path}/restore | Restore Key
[**RotationRequest**](EaaSAutoApi.md#RotationRequest) | **Post** /crypto/rotate | Rotate Data and Key

# **CreateKey**
> AutoKeyResponse CreateKey(ctx, path)
Create Key

Create a new encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**AutoKeyResponse**](AutoKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Decrypt**
> DecryptionResponse Decrypt(ctx, path, ciphertext, optional)
Decrypt

Decrypt ciphertext with a key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path to data key | 
  **ciphertext** | [**[]int32**](int32.md)| A value to be decrypted | 
 **optional** | ***EaaSAutoApiDecryptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EaaSAutoApiDecryptOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of the key with which to decrypt data | 

### Return type

[**DecryptionResponse**](DecryptionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKey**
> MessageResponse DeleteKey(ctx, path)
Delete Key

Delete an existing encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Encrypt**
> EncryptionResponse Encrypt(ctx, path, plaintext, optional)
Encrypt

Encrypt plaintext with a key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path to data key | 
  **plaintext** | **string**| A value to be encrypted | 
 **optional** | ***EaaSAutoApiEncryptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EaaSAutoApiEncryptOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of the key with which to encrypt data | 

### Return type

[**EncryptionResponse**](EncryptionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyMetadata**
> AutoKeyResponse GetKeyMetadata(ctx, path)
Get Key Metadata

Get metadata of an existing encryption/decryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The full key path, for example, mykeys/key01 | 

### Return type

[**AutoKeyResponse**](AutoKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreKey**
> RestoreKey(ctx, path)
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotationRequest**
> EncryptionResponse RotationRequest(ctx, path, ciphertext, startingVersion, optional)
Rotate Data and Key

Rotate data and optionally an existing encryption/decryption key.  If the starting version is the current version of the key, then DSV will rotate the key (create a new version of it) and re-encrypt the data using this new version.  If the starting version is NOT the current version of the key, and the ending version is not provided, then DSV will only re-encrypt the data using the current latest version of the key.  The starting and ending versions can also be below the latest one, so long as the starting is below the ending. In this case, DSV will re-encrypt the data using the version of the key specified by the ending version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path to data key | 
  **ciphertext** | [**[]int32**](int32.md)| A value to be rotated (re-encrypted) | 
  **startingVersion** | **string**| The starting version of the key with which to re-encrypt data | 
 **optional** | ***EaaSAutoApiRotationRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EaaSAutoApiRotationRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **endingVersion** | **optional.String**| The ending version of the key with which to re-encrypt data | 

### Return type

[**EncryptionResponse**](EncryptionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

