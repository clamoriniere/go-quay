// Code generated by go-swagger; DO NOT EDIT.

package repotoken

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// ListRepoTokensReader is a Reader for the ListRepoTokens structure.
type ListRepoTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRepoTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRepoTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRepoTokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListRepoTokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRepoTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRepoTokensNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRepoTokensOK creates a ListRepoTokensOK with default headers values
func NewListRepoTokensOK() *ListRepoTokensOK {
	return &ListRepoTokensOK{}
}

/*ListRepoTokensOK handles this case with default header values.

Successful invocation
*/
type ListRepoTokensOK struct {
}

func (o *ListRepoTokensOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/][%d] listRepoTokensOK ", 200)
}

func (o *ListRepoTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoTokensBadRequest creates a ListRepoTokensBadRequest with default headers values
func NewListRepoTokensBadRequest() *ListRepoTokensBadRequest {
	return &ListRepoTokensBadRequest{}
}

/*ListRepoTokensBadRequest handles this case with default header values.

Bad Request
*/
type ListRepoTokensBadRequest struct {
	Payload *models.APIError
}

func (o *ListRepoTokensBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/][%d] listRepoTokensBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepoTokensBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListRepoTokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTokensUnauthorized creates a ListRepoTokensUnauthorized with default headers values
func NewListRepoTokensUnauthorized() *ListRepoTokensUnauthorized {
	return &ListRepoTokensUnauthorized{}
}

/*ListRepoTokensUnauthorized handles this case with default header values.

Session required
*/
type ListRepoTokensUnauthorized struct {
	Payload *models.APIError
}

func (o *ListRepoTokensUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/][%d] listRepoTokensUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepoTokensUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListRepoTokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTokensForbidden creates a ListRepoTokensForbidden with default headers values
func NewListRepoTokensForbidden() *ListRepoTokensForbidden {
	return &ListRepoTokensForbidden{}
}

/*ListRepoTokensForbidden handles this case with default header values.

Unauthorized access
*/
type ListRepoTokensForbidden struct {
	Payload *models.APIError
}

func (o *ListRepoTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/][%d] listRepoTokensForbidden  %+v", 403, o.Payload)
}

func (o *ListRepoTokensForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListRepoTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTokensNotFound creates a ListRepoTokensNotFound with default headers values
func NewListRepoTokensNotFound() *ListRepoTokensNotFound {
	return &ListRepoTokensNotFound{}
}

/*ListRepoTokensNotFound handles this case with default header values.

Not found
*/
type ListRepoTokensNotFound struct {
	Payload *models.APIError
}

func (o *ListRepoTokensNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/][%d] listRepoTokensNotFound  %+v", 404, o.Payload)
}

func (o *ListRepoTokensNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListRepoTokensNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
