package clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteClustertemplatesAccountNameParams creates a new DeleteClustertemplatesAccountNameParams object
// with the default values initialized.
func NewDeleteClustertemplatesAccountNameParams() *DeleteClustertemplatesAccountNameParams {
	var ()
	return &DeleteClustertemplatesAccountNameParams{}
}

/*DeleteClustertemplatesAccountNameParams contains all the parameters to send to the API endpoint
for the delete clustertemplates account name operation typically these are written to a http.Request
*/
type DeleteClustertemplatesAccountNameParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the delete clustertemplates account name params
func (o *DeleteClustertemplatesAccountNameParams) WithName(name string) *DeleteClustertemplatesAccountNameParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClustertemplatesAccountNameParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}