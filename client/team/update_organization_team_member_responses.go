// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// UpdateOrganizationTeamMemberReader is a Reader for the UpdateOrganizationTeamMember structure.
type UpdateOrganizationTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrganizationTeamMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrganizationTeamMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrganizationTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrganizationTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationTeamMemberOK creates a UpdateOrganizationTeamMemberOK with default headers values
func NewUpdateOrganizationTeamMemberOK() *UpdateOrganizationTeamMemberOK {
	return &UpdateOrganizationTeamMemberOK{}
}

/*UpdateOrganizationTeamMemberOK handles this case with default header values.

Successful invocation
*/
type UpdateOrganizationTeamMemberOK struct {
}

func (o *UpdateOrganizationTeamMemberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}][%d] updateOrganizationTeamMemberOK ", 200)
}

func (o *UpdateOrganizationTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationTeamMemberBadRequest creates a UpdateOrganizationTeamMemberBadRequest with default headers values
func NewUpdateOrganizationTeamMemberBadRequest() *UpdateOrganizationTeamMemberBadRequest {
	return &UpdateOrganizationTeamMemberBadRequest{}
}

/*UpdateOrganizationTeamMemberBadRequest handles this case with default header values.

Bad Request
*/
type UpdateOrganizationTeamMemberBadRequest struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationTeamMemberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}][%d] updateOrganizationTeamMemberBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationTeamMemberBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationTeamMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationTeamMemberUnauthorized creates a UpdateOrganizationTeamMemberUnauthorized with default headers values
func NewUpdateOrganizationTeamMemberUnauthorized() *UpdateOrganizationTeamMemberUnauthorized {
	return &UpdateOrganizationTeamMemberUnauthorized{}
}

/*UpdateOrganizationTeamMemberUnauthorized handles this case with default header values.

Session required
*/
type UpdateOrganizationTeamMemberUnauthorized struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationTeamMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}][%d] updateOrganizationTeamMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationTeamMemberUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationTeamMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationTeamMemberForbidden creates a UpdateOrganizationTeamMemberForbidden with default headers values
func NewUpdateOrganizationTeamMemberForbidden() *UpdateOrganizationTeamMemberForbidden {
	return &UpdateOrganizationTeamMemberForbidden{}
}

/*UpdateOrganizationTeamMemberForbidden handles this case with default header values.

Unauthorized access
*/
type UpdateOrganizationTeamMemberForbidden struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}][%d] updateOrganizationTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationTeamMemberForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationTeamMemberNotFound creates a UpdateOrganizationTeamMemberNotFound with default headers values
func NewUpdateOrganizationTeamMemberNotFound() *UpdateOrganizationTeamMemberNotFound {
	return &UpdateOrganizationTeamMemberNotFound{}
}

/*UpdateOrganizationTeamMemberNotFound handles this case with default header values.

Not found
*/
type UpdateOrganizationTeamMemberNotFound struct {
	Payload *models.APIError
}

func (o *UpdateOrganizationTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}][%d] updateOrganizationTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationTeamMemberNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateOrganizationTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}