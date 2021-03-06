// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPrototype Description of a new prototype
//
// swagger:model NewPrototype
type NewPrototype struct {

	// activating user
	ActivatingUser *NewPrototypeActivatingUser `json:"activating_user,omitempty"`

	// delegate
	// Required: true
	Delegate *NewPrototypeDelegate `json:"delegate"`

	// Role that should be applied to the delegate
	// Required: true
	// Enum: [read write admin]
	Role *string `json:"role"`
}

// Validate validates this new prototype
func (m *NewPrototype) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivatingUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototype) validateActivatingUser(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivatingUser) { // not required
		return nil
	}

	if m.ActivatingUser != nil {
		if err := m.ActivatingUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activating_user")
			}
			return err
		}
	}

	return nil
}

func (m *NewPrototype) validateDelegate(formats strfmt.Registry) error {

	if err := validate.Required("delegate", "body", m.Delegate); err != nil {
		return err
	}

	if m.Delegate != nil {
		if err := m.Delegate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegate")
			}
			return err
		}
	}

	return nil
}

var newPrototypeTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newPrototypeTypeRolePropEnum = append(newPrototypeTypeRolePropEnum, v)
	}
}

const (

	// NewPrototypeRoleRead captures enum value "read"
	NewPrototypeRoleRead string = "read"

	// NewPrototypeRoleWrite captures enum value "write"
	NewPrototypeRoleWrite string = "write"

	// NewPrototypeRoleAdmin captures enum value "admin"
	NewPrototypeRoleAdmin string = "admin"
)

// prop value enum
func (m *NewPrototype) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, newPrototypeTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototype) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPrototype) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPrototype) UnmarshalBinary(b []byte) error {
	var res NewPrototype
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NewPrototypeActivatingUser Repository creating user to whom the rule should apply
//
// swagger:model NewPrototypeActivatingUser
type NewPrototypeActivatingUser struct {

	// The username for the activating_user
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this new prototype activating user
func (m *NewPrototypeActivatingUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototypeActivatingUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("activating_user"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPrototypeActivatingUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPrototypeActivatingUser) UnmarshalBinary(b []byte) error {
	var res NewPrototypeActivatingUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NewPrototypeDelegate Information about the user or team to which the rule grants access
//
// swagger:model NewPrototypeDelegate
type NewPrototypeDelegate struct {

	// Whether the delegate is a user or a team
	// Required: true
	// Enum: [user team]
	Kind *string `json:"kind"`

	// The name for the delegate team or user
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this new prototype delegate
func (m *NewPrototypeDelegate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var newPrototypeDelegateTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","team"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newPrototypeDelegateTypeKindPropEnum = append(newPrototypeDelegateTypeKindPropEnum, v)
	}
}

const (

	// NewPrototypeDelegateKindUser captures enum value "user"
	NewPrototypeDelegateKindUser string = "user"

	// NewPrototypeDelegateKindTeam captures enum value "team"
	NewPrototypeDelegateKindTeam string = "team"
)

// prop value enum
func (m *NewPrototypeDelegate) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, newPrototypeDelegateTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototypeDelegate) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("delegate"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("delegate"+"."+"kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *NewPrototypeDelegate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("delegate"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPrototypeDelegate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPrototypeDelegate) UnmarshalBinary(b []byte) error {
	var res NewPrototypeDelegate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
