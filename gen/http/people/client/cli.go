// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// people HTTP client CLI support package
//
// Command:
// $ goa gen github.com/akm/goa_v2_simple_example/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	people "github.com/akm/goa_v2_simple_example/gen/people"
)

// BuildListPayload builds the payload for the people list endpoint from CLI
// flags.
func BuildListPayload(peopleListQStart string, peopleListQEnd string, peopleListQOrder string, peopleListQSort string, peopleListQuery string) (*people.ListPayload, error) {
	var err error
	var qStart *int
	{
		if peopleListQStart != "" {
			var v int64
			v, err = strconv.ParseInt(peopleListQStart, 10, 64)
			val := int(v)
			qStart = &val
			if err != nil {
				err = fmt.Errorf("invalid value for qStart, must be INT")
			}
		}
	}
	var qEnd *int
	{
		if peopleListQEnd != "" {
			var v int64
			v, err = strconv.ParseInt(peopleListQEnd, 10, 64)
			val := int(v)
			qEnd = &val
			if err != nil {
				err = fmt.Errorf("invalid value for qEnd, must be INT")
			}
		}
	}
	var qOrder *string
	{
		if peopleListQOrder != "" {
			qOrder = &peopleListQOrder
		}
	}
	var qSort *string
	{
		if peopleListQSort != "" {
			qSort = &peopleListQSort
		}
	}
	var query *string
	{
		if peopleListQuery != "" {
			query = &peopleListQuery
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &people.ListPayload{
		QStart: qStart,
		QEnd:   qEnd,
		QOrder: qOrder,
		QSort:  qSort,
		Query:  query,
	}
	return payload, nil
}

// BuildCreatePayload builds the payload for the people create endpoint from
// CLI flags.
func BuildCreatePayload(peopleCreateBody string) (*people.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(peopleCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"memo\": \"Laudantium quae non voluptas.\",\n      \"name\": \"Ex quidem fuga nesciunt.\"\n   }'")
		}
	}
	if err != nil {
		return nil, err
	}
	v := &people.PersonPayload{
		Name: body.Name,
		Memo: body.Memo,
	}
	res := &people.CreatePayload{
		Person: v,
	}
	return res, nil
}

// BuildShowPayload builds the payload for the people show endpoint from CLI
// flags.
func BuildShowPayload(peopleShowID string) (*people.ShowPayload, error) {
	var id int64
	{
		id, err = strconv.ParseInt(peopleShowID, 10, 64)
	}
	payload := &people.ShowPayload{
		ID: id,
	}
	return payload, nil
}

// BuildUpdatePayload builds the payload for the people update endpoint from
// CLI flags.
func BuildUpdatePayload(peopleUpdateBody string, peopleUpdateID string) (*people.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(peopleUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"memo\": \"Quos nihil debitis.\",\n      \"name\": \"Accusantium officia nam qui eum ex.\"\n   }'")
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(peopleUpdateID, 10, 64)
	}
	if err != nil {
		return nil, err
	}
	v := &people.PersonPayload{
		Name: body.Name,
		Memo: body.Memo,
	}
	res := &people.UpdatePayload{
		Person: v,
	}
	res.ID = id
	return res, nil
}

// BuildDeletePayload builds the payload for the people delete endpoint from
// CLI flags.
func BuildDeletePayload(peopleDeleteID string) (*people.DeletePayload, error) {
	var id int64
	{
		id, err = strconv.ParseInt(peopleDeleteID, 10, 64)
	}
	payload := &people.DeletePayload{
		ID: id,
	}
	return payload, nil
}