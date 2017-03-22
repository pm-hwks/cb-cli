package blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostPrivateBlueprintParams creates a new PostPrivateBlueprintParams object
// with the default values initialized.
func NewPostPrivateBlueprintParams() *PostPrivateBlueprintParams {
	var ()
	return &PostPrivateBlueprintParams{}
}

/*PostPrivateBlueprintParams contains all the parameters to send to the API endpoint
for the post private blueprint operation typically these are written to a http.Request
*/
type PostPrivateBlueprintParams struct {

	/*Body*/
	Body *models_cloudbreak.BlueprintRequest
}

// WithBody adds the body to the post private blueprint params
func (o *PostPrivateBlueprintParams) WithBody(body *models_cloudbreak.BlueprintRequest) *PostPrivateBlueprintParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateBlueprintParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.BlueprintRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}