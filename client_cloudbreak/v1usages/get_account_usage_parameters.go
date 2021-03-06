// Code generated by go-swagger; DO NOT EDIT.

package v1usages

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

// NewGetAccountUsageParams creates a new GetAccountUsageParams object
// with the default values initialized.
func NewGetAccountUsageParams() *GetAccountUsageParams {
	var ()
	return &GetAccountUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountUsageParamsWithTimeout creates a new GetAccountUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountUsageParamsWithTimeout(timeout time.Duration) *GetAccountUsageParams {
	var ()
	return &GetAccountUsageParams{

		timeout: timeout,
	}
}

// NewGetAccountUsageParamsWithContext creates a new GetAccountUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountUsageParamsWithContext(ctx context.Context) *GetAccountUsageParams {
	var ()
	return &GetAccountUsageParams{

		Context: ctx,
	}
}

// NewGetAccountUsageParamsWithHTTPClient creates a new GetAccountUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountUsageParamsWithHTTPClient(client *http.Client) *GetAccountUsageParams {
	var ()
	return &GetAccountUsageParams{
		HTTPClient: client,
	}
}

/*GetAccountUsageParams contains all the parameters to send to the API endpoint
for the get account usage operation typically these are written to a http.Request
*/
type GetAccountUsageParams struct {

	/*Cloud*/
	Cloud *string
	/*Filterenddate*/
	Filterenddate *int64
	/*Since*/
	Since *int64
	/*User*/
	User *string
	/*Zone*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account usage params
func (o *GetAccountUsageParams) WithTimeout(timeout time.Duration) *GetAccountUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account usage params
func (o *GetAccountUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account usage params
func (o *GetAccountUsageParams) WithContext(ctx context.Context) *GetAccountUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account usage params
func (o *GetAccountUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account usage params
func (o *GetAccountUsageParams) WithHTTPClient(client *http.Client) *GetAccountUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account usage params
func (o *GetAccountUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloud adds the cloud to the get account usage params
func (o *GetAccountUsageParams) WithCloud(cloud *string) *GetAccountUsageParams {
	o.SetCloud(cloud)
	return o
}

// SetCloud adds the cloud to the get account usage params
func (o *GetAccountUsageParams) SetCloud(cloud *string) {
	o.Cloud = cloud
}

// WithFilterenddate adds the filterenddate to the get account usage params
func (o *GetAccountUsageParams) WithFilterenddate(filterenddate *int64) *GetAccountUsageParams {
	o.SetFilterenddate(filterenddate)
	return o
}

// SetFilterenddate adds the filterenddate to the get account usage params
func (o *GetAccountUsageParams) SetFilterenddate(filterenddate *int64) {
	o.Filterenddate = filterenddate
}

// WithSince adds the since to the get account usage params
func (o *GetAccountUsageParams) WithSince(since *int64) *GetAccountUsageParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get account usage params
func (o *GetAccountUsageParams) SetSince(since *int64) {
	o.Since = since
}

// WithUser adds the user to the get account usage params
func (o *GetAccountUsageParams) WithUser(user *string) *GetAccountUsageParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get account usage params
func (o *GetAccountUsageParams) SetUser(user *string) {
	o.User = user
}

// WithZone adds the zone to the get account usage params
func (o *GetAccountUsageParams) WithZone(zone *string) *GetAccountUsageParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get account usage params
func (o *GetAccountUsageParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cloud != nil {

		// query param cloud
		var qrCloud string
		if o.Cloud != nil {
			qrCloud = *o.Cloud
		}
		qCloud := qrCloud
		if qCloud != "" {
			if err := r.SetQueryParam("cloud", qCloud); err != nil {
				return err
			}
		}

	}

	if o.Filterenddate != nil {

		// query param filterenddate
		var qrFilterenddate int64
		if o.Filterenddate != nil {
			qrFilterenddate = *o.Filterenddate
		}
		qFilterenddate := swag.FormatInt64(qrFilterenddate)
		if qFilterenddate != "" {
			if err := r.SetQueryParam("filterenddate", qFilterenddate); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}

	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
