package design

import (
	. "goa.design/goa/v3/dsl"
)

var FraudActivity = Type("fraudactivity", func() {
	Description("Fraudulent Activities API")
	TypeName("FraudActivity")
	ContentType("application/json")

	Attribute("LendersRegisteredName", String, func() {
		Description("Name of Lender Reporting the Fraud") // registered with the Registrar of companies
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "1")
	})

	Attribute("LendersTradingName", String, func() {
		Description("Lenders Trading Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("LendersBranchName", String, func() {
		Description("Lenders Branch Name, where the fraud is reported ")
		MinLength(1)
		MaxLength(50)
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
	Attribute("LendersBranchCode", String, func() {
		Description("Lenders Branch Code")
		MinLength(1)
		MaxLength(7)
		Meta("rpc:tag", "4")
	})

	// Client Reference Number Linking to Account
	// involved in the Fraud in the Lender’s system.
	// This is for Client-Centric systems.
	Attribute("ClientNumber", String, func() {
		MinLength(1)
		MaxLength(20)
		Description("client reference number linking to account involved in fraud")
		Meta("rpc:tag", "5")
	})

	// Account Number Linking to Account involved
	// in the Fraud, in the Lender’s System.
	// The Account Number Is the Same as Client Number
	// for Account Centric Systems.
	Attribute("AccountNumber", String, func() {
		Description("account number linking to account involved in fraud")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "6")
	})

	Attribute("FraudType", String, "Fraud Type", func() {
		Description("brief description of fraud type")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "7")
	})

	Attribute("FraudIncidentDate", String, func() {
		Description("Date on which Fraud took place.")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "8")
	})

	Attribute("FraudReportDate", String, func() {
		Description("Date on which Fraud was reported")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "9")
	})

	Attribute("Amount", Int64, func() {
		Description("Amount involved in Fraud")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "10")
	})

	Attribute("LossAmount", Int64, func() {
		Description("Actual loss incurred in the Fraud as at reporting date")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "11")
	})

	Attribute("CurrencyCode", String, func() {
		Description("Currency code of the Amount involved in the Fraud") // ISO Currency Codes
		MinLength(3)
		MaxLength(3)
		Default("KES") // Default is Kenya Shillings
		Meta("rpc:tag", "12")
	})

	Attribute("IncidentDetails", String, func() {
		Description("Incident Details")
		MinLength(1)
		MaxLength(200)
		Meta("rpc:tag", "13")
	})

	Attribute("ForensicInformation", String, func() {
		Description("A brief on the Forensic evidence")
		MinLength(1)
		MaxLength(200)
		Meta("rpc:tag", "14")
	})

	Attribute("Surname", String, "Surname", func() {
		Description("Family Name or Surname of the Person Involved in the Fraud") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "15")
	})

	Attribute("Forename1", String, "The First Name", func() {
		Description("The First Name of the Person Involved in the Fraud") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "16")
	})

	Attribute("Forename2", String, "The Given Name", func() {
		Description("Given Name of the Person Involved in the Fraud.") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "17")
	})

	Attribute("Forename3", String, "Other Name or Initials", func() {
		Description("Other Name or Initials of the Person Involved in the Fraud") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "18")
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
		Meta("rpc:tag", "19")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumber", String, func() {
		Description("Primary Identification Document Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "20")
	})

	Attribute("AmountInKES", Int64, func() {
		Description("Equivalent Amount involved in the fraud in KES")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "21")
	})

})

var StoredFraudActivity = ResultType("application/vnd.goa.fraudactivity", func() {
	TypeName("StoredFraudActivity")
	Attributes(func() {
		Extend(FraudActivity)
		Required(
			"LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"AccountNumber",
			"Surname",
			"Forename1",
			"PrimaryIDocumentType",
			"PrimaryIDocumentNumber",
			"AmountInKES",
		)
	})

	View("default", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("Fraud Type")
		Attribute("FraudIncidentDate")
		Attribute("FraudReportDate")
		Attribute("Amount")
		Attribute("LossAmount")
		Attribute("CurrencyCode")
		Attribute("IncidentDetails")
		Attribute("ForensicInformation")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("AmountInKES")
	})

	View("mandatory", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AccountNumber")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("AmountInKES")
	})
})
