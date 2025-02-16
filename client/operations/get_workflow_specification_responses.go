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
)

// GetWorkflowSpecificationReader is a Reader for the GetWorkflowSpecification structure.
type GetWorkflowSpecificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowSpecificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowSpecificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetWorkflowSpecificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowSpecificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowSpecificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowSpecificationOK creates a GetWorkflowSpecificationOK with default headers values
func NewGetWorkflowSpecificationOK() *GetWorkflowSpecificationOK {
	return &GetWorkflowSpecificationOK{}
}

/*
GetWorkflowSpecificationOK describes a response with status code 200, with default header values.

Request succeeded. Workflow specification is returned.
*/
type GetWorkflowSpecificationOK struct {
	Payload *GetWorkflowSpecificationOKBody
}

// IsSuccess returns true when this get workflow specification o k response has a 2xx status code
func (o *GetWorkflowSpecificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow specification o k response has a 3xx status code
func (o *GetWorkflowSpecificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow specification o k response has a 4xx status code
func (o *GetWorkflowSpecificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow specification o k response has a 5xx status code
func (o *GetWorkflowSpecificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow specification o k response a status code equal to that given
func (o *GetWorkflowSpecificationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkflowSpecificationOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowSpecificationOK) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowSpecificationOK) GetPayload() *GetWorkflowSpecificationOKBody {
	return o.Payload
}

func (o *GetWorkflowSpecificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowSpecificationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowSpecificationForbidden creates a GetWorkflowSpecificationForbidden with default headers values
func NewGetWorkflowSpecificationForbidden() *GetWorkflowSpecificationForbidden {
	return &GetWorkflowSpecificationForbidden{}
}

/*
GetWorkflowSpecificationForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowSpecificationForbidden struct {
	Payload *GetWorkflowSpecificationForbiddenBody
}

// IsSuccess returns true when this get workflow specification forbidden response has a 2xx status code
func (o *GetWorkflowSpecificationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow specification forbidden response has a 3xx status code
func (o *GetWorkflowSpecificationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow specification forbidden response has a 4xx status code
func (o *GetWorkflowSpecificationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow specification forbidden response has a 5xx status code
func (o *GetWorkflowSpecificationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow specification forbidden response a status code equal to that given
func (o *GetWorkflowSpecificationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkflowSpecificationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkflowSpecificationForbidden) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkflowSpecificationForbidden) GetPayload() *GetWorkflowSpecificationForbiddenBody {
	return o.Payload
}

func (o *GetWorkflowSpecificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowSpecificationForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowSpecificationNotFound creates a GetWorkflowSpecificationNotFound with default headers values
func NewGetWorkflowSpecificationNotFound() *GetWorkflowSpecificationNotFound {
	return &GetWorkflowSpecificationNotFound{}
}

/*
GetWorkflowSpecificationNotFound describes a response with status code 404, with default header values.

Request failed. User does not exist.
*/
type GetWorkflowSpecificationNotFound struct {
	Payload *GetWorkflowSpecificationNotFoundBody
}

// IsSuccess returns true when this get workflow specification not found response has a 2xx status code
func (o *GetWorkflowSpecificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow specification not found response has a 3xx status code
func (o *GetWorkflowSpecificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow specification not found response has a 4xx status code
func (o *GetWorkflowSpecificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow specification not found response has a 5xx status code
func (o *GetWorkflowSpecificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow specification not found response a status code equal to that given
func (o *GetWorkflowSpecificationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkflowSpecificationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkflowSpecificationNotFound) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkflowSpecificationNotFound) GetPayload() *GetWorkflowSpecificationNotFoundBody {
	return o.Payload
}

func (o *GetWorkflowSpecificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowSpecificationNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowSpecificationInternalServerError creates a GetWorkflowSpecificationInternalServerError with default headers values
func NewGetWorkflowSpecificationInternalServerError() *GetWorkflowSpecificationInternalServerError {
	return &GetWorkflowSpecificationInternalServerError{}
}

/*
GetWorkflowSpecificationInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowSpecificationInternalServerError struct {
	Payload *GetWorkflowSpecificationInternalServerErrorBody
}

// IsSuccess returns true when this get workflow specification internal server error response has a 2xx status code
func (o *GetWorkflowSpecificationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow specification internal server error response has a 3xx status code
func (o *GetWorkflowSpecificationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow specification internal server error response has a 4xx status code
func (o *GetWorkflowSpecificationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow specification internal server error response has a 5xx status code
func (o *GetWorkflowSpecificationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workflow specification internal server error response a status code equal to that given
func (o *GetWorkflowSpecificationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkflowSpecificationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowSpecificationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowSpecificationInternalServerError) GetPayload() *GetWorkflowSpecificationInternalServerErrorBody {
	return o.Payload
}

func (o *GetWorkflowSpecificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowSpecificationInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetWorkflowSpecificationForbiddenBody get workflow specification forbidden body
swagger:model GetWorkflowSpecificationForbiddenBody
*/
type GetWorkflowSpecificationForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow specification forbidden body
func (o *GetWorkflowSpecificationForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification forbidden body based on context it is used
func (o *GetWorkflowSpecificationForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationInternalServerErrorBody get workflow specification internal server error body
swagger:model GetWorkflowSpecificationInternalServerErrorBody
*/
type GetWorkflowSpecificationInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow specification internal server error body
func (o *GetWorkflowSpecificationInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification internal server error body based on context it is used
func (o *GetWorkflowSpecificationInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationNotFoundBody get workflow specification not found body
swagger:model GetWorkflowSpecificationNotFoundBody
*/
type GetWorkflowSpecificationNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow specification not found body
func (o *GetWorkflowSpecificationNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification not found body based on context it is used
func (o *GetWorkflowSpecificationNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBody get workflow specification o k body
swagger:model GetWorkflowSpecificationOKBody
*/
type GetWorkflowSpecificationOKBody struct {

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`

	// specification
	Specification *GetWorkflowSpecificationOKBodySpecification `json:"specification,omitempty"`
}

// Validate validates this get workflow specification o k body
func (o *GetWorkflowSpecificationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSpecification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBody) validateSpecification(formats strfmt.Registry) error {
	if swag.IsZero(o.Specification) { // not required
		return nil
	}

	if o.Specification != nil {
		if err := o.Specification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflow specification o k body based on the context it is used
func (o *GetWorkflowSpecificationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBody) contextValidateSpecification(ctx context.Context, formats strfmt.Registry) error {

	if o.Specification != nil {
		if err := o.Specification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBodySpecification get workflow specification o k body specification
swagger:model GetWorkflowSpecificationOKBodySpecification
*/
type GetWorkflowSpecificationOKBodySpecification struct {

	// inputs
	Inputs *GetWorkflowSpecificationOKBodySpecificationInputs `json:"inputs,omitempty"`

	// outputs
	Outputs *GetWorkflowSpecificationOKBodySpecificationOutputs `json:"outputs,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// workflow
	Workflow *GetWorkflowSpecificationOKBodySpecificationWorkflow `json:"workflow,omitempty"`
}

// Validate validates this get workflow specification o k body specification
func (o *GetWorkflowSpecificationOKBodySpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(o.Inputs) { // not required
		return nil
	}

	if o.Inputs != nil {
		if err := o.Inputs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "inputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "inputs")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) validateOutputs(formats strfmt.Registry) error {
	if swag.IsZero(o.Outputs) { // not required
		return nil
	}

	if o.Outputs != nil {
		if err := o.Outputs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "outputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "outputs")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) validateWorkflow(formats strfmt.Registry) error {
	if swag.IsZero(o.Workflow) { // not required
		return nil
	}

	if o.Workflow != nil {
		if err := o.Workflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflow specification o k body specification based on the context it is used
func (o *GetWorkflowSpecificationOKBodySpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWorkflow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	if o.Inputs != nil {
		if err := o.Inputs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "inputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "inputs")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	if o.Outputs != nil {
		if err := o.Outputs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "outputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "outputs")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecification) contextValidateWorkflow(ctx context.Context, formats strfmt.Registry) error {

	if o.Workflow != nil {
		if err := o.Workflow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecification) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecification) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBodySpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBodySpecificationInputs get workflow specification o k body specification inputs
swagger:model GetWorkflowSpecificationOKBodySpecificationInputs
*/
type GetWorkflowSpecificationOKBodySpecificationInputs struct {

	// directories
	Directories []string `json:"directories"`

	// files
	Files []string `json:"files"`

	// options
	Options interface{} `json:"options,omitempty"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`
}

// Validate validates this get workflow specification o k body specification inputs
func (o *GetWorkflowSpecificationOKBodySpecificationInputs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification o k body specification inputs based on context it is used
func (o *GetWorkflowSpecificationOKBodySpecificationInputs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationInputs) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationInputs) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBodySpecificationInputs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBodySpecificationOutputs get workflow specification o k body specification outputs
swagger:model GetWorkflowSpecificationOKBodySpecificationOutputs
*/
type GetWorkflowSpecificationOKBodySpecificationOutputs struct {

	// directories
	Directories []string `json:"directories"`

	// files
	Files []string `json:"files"`
}

// Validate validates this get workflow specification o k body specification outputs
func (o *GetWorkflowSpecificationOKBodySpecificationOutputs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification o k body specification outputs based on context it is used
func (o *GetWorkflowSpecificationOKBodySpecificationOutputs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationOutputs) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationOutputs) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBodySpecificationOutputs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBodySpecificationWorkflow get workflow specification o k body specification workflow
swagger:model GetWorkflowSpecificationOKBodySpecificationWorkflow
*/
type GetWorkflowSpecificationOKBodySpecificationWorkflow struct {

	// file
	File string `json:"file,omitempty"`

	// specification
	Specification *GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification `json:"specification,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get workflow specification o k body specification workflow
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSpecification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) validateSpecification(formats strfmt.Registry) error {
	if swag.IsZero(o.Specification) { // not required
		return nil
	}

	if o.Specification != nil {
		if err := o.Specification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow" + "." + "specification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow" + "." + "specification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflow specification o k body specification workflow based on the context it is used
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) contextValidateSpecification(ctx context.Context, formats strfmt.Registry) error {

	if o.Specification != nil {
		if err := o.Specification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow" + "." + "specification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkflowSpecificationOK" + "." + "specification" + "." + "workflow" + "." + "specification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflow) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBodySpecificationWorkflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification get workflow specification o k body specification workflow specification
swagger:model GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification
*/
type GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification struct {

	// steps
	Steps []interface{} `json:"steps"`
}

// Validate validates this get workflow specification o k body specification workflow specification
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow specification o k body specification workflow specification based on context it is used
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification) UnmarshalBinary(b []byte) error {
	var res GetWorkflowSpecificationOKBodySpecificationWorkflowSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
