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

// AddOrganizationUsersReader is a Reader for the AddOrganizationUsers structure.
type AddOrganizationUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrganizationUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddOrganizationUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddOrganizationUsersOK creates a AddOrganizationUsersOK with default headers values
func NewAddOrganizationUsersOK() *AddOrganizationUsersOK {
	return &AddOrganizationUsersOK{}
}

/*AddOrganizationUsersOK handles this case with default header values.

successful operation
*/
type AddOrganizationUsersOK struct {
	Payload []*models_cloudbreak.UserResponseJSON
}

func (o *AddOrganizationUsersOK) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/name/{name}/addUsers][%d] addOrganizationUsersOK  %+v", 200, o.Payload)
}

func (o *AddOrganizationUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}