// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
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
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Site site
// swagger:model Site
type Site struct {

	// ASN
	// Maximum: 4.294967295e+09
	// Minimum: 1
	Asn *int64 `json:"asn,omitempty"`

	// Circuit count
	// Read Only: true
	CircuitCount int64 `json:"circuit_count,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Contact E-mail
	// Max Length: 254
	// Format: email
	ContactEmail strfmt.Email `json:"contact_email,omitempty"`

	// Contact name
	// Max Length: 50
	ContactName string `json:"contact_name,omitempty"`

	// Contact phone
	// Max Length: 20
	ContactPhone string `json:"contact_phone,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Facility
	// Max Length: 50
	Facility string `json:"facility,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Latitude
	Latitude *string `json:"latitude,omitempty"`

	// Longitude
	Longitude *string `json:"longitude,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Physical address
	// Max Length: 200
	PhysicalAddress string `json:"physical_address,omitempty"`

	// Prefix count
	// Read Only: true
	PrefixCount int64 `json:"prefix_count,omitempty"`

	// Rack count
	// Read Only: true
	RackCount int64 `json:"rack_count,omitempty"`

	// region
	Region *NestedRegion `json:"region,omitempty"`

	// Shipping address
	// Max Length: 200
	ShippingAddress string `json:"shipping_address,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// status
	Status *SiteStatus `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Time zone
	TimeZone string `json:"time_zone,omitempty"`

	// Virtualmachine count
	// Read Only: true
	VirtualmachineCount int64 `json:"virtualmachine_count,omitempty"`

	// Vlan count
	// Read Only: true
	VlanCount int64 `json:"vlan_count,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateAsn(formats strfmt.Registry) error {

	if swag.IsZero(m.Asn) { // not required
		return nil
	}

	if err := validate.MinimumInt("asn", "body", int64(*m.Asn), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("asn", "body", int64(*m.Asn), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_email", "body", string(m.ContactEmail), 254); err != nil {
		return err
	}

	if err := validate.FormatOf("contact_email", "body", "email", m.ContactEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactName(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactName) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_name", "body", string(m.ContactName), 50); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactPhone(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactPhone) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_phone", "body", string(m.ContactPhone), 20); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateFacility(formats strfmt.Registry) error {

	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if err := validate.MaxLength("facility", "body", string(m.Facility), 50); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateName(formats strfmt.Registry) error {

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

func (m *Site) validatePhysicalAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("physical_address", "body", string(m.PhysicalAddress), 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateShippingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_address", "body", string(m.ShippingAddress), 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", string(*m.Slug), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateTags(formats strfmt.Registry) error {

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

func (m *Site) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SiteStatus Status
// swagger:model SiteStatus
type SiteStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this site status
func (m *SiteStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *SiteStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteStatus) UnmarshalBinary(b []byte) error {
	var res SiteStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
