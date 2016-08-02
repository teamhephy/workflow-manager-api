package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Namespace namespace

swagger:model namespace
*/
type Namespace struct {

	/* daemon sets

	Required: true
	*/
	DaemonSets []*K8sResource `json:"daemonSets"`

	/* deployments

	Required: true
	*/
	Deployments []*K8sResource `json:"deployments"`

	/* events

	Required: true
	*/
	Events []*K8sResource `json:"events"`

	/* name

	Required: true
	*/
	Name string `json:"name"`

	/* pods

	Required: true
	*/
	Pods []*K8sResource `json:"pods"`

	/* replica sets

	Required: true
	*/
	ReplicaSets []*K8sResource `json:"replicaSets"`

	/* replication controllers

	Required: true
	*/
	ReplicationControllers []*K8sResource `json:"replicationControllers"`

	/* services

	Required: true
	*/
	Services []*K8sResource `json:"services"`
}

// Validate validates this namespace
func (m *Namespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaemonSets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeployments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplicaSets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplicationControllers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Namespace) validateDaemonSets(formats strfmt.Registry) error {

	if err := validate.Required("daemonSets", "body", m.DaemonSets); err != nil {
		return err
	}

	for i := 0; i < len(m.DaemonSets); i++ {

		if m.DaemonSets[i] != nil {

			if err := m.DaemonSets[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateDeployments(formats strfmt.Registry) error {

	if err := validate.Required("deployments", "body", m.Deployments); err != nil {
		return err
	}

	for i := 0; i < len(m.Deployments); i++ {

		if m.Deployments[i] != nil {

			if err := m.Deployments[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateEvents(formats strfmt.Registry) error {

	if err := validate.Required("events", "body", m.Events); err != nil {
		return err
	}

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if err := m.Events[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Namespace) validatePods(formats strfmt.Registry) error {

	if err := validate.Required("pods", "body", m.Pods); err != nil {
		return err
	}

	for i := 0; i < len(m.Pods); i++ {

		if m.Pods[i] != nil {

			if err := m.Pods[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateReplicaSets(formats strfmt.Registry) error {

	if err := validate.Required("replicaSets", "body", m.ReplicaSets); err != nil {
		return err
	}

	for i := 0; i < len(m.ReplicaSets); i++ {

		if m.ReplicaSets[i] != nil {

			if err := m.ReplicaSets[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateReplicationControllers(formats strfmt.Registry) error {

	if err := validate.Required("replicationControllers", "body", m.ReplicationControllers); err != nil {
		return err
	}

	for i := 0; i < len(m.ReplicationControllers); i++ {

		if m.ReplicationControllers[i] != nil {

			if err := m.ReplicationControllers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Namespace) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {

			if err := m.Services[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
