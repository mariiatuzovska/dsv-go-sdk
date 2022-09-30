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

// NewLeafParamsParams creates a new LeafParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLeafParamsParams() *LeafParamsParams {
	return &LeafParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLeafParamsParamsWithTimeout creates a new LeafParamsParams object
// with the ability to set a timeout on a request.
func NewLeafParamsParamsWithTimeout(timeout time.Duration) *LeafParamsParams {
	return &LeafParamsParams{
		timeout: timeout,
	}
}

// NewLeafParamsParamsWithContext creates a new LeafParamsParams object
// with the ability to set a context for a request.
func NewLeafParamsParamsWithContext(ctx context.Context) *LeafParamsParams {
	return &LeafParamsParams{
		Context: ctx,
	}
}

// NewLeafParamsParamsWithHTTPClient creates a new LeafParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLeafParamsParamsWithHTTPClient(client *http.Client) *LeafParamsParams {
	return &LeafParamsParams{
		HTTPClient: client,
	}
}

/* LeafParamsParams contains all the parameters to send to the API endpoint
   for the leaf params operation.

   Typically these are written to a http.Request.
*/
type LeafParamsParams struct {

	// Body.
	Body *models.SigningRequestInformation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the leaf params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LeafParamsParams) WithDefaults() *LeafParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the leaf params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LeafParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the leaf params params
func (o *LeafParamsParams) WithTimeout(timeout time.Duration) *LeafParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the leaf params params
func (o *LeafParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the leaf params params
func (o *LeafParamsParams) WithContext(ctx context.Context) *LeafParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the leaf params params
func (o *LeafParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the leaf params params
func (o *LeafParamsParams) WithHTTPClient(client *http.Client) *LeafParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the leaf params params
func (o *LeafParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the leaf params params
func (o *LeafParamsParams) WithBody(body *models.SigningRequestInformation) *LeafParamsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the leaf params params
func (o *LeafParamsParams) SetBody(body *models.SigningRequestInformation) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LeafParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
