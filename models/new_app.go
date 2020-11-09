// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewApp Description of a new organization application.
//
// swagger:model NewApp
type NewApp struct {

	// The URI for the application's homepage
	ApplicationURI string `json:"application_uri,omitempty"`

	// The e-mail address of the avatar to use for the application
	AvatarEmail string `json:"avatar_email,omitempty"`

	// The human-readable description for the application
	Description string `json:"description,omitempty"`

	// The name of the application
	// Required: true
	Name *string `json:"name"`

	// The URI for the application's OAuth redirect
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// Validate validates this new app
func (m *NewApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewApp) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewApp) UnmarshalBinary(b []byte) error {
	var res NewApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
