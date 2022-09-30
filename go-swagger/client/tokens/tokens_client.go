// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InitCertAuth(params *InitCertAuthParams, opts ...ClientOption) (*InitCertAuthOK, error)

	Revoke(params *RevokeParams, opts ...ClientOption) (*RevokeNoContent, error)

	Token(params *TokenParams, opts ...ClientOption) (*TokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InitCertAuth initiates authentication by certificate

  Request a challenge to decrypt to prove ownership of the private key for
authentication by certificate flow. Challenge id returned in response is
only valid for 5 minutes.
*/
func (a *Client) InitCertAuth(params *InitCertAuthParams, opts ...ClientOption) (*InitCertAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitCertAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "init-cert-auth",
		Method:             "POST",
		PathPattern:        "/certificate/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InitCertAuthReader{formats: a.formats},
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
	success, ok := result.(*InitCertAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for init-cert-auth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Revoke revokes refresh token

  Revoke an existing refresh token to prevent it from being used for authentication.
*/
func (a *Client) Revoke(params *RevokeParams, opts ...ClientOption) (*RevokeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "revoke",
		Method:             "POST",
		PathPattern:        "/revoke/{refreshtoken}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeReader{formats: a.formats},
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
	success, ok := result.(*RevokeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revoke: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Token authenticates

  Submit parameters to get a new access token for use on protected endpoints
*/
func (a *Client) Token(params *TokenParams, opts ...ClientOption) (*TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "token",
		Method:             "POST",
		PathPattern:        "/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TokenReader{formats: a.formats},
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
	success, ok := result.(*TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
