# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /roles/ | Create a Role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /roles/{name} | Delete a Role
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{name} | Get a Role
[**GetRoleByVersion**](RolesApi.md#GetRoleByVersion) | **Get** /roles/{name}/version/{version} | Get a Role By Version
[**RestoreRole**](RolesApi.md#RestoreRole) | **Get** /roles/{name}/restore | Restore a Role
[**SearchRoles**](RolesApi.md#SearchRoles) | **Get** /roles | Search for Roles
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /roles/{name} | Update a Role

# **CreateRole**
> RoleResponse CreateRole(ctx, body)
Create a Role

Creates a new role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleCreate**](RoleCreate.md)|  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> MessageResponse DeleteRole(ctx, name, optional)
Delete a Role

Delete a role by the role name. For roles linked to 3rd party providers, such as AWS or Azure, the role name must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full role name to lookup role by | 
 **optional** | ***RolesApiDeleteRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiDeleteRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> RoleResponse GetRole(ctx, name)
Get a Role

Retrieve an existing role by role name. For roles linked to 3rd party providers, such as AWS or Azure, the role name must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full role name to lookup role by | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleByVersion**
> RoleVersionResponse GetRoleByVersion(ctx, name, version)
Get a Role By Version

Retrieve an existing role by role name and versions. For roles linked to 3rd party providers, such as AWS or Azure, the role name must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full name to lookup role by | 
  **version** | **int64**| Versions to return | 

### Return type

[**RoleVersionResponse**](RoleVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreRole**
> RestoreRole(ctx, name)
Restore a Role

Restore a role by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full role name to lookup role by | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchRoles**
> RoleSearchResponse SearchRoles(ctx, optional)
Search for Roles

Search for one or more roles by role name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesApiSearchRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiSearchRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| Search pattern for names of roles to look up | 
 **limit** | **optional.Int64**| The maximum number of results per cursor | 
 **cursor** | **optional.String**| C ursor to next batch of results | 

### Return type

[**RoleSearchResponse**](RoleSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> RoleResponse UpdateRole(ctx, body, name)
Update a Role

Update an existing role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleDetailsModel**](RoleDetailsModel.md)|  | 
  **name** | **string**| Full role name to lookup role by | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

