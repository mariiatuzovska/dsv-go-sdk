# \TokensApi

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
  **clientCertificate** | **string**| Base64 encoded client certificate in PEM format | 

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Token**
> DsvAccessTokenResponse Token(ctx, grantType, optional)
Authenticate

Submit parameters to get a new access token for use on protected endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | **string**|  | 
 **optional** | ***TokensApiTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokensApiTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Username for password grant type | 
 **password** | **optional.String**| Password for password grant type | 
 **provider** | **optional.String**| Provider name for password grant flow for Thycotic One authentication | 
 **awsBody** | **optional.String**| Base64 encoded signed aws request body for aws_iam grant type | 
 **awsHeaders** | **optional.String**| Base64 encoded signed aws request headers for aws_iam grant type | 
 **jwt** | **optional.String**| JWT token for azure and gcp grant types | 
 **clientId** | **optional.String**| Client id for client_credentials grant type | 
 **clientSecret** | **optional.String**| Client secret for client_credentials grant type | 
 **refreshToken** | **optional.String**| Previously issued refresh token for the refresh_token grant type | 
 **certChallengeId** | **optional.String**| Challenge id for the certificate grant type | 
 **decryptedChallenge** | **optional.String**| Decrypted and base64 encoded challenge for the certificate grant type | 

### Return type

[**DsvAccessTokenResponse**](AccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

