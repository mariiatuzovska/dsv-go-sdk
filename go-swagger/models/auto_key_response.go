// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AutoKeyResponse AutoKeyResponse contains the metadata of a data key
//
// swagger:model AutoKeyResponse
type AutoKeyResponse struct {

	// Created date
	Created string `json:"created,omitempty"`

	// Who created the item
	CreatedBy string `json:"createdBy,omitempty"`

	// the id for this item
	ID string `json:"id,omitempty"`

	// Last updated date
	LastModified string `json:"lastModified,omitempty"`

	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// A path to a data-key managed by DSV
	Path string `json:"path,omitempty"`

	// Current version
	Version string `json:"version,omitempty"`
}

// Validate validates this auto key response
func (m *AutoKeyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auto key response based on context it is used
func (m *AutoKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AutoKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoKeyResponse) UnmarshalBinary(b []byte) error {
	var res AutoKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
