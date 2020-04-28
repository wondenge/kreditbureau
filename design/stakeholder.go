// Stakeholder file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var Stakeholder = Type("stakeholder", func() {
	Description("Stakeholder Information")
	TypeName("StakeholderInformationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var StakeholderResult = ResultType("vnd.goa.stakeholder", func() {
	Description("Stakeholder Information")
	TypeName("StakeholderInformationResult")
	ContentType("application/json")
	Reference(Stakeholder)

	Attributes(func() {

	})
})
