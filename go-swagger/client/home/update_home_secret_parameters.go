// Code generated by go-swagger; DO NOT EDIT.

package home

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

// NewUpdateHomeSecretParams creates a new UpdateHomeSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHomeSecretParams() *UpdateHomeSecretParams {
	return &UpdateHomeSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHomeSecretParamsWithTimeout creates a new UpdateHomeSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateHomeSecretParamsWithTimeout(timeout time.Duration) *UpdateHomeSecretParams {
	return &UpdateHomeSecretParams{
		timeout: timeout,
	}
}

// NewUpdateHomeSecretParamsWithContext creates a new UpdateHomeSecretParams object
// with the ability to set a context for a request.
func NewUpdateHomeSecretParamsWithContext(ctx context.Context) *UpdateHomeSecretParams {
	return &UpdateHomeSecretParams{
		Context: ctx,
	}
}

// NewUpdateHomeSecretParamsWithHTTPClient creates a new UpdateHomeSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHomeSecretParamsWithHTTPClient(client *http.Client) *UpdateHomeSecretParams {
	return &UpdateHomeSecretParams{
		HTTPClient: client,
	}
}

/* UpdateHomeSecretParams contains all the parameters to send to the API endpoint
   for the update home secret operation.

   Typically these are written to a http.Request.
*/
type UpdateHomeSecretParams struct {

	// Body.
	Body *models.RequestModelUpdate

	/* Path.

	   The full secret path i.e. servers/prod/webserver-01
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update home secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomeSecretParams) WithDefaults() *UpdateHomeSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update home secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomeSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update home secret params
func (o *UpdateHomeSecretParams) WithTimeout(timeout time.Duration) *UpdateHomeSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update home secret params
func (o *UpdateHomeSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update home secret params
func (o *UpdateHomeSecretParams) WithContext(ctx context.Context) *UpdateHomeSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update home secret params
func (o *UpdateHomeSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update home secret params
func (o *UpdateHomeSecretParams) WithHTTPClient(client *http.Client) *UpdateHomeSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update home secret params
func (o *UpdateHomeSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update home secret params
func (o *UpdateHomeSecretParams) WithBody(body *models.RequestModelUpdate) *UpdateHomeSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update home secret params
func (o *UpdateHomeSecretParams) SetBody(body *models.RequestModelUpdate) {
	o.Body = body
}

// WithPath adds the path to the update home secret params
func (o *UpdateHomeSecretParams) WithPath(path string) *UpdateHomeSecretParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the update home secret params
func (o *UpdateHomeSecretParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHomeSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
