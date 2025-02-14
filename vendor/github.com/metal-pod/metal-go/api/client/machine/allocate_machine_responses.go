// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// AllocateMachineReader is a Reader for the AllocateMachine structure.
type AllocateMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllocateMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAllocateMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllocateMachineOK creates a AllocateMachineOK with default headers values
func NewAllocateMachineOK() *AllocateMachineOK {
	return &AllocateMachineOK{}
}

/*AllocateMachineOK handles this case with default header values.

OK
*/
type AllocateMachineOK struct {
	Payload *models.V1MachineResponse
}

func (o *AllocateMachineOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachineOK  %+v", 200, o.Payload)
}

func (o *AllocateMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateMachineDefault creates a AllocateMachineDefault with default headers values
func NewAllocateMachineDefault(code int) *AllocateMachineDefault {
	return &AllocateMachineDefault{
		_statusCode: code,
	}
}

/*AllocateMachineDefault handles this case with default header values.

Error
*/
type AllocateMachineDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the allocate machine default response
func (o *AllocateMachineDefault) Code() int {
	return o._statusCode
}

func (o *AllocateMachineDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachine default  %+v", o._statusCode, o.Payload)
}

func (o *AllocateMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
