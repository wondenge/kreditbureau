// Delink IDs from an Account File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var Delink = Type("delink", func() {
	Description("Delink IDs from an Account File")
	TypeName("DelinkIDsResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var DelinkResult = ResultType("vnd.goa.delink", func() {
	Description("Delink IDs from an Account File")
	TypeName("DelinkIDsResult")
	ContentType("application/json")
	Reference(Delink)

	Attributes(func() {

	})
})
