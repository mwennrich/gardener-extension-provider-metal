// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineAllocation v1 machine allocation
// swagger:model v1.MachineAllocation
type V1MachineAllocation struct {

	// the console password which was generated while provisioning
	ConsolePassword string `json:"console_password,omitempty"`

	// the time when the machine was created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// a description for this machine
	Description string `json:"description,omitempty"`

	// the hostname which will be used when creating the machine
	// Required: true
	Hostname *string `json:"hostname"`

	// the image assigned to this machine
	// Read Only: true
	Image *V1ImageResponse `json:"image,omitempty"`

	// the name of the machine
	// Required: true
	Name *string `json:"name"`

	// the networks of this machine
	// Required: true
	Networks []*V1MachineNetwork `json:"networks"`

	// the project that this machine is assigned to
	// Required: true
	Project *string `json:"project"`

	// the public ssh keys to access the machine with
	// Required: true
	SSHPubKeys []string `json:"ssh_pub_keys"`

	// if the allocation of the machine was successful, this is set to true
	// Required: true
	Succeeded *bool `json:"succeeded"`

	// the tenant that this machine is assigned to
	// Required: true
	Tenant *string `json:"tenant"`

	// userdata to execute post installation tasks
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this v1 machine allocation
func (m *V1MachineAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPubKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSucceeded(formats); err != nil {
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

func (m *V1MachineAllocation) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineAllocation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MachineAllocation) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateSSHPubKeys(formats strfmt.Registry) error {

	if err := validate.Required("ssh_pub_keys", "body", m.SSHPubKeys); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateSucceeded(formats strfmt.Registry) error {

	if err := validate.Required("succeeded", "body", m.Succeeded); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocation) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineAllocation) UnmarshalBinary(b []byte) error {
	var res V1MachineAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
