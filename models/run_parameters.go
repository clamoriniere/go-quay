// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunParameters Optional run parameters for activating the build trigger
//
// swagger:model RunParameters
type RunParameters struct {

	// (SCM only) If specified, the name of the branch to build.
	BranchName string `json:"branch_name,omitempty"`

	// (Custom Only) If specified, the ref/SHA1 used to checkout a git repository.
	CommitSha string `json:"commit_sha,omitempty"`

	// (SCM Only) If specified, the ref to build.
	Refs interface{} `json:"refs,omitempty"`
}

// Validate validates this run parameters
func (m *RunParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunParameters) UnmarshalBinary(b []byte) error {
	var res RunParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
