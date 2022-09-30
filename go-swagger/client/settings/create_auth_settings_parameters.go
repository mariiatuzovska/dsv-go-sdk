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

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// NewCreateAuthSettingsParams creates a new CreateAuthSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAuthSettingsParams() *CreateAuthSettingsParams {
	return &CreateAuthSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthSettingsParamsWithTimeout creates a new CreateAuthSettingsParams object
// with the ability to set a timeout on a request.
func NewCreateAuthSettingsParamsWithTimeout(timeout time.Duration) *CreateAuthSettingsParams {
	return &CreateAuthSettingsParams{
		timeout: timeout,
	}
}

// NewCreateAuthSettingsParamsWithContext creates a new CreateAuthSettingsParams object
// with the ability to set a context for a request.
func NewCreateAuthSettingsParamsWithContext(ctx context.Context) *CreateAuthSettingsParams {
	return &CreateAuthSettingsParams{
		Context: ctx,
	}
}

// NewCreateAuthSettingsParamsWithHTTPClient creates a new CreateAuthSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAuthSettingsParamsWithHTTPClient(client *http.Client) *CreateAuthSettingsParams {
	return &CreateAuthSettingsParams{
		HTTPClient: client,
	}
}

/* CreateAuthSettingsParams contains all the parameters to send to the API endpoint
   for the create auth settings operation.

   Typically these are written to a http.Request.
*/
type CreateAuthSettingsParams struct {

	// Body.
	Body *models.AuthenticationCreateModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthSettingsParams) WithDefaults() *CreateAuthSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create auth settings params
func (o *CreateAuthSettingsParams) WithTimeout(timeout time.Duration) *CreateAuthSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create auth settings params
func (o *CreateAuthSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create auth settings params
func (o *CreateAuthSettingsParams) WithContext(ctx context.Context) *CreateAuthSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create auth settings params
func (o *CreateAuthSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create auth settings params
func (o *CreateAuthSettingsParams) WithHTTPClient(client *http.Client) *CreateAuthSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create auth settings params
func (o *CreateAuthSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create auth settings params
func (o *CreateAuthSettingsParams) WithBody(body *models.AuthenticationCreateModel) *CreateAuthSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create auth settings params
func (o *CreateAuthSettingsParams) SetBody(body *models.AuthenticationCreateModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
