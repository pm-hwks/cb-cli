// Code generated by go-swagger; DO NOT EDIT.

package v1organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetOrganizationReader is a Reader for the GetOrganization structure.
type GetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationOK creates a GetOrganizationOK with default headers values
func NewGetOrganizationOK() *GetOrganizationOK {
	return &GetOrganizationOK{}
}

/*GetOrganizationOK handles this case with default header values.

successful operation
*/
type GetOrganizationOK struct {
	Payload *models_cloudbreak.OrganizationResponse
}

func (o *GetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] getOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.OrganizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
