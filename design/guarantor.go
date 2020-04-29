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
	Description("")
	TypeName("Guarantor")
	ContentType("application/json")

	// Bank and Branch Code where facility being guaranteed is domiciled or was granted.
	// Branch’s sort code as provided by the KBA.
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// Client Reference Number linking Guarantor to Banking system.
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Account Number linking guarantor to Bank’s Accounting System for the Account being Guaranteed.
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Surname or Family Name of Guarantor.
	// Where the guarantor is an institution, then the company name shall be provided as the Surname.
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// First Name of Guarantor
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Other names of the guarantor
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Initials or other names
	Attribute("Forename3", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Date of Birth of guarantor.
	// Where the Guarantor is an institution, then the Date of Registration shall be
	// provided in the Date of Birth field.
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Gender of the Guarantor.
	// Options are :
	// M – Male
	// F - Female
	// I - Institutional
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Kenya Revenue authority Income Tax PIN Number
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Nationality of Guarantor.
	// Default to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Marital status of the guarantor, if an individual.
	// Options are:
	// M - Married
	// S - single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Type of guarantee being offered.
	// Options Are:
	// A - Directors guarantee
	// B - Personal guarantee
	// C - Corporate guarantee
	// D - Bank Guarantee
	Attribute("GuaranteeType", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The amount of limit on the guarantee.
	Attribute("GuaranteeLimit", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// The Primary Identification document given by the guarantor.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	// Company Registration Certificate
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Number of the Primary Id Provided.
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification Document Type is provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Any Other Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// The Number of Other Identification Document, where provided .
	// Mandatory if Other Identification Document Type is provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification Document Type is provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// The Email Address
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The guarantor’s employer, if employed.
	// Provide Non-Individual consumer Information Profile for the Employer
	Attribute("EmployerName", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Type Of employment.
	// Options Available :
	// Casual
	// Contract
	// Permanent
	// Pensionable
	// Self-Employed
	Attribute("EmploymentType", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// The Position in Organisation.
	// Mandatory if Employed.
	Attribute("EmployeePosition", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// The Department Stationed.
	// Mandatory if Employed.
	Attribute("EmployeeDepartment", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// The Range in which the Gross Monthly Salary in KES Falls.
	// Options are :
	// A – 0 To 50,000.00 KES
	// B - 50,000 To 100,000 KES
	// C - 100,000 To 200,000 KES
	// D - 200,000 To 250,000 KES
	// E - Over 250,000 KES
	//
	// Mandatory if Employed
	Attribute("SalaryBand", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// The Guarantor’s Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// The Guarantor’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Any other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Guarantor’s Postal Address Line1.
	// First Line of guarantor’s Full Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Guarantor’s Postal Address Line 2.
	// Second Line of guarantor’s Full Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Guarantor’s Town of Postal Address
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Guarantor’s Country of Postal Address
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "35")
	})

	// Guarantor’s Post Code of Address
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "36")
	})

	// Guarantor’s Residential Address e.g. street, estate
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "37")
	})

	// Guarantor’s Residential location e.g. House number or apartment number
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "38")
	})

	// Plot Land Ref (LR) No
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "39")
	})

	// Guarantor’s Town of residence
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "40")
	})

	// Guarantor’s Country of residence
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "41")
	})

})
