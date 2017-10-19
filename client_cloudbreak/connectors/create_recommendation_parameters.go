// Code generated by go-swagger; DO NOT EDIT.

package connectors

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewCreateRecommendationParams creates a new CreateRecommendationParams object
// with the default values initialized.
func NewCreateRecommendationParams() *CreateRecommendationParams {
	var ()
	return &CreateRecommendationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRecommendationParamsWithTimeout creates a new CreateRecommendationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRecommendationParamsWithTimeout(timeout time.Duration) *CreateRecommendationParams {
	var ()
	return &CreateRecommendationParams{

		timeout: timeout,
	}
}

// NewCreateRecommendationParamsWithContext creates a new CreateRecommendationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRecommendationParamsWithContext(ctx context.Context) *CreateRecommendationParams {
	var ()
	return &CreateRecommendationParams{

		Context: ctx,
	}
}

// NewCreateRecommendationParamsWithHTTPClient creates a new CreateRecommendationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRecommendationParamsWithHTTPClient(client *http.Client) *CreateRecommendationParams {
	var ()
	return &CreateRecommendationParams{
		HTTPClient: client,
	}
}

/*CreateRecommendationParams contains all the parameters to send to the API endpoint
for the create recommendation operation typically these are written to a http.Request
*/
type CreateRecommendationParams struct {

	/*Body*/
	Body *models_cloudbreak.RecommendationRequestJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create recommendation params
func (o *CreateRecommendationParams) WithTimeout(timeout time.Duration) *CreateRecommendationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create recommendation params
func (o *CreateRecommendationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create recommendation params
func (o *CreateRecommendationParams) WithContext(ctx context.Context) *CreateRecommendationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create recommendation params
func (o *CreateRecommendationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create recommendation params
func (o *CreateRecommendationParams) WithHTTPClient(client *http.Client) *CreateRecommendationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create recommendation params
func (o *CreateRecommendationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create recommendation params
func (o *CreateRecommendationParams) WithBody(body *models_cloudbreak.RecommendationRequestJSON) *CreateRecommendationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create recommendation params
func (o *CreateRecommendationParams) SetBody(body *models_cloudbreak.RecommendationRequestJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRecommendationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.RecommendationRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
