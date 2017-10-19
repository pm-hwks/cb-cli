// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewUpdatePrometheusAlertParams creates a new UpdatePrometheusAlertParams object
// with the default values initialized.
func NewUpdatePrometheusAlertParams() *UpdatePrometheusAlertParams {
	var ()
	return &UpdatePrometheusAlertParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrometheusAlertParamsWithTimeout creates a new UpdatePrometheusAlertParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePrometheusAlertParamsWithTimeout(timeout time.Duration) *UpdatePrometheusAlertParams {
	var ()
	return &UpdatePrometheusAlertParams{

		timeout: timeout,
	}
}

// NewUpdatePrometheusAlertParamsWithContext creates a new UpdatePrometheusAlertParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePrometheusAlertParamsWithContext(ctx context.Context) *UpdatePrometheusAlertParams {
	var ()
	return &UpdatePrometheusAlertParams{

		Context: ctx,
	}
}

// NewUpdatePrometheusAlertParamsWithHTTPClient creates a new UpdatePrometheusAlertParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePrometheusAlertParamsWithHTTPClient(client *http.Client) *UpdatePrometheusAlertParams {
	var ()
	return &UpdatePrometheusAlertParams{
		HTTPClient: client,
	}
}

/*UpdatePrometheusAlertParams contains all the parameters to send to the API endpoint
for the update prometheus alert operation typically these are written to a http.Request
*/
type UpdatePrometheusAlertParams struct {

	/*AlertID*/
	AlertID int64
	/*Body*/
	Body *models_autoscale.PromhetheusAlert
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithTimeout(timeout time.Duration) *UpdatePrometheusAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithContext(ctx context.Context) *UpdatePrometheusAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithHTTPClient(client *http.Client) *UpdatePrometheusAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithAlertID(alertID int64) *UpdatePrometheusAlertParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetAlertID(alertID int64) {
	o.AlertID = alertID
}

// WithBody adds the body to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithBody(body *models_autoscale.PromhetheusAlert) *UpdatePrometheusAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetBody(body *models_autoscale.PromhetheusAlert) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) WithClusterID(clusterID int64) *UpdatePrometheusAlertParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update prometheus alert params
func (o *UpdatePrometheusAlertParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrometheusAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", swag.FormatInt64(o.AlertID)); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models_autoscale.PromhetheusAlert)
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
