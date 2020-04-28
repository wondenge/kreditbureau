// Contact Upload File
package design

import (
	. "goa.design/goa/v3/dsl"
)

var ContactUpload = Type("contactupload", func() {
	Description("Contact Upload File")
	TypeName("ContactUploadResult")
	ContentType("application/json")

	Attributes(func() {

	})
})

var ContactUploadResult = ResultType("vnd.goa.contactupload", func() {
	Description("Contact Upload File")
	TypeName("ContactUploadResult")
	ContentType("application/json")
	Reference(ContactUpload)

	Attributes(func() {

	})
})
