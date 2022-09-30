// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new clients API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clients API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateClientCredential(params *CreateClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClientCredentialCreated, error)

	DeleteClientCredential(params *DeleteClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClientCredentialOK, error)

	GetBootstrapClientCredential(params *GetBootstrapClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBootstrapClientCredentialOK, error)

	GetClientCredential(params *GetClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientCredentialOK, error)

	RestoreClient(params *RestoreClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreClientOK, error)

	SearchClients(params *SearchClientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchClientsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateClientCredential creates a client credential

  Request a new client credential for a role and get back the client id and secret key.
*/
func (a *Client) CreateClientCredential(params *CreateClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClientCredentialCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClientCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createClientCredential",
		Method:             "POST",
		PathPattern:        "/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClientCredentialReader{formats: a.formats},
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
	success, ok := result.(*CreateClientCredentialCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createClientCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteClientCredential deletes a client credential

  Delete a client credential by its unique client id.
*/
func (a *Client) DeleteClientCredential(params *DeleteClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClientCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClientCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteClientCredential",
		Method:             "DELETE",
		PathPattern:        "/clients/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClientCredentialReader{formats: a.formats},
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
	success, ok := result.(*DeleteClientCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteClientCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBootstrapClientCredential gets a client bootstrap credential including secret

  Get a client credential by url.
*/
func (a *Client) GetBootstrapClientCredential(params *GetBootstrapClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBootstrapClientCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBootstrapClientCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBootstrapClientCredential",
		Method:             "GET",
		PathPattern:        "/clients/bootstrap/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBootstrapClientCredentialReader{formats: a.formats},
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
	success, ok := result.(*GetBootstrapClientCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBootstrapClientCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClientCredential gets a client credential

  Get a client credential by client id. The secret key is omitted.
*/
func (a *Client) GetClientCredential(params *GetClientCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClientCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClientCredential",
		Method:             "GET",
		PathPattern:        "/clients/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClientCredentialReader{formats: a.formats},
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
	success, ok := result.(*GetClientCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClientCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestoreClient restores a client

  Restore a client by ID.
*/
func (a *Client) RestoreClient(params *RestoreClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreClient",
		Method:             "GET",
		PathPattern:        "/clients/{clientId}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreClientReader{formats: a.formats},
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
	success, ok := result.(*RestoreClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchClients searches for client credentials

  Search for one or more client credentials associated with a particular role.
*/
func (a *Client) SearchClients(params *SearchClientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchClientsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchClientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchClients",
		Method:             "GET",
		PathPattern:        "/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchClientsReader{formats: a.formats},
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
	success, ok := result.(*SearchClientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchClients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
