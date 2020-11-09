// Code generated by go-swagger; DO NOT EDIT.

package signing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// GetRepoSignaturesReader is a Reader for the GetRepoSignatures structure.
type GetRepoSignaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepoSignaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepoSignaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRepoSignaturesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRepoSignaturesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepoSignaturesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepoSignaturesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepoSignaturesOK creates a GetRepoSignaturesOK with default headers values
func NewGetRepoSignaturesOK() *GetRepoSignaturesOK {
	return &GetRepoSignaturesOK{}
}

/*GetRepoSignaturesOK handles this case with default header values.

Successful invocation
*/
type GetRepoSignaturesOK struct {
}

func (o *GetRepoSignaturesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/signatures][%d] getRepoSignaturesOK ", 200)
}

func (o *GetRepoSignaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoSignaturesBadRequest creates a GetRepoSignaturesBadRequest with default headers values
func NewGetRepoSignaturesBadRequest() *GetRepoSignaturesBadRequest {
	return &GetRepoSignaturesBadRequest{}
}

/*GetRepoSignaturesBadRequest handles this case with default header values.

Bad Request
*/
type GetRepoSignaturesBadRequest struct {
	Payload *models.APIError
}

func (o *GetRepoSignaturesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/signatures][%d] getRepoSignaturesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoSignaturesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRepoSignaturesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoSignaturesUnauthorized creates a GetRepoSignaturesUnauthorized with default headers values
func NewGetRepoSignaturesUnauthorized() *GetRepoSignaturesUnauthorized {
	return &GetRepoSignaturesUnauthorized{}
}

/*GetRepoSignaturesUnauthorized handles this case with default header values.

Session required
*/
type GetRepoSignaturesUnauthorized struct {
	Payload *models.APIError
}

func (o *GetRepoSignaturesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/signatures][%d] getRepoSignaturesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepoSignaturesUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRepoSignaturesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoSignaturesForbidden creates a GetRepoSignaturesForbidden with default headers values
func NewGetRepoSignaturesForbidden() *GetRepoSignaturesForbidden {
	return &GetRepoSignaturesForbidden{}
}

/*GetRepoSignaturesForbidden handles this case with default header values.

Unauthorized access
*/
type GetRepoSignaturesForbidden struct {
	Payload *models.APIError
}

func (o *GetRepoSignaturesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/signatures][%d] getRepoSignaturesForbidden  %+v", 403, o.Payload)
}

func (o *GetRepoSignaturesForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRepoSignaturesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoSignaturesNotFound creates a GetRepoSignaturesNotFound with default headers values
func NewGetRepoSignaturesNotFound() *GetRepoSignaturesNotFound {
	return &GetRepoSignaturesNotFound{}
}

/*GetRepoSignaturesNotFound handles this case with default header values.

Not found
*/
type GetRepoSignaturesNotFound struct {
	Payload *models.APIError
}

func (o *GetRepoSignaturesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/signatures][%d] getRepoSignaturesNotFound  %+v", 404, o.Payload)
}

func (o *GetRepoSignaturesNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRepoSignaturesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
