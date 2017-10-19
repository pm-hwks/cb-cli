// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_autoscale"
)

// AddClusterReader is a Reader for the AddCluster structure.
type AddClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddClusterOK creates a AddClusterOK with default headers values
func NewAddClusterOK() *AddClusterOK {
	return &AddClusterOK{}
}

/*AddClusterOK handles this case with default header values.

successful operation
*/
type AddClusterOK struct {
	Payload *models_autoscale.ClusterSummary
}

func (o *AddClusterOK) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] addClusterOK  %+v", 200, o.Payload)
}

func (o *AddClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.ClusterSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
