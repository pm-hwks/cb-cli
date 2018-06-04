// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteInstancesStackV2Reader is a Reader for the DeleteInstancesStackV2 structure.
type DeleteInstancesStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstancesStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteInstancesStackV2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteInstancesStackV2Default creates a DeleteInstancesStackV2Default with default headers values
func NewDeleteInstancesStackV2Default(code int) *DeleteInstancesStackV2Default {
	return &DeleteInstancesStackV2Default{
		_statusCode: code,
	}
}

/*DeleteInstancesStackV2Default handles this case with default header values.

successful operation
*/
type DeleteInstancesStackV2Default struct {
	_statusCode int
}

// Code gets the status code for the delete instances stack v2 default response
func (o *DeleteInstancesStackV2Default) Code() int {
	return o._statusCode
}

func (o *DeleteInstancesStackV2Default) Error() string {
	return fmt.Sprintf("[DELETE /v2/stacks/{stackId}/deleteInstances][%d] deleteInstancesStackV2 default ", o._statusCode)
}

func (o *DeleteInstancesStackV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
