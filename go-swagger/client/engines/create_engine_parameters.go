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

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// NewCreateEngineParams creates a new CreateEngineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEngineParams() *CreateEngineParams {
	return &CreateEngineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEngineParamsWithTimeout creates a new CreateEngineParams object
// with the ability to set a timeout on a request.
func NewCreateEngineParamsWithTimeout(timeout time.Duration) *CreateEngineParams {
	return &CreateEngineParams{
		timeout: timeout,
	}
}

// NewCreateEngineParamsWithContext creates a new CreateEngineParams object
// with the ability to set a context for a request.
func NewCreateEngineParamsWithContext(ctx context.Context) *CreateEngineParams {
	return &CreateEngineParams{
		Context: ctx,
	}
}

// NewCreateEngineParamsWithHTTPClient creates a new CreateEngineParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEngineParamsWithHTTPClient(client *http.Client) *CreateEngineParams {
	return &CreateEngineParams{
		HTTPClient: client,
	}
}

/* CreateEngineParams contains all the parameters to send to the API endpoint
   for the create engine operation.

   Typically these are written to a http.Request.
*/
type CreateEngineParams struct {

	// Body.
	Body *models.EngineCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEngineParams) WithDefaults() *CreateEngineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEngineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create engine params
func (o *CreateEngineParams) WithTimeout(timeout time.Duration) *CreateEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create engine params
func (o *CreateEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create engine params
func (o *CreateEngineParams) WithContext(ctx context.Context) *CreateEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create engine params
func (o *CreateEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create engine params
func (o *CreateEngineParams) WithHTTPClient(client *http.Client) *CreateEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create engine params
func (o *CreateEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create engine params
func (o *CreateEngineParams) WithBody(body *models.EngineCreate) *CreateEngineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create engine params
func (o *CreateEngineParams) SetBody(body *models.EngineCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
