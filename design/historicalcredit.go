// Historical Credit Information Update File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var HistoricalCredit = Type("historicalcredit", func() {
	Description("Historical Credit Information Update File")
	TypeName("HistoricalCreditInformationUpdateResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var HistoricalCreditResult = ResultType("vnd.goa.historicalcredit", func() {
	Description("Historical Credit Information Update File")
	TypeName("HistoricalCreditInformationUpdateResult")
	ContentType("application/json")
	Reference(HistoricalCredit)

	Attributes(func() {

	})
})
