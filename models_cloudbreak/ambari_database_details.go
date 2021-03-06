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

// AmbariDatabaseDetails ambari database details
// swagger:model AmbariDatabaseDetails

type AmbariDatabaseDetails struct {

	// host of the Ambari database
	// Required: true
	// Pattern: ^[a-zA-Z0-9]([a-zA-Z0-9-\.]+)$
	Host *string `json:"host"`

	// name of the Ambari database
	// Required: true
	// Pattern: ^[^']+$
	Name *string `json:"name"`

	// password for the Ambari database
	// Required: true
	// Pattern: ^[^']+$
	Password *string `json:"password"`

	// port of the Ambari database
	// Required: true
	Port *int32 `json:"port"`

	// user name for the Ambari database
	// Required: true
	// Pattern: ^[^']+$
	UserName *string `json:"userName"`

	// vendor of the Ambari database
	// Required: true
	Vendor *string `json:"vendor"`
}

/* polymorph AmbariDatabaseDetails host false */

/* polymorph AmbariDatabaseDetails name false */

/* polymorph AmbariDatabaseDetails password false */

/* polymorph AmbariDatabaseDetails port false */

/* polymorph AmbariDatabaseDetails userName false */

/* polymorph AmbariDatabaseDetails vendor false */

// Validate validates this ambari database details
func (m *AmbariDatabaseDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmbariDatabaseDetails) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if err := validate.Pattern("host", "body", string(*m.Host), `^[a-zA-Z0-9]([a-zA-Z0-9-\.]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *AmbariDatabaseDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[^']+$`); err != nil {
		return err
	}

	return nil
}

func (m *AmbariDatabaseDetails) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.Pattern("password", "body", string(*m.Password), `^[^']+$`); err != nil {
		return err
	}

	return nil
}

func (m *AmbariDatabaseDetails) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *AmbariDatabaseDetails) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	if err := validate.Pattern("userName", "body", string(*m.UserName), `^[^']+$`); err != nil {
		return err
	}

	return nil
}

var ambariDatabaseDetailsTypeVendorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["POSTGRES","MYSQL","MARIADB","MSSQL","ORACLE11","ORACLE12","SQLANYWHERE","EMBEDDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ambariDatabaseDetailsTypeVendorPropEnum = append(ambariDatabaseDetailsTypeVendorPropEnum, v)
	}
}

const (
	// AmbariDatabaseDetailsVendorPOSTGRES captures enum value "POSTGRES"
	AmbariDatabaseDetailsVendorPOSTGRES string = "POSTGRES"
	// AmbariDatabaseDetailsVendorMYSQL captures enum value "MYSQL"
	AmbariDatabaseDetailsVendorMYSQL string = "MYSQL"
	// AmbariDatabaseDetailsVendorMARIADB captures enum value "MARIADB"
	AmbariDatabaseDetailsVendorMARIADB string = "MARIADB"
	// AmbariDatabaseDetailsVendorMSSQL captures enum value "MSSQL"
	AmbariDatabaseDetailsVendorMSSQL string = "MSSQL"
	// AmbariDatabaseDetailsVendorORACLE11 captures enum value "ORACLE11"
	AmbariDatabaseDetailsVendorORACLE11 string = "ORACLE11"
	// AmbariDatabaseDetailsVendorORACLE12 captures enum value "ORACLE12"
	AmbariDatabaseDetailsVendorORACLE12 string = "ORACLE12"
	// AmbariDatabaseDetailsVendorSQLANYWHERE captures enum value "SQLANYWHERE"
	AmbariDatabaseDetailsVendorSQLANYWHERE string = "SQLANYWHERE"
	// AmbariDatabaseDetailsVendorEMBEDDED captures enum value "EMBEDDED"
	AmbariDatabaseDetailsVendorEMBEDDED string = "EMBEDDED"
)

// prop value enum
func (m *AmbariDatabaseDetails) validateVendorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ambariDatabaseDetailsTypeVendorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AmbariDatabaseDetails) validateVendor(formats strfmt.Registry) error {

	if err := validate.Required("vendor", "body", m.Vendor); err != nil {
		return err
	}

	// value enum
	if err := m.validateVendorEnum("vendor", "body", *m.Vendor); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AmbariDatabaseDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariDatabaseDetails) UnmarshalBinary(b []byte) error {
	var res AmbariDatabaseDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
