// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

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

// NewGetAllStackV2Params creates a new GetAllStackV2Params object
// with the default values initialized.
func NewGetAllStackV2Params() *GetAllStackV2Params {

	return &GetAllStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllStackV2ParamsWithTimeout creates a new GetAllStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllStackV2ParamsWithTimeout(timeout time.Duration) *GetAllStackV2Params {

	return &GetAllStackV2Params{

		timeout: timeout,
	}
}

// NewGetAllStackV2ParamsWithContext creates a new GetAllStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllStackV2ParamsWithContext(ctx context.Context) *GetAllStackV2Params {

	return &GetAllStackV2Params{

		Context: ctx,
	}
}

// NewGetAllStackV2ParamsWithHTTPClient creates a new GetAllStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllStackV2ParamsWithHTTPClient(client *http.Client) *GetAllStackV2Params {

	return &GetAllStackV2Params{
		HTTPClient: client,
	}
}

/*GetAllStackV2Params contains all the parameters to send to the API endpoint
for the get all stack v2 operation typically these are written to a http.Request
*/
type GetAllStackV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all stack v2 params
func (o *GetAllStackV2Params) WithTimeout(timeout time.Duration) *GetAllStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all stack v2 params
func (o *GetAllStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all stack v2 params
func (o *GetAllStackV2Params) WithContext(ctx context.Context) *GetAllStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all stack v2 params
func (o *GetAllStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all stack v2 params
func (o *GetAllStackV2Params) WithHTTPClient(client *http.Client) *GetAllStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all stack v2 params
func (o *GetAllStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
