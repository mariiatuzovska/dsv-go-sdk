// Code generated by go-swagger; DO NOT EDIT.

package p_k_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// SignCertificateReader is a Reader for the SignCertificate structure.
type SignCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSignCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSignCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSignCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSignCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSignCertificateOK creates a SignCertificateOK with default headers values
func NewSignCertificateOK() *SignCertificateOK {
	return &SignCertificateOK{}
}

/* SignCertificateOK describes a response with status code 200, with default header values.

no error
*/
type SignCertificateOK struct {
	Payload *models.SignedLeafCertificate
}

// IsSuccess returns true when this sign certificate o k response has a 2xx status code
func (o *SignCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sign certificate o k response has a 3xx status code
func (o *SignCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign certificate o k response has a 4xx status code
func (o *SignCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sign certificate o k response has a 5xx status code
func (o *SignCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sign certificate o k response a status code equal to that given
func (o *SignCertificateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SignCertificateOK) Error() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateOK  %+v", 200, o.Payload)
}

func (o *SignCertificateOK) String() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateOK  %+v", 200, o.Payload)
}

func (o *SignCertificateOK) GetPayload() *models.SignedLeafCertificate {
	return o.Payload
}

func (o *SignCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignedLeafCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignCertificateBadRequest creates a SignCertificateBadRequest with default headers values
func NewSignCertificateBadRequest() *SignCertificateBadRequest {
	return &SignCertificateBadRequest{}
}

/* SignCertificateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type SignCertificateBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this sign certificate bad request response has a 2xx status code
func (o *SignCertificateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sign certificate bad request response has a 3xx status code
func (o *SignCertificateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign certificate bad request response has a 4xx status code
func (o *SignCertificateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this sign certificate bad request response has a 5xx status code
func (o *SignCertificateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this sign certificate bad request response a status code equal to that given
func (o *SignCertificateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SignCertificateBadRequest) Error() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *SignCertificateBadRequest) String() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *SignCertificateBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SignCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignCertificateUnauthorized creates a SignCertificateUnauthorized with default headers values
func NewSignCertificateUnauthorized() *SignCertificateUnauthorized {
	return &SignCertificateUnauthorized{}
}

/* SignCertificateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type SignCertificateUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this sign certificate unauthorized response has a 2xx status code
func (o *SignCertificateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sign certificate unauthorized response has a 3xx status code
func (o *SignCertificateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign certificate unauthorized response has a 4xx status code
func (o *SignCertificateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this sign certificate unauthorized response has a 5xx status code
func (o *SignCertificateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this sign certificate unauthorized response a status code equal to that given
func (o *SignCertificateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SignCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateUnauthorized  %+v", 401, o.Payload)
}

func (o *SignCertificateUnauthorized) String() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateUnauthorized  %+v", 401, o.Payload)
}

func (o *SignCertificateUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SignCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignCertificateForbidden creates a SignCertificateForbidden with default headers values
func NewSignCertificateForbidden() *SignCertificateForbidden {
	return &SignCertificateForbidden{}
}

/* SignCertificateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type SignCertificateForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this sign certificate forbidden response has a 2xx status code
func (o *SignCertificateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sign certificate forbidden response has a 3xx status code
func (o *SignCertificateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign certificate forbidden response has a 4xx status code
func (o *SignCertificateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sign certificate forbidden response has a 5xx status code
func (o *SignCertificateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sign certificate forbidden response a status code equal to that given
func (o *SignCertificateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SignCertificateForbidden) Error() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateForbidden  %+v", 403, o.Payload)
}

func (o *SignCertificateForbidden) String() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateForbidden  %+v", 403, o.Payload)
}

func (o *SignCertificateForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SignCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignCertificateInternalServerError creates a SignCertificateInternalServerError with default headers values
func NewSignCertificateInternalServerError() *SignCertificateInternalServerError {
	return &SignCertificateInternalServerError{}
}

/* SignCertificateInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SignCertificateInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this sign certificate internal server error response has a 2xx status code
func (o *SignCertificateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sign certificate internal server error response has a 3xx status code
func (o *SignCertificateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign certificate internal server error response has a 4xx status code
func (o *SignCertificateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this sign certificate internal server error response has a 5xx status code
func (o *SignCertificateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this sign certificate internal server error response a status code equal to that given
func (o *SignCertificateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SignCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *SignCertificateInternalServerError) String() string {
	return fmt.Sprintf("[POST /pki/sign][%d] signCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *SignCertificateInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SignCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
