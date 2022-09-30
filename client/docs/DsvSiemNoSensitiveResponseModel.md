# DsvSiemNoSensitiveResponseModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethod** | **string** | Authentication method (token) | [optional] [default to null]
**Endpoint** | **string** | Endpoint | [optional] [default to null]
**Failed** | **bool** | Failed is true if send has failed too many times, false otherwise | [optional] [default to null]
**FailedEvents** | **int64** | Number of failed send events | [optional] [default to null]
**Host** | **string** | Collect Server IP/FQDN | [optional] [default to null]
**Id** | **string** | The unique id for this item | [optional] [default to null]
**LoggingFormat** | **string** | Logging format (e.g. \&quot;rfc5424\&quot; for syslog) | [optional] [default to null]
**Name** | **string** | Name of registered SIEM endpoint, similar to path | [optional] [default to null]
**Pool** | **string** | Engine pool name, used when sending request to a DSV engine instance | [optional] [default to null]
**Port** | **int64** | Collect Server Port | [optional] [default to null]
**Protocol** | **string** | Type of protocol (\&quot;tcp\&quot;, \&quot;udp\&quot;, \&quot;http\&quot;, \&quot;https\&quot;, \&quot;tls\&quot;) | [optional] [default to null]
**SendToEngine** | **bool** | Denotes whether the endpoint should be accessed through a DSV engine instance | [optional] [default to null]
**SiemType** | **string** | Type of endpoint (\&quot;syslog\&quot;, \&quot;cef\&quot;, \&quot;json\&quot;, \&quot;splunk\&quot;) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


