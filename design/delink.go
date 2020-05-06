package design

import (
	. "goa.design/goa/v3/dsl"
)

var Delink = Type("delink", func() {
	Description("Delink IDs from an Account API")
	TypeName("Delink")
	ContentType("application/json")

	Attribute("Surname", String, func() {
		Description("Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "1")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Attribute("AccountNumber", String, func() {
		Description("Account Number to Delink from")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "5")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocumentTypeD", Int, func() {
		Description("Primary Identification Document Type to Delink")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "6")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumberD", String, func() {
		Description("Primary Identification Document Number to Delink")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "7")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocumentTypeL", Int, func() {
		Description("Primary Identification Document Type to link")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "8")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumberL", String, func() {
		Description("Primary identification Document Number to link")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "9")
	})
})

var StoredDelink = ResultType("application/vnd.goa.delink", func() {
	TypeName("StoredDelink")
	Attributes(func() {
		Extend(Delink)
		Required(
			"Surname", "Forename1", "AccountNumber", "PrimaryIDocumentTypeD", "PrimaryIDocumentNumberD", "PrimaryIDocumentTypeL", "PrimaryIDocumentNumberL")
	})
	View("default", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocumentTypeD")
		Attribute("PrimaryIDocumentNumberD")
		Attribute("PrimaryIDocumentTypeL")
		Attribute("PrimaryIDocumentNumberL")
	})
	View("mandatory", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocumentTypeD")
		Attribute("PrimaryIDocumentNumberD")
		Attribute("PrimaryIDocumentTypeL")
		Attribute("PrimaryIDocumentNumberL")
	})
})
