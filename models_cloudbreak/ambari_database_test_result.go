// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AmbariDatabaseTestResult ambari database test result
// swagger:model AmbariDatabaseTestResult

type AmbariDatabaseTestResult struct {

	// result of Ambari database test
	// Required: true
	Error *string `json:"error"`
}

/* polymorph AmbariDatabaseTestResult error false */

// Validate validates this ambari database test result
func (m *AmbariDatabaseTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmbariDatabaseTestResult) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AmbariDatabaseTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariDatabaseTestResult) UnmarshalBinary(b []byte) error {
	var res AmbariDatabaseTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
