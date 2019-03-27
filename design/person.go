package design

import . "goa.design/goa/dsl"

var PersonPayload = Type("PersonPayload", func() {
	Attribute("name", String)
	Attribute("memo", String)
	Required("name")
})

var Person = ResultType("application/vnd.person+json", func() {
	Attributes(func() {
		Attribute("id", Int64)
		Attribute("name", String)
		Attribute("memo", String)
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
	View("full", func() {
		Attribute("id")
		Attribute("name")
		Attribute("memo")
	})
})

var _ = Service("people", func() {
	HTTP(func() {
		Path("/people")
	})

	Error("not_found", ErrorResult, "Person not found")
	Error("validation_error", ErrorResult, "Person Validation Error")

	Method("list", func() {
		Description("list")
		Payload(func() {
			ListQueryPayload()
		})
		Result(func() {
			Attribute("people", ArrayOf(Person))
			Attribute(ListTotalCount, String)
			Required("people", ListTotalCount)
		})
		HTTP(func() {
			GET("")
			// ?_end=10&_order=DESC&_sort=id&_start=0
			ListQueryStrings()
			Response(StatusOK, func() {
				Header(ListTotalCount)
				Body("people")
			})
		})
	})

	Method("create", func() {
		Description("create")
		Payload(func() {
			Attribute("person", PersonPayload)
			Required("person")
		})
		Result(Person)
		HTTP(func() {
			POST("")
			Body("person")
			Response(StatusCreated)
			Response("validation_error", StatusBadRequest)
		})
	})

	Method("show", func() {
		Description("show")
		Payload(func() {
			Attribute("id", Int64)
			Required("id")
		})
		Result(Person)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("update", func() {
		Description("update")
		Payload(func() {
			Attribute("id", Int64)
			Attribute("person", PersonPayload)
			Required("id", "person")
		})
		Result(PersonPayload)
		HTTP(func() {
			PUT("/{id}")
			Body("person")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("validation_error", StatusBadRequest)
		})
	})

	Method("delete", func() {
		Description("delete")
		Payload(func() {
			Attribute("id", Int64)
			Required("id")
		})
		Result(Person)
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})
