// Code generated by go-swagger; DO NOT EDIT.

package v1organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeOrganizationUsersReader is a Reader for the ChangeOrganizationUsers structure.
type ChangeOrganizationUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeOrganizationUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewChangeOrganizationUsersDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewChangeOrganizationUsersDefault creates a ChangeOrganizationUsersDefault with default headers values
func NewChangeOrganizationUsersDefault(code int) *ChangeOrganizationUsersDefault {
	return &ChangeOrganizationUsersDefault{
		_statusCode: code,
	}
}

/*ChangeOrganizationUsersDefault handles this case with default header values.

successful operation
*/
type ChangeOrganizationUsersDefault struct {
	_statusCode int
}

// Code gets the status code for the change organization users default response
func (o *ChangeOrganizationUsersDefault) Code() int {
	return o._statusCode
}

func (o *ChangeOrganizationUsersDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] changeOrganizationUsers default ", o._statusCode)
}

func (o *ChangeOrganizationUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
