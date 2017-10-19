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

// GetClustersReader is a Reader for the GetClusters structure.
type GetClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClustersOK creates a GetClustersOK with default headers values
func NewGetClustersOK() *GetClustersOK {
	return &GetClustersOK{}
}

/*GetClustersOK handles this case with default header values.

successful operation
*/
type GetClustersOK struct {
	Payload []*models_autoscale.ClusterSummary
}

func (o *GetClustersOK) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] getClustersOK  %+v", 200, o.Payload)
}

func (o *GetClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
