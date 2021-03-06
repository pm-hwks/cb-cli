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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPostPrivateRecipeParams creates a new PostPrivateRecipeParams object
// with the default values initialized.
func NewPostPrivateRecipeParams() *PostPrivateRecipeParams {
	var ()
	return &PostPrivateRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateRecipeParamsWithTimeout creates a new PostPrivateRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateRecipeParamsWithTimeout(timeout time.Duration) *PostPrivateRecipeParams {
	var ()
	return &PostPrivateRecipeParams{

		timeout: timeout,
	}
}

// NewPostPrivateRecipeParamsWithContext creates a new PostPrivateRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateRecipeParamsWithContext(ctx context.Context) *PostPrivateRecipeParams {
	var ()
	return &PostPrivateRecipeParams{

		Context: ctx,
	}
}

// NewPostPrivateRecipeParamsWithHTTPClient creates a new PostPrivateRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateRecipeParamsWithHTTPClient(client *http.Client) *PostPrivateRecipeParams {
	var ()
	return &PostPrivateRecipeParams{
		HTTPClient: client,
	}
}

/*PostPrivateRecipeParams contains all the parameters to send to the API endpoint
for the post private recipe operation typically these are written to a http.Request
*/
type PostPrivateRecipeParams struct {

	/*Body*/
	Body *models_cloudbreak.RecipeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private recipe params
func (o *PostPrivateRecipeParams) WithTimeout(timeout time.Duration) *PostPrivateRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private recipe params
func (o *PostPrivateRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private recipe params
func (o *PostPrivateRecipeParams) WithContext(ctx context.Context) *PostPrivateRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private recipe params
func (o *PostPrivateRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private recipe params
func (o *PostPrivateRecipeParams) WithHTTPClient(client *http.Client) *PostPrivateRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private recipe params
func (o *PostPrivateRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private recipe params
func (o *PostPrivateRecipeParams) WithBody(body *models_cloudbreak.RecipeRequest) *PostPrivateRecipeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private recipe params
func (o *PostPrivateRecipeParams) SetBody(body *models_cloudbreak.RecipeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.RecipeRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
