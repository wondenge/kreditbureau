package design

import (
	. "goa.design/goa/v3/dsl"
)

var CreditApplication = Type("creditapplication", func() {

	Description("")
	TypeName("CreditApplication")
	ContentType("application/json")

	// The Name of Lender Reporting the Loan Application, as registered
	// with the Registrar of companies.
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

	// The Lenders Branch Name, where the Loan Application is reported to
	// have taken place.
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
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "4")
	})

	// The following Options available:
	// A – Individual Credit Consumer
	// B – Non-Individual Credit Consumer
	//
	// Mandatory
	// Lookup
	// Alphanumeric
	// Not more than 1 Character
	Attribute("ClientType", String, "Type of Client", func() {
		Meta("rpc:tag", "5")
	})

	// The Family Name or Surname
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 4.8.5
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "6")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 4.8.5
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "7")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "8")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "9")
	})

	// The Name as Registered with the Registrar of Companies.
	//
	// Alphanumeric
	// Not more than 70 Characters
	// Mandatory if 4.8.5 is Non-individual
	Attribute("RegisteredName", String, "Registered Name", func() {
		Meta("rpc:tag", "10")
	})

	// The Business or Trading Name
	//
	// Alphanumeric
	// Mandatory if 4.8.5 is Non-individual
	// Not more than 70 Characters
	Attribute("TradingName", String, "Trading Name", func() {
		Meta("rpc:tag", "11")
	})

	// The Applicant’s Primary Identification document Type Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	Attribute("PrimaryIDocument", String, "Primary Identification Document", func() {
		Meta("rpc:tag", "12")
	})

	// The Identification document Number.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDoc Number", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "13")
	})

	// The Applicant’s Secondary Identification Document Type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Data is Alpha Numeric
	// Not more than 3 Characters
	Attribute("SecondaryIDocument", String, "Secondary Identification Document", func() {
		Meta("rpc:tag", "14")
	})

	// Mandatory if Secondary Identification Document type is provided.
	//
	// For ID: Numeric Characters, between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters
	// Not more than 6 characters
	Attribute("SecondaryIDocument Number", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "15")
	})

	// The Applicant’s Other Identification document type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Data is Alpha Numeric
	// Not more than 3 Characters
	Attribute("OtherIDocument", String, "Other Identification Document", func() {
		Meta("rpc:tag", "16")
	})

	// Mandatory if Other Identification document type is provided.
	//
	// Data is Alphanumeric
	// Not more than 20 Characters
	Attribute("OtherIDocument Number", String, "Other Identification Document Number", func() {
		Meta("rpc:tag", "17")
	})

	// Client Reference Number Linking to Applicant’s account the Lender’s system.
	// This is for Client-Centric systems.
	//
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "18")
	})

	// Revenue Authority Personal Income Tax Number.
	//
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, "PIN Number", func() {
		Meta("rpc:tag", "19")
	})

	// Account Number Linking to Applicant’s Account in the Lender’s System.
	// The Account Number Is the Same as Client Number for Account Centric Systems.
	//
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "20")
	})

	// The Date the Application was made to the Lender.
	//
	// Date Field
	// Cannot be in the future
	// Not more than 8 characters
	Attribute("ApplicationDate", String, "Application Date", func() {
		Meta("rpc:tag", "21")
	})

	// Internal Application Reference number.
	//
	// Alphanumeric
	// Mandatory
	// Not more than 20 characters
	Attribute("ApplicationNumber", String, "Application Number", func() {
		Meta("rpc:tag", "22")
	})

	// The Type of Facility Options:
	// U - Unsecured
	// S - Secured
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 Character
	Attribute("FacilityApplicationType", String, "Facility application Type", func() {
		Meta("rpc:tag", "23")
	})

	// Amount applied for.
	//
	// Amount Field
	// Cannot be zero or negative
	// Not more than 16 Character
	// Mandatory Field
	Attribute("ApplicationAmount", String, "Application Amount", func() {
		Meta("rpc:tag", "24")
	})

	// ISO Currency in which facility is requested for.
	// Default is KES
	//
	// ISO
	// Mandatory
	// Alphanumeric
	// Not more than 3 characters
	Attribute("ApplicationCurrency", String, "Application Currency", func() {
		Meta("rpc:tag", "25")
	})

	// The Account Product Type :
	// A –Overdraft
	// B – Credit Cards
	// C – Business Working Capital Loans
	// D – Business Expansion Loans
	// E – Mortgage Loans
	// F – Asset Finance Loans
	// G – Trade Finance Facilities
	// H – Personal Loans
	// I – Mobile Loan
	// J – Insurance Premium Financing
	// K – Group Loans
	// L – Uncleared Effects
	//
	// Lookup
	// Mandatory
	// Not more than 1 character
	Attribute("ProductType", String, "Product Type", func() {
		Meta("rpc:tag", "26")
	})

	// Number of Months Applied for.
	// Conditional with Product Type
	// Not Mandatory for Credit Cards
	//
	// Number Field
	// Not more than 3 Characters
	// Cannot be zero or negative
	Attribute("LoanTerm", String, "Term of Loan", func() {
		Meta("rpc:tag", "27")
	})

	// The Current status of the Application.
	// Options are:
	// A - Pending
	// B - Awaiting Docs
	// C - Securities Perfection
	// D - Declined by the Bank
	// E - Withdrawn
	// F - Approved
	// G - Pending disbursement
	// H – Fully Disbursed
	// I – Customer Declines Offer
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 Character
	Attribute("ApplicationStatus", String, "Application Status", func() {
		Meta("rpc:tag", "28")
	})

	// Required if application is Declined.
	// A= Over indebted
	// B = Failed credit criteria
	// C= Failed verification – (e.g. Income/Employer could not be verified)
	// D = Lacks ability to repay
	// E = Weak Credit History
	// F = Insufficient Collateral
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 Character
	Attribute("ApplicationDeclineReason", String, "Application Decline Reason codes", func() {
		Meta("rpc:tag", "29")
	})

	// Date of status of Update
	// Date Field
	// Mandatory Field
	// Cannot be in the future
	Attribute("ApplicationStatusDate", String, "Application Status Date", func() {
		Meta("rpc:tag", "30")
	})

})

var StoredCreditApplication = ResultType("application/vnd.goa.creditapplication", func() {
	TypeName("StoredCreditApplication")
	Attributes(func() {
		Extend(CreditApplication)
		Required()
	})
	View("default", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("ClientType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDoc Number")
		Attribute("SecondaryIDocument")
		Attribute("SecondaryIDocument Number")
		Attribute("OtherIDocument")
		Attribute("OtherIDocument Number")
		Attribute("ClientNumber")
		Attribute("PINumber")
		Attribute("AccountNumber")
		Attribute("ApplicationDate")
		Attribute("ApplicationNumber")
		Attribute("FacilityApplicationType")
		Attribute("ApplicationAmount")
		Attribute("ApplicationCurrency")
		Attribute("ProductType")
		Attribute("LoanTerm")
		Attribute("ApplicationStatus")
		Attribute("ApplicationDeclineReason")
		Attribute("ApplicationStatusDate")
	})
})
