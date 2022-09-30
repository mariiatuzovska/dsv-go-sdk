// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_auto

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
)

// NewRestoreKeyParams creates a new RestoreKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreKeyParams() *RestoreKeyParams {
	return &RestoreKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreKeyParamsWithTimeout creates a new RestoreKeyParams object
// with the ability to set a timeout on a request.
func NewRestoreKeyParamsWithTimeout(timeout time.Duration) *RestoreKeyParams {
	return &RestoreKeyParams{
		timeout: timeout,
	}
}

// NewRestoreKeyParamsWithContext creates a new RestoreKeyParams object
// with the ability to set a context for a request.
func NewRestoreKeyParamsWithContext(ctx context.Context) *RestoreKeyParams {
	return &RestoreKeyParams{
		Context: ctx,
	}
}

// NewRestoreKeyParamsWithHTTPClient creates a new RestoreKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreKeyParamsWithHTTPClient(client *http.Client) *RestoreKeyParams {
	return &RestoreKeyParams{
		HTTPClient: client,
	}
}

/* RestoreKeyParams contains all the parameters to send to the API endpoint
   for the restore key operation.

   Typically these are written to a http.Request.
*/
type RestoreKeyParams struct {

	/* Path.

	   The full key path, for example, mykeys/key01
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreKeyParams) WithDefaults() *RestoreKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore key params
func (o *RestoreKeyParams) WithTimeout(timeout time.Duration) *RestoreKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore key params
func (o *RestoreKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore key params
func (o *RestoreKeyParams) WithContext(ctx context.Context) *RestoreKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore key params
func (o *RestoreKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore key params
func (o *RestoreKeyParams) WithHTTPClient(client *http.Client) *RestoreKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore key params
func (o *RestoreKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the restore key params
func (o *RestoreKeyParams) WithPath(path string) *RestoreKeyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the restore key params
func (o *RestoreKeyParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
