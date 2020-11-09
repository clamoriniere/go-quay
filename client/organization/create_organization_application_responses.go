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

// CreateOrganizationApplicationReader is a Reader for the CreateOrganizationApplication structure.
type CreateOrganizationApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationApplicationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrganizationApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateOrganizationApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrganizationApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationApplicationCreated creates a CreateOrganizationApplicationCreated with default headers values
func NewCreateOrganizationApplicationCreated() *CreateOrganizationApplicationCreated {
	return &CreateOrganizationApplicationCreated{}
}

/*CreateOrganizationApplicationCreated handles this case with default header values.

Successful creation
*/
type CreateOrganizationApplicationCreated struct {
}

func (o *CreateOrganizationApplicationCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/applications][%d] createOrganizationApplicationCreated ", 201)
}

func (o *CreateOrganizationApplicationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationApplicationBadRequest creates a CreateOrganizationApplicationBadRequest with default headers values
func NewCreateOrganizationApplicationBadRequest() *CreateOrganizationApplicationBadRequest {
	return &CreateOrganizationApplicationBadRequest{}
}

/*CreateOrganizationApplicationBadRequest handles this case with default header values.

Bad Request
*/
type CreateOrganizationApplicationBadRequest struct {
	Payload *models.APIError
}

func (o *CreateOrganizationApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/applications][%d] createOrganizationApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrganizationApplicationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOrganizationApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationApplicationUnauthorized creates a CreateOrganizationApplicationUnauthorized with default headers values
func NewCreateOrganizationApplicationUnauthorized() *CreateOrganizationApplicationUnauthorized {
	return &CreateOrganizationApplicationUnauthorized{}
}

/*CreateOrganizationApplicationUnauthorized handles this case with default header values.

Session required
*/
type CreateOrganizationApplicationUnauthorized struct {
	Payload *models.APIError
}

func (o *CreateOrganizationApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/applications][%d] createOrganizationApplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrganizationApplicationUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOrganizationApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationApplicationForbidden creates a CreateOrganizationApplicationForbidden with default headers values
func NewCreateOrganizationApplicationForbidden() *CreateOrganizationApplicationForbidden {
	return &CreateOrganizationApplicationForbidden{}
}

/*CreateOrganizationApplicationForbidden handles this case with default header values.

Unauthorized access
*/
type CreateOrganizationApplicationForbidden struct {
	Payload *models.APIError
}

func (o *CreateOrganizationApplicationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/applications][%d] createOrganizationApplicationForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationApplicationForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOrganizationApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationApplicationNotFound creates a CreateOrganizationApplicationNotFound with default headers values
func NewCreateOrganizationApplicationNotFound() *CreateOrganizationApplicationNotFound {
	return &CreateOrganizationApplicationNotFound{}
}

/*CreateOrganizationApplicationNotFound handles this case with default header values.

Not found
*/
type CreateOrganizationApplicationNotFound struct {
	Payload *models.APIError
}

func (o *CreateOrganizationApplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/applications][%d] createOrganizationApplicationNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationApplicationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOrganizationApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
