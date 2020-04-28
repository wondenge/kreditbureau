// Group Guarantee File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var GroupGuarantee = Type("groupguarantee", func() {
	Description("Group Guarantee File")
	TypeName("GroupGuaranteeResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var GroupGuaranteeResult = ResultType("vnd.goa.groupguarantee", func() {
	Description("Group Guarantee File")
	TypeName("GroupGuaranteeResult")
	ContentType("application/json")
	Reference(GroupGuarantee)

	Attributes(func() {

	})
})
