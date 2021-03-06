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

// Constraint constraint
// swagger:model Constraint

type Constraint struct {

	// name of a constraint template that defines the resource constraints for the hostgroup
	ConstraintTemplateName string `json:"constraintTemplateName,omitempty"`

	// number of hosts in the hostgroup
	// Required: true
	HostCount *int32 `json:"hostCount"`

	// name of an instance group where the hostgroup will be deployed
	InstanceGroupName string `json:"instanceGroupName,omitempty"`
}

/* polymorph Constraint constraintTemplateName false */

/* polymorph Constraint hostCount false */

/* polymorph Constraint instanceGroupName false */

// Validate validates this constraint
func (m *Constraint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Constraint) validateHostCount(formats strfmt.Registry) error {

	if err := validate.Required("hostCount", "body", m.HostCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Constraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Constraint) UnmarshalBinary(b []byte) error {
	var res Constraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
