// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// FindProjectsReader is a Reader for the FindProjects structure.
type FindProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindProjectsOK creates a FindProjectsOK with default headers values
func NewFindProjectsOK() *FindProjectsOK {
	return &FindProjectsOK{}
}

/*FindProjectsOK handles this case with default header values.

OK
*/
type FindProjectsOK struct {
	Payload []*models.V1ProjectResponse
}

func (o *FindProjectsOK) Error() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjectsOK  %+v", 200, o.Payload)
}

func (o *FindProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindProjectsDefault creates a FindProjectsDefault with default headers values
func NewFindProjectsDefault(code int) *FindProjectsDefault {
	return &FindProjectsDefault{
		_statusCode: code,
	}
}

/*FindProjectsDefault handles this case with default header values.

Error
*/
type FindProjectsDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the find projects default response
func (o *FindProjectsDefault) Code() int {
	return o._statusCode
}

func (o *FindProjectsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjects default  %+v", o._statusCode, o.Payload)
}

func (o *FindProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
