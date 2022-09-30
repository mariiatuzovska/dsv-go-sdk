// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// DeleteSecretReader is a Reader for the DeleteSecret structure.
type DeleteSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSecretOK creates a DeleteSecretOK with default headers values
func NewDeleteSecretOK() *DeleteSecretOK {
	return &DeleteSecretOK{}
}

/* DeleteSecretOK describes a response with status code 200, with default header values.

no error
*/
type DeleteSecretOK struct {
	Payload *models.MessageResponse
}

// IsSuccess returns true when this delete secret o k response has a 2xx status code
func (o *DeleteSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete secret o k response has a 3xx status code
func (o *DeleteSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret o k response has a 4xx status code
func (o *DeleteSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete secret o k response has a 5xx status code
func (o *DeleteSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret o k response a status code equal to that given
func (o *DeleteSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretOK) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretOK) GetPayload() *models.MessageResponse {
	return o.Payload
}

func (o *DeleteSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretNoContent creates a DeleteSecretNoContent with default headers values
func NewDeleteSecretNoContent() *DeleteSecretNoContent {
	return &DeleteSecretNoContent{}
}

/* DeleteSecretNoContent describes a response with status code 204, with default header values.

no error
*/
type DeleteSecretNoContent struct {
}

// IsSuccess returns true when this delete secret no content response has a 2xx status code
func (o *DeleteSecretNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete secret no content response has a 3xx status code
func (o *DeleteSecretNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret no content response has a 4xx status code
func (o *DeleteSecretNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete secret no content response has a 5xx status code
func (o *DeleteSecretNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret no content response a status code equal to that given
func (o *DeleteSecretNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteSecretNoContent) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretNoContent ", 204)
}

func (o *DeleteSecretNoContent) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretNoContent ", 204)
}

func (o *DeleteSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecretBadRequest creates a DeleteSecretBadRequest with default headers values
func NewDeleteSecretBadRequest() *DeleteSecretBadRequest {
	return &DeleteSecretBadRequest{}
}

/* DeleteSecretBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DeleteSecretBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete secret bad request response has a 2xx status code
func (o *DeleteSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret bad request response has a 3xx status code
func (o *DeleteSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret bad request response has a 4xx status code
func (o *DeleteSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete secret bad request response has a 5xx status code
func (o *DeleteSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret bad request response a status code equal to that given
func (o *DeleteSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteSecretBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSecretBadRequest) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSecretBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretUnauthorized creates a DeleteSecretUnauthorized with default headers values
func NewDeleteSecretUnauthorized() *DeleteSecretUnauthorized {
	return &DeleteSecretUnauthorized{}
}

/* DeleteSecretUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteSecretUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete secret unauthorized response has a 2xx status code
func (o *DeleteSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret unauthorized response has a 3xx status code
func (o *DeleteSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret unauthorized response has a 4xx status code
func (o *DeleteSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete secret unauthorized response has a 5xx status code
func (o *DeleteSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret unauthorized response a status code equal to that given
func (o *DeleteSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteSecretUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSecretUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSecretUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretForbidden creates a DeleteSecretForbidden with default headers values
func NewDeleteSecretForbidden() *DeleteSecretForbidden {
	return &DeleteSecretForbidden{}
}

/* DeleteSecretForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteSecretForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete secret forbidden response has a 2xx status code
func (o *DeleteSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret forbidden response has a 3xx status code
func (o *DeleteSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret forbidden response has a 4xx status code
func (o *DeleteSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete secret forbidden response has a 5xx status code
func (o *DeleteSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret forbidden response a status code equal to that given
func (o *DeleteSecretForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteSecretForbidden) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSecretForbidden) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSecretForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretNotFound creates a DeleteSecretNotFound with default headers values
func NewDeleteSecretNotFound() *DeleteSecretNotFound {
	return &DeleteSecretNotFound{}
}

/* DeleteSecretNotFound describes a response with status code 404, with default header values.

not found
*/
type DeleteSecretNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete secret not found response has a 2xx status code
func (o *DeleteSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret not found response has a 3xx status code
func (o *DeleteSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret not found response has a 4xx status code
func (o *DeleteSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete secret not found response has a 5xx status code
func (o *DeleteSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret not found response a status code equal to that given
func (o *DeleteSecretNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteSecretNotFound) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSecretNotFound) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSecretNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretInternalServerError creates a DeleteSecretInternalServerError with default headers values
func NewDeleteSecretInternalServerError() *DeleteSecretInternalServerError {
	return &DeleteSecretInternalServerError{}
}

/* DeleteSecretInternalServerError describes a response with status code 500, with default header values.

server error
*/
type DeleteSecretInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete secret internal server error response has a 2xx status code
func (o *DeleteSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret internal server error response has a 3xx status code
func (o *DeleteSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret internal server error response has a 4xx status code
func (o *DeleteSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete secret internal server error response has a 5xx status code
func (o *DeleteSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete secret internal server error response a status code equal to that given
func (o *DeleteSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteSecretInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSecretInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /secrets/{path}][%d] deleteSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSecretInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
