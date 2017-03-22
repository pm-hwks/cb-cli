package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostPublicStackParams creates a new PostPublicStackParams object
// with the default values initialized.
func NewPostPublicStackParams() *PostPublicStackParams {
	var ()
	return &PostPublicStackParams{}
}

/*PostPublicStackParams contains all the parameters to send to the API endpoint
for the post public stack operation typically these are written to a http.Request
*/
type PostPublicStackParams struct {

	/*Body*/
	Body *models_cloudbreak.StackRequest
}

// WithBody adds the body to the post public stack params
func (o *PostPublicStackParams) WithBody(body *models_cloudbreak.StackRequest) *PostPublicStackParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicStackParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.StackRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}