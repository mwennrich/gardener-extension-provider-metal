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

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*ListProjectsOK handles this case with default header values.

OK
*/
type ListProjectsOK struct {
	Payload []*models.V1ProjectResponse
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /v1/project][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsDefault creates a ListProjectsDefault with default headers values
func NewListProjectsDefault(code int) *ListProjectsDefault {
	return &ListProjectsDefault{
		_statusCode: code,
	}
}

/*ListProjectsDefault handles this case with default header values.

Error
*/
type ListProjectsDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the list projects default response
func (o *ListProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/project][%d] listProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
