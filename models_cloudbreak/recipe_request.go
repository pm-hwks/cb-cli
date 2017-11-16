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

// RecipeRequest recipe request
// swagger:model RecipeRequest

type RecipeRequest struct {

	// content of recipe
	Content string `json:"content,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// name of the resource
	// Max Length: 100
	// Min Length: 1
	// Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	Name string `json:"name,omitempty"`

	// type of recipe
	// Required: true
	RecipeType *string `json:"recipeType"`

	// recipe uri
	URI string `json:"uri,omitempty"`
}

/* polymorph RecipeRequest content false */

/* polymorph RecipeRequest description false */

/* polymorph RecipeRequest name false */

/* polymorph RecipeRequest recipeType false */

/* polymorph RecipeRequest uri false */

// Validate validates this recipe request
func (m *RecipeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeRequest) validateDescription(formats strfmt.Registry) error {

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

func (m *RecipeRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

var recipeRequestTypeRecipeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRE_AMBARI_START","POST_AMBARI_START","POST_CLUSTER_INSTALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recipeRequestTypeRecipeTypePropEnum = append(recipeRequestTypeRecipeTypePropEnum, v)
	}
}

const (
	// RecipeRequestRecipeTypePREAMBARISTART captures enum value "PRE_AMBARI_START"
	RecipeRequestRecipeTypePREAMBARISTART string = "PRE_AMBARI_START"
	// RecipeRequestRecipeTypePOSTAMBARISTART captures enum value "POST_AMBARI_START"
	RecipeRequestRecipeTypePOSTAMBARISTART string = "POST_AMBARI_START"
	// RecipeRequestRecipeTypePOSTCLUSTERINSTALL captures enum value "POST_CLUSTER_INSTALL"
	RecipeRequestRecipeTypePOSTCLUSTERINSTALL string = "POST_CLUSTER_INSTALL"
)

// prop value enum
func (m *RecipeRequest) validateRecipeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recipeRequestTypeRecipeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RecipeRequest) validateRecipeType(formats strfmt.Registry) error {

	if err := validate.Required("recipeType", "body", m.RecipeType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipeTypeEnum("recipeType", "body", *m.RecipeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecipeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeRequest) UnmarshalBinary(b []byte) error {
	var res RecipeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
