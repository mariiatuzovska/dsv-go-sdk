// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewRestorePolicyParams creates a new RestorePolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestorePolicyParams() *RestorePolicyParams {
	return &RestorePolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestorePolicyParamsWithTimeout creates a new RestorePolicyParams object
// with the ability to set a timeout on a request.
func NewRestorePolicyParamsWithTimeout(timeout time.Duration) *RestorePolicyParams {
	return &RestorePolicyParams{
		timeout: timeout,
	}
}

// NewRestorePolicyParamsWithContext creates a new RestorePolicyParams object
// with the ability to set a context for a request.
func NewRestorePolicyParamsWithContext(ctx context.Context) *RestorePolicyParams {
	return &RestorePolicyParams{
		Context: ctx,
	}
}

// NewRestorePolicyParamsWithHTTPClient creates a new RestorePolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestorePolicyParamsWithHTTPClient(client *http.Client) *RestorePolicyParams {
	return &RestorePolicyParams{
		HTTPClient: client,
	}
}

/* RestorePolicyParams contains all the parameters to send to the API endpoint
   for the restore policy operation.

   Typically these are written to a http.Request.
*/
type RestorePolicyParams struct {

	/* Path.

	   Full path to lookup policy
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestorePolicyParams) WithDefaults() *RestorePolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestorePolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore policy params
func (o *RestorePolicyParams) WithTimeout(timeout time.Duration) *RestorePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore policy params
func (o *RestorePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore policy params
func (o *RestorePolicyParams) WithContext(ctx context.Context) *RestorePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore policy params
func (o *RestorePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore policy params
func (o *RestorePolicyParams) WithHTTPClient(client *http.Client) *RestorePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore policy params
func (o *RestorePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the restore policy params
func (o *RestorePolicyParams) WithPath(path string) *RestorePolicyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the restore policy params
func (o *RestorePolicyParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RestorePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
