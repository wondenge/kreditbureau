package design

import (
	. "goa.design/goa/v3/dsl"
)

var ContactUpload = Type("contactupload", func() {
	Description("Contact Upload API")
	TypeName("ContactUpload")
	ContentType("application/json")

	// The date when the API is submitted to the Bureau.
	//
	// Date Field
	// Mandatory Field
	// Not more than 8 Characters
	// Can’t be in future
	Attribute("SubmissionDate", String, "Submission Date", func() {
		Meta("rpc:tag", "1")
	})

	// Acceptable Values:
	// A – Individual Credit Consumer
	// B – Non-Individual Credit Consumer
	//
	// Lookup
	// Alphanumeric
	// Mandatory Field
	// Not more than 1 character
	Attribute("ClientType", String, "Client Type", func() {
		Meta("rpc:tag", "2")
	})

	// The Family Name or Surname
	// This field is Mandatory if Option “A” is selected in Field 5.4.2
	//
	// Alphanumeric
	// More than 1 character
	// Conditional 5.4.2
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space,
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "3")
	})

	// The First Name
	// This field is Mandatory if Option “A” is selected in Field 5.4.2
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "4")
	})

	// The Given Name
	// This field is Mandatory if Option “A” is selected in Field 5.4.2
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "5")
	})

	// Other Name or Initials
	// This field is Mandatory if Option “A” is selected in Field 5.4.2
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "6")
	})

	// The Name as Registered by the Registrar of Companies
	//
	// Alphanumeric
	// Not more than 70 Characters
	// Conditional to 5.4.2
	Attribute("CompanyName", String, "Company Name", func() {
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
	// Mandatory Field
	// Not more than 3 Characters
	// If Option “B” is selected in option 5.4.2, option “005” must be selected
	Attribute("PrimaryIDocumentType", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "8")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocumentNumber", String, "Primary Identification Document Number", func() {
		Meta("rpc:tag", "9")
	})

	// Account Number of the Loan/Overdraft Facility.
	// This must be a number that exists in the CRB Database.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "10")
	})

	// The Consumers Main Telephone contact Number in the Form of: CCCAAANNNNNNN
	// Where:
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("MobileTelephoneNumber", String, "Mobile Telephone Number", func() {
		Meta("rpc:tag", "11")
	})

	// The Consumer’s Secondary Telephone Contact Number in the Form of: CCCAAANNNNNNN
	// Where:
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("HomeTelephoneNumber", String, "Home Telephone Number", func() {
		Meta("rpc:tag", "12")
	})

	// The Consumer’s Office Telephone Contact Number,
	// if consumer is employed in the Form of: CCCAAANNNNNNN
	// Where:
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("WorkTelephoneNumber", String, "Work Telephone Number", func() {
		Meta("rpc:tag", "13")
	})

	// Consumer’s Postal Address Line 1
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress1", String, "Postal Address 1", func() {
		Meta("rpc:tag", "14")
	})

	// Address Extra Details
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress2", String, "Postal Address 2", func() {
		Meta("rpc:tag", "15")
	})

	// Consumer’s Postal Address Town
	// This field is mandatory if field 5.4.12 or 5.4.13 is filled in.
	//
	// Alphanumeric
	// Not more than 30 Characters
	Attribute("PostalLocationTown", String, "Postal Location Town", func() {
		Meta("rpc:tag", "16")
	})

	// Country of Consumer’s Postal Address
	//
	// Alphanumeric
	// Not more than 2 Characters
	Attribute("PostalLocationCountry", String, "Postal Location Country", func() {
		Meta("rpc:tag", "17")
	})

	// Post Code of Address
	// This field is mandatory if field 5.4.12 or 5.4.13 is filled in.
	//
	// Alphanumeric
	// Not more than 10 characters
	Attribute("PostCode", String, "Post code", func() {
		Meta("rpc:tag", "18")
	})

	// Consumer’s Physical ( Residential Address) Line 1
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress1", String, "Physical Address 1", func() {
		Meta("rpc:tag", "19")
	})

	// Consumer’s Physical (residential Address) Line 2
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress2", String, "Physical Address 2", func() {
		Meta("rpc:tag", "20")
	})

	// Plot Land Ref (LR) No of Consumer’s residential Address
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("PlotNumber", String, "Plot Number", func() {
		Meta("rpc:tag", "21")
	})

	// Town of Consumer’s residential Address
	// This field is mandatory if Field 5.4.17 is filled in
	Attribute("LocationTown", String, "Location Town", func() {
		Meta("rpc:tag", "22")
	})

	// ISO Code of the Country of the Consumer’s residential Address.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 2 characters
	Attribute("LocationCountry", String, "Location Country", func() {
		Meta("rpc:tag", "23")
	})

	// Date When Consumer Moved to the Residential Address.
	//
	// Date Field
	// Not more than 8 characters
	// Date can’t be in Future
	Attribute("PhysicalAddressDate", String, "Date at Physical Address", func() {
		Meta("rpc:tag", "24")
	})

	// The Consumer Work Email Address, If employed or available.
	//
	// Check email format
	// Not more than 50 characters
	// Alphanumeric characters
	Attribute("Email", String, "Email", func() {
		Meta("rpc:tag", "25")
	})

})

var StoredContactUpload = ResultType("", func() {
	TypeName("StoredContactUpload")
	Attributes(func() {
		Extend(ContactUpload)
		Required()
	})
	View("", func() {
		Attribute("SubmissionDate")
		Attribute("ClientType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyName")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("AccountNumber")
		Attribute("MobileTelephoneNumber")
		Attribute("HomeTelephoneNumber")
		Attribute("WorkTelephoneNumber")
		Attribute("PostalAddress1")
		Attribute("PostalAddress2")
		Attribute("PostalLocationTown")
		Attribute("PostalLocationCountry")
		Attribute("PostCode")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalAddress2")
		Attribute("PlotNumber")
		Attribute("LocationTown")
		Attribute("LocationCountry")
		Attribute("PhysicalAddressDate")
		Attribute("Email")

	})
})
