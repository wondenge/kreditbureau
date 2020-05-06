package design

import (
	. "goa.design/goa/v3/dsl"
)

var CreditApplication = Type("creditapplication", func() {

	Description("")
	TypeName("CreditApplication")
	ContentType("application/json")

	Attribute("LendersRegisteredName", String, func() {
		Description("Name of Lender Reporting  Loan Application") // As registered with the Registrar of companies.
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
		Description("Lenders Branch Name")
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
		Meta("rpc:tag", "4")
	})

	// A – Individual Credit Consumer
	// B – Non-Individual Credit Consumer
	Attribute("ClientType", ArrayOf(Enum), func() {
		Description("Type of Client")
		Enum(IndividualCreditConsumer, CompanyCreditConsumer)
		Meta("rpc:tag", "5")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocument", Int, func() {
		Description("Primary Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "12")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDoc Number", String, func() {
		Description("Primary Identification Doc Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "13")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("SecondaryIDocument", String, func() {
		Description("Secondary Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "14")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("SecondaryIDocument Number", String, func() {
		Description("Secondary Identification Document Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "15")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("OtherIDocument", String, func() {
		Description("Other Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "16")

		Attribute("OtherIDocument Number", String, func() {
			Description("Other Identification Document Number")
			MinLength(1)
			MaxLength(20)
			Meta("rpc:tag", "17")
		})
		Required("OtherIDocument")
	})

	// This is for Client-Centric systems.
	Attribute("ClientNumber", String, func() {
		Description("Client Reference Number") // Linking to Applicant’s account the Lender’s system
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "18")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, func() {
		Description("Revenue Authority Personal Income Tax Number")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "19")
	})

	// Account Number Linking to Applicant’s Account in the Lender’s System.
	// The Account Number Is the Same as Client Number for Account Centric Systems.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "20")
	})

	// The Date the Application was made to the Lender.
	Attribute("ApplicationDate", String, func() {
		Description("Application Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "21")
	})

	Attribute("ApplicationNumber", String, func() {
		Description("Internal Application Reference number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "22")
	})

	// U - Unsecured
	// S - Secured
	Attribute("FacilityApplicationType", String, func() {
		Description("Facility application Type")
		Enum("U", "S")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "23")
	})

	Attribute("ApplicationAmount", Int64, func() {
		Description("Application Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(1)
		Meta("rpc:tag", "24")
	})

	Attribute("ApplicationCurrency", String, func() {
		Description("Application ISO Currency")
		MinLength(3)
		MaxLength(3)
		Default("KES") // Default is KES
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
	Attribute("ProductType", ArrayOf(Enum), func() {
		Description("Product Type")
		Enum(
			CAOverdraft,                   // A –Overdraft
			CACreditCards,                 // B – Credit Cards
			CABusinessWorkingCapitalLoans, // C – Business Working Capital Loans
			CABusinessExpansionLoans,      // D – Business Expansion Loans
			CAMortgageLoans,               // E – Mortgage Loans
			CAAssetFinanceLoans,           // F – Asset Finance Loans
			CATradeFinance,                // G – Trade Finance Facilities
			CAPersonalLoans,               // H – Personal Loans
			CAMobileLoan,                  // I – Mobile Loan
			CAInsurancePremiumFinancing,   // J – Insurance Premium Financing
			CAGroupLoans,                  // K – Group Loans
			CAUnclearedEffects,            // L – Uncleared Effects
		)
		Meta("rpc:tag", "26")
	})

	// A - Pending
	// B - Awaiting Docs
	// C - Securities Perfection
	// D - Declined by the Bank
	// E - Withdrawn
	// F - Approved
	// G - Pending disbursement
	// H – Fully Disbursed
	// I – Customer Declines Offer
	Attribute("ApplicationStatus", String, func() {
		Description("Current status of the Application")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "I")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("I")
		Meta("rpc:tag", "28")
	})

	// A= Over indebted
	// B = Failed credit criteria
	// C= Failed verification – (e.g. Income/Employer could not be verified)
	// D = Lacks ability to repay
	// E = Weak Credit History
	// F = Insufficient Collateral
	Attribute("ApplicationDeclineReason", String, func() {
		Description("Application Decline Reason codes")
		Enum("A", "B", "C", "D", "E", "F")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("F")
		Meta("rpc:tag", "29")
	})

	Attribute("ApplicationStatusDate", String, func() {
		Description("Application Status Date") // Date of status of Update
		Format(FormatDate)
		Meta("rpc:tag", "30")
	})

	// A = Additional Income provided
	// B = Credit Profile Updated
	// (e.g. if client presents proof that an account is paid up or status on credit bureau is out dated)
	// C= Additional deposit provided by client
	Attribute("StatusUpdateReasonCode", String, func() {
		Description("Status Update Reason Code")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Meta("rpc:tag", "31")
	})

})

var StoredCreditApplication = ResultType("application/vnd.goa.creditapplication", func() {
	TypeName("StoredCreditApplication")
	Attributes(func() {
		Extend(CreditApplication)
		Required("LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"ClientType",
			"PrimaryIDocument",
			"PrimaryIDoc Number",
			"ApplicationDate",
			"ApplicationNumber",
			"FacilityApplicationType",
			"ApplicationAmount",
			"ApplicationCurrency",
			"ProductType",
			"ApplicationStatus",
			"ApplicationStatusDate")
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
		Attribute("StatusUpdateReasonCode")
	})
	View("mandatory", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("ClientType")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDoc Number")
		Attribute("ApplicationDate")
		Attribute("ApplicationNumber")
		Attribute("FacilityApplicationType")
		Attribute("ApplicationAmount")
		Attribute("ApplicationCurrency")
		Attribute("ProductType")
		Attribute("ApplicationStatus")
		Attribute("ApplicationStatusDate")
	})
})

// A – Individual Credit Consumer
var IndividualCreditConsumer = Type("", func() {
	Description("Client Type A: Individual Credit Consumer")

	Attribute("Surname", String, func() {
		Description("Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "6")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "7")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "8")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "9")
	})

	Required("Surname", "Forename1", "Forename2", "Forename3")

})

// B – Non-Individual Credit Consumer
var CompanyCreditConsumer = Type("CompanyCreditConsumer", func() {
	Description("Client Type B: Non-Individual Credit Consumer")

	Attribute("RegisteredName", String, func() {
		Description("Name as Registered with the Registrar of Companies")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "10")
	})

	Attribute("TradingName", String, func() {
		Description("Business or Trading Name")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "11")
	})

	Required("RegisteredName", "TradingName")
})
