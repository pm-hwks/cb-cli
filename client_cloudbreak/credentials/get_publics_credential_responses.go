// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicsCredentialReader is a Reader for the GetPublicsCredential structure.
type GetPublicsCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsCredentialOK creates a GetPublicsCredentialOK with default headers values
func NewGetPublicsCredentialOK() *GetPublicsCredentialOK {
	return &GetPublicsCredentialOK{}
}

/*GetPublicsCredentialOK handles this case with default header values.

successful operation
*/
type GetPublicsCredentialOK struct {
	Payload []*models_cloudbreak.CredentialResponse
}

func (o *GetPublicsCredentialOK) Error() string {
	return fmt.Sprintf("[GET /credentials/account][%d] getPublicsCredentialOK  %+v", 200, o.Payload)
}

func (o *GetPublicsCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
