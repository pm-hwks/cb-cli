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

// TemplateResponse template response
// swagger:model TemplateResponse

type TemplateResponse struct {

	// aws specific parameters for template
	AwsParameters *AwsParameters `json:"awsParameters,omitempty"`

	// azure specific parameters for template
	AzureParameters *AzureParameters `json:"azureParameters,omitempty"`

	// type of cloud provider
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// custom instancetype definition
	CustomInstanceType *CustomInstanceType `json:"customInstanceType,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// gcp specific parameters for template
	GcpParameters *GcpParameters `json:"gcpParameters,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// type of the instance
	InstanceType string `json:"instanceType,omitempty"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// openstack specific parameters for template
	OpenStackParameters OpenStackParameters `json:"openStackParameters,omitempty"`

	// cloud specific parameters for template
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// resource is visible in account
	Public *bool `json:"public,omitempty"`

	// size of the root volume
	RootVolumeSize int32 `json:"rootVolumeSize,omitempty"`

	// cloud specific secret parameters for template
	SecretParameters map[string]interface{} `json:"secretParameters,omitempty"`

	// id of the topology the resource belongs to
	TopologyID int64 `json:"topologyId,omitempty"`

	// number of volumes
	// Required: true
	VolumeCount *int32 `json:"volumeCount"`

	// size of volumes
	// Required: true
	VolumeSize *int32 `json:"volumeSize"`

	// type of the volumes
	VolumeType string `json:"volumeType,omitempty"`

	// yarn specific parameters for template
	YarnParameters YarnParameters `json:"yarnParameters,omitempty"`
}

/* polymorph TemplateResponse awsParameters false */

/* polymorph TemplateResponse azureParameters false */

/* polymorph TemplateResponse cloudPlatform false */

/* polymorph TemplateResponse customInstanceType false */

/* polymorph TemplateResponse description false */

/* polymorph TemplateResponse gcpParameters false */

/* polymorph TemplateResponse id false */

/* polymorph TemplateResponse instanceType false */

/* polymorph TemplateResponse name false */

/* polymorph TemplateResponse openStackParameters false */

/* polymorph TemplateResponse parameters false */

/* polymorph TemplateResponse public false */

/* polymorph TemplateResponse rootVolumeSize false */

/* polymorph TemplateResponse secretParameters false */

/* polymorph TemplateResponse topologyId false */

/* polymorph TemplateResponse volumeCount false */

/* polymorph TemplateResponse volumeSize false */

/* polymorph TemplateResponse volumeType false */

/* polymorph TemplateResponse yarnParameters false */

// Validate validates this template response
func (m *TemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAzureParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomInstanceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGcpParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateResponse) validateAwsParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsParameters) { // not required
		return nil
	}

	if m.AwsParameters != nil {

		if err := m.AwsParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParameters")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateResponse) validateAzureParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureParameters) { // not required
		return nil
	}

	if m.AzureParameters != nil {

		if err := m.AzureParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParameters")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateResponse) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateCustomInstanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomInstanceType) { // not required
		return nil
	}

	if m.CustomInstanceType != nil {

		if err := m.CustomInstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customInstanceType")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateGcpParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.GcpParameters) { // not required
		return nil
	}

	if m.GcpParameters != nil {

		if err := m.GcpParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpParameters")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateVolumeCount(formats strfmt.Registry) error {

	if err := validate.Required("volumeCount", "body", m.VolumeCount); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateVolumeSize(formats strfmt.Registry) error {

	if err := validate.Required("volumeSize", "body", m.VolumeSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateResponse) UnmarshalBinary(b []byte) error {
	var res TemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
