// Daily Payment Information File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var DailyPayments = Type("dailypayments", func() {
	Description("Daily Payment Information File")
	TypeName("DailyPaymentInformationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var DailyPaymentsResult = ResultType("vnd.goa.dailypayments", func() {
	Description("Daily Payment Information File")
	TypeName("DailyPaymentInformationResult")
	ContentType("application/json")
	Reference(DailyPayments)

	Attributes(func() {

	})
})
