package design

import (
	. "goa.design/goa/v3/dsl"
)

var LinkDelinkIDs = Type("linkdelinkids", func() {
	Description("Link-Delink IDs API")
	TypeName("LinkDelinkIDs")
	ContentType("application/json")

	// Refer to the following options
	// 001 - Link
	// 002 – Delink
	//
	// Mandatory Field
	// Lookup
	// Alphanumeric
	// Not more than 3 Characters
	Attribute("Function", String, "Function", func() {
		Meta("rpc:tag", "1")
	})

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
	Attribute("IDocumentType1From", String, "Identification Document Type 1 From", func() {
		Meta("rpc:tag", "2")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("IDocumentNumber1From", String, "Identification Document Number 1 From", func() {
		Meta("rpc:tag", "3")
	})

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
	Attribute("IDocumentType2To", String, "Identification Document Type 2 To", func() {
		Meta("rpc:tag", "4")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("IDocumentNumber2To", String, "Identification Document Number 2 To", func() {
		Meta("rpc:tag", "5")
	})
})
