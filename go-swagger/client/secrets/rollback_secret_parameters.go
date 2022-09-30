// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRollbackSecretParams creates a new RollbackSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRollbackSecretParams() *RollbackSecretParams {
	return &RollbackSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackSecretParamsWithTimeout creates a new RollbackSecretParams object
// with the ability to set a timeout on a request.
func NewRollbackSecretParamsWithTimeout(timeout time.Duration) *RollbackSecretParams {
	return &RollbackSecretParams{
		timeout: timeout,
	}
}

// NewRollbackSecretParamsWithContext creates a new RollbackSecretParams object
// with the ability to set a context for a request.
func NewRollbackSecretParamsWithContext(ctx context.Context) *RollbackSecretParams {
	return &RollbackSecretParams{
		Context: ctx,
	}
}

// NewRollbackSecretParamsWithHTTPClient creates a new RollbackSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewRollbackSecretParamsWithHTTPClient(client *http.Client) *RollbackSecretParams {
	return &RollbackSecretParams{
		HTTPClient: client,
	}
}

/* RollbackSecretParams contains all the parameters to send to the API endpoint
   for the rollback secret operation.

   Typically these are written to a http.Request.
*/
type RollbackSecretParams struct {

	/* Path.

	   The full secret path i.e. servers/prod/webserver-01
	*/
	Path string

	/* Version.

	   Versions to return

	   Format: int64
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rollback secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackSecretParams) WithDefaults() *RollbackSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rollback secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rollback secret params
func (o *RollbackSecretParams) WithTimeout(timeout time.Duration) *RollbackSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback secret params
func (o *RollbackSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback secret params
func (o *RollbackSecretParams) WithContext(ctx context.Context) *RollbackSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback secret params
func (o *RollbackSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback secret params
func (o *RollbackSecretParams) WithHTTPClient(client *http.Client) *RollbackSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback secret params
func (o *RollbackSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the rollback secret params
func (o *RollbackSecretParams) WithPath(path string) *RollbackSecretParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the rollback secret params
func (o *RollbackSecretParams) SetPath(path string) {
	o.Path = path
}

// WithVersion adds the version to the rollback secret params
func (o *RollbackSecretParams) WithVersion(version int64) *RollbackSecretParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the rollback secret params
func (o *RollbackSecretParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", swag.FormatInt64(o.Version)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
