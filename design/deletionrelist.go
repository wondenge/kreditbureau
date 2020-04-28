package design

import (
	. "goa.design/goa/v3/dsl"
)

// Deletion/Relist File
//
// This file should be used whenever the default history is to be deleted.
// This file will be used by lenders to delete Default History or
// Relist Deleted historical data of a customer.
// The file will be used in case there is a court order or an agreement
// of the 2 parties i.e. Customer and Lender to delete default history.

var CustomerHistory = Type("customerhistory", func() {
	Description("Used by lenders to delete default history or relist deleted historical data of a customer")
	TypeName("CustomerHistory")
	ContentType("application/json")

	// 5.8.1
	// Mandatory for FIs and KDIC
	// Refer to LOV32
	Attribute("Function Type", String, "The type of the intended deletion", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Enum("A", "B", "C", "D")
		Meta("rpc:tag", "1")
	})

	// 5.8.2.
	// Mandatory for FIs and KDIC
	Attribute("Surname", String, "The Family Name or Surname", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "2")
	})

	// 5.8.3
	// Mandatory for FIs and KDIC
	Attribute("Forename 1", String, "The First Name", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "3")
	})

	// 5.8.4
	// Optional for FIs and KDIC
	Attribute("Forename 2", String, "The Given Name", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "4")
	})

	// 5.8.5
	// Optional for FIs and KDIC
	Attribute("Forename 3", String, "Other Name or Initials", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "5")
	})

	// 5.8.6
	// Mandatory for FIs and KDIC
	Attribute("ID Type", String, "The Type of Identification document provided on Opening the Account", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "6")
	})

	// 5.8.7
	// Mandatory for FIs and KDIC
	Attribute("ID Number", String, "The number of the Primary Identification document provided on opening the account", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "7")
	})

	// 5.8.8
	// Mandatory for FIs and KDIC
	Attribute("Account Number", String, "Specific Contract Unique Identifier; or account that contains a credit facility", func() {
		Pattern(`(?i)[a-z '\.]+`)
		Meta("rpc:tag", "8")
	})

	// 5.8.9
	// Optional for FIs and KDIC
	Attribute("Submission Date", String, "The date the record is to be effected.", func() {
		Meta("rpc:tag", "9")
	})

	// 5.8.10
	// Optional for FIs and KDIC
	// This should be rounded to 20days
	// This option is only applicable if Option B, "Default History" is selected in field 5.8.1
	Attribute("New Days in Arrears", String, "The new number of days the account has been in arrears.", func() {
		Meta("rpc:tag", "10")
	})

	// Explicitly define default view
	View("default", func() {
		Attribute("Function Type")
		Attribute("Surname")
		Attribute("Forename 1")
		Attribute("Forename 2")
		Attribute("Forename 3")
		Attribute("ID Type")
		Attribute("ID Number")
		Attribute("Account Number")
		Attribute("Submission Date")
		Attribute("New Days in Arrears")
	})

	Required("Function Type", "Surname", "Forename 1", "ID Type", "ID Number", "Account Number")

})

var StoredCustomerHistory = ResultType("application/vnd.goa.customerhistory", func() {
	Description("")
	TypeName("StoredCustomerHistory")
	ContentType("application/json")

	Attributes(func() {
		Extend(CustomerHistory)
		Required("Function Type", "Surname", "Forename 1", "ID Type", "ID Number", "Account Number")
	})

	// Explicitly define default view
	View("default", func() {
		Attribute("Function Type")
		Attribute("Surname")
		Attribute("Forename 1")
		Attribute("Forename 2")
		Attribute("Forename 3")
		Attribute("ID Type")
		Attribute("ID Number")
		Attribute("Account Number")
		Attribute("Submission Date")
		Attribute("New Days in Arrears")
	})
})
