// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// UpdateOrganizationApplicationReader is a Reader for the UpdateOrganizationApplication structure.
type UpdateOrganizationApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrganizationApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrganizationApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrganizationApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrganizationApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationApplicationOK creates a UpdateOrganizationApplicationOK with default headers values
func NewUpdateOrganizationApplicationOK() *UpdateOrganizationApplicationOK {
	return &UpdateOrganizationApplicationOK{}
}

/*UpdateOrganizationApplicationOK handles this case with default header values.

Successful invocation
*/
type UpdateOrganizationApplicationOK struct {
}

func (o *UpdateOrganizationApplicationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/applications/{client_id}][%d] updateOrganizationApplicationOK ", 200)
}

func (o *UpdateOrganizationApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationApplicationBadRequest creates a UpdateOrganizationApplicationBadRequest with default headers values
func NewUpdateOrganizationApplicationBadRequest() *UpdateOrganizationApplicationBadRequest {
	return &UpdateOrganizationApplicationBadRequest{}
}

/*UpdateOrganizationApplicationBadRequest handles this case with default header values.

Bad Request
*/
type UpdateOrganizationApplicationBadRequest struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationApplicationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/applications/{client_id}][%d] updateOrganizationApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationApplicationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationApplicationUnauthorized creates a UpdateOrganizationApplicationUnauthorized with default headers values
func NewUpdateOrganizationApplicationUnauthorized() *UpdateOrganizationApplicationUnauthorized {
	return &UpdateOrganizationApplicationUnauthorized{}
}

/*UpdateOrganizationApplicationUnauthorized handles this case with default header values.

Session required
*/
type UpdateOrganizationApplicationUnauthorized struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/applications/{client_id}][%d] updateOrganizationApplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationApplicationUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationApplicationForbidden creates a UpdateOrganizationApplicationForbidden with default headers values
func NewUpdateOrganizationApplicationForbidden() *UpdateOrganizationApplicationForbidden {
	return &UpdateOrganizationApplicationForbidden{}
}

/*UpdateOrganizationApplicationForbidden handles this case with default header values.

Unauthorized access
*/
type UpdateOrganizationApplicationForbidden struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationApplicationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/applications/{client_id}][%d] updateOrganizationApplicationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationApplicationForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationApplicationNotFound creates a UpdateOrganizationApplicationNotFound with default headers values
func NewUpdateOrganizationApplicationNotFound() *UpdateOrganizationApplicationNotFound {
	return &UpdateOrganizationApplicationNotFound{}
}

/*UpdateOrganizationApplicationNotFound handles this case with default header values.

Not found
*/
type UpdateOrganizationApplicationNotFound struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationApplicationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/applications/{client_id}][%d] updateOrganizationApplicationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationApplicationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}