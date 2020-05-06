package design

import (
	. "goa.design/goa/v3/dsl"
)

var ContactUpload = Type("contactupload", func() {
	Description("Contact Upload API")
	TypeName("ContactUpload")
	ContentType("application/json")

	// The date when the API is submitted to the Bureau.
	Attribute("SubmissionDate", String, func() {
		Description("Date when API was submitted to the Bureau")
		Format(FormatDate)
		MaxLength(8)
		Meta("rpc:tag", "1")
	})

	// ClientTypeA – Individual Credit Consumer
	// ClientTypeB – Non-Individual Credit Consumer
	Attribute("ClientType", ArrayOf(Enum), func() {
		Enum(ContactUploadCTA,
			ContactUploadCTB,
		)
		Meta("rpc:tag", "2")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocumentType", Int, func() {
		Description("Primary Identification Document Type")
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
	Attribute("PrimaryIDocumentNumber", String, func() {
		Description("Primary ID provided on Opening of the Account")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "9")
	})

	// This must be a number that exists in the CRB Database.
	Attribute("AccountNumber", String, func() {
		Description("Account Number of the Loan/Overdraft Facility")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "10")
	})

	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("MobileTelephoneNumber", String, func() {
		Description("Mobile Telephone Number") // Main Telephone contact Number ( CCCAAANNNNNNN )
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "11")
	})

	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	// Secondary Telephone Contact Number ( CCCAAANNNNNNN )
	Attribute("HomeTelephoneNumber", String, func() {
		Description("Home Telephone Number")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "12")
	})

	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	// Office Telephone Contact Number ( CCCAAANNNNNNN )
	Attribute("WorkTelephoneNumber", String, func() {
		Description("Work Telephone Number")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "13")
	})

	Attribute("PostalAddress1", String, func() {
		Description("Postal Address 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "14")
	})

	Attribute("PostalAddress2", String, func() {
		Description("Postal Address 2")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "15")
	})

	// This field is mandatory if field 5.4.12 or 5.4.13 is filled in.
	Attribute("PostalLocationTown", String, func() {
		Description("Postal Location Town")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "16")
	})

	Attribute("PostalLocationCountry", String, func() {
		Description("Postal Location Country")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "17")
	})

	// This field is mandatory if field 5.4.12 or 5.4.13 is filled in.
	Attribute("PostCode", String, func() {
		Description("Post Code of Address")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "18")
	})

	Attribute("PhysicalAddress1", String, func() {
		Description("Physical ( Residential Address) Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "19")
	})

	Attribute("PhysicalAddress2", String, func() {
		Description("Physical (residential Address) Line 2")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "20")
	})

	Attribute("PlotNumber", String, func() {
		Description("Plot Land Ref (LR) No")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "21")
	})

	// Town of Consumer’s residential Address
	// This field is mandatory if Field 5.4.17 is filled in
	Attribute("LocationTown", String, func() {
		Description("Location Town")
		Meta("rpc:tag", "22")
	})

	// ISO Code of the Country of the
	// Consumer’s residential Address.
	Attribute("LocationCountry", String, func() {
		Description("Location Country")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "23")
	})

	Attribute("PhysicalAddressDate", String, func() {
		Description("Date moved at Physical Address")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "24")
	})

	Attribute("Email", String, "Email", func() {
		Description("Consumer Work Email Address")
		MinLength(1)
		MaxLength(50)
		Format(FormatEmail)
		Meta("rpc:tag", "25")
	})

})

var StoredContactUpload = ResultType("application/vnd.goa.contactupload", func() {
	TypeName("StoredContactUpload")
	Attributes(func() {
		Extend(ContactUpload)
		Required(
			"SubmissionDate",
			"ClientType",
			"PrimaryIDocumentType",
			"PrimaryIDocumentNumber",
			"AccountNumber",
			"PostalLocationTown",
			"PostalLocationCountry",
			"LocationTown",
			"LocationCountry",
		)
	})
	View("default", func() {
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

	View("mandatory", func() {
		Attribute("SubmissionDate")
		Attribute("ClientType")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("AccountNumber")
		Attribute("PostalLocationTown")
		Attribute("PostalLocationCountry")
		Attribute("LocationTown")
		Attribute("LocationCountry")
	})
})

var ContactUploadCTA = Type("ContactUploadCTA", func() {
	Description("A – Individual Credit Consumer")

	Attribute("Surname", String, func() {
		Description("Surname") // Allow Hyphen, Apostrophe, Space.
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

	Required("Surname", "Forename1", "Forename2", "Forename3")

})

var ContactUploadCTB = Type("ContactUploadCTB", func() {
	Description("B – NonIndividual Credit Consumer")

	Attribute("RegisteredName", String, func() {
		Description("The Name as Registered by an Approved Registrar") // Company Name
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "7")
	})

	Attribute("TradingName", String, "Trading Name", func() {
		Description("The Business or Trading Name")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "8")
	})

	Required("RegisteredName")

})
