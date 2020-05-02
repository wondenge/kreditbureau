package design

import (
	. "goa.design/goa/v3/dsl"
)

var DailyPayment = Type("dailypayment", func() {
	Description("Daily Payment Information API")
	TypeName("DailyPayment")
	ContentType("application/json")

	// The date when the account is being affected and should
	// match the submission date.
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
	// Mandatory Field
	// Not more than 1 character
	Attribute("ClientType", String, "Client Type", func() {
		Meta("rpc:tag", "2")
	})

	// The Family Name or Surname
	// This field is Mandatory if Option “A” is selected in Field 5.1.2
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 5.1.2
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space,
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "3")
	})

	// The First Name
	// This field is Mandatory if Option “A” is selected in Field 5.1.2
	//
	// Alphanumeric
	// More than 1 character
	// Conditional to 5.1.2
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "4")
	})

	// The Given Name
	// This field is Mandatory if Option “A” is selected in Field 5.1.2
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "5")
	})

	// Other Name or Initials
	// This field is Mandatory if Option “A” is selected in Field 5.1.2
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "6")
	})

	// Name of the Non-Individual Entity
	//
	// Alphanumeric
	// Not more than 100 Characters
	// Conditional to 5.1.2
	Attribute("CompanyName", String, "Company Name", func() {
		Meta("rpc:tag", "7")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Alphanumeric
	// Mandatory
	// Not more than 3 Characters
	// If client type is B, option “005” Must be selected
	Attribute("PrimaryIDocumentType", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "8")
	})

	// The Number of the Primary Identification Document provided on
	// Opening of the Account
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "9")
	})

	// Account Number of the Loan/Overdraft Facility.
	// This must be a number that exists in the CRB Database.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "10")
	})

	// The currency code for the facility's currency ISO certified.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 3 characters
	Attribute("FacilityCurrency", String, "Currency of Facility", func() {
		Meta("rpc:tag", "11")
	})

	// For Loans, submit the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; submit the approved limits
	// For overdrafts with no limits, input the maximum amount overdrawn.
	// This field should not be Zero or Null.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 Characters
	Attribute("OriginalAmount", String, "Original Amount", func() {
		Meta("rpc:tag", "12")
	})

	// The sum total amount paid into the loan on the Payment Date.
	// The amount must be a positive value, which will result in
	// a reduction of Current Balance.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 Characters
	Attribute("PaymentAmount", String, "Payment Amount", func() {
		Meta("rpc:tag", "13")
	})

	// The date when a payment was made into the Loan.
	//
	// Mandatory Field
	// Date Field
	// Not more than 18 Characters
	// Can’t be in the future
	Attribute("PaymentDate", String, "Payment Date", func() {
		Meta("rpc:tag", "14")
	})

	// The new Current Balance in the loan account after the payment.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 Characters
	Attribute("CurrentBalance", String, "Current Balance", func() {
		Meta("rpc:tag", "15")
	})

	// The total number of instalments outstanding after this payment.
	// If the Loan is not in arrears after the payment, return report 0.
	//
	// Mandatory Field
	// Number field
	// Not more than 3 characters
	Attribute("InstallmentsInArrears", String, "Installments in Arrears", func() {
		Meta("rpc:tag", "16")
	})

	// The number of days that the loan is overdue.
	// If the loan is not in arrears after the payment, return 0.
	//
	// Mandatory Field
	// Number Field
	// Max Characters 5
	Attribute("DaysInArrears", String, "Days in Arrears", func() {
		Meta("rpc:tag", "17")
	})

	// The Account Status after the payment has been effected
	// A - Closed – No more admin processes running such as instalment demands
	// or interest charges to account, and no further facilities can be offered
	// on this account.
	// B - Dormant - no activity for 2 years. This applies for Overdraft/Current Accounts.
	// C - Write-Off – For facilities that don’t form part of the outstanding
	// portfolio in the Balance Sheet, but are still outstanding in the books
	// of accounts.
	// D - Legal -with legal officer in court
	// E - Collection- with collection bureau
	// F – Active - For facilities that form part of the outstanding portfolio,
	// and are reported in the Balance Sheet.
	// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
	// H –Settled – The facility has been cleared. This status can only be
	// used for revolving facilities
	// J – Called Up : The facility has been called up. Once the client
	// has paid up, the status should be updated to Option A, H; or otherwise Option L
	// K – Suspended – The facility has been put on hold for an indefinite period of time
	// L– Client Deceased
	// M – Deferred – This refers to facilities whose payments have been put on
	// hold for a definite period or in moratorium (Grace Period)
	// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
	// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved.
	//
	// Mandatory Field
	// Number Field
	// Not more than 1 Character
	// Lookup
	Attribute("AccountStatus", String, "Account Status", func() {
		Meta("rpc:tag", "18")
	})
})
