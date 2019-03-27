package bookshelf

import (
	"context"
	"log"

	people "github.com/akm/goa_v2_simple_example/gen/people"
)

// people service example implementation.
// The example methods log the requests and return zero values.
type peoplesrvc struct {
	logger *log.Logger
}

// NewPeople returns the people service implementation.
func NewPeople(logger *log.Logger) people.Service {
	return &peoplesrvc{logger}
}

// list
func (s *peoplesrvc) List(ctx context.Context, p *people.ListPayload) (res *people.ListResult, err error) {
	res = &people.ListResult{}
	s.logger.Print("people.list")
	return
}

// create
func (s *peoplesrvc) Create(ctx context.Context, p *people.CreatePayload) (res *people.Person, view string, err error) {
	res = &people.Person{}
	view = "default"
	s.logger.Print("people.create")
	return
}

// show
func (s *peoplesrvc) Show(ctx context.Context, p *people.ShowPayload) (res *people.Person, view string, err error) {
	res = &people.Person{}
	view = "default"
	s.logger.Print("people.show")
	return
}

// update
func (s *peoplesrvc) Update(ctx context.Context, p *people.UpdatePayload) (res *people.PersonPayload, err error) {
	res = &people.PersonPayload{}
	s.logger.Print("people.update")
	return
}

// delete
func (s *peoplesrvc) Delete(ctx context.Context, p *people.DeletePayload) (res *people.Person, view string, err error) {
	res = &people.Person{}
	view = "default"
	s.logger.Print("people.delete")
	return
}
