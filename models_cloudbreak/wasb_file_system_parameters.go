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

// WasbFileSystemParameters wasb file system parameters
// swagger:model WasbFileSystemParameters

type WasbFileSystemParameters struct {

	// account key
	AccountKey string `json:"accountKey,omitempty"`

	// account name
	AccountName string `json:"accountName,omitempty"`

	// as map
	AsMap map[string]string `json:"asMap,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph WasbFileSystemParameters accountKey false */

/* polymorph WasbFileSystemParameters accountName false */

/* polymorph WasbFileSystemParameters asMap false */

/* polymorph WasbFileSystemParameters type false */

// Validate validates this wasb file system parameters
func (m *WasbFileSystemParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wasbFileSystemParametersTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WASB_INTEGRATED","GCS","WASB","ADLS","S3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wasbFileSystemParametersTypeTypePropEnum = append(wasbFileSystemParametersTypeTypePropEnum, v)
	}
}

const (
	// WasbFileSystemParametersTypeWASBINTEGRATED captures enum value "WASB_INTEGRATED"
	WasbFileSystemParametersTypeWASBINTEGRATED string = "WASB_INTEGRATED"
	// WasbFileSystemParametersTypeGCS captures enum value "GCS"
	WasbFileSystemParametersTypeGCS string = "GCS"
	// WasbFileSystemParametersTypeWASB captures enum value "WASB"
	WasbFileSystemParametersTypeWASB string = "WASB"
	// WasbFileSystemParametersTypeADLS captures enum value "ADLS"
	WasbFileSystemParametersTypeADLS string = "ADLS"
	// WasbFileSystemParametersTypeS3 captures enum value "S3"
	WasbFileSystemParametersTypeS3 string = "S3"
)

// prop value enum
func (m *WasbFileSystemParameters) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wasbFileSystemParametersTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WasbFileSystemParameters) validateType(formats strfmt.Registry) error {

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
func (m *WasbFileSystemParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WasbFileSystemParameters) UnmarshalBinary(b []byte) error {
	var res WasbFileSystemParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}