// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// FindSwitchReader is a Reader for the FindSwitch structure.
type FindSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindSwitchOK creates a FindSwitchOK with default headers values
func NewFindSwitchOK() *FindSwitchOK {
	return &FindSwitchOK{}
}

/*FindSwitchOK handles this case with default header values.

OK
*/
type FindSwitchOK struct {
	Payload *models.V1SwitchResponse
}

func (o *FindSwitchOK) Error() string {
	return fmt.Sprintf("[GET /v1/switch/{id}][%d] findSwitchOK  %+v", 200, o.Payload)
}

func (o *FindSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindSwitchDefault creates a FindSwitchDefault with default headers values
func NewFindSwitchDefault(code int) *FindSwitchDefault {
	return &FindSwitchDefault{
		_statusCode: code,
	}
}

/*FindSwitchDefault handles this case with default header values.

Error
*/
type FindSwitchDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the find switch default response
func (o *FindSwitchDefault) Code() int {
	return o._statusCode
}

func (o *FindSwitchDefault) Error() string {
	return fmt.Sprintf("[GET /v1/switch/{id}][%d] findSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *FindSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
