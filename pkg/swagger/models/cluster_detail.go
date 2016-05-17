package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ClusterDetail cluster detail

swagger:model clusterDetail
*/
type ClusterDetail struct {

	/* components

	Required: true
	*/
	Components []*ComponentDetail `json:"components"`

	/* first seen
	 */
	FirstSeen *strfmt.DateTime `json:"firstSeen,omitempty"`

	/* id

	Required: true
	Min Length: 1
	*/
	ID string `json:"id"`

	/* last seen
	 */
	LastSeen *strfmt.DateTime `json:"lastSeen,omitempty"`
}

// Validate validates this cluster detail
func (m *ClusterDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDetail) validateComponents(formats strfmt.Registry) error {

	if err := validate.Required("components", "body", m.Components); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {

			if err := m.Components[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDetail) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}
