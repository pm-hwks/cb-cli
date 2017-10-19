// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPlatformVariantsReader is a Reader for the GetPlatformVariants structure.
type GetPlatformVariantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformVariantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformVariantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlatformVariantsOK creates a GetPlatformVariantsOK with default headers values
func NewGetPlatformVariantsOK() *GetPlatformVariantsOK {
	return &GetPlatformVariantsOK{}
}

/*GetPlatformVariantsOK handles this case with default header values.

successful operation
*/
type GetPlatformVariantsOK struct {
	Payload *models_cloudbreak.PlatformVariantsJSON
}

func (o *GetPlatformVariantsOK) Error() string {
	return fmt.Sprintf("[GET /connectors/variants][%d] getPlatformVariantsOK  %+v", 200, o.Payload)
}

func (o *GetPlatformVariantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformVariantsJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
