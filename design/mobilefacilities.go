// Mobile Facilities File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var MobileFacilities = Type("mobilefacilities", func() {
	Description("Mobile Facilities File")
	TypeName("MobileFacilitiesResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var MobileFacilitiesResult = ResultType("vnd.goa.mobilefacilities", func() {
	Description("Mobile Facilities File")
	TypeName("MobileFacilitiesResult")
	ContentType("application/json")
	Reference(MobileFacilities)

	Attributes(func() {

	})
})
