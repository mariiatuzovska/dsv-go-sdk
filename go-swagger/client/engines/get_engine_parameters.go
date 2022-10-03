// Code generated by go-swagger; DO NOT EDIT.

package engines

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

// NewGetEngineParams creates a new GetEngineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEngineParams() *GetEngineParams {
	return &GetEngineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEngineParamsWithTimeout creates a new GetEngineParams object
// with the ability to set a timeout on a request.
func NewGetEngineParamsWithTimeout(timeout time.Duration) *GetEngineParams {
	return &GetEngineParams{
		timeout: timeout,
	}
}

// NewGetEngineParamsWithContext creates a new GetEngineParams object
// with the ability to set a context for a request.
func NewGetEngineParamsWithContext(ctx context.Context) *GetEngineParams {
	return &GetEngineParams{
		Context: ctx,
	}
}

// NewGetEngineParamsWithHTTPClient creates a new GetEngineParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEngineParamsWithHTTPClient(client *http.Client) *GetEngineParams {
	return &GetEngineParams{
		HTTPClient: client,
	}
}

/* GetEngineParams contains all the parameters to send to the API endpoint
   for the get engine operation.

   Typically these are written to a http.Request.
*/
type GetEngineParams struct {

	/* Name.

	   Name of engine
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEngineParams) WithDefaults() *GetEngineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEngineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get engine params
func (o *GetEngineParams) WithTimeout(timeout time.Duration) *GetEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get engine params
func (o *GetEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get engine params
func (o *GetEngineParams) WithContext(ctx context.Context) *GetEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get engine params
func (o *GetEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get engine params
func (o *GetEngineParams) WithHTTPClient(client *http.Client) *GetEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get engine params
func (o *GetEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get engine params
func (o *GetEngineParams) WithName(name string) *GetEngineParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get engine params
func (o *GetEngineParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
