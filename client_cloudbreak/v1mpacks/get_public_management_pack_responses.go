// Code generated by go-swagger; DO NOT EDIT.

package v1mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicManagementPackReader is a Reader for the GetPublicManagementPack structure.
type GetPublicManagementPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicManagementPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicManagementPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicManagementPackOK creates a GetPublicManagementPackOK with default headers values
func NewGetPublicManagementPackOK() *GetPublicManagementPackOK {
	return &GetPublicManagementPackOK{}
}

/*GetPublicManagementPackOK handles this case with default header values.

successful operation
*/
type GetPublicManagementPackOK struct {
	Payload *models_cloudbreak.ManagementPackResponse
}

func (o *GetPublicManagementPackOK) Error() string {
	return fmt.Sprintf("[GET /v1/mpacks/account/{name}][%d] getPublicManagementPackOK  %+v", 200, o.Payload)
}

func (o *GetPublicManagementPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ManagementPackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
