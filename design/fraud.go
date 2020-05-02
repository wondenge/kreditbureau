package design

import (
	. "goa.design/goa/v3/dsl"
)

// L. Fraudulent Activities API
var Fraud = Type("fraud", func() {
	Description("Fraudulent Activities API")
	TypeName("Fraud")
	ContentType("application/json")

	// The Name of Lender Reporting the Fraud, as registered with the Registrar of companies.
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

	// The Lenders Branch Name, where the fraud is reported to have taken place.
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

	// Client Reference Number Linking to
	// Account involved in the Fraud in the Lender’s system.
	// This is for Client-Centric systems.
	//
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("Client Number", String, "Client Number", func() {
		Meta("rpc:tag", "5")
	})

	// Account Number Linking to Account involved in the Fraud, in the Lender’s System.
	//The Account Number Is the Same as Client Number for Account Centric Systems.
	//
	// Mandatory field
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("Account Number", String, "Account Number", func() {
		Meta("rpc:tag", "6")
	})

	// Brief Description of fraud type.
	//
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("Fraud Type", String, "Fraud Type", func() {
		Meta("rpc:tag", "7")
	})

	// Date on which Fraud took place.
	//
	// Date Field
	// Not more than 8 characters
	// Can’t be in future
	Attribute("FraudIncidentDate", String, "Fraud Incident Date", func() {
		Meta("rpc:tag", "8")
	})

	// Date on which Fraud was reported.
	//
	// Date Field
	// Not more than 8 Characters
	// Can’t be in future
	Attribute("FraudReportDate", String, "Fraud Report Date", func() {
		Meta("rpc:tag", "9")
	})

	// Amount involved in Fraud
	//
	// Amount Field
	// Not more than 16 characters
	Attribute("Amount", String, "Amount", func() {
		Meta("rpc:tag", "10")
	})

	// Actual loss incurred in the fraud as at reporting date.
	//
	// Currency field
	// Not more than 16 Characters
	Attribute("Loss Amount", String, "Loss Amount", func() {
		Meta("rpc:tag", "11")
	})

	// Currency code of the Amount involved in the Fraud.
	// Default is Kenya Shillings
	//
	// Alphanumeric
	// Not more than 3 Characters
	// ISO Currency Codes
	Attribute("Currency Code", String, "Currency Code", func() {
		Meta("rpc:tag", "12")
	})

	// Incident Details
	// Alphanumeric
	// Not more than 200 characters
	Attribute("Incident Details", String, "Incident Details", func() {
		Meta("rpc:tag", "13")
	})

	// A brief on the Forensic evidence
	//
	// Alphanumeric
	// Not more than 200 Characters
	Attribute("ForensicInformation", String, "Forensic Information", func() {
		Meta("rpc:tag", "14")
	})

	// The Family Name or Surname of the Person Involved in the Fraud.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space,
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "15")
	})

	// The First Name of the Person Involved in the Fraud.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "16")
	})

	// The Given Name of the Person Involved in the Fraud.
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "17")
	})

	// Other Name or Initials of the Person Involved in the Fraud.
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "18")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Mandatory Field
	// Data is Alpha Numeric
	// Not more than 3 Characters
	Attribute("PrimaryIDocumentType", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "19")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocumentNumber", String, "Primary Identification Document Number", func() {
		Meta("rpc:tag", "20")
	})

	// The Equivalent Amount involved in the fraud in KES.
	//
	// Mandatory Field
	// Currency Field
	// No more than 16 Characters
	Attribute("AmountInKES", String, "Amount in Kenya Shillings", func() {
		Meta("rpc:tag", "21")
	})

})
