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

// DeletePartitionReader is a Reader for the DeletePartition structure.
type DeletePartitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePartitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePartitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeletePartitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePartitionOK creates a DeletePartitionOK with default headers values
func NewDeletePartitionOK() *DeletePartitionOK {
	return &DeletePartitionOK{}
}

/*DeletePartitionOK handles this case with default header values.

OK
*/
type DeletePartitionOK struct {
	Payload *models.V1PartitionResponse
}

func (o *DeletePartitionOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartitionOK  %+v", 200, o.Payload)
}

func (o *DeletePartitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PartitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePartitionDefault creates a DeletePartitionDefault with default headers values
func NewDeletePartitionDefault(code int) *DeletePartitionDefault {
	return &DeletePartitionDefault{
		_statusCode: code,
	}
}

/*DeletePartitionDefault handles this case with default header values.

Error
*/
type DeletePartitionDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the delete partition default response
func (o *DeletePartitionDefault) Code() int {
	return o._statusCode
}

func (o *DeletePartitionDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartition default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePartitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
