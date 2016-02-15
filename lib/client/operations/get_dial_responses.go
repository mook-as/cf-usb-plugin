package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type GetDialReader struct {
	formats strfmt.Registry
}

func (o *GetDialReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetDialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDialOK creates a GetDialOK with default headers values
func NewGetDialOK() *GetDialOK {
	return &GetDialOK{}
}

/*GetDialOK

Sucessfull response
*/
type GetDialOK struct {
	Payload *models.Dial
}

func (o *GetDialOK) Error() string {
	return fmt.Sprintf("[GET /dials/{dial_id}][%d] getDialOK  %+v", 200, o.Payload)
}

func (o *GetDialOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dial)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDialNotFound creates a GetDialNotFound with default headers values
func NewGetDialNotFound() *GetDialNotFound {
	return &GetDialNotFound{}
}

/*GetDialNotFound

Not Found
*/
type GetDialNotFound struct {
}

func (o *GetDialNotFound) Error() string {
	return fmt.Sprintf("[GET /dials/{dial_id}][%d] getDialNotFound ", 404)
}

func (o *GetDialNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDialInternalServerError creates a GetDialInternalServerError with default headers values
func NewGetDialInternalServerError() *GetDialInternalServerError {
	return &GetDialInternalServerError{}
}

/*GetDialInternalServerError

Unexpected error
*/
type GetDialInternalServerError struct {
	Payload string
}

func (o *GetDialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dials/{dial_id}][%d] getDialInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDialInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
