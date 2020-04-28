// Accounts Merger File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var AccountsMerger = Type("accountsmerger", func() {
	Description("Accounts Merger File")
	TypeName("AccountsMergerResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var AccountsMergerResult = ResultType("vnd.goa.accountsmerger", func() {
	Description("Accounts Merger File")
	TypeName("AccountsMergerResult")
	ContentType("application/json")
	Reference(AccountsMerger)

	Attributes(func() {

	})
})
