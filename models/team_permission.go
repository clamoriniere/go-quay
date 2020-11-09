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

// TeamPermission Description of a team permission.
//
// swagger:model TeamPermission
type TeamPermission struct {

	// Role to use for the team
	// Required: true
	// Enum: [read write admin]
	Role *string `json:"role"`
}

// Validate validates this team permission
func (m *TeamPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var teamPermissionTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teamPermissionTypeRolePropEnum = append(teamPermissionTypeRolePropEnum, v)
	}
}

const (

	// TeamPermissionRoleRead captures enum value "read"
	TeamPermissionRoleRead string = "read"

	// TeamPermissionRoleWrite captures enum value "write"
	TeamPermissionRoleWrite string = "write"

	// TeamPermissionRoleAdmin captures enum value "admin"
	TeamPermissionRoleAdmin string = "admin"
)

// prop value enum
func (m *TeamPermission) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, teamPermissionTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TeamPermission) validateRole(formats strfmt.Registry) error {

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
func (m *TeamPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamPermission) UnmarshalBinary(b []byte) error {
	var res TeamPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
