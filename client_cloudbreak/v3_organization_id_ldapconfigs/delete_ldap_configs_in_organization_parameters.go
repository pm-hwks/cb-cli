// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_ldapconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLdapConfigsInOrganizationParams creates a new DeleteLdapConfigsInOrganizationParams object
// with the default values initialized.
func NewDeleteLdapConfigsInOrganizationParams() *DeleteLdapConfigsInOrganizationParams {
	var ()
	return &DeleteLdapConfigsInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLdapConfigsInOrganizationParamsWithTimeout creates a new DeleteLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLdapConfigsInOrganizationParamsWithTimeout(timeout time.Duration) *DeleteLdapConfigsInOrganizationParams {
	var ()
	return &DeleteLdapConfigsInOrganizationParams{

		timeout: timeout,
	}
}

// NewDeleteLdapConfigsInOrganizationParamsWithContext creates a new DeleteLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLdapConfigsInOrganizationParamsWithContext(ctx context.Context) *DeleteLdapConfigsInOrganizationParams {
	var ()
	return &DeleteLdapConfigsInOrganizationParams{

		Context: ctx,
	}
}

// NewDeleteLdapConfigsInOrganizationParamsWithHTTPClient creates a new DeleteLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLdapConfigsInOrganizationParamsWithHTTPClient(client *http.Client) *DeleteLdapConfigsInOrganizationParams {
	var ()
	return &DeleteLdapConfigsInOrganizationParams{
		HTTPClient: client,
	}
}

/*DeleteLdapConfigsInOrganizationParams contains all the parameters to send to the API endpoint
for the delete ldap configs in organization operation typically these are written to a http.Request
*/
type DeleteLdapConfigsInOrganizationParams struct {

	/*LdapConfigName*/
	LdapConfigName string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) WithTimeout(timeout time.Duration) *DeleteLdapConfigsInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) WithContext(ctx context.Context) *DeleteLdapConfigsInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) WithHTTPClient(client *http.Client) *DeleteLdapConfigsInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLdapConfigName adds the ldapConfigName to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) WithLdapConfigName(ldapConfigName string) *DeleteLdapConfigsInOrganizationParams {
	o.SetLdapConfigName(ldapConfigName)
	return o
}

// SetLdapConfigName adds the ldapConfigName to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) SetLdapConfigName(ldapConfigName string) {
	o.LdapConfigName = ldapConfigName
}

// WithOrganizationID adds the organizationID to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) WithOrganizationID(organizationID int64) *DeleteLdapConfigsInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete ldap configs in organization params
func (o *DeleteLdapConfigsInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLdapConfigsInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ldapConfigName
	if err := r.SetPathParam("ldapConfigName", o.LdapConfigName); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
