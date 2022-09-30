// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// PingEngineReader is a Reader for the PingEngine structure.
type PingEngineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingEngineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingEngineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingEngineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPingEngineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPingEngineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingEngineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPingEngineOK creates a PingEngineOK with default headers values
func NewPingEngineOK() *PingEngineOK {
	return &PingEngineOK{}
}

/* PingEngineOK describes a response with status code 200, with default header values.

no error
*/
type PingEngineOK struct {
	Payload *models.PingResponse
}

// IsSuccess returns true when this ping engine o k response has a 2xx status code
func (o *PingEngineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping engine o k response has a 3xx status code
func (o *PingEngineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping engine o k response has a 4xx status code
func (o *PingEngineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping engine o k response has a 5xx status code
func (o *PingEngineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping engine o k response a status code equal to that given
func (o *PingEngineOK) IsCode(code int) bool {
	return code == 200
}

func (o *PingEngineOK) Error() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineOK  %+v", 200, o.Payload)
}

func (o *PingEngineOK) String() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineOK  %+v", 200, o.Payload)
}

func (o *PingEngineOK) GetPayload() *models.PingResponse {
	return o.Payload
}

func (o *PingEngineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingEngineBadRequest creates a PingEngineBadRequest with default headers values
func NewPingEngineBadRequest() *PingEngineBadRequest {
	return &PingEngineBadRequest{}
}

/* PingEngineBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PingEngineBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this ping engine bad request response has a 2xx status code
func (o *PingEngineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping engine bad request response has a 3xx status code
func (o *PingEngineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping engine bad request response has a 4xx status code
func (o *PingEngineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping engine bad request response has a 5xx status code
func (o *PingEngineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ping engine bad request response a status code equal to that given
func (o *PingEngineBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PingEngineBadRequest) Error() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineBadRequest  %+v", 400, o.Payload)
}

func (o *PingEngineBadRequest) String() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineBadRequest  %+v", 400, o.Payload)
}

func (o *PingEngineBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PingEngineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingEngineUnauthorized creates a PingEngineUnauthorized with default headers values
func NewPingEngineUnauthorized() *PingEngineUnauthorized {
	return &PingEngineUnauthorized{}
}

/* PingEngineUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PingEngineUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this ping engine unauthorized response has a 2xx status code
func (o *PingEngineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping engine unauthorized response has a 3xx status code
func (o *PingEngineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping engine unauthorized response has a 4xx status code
func (o *PingEngineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping engine unauthorized response has a 5xx status code
func (o *PingEngineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ping engine unauthorized response a status code equal to that given
func (o *PingEngineUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PingEngineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *PingEngineUnauthorized) String() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *PingEngineUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PingEngineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingEngineForbidden creates a PingEngineForbidden with default headers values
func NewPingEngineForbidden() *PingEngineForbidden {
	return &PingEngineForbidden{}
}

/* PingEngineForbidden describes a response with status code 403, with default header values.

forbidden
*/
type PingEngineForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this ping engine forbidden response has a 2xx status code
func (o *PingEngineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping engine forbidden response has a 3xx status code
func (o *PingEngineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping engine forbidden response has a 4xx status code
func (o *PingEngineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping engine forbidden response has a 5xx status code
func (o *PingEngineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ping engine forbidden response a status code equal to that given
func (o *PingEngineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PingEngineForbidden) Error() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineForbidden  %+v", 403, o.Payload)
}

func (o *PingEngineForbidden) String() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineForbidden  %+v", 403, o.Payload)
}

func (o *PingEngineForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PingEngineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingEngineInternalServerError creates a PingEngineInternalServerError with default headers values
func NewPingEngineInternalServerError() *PingEngineInternalServerError {
	return &PingEngineInternalServerError{}
}

/* PingEngineInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PingEngineInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this ping engine internal server error response has a 2xx status code
func (o *PingEngineInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping engine internal server error response has a 3xx status code
func (o *PingEngineInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping engine internal server error response has a 4xx status code
func (o *PingEngineInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping engine internal server error response has a 5xx status code
func (o *PingEngineInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ping engine internal server error response a status code equal to that given
func (o *PingEngineInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PingEngineInternalServerError) Error() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *PingEngineInternalServerError) String() string {
	return fmt.Sprintf("[POST /engines/{name}/ping][%d] pingEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *PingEngineInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PingEngineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
