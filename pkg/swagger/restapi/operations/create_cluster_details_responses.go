package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
)

/*CreateClusterDetailsOK clusters details response

swagger:response createClusterDetailsOK
*/
type CreateClusterDetailsOK struct {

	// In: body
	Payload *models.ClusterDetail `json:"body,omitempty"`
}

// NewCreateClusterDetailsOK creates CreateClusterDetailsOK with default headers values
func NewCreateClusterDetailsOK() *CreateClusterDetailsOK {
	return &CreateClusterDetailsOK{}
}

// WithPayload adds the payload to the create cluster details o k response
func (o *CreateClusterDetailsOK) WithPayload(payload *models.ClusterDetail) *CreateClusterDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster details o k response
func (o *CreateClusterDetailsOK) SetPayload(payload *models.ClusterDetail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterDetailsOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateClusterDetailsDefault unexpected error

swagger:response createClusterDetailsDefault
*/
type CreateClusterDetailsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateClusterDetailsDefault creates CreateClusterDetailsDefault with default headers values
func NewCreateClusterDetailsDefault(code int) *CreateClusterDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateClusterDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create cluster details default response
func (o *CreateClusterDetailsDefault) WithStatusCode(code int) *CreateClusterDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create cluster details default response
func (o *CreateClusterDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create cluster details default response
func (o *CreateClusterDetailsDefault) WithPayload(payload *models.Error) *CreateClusterDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster details default response
func (o *CreateClusterDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterDetailsDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
