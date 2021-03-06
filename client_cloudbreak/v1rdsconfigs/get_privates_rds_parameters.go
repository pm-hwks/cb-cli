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

// NewGetPrivatesRdsParams creates a new GetPrivatesRdsParams object
// with the default values initialized.
func NewGetPrivatesRdsParams() *GetPrivatesRdsParams {

	return &GetPrivatesRdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivatesRdsParamsWithTimeout creates a new GetPrivatesRdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivatesRdsParamsWithTimeout(timeout time.Duration) *GetPrivatesRdsParams {

	return &GetPrivatesRdsParams{

		timeout: timeout,
	}
}

// NewGetPrivatesRdsParamsWithContext creates a new GetPrivatesRdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivatesRdsParamsWithContext(ctx context.Context) *GetPrivatesRdsParams {

	return &GetPrivatesRdsParams{

		Context: ctx,
	}
}

// NewGetPrivatesRdsParamsWithHTTPClient creates a new GetPrivatesRdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivatesRdsParamsWithHTTPClient(client *http.Client) *GetPrivatesRdsParams {

	return &GetPrivatesRdsParams{
		HTTPClient: client,
	}
}

/*GetPrivatesRdsParams contains all the parameters to send to the API endpoint
for the get privates rds operation typically these are written to a http.Request
*/
type GetPrivatesRdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get privates rds params
func (o *GetPrivatesRdsParams) WithTimeout(timeout time.Duration) *GetPrivatesRdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get privates rds params
func (o *GetPrivatesRdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get privates rds params
func (o *GetPrivatesRdsParams) WithContext(ctx context.Context) *GetPrivatesRdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get privates rds params
func (o *GetPrivatesRdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get privates rds params
func (o *GetPrivatesRdsParams) WithHTTPClient(client *http.Client) *GetPrivatesRdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get privates rds params
func (o *GetPrivatesRdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivatesRdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
