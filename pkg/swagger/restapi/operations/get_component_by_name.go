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
)

// GetComponentByNameHandlerFunc turns a function with the right signature into a get component by name handler
type GetComponentByNameHandlerFunc func(GetComponentByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetComponentByNameHandlerFunc) Handle(params GetComponentByNameParams) middleware.Responder {
	return fn(params)
}

// GetComponentByNameHandler interface for that can handle valid get component by name params
type GetComponentByNameHandler interface {
	Handle(GetComponentByNameParams) middleware.Responder
}

// NewGetComponentByName creates a new http.Handler for the get component by name operation
func NewGetComponentByName(ctx *middleware.Context, handler GetComponentByNameHandler) *GetComponentByName {
	return &GetComponentByName{Context: ctx, Handler: handler}
}

/*GetComponentByName swagger:route GET /versions/{train}/{component} getComponentByName

list the releases of a component

*/
type GetComponentByName struct {
	Context *middleware.Context
	Handler GetComponentByNameHandler
}

func (o *GetComponentByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetComponentByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetComponentByNameOKBodyBody get component by name o k body body

swagger:model GetComponentByNameOKBodyBody
*/
type GetComponentByNameOKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.ComponentDetail `json:"data"`
}

// Validate validates this get component by name o k body body
func (o *GetComponentByNameOKBodyBody) Validate(formats strfmt.Registry) error {
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

func (o *GetComponentByNameOKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getComponentByNameOK"+"."+"data", "body", o.Data); err != nil {
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