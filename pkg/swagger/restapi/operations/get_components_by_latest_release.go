package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
	"github.com/go-swagger/go-swagger/errors"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// GetComponentsByLatestReleaseHandlerFunc turns a function with the right signature into a get components by latest release handler
type GetComponentsByLatestReleaseHandlerFunc func(GetComponentsByLatestReleaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetComponentsByLatestReleaseHandlerFunc) Handle(params GetComponentsByLatestReleaseParams) middleware.Responder {
	return fn(params)
}

// GetComponentsByLatestReleaseHandler interface for that can handle valid get components by latest release params
type GetComponentsByLatestReleaseHandler interface {
	Handle(GetComponentsByLatestReleaseParams) middleware.Responder
}

// NewGetComponentsByLatestRelease creates a new http.Handler for the get components by latest release operation
func NewGetComponentsByLatestRelease(ctx *middleware.Context, handler GetComponentsByLatestReleaseHandler) *GetComponentsByLatestRelease {
	return &GetComponentsByLatestRelease{Context: ctx, Handler: handler}
}

/*GetComponentsByLatestRelease swagger:route POST /versions/latest getComponentsByLatestRelease

list the latest release version of the components

*/
type GetComponentsByLatestRelease struct {
	Context *middleware.Context
	Handler GetComponentsByLatestReleaseHandler
}

func (o *GetComponentsByLatestRelease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetComponentsByLatestReleaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetComponentsByLatestReleaseBody get components by latest release body

swagger:model GetComponentsByLatestReleaseBody
*/
type GetComponentsByLatestReleaseBody struct {

	/* data
	 */
	Data []*models.ComponentVersion `json:"data,omitempty"`
}

// Validate validates this get components by latest release body
func (o *GetComponentsByLatestReleaseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*GetComponentsByLatestReleaseOKBodyBody get components by latest release o k body body

swagger:model GetComponentsByLatestReleaseOKBodyBody
*/
type GetComponentsByLatestReleaseOKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.ComponentVersion `json:"data"`
}

// Validate validates this get components by latest release o k body body
func (o *GetComponentsByLatestReleaseOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseOKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getComponentsByLatestReleaseOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
