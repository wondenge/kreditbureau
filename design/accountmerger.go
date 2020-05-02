package design

import (
	. "goa.design/goa/v3/dsl"
)

var AccountMerger = Type("accountmerger", func() {
	Description("Account Merger API")
	TypeName("AccountMerger")
	ContentType("application/json")

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Alphanumeric
	// Mandatory
	// Not more than 3 Characters
	Attribute("IDocumentType", String, "Identification Document Type", func() {
		Meta("rpc:tag", "1")
	})

	// The Number of the Primary Identification Document provided on
	// Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("IDocumentNumber", String, "Identification Document Number", func() {
		Meta("rpc:tag", "2")
	})

	// The Family Name or Surname.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("SurnameOrCompanyName", String, "Surname/Company Name", func() {
		Meta("rpc:tag", "3")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe
	Attribute("Forename1OrTradingName", String, "The First Name", func() {
		Meta("rpc:tag", "4")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "5")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "6")
	})

	// Specific Contract Unique Identifier;
	// or account that contains a credit facility.
	//
	// Alphanumeric
	// Mandatory
	// Not more than 20 characters
	Attribute("NewAccountNumber", String, "New Account Number", func() {
		Meta("rpc:tag", "7")
	})

	// Contract Unique Identifier used previously to report this facility.
	//
	// Alphanumeric
	// Mandatory
	// Not more 20 characters
	Attribute("OldAccountNumber", String, "Old Account Number", func() {
		Meta("rpc:tag", "8")
	})
})

var StoredAccountMerger = ResultType("application/vnd.goa.accountmerger", func() {
	TypeName("StoredAccountMerger")
	Attributes(func() {
		Extend(AccountMerger)
		Required(
			"IDocumentType",
			"IDocumentNumber",
			"SurnameOrCompanyName",
			"Forename1OrTradingName",
			"NewAccountNumber",
			"OldAccountNumber",
		)
	})

	View("default", func() {
		Attribute("IDocumentType")
		Attribute("IDocumentNumber")
		Attribute("SurnameOrCompanyName")
		Attribute("Forename1OrTradingName")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("NewAccountNumber")
		Attribute("OldAccountNumber")
	})
})
