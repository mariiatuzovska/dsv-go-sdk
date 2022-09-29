# DSV Go SDK

The purpose of this application is to provide a simple service for storing and getting secrets.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Generate SDK

1. Install [swagger-codegen](https://github.com/swagger-api/swagger-codegen).
2. Move `swagger.json` to the current folder.
3. Generate models: `swagger-codegen generate -i ./swagger.json -l go -o ./models --model-name-prefix DSV -Dmodels`.
4. Generate API client: `swagger-codegen generate -i ./swagger.json -l go -o ./client`.

## Documentation for API Endpoints

All URIs are relative to *https://secretsvaultcloud.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditApi* | [**DownloadAudit**](client/docs/AuditApi.md#downloadaudit) | **Get** /download/audit | Download Audit Records
*AuditApi* | [**FindAudit**](client/docs/AuditApi.md#findaudit) | **Get** /audit | Find Audit Records
*BreakGlassApi* | [**ApplyRequest**](client/docs/BreakGlassApi.md#applyrequest) | **Post** /breakglass/apply | Apply
*BreakGlassApi* | [**GenerateRequest**](client/docs/BreakGlassApi.md#generaterequest) | **Post** /breakglass/generate | Generate
*BreakGlassApi* | [**GetStatus**](client/docs/BreakGlassApi.md#getstatus) | **Get** /breakglass | Get Status
*ClientsApi* | [**CreateClientCredential**](client/docs/ClientsApi.md#createclientcredential) | **Post** /clients | Create a Client Credential
*ClientsApi* | [**DeleteClientCredential**](client/docs/ClientsApi.md#deleteclientcredential) | **Delete** /clients/{clientId} | Delete a Client Credential
*ClientsApi* | [**GetBootstrapClientCredential**](client/docs/ClientsApi.md#getbootstrapclientcredential) | **Get** /clients/bootstrap/{clientId} | Get a Client Bootstrap Credential including secret
*ClientsApi* | [**GetClientCredential**](client/docs/ClientsApi.md#getclientcredential) | **Get** /clients/{clientId} | Get a Client Credential
*ClientsApi* | [**RestoreClient**](client/docs/ClientsApi.md#restoreclient) | **Get** /clients/{clientId}/restore | Restore a Client
*ClientsApi* | [**SearchClients**](client/docs/ClientsApi.md#searchclients) | **Get** /clients | Search for Client Credentials
*ConfigApi* | [**GetConfig**](client/docs/ConfigApi.md#getconfig) | **Get** /config | Get Config
*ConfigApi* | [**GetConfigByVersion**](client/docs/ConfigApi.md#getconfigbyversion) | **Get** /config/version/{version} | Get Config By Version
*ConfigApi* | [**PostConfig**](client/docs/ConfigApi.md#postconfig) | **Post** /config | Create Config
*EaaSAutoApi* | [**CreateKey**](client/docs/EaaSAutoApi.md#createkey) | **Post** /crypto/key/{path} | Create Key
*EaaSAutoApi* | [**Decrypt**](client/docs/EaaSAutoApi.md#decrypt) | **Post** /crypto/decrypt | Decrypt
*EaaSAutoApi* | [**DeleteKey**](client/docs/EaaSAutoApi.md#deletekey) | **Delete** /crypto/key/{path} | Delete Key
*EaaSAutoApi* | [**Encrypt**](client/docs/EaaSAutoApi.md#encrypt) | **Post** /crypto/encrypt | Encrypt
*EaaSAutoApi* | [**GetKeyMetadata**](client/docs/EaaSAutoApi.md#getkeymetadata) | **Get** /crypto/key/{path} | Get Key Metadata
*EaaSAutoApi* | [**RestoreKey**](client/docs/EaaSAutoApi.md#restorekey) | **Put** /crypto/key/{path}/restore | Restore Key
*EaaSAutoApi* | [**RotationRequest**](client/docs/EaaSAutoApi.md#rotationrequest) | **Post** /crypto/rotate | Rotate Data and Key
*EaaSManualApi* | [**DecryptWithManualKey**](client/docs/EaaSManualApi.md#decryptwithmanualkey) | **Post** /crypto/manual/decrypt | Decrypt
*EaaSManualApi* | [**DeleteManualKey**](client/docs/EaaSManualApi.md#deletemanualkey) | **Delete** /crypto/manual/key/{path} | Delete Key
*EaaSManualApi* | [**EncryptWithManualKey**](client/docs/EaaSManualApi.md#encryptwithmanualkey) | **Post** /crypto/manual/encrypt | Encrypt
*EaaSManualApi* | [**ReadManualKey**](client/docs/EaaSManualApi.md#readmanualkey) | **Get** /crypto/manual/key/{path} | Read Key
*EaaSManualApi* | [**RestoreManualKey**](client/docs/EaaSManualApi.md#restoremanualkey) | **Put** /crypto/manual/key/{path}/restore | Restore Key
*EaaSManualApi* | [**UpdateKey**](client/docs/EaaSManualApi.md#updatekey) | **Put** /crypto/manual/key/{path} | Update Key
*EaaSManualApi* | [**UploadKey**](client/docs/EaaSManualApi.md#uploadkey) | **Post** /crypto/manual/key/{path} | Upload Key
*EnginesApi* | [**CreateEngine**](client/docs/EnginesApi.md#createengine) | **Post** /engines | Create Engine
*EnginesApi* | [**DeleteEngine**](client/docs/EnginesApi.md#deleteengine) | **Delete** /engines/{name} | Delete Engine
*EnginesApi* | [**GetEngine**](client/docs/EnginesApi.md#getengine) | **Get** /engines/{name} | Get Engine
*EnginesApi* | [**ListEngines**](client/docs/EnginesApi.md#listengines) | **Get** /engines | List Engines
*EnginesApi* | [**PingEngine**](client/docs/EnginesApi.md#pingengine) | **Post** /engines/{name}/ping | Ping Engine
*GroupsApi* | [**AddMember**](client/docs/GroupsApi.md#addmember) | **Post** /groups/{name}/members | Add Members To Group
*GroupsApi* | [**CreateGroup**](client/docs/GroupsApi.md#creategroup) | **Post** /groups/ | Create Group
*GroupsApi* | [**DeleteGroup**](client/docs/GroupsApi.md#deletegroup) | **Delete** /groups/{name} | Delete Group
*GroupsApi* | [**DeleteMember**](client/docs/GroupsApi.md#deletemember) | **Delete** /groups/{name}/members | Delete Members From Group
*GroupsApi* | [**GetGroup**](client/docs/GroupsApi.md#getgroup) | **Get** /groups/{name} | Get Group
*GroupsApi* | [**RestoreGroup**](client/docs/GroupsApi.md#restoregroup) | **Get** /groups/{name}/restore | Restore Group
*GroupsApi* | [**SearchGroups**](client/docs/GroupsApi.md#searchgroups) | **Get** /groups | Search Groups
*HomeApi* | [**CreateHomeSecret**](client/docs/HomeApi.md#createhomesecret) | **Post** /home/{principalName}/{path} | Create home secrets
*HomeApi* | [**DeleteHomeSecret**](client/docs/HomeApi.md#deletehomesecret) | **Delete** /home/{principalName}/{path} | Delete Home Secret
*HomeApi* | [**GetHomeSecret**](client/docs/HomeApi.md#gethomesecret) | **Get** /home/{principalName}/{path} | Get Home
*HomeApi* | [**GetHomeSecretByVersion**](client/docs/HomeApi.md#gethomesecretbyversion) | **Get** /home/{principalName}/{path}/version/{version} | Get Home Secret By Version
*HomeApi* | [**GetHomeSecretDescription**](client/docs/HomeApi.md#gethomesecretdescription) | **Get** /home/{principalName}/{path}::description | Get home Secret Description
*HomeApi* | [**RestoreHomeSecret**](client/docs/HomeApi.md#restorehomesecret) | **Get** /home/{principalName}/{path}/restore | Restore home Secret
*HomeApi* | [**RollbackHomeSecret**](client/docs/HomeApi.md#rollbackhomesecret) | **Put** /home/{principalName}/{path}/rollback/{version} | Rollback a Home Secret
*HomeApi* | [**SearchHomeSecrets**](client/docs/HomeApi.md#searchhomesecrets) | **Get** /home/{principalName} | Search for Home Secrets
*HomeApi* | [**UpdateHomeSecret**](client/docs/HomeApi.md#updatehomesecret) | **Put** /home/{principalName}/{path} | Update home Secret
*KeyApi* | [**Masterkeys**](client/docs/KeyApi.md#masterkeys) | **Put** /config/keys | Update Master Key
*PKIApi* | [**LeafParams**](client/docs/PKIApi.md#leafparams) | **Post** /pki/leaf | Create Leaf Certificate and Private Key
*PKIApi* | [**RegisterRoot**](client/docs/PKIApi.md#registerroot) | **Post** /pki/register | Register Root CA
*PKIApi* | [**RootCARegistration**](client/docs/PKIApi.md#rootcaregistration) | **Post** /pki/root | Generate Root Certificate
*PKIApi* | [**SSHCertParams**](client/docs/PKIApi.md#sshcertparams) | **Post** /pki/ssh-cert | Create SSH Certificate
*PKIApi* | [**SignCertificate**](client/docs/PKIApi.md#signcertificate) | **Post** /pki/sign | Create Signed Certificate
*PoliciesApi* | [**CreatePolicy**](client/docs/PoliciesApi.md#createpolicy) | **Post** /config/policies/ | Create Policy
*PoliciesApi* | [**DeletePolicy**](client/docs/PoliciesApi.md#deletepolicy) | **Delete** /config/policies/{path} | Delete Policy
*PoliciesApi* | [**GetPolicy**](client/docs/PoliciesApi.md#getpolicy) | **Get** /config/policies/{path} | Get Policy
*PoliciesApi* | [**GetPolicyByVersion**](client/docs/PoliciesApi.md#getpolicybyversion) | **Get** /config/policies/{path}/version/{version} | Get a list of policies by version
*PoliciesApi* | [**RestorePolicy**](client/docs/PoliciesApi.md#restorepolicy) | **Get** /config/policies/{path}/restore | Restore Policy
*PoliciesApi* | [**RollbackPolicy**](client/docs/PoliciesApi.md#rollbackpolicy) | **Put** /config/policies/{path}/rollback/{version} | Rollback Policy
*PoliciesApi* | [**SearchFilter**](client/docs/PoliciesApi.md#searchfilter) | **Get** /config/policies | Search Policies
*PoliciesApi* | [**UpdatePolicy**](client/docs/PoliciesApi.md#updatepolicy) | **Put** /config/policies/{path} | Update Policy
*PoolsApi* | [**CreatePool**](client/docs/PoolsApi.md#createpool) | **Post** /pools | Create Pool
*PoolsApi* | [**DeletePool**](client/docs/PoolsApi.md#deletepool) | **Delete** /pools/{name} | Delete Pool
*PoolsApi* | [**GetPool**](client/docs/PoolsApi.md#getpool) | **Get** /pools/{name} | Get Pool
*PoolsApi* | [**ListPools**](client/docs/PoolsApi.md#listpools) | **Get** /pools | List Pools
*RolesApi* | [**CreateRole**](client/docs/RolesApi.md#createrole) | **Post** /roles/ | Create a Role
*RolesApi* | [**DeleteRole**](client/docs/RolesApi.md#deleterole) | **Delete** /roles/{name} | Delete a Role
*RolesApi* | [**GetRole**](client/docs/RolesApi.md#getrole) | **Get** /roles/{name} | Get a Role
*RolesApi* | [**GetRoleByVersion**](client/docs/RolesApi.md#getrolebyversion) | **Get** /roles/{name}/version/{version} | Get a Role By Version
*RolesApi* | [**RestoreRole**](client/docs/RolesApi.md#restorerole) | **Get** /roles/{name}/restore | Restore a Role
*RolesApi* | [**SearchRoles**](client/docs/RolesApi.md#searchroles) | **Get** /roles | Search for Roles
*RolesApi* | [**UpdateRole**](client/docs/RolesApi.md#updaterole) | **Put** /roles/{name} | Update a Role
*SIEMApi* | [**SearchSiems**](client/docs/SIEMApi.md#searchsiems) | **Get** /config/siem | Search SIEM Endpoints
*SIEMApi* | [**SiemCreate**](client/docs/SIEMApi.md#siemcreate) | **Post** /config/siem | Create SIEM Endpoint
*SIEMApi* | [**SiemDelete**](client/docs/SIEMApi.md#siemdelete) | **Delete** /config/siem/{name} | Delete SIEM Endpoint
*SIEMApi* | [**SiemGet**](client/docs/SIEMApi.md#siemget) | **Get** /config/siem/{name} | Get SIEM Endpoint
*SIEMApi* | [**SiemUpdate**](client/docs/SIEMApi.md#siemupdate) | **Put** /config/siem/{name} | Update SIEM Endpoint
*SecretsApi* | [**CreateSecret**](client/docs/SecretsApi.md#createsecret) | **Post** /secrets/{path} | Create Secret
*SecretsApi* | [**DeleteSecret**](client/docs/SecretsApi.md#deletesecret) | **Delete** /secrets/{path} | Delete Secret
*SecretsApi* | [**GetSecret**](client/docs/SecretsApi.md#getsecret) | **Get** /secrets/{path} | Get Secret
*SecretsApi* | [**GetSecretByVersion**](client/docs/SecretsApi.md#getsecretbyversion) | **Get** /secrets/{path}/version/{version} | Get Secret By Version
*SecretsApi* | [**GetSecretDescription**](client/docs/SecretsApi.md#getsecretdescription) | **Get** /secrets/{path}::description | Get Secret Description
*SecretsApi* | [**ListSecretPaths**](client/docs/SecretsApi.md#listsecretpaths) | **Get** /secrets/{path}::listpaths | List Secret Paths
*SecretsApi* | [**RollbackSecret**](client/docs/SecretsApi.md#rollbacksecret) | **Put** /secrets/{path}/rollback/{version} | Rollback a Secret
*SecretsApi* | [**SearchSecrets**](client/docs/SecretsApi.md#searchsecrets) | **Get** /secrets | Search for Secrets
*SecretsApi* | [**UpdateSecret**](client/docs/SecretsApi.md#updatesecret) | **Put** /secrets/{path} | Update Secret
*SettingsApi* | [**CreateAuthSettings**](client/docs/SettingsApi.md#createauthsettings) | **Post** /config/auth/ | Create Authentication Provider
*SettingsApi* | [**DeleteAuthSettings**](client/docs/SettingsApi.md#deleteauthsettings) | **Delete** /config/auth/{name} | Delete Authentication Provider
*SettingsApi* | [**GetAuthSettings**](client/docs/SettingsApi.md#getauthsettings) | **Get** /config/auth/{name} | Get Authentication Provider
*SettingsApi* | [**GetAuthSettingsByVersion**](client/docs/SettingsApi.md#getauthsettingsbyversion) | **Get** /config/auth/{name}/version/{version} | Get a list of Authentication Settings by version
*SettingsApi* | [**RestoreAuthSettings**](client/docs/SettingsApi.md#restoreauthsettings) | **Get** /config/auth/{name}/restore | Restore Authentication Provider
*SettingsApi* | [**RollbackAuthSettings**](client/docs/SettingsApi.md#rollbackauthsettings) | **Put** /config/auth/{name}/rollback/{version} | Rollback Authentication Provider
*SettingsApi* | [**SearchSettings**](client/docs/SettingsApi.md#searchsettings) | **Get** /config/auth | Search Authentication Providers
*SettingsApi* | [**UpdateAuthSettings**](client/docs/SettingsApi.md#updateauthsettings) | **Put** /config/auth/{name} | Update Authentication Provider
*TasksApi* | [**GetTaskStatus**](client/docs/TasksApi.md#gettaskstatus) | **Get** /task/status/{id} | Get background task status
*TokensApi* | [**InitCertAuth**](client/docs/TokensApi.md#initcertauth) | **Post** /certificate/auth | Initiate authentication by certificate
*TokensApi* | [**Revoke**](client/docs/TokensApi.md#revoke) | **Post** /revoke/{refreshtoken} | Revoke Refresh Token
*TokensApi* | [**Token**](client/docs/TokensApi.md#token) | **Post** /token | Authenticate
*UsageApi* | [**GetUsage**](client/docs/UsageApi.md#getusage) | **Get** /usage | Get Usage
*UsersApi* | [**AddToGroups**](client/docs/UsersApi.md#addtogroups) | **Post** /users/{name}/groups | Add Member To Groups
*UsersApi* | [**ChangePassword**](client/docs/UsersApi.md#changepassword) | **Post** /users/{name}/password | Change Password
*UsersApi* | [**CreateUser**](client/docs/UsersApi.md#createuser) | **Post** /users/ | Create a User
*UsersApi* | [**DeleteUser**](client/docs/UsersApi.md#deleteuser) | **Delete** /users/{name} | Delete a User
*UsersApi* | [**GetMember**](client/docs/UsersApi.md#getmember) | **Get** /users/{name}/groups | Get Member Group
*UsersApi* | [**GetUser**](client/docs/UsersApi.md#getuser) | **Get** /users/{name} | Get a User
*UsersApi* | [**GetUserByVersion**](client/docs/UsersApi.md#getuserbyversion) | **Get** /users/{name}/version/{version} | Get a User By Version
*UsersApi* | [**RestoreUser**](client/docs/UsersApi.md#restoreuser) | **Get** /users/{name}/restore | Restore a User
*UsersApi* | [**SearchUsers**](client/docs/UsersApi.md#searchusers) | **Get** /users | Search for Users
*UsersApi* | [**UpdateUser**](client/docs/UsersApi.md#updateuser) | **Put** /users/{name} | Update a User

## Documentation For Models

 - [AccessTokenResponse](client/docs/AccessTokenResponse.md)
 - [AddMemberResponse](client/docs/AddMemberResponse.md)
 - [AddToGroupsRequest](client/docs/AddToGroupsRequest.md)
 - [AddToGroupsResponse](client/docs/AddToGroupsResponse.md)
 - [ApplyResponse](client/docs/ApplyResponse.md)
 - [Audit](client/docs/Audit.md)
 - [AuthProperties](client/docs/AuthProperties.md)
 - [Authentication](client/docs/Authentication.md)
 - [AuthenticationDetailsModel](client/docs/AuthenticationDetailsModel.md)
 - [AuthenticationProviderPropertiesModel](client/docs/AuthenticationProviderPropertiesModel.md)
 - [AuthenticationSettingsCreateModel](client/docs/AuthenticationSettingsCreateModel.md)
 - [AuthenticationSettingsResponse](client/docs/AuthenticationSettingsResponse.md)
 - [AuthenticationSettingsSearchResponse](client/docs/AuthenticationSettingsSearchResponse.md)
 - [AuthenticationSettingsVersionResponse](client/docs/AuthenticationSettingsVersionResponse.md)
 - [AutoKey](client/docs/AutoKey.md)
 - [AutoKeyResponse](client/docs/AutoKeyResponse.md)
 - [CertificateAuthBody](client/docs/CertificateAuthBody.md)
 - [ClientCreate](client/docs/ClientCreate.md)
 - [ClientCredentialsResponse](client/docs/ClientCredentialsResponse.md)
 - [ClientSearchModel](client/docs/ClientSearchModel.md)
 - [Condition](client/docs/Condition.md)
 - [ConfigResponse](client/docs/ConfigResponse.md)
 - [ConfigVersionResponse](client/docs/ConfigVersionResponse.md)
 - [CreateGroup](client/docs/CreateGroup.md)
 - [DecryptionResponse](client/docs/DecryptionResponse.md)
 - [DefaultPolicy](client/docs/DefaultPolicy.md)
 - [EncryptionResponse](client/docs/EncryptionResponse.md)
 - [EngineCreate](client/docs/EngineCreate.md)
 - [EngineCreateResponse](client/docs/EngineCreateResponse.md)
 - [EngineGetResponse](client/docs/EngineGetResponse.md)
 - [EngineListResult](client/docs/EngineListResult.md)
 - [EnginePingResponse](client/docs/EnginePingResponse.md)
 - [EngineSearchResponse](client/docs/EngineSearchResponse.md)
 - [GenerateResponse](client/docs/GenerateResponse.md)
 - [Group](client/docs/Group.md)
 - [GroupMemberInfo](client/docs/GroupMemberInfo.md)
 - [GroupResponse](client/docs/GroupResponse.md)
 - [GroupSearch](client/docs/GroupSearch.md)
 - [History](client/docs/History.md)
 - [HttpError](client/docs/HttpError.md)
 - [InitiateCertAuthResponse](client/docs/InitiateCertAuthResponse.md)
 - [Key](client/docs/Key.md)
 - [LogSearchResponse](client/docs/LogSearchResponse.md)
 - [ManualKeyData](client/docs/ManualKeyData.md)
 - [Masterkey](client/docs/Masterkey.md)
 - [MemberRequest](client/docs/MemberRequest.md)
 - [MemberResponse](client/docs/MemberResponse.md)
 - [MessageResponse](client/docs/MessageResponse.md)
 - [PageInfo](client/docs/PageInfo.md)
 - [PasswordChangeModel](client/docs/PasswordChangeModel.md)
 - [PolicyCreate](client/docs/PolicyCreate.md)
 - [PolicyResponse](client/docs/PolicyResponse.md)
 - [PolicySearchResponse](client/docs/PolicySearchResponse.md)
 - [PolicyUpdate](client/docs/PolicyUpdate.md)
 - [PolicyVersionResponse](client/docs/PolicyVersionResponse.md)
 - [Pool](client/docs/Pool.md)
 - [PoolCreate](client/docs/PoolCreate.md)
 - [PoolListResult](client/docs/PoolListResult.md)
 - [PostConfigModel](client/docs/PostConfigModel.md)
 - [ResponseCertificate](client/docs/ResponseCertificate.md)
 - [ResponseRootCertificate](client/docs/ResponseRootCertificate.md)
 - [ResponseSshCertificate](client/docs/ResponseSshCertificate.md)
 - [RoleCreate](client/docs/RoleCreate.md)
 - [RoleDetailsModel](client/docs/RoleDetailsModel.md)
 - [RoleResponse](client/docs/RoleResponse.md)
 - [RoleSearchResponse](client/docs/RoleSearchResponse.md)
 - [RoleVersionResponse](client/docs/RoleVersionResponse.md)
 - [RootCaRegistration](client/docs/RootCaRegistration.md)
 - [RootCaSecret](client/docs/RootCaSecret.md)
 - [SecretCreate](client/docs/SecretCreate.md)
 - [SecretDescription](client/docs/SecretDescription.md)
 - [SecretListPathsResponse](client/docs/SecretListPathsResponse.md)
 - [SecretResponse](client/docs/SecretResponse.md)
 - [SecretSearch](client/docs/SecretSearch.md)
 - [SecretUpdate](client/docs/SecretUpdate.md)
 - [SecretVersionResponse](client/docs/SecretVersionResponse.md)
 - [Settings](client/docs/Settings.md)
 - [SiemCreateUpdateRequestModel](client/docs/SiemCreateUpdateRequestModel.md)
 - [SiemNoSensitiveResponseModel](client/docs/SiemNoSensitiveResponseModel.md)
 - [SiemResponse](client/docs/SiemResponse.md)
 - [SiemSearchResponse](client/docs/SiemSearchResponse.md)
 - [SignedLeafCertificate](client/docs/SignedLeafCertificate.md)
 - [SigningRequest](client/docs/SigningRequest.md)
 - [SigningRequestInformation](client/docs/SigningRequestInformation.md)
 - [SshCertInformation](client/docs/SshCertInformation.md)
 - [StatusResponse](client/docs/StatusResponse.md)
 - [TaskResult](client/docs/TaskResult.md)
 - [TaskState](client/docs/TaskState.md)
 - [TokenBody](client/docs/TokenBody.md)
 - [UpdateKeyRequest](client/docs/UpdateKeyRequest.md)
 - [UsageResponseGeneral](client/docs/UsageResponseGeneral.md)
 - [UserCreateModel](client/docs/UserCreateModel.md)
 - [UserResponse](client/docs/UserResponse.md)
 - [UserSearch](client/docs/UserSearch.md)
 - [UserUpdateModel](client/docs/UserUpdateModel.md)
 - [UserVersionResponse](client/docs/UserVersionResponse.md)

## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```