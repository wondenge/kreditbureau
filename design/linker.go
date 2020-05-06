package design

import (
	. "goa.design/goa/v3/dsl"
)

var LinkDelink = Type("linkdelink", func() {
	Description("Link-Delink IDs API")
	TypeName("LinkDelink")
	ContentType("application/json")

	// 001 - Link
	// 002 – Delink
	Attribute("Function", Int, func() {
		Description("Function")
		Enum(001, 002)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(002)
		Meta("rpc:tag", "1")
	})

	// ID Acceptable Values and Checks.
	// 001 - National ID ( Numeric Characters, Between 1-8 characters)
	// 002 - Passport ( Alphanumeric)
	// 003 - Alien Registration ( Numeric characters, Not more than 6 characters)
	// 004 - Service ID ( Numeric Characters, Not more than 6 characters)
	// 005 - Company Registration Number
	Attribute("IDocumentType1From", Int, func() {
		Description("Identification Document Type 1 From")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "2")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("IDocumentNumber1From", String, func() {
		Description("Identification Document Number 1 From")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "3")
	})

	// ID Acceptable Values and Checks.
	// 001 - National ID ( Numeric Characters, Between 1-8 characters)
	// 002 - Passport ( Alphanumeric)
	// 003 - Alien Registration ( Numeric characters, Not more than 6 characters)
	// 004 - Service ID ( Numeric Characters, Not more than 6 characters)
	// 005 - Company Registration Number
	Attribute("IDocumentType2To", Int, func() {
		Description("Identification Document Type 2 To")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "4")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("IDocumentNumber2To", String, func() {
		Description("Identification Document Number 2 To")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "5")
	})
})

var StoredLinkDelink = ResultType("application/vnd.goa.linkdelink", func() {
	TypeName("StoredLinkDelink")
	Attributes(func() {
		Extend(LinkDelink)
		Required(
			"Function", "IDocumentType1From", "IDocumentNumber1From", "IDocumentType2To", "IDocumentNumber2To")
	})

	View("default", func() {
		Attribute("Function")
		Attribute("IDocumentType1From")
		Attribute("IDocumentNumber1From")
		Attribute("IDocumentType2To")
		Attribute("IDocumentNumber2To")
	})

	View("mandatory", func() {
		Attribute("Function")
		Attribute("IDocumentType1From")
		Attribute("IDocumentNumber1From")
		Attribute("IDocumentType2To")
		Attribute("IDocumentNumber2To")
	})
})
