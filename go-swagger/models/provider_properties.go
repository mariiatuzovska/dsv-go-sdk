// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProviderProperties AuthenticationSettings is the 3rd party authentication providers.
//
// swagger:model ProviderProperties
type ProviderProperties struct {

	// account ID
	AccountID string `json:"accountId,omitempty"`

	// base URI
	BaseURI string `json:"baseUri,omitempty"`

	// client email
	ClientEmail string `json:"clientEmail,omitempty"`

	// client ID
	ClientID string `json:"clientId,omitempty"`

	// client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`

	// private key ID
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// project ID
	ProjectID string `json:"projectId,omitempty"`

	// send welcome email
	SendWelcomeEmail bool `json:"sendWelcomeEmail,omitempty"`

	// tenant ID
	TenantID string `json:"tenantId,omitempty"`

	// token URI
	TokenURI string `json:"tokenUri,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// username claim
	UsernameClaim string `json:"usernameClaim,omitempty"`
}

// Validate validates this provider properties
func (m *ProviderProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this provider properties based on context it is used
func (m *ProviderProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderProperties) UnmarshalBinary(b []byte) error {
	var res ProviderProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
