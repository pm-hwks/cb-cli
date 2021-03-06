// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RdsBuildResult rds build result
// swagger:model RdsBuildResult

type RdsBuildResult struct {

	// name of the created dbs
	// Required: true
	Results map[string]string `json:"results"`
}

/* polymorph RdsBuildResult results false */

// Validate validates this rds build result
func (m *RdsBuildResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsBuildResult) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsBuildResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsBuildResult) UnmarshalBinary(b []byte) error {
	var res RdsBuildResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
