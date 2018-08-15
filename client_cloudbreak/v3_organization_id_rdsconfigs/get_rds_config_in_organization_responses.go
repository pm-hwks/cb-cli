// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetRdsConfigInOrganizationReader is a Reader for the GetRdsConfigInOrganization structure.
type GetRdsConfigInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRdsConfigInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRdsConfigInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRdsConfigInOrganizationOK creates a GetRdsConfigInOrganizationOK with default headers values
func NewGetRdsConfigInOrganizationOK() *GetRdsConfigInOrganizationOK {
	return &GetRdsConfigInOrganizationOK{}
}

/*GetRdsConfigInOrganizationOK handles this case with default header values.

successful operation
*/
type GetRdsConfigInOrganizationOK struct {
	Payload *models_cloudbreak.RDSConfigResponse
}

func (o *GetRdsConfigInOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v3/{organizationId}/rdsconfigs/{name}][%d] getRdsConfigInOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetRdsConfigInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RDSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
