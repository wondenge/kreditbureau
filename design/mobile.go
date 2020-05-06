package design

import (
	. "goa.design/goa/v3/dsl"
)

var MobileFacility = Type("mobilefacility", func() {
	Description("Mobile Facilities API")
	TypeName("MobileFacility")
	ContentType("application/json")

	Attribute("Surname", String, func() {
		Description("Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "1")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	// Name of a business if the individual runs a business as a sole proprietor that is not registered;
	// or if the sole proprietor took a credit facility to finance the business.
	Attribute("TradingAs", String, func() {
		Description("Name of a business run by sole proprietor")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	Attribute("DateOfBirth", String, func() {
		Description("The Date of Birth of the Customer") // Not a future date, Not under 18, Not over 100 years
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "6")
	})

	Attribute("ClientNumber", String, func() {
		Description("Unique identification of a client") // Within the lenders Core System
		Minimum(1)
		MaxLength(20)
		Meta("rpc:tag", "7")
	})

	Attribute("AccountNumber", String, func() {
		Description("Account Number to uniquely identify specific mobile facility")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "8")
	})

	Attribute("Gender", String, func() {
		Description("Gender of the Borrower")
		Enum("M", "F")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "9")
	})

	Attribute("Nationality", String, func() {
		Description("ISO Country Code for the Consumer’s Nationality")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "10")
	})

	// 	001 = NationalID IDType
	//	002 = Passport
	//	003 = AlienRegistration
	//	004 = ServiceID
	Attribute("PrimaryIDocumentType", Int, func() {
		Description("Primary Identification Document Type")
		Enum(001, 002, 003, 004)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "11")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumber", String, func() {
		Description("Primary Identification Document Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "12")
	})

	// The Mobile Phone Number of the Client used
	// to grant the loan facility.
	Attribute("MobilePhoneNumber", String, func() {
		Description("Mobile Phone Number")
		MaxLength(15)
		Meta("rpc:tag", "13")
	})

	// The Account Product Type :
	// 	A  = "Overdraft"
	//	B  = "Credit Cards"
	//	C  = "Business Working Capital Loans"
	//	D  = "Business Expansion Loans"
	//	E  = "Mortgage Loans"
	//	F  = "Asset Finance Loans"
	//	G  = "Trade Finance Facilities"
	//	H  = "Personal Loans"
	//	I  = "Mobile Loan"
	//	J  = "Insurance Premium Financing"
	//	K  = "Group Loans"
	//	L  = "Uncleared Effects"
	// The default product type be Mobile loan
	Attribute("AccountProductType", String, func() {
		Description("Account Product Type")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("L")
		Meta("rpc:tag", "14")
	})

	// If the account is not in arrears, report the next instalment date.
	// If the account is in arrears, report the overdue date.
	Attribute("InstalmentDueDate", String, func() {
		Description("Instalment Due Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "15")
	})

	Attribute("OriginalAmount", Int, func() {
		Description("The principal Amount issued.")
		MinLength(1)
		MaxLength(16)
		Minimum(1)
		Meta("rpc:tag", "16")
	})

	Attribute("CurrencyOfFacility", String, func() {
		Description("The currency code facility's currency")
		MinLength(3)
		MaxLength(3)
		Default("KES")
		Meta("rpc:tag", "17")
	})

	// The current balance in Kenya shillings,
	// if the Currency is not Kenya shillings.
	Attribute("CurrentBalanceInKES", Int64, func() {
		Description("The current balance in Kenya shillings")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "18")
	})

	// The current balance in the mobile account
	// in the original currency.
	Attribute("CurrentBalance", Int64, func() {
		Description("The current balance in the mobile account")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "19")
	})

	// If account is not in arrears report Zero (0)
	Attribute("OverdueBalance", Int64, func() {
		Description("Overdue Amount in a facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "20")
	})

	// This Field is Mandatory if Field 5.2.20 is greater than Zero (0)
	Attribute("OverdueDate", String, func() {
		Description("The date when the account fell Overdue") // Date can’t be in the future
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "21")
	})

	// The number of days the account has been in arrears
	// (Difference between reporting date and overdue date, calculated in days)
	// If the account is not in arrears, report Zero (0).
	Attribute("DaysInArrears", Int, func() {
		Description("Number of days account has been in arrears ")
		MinLength(1)
		MaxLength(5)
		Minimum(0)
		Meta("rpc:tag", "22")
	})

	// If Field 5.2.22 is Zero (0) return 0.
	// This field cannot be 0 if Field 5.2.22 is greater than 0.
	// Cannot be 0 if field 5.2.22 is greater than 0
	Attribute("InstalmentsInArrears", Int, func() {
		Description("Number of missed facility Instalments")
		MinLength(1)
		MaxLength(5)
		Minimum(0)
		Meta("rpc:tag", "23")
	})

	// Classification of the Assets.
	// 	A  = "Normal (0-30)"
	//	B  = "Watch (>30-90)"
	//	C  = "Substandard(>90-180)"
	//	D  = "Doubtful(>180-360)"
	//	E  = "Loss(>360)"
	Attribute("PrudentialRiskClassification", String, func() {
		Description("Prudential Risk Classification")
		Enum("A", "B", "C", "D", "E")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("E")
		Meta("rpc:tag", "24")
	})

	// The Account status as at the time of reporting.
	//
	// A AccountStatus = "Closed"
	// No more admin processes running such as instalment demands or interest charges
	// to account, and no further facilities can be offered on this account.
	//
	// B = "Dormant"
	// No activity for 2 years. This applies for Overdraft/Current Accounts.
	//
	//
	// C = "Write-Off"
	// For facilities that don’t form part of the outstanding portfolio in the
	// Balance Sheet, but are still outstanding in the books of accounts.
	//
	// D = "Legal"
	// With legal officer in court
	//
	// E = "Collection"
	// With collection bureau
	//
	// F = "Active"
	// For facilities that form part of the outstanding portfolio,
	// and are  reported in the Balance Sheet.
	//
	// G = "Facility Rescheduled"
	// For Rescheduled/Restructured Facilities
	//
	// H = "Settled"
	// The facility has been cleared. This status can only be used
	// for revolving facilities.
	//
	// J = "Called Up"
	// The facility has been called up.Once the client has paid up, the
	// status should be updated to Option A, H; or otherwise Option L
	//
	// K = "Suspended"
	// The facility has been put on hold for an indefinite period of time
	//
	// L = "Client Deceased"
	//
	// M = "Deferred"
	// This refers to facilities whose payments have been put on hold
	// for a definite period or in moratorium (Grace Period)
	//
	// N = "Not Updated"
	// This status is reserved for CRBs (if last record status is not CLOSED)
	//
	// P = "Disputed"
	// Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved.
	Attribute("AccountStatus", String, func() {
		Description("Account Status")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("P")
		Meta("rpc:tag", "25")
	})

	// .
	Attribute("AccountStatusDate", String, func() {
		Description("The date of the status")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "26")
	})

	// This is the initial contractual period.
	Attribute("RepaymentPeriod", Int, "Repayment Period", func() {
		Description("Repayment Period for the Facility in months")
		MinLength(1)
		MaxLength(3)
		Minimum(0)
		Meta("rpc:tag", "27")
	})

	Attribute("DisbursementDate", String, func() {
		Description("Date of facility drawdown")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "28")
	})

	Attribute("InstalmentAmount", Int64, func() {
		Description("Instalment amount for the loan")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "29")
	})

	// This field should be null if Option “G” is selected in Field 5.2.14.
	Attribute("LatestPaymentDate", String, func() {
		Description("Date when payments were last received into the facility")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "30")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last Payment Amount received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // Report Zero (0) if no payment was received.
		Meta("rpc:tag", "31")
	})
})

