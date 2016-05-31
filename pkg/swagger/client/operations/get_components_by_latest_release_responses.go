package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager/pkg/swagger/models"
)

// GetComponentsByLatestReleaseReader is a Reader for the GetComponentsByLatestRelease structure.
type GetComponentsByLatestReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetComponentsByLatestReleaseReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetComponentsByLatestReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetComponentsByLatestReleaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetComponentsByLatestReleaseOK creates a GetComponentsByLatestReleaseOK with default headers values
func NewGetComponentsByLatestReleaseOK() *GetComponentsByLatestReleaseOK {
	return &GetComponentsByLatestReleaseOK{}
}

/*GetComponentsByLatestReleaseOK handles this case with default header values.

component releases response
*/
type GetComponentsByLatestReleaseOK struct {
	Payload GetComponentsByLatestReleaseOKBodyBody
}

func (o *GetComponentsByLatestReleaseOK) Error() string {
	return fmt.Sprintf("[POST /v3/versions/latest][%d] getComponentsByLatestReleaseOK  %+v", 200, o.Payload)
}

func (o *GetComponentsByLatestReleaseOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentsByLatestReleaseDefault creates a GetComponentsByLatestReleaseDefault with default headers values
func NewGetComponentsByLatestReleaseDefault(code int) *GetComponentsByLatestReleaseDefault {
	return &GetComponentsByLatestReleaseDefault{
		_statusCode: code,
	}
}

/*GetComponentsByLatestReleaseDefault handles this case with default header values.

unexpected error
*/
type GetComponentsByLatestReleaseDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get components by latest release default response
func (o *GetComponentsByLatestReleaseDefault) Code() int {
	return o._statusCode
}

func (o *GetComponentsByLatestReleaseDefault) Error() string {
	return fmt.Sprintf("[POST /v3/versions/latest][%d] getComponentsByLatestRelease default  %+v", o._statusCode, o.Payload)
}

func (o *GetComponentsByLatestReleaseDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetComponentsByLatestReleaseBody get components by latest release body

swagger:model GetComponentsByLatestReleaseBody
*/
type GetComponentsByLatestReleaseBody struct {

	/* data
	 */
	Data []*models.ComponentVersion `json:"data,omitempty"`
}

// Validate validates this get components by latest release body
func (o *GetComponentsByLatestReleaseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*GetComponentsByLatestReleaseOKBodyBody get components by latest release o k body body

swagger:model GetComponentsByLatestReleaseOKBodyBody
*/
type GetComponentsByLatestReleaseOKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.ComponentVersion `json:"data"`
}

// Validate validates this get components by latest release o k body body
func (o *GetComponentsByLatestReleaseOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseOKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getComponentsByLatestReleaseOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
