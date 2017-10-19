// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivatesLdapReader is a Reader for the GetPrivatesLdap structure.
type GetPrivatesLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesLdapOK creates a GetPrivatesLdapOK with default headers values
func NewGetPrivatesLdapOK() *GetPrivatesLdapOK {
	return &GetPrivatesLdapOK{}
}

/*GetPrivatesLdapOK handles this case with default header values.

successful operation
*/
type GetPrivatesLdapOK struct {
	Payload []*models_cloudbreak.LdapConfigResponse
}

func (o *GetPrivatesLdapOK) Error() string {
	return fmt.Sprintf("[GET /ldap/user][%d] getPrivatesLdapOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
