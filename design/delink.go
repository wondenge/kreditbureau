package design

import (
	. "goa.design/goa/v3/dsl"
)

var Delink = Type("delink", func() {
	Description("Delink IDs from an Account API")
	TypeName("Delink")
	ContentType("application/json")

	// The Family or Surname
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "1")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than Max Characters 50
	// Allow Hyphen, Apostrophe.
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "4")
	})

	// The Account Number to Delink from
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "5")
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
	Attribute("PrimaryIDocumentTypeD", String, "Primary Identification Document Type to Delink", func() {
		Meta("rpc:tag", "6")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters,  Not more than 6 characters
	// For Service ID:  Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocumentNumberD", String, "Primary Identification Document Number to Delink", func() {
		Meta("rpc:tag", "7")
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
	Attribute("PrimaryIDocumentTypeL", String, "Primary Identification Document Type to link", func() {
		Meta("rpc:tag", "8")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than characters allowed
	Attribute("PrimaryIDocumentNumberL", String, "Primary identification Document Number to link", func() {
		Meta("rpc:tag", "9")
	})
})

var StoredDelink = ResultType("", func() {
	TypeName("StoredDelink")
	Attributes(func() {
		Extend(Delink)
		Required()
	})
	View("", func() {
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
