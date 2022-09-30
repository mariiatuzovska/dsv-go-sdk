// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewRollbackAuthSettingsParams creates a new RollbackAuthSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRollbackAuthSettingsParams() *RollbackAuthSettingsParams {
	return &RollbackAuthSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackAuthSettingsParamsWithTimeout creates a new RollbackAuthSettingsParams object
// with the ability to set a timeout on a request.
func NewRollbackAuthSettingsParamsWithTimeout(timeout time.Duration) *RollbackAuthSettingsParams {
	return &RollbackAuthSettingsParams{
		timeout: timeout,
	}
}

// NewRollbackAuthSettingsParamsWithContext creates a new RollbackAuthSettingsParams object
// with the ability to set a context for a request.
func NewRollbackAuthSettingsParamsWithContext(ctx context.Context) *RollbackAuthSettingsParams {
	return &RollbackAuthSettingsParams{
		Context: ctx,
	}
}

// NewRollbackAuthSettingsParamsWithHTTPClient creates a new RollbackAuthSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRollbackAuthSettingsParamsWithHTTPClient(client *http.Client) *RollbackAuthSettingsParams {
	return &RollbackAuthSettingsParams{
		HTTPClient: client,
	}
}

/* RollbackAuthSettingsParams contains all the parameters to send to the API endpoint
   for the rollback auth settings operation.

   Typically these are written to a http.Request.
*/
type RollbackAuthSettingsParams struct {

	/* Name.

	   Full name to lookup authentication settings by.
	*/
	Name string

	/* Version.

	   Versions to return.

	   Format: int64
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rollback auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackAuthSettingsParams) WithDefaults() *RollbackAuthSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rollback auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackAuthSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rollback auth settings params
func (o *RollbackAuthSettingsParams) WithTimeout(timeout time.Duration) *RollbackAuthSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback auth settings params
func (o *RollbackAuthSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback auth settings params
func (o *RollbackAuthSettingsParams) WithContext(ctx context.Context) *RollbackAuthSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback auth settings params
func (o *RollbackAuthSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback auth settings params
func (o *RollbackAuthSettingsParams) WithHTTPClient(client *http.Client) *RollbackAuthSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback auth settings params
func (o *RollbackAuthSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the rollback auth settings params
func (o *RollbackAuthSettingsParams) WithName(name string) *RollbackAuthSettingsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the rollback auth settings params
func (o *RollbackAuthSettingsParams) SetName(name string) {
	o.Name = name
}

// WithVersion adds the version to the rollback auth settings params
func (o *RollbackAuthSettingsParams) WithVersion(version int64) *RollbackAuthSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the rollback auth settings params
func (o *RollbackAuthSettingsParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackAuthSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
