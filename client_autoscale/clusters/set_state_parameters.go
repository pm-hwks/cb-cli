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

// NewSetStateParams creates a new SetStateParams object
// with the default values initialized.
func NewSetStateParams() *SetStateParams {
	var ()
	return &SetStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetStateParamsWithTimeout creates a new SetStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetStateParamsWithTimeout(timeout time.Duration) *SetStateParams {
	var ()
	return &SetStateParams{

		timeout: timeout,
	}
}

// NewSetStateParamsWithContext creates a new SetStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetStateParamsWithContext(ctx context.Context) *SetStateParams {
	var ()
	return &SetStateParams{

		Context: ctx,
	}
}

// NewSetStateParamsWithHTTPClient creates a new SetStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetStateParamsWithHTTPClient(client *http.Client) *SetStateParams {
	var ()
	return &SetStateParams{
		HTTPClient: client,
	}
}

/*SetStateParams contains all the parameters to send to the API endpoint
for the set state operation typically these are written to a http.Request
*/
type SetStateParams struct {

	/*Body*/
	Body *models_autoscale.ClusterState
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set state params
func (o *SetStateParams) WithTimeout(timeout time.Duration) *SetStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set state params
func (o *SetStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set state params
func (o *SetStateParams) WithContext(ctx context.Context) *SetStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set state params
func (o *SetStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set state params
func (o *SetStateParams) WithHTTPClient(client *http.Client) *SetStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set state params
func (o *SetStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set state params
func (o *SetStateParams) WithBody(body *models_autoscale.ClusterState) *SetStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set state params
func (o *SetStateParams) SetBody(body *models_autoscale.ClusterState) {
	o.Body = body
}

// WithClusterID adds the clusterID to the set state params
func (o *SetStateParams) WithClusterID(clusterID int64) *SetStateParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set state params
func (o *SetStateParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *SetStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ClusterState)
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
