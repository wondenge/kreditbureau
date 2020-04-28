// Fraudulent Activities file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var FraudActivities = Type("fraudactivities", func() {
	Description("Fraudulent Activities Information")
	TypeName("FraudulentActivitiesInformationResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var FraudActivitiesResult = ResultType("vnd.goa.fraudactivities", func() {
	Description("Fraudulent Activities Information")
	TypeName("FraudulentActivitiesInformationResult")
	ContentType("application/json")
	Reference(FraudActivities)

	Attributes(func() {

	})
})
