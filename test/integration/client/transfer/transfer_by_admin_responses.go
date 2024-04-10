// Code generated by go-swagger; DO NOT EDIT.

package transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anoideaopen/channel-transfer/test/integration/models"
)

// TransferByAdminReader is a Reader for the TransferByAdmin structure.
type TransferByAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferByAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransferByAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTransferByAdminInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTransferByAdminDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTransferByAdminOK creates a TransferByAdminOK with default headers values
func NewTransferByAdminOK() *TransferByAdminOK {
	return &TransferByAdminOK{}
}

/*
TransferByAdminOK handles this case with default header values.

A successful response.
*/
type TransferByAdminOK struct {
	Payload *models.ChannelTransferTransferStatusResponse
}

func (o *TransferByAdminOK) Error() string {
	return fmt.Sprintf("[POST /v1/transfer/admin][%d] transferByAdminOK  %+v", 200, o.Payload)
}

func (o *TransferByAdminOK) GetPayload() *models.ChannelTransferTransferStatusResponse {
	return o.Payload
}

func (o *TransferByAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelTransferTransferStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferByAdminInternalServerError creates a TransferByAdminInternalServerError with default headers values
func NewTransferByAdminInternalServerError() *TransferByAdminInternalServerError {
	return &TransferByAdminInternalServerError{}
}

/*
TransferByAdminInternalServerError handles this case with default header values.

Internal server error
*/
type TransferByAdminInternalServerError struct {
	Payload *models.ChannelTransferErrorResponse
}

func (o *TransferByAdminInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/transfer/admin][%d] transferByAdminInternalServerError  %+v", 500, o.Payload)
}

func (o *TransferByAdminInternalServerError) GetPayload() *models.ChannelTransferErrorResponse {
	return o.Payload
}

func (o *TransferByAdminInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelTransferErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferByAdminDefault creates a TransferByAdminDefault with default headers values
func NewTransferByAdminDefault(code int) *TransferByAdminDefault {
	return &TransferByAdminDefault{
		_statusCode: code,
	}
}

/*
TransferByAdminDefault handles this case with default header values.

An unexpected error response.
*/
type TransferByAdminDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the transfer by admin default response
func (o *TransferByAdminDefault) Code() int {
	return o._statusCode
}

func (o *TransferByAdminDefault) Error() string {
	return fmt.Sprintf("[POST /v1/transfer/admin][%d] transferByAdmin default  %+v", o._statusCode, o.Payload)
}

func (o *TransferByAdminDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *TransferByAdminDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
