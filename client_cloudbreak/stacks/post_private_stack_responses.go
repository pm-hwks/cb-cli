// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPrivateStackReader is a Reader for the PostPrivateStack structure.
type PostPrivateStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateStackOK creates a PostPrivateStackOK with default headers values
func NewPostPrivateStackOK() *PostPrivateStackOK {
	return &PostPrivateStackOK{}
}

/*PostPrivateStackOK handles this case with default header values.

successful operation
*/
type PostPrivateStackOK struct {
	Payload *models_cloudbreak.StackResponse
}

func (o *PostPrivateStackOK) Error() string {
	return fmt.Sprintf("[POST /stacks/user][%d] postPrivateStackOK  %+v", 200, o.Payload)
}

func (o *PostPrivateStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
