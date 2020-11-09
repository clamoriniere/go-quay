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

// UserView Describes a user
//
// swagger:model UserView
type UserView struct {

	// true if this user data represents a guest user
	// Required: true
	Anonymous *bool `json:"anonymous"`

	// Avatar data representing the user's icon
	// Required: true
	Avatar interface{} `json:"avatar"`

	// Whether the user has permission to create repositories
	CanCreateRepo bool `json:"can_create_repo,omitempty"`

	// The user's email address
	Email string `json:"email,omitempty"`

	// The list of external login providers against which the user has authenticated
	Logins []interface{} `json:"logins"`

	// Information about the organizations in which the user is a member
	Organizations []interface{} `json:"organizations"`

	// If true, the user's namespace is the preferred namespace to display
	PreferredNamespace bool `json:"preferred_namespace,omitempty"`

	// Whether the user's email address has been verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this user view
func (m *UserView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnonymous(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserView) validateAnonymous(formats strfmt.Registry) error {

	if err := validate.Required("anonymous", "body", m.Anonymous); err != nil {
		return err
	}

	return nil
}

func (m *UserView) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserView) UnmarshalBinary(b []byte) error {
	var res UserView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
