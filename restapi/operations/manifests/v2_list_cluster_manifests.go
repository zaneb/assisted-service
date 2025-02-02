// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2ListClusterManifestsHandlerFunc turns a function with the right signature into a v2 list cluster manifests handler
type V2ListClusterManifestsHandlerFunc func(V2ListClusterManifestsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2ListClusterManifestsHandlerFunc) Handle(params V2ListClusterManifestsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2ListClusterManifestsHandler interface for that can handle valid v2 list cluster manifests params
type V2ListClusterManifestsHandler interface {
	Handle(V2ListClusterManifestsParams, interface{}) middleware.Responder
}

// NewV2ListClusterManifests creates a new http.Handler for the v2 list cluster manifests operation
func NewV2ListClusterManifests(ctx *middleware.Context, handler V2ListClusterManifestsHandler) *V2ListClusterManifests {
	return &V2ListClusterManifests{Context: ctx, Handler: handler}
}

/* V2ListClusterManifests swagger:route GET /v2/clusters/{cluster_id}/manifests manifests v2ListClusterManifests

Lists manifests for customizing cluster installation.

*/
type V2ListClusterManifests struct {
	Context *middleware.Context
	Handler V2ListClusterManifestsHandler
}

func (o *V2ListClusterManifests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2ListClusterManifestsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
