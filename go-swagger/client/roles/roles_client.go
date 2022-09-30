// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleCreated, error)

	DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleOK, *DeleteRoleNoContent, error)

	GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error)

	GetRoleByVersion(params *GetRoleByVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleByVersionOK, error)

	RestoreRole(params *RestoreRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreRoleOK, error)

	SearchRoles(params *SearchRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchRolesOK, error)

	UpdateRole(params *UpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRole creates a role

  Creates a new role.
*/
func (a *Client) CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRole",
		Method:             "POST",
		PathPattern:        "/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRoleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRole deletes a role

  Delete a role by the role name. For roles linked to 3rd party providers, such as AWS or Azure, the role name
must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db
*/
func (a *Client) DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleOK, *DeleteRoleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRole",
		Method:             "DELETE",
		PathPattern:        "/roles/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteRoleOK:
		return value, nil, nil
	case *DeleteRoleNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRole gets a role

  Retrieve an existing role by role name. For roles linked to 3rd party providers, such as AWS or Azure, the role name
must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db
*/
func (a *Client) GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRole",
		Method:             "GET",
		PathPattern:        "/roles/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRoleByVersion gets a role by version

  Retrieve an existing role by role name and versions. For roles linked to 3rd party providers, such as AWS or Azure, the role name
must be prefixed with the provider name from configuration in the format of <providername>:<rolename> i.e. aws-dev:db
*/
func (a *Client) GetRoleByVersion(params *GetRoleByVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRoleByVersion",
		Method:             "GET",
		PathPattern:        "/roles/{name}/version/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRoleByVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRoleByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRoleByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestoreRole restores a role

  Restore a role by path.
*/
func (a *Client) RestoreRole(params *RestoreRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreRole",
		Method:             "GET",
		PathPattern:        "/roles/{name}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchRoles searches for roles

  Search for one or more roles by role name.
*/
func (a *Client) SearchRoles(params *SearchRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchRoles",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRole updates a role

  Update an existing role.
*/
func (a *Client) UpdateRole(params *UpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRole",
		Method:             "PUT",
		PathPattern:        "/roles/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
