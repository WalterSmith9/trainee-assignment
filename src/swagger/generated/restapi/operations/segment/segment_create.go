// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SegmentCreateHandlerFunc turns a function with the right signature into a segment create handler
type SegmentCreateHandlerFunc func(SegmentCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SegmentCreateHandlerFunc) Handle(params SegmentCreateParams) middleware.Responder {
	return fn(params)
}

// SegmentCreateHandler interface for that can handle valid segment create params
type SegmentCreateHandler interface {
	Handle(SegmentCreateParams) middleware.Responder
}

// NewSegmentCreate creates a new http.Handler for the segment create operation
func NewSegmentCreate(ctx *middleware.Context, handler SegmentCreateHandler) *SegmentCreate {
	return &SegmentCreate{Context: ctx, Handler: handler}
}

/*
	SegmentCreate swagger:route POST /segment Segment segmentCreate

# Create new segment

Accepts information about segment and stores it
*/
type SegmentCreate struct {
	Context *middleware.Context
	Handler SegmentCreateHandler
}

func (o *SegmentCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSegmentCreateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}