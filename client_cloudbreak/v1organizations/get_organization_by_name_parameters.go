// Code generated by go-swagger; DO NOT EDIT.

package v1organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrganizationByNameParams creates a new GetOrganizationByNameParams object
// with the default values initialized.
func NewGetOrganizationByNameParams() *GetOrganizationByNameParams {
	var ()
	return &GetOrganizationByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationByNameParamsWithTimeout creates a new GetOrganizationByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationByNameParamsWithTimeout(timeout time.Duration) *GetOrganizationByNameParams {
	var ()
	return &GetOrganizationByNameParams{

		timeout: timeout,
	}
}

// NewGetOrganizationByNameParamsWithContext creates a new GetOrganizationByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationByNameParamsWithContext(ctx context.Context) *GetOrganizationByNameParams {
	var ()
	return &GetOrganizationByNameParams{

		Context: ctx,
	}
}

// NewGetOrganizationByNameParamsWithHTTPClient creates a new GetOrganizationByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationByNameParamsWithHTTPClient(client *http.Client) *GetOrganizationByNameParams {
	var ()
	return &GetOrganizationByNameParams{
		HTTPClient: client,
	}
}

/*GetOrganizationByNameParams contains all the parameters to send to the API endpoint
for the get organization by name operation typically these are written to a http.Request
*/
type GetOrganizationByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization by name params
func (o *GetOrganizationByNameParams) WithTimeout(timeout time.Duration) *GetOrganizationByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization by name params
func (o *GetOrganizationByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization by name params
func (o *GetOrganizationByNameParams) WithContext(ctx context.Context) *GetOrganizationByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization by name params
func (o *GetOrganizationByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization by name params
func (o *GetOrganizationByNameParams) WithHTTPClient(client *http.Client) *GetOrganizationByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization by name params
func (o *GetOrganizationByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get organization by name params
func (o *GetOrganizationByNameParams) WithName(name string) *GetOrganizationByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get organization by name params
func (o *GetOrganizationByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
