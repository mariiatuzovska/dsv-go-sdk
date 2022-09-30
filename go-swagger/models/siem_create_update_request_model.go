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

// SiemCreateUpdateRequestModel siem create update request model
//
// swagger:model SiemCreateUpdateRequestModel
type SiemCreateUpdateRequestModel struct {

	// Denotes whether the endpoint can use self signed root certifcate for handshake (only for https protocol)
	AllowSelfSigned bool `json:"allowSelfSigned,omitempty"`

	// Authentication token
	// Required: true
	Auth *string `json:"auth"`

	// Authentication method (token)
	// Required: true
	AuthMethod *string `json:"authMethod"`

	// Endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// Collect Server IP/FQDN
	// Required: true
	Host *string `json:"host"`

	// Logging Format (i.e. syslog (RFC 5424))
	// Required: true
	LoggingFormat *string `json:"loggingFormat"`

	// Name of registered SIEM endpoint, similar to path
	// Required: true
	Name *string `json:"name"`

	// Engine pool name, used when sending request to a DSV engine instance
	Pool string `json:"pool,omitempty"`

	// Port
	// Required: true
	Port *int64 `json:"port"`

	// Type of protocol (i.e. TCP, UDP)
	// Required: true
	Protocol *string `json:"protocol"`

	// Denotes whether the endpoint should be accessed through a DSV engine instance
	SendToEngine bool `json:"sendToEngine,omitempty"`

	// Type of endpoint ("syslog", "cef", "json", "splunk")
	// Required: true
	SiemType *string `json:"siemType"`
}

// Validate validates this siem create update request model
func (m *SiemCreateUpdateRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoggingFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiemType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiemCreateUpdateRequestModel) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateAuthMethod(formats strfmt.Registry) error {

	if err := validate.Required("authMethod", "body", m.AuthMethod); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateLoggingFormat(formats strfmt.Registry) error {

	if err := validate.Required("loggingFormat", "body", m.LoggingFormat); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *SiemCreateUpdateRequestModel) validateSiemType(formats strfmt.Registry) error {

	if err := validate.Required("siemType", "body", m.SiemType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this siem create update request model based on context it is used
func (m *SiemCreateUpdateRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiemCreateUpdateRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiemCreateUpdateRequestModel) UnmarshalBinary(b []byte) error {
	var res SiemCreateUpdateRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
