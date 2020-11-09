// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// GetUserPermissionsReader is a Reader for the GetUserPermissions structure.
type GetUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserPermissionsOK creates a GetUserPermissionsOK with default headers values
func NewGetUserPermissionsOK() *GetUserPermissionsOK {
	return &GetUserPermissionsOK{}
}

/*GetUserPermissionsOK handles this case with default header values.

Successful invocation
*/
type GetUserPermissionsOK struct {
}

func (o *GetUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/{username}][%d] getUserPermissionsOK ", 200)
}

func (o *GetUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserPermissionsBadRequest creates a GetUserPermissionsBadRequest with default headers values
func NewGetUserPermissionsBadRequest() *GetUserPermissionsBadRequest {
	return &GetUserPermissionsBadRequest{}
}

/*GetUserPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetUserPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *GetUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/{username}][%d] getUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserPermissionsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUserPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPermissionsUnauthorized creates a GetUserPermissionsUnauthorized with default headers values
func NewGetUserPermissionsUnauthorized() *GetUserPermissionsUnauthorized {
	return &GetUserPermissionsUnauthorized{}
}

/*GetUserPermissionsUnauthorized handles this case with default header values.

Session required
*/
type GetUserPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/{username}][%d] getUserPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserPermissionsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPermissionsForbidden creates a GetUserPermissionsForbidden with default headers values
func NewGetUserPermissionsForbidden() *GetUserPermissionsForbidden {
	return &GetUserPermissionsForbidden{}
}

/*GetUserPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type GetUserPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *GetUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/{username}][%d] getUserPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserPermissionsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUserPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPermissionsNotFound creates a GetUserPermissionsNotFound with default headers values
func NewGetUserPermissionsNotFound() *GetUserPermissionsNotFound {
	return &GetUserPermissionsNotFound{}
}

/*GetUserPermissionsNotFound handles this case with default header values.

Not found
*/
type GetUserPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *GetUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/{username}][%d] getUserPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserPermissionsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUserPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
