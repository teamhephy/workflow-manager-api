package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Data data

swagger:model data
*/
type Data struct {

	/* description

	Min Length: 1
	*/
	Description string `json:"description,omitempty"`

	/* fixes

	Min Length: 1
	*/
	Fixes string `json:"fixes,omitempty"`
}

// Validate validates this data
func (m *Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFixes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Data) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateFixes(formats strfmt.Registry) error {

	if swag.IsZero(m.Fixes) { // not required
		return nil
	}

	if err := validate.MinLength("fixes", "body", string(m.Fixes), 1); err != nil {
		return err
	}

	return nil
}
