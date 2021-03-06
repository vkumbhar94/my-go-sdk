// Code generated by go-swagger; DO NOT EDIT.

package user_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vkumbhar94/my-go-sdk/models"
)

// GetUserUsingGETReader is a Reader for the GetUserUsingGET structure.
type GetUserUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserUsingGETOK creates a GetUserUsingGETOK with default headers values
func NewGetUserUsingGETOK() *GetUserUsingGETOK {
	return &GetUserUsingGETOK{}
}

/*GetUserUsingGETOK handles this case with default header values.

OK
*/
type GetUserUsingGETOK struct {
	Payload *models.User
}

func (o *GetUserUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /rest/user/{userName}][%d] getUserUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUserUsingGETOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUsingGETUnauthorized creates a GetUserUsingGETUnauthorized with default headers values
func NewGetUserUsingGETUnauthorized() *GetUserUsingGETUnauthorized {
	return &GetUserUsingGETUnauthorized{}
}

/*GetUserUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUserUsingGETUnauthorized struct {
}

func (o *GetUserUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /rest/user/{userName}][%d] getUserUsingGETUnauthorized ", 401)
}

func (o *GetUserUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserUsingGETForbidden creates a GetUserUsingGETForbidden with default headers values
func NewGetUserUsingGETForbidden() *GetUserUsingGETForbidden {
	return &GetUserUsingGETForbidden{}
}

/*GetUserUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetUserUsingGETForbidden struct {
}

func (o *GetUserUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /rest/user/{userName}][%d] getUserUsingGETForbidden ", 403)
}

func (o *GetUserUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserUsingGETNotFound creates a GetUserUsingGETNotFound with default headers values
func NewGetUserUsingGETNotFound() *GetUserUsingGETNotFound {
	return &GetUserUsingGETNotFound{}
}

/*GetUserUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetUserUsingGETNotFound struct {
}

func (o *GetUserUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /rest/user/{userName}][%d] getUserUsingGETNotFound ", 404)
}

func (o *GetUserUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
