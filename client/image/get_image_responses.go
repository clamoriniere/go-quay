// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// GetImageReader is a Reader for the GetImage structure.
type GetImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetImageOK creates a GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {
	return &GetImageOK{}
}

/*GetImageOK handles this case with default header values.

Successful invocation
*/
type GetImageOK struct {
}

func (o *GetImageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{image_id}][%d] getImageOK ", 200)
}

func (o *GetImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetImageBadRequest creates a GetImageBadRequest with default headers values
func NewGetImageBadRequest() *GetImageBadRequest {
	return &GetImageBadRequest{}
}

/*GetImageBadRequest handles this case with default header values.

Bad Request
*/
type GetImageBadRequest struct {
	Payload *models.APIError
}

func (o *GetImageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{image_id}][%d] getImageBadRequest  %+v", 400, o.Payload)
}

func (o *GetImageBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageUnauthorized creates a GetImageUnauthorized with default headers values
func NewGetImageUnauthorized() *GetImageUnauthorized {
	return &GetImageUnauthorized{}
}

/*GetImageUnauthorized handles this case with default header values.

Session required
*/
type GetImageUnauthorized struct {
	Payload *models.APIError
}

func (o *GetImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{image_id}][%d] getImageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetImageUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageForbidden creates a GetImageForbidden with default headers values
func NewGetImageForbidden() *GetImageForbidden {
	return &GetImageForbidden{}
}

/*GetImageForbidden handles this case with default header values.

Unauthorized access
*/
type GetImageForbidden struct {
	Payload *models.APIError
}

func (o *GetImageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{image_id}][%d] getImageForbidden  %+v", 403, o.Payload)
}

func (o *GetImageForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageNotFound creates a GetImageNotFound with default headers values
func NewGetImageNotFound() *GetImageNotFound {
	return &GetImageNotFound{}
}

/*GetImageNotFound handles this case with default header values.

Not found
*/
type GetImageNotFound struct {
	Payload *models.APIError
}

func (o *GetImageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{image_id}][%d] getImageNotFound  %+v", 404, o.Payload)
}

func (o *GetImageNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
