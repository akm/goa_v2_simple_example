// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// people endpoints
//
// Command:
// $ goa gen github.com/akm/goa_v2_simple_example/design

package people

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "people" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Create goa.Endpoint
	Show   goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "people" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Create: NewCreateEndpoint(s),
		Show:   NewShowEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Delete: NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "people" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Create = m(e.Create)
	e.Show = m(e.Show)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "people".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "people".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		res, view, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPerson(res, view)
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "people".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPerson(res, view)
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "people".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "people".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		res, view, err := s.Delete(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPerson(res, view)
		return vres, nil
	}
}
