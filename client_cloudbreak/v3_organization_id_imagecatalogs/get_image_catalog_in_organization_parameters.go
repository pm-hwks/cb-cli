// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_imagecatalogs

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

// NewGetImageCatalogInOrganizationParams creates a new GetImageCatalogInOrganizationParams object
// with the default values initialized.
func NewGetImageCatalogInOrganizationParams() *GetImageCatalogInOrganizationParams {
	var ()
	return &GetImageCatalogInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageCatalogInOrganizationParamsWithTimeout creates a new GetImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageCatalogInOrganizationParamsWithTimeout(timeout time.Duration) *GetImageCatalogInOrganizationParams {
	var ()
	return &GetImageCatalogInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetImageCatalogInOrganizationParamsWithContext creates a new GetImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageCatalogInOrganizationParamsWithContext(ctx context.Context) *GetImageCatalogInOrganizationParams {
	var ()
	return &GetImageCatalogInOrganizationParams{

		Context: ctx,
	}
}

// NewGetImageCatalogInOrganizationParamsWithHTTPClient creates a new GetImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageCatalogInOrganizationParamsWithHTTPClient(client *http.Client) *GetImageCatalogInOrganizationParams {
	var ()
	return &GetImageCatalogInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetImageCatalogInOrganizationParams contains all the parameters to send to the API endpoint
for the get image catalog in organization operation typically these are written to a http.Request
*/
type GetImageCatalogInOrganizationParams struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) WithTimeout(timeout time.Duration) *GetImageCatalogInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) WithContext(ctx context.Context) *GetImageCatalogInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) WithHTTPClient(client *http.Client) *GetImageCatalogInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) WithName(name string) *GetImageCatalogInOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) WithOrganizationID(organizationID int64) *GetImageCatalogInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get image catalog in organization params
func (o *GetImageCatalogInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageCatalogInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
