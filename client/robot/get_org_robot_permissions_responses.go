// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/clamoriniere/go-quay/models"
)

// GetOrgRobotPermissionsReader is a Reader for the GetOrgRobotPermissions structure.
type GetOrgRobotPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgRobotPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgRobotPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgRobotPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgRobotPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgRobotPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgRobotPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgRobotPermissionsOK creates a GetOrgRobotPermissionsOK with default headers values
func NewGetOrgRobotPermissionsOK() *GetOrgRobotPermissionsOK {
	return &GetOrgRobotPermissionsOK{}
}

/*GetOrgRobotPermissionsOK handles this case with default header values.

Successful invocation
*/
type GetOrgRobotPermissionsOK struct {
}

func (o *GetOrgRobotPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions][%d] getOrgRobotPermissionsOK ", 200)
}

func (o *GetOrgRobotPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrgRobotPermissionsBadRequest creates a GetOrgRobotPermissionsBadRequest with default headers values
func NewGetOrgRobotPermissionsBadRequest() *GetOrgRobotPermissionsBadRequest {
	return &GetOrgRobotPermissionsBadRequest{}
}

/*GetOrgRobotPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetOrgRobotPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrgRobotPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions][%d] getOrgRobotPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgRobotPermissionsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetOrgRobotPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotPermissionsUnauthorized creates a GetOrgRobotPermissionsUnauthorized with default headers values
func NewGetOrgRobotPermissionsUnauthorized() *GetOrgRobotPermissionsUnauthorized {
	return &GetOrgRobotPermissionsUnauthorized{}
}

/*GetOrgRobotPermissionsUnauthorized handles this case with default header values.

Session required
*/
type GetOrgRobotPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrgRobotPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions][%d] getOrgRobotPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgRobotPermissionsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetOrgRobotPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotPermissionsForbidden creates a GetOrgRobotPermissionsForbidden with default headers values
func NewGetOrgRobotPermissionsForbidden() *GetOrgRobotPermissionsForbidden {
	return &GetOrgRobotPermissionsForbidden{}
}

/*GetOrgRobotPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrgRobotPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *GetOrgRobotPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions][%d] getOrgRobotPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgRobotPermissionsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetOrgRobotPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotPermissionsNotFound creates a GetOrgRobotPermissionsNotFound with default headers values
func NewGetOrgRobotPermissionsNotFound() *GetOrgRobotPermissionsNotFound {
	return &GetOrgRobotPermissionsNotFound{}
}

/*GetOrgRobotPermissionsNotFound handles this case with default header values.

Not found
*/
type GetOrgRobotPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *GetOrgRobotPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions][%d] getOrgRobotPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRobotPermissionsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetOrgRobotPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
