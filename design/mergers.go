package design

import (
	. "goa.design/goa/v3/dsl"
)

var AccountMerger = Type("accountmerger", func() {
	Description("Account Merger API")
	TypeName("AccountMerger")
	ContentType("application/json")

	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	Attribute("IDocumentType", Int, func() {
		Description("Identification Document Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "1")
	})

	// ID Acceptable Values and Checks.
	// 001 - National ID ( Numeric Characters, Between 1-8 characters)
	// 002 - Passport ( Alphanumeric)
	// 003 - Alien Registration ( Numeric characters, Not more than 6 characters)
	// 004 - Service ID ( Numeric Characters, Not more than 6 characters)
	// 005 - Company Registration Number
	Attribute("IDocumentNumber", String, func() {
		Description("Primary ID Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "2")
	})

	Attribute("Surname", String, func() {
		Description("The Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "6")
	})

	// Specific Contract Unique Identifier;
	// or account that contains a credit facility.
	Attribute("NewAccountNumber", String, func() {
		Description("New Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "7")
	})

	// Contract Unique Identifier used previously
	// to report this facility.
	Attribute("OldAccountNumber", String, func() {
		Description("Old Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "8")
	})
})

var StoredAccountMerger = ResultType("application/vnd.goa.accountmerger", func() {
	TypeName("StoredAccountMerger")
	Attributes(func() {
		Extend(AccountMerger)
		Required("IDocumentType",
			"IDocumentNumber",
			"Surname",
			"Forename",
			"NewAccountNumber",
			"OldAccountNumber",
		)
	})

	View("default", func() {
		Attribute("IDocumentType")
		Attribute("IDocumentNumber")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("NewAccountNumber")
		Attribute("OldAccountNumber")
	})
	View("mandatory", func() {
		Attribute("IDocumentType")
		Attribute("IDocumentNumber")
		Attribute("Surname")
		Attribute("Forename")
		Attribute("NewAccountNumber")
		Attribute("OldAccountNumber")
	})
})
