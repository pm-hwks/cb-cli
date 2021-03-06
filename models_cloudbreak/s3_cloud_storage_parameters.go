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

// S3CloudStorageParameters s3 cloud storage parameters
// swagger:model S3CloudStorageParameters

type S3CloudStorageParameters struct {

	// instance profile
	// Required: true
	InstanceProfile *string `json:"instanceProfile"`
}

/* polymorph S3CloudStorageParameters instanceProfile false */

// Validate validates this s3 cloud storage parameters
func (m *S3CloudStorageParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3CloudStorageParameters) validateInstanceProfile(formats strfmt.Registry) error {

	if err := validate.Required("instanceProfile", "body", m.InstanceProfile); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3CloudStorageParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3CloudStorageParameters) UnmarshalBinary(b []byte) error {
	var res S3CloudStorageParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
