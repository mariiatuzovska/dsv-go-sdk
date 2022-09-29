# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitCertAuth**](TokensApi.md#InitCertAuth) | **Post** /certificate/auth | Initiate authentication by certificate
[**Revoke**](TokensApi.md#Revoke) | **Post** /revoke/{refreshtoken} | Revoke Refresh Token
[**Token**](TokensApi.md#Token) | **Post** /token | Authenticate

# **InitCertAuth**
> DsvInitiateCertAuthResponse InitCertAuth(ctx, clientCertificate)
Initiate authentication by certificate

Request a challenge to decrypt to prove ownership of the private key for authentication by certificate flow. Challenge id returned in response is only valid for 5 minutes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientCertificate** | **string**|  | 

### Return type

[**DsvInitiateCertAuthResponse**](InitiateCertAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Revoke**
> Revoke(ctx, refreshtoken)
Revoke Refresh Token

Revoke an existing refresh token to prevent it from being used for authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refreshtoken** | **string**| Refresh token to revoke | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Token**
> DsvAccessTokenResponse Token(ctx, grantType, username, password, provider, awsBody, awsHeaders, jwt, clientId, clientSecret, refreshToken, certChallengeId, decryptedChallenge)
Authenticate

Submit parameters to get a new access token for use on protected endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | **string**|  | 
  **username** | **string**|  | 
  **password** | **string**|  | 
  **provider** | **string**|  | 
  **awsBody** | **string**|  | 
  **awsHeaders** | **string**|  | 
  **jwt** | **string**|  | 
  **clientId** | **string**|  | 
  **clientSecret** | **string**|  | 
  **refreshToken** | **string**|  | 
  **certChallengeId** | **string**|  | 
  **decryptedChallenge** | **string**|  | 

### Return type

[**DsvAccessTokenResponse**](AccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

