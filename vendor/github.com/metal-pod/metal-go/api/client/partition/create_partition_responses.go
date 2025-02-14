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

// CreatePartitionReader is a Reader for the CreatePartition structure.
type CreatePartitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePartitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePartitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreatePartitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreatePartitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePartitionCreated creates a CreatePartitionCreated with default headers values
func NewCreatePartitionCreated() *CreatePartitionCreated {
	return &CreatePartitionCreated{}
}

/*CreatePartitionCreated handles this case with default header values.

Created
*/
type CreatePartitionCreated struct {
	Payload *models.V1PartitionResponse
}

func (o *CreatePartitionCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/partition][%d] createPartitionCreated  %+v", 201, o.Payload)
}

func (o *CreatePartitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PartitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePartitionConflict creates a CreatePartitionConflict with default headers values
func NewCreatePartitionConflict() *CreatePartitionConflict {
	return &CreatePartitionConflict{}
}

/*CreatePartitionConflict handles this case with default header values.

Conflict
*/
type CreatePartitionConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *CreatePartitionConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/partition][%d] createPartitionConflict  %+v", 409, o.Payload)
}

func (o *CreatePartitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePartitionDefault creates a CreatePartitionDefault with default headers values
func NewCreatePartitionDefault(code int) *CreatePartitionDefault {
	return &CreatePartitionDefault{
		_statusCode: code,
	}
}

/*CreatePartitionDefault handles this case with default header values.

Error
*/
type CreatePartitionDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the create partition default response
func (o *CreatePartitionDefault) Code() int {
	return o._statusCode
}

func (o *CreatePartitionDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/partition][%d] createPartition default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePartitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
