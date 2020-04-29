package design

import (
	. "goa.design/goa/v3/dsl"
)

// C. Individual Consumer Information
//
// The Individual consumer record refers to the individual customer of the
// institutions and contains a profile of the customer from the account opening
// information and any other information the customer provides to the institution.
var Individual = Type("individual", func() {
	Description("")
	TypeName("Individual")
	ContentType("application/json")

	// The Family Name or Surname
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "1")
	})

	//The First Name
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	Attribute("Forename3", String, "", func() {
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

	// Any Additional Identification Document Provided.
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

	// The Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// The Any other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// Client’s Postal Address Line1
	// This is the first line of the Full Client’s Postal address
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Client’s Postal Address Line 2
	// This is the second line of the Full Client’s Postal address
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Town of Client’s Postal Address
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Country of Client’s Postal Address
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// Post Code of Client’s Address
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Client’s residential Address e.g. Street Address, Estate or village
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// House or Apartment number of Client’s Residence
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Plot Land Ref (LR) No of Client’s residence
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Town of Client’s residence
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Country of Client’s residence
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Client’s Email Address if provided
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "34")
	})

})