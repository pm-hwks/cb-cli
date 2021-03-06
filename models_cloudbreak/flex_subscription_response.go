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

// FlexSubscriptionResponse flex subscription response
// swagger:model FlexSubscriptionResponse

type FlexSubscriptionResponse struct {

	// account id of the resource owner that is provided by OAuth provider
	// Read Only: true
	Account string `json:"account,omitempty"`

	// id of the resource
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// id of the resource owner that is provided by OAuth provider
	// Read Only: true
	Owner string `json:"owner,omitempty"`

	// resource is visible in account
	// Read Only: true
	PublicInAccount *bool `json:"publicInAccount,omitempty"`

	// The associated SmartSense subscription Cloudbreak domain object json representation.
	// Read Only: true
	SmartSenseSubscription *SmartSenseSubscriptionJSON `json:"smartSenseSubscription,omitempty"`

	// Identifier of SmartSense subscription Cloudbreak domain object json representation.
	// Read Only: true
	SmartSenseSubscriptionID int64 `json:"smartSenseSubscriptionId,omitempty"`

	// Identifier of Flex subscription.
	// Read Only: true
	// Pattern: ^(FLEX-[0-9]{10}$)
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// true if the flex subscription is the default one
	UsedAsDefault *bool `json:"usedAsDefault,omitempty"`

	// true if the flex subscription was used for the controller
	UsedForController *bool `json:"usedForController,omitempty"`
}

/* polymorph FlexSubscriptionResponse account false */

/* polymorph FlexSubscriptionResponse id false */

/* polymorph FlexSubscriptionResponse name false */

/* polymorph FlexSubscriptionResponse owner false */

/* polymorph FlexSubscriptionResponse publicInAccount false */

/* polymorph FlexSubscriptionResponse smartSenseSubscription false */

/* polymorph FlexSubscriptionResponse smartSenseSubscriptionId false */

/* polymorph FlexSubscriptionResponse subscriptionId false */

/* polymorph FlexSubscriptionResponse usedAsDefault false */

/* polymorph FlexSubscriptionResponse usedForController false */

// Validate validates this flex subscription response
func (m *FlexSubscriptionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSmartSenseSubscription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexSubscriptionResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FlexSubscriptionResponse) validateSmartSenseSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.SmartSenseSubscription) { // not required
		return nil
	}

	if m.SmartSenseSubscription != nil {

		if err := m.SmartSenseSubscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smartSenseSubscription")
			}
			return err
		}
	}

	return nil
}

func (m *FlexSubscriptionResponse) validateSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.Pattern("subscriptionId", "body", string(m.SubscriptionID), `^(FLEX-[0-9]{10}$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexSubscriptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexSubscriptionResponse) UnmarshalBinary(b []byte) error {
	var res FlexSubscriptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
