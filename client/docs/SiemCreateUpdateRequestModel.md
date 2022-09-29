# SiemCreateUpdateRequestModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **string** | Authentication token | [default to null]
**AuthMethod** | **string** | Authentication method (token) | [default to null]
**Endpoint** | **string** | Endpoint | [optional] [default to null]
**Host** | **string** | Collect Server IP/FQDN | [default to null]
**LoggingFormat** | **string** | Logging Format (i.e. syslog (RFC 5424)) | [default to null]
**Name** | **string** | Name of registered SIEM endpoint, similar to path | [default to null]
**Pool** | **string** | Engine pool name, used when sending request to a DSV engine instance | [optional] [default to null]
**Port** | **int64** | Port | [default to null]
**Protocol** | **string** | Type of protocol (i.e. TCP, UDP) | [default to null]
**SendToEngine** | **bool** | Denotes whether the endpoint should be accessed through a DSV engine instance | [optional] [default to null]
**SiemType** | **string** | Type of endpoint (\&quot;syslog\&quot;, \&quot;cef\&quot;, \&quot;json\&quot;, \&quot;splunk\&quot;) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

