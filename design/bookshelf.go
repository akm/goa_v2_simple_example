package design

import . "goa.design/goa/dsl"

// API describes the global properties of the API server.
var _ = API("bookshelf", func() {
	Title("Bookshelf Service")
	Description("Support to rent books in your office bookshelf")
})
