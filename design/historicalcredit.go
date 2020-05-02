package design

import (
	. "goa.design/goa/v3/dsl"
)

var HistoricalCredit = Type("historicalcredit", func() {
	Description("Historical Credit Information Update API")
	TypeName("HistoricalCredit")
	ContentType("application/json")

	// The date when the API was submitted to the Bureau.
	//
	// Date Field
	// Mandatory Field
	// Can’t be in future
	// Not more than 8 Characters
	Attribute("SnapshotDate", String, "Snapshot Date", func() {
		Meta("rpc:tag", "1")
	})

	// Acceptable Values:
	// A – Individual Credit Consumer
	// B – Non-Individual Credit Consumer
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	// Mandatory Field
	Attribute("ClientType", String, "Client Type", func() {
		Meta("rpc:tag", "2")
	})

	// Mandatory if Client Type is Individual.
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 5.3.2
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "3")
	})

	// The First Name
	// This field is Mandatory if Option “A” is selected in Field 5.3.2
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 5.3.2
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "4")
	})

	// The Given Name
	// This field is Mandatory if Option “A” is selected in Field 5.3.2
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "5")
	})

	// Other Name or Initials
	// This field is Mandatory if Option “A” is selected in Field 5.3.2
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "6")
	})

	// The Name as Registered  by an Approved Registrar
	// This field is Mandatory if Option “B” is selected in Field 5.3.2
	//
	// Alphanumeric
	// Not more than 70 Characters
	Attribute("RegisteredName", String, "Registered Name", func() {
		Meta("rpc:tag", "7")
	})

	// The Business or Trading Name
	//
	// Alphanumeric
	// Not more than 70 Characters
	Attribute("TradingName", String, "Trading Name", func() {
		Meta("rpc:tag", "8")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Alphanumeric
	// Mandatory Field
	// Not more than 3 characters
	// If Option “B” in 5.3.2, then option “005” Must be selected
	Attribute("PrimaryIDocument Type", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "9")
	})

	// The Number of the Primary Identification Document provided on Opening of the Account
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("Primary IDocument Number", String, "Primary Identification Document Number", func() {
		Meta("rpc:tag", "10")
	})

	// Income Tax Number
	//
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	// - Rest are Numeric
	Attribute("PINumber", String, "PIN Number", func() {
		Meta("rpc:tag", "11")
	})

	// Account Number of the Loan/Overdraft Facility.
	// This must be a number that exists in the CRB Database.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "12")
	})

	// The currency code for the facility's currency.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 3 characters
	Attribute("FacilityCurrency", String, "Currency of Facility", func() {
		Meta("rpc:tag", "13")
	})

	// For Loans, the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; the approved limits.
	// For overdrafts with no limits, input the maximum amount overdrawn.
	//
	// Mandatory Field
	// Currency Field
	// Not more than 16 characters
	Attribute("OriginalAmount", String, "Original Amount", func() {
		Meta("rpc:tag", "14")
	})

	// The current balance in the account.
	//
	// Mandatory Field
	// Currency Field
	// Not more than 16 Characters
	Attribute("CurrentBalance", String, "Current Balance", func() {
		Meta("rpc:tag", "15")
	})

	// Arrears amount in a facility.
	// If the account is not in arrears, return 0.
	//
	// Mandatory Field
	// Currency Field
	// Not more than 16 Characters
	Attribute("OverdueBalance", String, "Overdue Balance", func() {
		Meta("rpc:tag", "16")
	})

	// The number of days the account has been in arrears (The Difference between
	// the reporting date and the overdue date, calculated in days).
	// If the account is not in arears, return 0.
	//
	// Number Field
	// Not more than 5 characters
	// Mandatory Field
	Attribute("DaysInArrears", String, "Days in Arrears", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of missed Instalments in the Facility.
	// If Field 5.3.17 is Zero (0) return 0.
	// This field cannot be 0 if Field 5.3.17 is greater than 0.
	//
	// Number Field
	// Not more than 3 characters
	// Mandatory Field
	Attribute("InstalmentsInArrears", String, "Instalments in Arrears", func() {
		Meta("rpc:tag", "18")
	})

	// The Account status as at the time of reporting.
	// A - Closed – No more admin processes running such as instalment demands or
	// interest charges to account, and no further facilities can be offered on this
	// account.
	// B - Dormant - no activity for 2 years.
	// This applies for Overdraft/Current Accounts.
	// C - Write-Off – For facilities that don’t form part of the outstanding portfolio
	// in the Balance Sheet, but are still outstanding in the books of accounts.
	// D - Legal -with legal officer in court
	// E - Collection- with collection bureau
	// F – Active - For facilities that form part of the outstanding portfolio, and
	// are reported in the Balance Sheet.
	// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
	// H –Settled – The facility has been cleared.
	// This status can only be used for revolving facilities
	// J – Called Up: The facility has been called up.
	// Once the client has paid up, the status should be updated to Option A, H; or otherwise Option L
	// K – Suspended – The facility has been put on hold for an indefinite period of time
	// L– Client Deceased
	// M – Deferred – This refers to facilities whose payments have been put on hold
	// for a definite period or in moratorium (Grace Period)
	// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
	// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved.
	//
	// Alphanumeric
	// Mandatory Field
	// Not more than 1 character
	// lookup
	Attribute("AccountStatus", String, "Account Status", func() {
		Meta("rpc:tag", "19")
	})

	// Classification of the Account
	// A = Normal
	// B = Watch
	// C = Substandard
	// D = Doubtful
	// E = Loss
	// Mandatory Field
	// Alphanumeric
	// Not more than 1 character
	// Lookup
	Attribute("PrudentialAssetClassification", String, "Prudential Asset Classification", func() {
		Meta("rpc:tag", "20")
	})

	// The date when the facility last received a payment.
	// If no amount has been paid into the account, leave the field blank.
	// This date is equal or less than field 5.3.1
	//
	// Not more than 8 characters
	// Not in future
	// Date field
	Attribute("LastPaymentDate", String, "Date of Last Payment", func() {
		Meta("rpc:tag", "21")
	})

	// The Account Product Type options:
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
	// Mandatory
	// Alphanumeric
	// Not more than 1 character
	Attribute("AccountProductType", String, "Account Product Type", func() {
		Meta("rpc:tag", "22")
	})

	// Instalment amount for Loans.
	//
	// Currency Field
	// Not more than 16 characters
	Attribute("InstalmentAmount", String, "Instalment Amount", func() {
		Meta("rpc:tag", "23")
	})

	// Any additional information on the facility.
	//
	// Alphanumeric
	// Not more than 100 Characters
	Attribute("AdditionalInformation", String, "Additional Information", func() {
		Meta("rpc:tag", "24")
	})

})
