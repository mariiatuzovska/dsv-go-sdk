// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// GetUserByVersionReader is a Reader for the GetUserByVersion structure.
type GetUserByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserByVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserByVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserByVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserByVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserByVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserByVersionOK creates a GetUserByVersionOK with default headers values
func NewGetUserByVersionOK() *GetUserByVersionOK {
	return &GetUserByVersionOK{}
}

/* GetUserByVersionOK describes a response with status code 200, with default header values.

no error
*/
type GetUserByVersionOK struct {
	Payload *models.UserVersionResponse
}

// IsSuccess returns true when this get user by version o k response has a 2xx status code
func (o *GetUserByVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user by version o k response has a 3xx status code
func (o *GetUserByVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version o k response has a 4xx status code
func (o *GetUserByVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user by version o k response has a 5xx status code
func (o *GetUserByVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by version o k response a status code equal to that given
func (o *GetUserByVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserByVersionOK) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionOK  %+v", 200, o.Payload)
}

func (o *GetUserByVersionOK) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionOK  %+v", 200, o.Payload)
}

func (o *GetUserByVersionOK) GetPayload() *models.UserVersionResponse {
	return o.Payload
}

func (o *GetUserByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByVersionBadRequest creates a GetUserByVersionBadRequest with default headers values
func NewGetUserByVersionBadRequest() *GetUserByVersionBadRequest {
	return &GetUserByVersionBadRequest{}
}

/* GetUserByVersionBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetUserByVersionBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get user by version bad request response has a 2xx status code
func (o *GetUserByVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by version bad request response has a 3xx status code
func (o *GetUserByVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version bad request response has a 4xx status code
func (o *GetUserByVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by version bad request response has a 5xx status code
func (o *GetUserByVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by version bad request response a status code equal to that given
func (o *GetUserByVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserByVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserByVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserByVersionBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetUserByVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByVersionUnauthorized creates a GetUserByVersionUnauthorized with default headers values
func NewGetUserByVersionUnauthorized() *GetUserByVersionUnauthorized {
	return &GetUserByVersionUnauthorized{}
}

/* GetUserByVersionUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetUserByVersionUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get user by version unauthorized response has a 2xx status code
func (o *GetUserByVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by version unauthorized response has a 3xx status code
func (o *GetUserByVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version unauthorized response has a 4xx status code
func (o *GetUserByVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by version unauthorized response has a 5xx status code
func (o *GetUserByVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by version unauthorized response a status code equal to that given
func (o *GetUserByVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserByVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserByVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserByVersionUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetUserByVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByVersionForbidden creates a GetUserByVersionForbidden with default headers values
func NewGetUserByVersionForbidden() *GetUserByVersionForbidden {
	return &GetUserByVersionForbidden{}
}

/* GetUserByVersionForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetUserByVersionForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get user by version forbidden response has a 2xx status code
func (o *GetUserByVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by version forbidden response has a 3xx status code
func (o *GetUserByVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version forbidden response has a 4xx status code
func (o *GetUserByVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by version forbidden response has a 5xx status code
func (o *GetUserByVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by version forbidden response a status code equal to that given
func (o *GetUserByVersionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserByVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetUserByVersionForbidden) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetUserByVersionForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetUserByVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByVersionNotFound creates a GetUserByVersionNotFound with default headers values
func NewGetUserByVersionNotFound() *GetUserByVersionNotFound {
	return &GetUserByVersionNotFound{}
}

/* GetUserByVersionNotFound describes a response with status code 404, with default header values.

not found
*/
type GetUserByVersionNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get user by version not found response has a 2xx status code
func (o *GetUserByVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by version not found response has a 3xx status code
func (o *GetUserByVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version not found response has a 4xx status code
func (o *GetUserByVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by version not found response has a 5xx status code
func (o *GetUserByVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by version not found response a status code equal to that given
func (o *GetUserByVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserByVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetUserByVersionNotFound) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetUserByVersionNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetUserByVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByVersionInternalServerError creates a GetUserByVersionInternalServerError with default headers values
func NewGetUserByVersionInternalServerError() *GetUserByVersionInternalServerError {
	return &GetUserByVersionInternalServerError{}
}

/* GetUserByVersionInternalServerError describes a response with status code 500, with default header values.

server error
*/
type GetUserByVersionInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get user by version internal server error response has a 2xx status code
func (o *GetUserByVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by version internal server error response has a 3xx status code
func (o *GetUserByVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by version internal server error response has a 4xx status code
func (o *GetUserByVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user by version internal server error response has a 5xx status code
func (o *GetUserByVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user by version internal server error response a status code equal to that given
func (o *GetUserByVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserByVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserByVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{name}/version/{version}][%d] getUserByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserByVersionInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetUserByVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
