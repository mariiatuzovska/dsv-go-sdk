# {{classname}}

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](GroupsApi.md#AddMember) | **Post** /groups/{name}/members | Add Members To Group
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /groups/ | Create Group
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /groups/{name} | Delete Group
[**DeleteMember**](GroupsApi.md#DeleteMember) | **Delete** /groups/{name}/members | Delete Members From Group
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /groups/{name} | Get Group
[**RestoreGroup**](GroupsApi.md#RestoreGroup) | **Get** /groups/{name}/restore | Restore Group
[**SearchGroups**](GroupsApi.md#SearchGroups) | **Get** /groups | Search Groups

# **AddMember**
> DsvAddMemberResponse AddMember(ctx, body, name)
Add Members To Group

Add one or more members to a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvMemberRequest**](DsvMemberRequest.md)|  | 
  **name** | **string**| Group name | 

### Return type

[**DsvAddMemberResponse**](AddMemberResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> DsvAddMemberResponse CreateGroup(ctx, body)
Create Group

Create a new group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvCreateGroup**](DsvCreateGroup.md)|  | 

### Return type

[**DsvAddMemberResponse**](AddMemberResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DsvMessageResponse DeleteGroup(ctx, name, optional)
Delete Group

Delete Group if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Group name | 
 **optional** | ***GroupsApiDeleteGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiDeleteGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Delete immediately | 

### Return type

[**DsvMessageResponse**](MessageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMember**
> DeleteMember(ctx, body, name)
Delete Members From Group

Delete one or more members from a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvMemberRequest**](DsvMemberRequest.md)|  | 
  **name** | **string**| Group name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> DsvGroupResponse GetGroup(ctx, name)
Get Group

Retrieve an existing group by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Group name | 

### Return type

[**DsvGroupResponse**](GroupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreGroup**
> RestoreGroup(ctx, name)
Restore Group

Restore a soft-deleted group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Group name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGroups**
> DsvGroupSearch SearchGroups(ctx, optional)
Search Groups

Search for one or more groups by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiSearchGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiSearchGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| Partial search term for search by path | 
 **limit** | **optional.Int64**| Limit for the number of results per page (cursor) | 
 **cursor** | **optional.String**| Cursor to next batch of results | 
 **sort** | **optional.String**| Sort results ascending (asc) or descending (desc) order by lastModified attribute on field search. Default is desc | 
 **sortedBy** | **optional.String**| SortedBy order the result by name, created or lastModified attribute on field search. Default is lastModified | 

### Return type

[**DsvGroupSearch**](GroupSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

