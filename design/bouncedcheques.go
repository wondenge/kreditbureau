// Bounced Cheque file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var BouncedCheque = Type("bouncedcheques", func() {
	Description("Bounced Cheque information")
	TypeName("BouncedChequeInformationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var BouncedChequeResult = ResultType("vnd.goa.bouncedcheques", func() {
	Description("Bounced Cheque information")
	TypeName("BouncedChequeInformationResult")
	ContentType("application/json")
	Reference(BouncedCheque)

	Attributes(func() {

	})
})
