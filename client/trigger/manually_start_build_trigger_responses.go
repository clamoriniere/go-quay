// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// ManuallyStartBuildTriggerReader is a Reader for the ManuallyStartBuildTrigger structure.
type ManuallyStartBuildTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManuallyStartBuildTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewManuallyStartBuildTriggerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewManuallyStartBuildTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewManuallyStartBuildTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewManuallyStartBuildTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewManuallyStartBuildTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManuallyStartBuildTriggerCreated creates a ManuallyStartBuildTriggerCreated with default headers values
func NewManuallyStartBuildTriggerCreated() *ManuallyStartBuildTriggerCreated {
	return &ManuallyStartBuildTriggerCreated{}
}

/*ManuallyStartBuildTriggerCreated handles this case with default header values.

Successful creation
*/
type ManuallyStartBuildTriggerCreated struct {
}

func (o *ManuallyStartBuildTriggerCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerCreated ", 201)
}

func (o *ManuallyStartBuildTriggerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewManuallyStartBuildTriggerBadRequest creates a ManuallyStartBuildTriggerBadRequest with default headers values
func NewManuallyStartBuildTriggerBadRequest() *ManuallyStartBuildTriggerBadRequest {
	return &ManuallyStartBuildTriggerBadRequest{}
}

/*ManuallyStartBuildTriggerBadRequest handles this case with default header values.

Bad Request
*/
type ManuallyStartBuildTriggerBadRequest struct {
	Payload *models.APIError
}

func (o *ManuallyStartBuildTriggerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *ManuallyStartBuildTriggerBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ManuallyStartBuildTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManuallyStartBuildTriggerUnauthorized creates a ManuallyStartBuildTriggerUnauthorized with default headers values
func NewManuallyStartBuildTriggerUnauthorized() *ManuallyStartBuildTriggerUnauthorized {
	return &ManuallyStartBuildTriggerUnauthorized{}
}

/*ManuallyStartBuildTriggerUnauthorized handles this case with default header values.

Session required
*/
type ManuallyStartBuildTriggerUnauthorized struct {
	Payload *models.APIError
}

func (o *ManuallyStartBuildTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *ManuallyStartBuildTriggerUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ManuallyStartBuildTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManuallyStartBuildTriggerForbidden creates a ManuallyStartBuildTriggerForbidden with default headers values
func NewManuallyStartBuildTriggerForbidden() *ManuallyStartBuildTriggerForbidden {
	return &ManuallyStartBuildTriggerForbidden{}
}

/*ManuallyStartBuildTriggerForbidden handles this case with default header values.

Unauthorized access
*/
type ManuallyStartBuildTriggerForbidden struct {
	Payload *models.APIError
}

func (o *ManuallyStartBuildTriggerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerForbidden  %+v", 403, o.Payload)
}

func (o *ManuallyStartBuildTriggerForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ManuallyStartBuildTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManuallyStartBuildTriggerNotFound creates a ManuallyStartBuildTriggerNotFound with default headers values
func NewManuallyStartBuildTriggerNotFound() *ManuallyStartBuildTriggerNotFound {
	return &ManuallyStartBuildTriggerNotFound{}
}

/*ManuallyStartBuildTriggerNotFound handles this case with default header values.

Not found
*/
type ManuallyStartBuildTriggerNotFound struct {
	Payload *models.APIError
}

func (o *ManuallyStartBuildTriggerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerNotFound  %+v", 404, o.Payload)
}

func (o *ManuallyStartBuildTriggerNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ManuallyStartBuildTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
