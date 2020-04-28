// Collateral Register file
package design

import (
	. "goa.design/goa/v3/dsl"
)

var CollateralRegister = Type("collateralregister", func() {
	Description("Collateral Register")
	TypeName("CollateralRegisterResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var CollateralRegisterResult = ResultType("vnd.goa.collateralregister", func() {
	Description("Collateral Register")
	TypeName("CollateralRegisterResult")
	ContentType("application/json")
	Reference(CollateralRegister)

	Attributes(func() {

	})
})
