// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

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

// NewGetPrivateRdsParams creates a new GetPrivateRdsParams object
// with the default values initialized.
func NewGetPrivateRdsParams() *GetPrivateRdsParams {
	var ()
	return &GetPrivateRdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateRdsParamsWithTimeout creates a new GetPrivateRdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateRdsParamsWithTimeout(timeout time.Duration) *GetPrivateRdsParams {
	var ()
	return &GetPrivateRdsParams{

		timeout: timeout,
	}
}

// NewGetPrivateRdsParamsWithContext creates a new GetPrivateRdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateRdsParamsWithContext(ctx context.Context) *GetPrivateRdsParams {
	var ()
	return &GetPrivateRdsParams{

		Context: ctx,
	}
}

// NewGetPrivateRdsParamsWithHTTPClient creates a new GetPrivateRdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateRdsParamsWithHTTPClient(client *http.Client) *GetPrivateRdsParams {
	var ()
	return &GetPrivateRdsParams{
		HTTPClient: client,
	}
}

/*GetPrivateRdsParams contains all the parameters to send to the API endpoint
for the get private rds operation typically these are written to a http.Request
*/
type GetPrivateRdsParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private rds params
func (o *GetPrivateRdsParams) WithTimeout(timeout time.Duration) *GetPrivateRdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private rds params
func (o *GetPrivateRdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private rds params
func (o *GetPrivateRdsParams) WithContext(ctx context.Context) *GetPrivateRdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private rds params
func (o *GetPrivateRdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private rds params
func (o *GetPrivateRdsParams) WithHTTPClient(client *http.Client) *GetPrivateRdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private rds params
func (o *GetPrivateRdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get private rds params
func (o *GetPrivateRdsParams) WithName(name string) *GetPrivateRdsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get private rds params
func (o *GetPrivateRdsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateRdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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