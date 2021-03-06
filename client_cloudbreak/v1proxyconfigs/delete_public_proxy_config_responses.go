// Code generated by go-swagger; DO NOT EDIT.

package v1proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePublicProxyConfigReader is a Reader for the DeletePublicProxyConfig structure.
type DeletePublicProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePublicProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePublicProxyConfigDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePublicProxyConfigDefault creates a DeletePublicProxyConfigDefault with default headers values
func NewDeletePublicProxyConfigDefault(code int) *DeletePublicProxyConfigDefault {
	return &DeletePublicProxyConfigDefault{
		_statusCode: code,
	}
}

/*DeletePublicProxyConfigDefault handles this case with default header values.

successful operation
*/
type DeletePublicProxyConfigDefault struct {
	_statusCode int
}

// Code gets the status code for the delete public proxy config default response
func (o *DeletePublicProxyConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeletePublicProxyConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/proxyconfigs/account/{name}][%d] deletePublicProxyConfig default ", o._statusCode)
}

func (o *DeletePublicProxyConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
