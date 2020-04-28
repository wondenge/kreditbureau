// Guarantor file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var Guarantor = Type("guarantor", func() {
	Description("Guarantor Information")
	TypeName("GuarantorInformationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var GuarantorResult = ResultType("vnd.goa.guarantor", func() {
	Description("Guarantor Information")
	TypeName("GuarantorInformationResult")
	ContentType("application/json")
	Reference(Guarantor)

	Attributes(func() {

	})
})
