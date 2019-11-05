// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerOutletTemplate writable power outlet template
// swagger:model WritablePowerOutletTemplate
type WritablePowerOutletTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Feed leg
	//
	// Phase (for three-phase feeds)
	// Enum: [1 2 3]
	FeedLeg *int64 `json:"feed_leg,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// power port
	PowerPort *PowerPortTemplate `json:"power_port,omitempty"`
}

// Validate validates this writable power outlet template
func (m *WritablePowerOutletTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerOutletTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

var writablePowerOutletTemplateTypeFeedLegPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTemplateTypeFeedLegPropEnum = append(writablePowerOutletTemplateTypeFeedLegPropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerOutletTemplate) validateFeedLegEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writablePowerOutletTemplateTypeFeedLegPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutletTemplate) validateFeedLeg(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeedLegEnum("feed_leg", "body", *m.FeedLeg); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validatePowerPort(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if m.PowerPort != nil {
		if err := m.PowerPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerOutletTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerOutletTemplate) UnmarshalBinary(b []byte) error {
	var res WritablePowerOutletTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
