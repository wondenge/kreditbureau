package design

import (
	. "goa.design/goa/v3/dsl"
)

// D. Individual Consumer Employment Information
//
// The Institution’s customers could be employed or self-employed.
// Details of the customer’s employment are captured in the employment information record.
// Where a customer has provided employment record details, the institution is required to
// provide one record for the employment details as laid out below
var Employment = Type("employment", func() {
	Description("Individual Consumer, Employer and Account API")
	TypeName("Employment")
	ContentType("application/json")

	// The Family Name or Surname
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "1")
	})

	//The First Name of Customer.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50
	// Allow Hyphen, Apostrophe
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50
	// Allow Hyphen, Apostrophe,
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "4")
	})

	// Options are :
	// Mr., Mrs., Miss, Ms, Dr. , Prof., Hon.
	Attribute("Salutation", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Date of Birth of the Customer.
	// Cannot be in the Future.
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Client Reference Number linking client to Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number linking client to Banking system
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Male/Female ( M/F)
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Kenya Revenue Authority Income Tax PIN Number.
	// May be required at a later Date.
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// The Nationality of the Customer- Defaults to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Options to Select From will be provided
	// Options Available :
	// M – Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Primary Identification document Provided on Opening the Account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	//
	// For Kenya nationals, default should be the National Identification.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Identification document specified above.
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification document is Provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Any Other Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of Other Identification Document, where provided.
	// Mandatory if Other Identification Document is Provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Any Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification document is Provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// Client’s Email Address if provided
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Name of the Employer if not self-employed.
	Attribute("EmployerName", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Type Of employment.
	// Options Available :
	// Casual Contract
	// Permanent,
	// Pensionable
	//Self-Employed
	Attribute("EmploymentType", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The Position in Organisation
	Attribute("EmployeePosition", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// The Department Stationed
	Attribute("EmployeeDepartment", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// The band within which the Customer’s Gross Monthly Salary falls.
	// Options given are :
	// A - 0 To 50,000 KES
	// B - 50,000 To 100,000 KES
	// C – 100,000 To 200,000 KES
	// D – 200,000 To 250,000 KES
	// E - Over 250,000 KES
	Attribute("SalaryBand", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Employer’s Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// The Employer’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Any other Employer Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Employer’s Postal Address Line1.
	// First Line of the full employer’s Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Employer’s Postal Address Line 2.
	// Second Line of the full employer’s Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Employer’s Town of Postal Address.
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Employer’s Country of Postal Address.
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Post Code of Employer’s Address.
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Employer’s registered Office Street Address.
	// The Location ( Street Name) of the Employer’s Offices.
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "35")
	})

	// Employer’s Office building and office number.
	// The Location ( Building or Apartment No) of the address Employer’s Office
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "36")
	})

	// Office Plot Land Ref (LR) Number.
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "37")
	})

	// Employers Address Town.
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "38")
	})

	// Employer’s Location Country
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "39")
	})

})
