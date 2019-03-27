package design

// import . "goa.design/goa/dsl"

// var BookPayload = Type("BookPayload", func() {
// 	Attribute("owner_name", String)
// 	Attribute("title", String)
// 	Attribute("author_names", ArrayOf(String))
// 	Attribute("publisher_name", String)
// 	Attribute("tags", ArrayOf(String))
// 	Required("title", "owner")
// })

// var Book = ResultType("application/vnd.book+json", func() {
// 	Attributes(func() {
// 		Attribute("owner", RefInt64)
// 		Attribute("title", String)
// 		Attribute("authors", ArrayOf(RefInt64))
// 		Attribute("publisher", RefInt64)
// 		Attribute("tags", ArrayOf(String))
// 	})
// 	View("default", func() {
// 		Attribute("id")
// 		Attribute("owner")
// 		Attribute("title")
// 	})
// 	View("full", func() {
// 		Attribute("id")
// 		Attribute("owner")
// 		Attribute("title")
// 		Attribute("authors")
// 		Attribute("publisher")
// 		Attribute("tags")
// 	})
// })
