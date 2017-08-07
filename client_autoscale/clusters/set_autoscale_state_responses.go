package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// SetAutoscaleStateReader is a Reader for the SetAutoscaleState structure.
type SetAutoscaleStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SetAutoscaleStateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAutoscaleStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAutoscaleStateOK creates a SetAutoscaleStateOK with default headers values
func NewSetAutoscaleStateOK() *SetAutoscaleStateOK {
	return &SetAutoscaleStateOK{}
}

/*SetAutoscaleStateOK handles this case with default header values.

successful operation
*/
type SetAutoscaleStateOK struct {
	Payload *models_autoscale.ClusterSummary
}

func (o *SetAutoscaleStateOK) Error() string {
	return fmt.Sprintf("[POST /clusters/{clusterId}/autoscale][%d] setAutoscaleStateOK  %+v", 200, o.Payload)
}

func (o *SetAutoscaleStateOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.ClusterSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
