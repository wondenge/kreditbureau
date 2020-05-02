package design

import (
	. "goa.design/goa/v3/dsl"
)

var MobileFacility = Type("mobilefacility", func() {
	Description("Mobile Facilities API")
	TypeName("MobileFacility")
	ContentType("application/json")

	// The Family Name or Surname
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "1")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "4")
	})

	// Name of a business if the individual runs a business as a
	// sole proprietor that is not registered; or if the sole proprietor
	// took a credit facility to finance the business.
	//
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("TradingAs", String, "Trading As", func() {
		Meta("rpc:tag", "5")
	})

	// The Date of Birth of the Customer.
	//
	// Mandatory - Issue Warning
	// Not a future date
	// Not under 18
	// Not over 100 years
	Attribute("DateOfBirth", String, "Date of Birth", func() {
		Meta("rpc:tag", "6")
	})

	// Unique identification of a client within the lenders Core System.
	//
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "7")
	})

	// Number that uniquely identifies the specific Mobile Facility.
	//
	// Mandatory field
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "8")
	})

	// Gender of the Borrower.
	// 	M = "Male"
	//	F = "Female"
	//
	// Mandatory with Warning
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	Attribute("Gender", String, "Gender", func() {
		Meta("rpc:tag", "9")
	})

	// ISO Country Code for the Consumer’s Nationality.
	//
	// Mandatory Field – Issue Warning
	// Data is Alphanumeric
	// Not more than 2 characters
	// Matches ISO CODE
	Attribute("Nationality", String, "Nationality", func() {
		Meta("rpc:tag", "10")
	})

	// 	001 = NationalID IDType
	//	002 = Passport
	//	003 = AlienRegistration
	//	004 = ServiceID
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not more than 3 characters
	Attribute("PrimaryIDocumentType", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "11")
	})

	// The Number of the Primary Identification Document provided on
	// Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocumentNumber", String, "Primary Identification Document Number", func() {
		Meta("rpc:tag", "12")
	})

	// The Mobile Phone Number of the Client used to grant the loan facility.
	// Alphanumeric
	// Mandatory Field
	// Not more than 15 characters
	Attribute("MobilePhoneNumber", String, "Mobile Phone Number", func() {
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
	//
	// Mandatory Field
	// Lookup
	// Not more than 1 character
	Attribute("AccountProductType", String, "Account Product Type", func() {
		Meta("rpc:tag", "14")
	})

	// If the account is not in arrears, report the next instalment date.
	// If the account is in arrears, report the overdue date.
	//
	// Mandatory Field
	// Date Field
	// Not more than 8 characters
	Attribute("InstalmentDueDate", String, "Instalment Due Date", func() {
		Meta("rpc:tag", "15")
	})

	// The principal Amount issued
	//
	// Currency Field
	// Not more than 16 characters
	// Mandatory Field
	Attribute("OriginalAmount", String, "Original Amount", func() {
		Meta("rpc:tag", "16")
	})

	// The currency code facility's currency.
	//
	// Alphanumeric
	// Not more than 3 Characters
	// Mandatory Field
	Attribute("CurrencyOfFacility", String, "Currency of Facility", func() {
		Meta("rpc:tag", "17")
	})

	// The current balance in Kenya shillings, if the Currency is not Kenya shillings.
	//
	// Currency Field
	// Not more than 16 characters
	// Mandatory Field
	Attribute("CurrentBalanceInKES", String, "Current Balance in Kenya Shillings", func() {
		Meta("rpc:tag", "18")
	})

	// The current balance in the mobile account in the original currency.
	//
	// Mandatory Field
	// Currency Field
	// Not more than 16 characters
	Attribute("CurrentBalance", String, "Current Balance", func() {
		Meta("rpc:tag", "19")
	})

	// Overdue Amount in a facility.
	// If account is not in arrears report Zero (0)
	//
	// Currency Field
	// Not more than 16 characters
	// Mandatory field
	Attribute("OverdueBalance", String, "Overdue Balance", func() {
		Meta("rpc:tag", "20")
	})

	// The date when the account fell Overdue
	// This Field is Mandatory if Field 5.2.20 is greater than Zero (0)
	//
	// Date Field
	// Not more than 8 characters
	// Date can’t be in the future
	Attribute("OverdueDate", String, "Overdue Date", func() {
		Meta("rpc:tag", "21")
	})

	// The number of days the account has been in arrears (Difference between
	// reporting date and overdue date, calculated in days)
	// If the account is not in arrears, report Zero (0).
	//
	// Numeric Field
	// Not more than 5 characters
	// Mandatory field
	Attribute("DaysInArrears", String, "Number of Days in Arrears", func() {
		Meta("rpc:tag", "22")
	})

	// The Number of missed Instalments in the Facility.
	// If Field 5.2.22 is Zero (0) return 0.
	// This field cannot be 0 if Field 5.2.22 is greater than 0.
	//
	// Mandatory Field
	// Number Field
	// Cannot be 0 if field 5.2.22 is greater than 0
	// Not more than 5 characters
	Attribute("InstalmentsInArrears", String, "Number of Instalments in Arrears", func() {
		Meta("rpc:tag", "23")
	})

	// Classification of the Assets.
	// 	A  = "Normal (0-30)"
	//	B  = "Watch (>30-90)"
	//	C  = "Substandard(>90-180)"
	//	D  = "Doubtful(>180-360)"
	//	E  = "Loss(>360)"
	//
	// Alphanumeric Field
	// Not more than 1 character
	// Lookup
	// Mandatory Field
	Attribute("PrudentialRiskClassification", String, "Prudential Risk Classification", func() {
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
	//
	// Lookup
	// Not more than 1 character
	// Mandatory Field
	Attribute("AccountStatus", String, "Account Status", func() {
		Meta("rpc:tag", "25")
	})

	// The date of the status.
	//
	// Mandatory field
	// Date Field
	// Not more than 8 fields
	Attribute("AccountStatusDate", String, "Account Status Date", func() {
		Meta("rpc:tag", "26")
	})

	// Repayment Period for the Facility in months.
	// This is the initial contractual period.
	//
	// Number Field
	// Mandatory Field
	// Not more than 3 Characters
	Attribute("RepaymentPeriod", String, "Repayment Period", func() {
		Meta("rpc:tag", "27")
	})

	// Date of facility drawdown.
	//
	// Date Field
	// Cannot be in the future
	// Not more than 8 Character
	Attribute("DisbursementDate", String, "Disbursement Date", func() {
		Meta("rpc:tag", "28")
	})

	// Instalment amount for the loan.
	// Amount Field
	// Mandatory Field
	// Not more than 16 Characters
	Attribute("InstalmentAmount", String, "Instalment Amount", func() {
		Meta("rpc:tag", "29")
	})

	// The date when payments were last received into the facility.
	// This field should be null if Option “G” is selected in Field 5.2.14.
	//
	// Date Field
	// Date not in future
	// Not more than 8 characters
	// Conditional
	Attribute("LatestPaymentDate", String, "Date of Latest Payment", func() {
		Meta("rpc:tag", "30")
	})

	// Last Payment Amount received into the facility Report Zero (0) if no payment was received.
	//
	// Amount Field
	// Not more than 16 Characters
	// Mandatory Field
	Attribute("LastPaymentAmount", String, "Last Payment Amount", func() {
		Meta("rpc:tag", "31")
	})
})

var StoredMobileFacility = ResultType("application/vnd.goa.mobilefacility", func() {
	TypeName("StoredMobileFacility")
	Attributes(func() {
		Extend(MobileFacility)
		Required()
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
})
