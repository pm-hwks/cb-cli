// Code generated by go-swagger; DO NOT EDIT.

package smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivateSmartSenseSubscriptionsReader is a Reader for the GetPrivateSmartSenseSubscriptions structure.
type GetPrivateSmartSenseSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSmartSenseSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSmartSenseSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSmartSenseSubscriptionsOK creates a GetPrivateSmartSenseSubscriptionsOK with default headers values
func NewGetPrivateSmartSenseSubscriptionsOK() *GetPrivateSmartSenseSubscriptionsOK {
	return &GetPrivateSmartSenseSubscriptionsOK{}
}

/*GetPrivateSmartSenseSubscriptionsOK handles this case with default header values.

successful operation
*/
type GetPrivateSmartSenseSubscriptionsOK struct {
	Payload []*models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *GetPrivateSmartSenseSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /smartsensesubscriptions/user][%d] getPrivateSmartSenseSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSmartSenseSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
