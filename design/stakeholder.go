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
	Description("")
	TypeName("Stakeholder")
	ContentType("application/json")

	// The Family Name of stakeholder, or the registered company name
	// if the stakeholder is a company or institution.
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The First Name of stakeholder.
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name of stakeholder.
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	Attribute("Forename3", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Options – Mr., Mrs., Miss, Ms, Dr. , Prof., Hon.
	Attribute("Salutation", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Date of Birth or date of registration for non-individual stakeholders.
	// Cannot be in the Future
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Client Reference Number linking stakeholder to Core Banking system.
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number linking stakeholder to Banking system.
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Gender of the stakeholder.
	// Options Available :
	// M – Male
	// F - Female
	// I - Institution or Organisation
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// The Kenya Revenue Authority Income Tax PIN Number.
	// May be required at a later Date
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Nationality of the Individual stakeholder or the country
	// of registration for non-individual stakeholders.
	// Default - Kenyan
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Applicable to non-individual stakeholders only.
	// Options to Select From :
	// M - Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Primary Identification document Provided by stake holder.
	// Options Are:
	// National ID
	// Passport
	// lien Registration
	// Service ID
	// Company Registration Certificate
	//
	// National Identification is the preferred document but the other are acceptable.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	//The company registration Number is the Registration Number of the Institutional stake holder.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Id Provided
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Any Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	//
	// For Individual Stakeholders
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification Document Type is Provided.
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

	// The Number of Other Identification Document, where provided .
	// Mandatory if Other Identification Document Type is Provided.
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
	// Mandatory if Additional Identification Document Type is Provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Company Registration Certificate Number
	Attribute("CompanyRegistrationNo", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// The Company Income Tax PIN Number
	Attribute("CompanyPINNo", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The Income Tax VAT registration Number.
	Attribute("CompanyVATNo", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Type of stake held in the Company.
	// Options are :
	// A - Managing Director
	// B - Director
	// C - Share Holder
	// D - Company Secretary
	// E - Senior Management
	// For limited liability borrowers, A Profile Record for all directors
	// must be provided. Shareholders’ Profiles are not required.
	Attribute("StakeholderType", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Number of Shares held in the Company.
	// N/A for limited liability companies
	Attribute("NofShares", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// The Percentage of shares Held in the Company.
	// Percentage of shareholding in the Company for limited liability companies.
	// Not required if no of shares is provided.
	Attribute("PercentageShares", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// The Stake holder’s Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// The Stake holder’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Stake holder’s other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Stake holder’s Postal Address Line1.
	// First Line of Full stakeholder Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Stake holder’s Postal Address Line 2.
	// Second Line of Full stakeholder Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Town of Stake holder’s Postal Address
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Country of Stake holder’s Postal Address
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Post Code of Stake holder’s Address
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "35")
	})

	// Stake holder’s Residence or registered Office if a non-individual.
	// Location of stakeholder ( e.g. Office Address or Building)
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "36")
	})

	// Street of Stake holder’s residence or registered office.
	// Location street of Stakeholder’s location.
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "37")
	})

	// Location street of Stakeholder’s location.
	Attribute("PhysicalAddress3", String, "", func() {
		Meta("rpc:tag", "38")
	})

	// Plot Land Ref (LR) No of Stake holder’s office/residence
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "39")
	})

	// Stake holder’s Town of location.
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "40")
	})

	// Stake holder’s Country of location.
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "41")
	})
})
