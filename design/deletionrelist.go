package design

import (
	. "goa.design/goa/v3/dsl"
)

// Deletion/Relist API
//
// This API should be used whenever the default history is to be deleted.
// This API will be used by lenders to delete Default History or
// Relist Deleted historical data of a customer.
// The API will be used in case there is a court order or an agreement
// of the 2 parties i.e. Customer and Lender to delete default history.

var DeletionRelist = Type("deletionrelist", func() {
	Description("Deletion/Relist API")
	TypeName("DeletionRelist")
	ContentType("application/json")

	// Refer to acceptable options
	// A-Complete
	// B-Default History
	// C-Specific Month Update
	// D- Relist
	//
	// Mandatory
	// Lookup
	// Not more than 1 character
	Attribute("FunctionType", String, "Function Type", func() {
		Meta("rpc:tag", "1")
	})

	// The Family Name or Surname.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Surname", String, "The Family Name or Surname", func() {
		Meta("rpc:tag", "2")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "3")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "4")
	})

	// Other Name or Initials.
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "5")
	})

	// The Type of Identification document Provided on Opening the Account.
	//
	// Alphanumeric
	// Mandatory
	// Not more than 3 Characters
	Attribute("IDocumentType", String, "Identification Document Type", func() {
		Meta("rpc:tag", "6")
	})

	// The number of the Primary Identification document provided on opening the account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("IDocumentNumber", String, "Identification Document Number", func() {
		Meta("rpc:tag", "7")
	})

	// Specific Contract Unique Identifier; or account that contains a credit facility.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "8")
	})

	// The date the record is to be affected.
	//
	// Date Field
	// Not more than 8 Characters
	Attribute("SubmissionDate", String, "Submission Date", func() {
		Meta("rpc:tag", "9")
	})

	// The new number of days the account has been in arrears.
	// This should be rounded to 20 days.
	// This option is only applicable if Option B is selected in field 5.8.1
	//
	// Number Field
	// Not more than 5 characters
	Attribute("DaysInArrears", String, "New Number of Days in arrears", func() {
		Meta("rpc:tag", "10")
	})

})
