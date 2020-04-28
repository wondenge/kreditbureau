// Credit Application file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var CreditApplication = Type("creditapplication", func() {
	Description("Credit Application File")
	TypeName("CreditApplicationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var CreditApplicationResult = ResultType("vnd.goa.creditapplication", func() {
	Description("Credit Application File")
	TypeName("CreditApplicationResult")
	ContentType("application/json")
	Reference(CreditApplication)

	Attributes(func() {

	})
})
