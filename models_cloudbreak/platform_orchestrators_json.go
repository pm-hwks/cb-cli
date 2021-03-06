// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlatformOrchestratorsJSON platform orchestrators Json
// swagger:model PlatformOrchestratorsJson

type PlatformOrchestratorsJSON struct {

	// default orchestrators
	Defaults map[string]string `json:"defaults,omitempty"`

	// orchestrators
	Orchestrators map[string][]string `json:"orchestrators,omitempty"`
}

/* polymorph PlatformOrchestratorsJson defaults false */

/* polymorph PlatformOrchestratorsJson orchestrators false */

// Validate validates this platform orchestrators Json
func (m *PlatformOrchestratorsJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrchestrators(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformOrchestratorsJSON) validateOrchestrators(formats strfmt.Registry) error {

	if swag.IsZero(m.Orchestrators) { // not required
		return nil
	}

	if swag.IsZero(m.Orchestrators) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformOrchestratorsJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformOrchestratorsJSON) UnmarshalBinary(b []byte) error {
	var res PlatformOrchestratorsJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
