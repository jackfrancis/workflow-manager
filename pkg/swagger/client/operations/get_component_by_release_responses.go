package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager/pkg/swagger/models"
)

// GetComponentByReleaseReader is a Reader for the GetComponentByRelease structure.
type GetComponentByReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetComponentByReleaseReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetComponentByReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetComponentByReleaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetComponentByReleaseOK creates a GetComponentByReleaseOK with default headers values
func NewGetComponentByReleaseOK() *GetComponentByReleaseOK {
	return &GetComponentByReleaseOK{}
}

/*GetComponentByReleaseOK handles this case with default header values.

component release response
*/
type GetComponentByReleaseOK struct {
	Payload *models.ComponentVersion
}

func (o *GetComponentByReleaseOK) Error() string {
	return fmt.Sprintf("[GET /v3/versions/{train}/{component}/{release}][%d] getComponentByReleaseOK  %+v", 200, o.Payload)
}

func (o *GetComponentByReleaseOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentByReleaseDefault creates a GetComponentByReleaseDefault with default headers values
func NewGetComponentByReleaseDefault(code int) *GetComponentByReleaseDefault {
	return &GetComponentByReleaseDefault{
		_statusCode: code,
	}
}

/*GetComponentByReleaseDefault handles this case with default header values.

unexpected error
*/
type GetComponentByReleaseDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get component by release default response
func (o *GetComponentByReleaseDefault) Code() int {
	return o._statusCode
}

func (o *GetComponentByReleaseDefault) Error() string {
	return fmt.Sprintf("[GET /v3/versions/{train}/{component}/{release}][%d] getComponentByRelease default  %+v", o._statusCode, o.Payload)
}

func (o *GetComponentByReleaseDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
