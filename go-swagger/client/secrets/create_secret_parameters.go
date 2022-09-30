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

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// NewCreateSecretParams creates a new CreateSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSecretParams() *CreateSecretParams {
	return &CreateSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecretParamsWithTimeout creates a new CreateSecretParams object
// with the ability to set a timeout on a request.
func NewCreateSecretParamsWithTimeout(timeout time.Duration) *CreateSecretParams {
	return &CreateSecretParams{
		timeout: timeout,
	}
}

// NewCreateSecretParamsWithContext creates a new CreateSecretParams object
// with the ability to set a context for a request.
func NewCreateSecretParamsWithContext(ctx context.Context) *CreateSecretParams {
	return &CreateSecretParams{
		Context: ctx,
	}
}

// NewCreateSecretParamsWithHTTPClient creates a new CreateSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSecretParamsWithHTTPClient(client *http.Client) *CreateSecretParams {
	return &CreateSecretParams{
		HTTPClient: client,
	}
}

/* CreateSecretParams contains all the parameters to send to the API endpoint
   for the create secret operation.

   Typically these are written to a http.Request.
*/
type CreateSecretParams struct {

	// Body.
	Body *models.RequestModelCreate

	/* Path.

	   The full secret path i.e. servers/prod/webserver-01
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecretParams) WithDefaults() *CreateSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create secret params
func (o *CreateSecretParams) WithTimeout(timeout time.Duration) *CreateSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create secret params
func (o *CreateSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create secret params
func (o *CreateSecretParams) WithContext(ctx context.Context) *CreateSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create secret params
func (o *CreateSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create secret params
func (o *CreateSecretParams) WithHTTPClient(client *http.Client) *CreateSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create secret params
func (o *CreateSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create secret params
func (o *CreateSecretParams) WithBody(body *models.RequestModelCreate) *CreateSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create secret params
func (o *CreateSecretParams) SetBody(body *models.RequestModelCreate) {
	o.Body = body
}

// WithPath adds the path to the create secret params
func (o *CreateSecretParams) WithPath(path string) *CreateSecretParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the create secret params
func (o *CreateSecretParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
