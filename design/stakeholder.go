package design

import (
	. "goa.design/goa/v3/dsl"
)

// F. Stakeholder Information
//
// Where the Stakeholder is an institution or organisation, the names should be provided
// should be the company registered name only. The Other Names (e.g. surname, first Names)
// should either not be provided or the Company name can be given in all the name fields.
var Stakeholder = Type("stakeholder", func() {
	Description("Stakeholder API")
	TypeName("Stakeholder")
	ContentType("application/json")

	// The Stakeholder’s Family Name.
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 4.3.7
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "1")
	})

	// The Stakeholder’s First Name
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 4.3.7
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe
	Attribute("Forename1", String, "Stakeholder’s First Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Stakeholder’s Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename2", String, "Stakeholder’s Given Name", func() {
		Meta("rpc:tag", "3")
	})

	// Stakeholder’s Other Name or Initials.
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe
	Attribute("Forename3", String, "Stakeholder’s Other Name or Initials", func() {
		Meta("rpc:tag", "4")
	})

	// Name of the Company/Corporate entity that is a stakeholder to the Subject Account.
	//
	// Alphanumeric
	// Not more than 70 characters
	// Conditional to 4.3.7
	Attribute("CompanyOrCorporateName", String, "Company/Corporate Name", func() {
		Meta("rpc:tag", "5")
	})

	// Stakeholder’s Date of Birth (individual) Or Stakeholder’s Date of Registration (non -individual).
	//
	// Mandatory Field
	// Date Field
	// Not more than 8 Characters
	// Cannot be in future
	Attribute("DoBOrDoR", String, "Date Of Birth Or Date of Registration", func() {
		Meta("rpc:tag", "6")
	})

	// Stakeholder’s Gender, Male/Female/Institution
	// Options :
	// M – Male
	// F – Female
	// I – Institution
	//
	// Mandatory
	// Lookup
	// Not more than 1 character
	Attribute("Gender", String, "Gender", func() {
		Meta("rpc:tag", "7")
	})

	// ISO Country Code for Stakeholder’s Nationality.
	//
	// Mandatory Field – Issue Warning
	// Alphanumeric
	// Not more than 2 characters
	// Matches ISO CODE
	Attribute("Nationality", String, "Nationality", func() {
		Meta("rpc:tag", "8")
	})

	// Client reference Number in Core Lenders system for the Account in
	// which this Stakeholder has linkage . For Client-Centric Systems.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "9")
	})

	// Account Number in Core Lenders system for the Account in which this Stakeholder has linkage.
	// Same as Client Number for Account Centric Systems
	//
	// Mandatory field
	// Alphanumeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "10")
	})

	// The Stakeholder’s Primary Identification document
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 3 characters
	// If Gender = I, then stakeholder identification = 005
	Attribute("PrimaryIDocument", String, "Primary Identification Document", func() {
		Meta("rpc:tag", "11")
	})

	// The Number of the Stakeholder’s Primary Identification Document.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "12")
	})

	// If Stakeholder provided a second Identification Document, then the Document
	// Type for this secondary identification.
	// Options are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	//
	// Alphanumeric
	// Not more than 3 characters
	Attribute("SecondaryIDocumentType", String, "Secondary Identification Document Type", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Stakeholder’s Secondary Identification Document Mandatory if
	// secondary identification Document type is provided
	//
	// Conditional to 4.3.13
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("SecondaryIDocumentNumber", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "14")
	})

	// If Stakeholder provided a third Identification Document, then the Document
	// Type for this other identification.
	// Options are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	//
	// Alphanumeric
	// Not more than 3 characters
	Attribute("OtherIDocumentType", String, "Other Identification Document Type", func() {
		Meta("rpc:tag", "15")
	})

	// The Number of the Stakeholder’s Other Identification Document.
	// Mandatory if Other Identification
	// Document Type is provided.
	//
	// Conditional Field to 4.3.15
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("OtherIDocNumber", String, "Other Identification Doc Number", func() {
		Meta("rpc:tag", "16")
	})

	// Stakeholder’s Email Address, if Available.
	//
	// Validate Email Syntax
	// Not more than 50 characters
	// Alphanumeric
	Attribute("Email", String, "E-Mail", func() {
		Meta("rpc:tag", "17")
	})

	// The company Registration Certificate Number.
	//
	// Conditional to 4.3.11
	// Not more than 20 Characters
	// Alphanumeric
	// Should be the same with field 4.3.12 if option 005 is provided in 4.3.11
	Attribute("CompanyRegistrationNumber", String, "Company Registration Number", func() {
		Meta("rpc:tag", "18")
	})

	// The Registration Certificate Number - format to conform to Registrar of Companies in Kenya.
	// Required, only if applicable because the company type changes, the history
	// details of the previous type can be matched to the new company type.
	Attribute("PreviousRegistrationNumber", String, "Previous Registration Number", func() {
		Meta("rpc:tag", "19")
	})

	// The Company Income Tax PIN Number
	//
	// Mandatory if Gender is “I” in field 4.3.7
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("CompanyPINumber", String, "Company PIN Number", func() {
		Meta("rpc:tag", "20")
	})

	// Income Tax VAT Number
	//
	// Alphanumeric Field
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("CompanyVATNumber", String, "Company VAT Number", func() {
		Meta("rpc:tag", "21")
	})

	// Type of stake held
	// Options are :
	// A – Director
	// B - Shareholder
	// C - Partner
	// D - Proprietor
	// E - Trust Beneficiary
	// F - Power of Attorney
	// G - Guarantor
	//
	// Alphanumeric
	// Lookup
	// Mandatory with warning
	Attribute("StakeholderType", String, "Stakeholder Type", func() {
		Meta("rpc:tag", "22")
	})

	// If stakeholder is a shareholder, the Shareholding in the company.
	//
	// Number Format
	// Not more than 6 characters
	Attribute("PercentageShares", String, "Percentage of Shares in company", func() {
		Meta("rpc:tag", "23")
	})

	// The Primary Stakeholder’s Telephone contact Number in the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("MobileTelephoneNumber", String, "Mobile Telephone Number", func() {
		Meta("rpc:tag", "24")
	})

	// The Secondary Telephone contact Number in the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("HomeTelephoneNumber", String, "Home Telephone Number", func() {
		Meta("rpc:tag", "25")
	})

	// The Any other Telephone contact Number in the
	// Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("WorkTelephoneNumber", String, "Work Telephone Number", func() {
		Meta("rpc:tag", "26")
	})

	// Stakeholder’s Postal Address Line 1
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("PostalAddress1", String, "Postal Address 1", func() {
		Meta("rpc:tag", "27")
	})

	// Stakeholder’s Postal Address Line 2
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("PostalAddress2", String, "Postal Address 2", func() {
		Meta("rpc:tag", "28")
	})

	// Town of Stakeholder’s Postal Address
	//
	// Alphanumeric
	// Not more than 30 characters
	Attribute("Town", String, "Town", func() {
		Meta("rpc:tag", "29")
	})

	// Country of Stakeholder’s Postal Address.
	//
	// Alphanumeric
	// Not more than 2 characters
	Attribute("Country", String, "Country", func() {
		Meta("rpc:tag", "30")
	})

	// Post Code of Stakeholder’s Postal Address.
	//
	// Alphanumeric
	// Not more than 10 characters
	Attribute("PostCode", String, "Post Code", func() {
		Meta("rpc:tag", "31")
	})

	// Stakeholder’s Physical Address Line 1
	// If stakeholder is an individual; this would be the Residential Address.
	// If the stakeholder is an institution, this would be the registered Office Address.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress1", String, "Physical Address 1", func() {
		Meta("rpc:tag", "32")
	})

	// Stakeholder’s Physical Address Line 2
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress2", String, "Physical Address 2", func() {
		Meta("rpc:tag", "33")
	})

	// Stakeholder’s Physical Address Line 3
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress3", String, "Physical Address 3", func() {
		Meta("rpc:tag", "34")
	})

	// Plot Land Ref (LR) No of Stakeholder’s Physical Address.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("PlotNumber", String, "Plot Number", func() {
		Meta("rpc:tag", "35")
	})

	// Town of Stakeholder’s Physical Address.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 30 characters
	Attribute("PhysicalLocationTown", String, "Physical Location Town", func() {
		Meta("rpc:tag", "36")
	})

	// ISO Country Code of the Stakeholder’s Physical Address.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 2 characters
	// Lookup
	Attribute("PhysicalLocationCountry", String, "Physical Location Country", func() {
		Meta("rpc:tag", "37")
	})
})

var StoredStakeholder = ResultType("application/vnd.goa.stakeholder", func() {
	TypeName("StoredStakeholder")
	Attributes(func() {
		Extend(Stakeholder)
		Required()
	})
	View("", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyOrCorporateName")
		Attribute("DoBOrDoR")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocumentType")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocumentType")
		Attribute("OtherIDocNumber")
		Attribute("Email")
		Attribute("CompanyRegistrationNumber")
		Attribute("PreviousRegistrationNumber")
		Attribute("CompanyPINumber")
		Attribute("CompanyVATNumber")
		Attribute("StakeholderType")
		Attribute("PercentageShares")
		Attribute("MobileTelephoneNumber")
		Attribute("HomeTelephoneNumber")
		Attribute("WorkTelephoneNumber")
		Attribute("PostalAddress1")
		Attribute("PostalAddress2")
		Attribute("Town")
		Attribute("Country")
		Attribute("PostCode")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalAddress2")
		Attribute("PhysicalAddress3")
		Attribute("PlotNumber")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
	})
})
