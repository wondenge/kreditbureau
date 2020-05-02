package design

import (
	. "goa.design/goa/v3/dsl"
)

var AccountsMerger = Type("accountsmerger", func() {
	Description("Accounts Merger API")
	TypeName("AccountsMerger")
	ContentType("application/json")

	Attribute("IDocumentType", String, "Identification Document Type", func() {
		Meta("rpc:tag", "1")
	})
	Attribute("IDocumentNumber", String, "Identification Document Number", func() {
		Meta("rpc:tag", "2")
	})
	Attribute("SurnameOrCompanyName", String, "Surname/Company Name", func() {
		Meta("rpc:tag", "3")
	})
	Attribute("Forename1OrTradingName", String, "The First Name", func() {
		Meta("rpc:tag", "4")
	})
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "5")
	})
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "6")
	})
	Attribute("NewAccountNumber", String, "New Account Number", func() {
		Meta("rpc:tag", "7")
	})
	Attribute("OldAccountNumber", String, "Old Account Number", func() {
		Meta("rpc:tag", "8")
	})
})