var StoredMobileFacility = ResultType("application/vnd.goa.mobilefacility", func() {
	TypeName("StoredMobileFacility")
	Attributes(func() {
		Extend(MobileFacility)
		Required(
			"Surname",
			"Forename1",
			"TradingAs",
			"DateOfBirth",
			"AccountNumber",
			"Gender",
			"Nationality",
			"PrimaryIDocumentType",
			"PrimaryIDocumentNumber",
			"MobilePhoneNumber",
			"AccountProductType",
			"InstalmentDueDate",
			"OriginalAmount",
			"CurrencyOfFacility",
			"CurrentBalanceInKES",
			"CurrentBalance",
			"OverdueBalance",
			"OverdueDate",
			"DaysInArrears",
			"InstalmentsInArrears",
			"PrudentialRiskClassification",
			"AccountStatus",
			"AccountStatusDate",
			"RepaymentPeriod",
			"DisbursementDate",
			"InstalmentAmount",
			"LatestPaymentDate",
			"LastPaymentAmount",

		)
	})
	View("default", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("TradingAs")
		Attribute("DateOfBirth")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("MobilePhoneNumber")
		Attribute("AccountProductType")
		Attribute("InstalmentDueDate")
		Attribute("OriginalAmount")
		Attribute("CurrencyOfFacility")
		Attribute("CurrentBalanceInKES")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("OverdueDate")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("PrudentialRiskClassification")
		Attribute("AccountStatus")
		Attribute("AccountStatusDate")
		Attribute("RepaymentPeriod")
		Attribute("DisbursementDate")
		Attribute("InstalmentAmount")
		Attribute("LatestPaymentDate")
		Attribute("LastPaymentAmount")
	})

	View("mandatory", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("TradingAs")
		Attribute("DateOfBirth")
		Attribute("AccountNumber")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocumentNumber")
		Attribute("MobilePhoneNumber")
		Attribute("AccountProductType")
		Attribute("InstalmentDueDate")
		Attribute("OriginalAmount")
		Attribute("CurrencyOfFacility")
		Attribute("CurrentBalanceInKES")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("OverdueDate")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("PrudentialRiskClassification")
		Attribute("AccountStatus")
		Attribute("AccountStatusDate")
		Attribute("RepaymentPeriod")
		Attribute("DisbursementDate")
		Attribute("InstalmentAmount")
		Attribute("LatestPaymentDate")
		Attribute("LastPaymentAmount")
	})
})
