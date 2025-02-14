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

// V1FindNetworksRequest v1 find networks request
// swagger:model v1.FindNetworksRequest
type V1FindNetworksRequest struct {

	// destinationprefixes
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nat
	// Required: true
	Nat *bool `json:"nat"`

	// parentnetworkid
	// Required: true
	Parentnetworkid *string `json:"parentnetworkid"`

	// partitionid
	// Required: true
	Partitionid *string `json:"partitionid"`

	// prefixes
	// Required: true
	Prefixes []string `json:"prefixes"`

	// primary
	// Required: true
	Primary *bool `json:"primary"`

	// projectid
	// Required: true
	Projectid *string `json:"projectid"`

	// tenantid
	// Required: true
	Tenantid *string `json:"tenantid"`

	// underlay
	// Required: true
	Underlay *bool `json:"underlay"`

	// vrf
	// Required: true
	Vrf *int64 `json:"vrf"`
}

// Validate validates this v1 find networks request
func (m *V1FindNetworksRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationprefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentnetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderlay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FindNetworksRequest) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateParentnetworkid(formats strfmt.Registry) error {

	if err := validate.Required("parentnetworkid", "body", m.Parentnetworkid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validatePartitionid(formats strfmt.Registry) error {

	if err := validate.Required("partitionid", "body", m.Partitionid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validatePrimary(formats strfmt.Registry) error {

	if err := validate.Required("primary", "body", m.Primary); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateTenantid(formats strfmt.Registry) error {

	if err := validate.Required("tenantid", "body", m.Tenantid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

func (m *V1FindNetworksRequest) validateVrf(formats strfmt.Registry) error {

	if err := validate.Required("vrf", "body", m.Vrf); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FindNetworksRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FindNetworksRequest) UnmarshalBinary(b []byte) error {
	var res V1FindNetworksRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
