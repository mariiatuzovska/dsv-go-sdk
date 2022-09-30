// Code generated by go-swagger; DO NOT EDIT.

package p_k_i

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

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// NewRegisterRootParams creates a new RegisterRootParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterRootParams() *RegisterRootParams {
	return &RegisterRootParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterRootParamsWithTimeout creates a new RegisterRootParams object
// with the ability to set a timeout on a request.
func NewRegisterRootParamsWithTimeout(timeout time.Duration) *RegisterRootParams {
	return &RegisterRootParams{
		timeout: timeout,
	}
}

// NewRegisterRootParamsWithContext creates a new RegisterRootParams object
// with the ability to set a context for a request.
func NewRegisterRootParamsWithContext(ctx context.Context) *RegisterRootParams {
	return &RegisterRootParams{
		Context: ctx,
	}
}

// NewRegisterRootParamsWithHTTPClient creates a new RegisterRootParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterRootParamsWithHTTPClient(client *http.Client) *RegisterRootParams {
	return &RegisterRootParams{
		HTTPClient: client,
	}
}

/* RegisterRootParams contains all the parameters to send to the API endpoint
   for the register root operation.

   Typically these are written to a http.Request.
*/
type RegisterRootParams struct {

	// Body.
	Body *models.RootCASecret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterRootParams) WithDefaults() *RegisterRootParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterRootParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register root params
func (o *RegisterRootParams) WithTimeout(timeout time.Duration) *RegisterRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register root params
func (o *RegisterRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register root params
func (o *RegisterRootParams) WithContext(ctx context.Context) *RegisterRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register root params
func (o *RegisterRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register root params
func (o *RegisterRootParams) WithHTTPClient(client *http.Client) *RegisterRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register root params
func (o *RegisterRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register root params
func (o *RegisterRootParams) WithBody(body *models.RootCASecret) *RegisterRootParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register root params
func (o *RegisterRootParams) SetBody(body *models.RootCASecret) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
