package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetComponentByLatestReleaseParams creates a new GetComponentByLatestReleaseParams object
// with the default values initialized.
func NewGetComponentByLatestReleaseParams() GetComponentByLatestReleaseParams {
	var ()
	return GetComponentByLatestReleaseParams{}
}

// GetComponentByLatestReleaseParams contains all the bound params for the get component by latest release operation
// typically these are obtained from a http.Request
//
// swagger:parameters getComponentByLatestRelease
type GetComponentByLatestReleaseParams struct {
	/*A component is a single deis component, e.g., deis-router
	  Required: true
	  In: path
	*/
	Component string
	/*A train is a release cadence type, e.g., "beta" or "stable"
	  Required: true
	  In: path
	*/
	Train string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetComponentByLatestReleaseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rComponent, rhkComponent, _ := route.Params.GetOK("component")
	if err := o.bindComponent(rComponent, rhkComponent, route.Formats); err != nil {
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

func (o *GetComponentByLatestReleaseParams) bindComponent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Component = raw

	return nil
}

func (o *GetComponentByLatestReleaseParams) bindTrain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Train = raw

	return nil
}
