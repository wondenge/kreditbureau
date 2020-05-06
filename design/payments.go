package design

import (
	. "goa.design/goa/v3/dsl"
)

var DailyPayment = Type("dailypayment", func() {
	Description("Daily Payment Information API")
	TypeName("DailyPayment")
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
		Enum(DailyPaymentCTA,
			DailyPaymentCTB,
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
		Meta("rpc:tag", "8")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocNumber", String, func() {
		Description("Primary ID Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "9")
	})

	// This must be a number that exists in the CRB Database.
	Attribute("AccountNumber", String, "Account Number", func() {
		Description("Account Number of the Loan/Overdraft Facility")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "10")
	})

	Attribute("FacilityCurrency", String, "Currency of Facility", func() {
		Description("Currency code for the facility's currency")
		MinLength(3)
		MaxLength(3)
		Default("KES")
		Meta("rpc:tag", "11")
	})

	// For Loans, submit the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; submit the approved limits
	// For overdrafts with no limits, input the maximum amount overdrawn.
	Attribute("OriginalAmount", Int64, func() {
		Description("Original Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(1) // This field should not be Zero or Null.
		Meta("rpc:tag", "12")
	})

	Attribute("PaymentAmount", Int64, "Payment Amount", func() {
		Description("Total amount paid into the loan on the payment date")
		MinLength(1)
		MaxLength(16)
		Minimum(1) // The amount must be a positive value, which will result in a reduction of Current Balance.
		Meta("rpc:tag", "13")
	})

	Attribute("PaymentDate", String, func() {
		Description("Date payment was made into the Loan")
		MaxLength(18)
		Format(FormatDate)
		Meta("rpc:tag", "14")
	})

	// The new Current Balance in the loan account after the payment.
	Attribute("CurrentBalance", Int64, func() {
		Description("New Current Balance in the loan account")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "15")
	})

	// If the Loan is not in arrears after the payment, return report 0.
	Attribute("InstallmentsInArrears", String, func() {
		Description("Total number of instalments outstanding") // Installments in Arrears
		MinLength(1)
		MaxLength(3)
		Minimum(0)
		Meta("rpc:tag", "16")
	})

	// If the loan is not in arrears after the payment, return 0.
	Attribute("DaysInArrears", Int, func() {
		Description("Number of days loan is overdue")
		MinLength(1)
		MaxLength(5)
		Minimum(0)
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
	Attribute("AccountStatus", String, func() {
		Description("Account Status")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("P")
		Meta("rpc:tag", "18")
	})
})

var StoredDailyPayment = ResultType("application/vnd.goa.dailypayment", func() {
	TypeName("StoredDailyPayment")
	Attributes(func() {
		Extend(DailyPayment)
		Required(
			"SnapshotDate",
			"ClientType",
			"PrimaryIDocumentType",
			"PrimaryIDocNumber",
			"AccountNumber",
			"FacilityCurrency",
			"OriginalAmount",
			"PaymentAmount",
			"PaymentDate",
			"CurrentBalance",
			"InstallmentsInArrears",
			"DaysInArrears",
			"AccountStatus",
		)
	})
	View("default", func() {
		Attribute("SnapshotDate")
		Attribute("ClientType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyName")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocNumber")
		Attribute("AccountNumber")
		Attribute("FacilityCurrency")
		Attribute("OriginalAmount")
		Attribute("PaymentAmount")
		Attribute("PaymentDate")
		Attribute("CurrentBalance")
		Attribute("InstallmentsInArrears")
		Attribute("DaysInArrears")
		Attribute("AccountStatus")
	})

	View("mandatory", func() {
		Attribute("SnapshotDate")
		Attribute("ClientType")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocNumber")
		Attribute("AccountNumber")
		Attribute("FacilityCurrency")
		Attribute("OriginalAmount")
		Attribute("PaymentAmount")
		Attribute("PaymentDate")
		Attribute("CurrentBalance")
		Attribute("InstallmentsInArrears")
		Attribute("DaysInArrears")
		Attribute("AccountStatus")
	})
})

var DailyPaymentCTA = Type("DailyPaymentCTA", func() {
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

var DailyPaymentCTB = Type("DailyPaymentCTB", func() {
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
