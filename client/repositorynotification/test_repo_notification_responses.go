// Code generated by go-swagger; DO NOT EDIT.

package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// TestRepoNotificationReader is a Reader for the TestRepoNotification structure.
type TestRepoNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestRepoNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTestRepoNotificationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestRepoNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTestRepoNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTestRepoNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestRepoNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestRepoNotificationCreated creates a TestRepoNotificationCreated with default headers values
func NewTestRepoNotificationCreated() *TestRepoNotificationCreated {
	return &TestRepoNotificationCreated{}
}

/*TestRepoNotificationCreated handles this case with default header values.

Successful creation
*/
type TestRepoNotificationCreated struct {
}

func (o *TestRepoNotificationCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/notification/{uuid}/test][%d] testRepoNotificationCreated ", 201)
}

func (o *TestRepoNotificationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestRepoNotificationBadRequest creates a TestRepoNotificationBadRequest with default headers values
func NewTestRepoNotificationBadRequest() *TestRepoNotificationBadRequest {
	return &TestRepoNotificationBadRequest{}
}

/*TestRepoNotificationBadRequest handles this case with default header values.

Bad Request
*/
type TestRepoNotificationBadRequest struct {
	Payload *models.APIError
}

func (o *TestRepoNotificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/notification/{uuid}/test][%d] testRepoNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *TestRepoNotificationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *TestRepoNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestRepoNotificationUnauthorized creates a TestRepoNotificationUnauthorized with default headers values
func NewTestRepoNotificationUnauthorized() *TestRepoNotificationUnauthorized {
	return &TestRepoNotificationUnauthorized{}
}

/*TestRepoNotificationUnauthorized handles this case with default header values.

Session required
*/
type TestRepoNotificationUnauthorized struct {
	Payload *models.APIError
}

func (o *TestRepoNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/notification/{uuid}/test][%d] testRepoNotificationUnauthorized  %+v", 401, o.Payload)
}

func (o *TestRepoNotificationUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *TestRepoNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestRepoNotificationForbidden creates a TestRepoNotificationForbidden with default headers values
func NewTestRepoNotificationForbidden() *TestRepoNotificationForbidden {
	return &TestRepoNotificationForbidden{}
}

/*TestRepoNotificationForbidden handles this case with default header values.

Unauthorized access
*/
type TestRepoNotificationForbidden struct {
	Payload *models.APIError
}

func (o *TestRepoNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/notification/{uuid}/test][%d] testRepoNotificationForbidden  %+v", 403, o.Payload)
}

func (o *TestRepoNotificationForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *TestRepoNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestRepoNotificationNotFound creates a TestRepoNotificationNotFound with default headers values
func NewTestRepoNotificationNotFound() *TestRepoNotificationNotFound {
	return &TestRepoNotificationNotFound{}
}

/*TestRepoNotificationNotFound handles this case with default header values.

Not found
*/
type TestRepoNotificationNotFound struct {
	Payload *models.APIError
}

func (o *TestRepoNotificationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/notification/{uuid}/test][%d] testRepoNotificationNotFound  %+v", 404, o.Payload)
}

func (o *TestRepoNotificationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *TestRepoNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
