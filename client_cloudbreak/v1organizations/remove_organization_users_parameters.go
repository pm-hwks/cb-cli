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

// NewRemoveOrganizationUsersParams creates a new RemoveOrganizationUsersParams object
// with the default values initialized.
func NewRemoveOrganizationUsersParams() *RemoveOrganizationUsersParams {
	var ()
	return &RemoveOrganizationUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveOrganizationUsersParamsWithTimeout creates a new RemoveOrganizationUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveOrganizationUsersParamsWithTimeout(timeout time.Duration) *RemoveOrganizationUsersParams {
	var ()
	return &RemoveOrganizationUsersParams{

		timeout: timeout,
	}
}

// NewRemoveOrganizationUsersParamsWithContext creates a new RemoveOrganizationUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveOrganizationUsersParamsWithContext(ctx context.Context) *RemoveOrganizationUsersParams {
	var ()
	return &RemoveOrganizationUsersParams{

		Context: ctx,
	}
}

// NewRemoveOrganizationUsersParamsWithHTTPClient creates a new RemoveOrganizationUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveOrganizationUsersParamsWithHTTPClient(client *http.Client) *RemoveOrganizationUsersParams {
	var ()
	return &RemoveOrganizationUsersParams{
		HTTPClient: client,
	}
}

/*RemoveOrganizationUsersParams contains all the parameters to send to the API endpoint
for the remove organization users operation typically these are written to a http.Request
*/
type RemoveOrganizationUsersParams struct {

	/*Body*/
	Body []string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove organization users params
func (o *RemoveOrganizationUsersParams) WithTimeout(timeout time.Duration) *RemoveOrganizationUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove organization users params
func (o *RemoveOrganizationUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove organization users params
func (o *RemoveOrganizationUsersParams) WithContext(ctx context.Context) *RemoveOrganizationUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove organization users params
func (o *RemoveOrganizationUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove organization users params
func (o *RemoveOrganizationUsersParams) WithHTTPClient(client *http.Client) *RemoveOrganizationUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove organization users params
func (o *RemoveOrganizationUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove organization users params
func (o *RemoveOrganizationUsersParams) WithBody(body []string) *RemoveOrganizationUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove organization users params
func (o *RemoveOrganizationUsersParams) SetBody(body []string) {
	o.Body = body
}

// WithName adds the name to the remove organization users params
func (o *RemoveOrganizationUsersParams) WithName(name string) *RemoveOrganizationUsersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove organization users params
func (o *RemoveOrganizationUsersParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveOrganizationUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
