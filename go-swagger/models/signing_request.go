// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SigningRequest signing request
//
// swagger:model SigningRequest
type SigningRequest struct {

	// Certificate Signing Request
	// Required: true
	CSR *string `json:"csr"`

	// Boolean indicating whether to return a root certificate
	Chain bool `json:"chain,omitempty"`

	// Path to secret - registered root CA
	// Required: true
	RootCAPath *string `json:"rootCAPath"`

	// A list of Subject Alternative Names for a certificate (each domain must be present in the list of allowed domains)
	SubjectAltNames []string `json:"subjectAltNames"`

	// TTL for a generated certificate (in hours, cannot exceed the maximum TTL specified in root CA secret)
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this signing request
func (m *SigningRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCSR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCAPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigningRequest) validateCSR(formats strfmt.Registry) error {

	if err := validate.Required("csr", "body", m.CSR); err != nil {
		return err
	}

	return nil
}

func (m *SigningRequest) validateRootCAPath(formats strfmt.Registry) error {

	if err := validate.Required("rootCAPath", "body", m.RootCAPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signing request based on context it is used
func (m *SigningRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SigningRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningRequest) UnmarshalBinary(b []byte) error {
	var res SigningRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
