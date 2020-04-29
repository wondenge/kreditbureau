package design

import (
	. "goa.design/goa/v3/dsl"
)

// E. Non Individual Consumer Information
var Company = Type("company", func() {
	Description("")
	TypeName("Company")
	ContentType("application/json")

	// The Name as Registered with the Registrar of Companies
	Attribute("RegisteredName", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Business or Trading Name
	Attribute("TradingName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Date Registered with the Registrar of Companies
	Attribute("RegistrationDate", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Registration Certificate Number
	Attribute("RegistrationNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Country of Registration of the Company, Defaults to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Client Reference Number linking Company to Core Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// The Account Number Linking Client to Bank’s Accounting System
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Type of Organisation.
	// The type of Organisation.
	// Options Available :
	// A - Limited Company
	// B - Sole Proprietor
	// C - NGO
	// D - Other
	Attribute("CompanyType", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Industry Code for the Line of business as per the Central Bank Returns – Options are :
	// Agriculture
	// Manufacturing
	// Building/Construction
	// Mining/Quarrying
	// Energy /water
	// Trade
	// Tourism/ Restaurant/Hotels
	// Transport/ Communications
	// Real Estate
	Attribute("IndustryCode", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// The Income Tax PIN Number.
	// May be Mandatory Later
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Value Added Tax Registration Number
	Attribute("VATNumber", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// 1 is a sole proprietor
	Attribute("NofDirectors", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// 1 if a sole proprietor
	Attribute("NofShareholders", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Business Email Address
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Company Website url
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Status of Trading of the company –
	// Options Available :
	// Dormant
	// Actively Trading
	// Under Management
	// Dissolved
	Attribute("Status", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Options :
	// Actively Trading
	// Under management
	// Dissolved
	Attribute("TradingStatus", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// Date of the Current Status.
	// Provide the Date of the Last status.
	Attribute("StatusDate", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// The consumer’s Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// The consumer’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// The consumer’s other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// Consumer Postal Address Line1.
	// First Line of Company’s Full Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Consumer Postal Address Line 2.
	// Second Line of Company’s Full Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// Town of consumer Postal Address.
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Country of consumer Postal Address.
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Post Code of consumer Address.
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Consumer Registered Office location building.
	// Location ( e.g. Building or Apartment) housing the company’s Registered Office.
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// Consumer Office location Street.
	// Street Name of Company’s Registered office.
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Office location Plot Land Ref (LR) No.
	// Land Reference Number where Office is located
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Office Location Town
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Office Location Country
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "31")
	})
})
