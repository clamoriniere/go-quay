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

// RemoveOrganizationMemberReader is a Reader for the RemoveOrganizationMember structure.
type RemoveOrganizationMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOrganizationMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveOrganizationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveOrganizationMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveOrganizationMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveOrganizationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveOrganizationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveOrganizationMemberNoContent creates a RemoveOrganizationMemberNoContent with default headers values
func NewRemoveOrganizationMemberNoContent() *RemoveOrganizationMemberNoContent {
	return &RemoveOrganizationMemberNoContent{}
}

/*RemoveOrganizationMemberNoContent handles this case with default header values.

Deleted
*/
type RemoveOrganizationMemberNoContent struct {
}

func (o *RemoveOrganizationMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberNoContent ", 204)
}

func (o *RemoveOrganizationMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOrganizationMemberBadRequest creates a RemoveOrganizationMemberBadRequest with default headers values
func NewRemoveOrganizationMemberBadRequest() *RemoveOrganizationMemberBadRequest {
	return &RemoveOrganizationMemberBadRequest{}
}

/*RemoveOrganizationMemberBadRequest handles this case with default header values.

Bad Request
*/
type RemoveOrganizationMemberBadRequest struct {
	Payload *models.APIError
}

func (o *RemoveOrganizationMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrganizationMemberBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RemoveOrganizationMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrganizationMemberUnauthorized creates a RemoveOrganizationMemberUnauthorized with default headers values
func NewRemoveOrganizationMemberUnauthorized() *RemoveOrganizationMemberUnauthorized {
	return &RemoveOrganizationMemberUnauthorized{}
}

/*RemoveOrganizationMemberUnauthorized handles this case with default header values.

Session required
*/
type RemoveOrganizationMemberUnauthorized struct {
	Payload *models.APIError
}

func (o *RemoveOrganizationMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveOrganizationMemberUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RemoveOrganizationMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrganizationMemberForbidden creates a RemoveOrganizationMemberForbidden with default headers values
func NewRemoveOrganizationMemberForbidden() *RemoveOrganizationMemberForbidden {
	return &RemoveOrganizationMemberForbidden{}
}

/*RemoveOrganizationMemberForbidden handles this case with default header values.

Unauthorized access
*/
type RemoveOrganizationMemberForbidden struct {
	Payload *models.APIError
}

func (o *RemoveOrganizationMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrganizationMemberForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RemoveOrganizationMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrganizationMemberNotFound creates a RemoveOrganizationMemberNotFound with default headers values
func NewRemoveOrganizationMemberNotFound() *RemoveOrganizationMemberNotFound {
	return &RemoveOrganizationMemberNotFound{}
}

/*RemoveOrganizationMemberNotFound handles this case with default header values.

Not found
*/
type RemoveOrganizationMemberNotFound struct {
	Payload *models.APIError
}

func (o *RemoveOrganizationMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *RemoveOrganizationMemberNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RemoveOrganizationMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
