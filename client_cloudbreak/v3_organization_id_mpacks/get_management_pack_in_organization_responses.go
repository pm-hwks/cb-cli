// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetManagementPackInOrganizationReader is a Reader for the GetManagementPackInOrganization structure.
type GetManagementPackInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagementPackInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetManagementPackInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagementPackInOrganizationOK creates a GetManagementPackInOrganizationOK with default headers values
func NewGetManagementPackInOrganizationOK() *GetManagementPackInOrganizationOK {
	return &GetManagementPackInOrganizationOK{}
}

/*GetManagementPackInOrganizationOK handles this case with default header values.

successful operation
*/
type GetManagementPackInOrganizationOK struct {
	Payload *models_cloudbreak.ManagementPackResponse
}

func (o *GetManagementPackInOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v3/{organizationId}/mpacks/{name}][%d] getManagementPackInOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetManagementPackInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ManagementPackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
