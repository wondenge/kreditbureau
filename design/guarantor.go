package design

import (
	. "goa.design/goa/v3/dsl"
)

// H. Guarantor Information
//
// Where the Guarantor is an institution or organisation, the names should be provided
// should be the company registered name only. The Other Names (e.g. surname, first Names)
// should either not be provided or the Company name can be given in all the name fields.

var Guarantor = Type("guarantor", func() {
	Description("Guarantor API")
	TypeName("Guarantor")
	ContentType("application/json")

	// The Name of Lender as registered with the Registrar of companies.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersRegisteredName", String, "Lenders Registered Name", func() {
		Meta("rpc:tag", "1")
	})

	// The Lenders Trading Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersTradingName", String, "Lenders Trading Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Lenders Branch Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersBranchName", String, "Lenders Branch Name", func() {
		Meta("rpc:tag", "3")
	})

	// The format of the branch code is IXXXYYY
	// Where
	// I – the Institution type code
	// B for Banks
	// D For MFBs
	// S For Saccos
	// M for MFIs
	// L for Leasing Companies
	// XXX is the Lenders Institution Code left padded with Zeros
	// e.g. 098 for a bank whose code is 98.
	// YYY is the Lenders Branch Code left padded with Zeros
	// e.g. 009 for a branch whose code is 9.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 7 characters
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "4")
	})

	// Client Reference Number Linking guarantor to Account being guaranteed in the Lenders systems.
	// This is for Client-Centric systems.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "5")
	})

	// Account Number Linking Guarantor to Account Being Guaranteed to Core Lenders system.
	//The Account Number Is the Same as Client Number for Account Centric Systems.
	//
	// Mandatory field
	// Alphanumeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "6")
	})

	// Guarantor’s Surname.
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	// Mandatory if option M/F is selected in field 4.4.13
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "7")
	})

	// First Name of Guarantor
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	// Mandatory if option M/F is selected in field 4.4.13
	Attribute("Forename 1", String, "First Name of Guarantor", func() {
		Meta("rpc:tag", "8")
	})

	// Guarantor’s Second Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space,
	// Conditional to field 4.4.13
	Attribute("Forename 2", String, "Guarantor’s Second Name", func() {
		Meta("rpc:tag", "9")
	})

	// Guarantor’s Other Names
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space,
	// Conditional to field 4.4.13
	Attribute("Forename 3", String, "Guarantor’s Other Names", func() {
		Meta("rpc:tag", "10")
	})

	// Name of the Company/Corporate entity that is a Guarantor to the Subject Account.
	//
	// Alphanumeric
	// Not more than 70 Characters
	Attribute("CompanyOrCorporateName", String, "Company/Corporate Name", func() {
		Meta("rpc:tag", "11")
	})

	// Guarantor’s Date of Birth if it is an individual Guarantor’s date
	// of Registration if it is an Institution.
	//
	// Mandatory – Issue Warning
	// Not a future date
	// If an individual, Client Type A,
	// - Not under 18
	// - And Not over 100 years
	// Date Format YYYYMMDD – Issue Warning
	// Registration Date Not in the Future
	// Not more than 8 Numeric characters
	Attribute("DOBorDOR", String, "Date Of Birth / Date of Registration", func() {
		Meta("rpc:tag", "12")
	})

	// Guarantor’s Gender.
	// Options Available :
	// M - Male
	// F - Female
	// I - Institutional
	//
	// Mandatory
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	Attribute("Gender", String, "Gender", func() {
		Meta("rpc:tag", "13")
	})

	// ISO Code for the Country of Guarantor’s Nationality.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 2 characters
	Attribute("Nationality", String, "Nationality", func() {
		Meta("rpc:tag", "14")
	})

	// The Guarantor Marital status.
	// Options Available are :
	// M - Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	//
	// Lookup
	// Alphanumeric
	// Not More than 1 character
	Attribute("MaritalStatus", String, "Marital Status", func() {
		Meta("rpc:tag", "15")
	})

	// The Type of Guarantee
	// Options are :
	// A - Directors guarantee
	// B - Personal guarantee
	// C - Corporate guarantee
	// D - Bank Guarantee
	// F – Personal Guarantee with Collateral
	//
	// Mandatory
	//Alphanumeric
	//Lookup
	Attribute("GuaranteeType", String, "Guarantee Type", func() {
		Meta("rpc:tag", "16")
	})

	// Clients relations with the Guarantor
	// Conditional with Guarantee Type
	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// - if Guarantee Type = B
	// B = Shareholder, Partner, etc.
	// - if Guarantee Type = A or C
	// C = Other
	//
	// Alphanumeric
	// Lookup
	Attribute("GuarantorRelationship", String, "Guarantor Relationship", func() {
		Meta("rpc:tag", "17")
	})

	// The Amount of Limit of the Guarantee if It has a limit
	//
	// Amount field
	// Mandatory
	// Not more than 16 Characters
	Attribute("GuaranteeLimit", String, "Guarantee Limit", func() {
		Meta("rpc:tag", "18")
	})

	// The Guarantor’s Primary Identification document Type Provided to the Lender.
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
	Attribute("PrimaryIDocument", String, "Primary Identification Document", func() {
		Meta("rpc:tag", "19")
	})

	// The Identification document Number.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "20")
	})

	// The Guarantor’s Secondary Identification document Type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Alphanumeric
	// Not more than 3 characters
	Attribute("SecondaryIDocument", String, "Secondary Identification Document", func() {
		Meta("rpc:tag", "21")
	})

	// Mandatory if Secondary Identification Document type is provided.
	//
	// Conditional Field to 4.4.21
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("SecondaryIDocumentNumber", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "22")
	})

	// The Guarantor’s Other Identification document type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Alphanumeric
	// Not more than 3 characters
	Attribute("OtherIDocument", String, "Other Identification Document", func() {
		Meta("rpc:tag", "23")
	})

	// Mandatory if Other Identification document type is provided.
	//
	// Conditional Field to 4.4.23
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("OtherIDocumentNumber", String, "Other Identification Document Number", func() {
		Meta("rpc:tag", "24")
	})

	// Guarantor’s Employer, if Guarantor is employed.
	//
	// Alphanumeric
	// Max Characters 50
	Attribute("EmployerName", String, "Employer Name", func() {
		Meta("rpc:tag", "25")
	})

	// If guarantor is employed, then Type Of employment of the guarantor.
	// Options Available :
	// A - Casual
	// B - Contract
	// C - Permanent
	// D - Pensionable
	// E - Self-Employed
	// Mandatory if guarantor is employed.
	//
	// Alphanumeric
	// Lookup
	// Not more than 1 character
	Attribute("EmploymentType", String, "Employment type", func() {
		Meta("rpc:tag", "26")
	})

	// The band within which the Customer’s Gross Monthly Salary falls.
	// Options given are:
	// A – KES 0/- To 50,000/-
	// B – KES 50,001/- To 100,000/-
	// C - KES 100,001/- To 200,000/-
	// D - KES 200,001/- To 500,000/-
	// E - Over KES 500,000/-
	//
	// Alphanumeric
	// Lookup
	// Not more than 1 character
	Attribute("IncomeBand", String, "Income Band", func() {
		Meta("rpc:tag", "27")
	})

	// The Guarantor’s Primary Telephone contact Number in the Form of:
	// CCCAAANNNNNNN
	// Where:
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("MobileTelephoneNumber", String, "Mobile Telephone Number", func() {
		Meta("rpc:tag", "28")
	})

	// The Guarantor’s Secondary Telephone contact Number, if provided in the Form of :
	// CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("HomeTelephoneNumber", String, "Home Telephone Number", func() {
		Meta("rpc:tag", "29")
	})

	// IF Guarantor is employed then the employer’s Telephone contact Number in the Form of :
	// CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("WorkTelephoneNumber", String, "Work Telephone Number", func() {
		Meta("rpc:tag", "30")
	})

	// Guarantor’s Postal Address Line 1
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress1", String, "Postal Address 1", func() {
		Meta("rpc:tag", "31")
	})

	// Postal Address Extra details
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress2", String, "Postal Address 2", func() {
		Meta("rpc:tag", "32")
	})

	// Town of Postal Address
	//
	// Mandatory
	// Alphanumeric
	// Not more than 30 characters
	Attribute("Town", String, "Town", func() {
		Meta("rpc:tag", "33")
	})

	// Country of Guarantor’s Postal Address.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 30 characters
	Attribute("Country", String, "Country", func() {
		Meta("rpc:tag", "34")
	})

	// Post Code of Postal Address
	//
	// Alphanumeric
	// Not more than 10 characters
	Attribute("PostCode", String, "Post code", func() {
		Meta("rpc:tag", "35")
	})

	// Guarantor’s Physical Address Line 1
	// If the Guarantor is an Individual then this would be the Residential
	// Address of the Guarantor. If the Guarantor is an Institution, then
	// this would be the Guarantor’s registered Office Physical Address.
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress1", String, "Physical Address 1", func() {
		Meta("rpc:tag", "36")
	})

	// Guarantor’s Physical Address Line 2
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress2", String, "Physical Address 2", func() {
		Meta("rpc:tag", "37")
	})

	// Plot Land Ref (LR) No of Guarantor’s Physical Address
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PlotNumber", String, "Plot Number", func() {
		Meta("rpc:tag", "38")
	})

	// Town of Guarantor’s Physical Address.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 30 characters
	Attribute("PhysicalLocationTown", String, "Physical Location Town", func() {
		Meta("rpc:tag", "39")
	})

	// ISO Country Code of the Country of Residence or Registered Office of the Guarantor.
	//
	// Mandatory
	// Alphanumeric
	// Not more 2 characters
	Attribute("PhysicalLocationCountry", String, "Physical Location Country", func() {
		Meta("rpc:tag", "40")
	})

	// Email address of the Guarantor.
	//
	// Alphanumeric
	// Not more than 50 characters
	// Email syntax
	Attribute("Email", String, "Email", func() {
		Meta("rpc:tag", "41")
	})

})
