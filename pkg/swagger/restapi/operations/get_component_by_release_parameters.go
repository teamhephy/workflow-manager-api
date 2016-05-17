package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetComponentByReleaseParams creates a new GetComponentByReleaseParams object
// with the default values initialized.
func NewGetComponentByReleaseParams() GetComponentByReleaseParams {
	var ()
	return GetComponentByReleaseParams{}
}

// GetComponentByReleaseParams contains all the bound params for the get component by release operation
// typically these are obtained from a http.Request
//
// swagger:parameters getComponentByRelease
type GetComponentByReleaseParams struct {
	/*A component is a single deis component, e.g., deis-router
	  Required: true
	  In: path
	*/
	Component string
	/*The release version of the deis component, eg., 2.0.0-beta2
	  Required: true
	  In: path
	*/
	Release string
	/*A train is a release cadence type, e.g., "beta" or "stable"
	  Required: true
	  In: path
	*/
	Train string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetComponentByReleaseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rComponent, rhkComponent, _ := route.Params.GetOK("component")
	if err := o.bindComponent(rComponent, rhkComponent, route.Formats); err != nil {
		res = append(res, err)
	}

	rRelease, rhkRelease, _ := route.Params.GetOK("release")
	if err := o.bindRelease(rRelease, rhkRelease, route.Formats); err != nil {
		res = append(res, err)
	}

	rTrain, rhkTrain, _ := route.Params.GetOK("train")
	if err := o.bindTrain(rTrain, rhkTrain, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentByReleaseParams) bindComponent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Component = raw

	return nil
}

func (o *GetComponentByReleaseParams) bindRelease(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Release = raw

	return nil
}

func (o *GetComponentByReleaseParams) bindTrain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Train = raw

	return nil
}
