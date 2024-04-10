// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChannelTransferTransferBeginCustomerRequest TransferBeginAdminRequest запрос о переводе токенов от Владельца токенов
//
// swagger:model channel_transferTransferBeginCustomerRequest
type ChannelTransferTransferBeginCustomerRequest struct {

	// Сколько переводят токенов
	Amount string `json:"amount,omitempty"`

	// Канал, в который переводят токены
	ChannelTo string `json:"channelTo,omitempty"`

	// Информация о транзакции
	Generals *ChannelTransferGeneralParams `json:"generals,omitempty"`

	// ID трансфера (должно быть уникально)
	IDTransfer string `json:"idTransfer,omitempty"`

	// options
	Options []*ProtobufOption `json:"options"`

	// Токены которые хотят перевести
	Token string `json:"token,omitempty"`
}

// Validate validates this channel transfer transfer begin customer request
func (m *ChannelTransferTransferBeginCustomerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenerals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelTransferTransferBeginCustomerRequest) validateGenerals(formats strfmt.Registry) error {

	if swag.IsZero(m.Generals) { // not required
		return nil
	}

	if m.Generals != nil {
		if err := m.Generals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generals")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelTransferTransferBeginCustomerRequest) validateOptions(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelTransferTransferBeginCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelTransferTransferBeginCustomerRequest) UnmarshalBinary(b []byte) error {
	var res ChannelTransferTransferBeginCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
