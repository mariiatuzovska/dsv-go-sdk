# DsvSigningRequestInformation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **bool** | Boolean indicating whether to return a root certificate | [optional] [default to null]
**CommonName** | **string** |  | [default to null]
**Country** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**EmailAddress** | **string** |  | [optional] [default to null]
**Locality** | **string** |  | [optional] [default to null]
**Organization** | **string** |  | [optional] [default to null]
**OrganizationalUnit** | **string** |  | [optional] [default to null]
**RootCAPath** | **string** | The name of the secret containing the root CA certificate | [default to null]
**State** | **string** |  | [optional] [default to null]
**StorePath** | **string** | The name of the secret in which to store the generated certificate and private key | [optional] [default to null]
**Ttl** | **int64** | TTL for a generated certificate (in hours, cannot exceed the maximum TTL specified in root CA secret) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


