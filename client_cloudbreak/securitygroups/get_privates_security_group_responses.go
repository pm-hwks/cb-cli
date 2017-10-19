// Code generated by go-swagger; DO NOT EDIT.

package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivatesSecurityGroupReader is a Reader for the GetPrivatesSecurityGroup structure.
type GetPrivatesSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesSecurityGroupOK creates a GetPrivatesSecurityGroupOK with default headers values
func NewGetPrivatesSecurityGroupOK() *GetPrivatesSecurityGroupOK {
	return &GetPrivatesSecurityGroupOK{}
}

/*GetPrivatesSecurityGroupOK handles this case with default header values.

successful operation
*/
type GetPrivatesSecurityGroupOK struct {
	Payload []*models_cloudbreak.SecurityGroupResponse
}

func (o *GetPrivatesSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /securitygroups/user][%d] getPrivatesSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
