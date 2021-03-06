// Code generated by go-swagger; DO NOT EDIT.

package v1flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePublicFlexSubscriptionByNameReader is a Reader for the DeletePublicFlexSubscriptionByName structure.
type DeletePublicFlexSubscriptionByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePublicFlexSubscriptionByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePublicFlexSubscriptionByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePublicFlexSubscriptionByNameDefault creates a DeletePublicFlexSubscriptionByNameDefault with default headers values
func NewDeletePublicFlexSubscriptionByNameDefault(code int) *DeletePublicFlexSubscriptionByNameDefault {
	return &DeletePublicFlexSubscriptionByNameDefault{
		_statusCode: code,
	}
}

/*DeletePublicFlexSubscriptionByNameDefault handles this case with default header values.

successful operation
*/
type DeletePublicFlexSubscriptionByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete public flex subscription by name default response
func (o *DeletePublicFlexSubscriptionByNameDefault) Code() int {
	return o._statusCode
}

func (o *DeletePublicFlexSubscriptionByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/flexsubscriptions/account/{name}][%d] deletePublicFlexSubscriptionByName default ", o._statusCode)
}

func (o *DeletePublicFlexSubscriptionByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
