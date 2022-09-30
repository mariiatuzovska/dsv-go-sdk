// Code generated by go-swagger; DO NOT EDIT.

package s_i_e_m

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

// NewSiemGetParams creates a new SiemGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSiemGetParams() *SiemGetParams {
	return &SiemGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSiemGetParamsWithTimeout creates a new SiemGetParams object
// with the ability to set a timeout on a request.
func NewSiemGetParamsWithTimeout(timeout time.Duration) *SiemGetParams {
	return &SiemGetParams{
		timeout: timeout,
	}
}

// NewSiemGetParamsWithContext creates a new SiemGetParams object
// with the ability to set a context for a request.
func NewSiemGetParamsWithContext(ctx context.Context) *SiemGetParams {
	return &SiemGetParams{
		Context: ctx,
	}
}

// NewSiemGetParamsWithHTTPClient creates a new SiemGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSiemGetParamsWithHTTPClient(client *http.Client) *SiemGetParams {
	return &SiemGetParams{
		HTTPClient: client,
	}
}

/* SiemGetParams contains all the parameters to send to the API endpoint
   for the siem get operation.

   Typically these are written to a http.Request.
*/
type SiemGetParams struct {

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the siem get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiemGetParams) WithDefaults() *SiemGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the siem get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiemGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the siem get params
func (o *SiemGetParams) WithTimeout(timeout time.Duration) *SiemGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the siem get params
func (o *SiemGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the siem get params
func (o *SiemGetParams) WithContext(ctx context.Context) *SiemGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the siem get params
func (o *SiemGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the siem get params
func (o *SiemGetParams) WithHTTPClient(client *http.Client) *SiemGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the siem get params
func (o *SiemGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the siem get params
func (o *SiemGetParams) WithName(name string) *SiemGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the siem get params
func (o *SiemGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SiemGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
