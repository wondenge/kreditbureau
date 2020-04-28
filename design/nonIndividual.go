// Non-Individual Consumer and Account file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var NonIndividual = Type("nonindividual", func() {
	Description("Non-Individual Consumer and Account information")
	TypeName("NonIndividualConsumerResult") // Override generated type name
	ContentType("application/json")         // Override Content-Type header

	Attributes(func() {

	})
})

var NonIndividualResult = ResultType("vnd.goa.nonindividual", func() {
	Description("Non-Individual Consumer and Account information")
	TypeName("NonIndividualConsumerResult") // Override generated type name
	ContentType("application/json")         // Override Content-Type header
	Reference(NonIndividual)

	Attributes(func() {

	})
})
