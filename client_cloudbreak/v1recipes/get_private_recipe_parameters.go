// Code generated by go-swagger; DO NOT EDIT.

package v1recipes

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

// NewGetPrivateRecipeParams creates a new GetPrivateRecipeParams object
// with the default values initialized.
func NewGetPrivateRecipeParams() *GetPrivateRecipeParams {
	var ()
	return &GetPrivateRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateRecipeParamsWithTimeout creates a new GetPrivateRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateRecipeParamsWithTimeout(timeout time.Duration) *GetPrivateRecipeParams {
	var ()
	return &GetPrivateRecipeParams{

		timeout: timeout,
	}
}

// NewGetPrivateRecipeParamsWithContext creates a new GetPrivateRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateRecipeParamsWithContext(ctx context.Context) *GetPrivateRecipeParams {
	var ()
	return &GetPrivateRecipeParams{

		Context: ctx,
	}
}

// NewGetPrivateRecipeParamsWithHTTPClient creates a new GetPrivateRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateRecipeParamsWithHTTPClient(client *http.Client) *GetPrivateRecipeParams {
	var ()
	return &GetPrivateRecipeParams{
		HTTPClient: client,
	}
}

/*GetPrivateRecipeParams contains all the parameters to send to the API endpoint
for the get private recipe operation typically these are written to a http.Request
*/
type GetPrivateRecipeParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private recipe params
func (o *GetPrivateRecipeParams) WithTimeout(timeout time.Duration) *GetPrivateRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private recipe params
func (o *GetPrivateRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private recipe params
func (o *GetPrivateRecipeParams) WithContext(ctx context.Context) *GetPrivateRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private recipe params
func (o *GetPrivateRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private recipe params
func (o *GetPrivateRecipeParams) WithHTTPClient(client *http.Client) *GetPrivateRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private recipe params
func (o *GetPrivateRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get private recipe params
func (o *GetPrivateRecipeParams) WithName(name string) *GetPrivateRecipeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get private recipe params
func (o *GetPrivateRecipeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
