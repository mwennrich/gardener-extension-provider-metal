// Code generated by go-swagger; DO NOT EDIT.

package partition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// UpdatePartitionReader is a Reader for the UpdatePartition structure.
type UpdatePartitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePartitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePartitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewUpdatePartitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdatePartitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePartitionOK creates a UpdatePartitionOK with default headers values
func NewUpdatePartitionOK() *UpdatePartitionOK {
	return &UpdatePartitionOK{}
}

/*UpdatePartitionOK handles this case with default header values.

OK
*/
type UpdatePartitionOK struct {
	Payload *models.V1PartitionResponse
}

func (o *UpdatePartitionOK) Error() string {
	return fmt.Sprintf("[POST /v1/partition][%d] updatePartitionOK  %+v", 200, o.Payload)
}

func (o *UpdatePartitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PartitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePartitionConflict creates a UpdatePartitionConflict with default headers values
func NewUpdatePartitionConflict() *UpdatePartitionConflict {
	return &UpdatePartitionConflict{}
}

/*UpdatePartitionConflict handles this case with default header values.

Conflict
*/
type UpdatePartitionConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *UpdatePartitionConflict) Error() string {
	return fmt.Sprintf("[POST /v1/partition][%d] updatePartitionConflict  %+v", 409, o.Payload)
}

func (o *UpdatePartitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePartitionDefault creates a UpdatePartitionDefault with default headers values
func NewUpdatePartitionDefault(code int) *UpdatePartitionDefault {
	return &UpdatePartitionDefault{
		_statusCode: code,
	}
}

/*UpdatePartitionDefault handles this case with default header values.

Error
*/
type UpdatePartitionDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the update partition default response
func (o *UpdatePartitionDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePartitionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/partition][%d] updatePartition default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePartitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
