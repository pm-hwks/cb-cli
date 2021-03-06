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

// SupportedDatabaseEntryResponse supported database entry response
// swagger:model SupportedDatabaseEntryResponse

type SupportedDatabaseEntryResponse struct {

	// Name of the database
	DatabaseName string `json:"databaseName,omitempty"`

	// Display name of the database
	DisplayName string `json:"displayName,omitempty"`

	// Jdbc prefix of the database
	JdbcPrefix string `json:"jdbcPrefix,omitempty"`

	// Supported version types currently only for Oracle
	// Unique: true
	Versions []string `json:"versions"`
}

/* polymorph SupportedDatabaseEntryResponse databaseName false */

/* polymorph SupportedDatabaseEntryResponse displayName false */

/* polymorph SupportedDatabaseEntryResponse jdbcPrefix false */

/* polymorph SupportedDatabaseEntryResponse versions false */

// Validate validates this supported database entry response
func (m *SupportedDatabaseEntryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedDatabaseEntryResponse) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	if err := validate.UniqueItems("versions", "body", m.Versions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedDatabaseEntryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedDatabaseEntryResponse) UnmarshalBinary(b []byte) error {
	var res SupportedDatabaseEntryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
