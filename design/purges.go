package design

import (
	. "goa.design/goa/v3/dsl"
)

var DeletionRelist = Type("deletionrelist", func() {
	Description("Deletion/Relist API")
	TypeName("DeletionRelist")
	ContentType("application/json")

	// Refer to acceptable options
	// A-Complete
	// B-Default History
	// C-Specific Month Update
	// D- Relist
	Attribute("FunctionType", String, func() {
		Description("Function Type")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("D")
		Meta("rpc:tag", "1")
	})

	Attribute("Surname", String, func() {
		Description("The Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	// 001 - National ID ( Numeric Characters, Between 1-8 characters)
	// 002 - Passport ( Alphanumeric)
	// 003 - Alien Registration ( Numeric characters, Not more than 6 characters)
	// 004 - Service ID ( Numeric Characters, Not more than 6 characters)
	// 005 - Company Registration Number
	Attribute("IDocumentType", Int, func() {
		Description("Identification Document Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "6")
	})

	// The number of the Primary Identification
	// document provided on opening the account.
	Attribute("IDocumentNumber", String, func() {
		Description("Primary ID Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "7")
	})

	// Specific Contract Unique Identifier;
	// or account that contains a credit facility.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "8")
	})

	Attribute("SubmissionDate", String, func() {
		Description("Submission Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "9")
	})

	// The new number of days the account has been in arrears.
	Attribute("DaysInArrears", Int, func() {
		Description("New Number of Days in arrears")
		MinLength(1)
		MaxLength(5)
		Minimum(0)
		Meta("rpc:tag", "10")
	})

})

var StoredDeletionRelist = ResultType("application/vnd.goa.deletionrelist", func() {
	TypeName("StoredDeletionRelist")
	Attributes(func() {
		Extend(DeletionRelist)
		Required(
			"FunctionType", "Surname", "Forename1", "IDocumentType", "IDocumentNumber", "AccountNumber")
	})

	View("default", func() {
		Attribute("FunctionType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("IDocumentType")
		Attribute("IDocumentNumber")
		Attribute("AccountNumber")
		Attribute("SubmissionDate")
		Attribute("DaysInArrears")
	})

	View("mandatory", func() {
		Attribute("FunctionType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("IDocumentType")
		Attribute("IDocumentNumber")
		Attribute("AccountNumber")
	})

})
