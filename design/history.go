package design

import (
	. "goa.design/goa/v3/dsl"
)

var CreditHistory = Type("credithistory ", func() {
	Description("Historical Credit Information Update API")
	TypeName("CreditHistory")
	ContentType("application/json")

	Attribute("SnapshotDate", String, func() {
		Description("Date when API was submitted to the Bureau")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "1")
	})

	// ClientTypeA – Individual Credit Consumer
	// ClientTypeB – Non-Individual Credit Consumer
	Attribute("ClientType", ArrayOf(Enum), func() {
		Enum(CreditHistoryCTA,
			CreditHistoryCTB,
		)
		Meta("rpc:tag", "2")
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
		Meta("rpc:tag", "9")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumber", String, func() {
		Description("Primary ID provided on Opening of the Account")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "10")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	// - Rest are Numeric
	Attribute("PINumber", String, func() {
		Description("Income Tax Number")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "11")
	})

	// This must be a number that exists in the CRB Database.
	Attribute("AccountNumber", String, func() {
		Description("Account Number of the Loan/Overdraft Facility")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "12")
	})

	Attribute("FacilityCurrency", String, func() {
		Description("The currency code for the facility's currency")
		MinLength(3)
		MaxLength(3)
		Default("KES")
		Meta("rpc:tag", "13")
	})

	// For Loans, the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; the approved limits.
	// For overdrafts with no limits, input the maximum amount overdrawn.
	Attribute("OriginalAmount", Int64, func() {
		Description("Original Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(1)
		Meta("rpc:tag", "14")
	})

	// The current balance in the account.
	Attribute("CurrentBalance", Int64, func() {
		Description("The current balance in the account")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "15")
	})

	// If the account is not in arrears, return 0.
	Attribute("OverdueBalance", Int64, func() {
		Description("Arrears amount in a facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "16")
	})

	// The number of days the account has been in arrears
	// (The Difference between the reporting date and the overdue date, calculated in days).
	// If the account is not in arrears, return 0.
	Attribute("DaysInArrears", Int, func() {
		Description("The number of days the account has been in arrears")
		MinLength(1)
		MaxLength(5)
		Minimum(0)
		Meta("rpc:tag", "17")
	})

	// If Field 5.3.17 is Zero (0) return 0.
	// This field cannot be 0 if Field 5.3.17 is greater than 0.
	Attribute("InstalmentsInArrears", Int, func() {
		Description("The Number of missed Instalments in the Facility")
		MinLength(1)
		MaxLength(3)
		Minimum(0)
		Meta("rpc:tag", "18")
	})

	// The Account status as at the time of reporting.
	// A - Closed – No more admin processes running such as instalment demands or
	// interest charges to account, and no further facilities can be offered on this
	// account.
	// B - Dormant - no activity for 2 years. This applies for Overdraft/Current Accounts.
	// C - Write-Off – For facilities that don’t form part of the outstanding portfolio
	// in the Balance Sheet, but are still outstanding in the books of accounts.
	// D - Legal -with legal officer in court
	// E - Collection- with collection bureau
	// F – Active - For facilities that form part of the outstanding portfolio, and are reported in the Balance Sheet.
	// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
	// H –Settled – The facility has been cleared. This status can only be used for revolving facilities
	// J – Called Up: The facility has been called up.
	// Once the client has paid up, the status should be updated to Option A, H; or otherwise Option L
	// K – Suspended – The facility has been put on hold for an indefinite period of time
	// L– Client Deceased
	// M – Deferred – This refers to facilities whose payments have been put on hold
	// for a definite period or in moratorium (Grace Period)
	// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
	// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
	//
	// The account cannot be updated until the dispute is resolved.
	Attribute("AccountStatus", String, func() {
		Description("Account Status")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("P")
		Meta("rpc:tag", "19")
	})

	// Classification of the Account
	// A = Normal
	// B = Watch
	// C = Substandard
	// D = Doubtful
	// E = Loss
	Attribute("PrudentialAssetClassification", String, func() {
		Description("Prudential Asset Classification")
		Enum("A", "B", "C", "D", "E")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("E")
		Meta("rpc:tag", "20")
	})

	//
	// If no amount has been paid into the account, leave the field blank.
	// This date is equal or less than field 5.3.1
	Attribute("LastPaymentDate", String, func() {
		Description("The date when the facility last received a payment.")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "21")
	})

	// The Account Product Type options:
	// A – Overdraft
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
	Attribute("AccountProductType", String, "Account Product Type", func() {
		Description("Account Product Type")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("L")
		Meta("rpc:tag", "22")
	})

	Attribute("InstalmentAmount", Int64, func() {
		Description("Instalment amount for Loans.")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "23")
	})

	Attribute("AdditionalInformation", String, func() {
		Description("Additional Information on the facility")
		MinLength(1)
		MaxLength(100)
		Meta("rpc:tag", "24")
	})

})

var StoredCreditHistory = ResultType("application/vnd.goa.credithistory", func() {
	TypeName("StoredCreditHistory")
	Attributes(func() {
		Extend(CreditHistory)
		Required(
			"SnapshotDate",
			"ClientType",
			"PrimaryIDocumentType",
			"PrimaryIDocumentNumber",
			"AccountNumber",
			"FacilityCurrency",
			"OriginalAmount",
			"CurrentBalance",
			"OverdueBalance",
			"DaysInArrears",
			"InstalmentsInArrears",
			"AccountStatus",
			"PrudentialAssetClassification",
			"LastPaymentDate",
			"AccountProductType",
			"InstalmentAmount",

		)
	})

	View("default", func() {
		Attribute("SnapshotDate")
		Attribute("ClientType")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("PINumber")
		Attribute("AccountNumber")
		Attribute("FacilityCurrency")
		Attribute("OriginalAmount")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("AccountStatus")
		Attribute("PrudentialAssetClassification")
		Attribute("LastPaymentDate")
		Attribute("AccountProductType")
		Attribute("InstalmentAmount")
		Attribute("AdditionalInformation")
	})

	View("mandatory", func() {
		Attribute("SnapshotDate")
		Attribute("ClientType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("AccountNumber")
		Attribute("FacilityCurrency")
		Attribute("OriginalAmount")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("AccountStatus")
		Attribute("PrudentialAssetClassification")
		Attribute("LastPaymentDate")
		Attribute("AccountProductType")
		Attribute("InstalmentAmount")
	})
})

var CreditHistoryCTA = Type("CreditHistoryCTA", func() {
	Description("A – Individual Credit Consumer")

	Attribute("Surname", String, func() {
		Description("Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "6")
	})

	Required("Surname", "Forename1", "Forename2", "Forename3")

})

var CreditHistoryCTB = Type("CreditHistoryCTB", func() {
	Description("B – NonIndividual Credit Consumer")

	Attribute("RegisteredName", String, func() {
		Description("The Name as Registered by an Approved Registrar") // Company Name
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "7")
	})

	Attribute("TradingName", String, "Trading Name", func() {
		Description("The Business or Trading Name")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "8")
	})

	Required("RegisteredName")

})
