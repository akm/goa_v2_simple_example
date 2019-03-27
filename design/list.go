package design

import (
	. "goa.design/goa/dsl"
)

const ListTotalCount = "X-Total-Count"

func ListQueryPayload() {
	Attribute("q_start", Int)
	Attribute("q_end", Int)
	Attribute("q_order", String, func() {
		Enum("ASC", "DESC")
	})
	Attribute("q_sort", String)
	Attribute("query", String)
}

// https://github.com/marmelab/react-admin/blob/master/packages/ra-data-json-server/src/index.js#L18
// https://github.com/typicode/json-server#slice
func ListQueryStrings() {
	Param("q_start:_start")
	Param("q_end:_end")
	Param("q_order:_order")
	Param("q_sort:_sort")
	Param("query:q")
}
