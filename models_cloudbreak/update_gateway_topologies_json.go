// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateGatewayTopologiesJSON update gateway topologies Json
// swagger:model UpdateGatewayTopologiesJson

type UpdateGatewayTopologiesJSON struct {

	// topologies
	Topologies []*GatewayTopologyJSON `json:"topologies"`
}

/* polymorph UpdateGatewayTopologiesJson topologies false */

// Validate validates this update gateway topologies Json
func (m *UpdateGatewayTopologiesJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopologies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateGatewayTopologiesJSON) validateTopologies(formats strfmt.Registry) error {

	if swag.IsZero(m.Topologies) { // not required
		return nil
	}

	for i := 0; i < len(m.Topologies); i++ {

		if swag.IsZero(m.Topologies[i]) { // not required
			continue
		}

		if m.Topologies[i] != nil {

			if err := m.Topologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateGatewayTopologiesJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateGatewayTopologiesJSON) UnmarshalBinary(b []byte) error {
	var res UpdateGatewayTopologiesJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
