// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateStackV2Reader is a Reader for the DeletePrivateStackV2 structure.
type DeletePrivateStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateStackV2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateStackV2Default creates a DeletePrivateStackV2Default with default headers values
func NewDeletePrivateStackV2Default(code int) *DeletePrivateStackV2Default {
	return &DeletePrivateStackV2Default{
		_statusCode: code,
	}
}

/*DeletePrivateStackV2Default handles this case with default header values.

successful operation
*/
type DeletePrivateStackV2Default struct {
	_statusCode int
}

// Code gets the status code for the delete private stack v2 default response
func (o *DeletePrivateStackV2Default) Code() int {
	return o._statusCode
}

func (o *DeletePrivateStackV2Default) Error() string {
	return fmt.Sprintf("[DELETE /v2/stacks/user/{name}][%d] deletePrivateStackV2 default ", o._statusCode)
}

func (o *DeletePrivateStackV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
