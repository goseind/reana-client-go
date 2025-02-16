// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LaunchReader is a Reader for the Launch structure.
type LaunchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LaunchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLaunchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLaunchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLaunchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLaunchOK creates a LaunchOK with default headers values
func NewLaunchOK() *LaunchOK {
	return &LaunchOK{}
}

/*
LaunchOK describes a response with status code 200, with default header values.

Request succeeded. Information of the workflow launched.
*/
type LaunchOK struct {
	Payload *LaunchOKBody
}

// IsSuccess returns true when this launch o k response has a 2xx status code
func (o *LaunchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this launch o k response has a 3xx status code
func (o *LaunchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch o k response has a 4xx status code
func (o *LaunchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this launch o k response has a 5xx status code
func (o *LaunchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this launch o k response a status code equal to that given
func (o *LaunchOK) IsCode(code int) bool {
	return code == 200
}

func (o *LaunchOK) Error() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchOK  %+v", 200, o.Payload)
}

func (o *LaunchOK) String() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchOK  %+v", 200, o.Payload)
}

func (o *LaunchOK) GetPayload() *LaunchOKBody {
	return o.Payload
}

func (o *LaunchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LaunchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchBadRequest creates a LaunchBadRequest with default headers values
func NewLaunchBadRequest() *LaunchBadRequest {
	return &LaunchBadRequest{}
}

/*
LaunchBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type LaunchBadRequest struct {
	Payload *LaunchBadRequestBody
}

// IsSuccess returns true when this launch bad request response has a 2xx status code
func (o *LaunchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this launch bad request response has a 3xx status code
func (o *LaunchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch bad request response has a 4xx status code
func (o *LaunchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this launch bad request response has a 5xx status code
func (o *LaunchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this launch bad request response a status code equal to that given
func (o *LaunchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *LaunchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchBadRequest  %+v", 400, o.Payload)
}

func (o *LaunchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchBadRequest  %+v", 400, o.Payload)
}

func (o *LaunchBadRequest) GetPayload() *LaunchBadRequestBody {
	return o.Payload
}

func (o *LaunchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LaunchBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInternalServerError creates a LaunchInternalServerError with default headers values
func NewLaunchInternalServerError() *LaunchInternalServerError {
	return &LaunchInternalServerError{}
}

/*
LaunchInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type LaunchInternalServerError struct {
	Payload *LaunchInternalServerErrorBody
}

// IsSuccess returns true when this launch internal server error response has a 2xx status code
func (o *LaunchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this launch internal server error response has a 3xx status code
func (o *LaunchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch internal server error response has a 4xx status code
func (o *LaunchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this launch internal server error response has a 5xx status code
func (o *LaunchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this launch internal server error response a status code equal to that given
func (o *LaunchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *LaunchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchInternalServerError  %+v", 500, o.Payload)
}

func (o *LaunchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/launch][%d] launchInternalServerError  %+v", 500, o.Payload)
}

func (o *LaunchInternalServerError) GetPayload() *LaunchInternalServerErrorBody {
	return o.Payload
}

func (o *LaunchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LaunchInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LaunchBadRequestBody launch bad request body
swagger:model LaunchBadRequestBody
*/
type LaunchBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this launch bad request body
func (o *LaunchBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this launch bad request body based on context it is used
func (o *LaunchBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LaunchBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LaunchBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LaunchBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LaunchBody launch body
swagger:model LaunchBody
*/
type LaunchBody struct {

	// Workflow name.
	Name string `json:"name,omitempty"`

	// Workflow parameters.
	Parameters string `json:"parameters,omitempty"`

	// Path to the workflow specification file to be used.
	Specification string `json:"specification,omitempty"`

	// Remote origin URL where the REANA specification file is hosted.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this launch body
func (o *LaunchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LaunchBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this launch body based on context it is used
func (o *LaunchBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LaunchBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LaunchBody) UnmarshalBinary(b []byte) error {
	var res LaunchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LaunchInternalServerErrorBody launch internal server error body
swagger:model LaunchInternalServerErrorBody
*/
type LaunchInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this launch internal server error body
func (o *LaunchInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this launch internal server error body based on context it is used
func (o *LaunchInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LaunchInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LaunchInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LaunchInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LaunchOKBody launch o k body
swagger:model LaunchOKBody
*/
type LaunchOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this launch o k body
func (o *LaunchOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this launch o k body based on context it is used
func (o *LaunchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LaunchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LaunchOKBody) UnmarshalBinary(b []byte) error {
	var res LaunchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
