# TokenBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | [default to null]
**Username** | **string** | Username for password grant type | [optional] [default to null]
**Password** | **string** | Password for password grant type | [optional] [default to null]
**Provider** | **string** | Provider name for password grant flow for Thycotic One authentication | [optional] [default to null]
**AwsBody** | **string** | Base64 encoded signed aws request body for aws_iam grant type | [optional] [default to null]
**AwsHeaders** | **string** | Base64 encoded signed aws request headers for aws_iam grant type | [optional] [default to null]
**Jwt** | **string** | JWT token for azure and gcp grant types | [optional] [default to null]
**ClientId** | **string** | Client id for client_credentials grant type | [optional] [default to null]
**ClientSecret** | **string** | Client secret for client_credentials grant type | [optional] [default to null]
**RefreshToken** | **string** | Previously issued refresh token for the refresh_token grant type | [optional] [default to null]
**CertChallengeId** | **string** | Challenge id for the certificate grant type | [optional] [default to null]
**DecryptedChallenge** | **string** | Decrypted and base64 encoded challenge for the certificate grant type | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

