# RootCaRegistration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | **string** |  | [default to null]
**Country** | **string** |  | [optional] [default to null]
**Crl** | **string** | URL of the CRL from which the revocation of leaf certificates can be checked | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Domains** | **[]string** | List of domains for which certificate signing is allowed | [default to null]
**EmailAddress** | **string** |  | [optional] [default to null]
**Locality** | **string** |  | [optional] [default to null]
**MaxTTL** | **int64** | Maximum TTL of a signed certificate issued from a given root CA (in hours) | [default to null]
**Organization** | **string** |  | [optional] [default to null]
**OrganizationalUnit** | **string** |  | [optional] [default to null]
**RootCAPath** | **string** | The name of the secret containing the root CA certificate | [default to null]
**State** | **string** |  | [optional] [default to null]
**StorePath** | **string** | The name of the secret in which to store the generated certificate and private key | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

