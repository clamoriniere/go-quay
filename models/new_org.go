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

// NewOrg Description of a new organization.
//
// swagger:model NewOrg
type NewOrg struct {

	// Organization contact email
	Email string `json:"email,omitempty"`

	// Organization username
	// Required: true
	Name *string `json:"name"`

	// The (may be disabled) recaptcha response code for verification
	RecaptchaResponse string `json:"recaptcha_response,omitempty"`
}

// Validate validates this new org
func (m *NewOrg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewOrg) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewOrg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewOrg) UnmarshalBinary(b []byte) error {
	var res NewOrg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
