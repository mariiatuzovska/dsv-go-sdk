# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeafParams**](PKIApi.md#LeafParams) | **Post** /pki/leaf | Create Leaf Certificate and Private Key
[**RegisterRoot**](PKIApi.md#RegisterRoot) | **Post** /pki/register | Register Root CA
[**RootCARegistration**](PKIApi.md#RootCARegistration) | **Post** /pki/root | Generate Root Certificate
[**SSHCertParams**](PKIApi.md#SSHCertParams) | **Post** /pki/ssh-cert | Create SSH Certificate
[**SignCertificate**](PKIApi.md#SignCertificate) | **Post** /pki/sign | Create Signed Certificate

# **LeafParams**
> ResponseCertificate LeafParams(ctx, body)
Create Leaf Certificate and Private Key

Create and return a signed certificate with a private key based on a registered root CA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SigningRequestInformation**](SigningRequestInformation.md)|  | 

### Return type

[**ResponseCertificate**](ResponseCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterRoot**
> ResponseRootCertificate RegisterRoot(ctx, body)
Register Root CA

Register a root CA as a secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RootCaSecret**](RootCaSecret.md)|  | 

### Return type

[**ResponseRootCertificate**](ResponseRootCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RootCARegistration**
> ResponseCertificate RootCARegistration(ctx, body)
Generate Root Certificate

Create and return a new root certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RootCaRegistration**](RootCaRegistration.md)|  | 

### Return type

[**ResponseCertificate**](ResponseCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSHCertParams**
> ResponseSshCertificate SSHCertParams(ctx, body)
Create SSH Certificate

Create, store and return a signed SSH certificate using a root private key and SHH-compatible leaf public key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SshCertInformation**](SshCertInformation.md)|  | 

### Return type

[**ResponseSshCertificate**](ResponseSSHCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignCertificate**
> SignedLeafCertificate SignCertificate(ctx, body)
Create Signed Certificate

Create and return a signed certificate based on a registered root CA with a given CSR.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SigningRequest**](SigningRequest.md)|  | 

### Return type

[**SignedLeafCertificate**](SignedLeafCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

