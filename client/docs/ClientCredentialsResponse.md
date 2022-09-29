# ClientCredentialsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TTL** | **int64** | TTL expiration in seconds | [optional] [default to null]
**Accessed** | **string** | Url already used or not | [optional] [default to null]
**ClientId** | **string** | Unique uuid of client credentials | [optional] [default to null]
**ClientSecret** | **string** | Secret key returned on create | [optional] [default to null]
**Created** | **string** | Created date | [optional] [default to null]
**CreatedBy** | **string** | Who created | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**ExpiredAt** | **string** | ExpiredAt expiration time | [optional] [default to null]
**Id** | **string** | the id for this item | [optional] [default to null]
**Role** | **string** | Assigned role for determining access | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Url** | **bool** | If Url requested | [optional] [default to null]
**UrlExpires** | **string** | Url expiration time | [optional] [default to null]
**UrlPath** | **string** | URL Path | [optional] [default to null]
**UrlTTL** | **int64** | Url expiration in seconds | [optional] [default to null]
**UsedCount** | **int64** |  | [optional] [default to null]
**UsesLimit** | **int64** | Uses  the number of times the client credential can be read.  if set to 0, it can be used infinitely.  default is 0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

