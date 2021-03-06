// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetDisktypesReader is a Reader for the GetDisktypes structure.
type GetDisktypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDisktypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDisktypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDisktypesOK creates a GetDisktypesOK with default headers values
func NewGetDisktypesOK() *GetDisktypesOK {
	return &GetDisktypesOK{}
}

/*GetDisktypesOK handles this case with default header values.

successful operation
*/
type GetDisktypesOK struct {
	Payload *models_cloudbreak.PlatformDisksJSON
}

func (o *GetDisktypesOK) Error() string {
	return fmt.Sprintf("[GET /v1/connectors/disktypes][%d] getDisktypesOK  %+v", 200, o.Payload)
}

func (o *GetDisktypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformDisksJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
