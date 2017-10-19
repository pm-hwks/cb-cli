// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

	"github.com/hortonworks/cb-cli/models_autoscale"
)

// NewModifyClusterParams creates a new ModifyClusterParams object
// with the default values initialized.
func NewModifyClusterParams() *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyClusterParamsWithTimeout creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyClusterParamsWithTimeout(timeout time.Duration) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		timeout: timeout,
	}
}

// NewModifyClusterParamsWithContext creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyClusterParamsWithContext(ctx context.Context) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		Context: ctx,
	}
}

// NewModifyClusterParamsWithHTTPClient creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyClusterParamsWithHTTPClient(client *http.Client) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{
		HTTPClient: client,
	}
}

/*ModifyClusterParams contains all the parameters to send to the API endpoint
for the modify cluster operation typically these are written to a http.Request
*/
type ModifyClusterParams struct {

	/*Body*/
	Body *models_autoscale.ClusterRequestJSON
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify cluster params
func (o *ModifyClusterParams) WithTimeout(timeout time.Duration) *ModifyClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify cluster params
func (o *ModifyClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify cluster params
func (o *ModifyClusterParams) WithContext(ctx context.Context) *ModifyClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify cluster params
func (o *ModifyClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify cluster params
func (o *ModifyClusterParams) WithHTTPClient(client *http.Client) *ModifyClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify cluster params
func (o *ModifyClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify cluster params
func (o *ModifyClusterParams) WithBody(body *models_autoscale.ClusterRequestJSON) *ModifyClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify cluster params
func (o *ModifyClusterParams) SetBody(body *models_autoscale.ClusterRequestJSON) {
	o.Body = body
}

// WithClusterID adds the clusterID to the modify cluster params
func (o *ModifyClusterParams) WithClusterID(clusterID int64) *ModifyClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the modify cluster params
func (o *ModifyClusterParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ClusterRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
