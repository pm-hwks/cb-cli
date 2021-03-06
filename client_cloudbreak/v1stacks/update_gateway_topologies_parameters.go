// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewUpdateGatewayTopologiesParams creates a new UpdateGatewayTopologiesParams object
// with the default values initialized.
func NewUpdateGatewayTopologiesParams() *UpdateGatewayTopologiesParams {
	var ()
	return &UpdateGatewayTopologiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGatewayTopologiesParamsWithTimeout creates a new UpdateGatewayTopologiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGatewayTopologiesParamsWithTimeout(timeout time.Duration) *UpdateGatewayTopologiesParams {
	var ()
	return &UpdateGatewayTopologiesParams{

		timeout: timeout,
	}
}

// NewUpdateGatewayTopologiesParamsWithContext creates a new UpdateGatewayTopologiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGatewayTopologiesParamsWithContext(ctx context.Context) *UpdateGatewayTopologiesParams {
	var ()
	return &UpdateGatewayTopologiesParams{

		Context: ctx,
	}
}

// NewUpdateGatewayTopologiesParamsWithHTTPClient creates a new UpdateGatewayTopologiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGatewayTopologiesParamsWithHTTPClient(client *http.Client) *UpdateGatewayTopologiesParams {
	var ()
	return &UpdateGatewayTopologiesParams{
		HTTPClient: client,
	}
}

/*UpdateGatewayTopologiesParams contains all the parameters to send to the API endpoint
for the update gateway topologies operation typically these are written to a http.Request
*/
type UpdateGatewayTopologiesParams struct {

	/*Body*/
	Body *models_cloudbreak.UpdateGatewayTopologiesJSON
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) WithTimeout(timeout time.Duration) *UpdateGatewayTopologiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) WithContext(ctx context.Context) *UpdateGatewayTopologiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) WithHTTPClient(client *http.Client) *UpdateGatewayTopologiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) WithBody(body *models_cloudbreak.UpdateGatewayTopologiesJSON) *UpdateGatewayTopologiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) SetBody(body *models_cloudbreak.UpdateGatewayTopologiesJSON) {
	o.Body = body
}

// WithID adds the id to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) WithID(id int64) *UpdateGatewayTopologiesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update gateway topologies params
func (o *UpdateGatewayTopologiesParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGatewayTopologiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.UpdateGatewayTopologiesJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
