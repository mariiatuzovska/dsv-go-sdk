// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// SearchSettingsReader is a Reader for the SearchSettings structure.
type SearchSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchSettingsOK creates a SearchSettingsOK with default headers values
func NewSearchSettingsOK() *SearchSettingsOK {
	return &SearchSettingsOK{}
}

/* SearchSettingsOK describes a response with status code 200, with default header values.

no error
*/
type SearchSettingsOK struct {
	Payload *models.SearchResponse
}

// IsSuccess returns true when this search settings o k response has a 2xx status code
func (o *SearchSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search settings o k response has a 3xx status code
func (o *SearchSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings o k response has a 4xx status code
func (o *SearchSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search settings o k response has a 5xx status code
func (o *SearchSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search settings o k response a status code equal to that given
func (o *SearchSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchSettingsOK) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsOK  %+v", 200, o.Payload)
}

func (o *SearchSettingsOK) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsOK  %+v", 200, o.Payload)
}

func (o *SearchSettingsOK) GetPayload() *models.SearchResponse {
	return o.Payload
}

func (o *SearchSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSettingsBadRequest creates a SearchSettingsBadRequest with default headers values
func NewSearchSettingsBadRequest() *SearchSettingsBadRequest {
	return &SearchSettingsBadRequest{}
}

/* SearchSettingsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type SearchSettingsBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this search settings bad request response has a 2xx status code
func (o *SearchSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search settings bad request response has a 3xx status code
func (o *SearchSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings bad request response has a 4xx status code
func (o *SearchSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search settings bad request response has a 5xx status code
func (o *SearchSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search settings bad request response a status code equal to that given
func (o *SearchSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchSettingsBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SearchSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSettingsUnauthorized creates a SearchSettingsUnauthorized with default headers values
func NewSearchSettingsUnauthorized() *SearchSettingsUnauthorized {
	return &SearchSettingsUnauthorized{}
}

/* SearchSettingsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type SearchSettingsUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this search settings unauthorized response has a 2xx status code
func (o *SearchSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search settings unauthorized response has a 3xx status code
func (o *SearchSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings unauthorized response has a 4xx status code
func (o *SearchSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search settings unauthorized response has a 5xx status code
func (o *SearchSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search settings unauthorized response a status code equal to that given
func (o *SearchSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchSettingsUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SearchSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSettingsForbidden creates a SearchSettingsForbidden with default headers values
func NewSearchSettingsForbidden() *SearchSettingsForbidden {
	return &SearchSettingsForbidden{}
}

/* SearchSettingsForbidden describes a response with status code 403, with default header values.

forbidden
*/
type SearchSettingsForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this search settings forbidden response has a 2xx status code
func (o *SearchSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search settings forbidden response has a 3xx status code
func (o *SearchSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings forbidden response has a 4xx status code
func (o *SearchSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search settings forbidden response has a 5xx status code
func (o *SearchSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search settings forbidden response a status code equal to that given
func (o *SearchSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsForbidden  %+v", 403, o.Payload)
}

func (o *SearchSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsForbidden  %+v", 403, o.Payload)
}

func (o *SearchSettingsForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SearchSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSettingsNotFound creates a SearchSettingsNotFound with default headers values
func NewSearchSettingsNotFound() *SearchSettingsNotFound {
	return &SearchSettingsNotFound{}
}

/* SearchSettingsNotFound describes a response with status code 404, with default header values.

not found
*/
type SearchSettingsNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this search settings not found response has a 2xx status code
func (o *SearchSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search settings not found response has a 3xx status code
func (o *SearchSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings not found response has a 4xx status code
func (o *SearchSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search settings not found response has a 5xx status code
func (o *SearchSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search settings not found response a status code equal to that given
func (o *SearchSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SearchSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SearchSettingsNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SearchSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSettingsInternalServerError creates a SearchSettingsInternalServerError with default headers values
func NewSearchSettingsInternalServerError() *SearchSettingsInternalServerError {
	return &SearchSettingsInternalServerError{}
}

/* SearchSettingsInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SearchSettingsInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this search settings internal server error response has a 2xx status code
func (o *SearchSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search settings internal server error response has a 3xx status code
func (o *SearchSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search settings internal server error response has a 4xx status code
func (o *SearchSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search settings internal server error response has a 5xx status code
func (o *SearchSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search settings internal server error response a status code equal to that given
func (o *SearchSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/auth][%d] searchSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchSettingsInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SearchSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
