// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_rdsconfigs

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

// NewGetRdsConfigInOrganizationParams creates a new GetRdsConfigInOrganizationParams object
// with the default values initialized.
func NewGetRdsConfigInOrganizationParams() *GetRdsConfigInOrganizationParams {
	var ()
	return &GetRdsConfigInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRdsConfigInOrganizationParamsWithTimeout creates a new GetRdsConfigInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRdsConfigInOrganizationParamsWithTimeout(timeout time.Duration) *GetRdsConfigInOrganizationParams {
	var ()
	return &GetRdsConfigInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetRdsConfigInOrganizationParamsWithContext creates a new GetRdsConfigInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRdsConfigInOrganizationParamsWithContext(ctx context.Context) *GetRdsConfigInOrganizationParams {
	var ()
	return &GetRdsConfigInOrganizationParams{

		Context: ctx,
	}
}

// NewGetRdsConfigInOrganizationParamsWithHTTPClient creates a new GetRdsConfigInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRdsConfigInOrganizationParamsWithHTTPClient(client *http.Client) *GetRdsConfigInOrganizationParams {
	var ()
	return &GetRdsConfigInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetRdsConfigInOrganizationParams contains all the parameters to send to the API endpoint
for the get rds config in organization operation typically these are written to a http.Request
*/
type GetRdsConfigInOrganizationParams struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) WithTimeout(timeout time.Duration) *GetRdsConfigInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) WithContext(ctx context.Context) *GetRdsConfigInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) WithHTTPClient(client *http.Client) *GetRdsConfigInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) WithName(name string) *GetRdsConfigInOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) WithOrganizationID(organizationID int64) *GetRdsConfigInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get rds config in organization params
func (o *GetRdsConfigInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRdsConfigInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
