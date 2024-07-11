// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChannelTransferGeneralParams GeneralParams general parameters for every request
//
// swagger:model channel_transferGeneralParams
type ChannelTransferGeneralParams struct {

	// chaincode name
	Chaincode string `json:"chaincode,omitempty"`

	// channel name
	Channel string `json:"channel,omitempty"`

	// method name
	MethodName string `json:"methodName,omitempty"`

	// request nonce
	Nonce string `json:"nonce,omitempty"`

	// options
	Options []*ProtobufOption `json:"options"`

	// public key of user, signed the request
	PublicKey string `json:"publicKey,omitempty"`

	// request ID (may be empty)
	RequestID string `json:"requestId,omitempty"`

	// request sign
	Sign string `json:"sign,omitempty"`
}

// Validate validates this channel transfer general params
func (m *ChannelTransferGeneralParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelTransferGeneralParams) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this channel transfer general params based on the context it is used
func (m *ChannelTransferGeneralParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelTransferGeneralParams) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {

			if swag.IsZero(m.Options[i]) { // not required
				return nil
			}

			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelTransferGeneralParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelTransferGeneralParams) UnmarshalBinary(b []byte) error {
	var res ChannelTransferGeneralParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
