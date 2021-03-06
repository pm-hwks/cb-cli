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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteStackV2Params creates a new DeleteStackV2Params object
// with the default values initialized.
func NewDeleteStackV2Params() *DeleteStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStackV2ParamsWithTimeout creates a new DeleteStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStackV2ParamsWithTimeout(timeout time.Duration) *DeleteStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteStackV2ParamsWithContext creates a new DeleteStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStackV2ParamsWithContext(ctx context.Context) *DeleteStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteStackV2ParamsWithHTTPClient creates a new DeleteStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStackV2ParamsWithHTTPClient(client *http.Client) *DeleteStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,
		HTTPClient:         client,
	}
}

/*DeleteStackV2Params contains all the parameters to send to the API endpoint
for the delete stack v2 operation typically these are written to a http.Request
*/
type DeleteStackV2Params struct {

	/*DeleteDependencies*/
	DeleteDependencies *bool
	/*Forced*/
	Forced *bool
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete stack v2 params
func (o *DeleteStackV2Params) WithTimeout(timeout time.Duration) *DeleteStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete stack v2 params
func (o *DeleteStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete stack v2 params
func (o *DeleteStackV2Params) WithContext(ctx context.Context) *DeleteStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete stack v2 params
func (o *DeleteStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete stack v2 params
func (o *DeleteStackV2Params) WithHTTPClient(client *http.Client) *DeleteStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete stack v2 params
func (o *DeleteStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteDependencies adds the deleteDependencies to the delete stack v2 params
func (o *DeleteStackV2Params) WithDeleteDependencies(deleteDependencies *bool) *DeleteStackV2Params {
	o.SetDeleteDependencies(deleteDependencies)
	return o
}

// SetDeleteDependencies adds the deleteDependencies to the delete stack v2 params
func (o *DeleteStackV2Params) SetDeleteDependencies(deleteDependencies *bool) {
	o.DeleteDependencies = deleteDependencies
}

// WithForced adds the forced to the delete stack v2 params
func (o *DeleteStackV2Params) WithForced(forced *bool) *DeleteStackV2Params {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete stack v2 params
func (o *DeleteStackV2Params) SetForced(forced *bool) {
	o.Forced = forced
}

// WithID adds the id to the delete stack v2 params
func (o *DeleteStackV2Params) WithID(id int64) *DeleteStackV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete stack v2 params
func (o *DeleteStackV2Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteDependencies != nil {

		// query param deleteDependencies
		var qrDeleteDependencies bool
		if o.DeleteDependencies != nil {
			qrDeleteDependencies = *o.DeleteDependencies
		}
		qDeleteDependencies := swag.FormatBool(qrDeleteDependencies)
		if qDeleteDependencies != "" {
			if err := r.SetQueryParam("deleteDependencies", qDeleteDependencies); err != nil {
				return err
			}
		}

	}

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

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
