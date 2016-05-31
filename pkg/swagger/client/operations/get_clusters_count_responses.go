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

// GetClustersCountReader is a Reader for the GetClustersCount structure.
type GetClustersCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetClustersCountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClustersCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetClustersCountOK creates a GetClustersCountOK with default headers values
func NewGetClustersCountOK() *GetClustersCountOK {
	return &GetClustersCountOK{}
}

/*GetClustersCountOK handles this case with default header values.

clusters count response
*/
type GetClustersCountOK struct {
	Payload int64
}

func (o *GetClustersCountOK) Error() string {
	return fmt.Sprintf("[GET /v3/clusters/count][%d] getClustersCountOK  %+v", 200, o.Payload)
}

func (o *GetClustersCountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersCountDefault creates a GetClustersCountDefault with default headers values
func NewGetClustersCountDefault(code int) *GetClustersCountDefault {
	return &GetClustersCountDefault{
		_statusCode: code,
	}
}

/*GetClustersCountDefault handles this case with default header values.

unexpected error
*/
type GetClustersCountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get clusters count default response
func (o *GetClustersCountDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersCountDefault) Error() string {
	return fmt.Sprintf("[GET /v3/clusters/count][%d] getClustersCount default  %+v", o._statusCode, o.Payload)
}

func (o *GetClustersCountDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
