// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestableReceiver testable receiver
//
// swagger:model testableReceiver
type TestableReceiver struct {

	// discord configs
	DiscordConfigs []string `json:"discord_configs"`

	// group interval
	// Required: true
	GroupInterval *string `json:"group_interval"`

	// group wait
	// Required: true
	GroupWait *string `json:"group_wait"`

	// hook
	// Required: true
	Hook *string `json:"hook"`

	// name
	// Required: true
	Name *string `json:"name"`

	// repeat interval
	// Required: true
	RepeatInterval *string `json:"repeat_interval"`

	// type
	// Required: true
	Type *string `json:"type"`

	// uid
	// Required: true
	UID *string `json:"uid"`
}

// Validate validates this testable receiver
func (m *TestableReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupWait(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepeatInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestableReceiver) validateGroupInterval(formats strfmt.Registry) error {

	if err := validate.Required("group_interval", "body", m.GroupInterval); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateGroupWait(formats strfmt.Registry) error {

	if err := validate.Required("group_wait", "body", m.GroupWait); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateHook(formats strfmt.Registry) error {

	if err := validate.Required("hook", "body", m.Hook); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateRepeatInterval(formats strfmt.Registry) error {

	if err := validate.Required("repeat_interval", "body", m.RepeatInterval); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TestableReceiver) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this testable receiver based on context it is used
func (m *TestableReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestableReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestableReceiver) UnmarshalBinary(b []byte) error {
	var res TestableReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
