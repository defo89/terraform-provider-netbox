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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceInterface writable device interface
// swagger:model WritableDeviceInterface
type WritableDeviceInterface struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	//         Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses string `json:"count_ipaddresses,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// Form factor
	// Read Only: true
	FormFactor string `json:"form_factor,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Parent LAG
	Lag *int64 `json:"lag,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// OOB Management
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Mode
	// Enum: [100 200 300]
	Mode *int64 `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [0 200 800 1000 1120 1130 1150 1170 1050 1100 1200 1300 1310 1320 1350 1400 1420 1500 1510 1650 1520 1550 1600 1700 1750 1800 2600 2610 2620 2630 2640 2810 2820 2830 6100 6200 6300 6400 6500 6600 6700 3010 3020 3040 3080 3160 3320 3400 7010 7020 7030 7040 7050 7060 7070 7080 7090 4000 4010 4040 4050 5000 5050 5100 5150 5200 5300 5310 5320 5330 32767]
	Type int64 `json:"type,omitempty"`

	// Untagged VLAN
	UntaggedVlan *int64 `json:"untagged_vlan,omitempty"`
}

// Validate validates this writable device interface
func (m *WritableDeviceInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceInterface) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

var writableDeviceInterfaceTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeConnectionStatusPropEnum = append(writableDeviceInterfaceTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableDeviceInterface) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

var writableDeviceInterfaceTypeModePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[100,200,300]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeModePropEnum = append(writableDeviceInterfaceTypeModePropEnum, v)
	}
}

// prop value enum
func (m *WritableDeviceInterface) validateModeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

var writableDeviceInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,200,800,1000,1120,1130,1150,1170,1050,1100,1200,1300,1310,1320,1350,1400,1420,1500,1510,1650,1520,1550,1600,1700,1750,1800,2600,2610,2620,2630,2640,2810,2820,2830,6100,6200,6300,6400,6500,6600,6700,3010,3020,3040,3080,3160,3320,3400,7010,7020,7030,7040,7050,7060,7070,7080,7090,4000,4010,4040,4050,5000,5050,5100,5150,5200,5300,5310,5320,5330,32767]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeTypePropEnum = append(writableDeviceInterfaceTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableDeviceInterface) validateTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceInterface) UnmarshalBinary(b []byte) error {
	var res WritableDeviceInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
