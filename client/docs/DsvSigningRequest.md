# DsvSigningRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **bool** | Boolean indicating whether to return a root certificate | [optional] [default to null]
**Csr** | **string** | Certificate Signing Request | [default to null]
**RootCAPath** | **string** | Path to secret - registered root CA | [default to null]
**SubjectAltNames** | **[]string** | A list of Subject Alternative Names for a certificate (each domain must be present in the list of allowed domains) | [optional] [default to null]
**Ttl** | **int64** | TTL for a generated certificate (in hours, cannot exceed the maximum TTL specified in root CA secret) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


