package design

import (
	. "goa.design/goa/v3/dsl"
)

// Individual Consumer, Employer and Account API
var ICEAccount = Type("icea", func() {
	Description("Individual Consumer, Employer and Account API")
	TypeName("ICEAccount")
	ContentType("application/json")

	// The Family Name or Surname.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "1")
	})

	// The First Name.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory Field
	// Not more than 50
	// Allow Hyphen, Apostrophe.
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50
	// Allow Hyphen, Apostrophe
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50
	// Allow Hyphen, Apostrophe.
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "4")
	})

	// Name of a business if the individual runs a business as a sole proprietor
	// that is not registered; or if the sole proprietor took a credit
	// facility to finance the business.
	//
	// Alphanumeric
	// Not more than 50
	Attribute("TradingAs", String, "Trading As", func() {
		Meta("rpc:tag", "5")
	})

	// Assist in matching for unique identification and detection of possible fraud.
	// The field needs to be mandatory if the Id No cannot be validated or changes
	// in case of Passport renewal & Passport number changes
	//
	// Mandatory - Issue Warning
	// Not a future date
	// Not under 18
	// Not over 100 years
	Attribute("DoB", String, "Date of Birth", func() {
		Meta("rpc:tag", "6")
	})

	// Client reference Number in Core Lenders system.
	// Client Number if Lender is Client Centric Account Number if
	// Lender is not Client Centric.
	//
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number Linking Customer to Core Lenders system Same as
	// Client Number For Account-Centric System.
	//
	// Mandatory field
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "8")
	})

	// Contract Unique Identifier used previously to report this facility.
	//
	// Alphanumeric Field
	// Not more than 20 characters
	Attribute("OldAccountNumber", String, "Old Account Number", func() {
		Meta("rpc:tag", "9")
	})

	// Male/Female ( M/F)
	//
	// Mandatory
	// Lookup
	// Alphanumeric
	Attribute("Gender", String, "Gender", func() {
		Meta("rpc:tag", "10")
	})

	// ISO Country Code for the Consumer’s Nationality.
	//
	// Mandatory Field – Issue Warning
	// Data is Alphanumeric
	// Not more than 2 characters
	// Matches ISO CODE
	Attribute("Nationality", String, "Nationality", func() {
		Meta("rpc:tag", "11")
	})

	// Consumer’s Marital status
	// Options:
	// M – Married
	// S - Single
	// D - Divorced
	// W - Widowed
	//
	// Lookup
	// Not more than 1 character
	// Alphanumeric
	Attribute("MaritalStatus", String, "Marital Status", func() {
		Meta("rpc:tag", "12")
	})

	// The Type of Primary Identification document Provided on Opening the Account.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not more than 3 characters
	// Lookup
	Attribute("PrimaryIDocumentType", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Identification Document Provided on Opening the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "14")
	})

	// If an additional Identification Document is available, then the Type
	// of Additional Identification Document provided.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Lookup
	// Data is Alphanumeric
	// Not more than 3 characters
	Attribute("SecondaryIDocumentType", String, "Secondary Identification Document Type", func() {
		Meta("rpc:tag", "15")
	})

	// The Number of the Secondary Identification Document Provided.
	// Mandatory if Secondary
	// Identification Document Type is Provided.
	// Conditional field based on field 4.2.15
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("SecondaryIDocumentNumber", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "16")
	})

	// If an additional Identification Document is available, then the Type
	// of Additional Identification Document
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	//
	// Data is Alpha Numeric
	// Not more than 3 characters
	// Lookup
	Attribute("OtherIDocumentType", String, "Other Identification Document Type", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of the Other Identification Document Provided.
	// Mandatory if Other Identification Document Type is Provided.
	// Conditional field based on field 4.2.17
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("OtherIDocumentNumber", String, "Other Identification Document Number", func() {
		Meta("rpc:tag", "18")
	})

	// The Country Code where the Passport was Registered.
	//
	// Alphanumeric
	// Not more than 2 characters
	Attribute("PassportCountryCode", String, "Passport Country Code", func() {
		Meta("rpc:tag", "19")
	})

	// The Consumers Main Telephone contact Number in the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("MobileTelephoneNumber", String, "Mobile Telephone Number", func() {
		Meta("rpc:tag", "20")
	})

	// The Consumer’s Secondary Telephone contact Number in the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	Attribute("HomeTelephoneNumber", String, "Home Telephone Number", func() {
		Meta("rpc:tag", "21")
	})

	// The Consumer’s Office Telephone contact Number, if consumer
	// is employed in the Form of : CCCAAANNNNNNN
	//
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("WorkTelephoneNumber", String, "Work Telephone Number", func() {
		Meta("rpc:tag", "22")
	})

	// Consumer’s Postal Address Line 1
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress1", String, "Postal Address 1", func() {
		Meta("rpc:tag", "23")
	})

	// Address Extra Details
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PostalAddress2", String, "Postal Address 2", func() {
		Meta("rpc:tag", "24")
	})

	// Consumer’s Postal Address Town
	//
	// Alphanumeric
	// Not more than 30 characters
	Attribute("PostalLocationTown", String, "Postal Location Town", func() {
		Meta("rpc:tag", "25")
	})

	// Country of Consumer’s Postal Address
	//
	// Alphanumeric
	// Not more than 2 characters
	Attribute("PostalLocationCountry", String, "Postal Location Country", func() {
		Meta("rpc:tag", "26")
	})

	// Post Code of Address
	//
	// Alphanumeric
	// Not more than 10 characters
	Attribute("PostCode", String, "Post code", func() {
		Meta("rpc:tag", "27")
	})

	// Consumer’s Physical ( Residential Address) Line 1
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress1", String, "Physical Address 1", func() {
		Meta("rpc:tag", "28")
	})

	// Consumer’s Physical (residential Address) Line 2
	//
	// Alphanumeric
	// not more than 50 characters
	Attribute("PhysicalAddress2", String, "Physical Address 2", func() {
		Meta("rpc:tag", "29")
	})

	// Plot Land Ref (LR) No of Consumer’s residential Address
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("PlotNumber", String, "Plot Number", func() {
		Meta("rpc:tag", "30")
	})

	// Town of Consumer’s residential Address
	//
	// Alphanumeric
	// Not more than 30 characters
	Attribute("LocationTown", String, "Location Town", func() {
		Meta("rpc:tag", "31")
	})

	// ISO Code of the Country of the Consumer’s residential Address.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 2 characters
	Attribute("LocationCountry", String, "Location Country", func() {
		Meta("rpc:tag", "32")
	})

	// The type of housing the borrower lives in
	// A – Owned
	// B – Rented
	// C – Mortgaged
	// D – Family Owned
	//
	// Lookup
	// Alphanumeric Field
	Attribute("ResidencyType", String, "Type of Residency", func() {
		Meta("rpc:tag", "33")
	})

	// Income Tax PIN No.
	//
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, "PIN Number", func() {
		Meta("rpc:tag", "34")
	})

	// The Consumer Work Email Address, If employed or available.
	//
	// Validate Email Syntax
	// Not more than 50 characters
	// Alphanumeric
	Attribute("ConsumerEmail", String, "Consumer E-Mail", func() {
		Meta("rpc:tag", "35")
	})

	// If the consumer is employed, the Employer Name.
	// The field is conditional based on the Employment Type Field.
	// It is not mandatory for Self-Employed and Pensioners.
	//
	// Alphanumeric Field
	// Not more than 50 characters
	Attribute("EmployerName", String, "Employer Name", func() {
		Meta("rpc:tag", "36")
	})

	// The Occupational Industry Type of the Borrower
	// Options are :
	// 001 -Agriculture
	// 002 -Manufacturing
	// 003 -Building/Construction
	// 004 - Mining/Quarrying
	// 005 - Energy /water
	// 006 -Trade
	// 007 -Tourism/ Restaurant/Hotels
	// 008 -Transport/ Communications
	// 009 -Real Estate
	// 010 - Finance
	// 011 - Government
	//
	// Mandatory Field
	// Alphanumeric Field
	// 3 characters only
	Attribute("OccupationalIndustryType", String, "Occupational Industry Type", func() {
		Meta("rpc:tag", "37")
	})

	// If the Consumer is employed, the Date Consumer Was employed.
	//
	// Date Field
	// Should not be in the future
	// Not more than 8 characters
	Attribute("EmploymentDate", String, "Employment Date", func() {
		Meta("rpc:tag", "38")
	})

	// Type Of employment – Options
	// Available :
	// 001=Casual
	// 002= Contract
	// 003= Permanent
	// 004=Pensioner
	// 005= Self-Employed
	//
	// Lookup
	// Alphanumeric field
	// 3 characters only
	Attribute("EmploymentType", String, "Employment type", func() {
		Meta("rpc:tag", "39")
	})

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	// The Name of Lender as registered with the Registrar of companies.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersRegisteredName", String, "Lenders Registered Name", func() {
		Meta("rpc:tag", "41")
	})

	// The Lenders Trading Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersTradingName", String, "Lenders Trading Name", func() {
		Meta("rpc:tag", "42")
	})

	// The Lenders Branch Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersBranchName", String, "Lenders Branch Name", func() {
		Meta("rpc:tag", "43")
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
	// IXXX MUST match the API name
	// XBB (code of the institution making the submission)
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "44")
	})

	// If the Account is operated jointly or by a single individual.
	// Options available:
	// J = Joint
	// S = Single
	//
	// Mandatory Field
	// Lookup
	// Not more than 1 character
	Attribute("AJSIndicator", String, "Account Joint/Single Indicator", func() {
		Meta("rpc:tag", "45")
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
	// Mandatory Field
	// Lookup
	// Not more than 1 character
	Attribute("AccountProductType", String, "Account Product Type", func() {
		Meta("rpc:tag", "46")
	})

	// Date the Account was opened.
	//
	// Mandatory Field
	// Date Field
	// Cannot be in the future
	// Not more than 8 characters
	Attribute("DateAccountOpened", String, "Date Account Opened", func() {
		Meta("rpc:tag", "47")
	})

	// If the account is not in arrears, give the next instalment date.
	// If the account is in arrears, give the overdue date.
	// For Overdrafts, Credit Cards and Trade Finance products, give the expiry date.
	//
	// Mandatory Field
	// Date Field
	// Not more than 8 characters
	Attribute("InstalmentDueDate", String, "Instalment Due Date", func() {
		Meta("rpc:tag", "48")
	})

	// For Loans, give the approved loan amount.
	// For Overdrafts, Credit cards and Trade Finance Products, issue the approved limits.
	//
	// Amount Field
	// Not more than 16 characters
	// Mandatory Field
	Attribute("OriginalAmount", String, "Original Amount", func() {
		Meta("rpc:tag", "49")
	})

	// The ISO Currency Code for the Currency in which the facility is given.
	// Default is KES for Kenya Shillings.
	// Alphanumeric
	// Not more than 3 Characters
	// Mandatory Field
	Attribute("FacilityCurrency", String, "Currency of Facility", func() {
		Meta("rpc:tag", "50")
	})

	// The current balance in Kenya shillings, if the Currency is not Kenya shillings.
	// Same as the Original Amount if the facility was given in Kenya Shillings.
	// Use the exchange rate as at the reporting (month end) date.
	//
	// Amount Field
	// Not more than 16 characters
	// Mandatory Field
	Attribute("CurrentBalanceInKES", String, "Current Balance in Kenya Shillings", func() {
		Meta("rpc:tag", "51")
	})

	// Balance in the account as at reporting (end month) date
	// For Overdrafts:
	// If the overdraft has been utilized, give the debit (negative) balance due.
	// If the overdraft has not been utilized, or if a previously submitted
	// overdraft has been paid up, give a Zero (0)
	//
	// Amount Field
	// Not more than 16 characters
	// Mandatory Field
	Attribute("Current Balance", String, "Current Balance", func() {
		Meta("rpc:tag", "52")
	})

	// Arrears amount in a facility.
	// Mandatory if Overdue date is not null,
	// Number of Days in Arrears
	// & Number of Instalments in Arrears > 0
	Attribute("Overdue Balance", String, "Overdue Balance", func() {
		Meta("rpc:tag", "53")
	})

	// If the Account is in arrears, the Date when the Account fell overdue.
	// For Overdrafts, Trade Finance Products and Credit Cards, the
	// date of expiry or date it went over the limit Conditional to Overdue Balance,
	// Number of Days In Arrears and Number of Instalments In Arrears.
	//
	// Date Field
	// Not more than 8 characters
	// Date can’t be in the future
	Attribute("Overdue Date", String, "Overdue Date", func() {
		Meta("rpc:tag", "54")
	})

	// 000 - If the account is not in arrears to accommodate performing loans
	// Count the number of days account is in arrears (e.g. 120 days in arrears)
	//
	// Numeric Field
	// Not more than 5 characters
	// Mandatory field
	Attribute("DaysInArrears", String, "Number of Days In Arrears", func() {
		Meta("rpc:tag", "55")
	})

	// Number of missed Instalments in the Facility.
	// Required if the Account is in arrears.
	// Conditional to Number of Days In Arrears, Overdue Date, Overdue Balance
	//
	// Numeric Field
	// Not more than 3 characters
	// Mandatory field
	Attribute("InstalmentsInArrears", String, "Number of Instalments In Arrears", func() {
		Meta("rpc:tag", "56")
	})

	// This is as per CBK guidelines on classification of assets
	// Options available :
	// A = Normal (0-29)
	// B = Watch(31-89)
	// C = Substandard (91-179)
	// D = Doubtful( 181-360)
	// E = Loss (>360)
	//
	// Lookup
	// Alphanumeric Field
	// Not more than 1 character
	// Mandatory Field
	Attribute("PrudentialRiskClassification", String, "Prudential Risk Classification", func() {
		Meta("rpc:tag", "57")
	})

	// The Account status as at the time of reporting.
	// Options Available:
	// A - Closed – No more admin processes running such as instalment
	// demands or interest charges to account, and no further facilities
	// can be offered on this account.
	// B - Dormant - no activity for 2 years. This applies for Overdraft/Current Accounts.
	// C - Write-Off – For facilities that don’t form part of the outstanding
	// portfolio in the Balance Sheet, but are still outstanding in the books of accounts.
	// D - Legal -with legal officer in court
	// E - Collection- with collection bureau
	// F – Active - For facilities that form part of the outstanding portfolio, and are
	// reported in the Balance Sheet.
	// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
	// H –Settled – The facility has been cleared.
	// This status can only be used for revolving facilities
	// J – Called Up : The facility has been called up.
	// Once the client has paid up, the status should be updated to Option A, H;
	// or otherwise Option L
	// K – Suspended – The facility has been put on hold for an indefinite period of time
	// L– Client Deceased
	// M – Deferred – This refers to facilities whose payments have been put on hold for
	// a definite period or in moratorium (Grace Period)
	// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
	// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved.
	//
	// Lookup
	// Mandatory Field
	// Not more than 1 character
	Attribute("AccountStatus", String, "Account Status", func() {
		Meta("rpc:tag", "58")
	})

	// The date of the status Change.
	//
	// Date Field
	// Mandatory Field
	// Date Cannot be in the future
	Attribute("AccountStatusDate", String, "Account Status Date", func() {
		Meta("rpc:tag", "59")
	})

	// If the status of the Account is Closed, The reason for the account closure.
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("AccountClosureReason", String, "Account Closure Reason", func() {
		Meta("rpc:tag", "60")
	})

	// The repayment Period, in months for the Facility.
	// 999 for Credit Cards
	//
	// Mandatory field
	// Number Field
	// Not more than 3 characters
	Attribute("RepaymentPeriod", String, "Repayment Period", func() {
		Meta("rpc:tag", "61")
	})

	// Mandatory if Account Status = Deferred.
	// Deferred Payment Date must be in the future
	//
	// Date Field
	// Not more than 8 Characters
	// Must Be in Future
	Attribute("DeferredPaymentDate", String, "Deferred Payment Date", func() {
		Meta("rpc:tag", "62")
	})

	// Conditional to Account Status and Deferred Payment Date.
	//
	// Amount Field
	// Not more than 16 characters
	Attribute("DeferredPaymentAmount", String, "Deferred Payment Amount", func() {
		Meta("rpc:tag", "63")
	})

	// The Frequency in which Instalments are to be paid on the facility.
	// Options available:
	// W - Weekly
	// M - Monthly
	// Q – Quarterly
	// S – Semi-Annually
	// Y – Annually ( Yearly)
	// B – Bullet
	// T – Tri-annual (Three times in a year)
	// C – Biennial (Once in two years)
	// D – Daily
	// F - Fortnightly
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	Attribute("PaymentFrequency", String, "Payment Frequency", func() {
		Meta("rpc:tag", "64")
	})

	// Date of facility drawdown.
	//
	// Mandatory Field
	// Not more than 8 characters
	// Date Field
	// Cannot be in the future
	Attribute("DisbursementDate", String, "Disbursement Date", func() {
		Meta("rpc:tag", "65")
	})

	// Instalment amount of loans.
	// Optional for Overdrafts, Credit Cards and Trade Finance facilities.
	//
	// Amount Field
	// Not more than 16 characters
	Attribute("NextInstalmentAmount", String, "Next Instalment Amount", func() {
		Meta("rpc:tag", "66")
	})

	// Last date when payments were received into the facility.
	// Not mandatory for Trade Finance Products.
	// Null if no payment has been received.
	//
	// Date Field
	// Not more than 8 characters
	// Date cannot be in the future
	Attribute("LatestPaymentDate", String, "Date of Latest Payment", func() {
		Meta("rpc:tag", "67")
	})

	// Last payment received into the facility.
	// Not mandatory for Trade Finance Products.
	// Mandatory if Date of Latest Payment is not null.
	// Null if no payment has been received.
	//
	// Amount Field
	// Not more than 16 characters
	Attribute("LastPaymentAmount", String, "Last Payment Amount", func() {
		Meta("rpc:tag", "68")
	})

	// Indicator of whether the facilities are secured or Not.
	// Options Available :
	// U - Unsecured
	// S - Fully Secured
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	// Mandatory
	Attribute("SecurityType", String, "Type of Security", func() {
		Meta("rpc:tag", "69")
	})

	// The Unique Identifier of the Group the borrower belongs to under
	// which the facility was granted.
	//
	// Alphanumeric
	// Not more than 15 Characters
	Attribute("GroupID", String, "Group ID", func() {
		Meta("rpc:tag", "70")
	})

})
