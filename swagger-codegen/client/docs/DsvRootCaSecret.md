# DsvRootCaSecret

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Certificate of the root CA that contains information about it and public key | [default to null]
**Crl** | **string** | URL of the CRL from which the revocation of leaf certificates can be checked | [optional] [default to null]
**Domains** | **[]string** | List of domains for which certificate signing is allowed | [default to null]
**MaxTTL** | **int64** | Maximum TTL of a signed certificate issued from a given root CA (in hours) | [default to null]
**PrivateKey** | **string** | Private key of the root CA | [default to null]
**RootCAPath** | **string** | RootCAPath to secret, which also serves as an identifier of the root CA | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


