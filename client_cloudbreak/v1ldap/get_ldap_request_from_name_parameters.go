// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

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

// NewGetLdapRequestFromNameParams creates a new GetLdapRequestFromNameParams object
// with the default values initialized.
func NewGetLdapRequestFromNameParams() *GetLdapRequestFromNameParams {
	var ()
	return &GetLdapRequestFromNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapRequestFromNameParamsWithTimeout creates a new GetLdapRequestFromNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapRequestFromNameParamsWithTimeout(timeout time.Duration) *GetLdapRequestFromNameParams {
	var ()
	return &GetLdapRequestFromNameParams{

		timeout: timeout,
	}
}

// NewGetLdapRequestFromNameParamsWithContext creates a new GetLdapRequestFromNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapRequestFromNameParamsWithContext(ctx context.Context) *GetLdapRequestFromNameParams {
	var ()
	return &GetLdapRequestFromNameParams{

		Context: ctx,
	}
}

// NewGetLdapRequestFromNameParamsWithHTTPClient creates a new GetLdapRequestFromNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapRequestFromNameParamsWithHTTPClient(client *http.Client) *GetLdapRequestFromNameParams {
	var ()
	return &GetLdapRequestFromNameParams{
		HTTPClient: client,
	}
}

/*GetLdapRequestFromNameParams contains all the parameters to send to the API endpoint
for the get ldap request from name operation typically these are written to a http.Request
*/
type GetLdapRequestFromNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) WithTimeout(timeout time.Duration) *GetLdapRequestFromNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) WithContext(ctx context.Context) *GetLdapRequestFromNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) WithHTTPClient(client *http.Client) *GetLdapRequestFromNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) WithName(name string) *GetLdapRequestFromNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get ldap request from name params
func (o *GetLdapRequestFromNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapRequestFromNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
