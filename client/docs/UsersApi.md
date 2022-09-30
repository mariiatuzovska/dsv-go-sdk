# \UsersApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToGroups**](UsersApi.md#AddToGroups) | **Post** /users/{name}/groups | Add Member To Groups
[**ChangePassword**](UsersApi.md#ChangePassword) | **Post** /users/{name}/password | Change Password
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users/ | Create a User
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{name} | Delete a User
[**GetMember**](UsersApi.md#GetMember) | **Get** /users/{name}/groups | Get Member Group
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{name} | Get a User
[**GetUserByVersion**](UsersApi.md#GetUserByVersion) | **Get** /users/{name}/version/{version} | Get a User By Version
[**RestoreUser**](UsersApi.md#RestoreUser) | **Get** /users/{name}/restore | Restore a User
[**SearchUsers**](UsersApi.md#SearchUsers) | **Get** /users | Search for Users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users/{name} | Update a User


# **AddToGroups**
> DsvAddToGroupsResponse AddToGroups(ctx, name, body)
Add Member To Groups

Add a user to one or more groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 
  **body** | [**DsvAddToGroupsRequest**](DsvAddToGroupsRequest.md)|  | 

### Return type

[**DsvAddToGroupsResponse**](AddToGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePassword**
> DsvMessageResponse ChangePassword(ctx, name, body)
Change Password

Allows the user to change their own password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 
  **body** | [**DsvPasswordChangeModel**](DsvPasswordChangeModel.md)|  | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> DsvUserResponse CreateUser(ctx, body)
Create a User

Create a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvUserCreateModel**](DsvUserCreateModel.md)|  | 

### Return type

[**DsvUserResponse**](UserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DsvMessageResponse DeleteUser(ctx, name, optional)
Delete a User

Retrieve an existing user by user name. For users linked to 3rd party providers, such as AWS or Azure, the user name must be prefixed with the provider name from configuration in the format of <providername>:<username> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 
 **optional** | ***UsersApiDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiDeleteUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMember**
> DsvMemberResponse GetMember(ctx, name)
Get Member Group

Retrieve a member's group by member name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Member name to lookup member | 

### Return type

[**DsvMemberResponse**](MemberResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> DsvUserResponse GetUser(ctx, name)
Get a User

Retrieve an existing user by user name. For users linked to 3rd party providers, such as AWS or Azure, the user name must be prefixed with the provider name from configuration in the format of <providername>:<username> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 

### Return type

[**DsvUserResponse**](UserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserByVersion**
> DsvUserVersionResponse GetUserByVersion(ctx, name, version)
Get a User By Version

Retrieve an existing user by user name and version. For users linked to 3rd party providers, such as AWS or Azure, the user name must be prefixed with the provider name from configuration in the format of <providername>:<username> i.e. aws-dev:db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 
  **version** | **int64**| Versions to return | 

### Return type

[**DsvUserVersionResponse**](UserVersionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreUser**
> RestoreUser(ctx, name)
Restore a User

Restore a user by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsers**
> DsvUserSearch SearchUsers(ctx, optional)
Search for Users

Search for one or more users by their name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiSearchUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchField** | **optional.String**| User field for advanced searching | 
 **searchComparison** | **optional.String**| Comparison type (equal, contains, begins_with) for advanced searching | 
 **searchType** | **optional.String**| Attribute type (string, number) to search on | 
 **sort** | **optional.String**| Sort results ascending (asc) or descending (desc) order by lastModified attribute on field search. Default is desc | 
 **sortedBy** | **optional.String**| SortedBy order the result by name, created or lastModified attribute on field search. Default is lastModified | 
 **searchTerm** | **optional.String**| Search pattern for names of users to look up | 
 **limit** | **optional.Int64**| The maximum number of results per cursor | 
 **cursor** | **optional.String**| Cursor to next batch of results | 

### Return type

[**DsvUserSearch**](UserSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> DsvUserResponse UpdateUser(ctx, name, body)
Update a User

Update an existing user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full user name to lookup user by | 
  **body** | [**DsvUserUpdateModel**](DsvUserUpdateModel.md)|  | 

### Return type

[**DsvUserResponse**](UserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

