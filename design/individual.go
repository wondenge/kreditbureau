// Individual Consumer, Employer and Account file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var Individual = Type("individual", func() {
	Description("Individual Consumer, Employment and Account information")
	TypeName("IndividualConsumerResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var IndividualResult = ResultType("vnd.goa.individual", func() {
	Description("Individual Consumer, Employment and Account information")
	TypeName("IndividualConsumerResult")
	ContentType("application/json")
	Reference(Individual)

	Attributes(func() {

	})
})
