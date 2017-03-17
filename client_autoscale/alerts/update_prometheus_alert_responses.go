package alerts

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

// UpdatePrometheusAlertReader is a Reader for the UpdatePrometheusAlert structure.
type UpdatePrometheusAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdatePrometheusAlertReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePrometheusAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePrometheusAlertOK creates a UpdatePrometheusAlertOK with default headers values
func NewUpdatePrometheusAlertOK() *UpdatePrometheusAlertOK {
	return &UpdatePrometheusAlertOK{}
}

/*UpdatePrometheusAlertOK handles this case with default header values.

successful operation
*/
type UpdatePrometheusAlertOK struct {
	Payload *models_autoscale.PromhetheusAlert
}

func (o *UpdatePrometheusAlertOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/{clusterId}/alerts/prometheus/{alertId}][%d] updatePrometheusAlertOK  %+v", 200, o.Payload)
}

func (o *UpdatePrometheusAlertOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.PromhetheusAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}