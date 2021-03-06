// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Daily daily
// swagger:model daily
type Daily struct {

	// counter
	// Read Only: true
	Counter int64 `json:"counter,omitempty"`

	// date
	// Min Length: 1
	Date string `json:"date,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// message
	// Required: true
	// Min Length: 1
	Message *string `json:"message"`

	// pray
	// Required: true
	// Min Length: 1
	Pray *string `json:"pray"`

	// title
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`

	// verse
	// Required: true
	// Min Length: 1
	Verse *string `json:"verse"`
}

// Validate validates this daily
func (m *Daily) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePray(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Daily) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.MinLength("date", "body", string(m.Date), 1); err != nil {
		return err
	}

	return nil
}

func (m *Daily) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", string(*m.Message), 1); err != nil {
		return err
	}

	return nil
}

func (m *Daily) validatePray(formats strfmt.Registry) error {

	if err := validate.Required("pray", "body", m.Pray); err != nil {
		return err
	}

	if err := validate.MinLength("pray", "body", string(*m.Pray), 1); err != nil {
		return err
	}

	return nil
}

func (m *Daily) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", string(*m.Title), 1); err != nil {
		return err
	}

	return nil
}

func (m *Daily) validateVerse(formats strfmt.Registry) error {

	if err := validate.Required("verse", "body", m.Verse); err != nil {
		return err
	}

	if err := validate.MinLength("verse", "body", string(*m.Verse), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Daily) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Daily) UnmarshalBinary(b []byte) error {
	var res Daily
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
