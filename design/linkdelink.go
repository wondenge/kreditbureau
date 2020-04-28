// Link-Delink IDs File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var LinkDelink = Type("linkdelink", func() {
	Description("Link-Delink IDs File")
	TypeName("LinkDelinkIDsResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var LinkDelinkResult = ResultType("vnd.goa.linkdelink", func() {
	Description("Link-Delink IDs File")
	TypeName("LinkDelinkIDsResult")
	ContentType("application/json")
	Reference(LinkDelink)

	Attributes(func() {

	})
})
