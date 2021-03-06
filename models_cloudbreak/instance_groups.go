// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceGroups instance groups
// swagger:model instanceGroups

type InstanceGroups struct {

	// name of the instance group
	// Required: true
	Group *string `json:"group"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// cloud specific parameters for instance group
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// instancegroup related securitygroup
	SecurityGroup *SecurityGroupRequest `json:"securityGroup,omitempty"`

	// security group resource id for the instance group
	SecurityGroupID int64 `json:"securityGroupId,omitempty"`

	// instancegroup related template
	Template *TemplateRequest `json:"template,omitempty"`

	// referenced template id
	TemplateID int64 `json:"templateId,omitempty"`

	// type of the instance group
	Type string `json:"type,omitempty"`
}

/* polymorph instanceGroups group false */

/* polymorph instanceGroups nodeCount false */

/* polymorph instanceGroups parameters false */

/* polymorph instanceGroups securityGroup false */

/* polymorph instanceGroups securityGroupId false */

/* polymorph instanceGroups template false */

/* polymorph instanceGroups templateId false */

/* polymorph instanceGroups type false */

// Validate validates this instance groups
func (m *InstanceGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroups) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroups) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(*m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(*m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroups) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {

		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroups) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {

		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

var instanceGroupsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GATEWAY","CORE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupsTypeTypePropEnum = append(instanceGroupsTypeTypePropEnum, v)
	}
}

const (
	// InstanceGroupsTypeGATEWAY captures enum value "GATEWAY"
	InstanceGroupsTypeGATEWAY string = "GATEWAY"
	// InstanceGroupsTypeCORE captures enum value "CORE"
	InstanceGroupsTypeCORE string = "CORE"
)

// prop value enum
func (m *InstanceGroups) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroups) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroups) UnmarshalBinary(b []byte) error {
	var res InstanceGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
