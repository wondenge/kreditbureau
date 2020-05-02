package design

import (
	. "goa.design/goa/v3/dsl"
)

// Non Individual Consumer and Account API
var NonICEAccount = Type("company", func() {
	Description("Non Individual Consumer and Account API")
	TypeName("NonICEAccount")
	ContentType("application/json")

	// The Name as Registered with the Registrar of Companies.
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not More than 150 Characters
	Attribute("RegisteredName", String, "Registered Name", func() {
		Meta("rpc:tag", "1")
	})

	// The Business or Trading Name.
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not More than 150 Characters
	Attribute("TradingName", String, "Trading Name", func() {
		Meta("rpc:tag", "2")
	})

	// Date Registered with the Registrar of Companies.
	//
	// Mandatory Field – Issue warning
	// Date Format
	// YYYYMMDD – Issue Warning
	// Date Not in the Future
	// Not more than 8 Numeric characters
	Attribute("RegistrationDate", String, "Registration Date", func() {
		Meta("rpc:tag", "3")
	})

	// The Registration Certificate Number.
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not more than 25 characters
	Attribute("RegistrationNumber", String, "Registration Number", func() {
		Meta("rpc:tag", "4")
	})

	// The Registration Certificate Number - format to conform to Registrar of Companies in Kenya.
	//
	// Required, only if applicable because the company type changes, the history
	// details of the previous type can be matched to the new company type.
	// Data is Alphanumeric
	// Not more than 25 characters
	Attribute("PreviousRegistrationNumber", String, "Previous RegistrationNumber", func() {
		Meta("rpc:tag", "5")
	})

	// The Country of Registration of the Company, Defaults to Kenyan.
	//
	// Mandatory Field – Issue Warning
	// Data is Alphanumeric
	// Not more than 2 characters
	// Matches COUNTRY ISO CODE
	Attribute("Nationality", String, "Nationality", func() {
		Meta("rpc:tag", "6")
	})

	// The Client Reference Number linking Company to Core Banking system.
	// Client reference Number in Core Lenders system.
	// Client Number if Lender is Client Centric.
	//
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "7")
	})

	// The Account Number Linking Client to Bank’s Accounting System.
	// Account Number if Lender is not Client Centric.
	//
	// Mandatory field
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "8")
	})

	// Old Account Number linking to the previous MIS System
	//
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("OldAccountNumber", String, "Old Account Number", func() {
		Meta("rpc:tag", "9")
	})

	// The Type of Organisation. Options Available :
	// A – Public Limited company
	// B - Private Limited company
	// C - Partnership
	// D - Sole Proprietor
	// E – Society
	// F - Club/Association
	// G – Investment Group
	// H - Public Corporation
	// J - NGO
	// L - Church
	// M - School
	//
	// Mandatory – Issue Warning
	// Lookup Field – Company Type
	// Not more than 1 Character
	Attribute("CompanyType", String, "Company Type", func() {
		Meta("rpc:tag", "10")
	})

	// Industry Code for the Borrower’s Line of business.
	// Use the highest level of Central Bank of Kenya Codes.
	// Options are :
	// 001 - Agriculture
	// 002 - Manufacturing
	// 003 - Building/Construction
	// 004 - Mining/Quarrying
	// 005 - Energy /water
	// 006 -Trade
	// 007 - Tourism/ Restaurant/Hotels
	// 008 - Transport/ Communications
	// 009 - Real Estate
	// 010 - Financial Services
	// 011 – Government
	// 012 - Other
	//
	// Mandatory
	// Lookup Field – Industry Code
	// 3 characters only
	Attribute("IndustryCode", String, "Industry Code", func() {
		Meta("rpc:tag", "11")
	})

	// The Borrower’s Annual Gross Turnover Range, in Kenya Shillings.
	// Use the Central Bank of Kenya range options:
	//
	// Mandatory – Issue Warning
	// Data is Amount
	// Not Negative
	Attribute("AnnualTurnoverAmount", String, "Annual Turnover Amount", func() {
		Meta("rpc:tag", "12")
	})

	// The Income Tax PIN Number.
	//
	// Mandatory Field with Warning
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINNumber", String, "Income Tax PIN No.", func() {
		Meta("rpc:tag", "13")
	})

	// Value Added Tax Registration Number.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("VATNumber", String, "Value Added Tax Registration Number", func() {
		Meta("rpc:tag", "14")
	})

	// Indicate the number of persons who hold stake/shares in a non-individual entity.
	// 1 for Sole Proprietor.
	//
	// Mandatory – Issue Warning
	// Greater than or Equal to 1
	// Validate against Stakeholder API
	Attribute("NofStakeholders", String, "Number of Stakeholders", func() {
		Meta("rpc:tag", "15")
	})

	// Trading Status
	// Status of Trading of the consumer
	// Options :
	// 001=Dormant
	// 002=Actively Trading
	// 003=Under Receivership/Liquidation
	// 004=Liquidated
	// 005=Under Management
	// 006=Dissolved
	//
	// Mandatory with Warning
	// Alphanumeric
	// 3 characters only
	Attribute("TradingStatus", String, "Trading Status", func() {
		Meta("rpc:tag", "16")
	})

	// Trading Status Date.
	// Date of the Current Trading Status.
	//
	// Date format.
	// Not a future date.
	// Not a date before the company registration date.
	Attribute("TradingStatusDate", String, "Trading Status Date", func() {
		Meta("rpc:tag", "17")
	})

	// The Primary Telephone contact Number in the Format of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Mandatory-Issue warning
	// Alphanumeric
	// Not more than 15 characters
	// Check for Dummy Data
	// Check the phone format
	Attribute("MainTelephoneNumber", String, "Main Telephone Number", func() {
		Meta("rpc:tag", "18")
	})

	// Any other Telephone contact Number for the Consumer in the Format of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	//
	// Alphanumeric
	// Not more than 15 characters
	// Check for Dummy Data
	// Check the phone format
	Attribute("OtherTelephoneNumber", String, "Other Telephone Number", func() {
		Meta("rpc:tag", "19")
	})

	// Consumer Postal Address Line 1.
	// First Line of Company’s Full Postal Address.
	//
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("PostalAddress1", String, "Consumer’s Postal Address Line 1", func() {
		Meta("rpc:tag", "20")
	})

	// Consumer Postal Address Line 2.
	// Second Line of Company’s Full Postal Address.
	//
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("PostalAddress2", String, "Postal Address Extra Details.", func() {
		Meta("rpc:tag", "21")
	})

	// Town of consumer Postal Address.
	//
	// Alphanumeric
	// Not more than 30 Characters
	Attribute("Town", String, "Postal Town of Postal Address", func() {
		Meta("rpc:tag", "22")
	})

	// Country of consumer Postal Address.
	//
	// Alphanumeric
	// ISO CODE: Country
	// Not more than 2 characters
	Attribute("Country", String, "Country of Operation for the Consumer", func() {
		Meta("rpc:tag", "23")
	})

	// Post Code of consumer Address.
	//
	// Alphanumeric
	// Not more than 10 characters
	Attribute("PostCode", String, "Post Code of Consumer’s Postal Address", func() {
		Meta("rpc:tag", "24")
	})

	// Physical Address Line 1 of the Consumer’s business Location ( Office Location)
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Physical Address Line 2 of the Consumer’s business Location ( Office Location)
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("PhysicalAddress2", String, "Physical Address 2", func() {
		Meta("rpc:tag", "26")
	})

	// Plot Land Ref (LR) No of the Consumer’s business Location ( Office Location)
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("PlotNumber", String, "Plot Number", func() {
		Meta("rpc:tag", "27")
	})

	// Town of the Consumer’s business Location ( Office Location)
	//
	// Mandatory field
	// Alphanumeric
	// Not more than 30 characters
	Attribute("PhysicalLocationTown", String, "Physical Location Town", func() {
		Meta("rpc:tag", "28")
	})

	// Country of the Office Location.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 2 characters
	Attribute("PhysicalLocationCountry", String, "Physical Location Country", func() {
		Meta("rpc:tag", "29")
	})

	// The Name of Lender as registered with the Registrar of companies.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersRegisteredName", String, "Lenders Registered Name", func() {
		Meta("rpc:tag", "30")
	})

	// The Lenders Trading Name.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersTradingName", String, "Lenders Trading Name", func() {
		Meta("rpc:tag", "31")
	})

	// The Lenders Branch Name.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersBranchName", String, "Lenders Branch Name", func() {
		Meta("rpc:tag", "32")
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
	// Mandatory
	// Alphanumeric
	// Not more than 7 characters
	// IXXX MUST match the API name XBB (code of the institution making the submission)
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "33")
	})

	// Whether the Account is operated Jointly or By a single individual.
	// Options available :
	// J = Joint
	// S = Single
	//
	// Data is Alphanumeric
	// Not more than 1 character
	// Lookup-Account Joint/Single
	// Indicator
	Attribute("AccountIndicator", String, "Account Joint/Single Indicator", func() {
		Meta("rpc:tag", "34")
	})

	// Options Available :
	// A – Corporate
	// B – SME
	// C - NGO/Church/Society
	//
	// Alphanumeric
	// Not more than 1 character
	// Lookup-Class of Organisation
	Attribute("OrganisationClass", String, "Class of Organisation", func() {
		Meta("rpc:tag", "35")
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
	// Mandatory with Warning
	// Alphanumeric
	// Lookup- Account Product Type
	// Not more than 1 character
	// Option B,H, I, K not allowed
	Attribute("AccountProductType", String, "Account Product Type", func() {
		Meta("rpc:tag", "36")
	})

	// Date When the Account was opened.
	//
	// Mandatory
	// Date Field
	// Not in the Future
	// - Date should not be more than the snapshot date
	Attribute("DateAccountOpened", String, "Date Account Opened", func() {
		Meta("rpc:tag", "37")
	})

	// If the account is not in arrears, give the next instalment date.
	// If the account is in arrears, give the overdue date
	// For Overdrafts, Credit Cards and Trade Finance products, give the expiry date.
	//
	// Mandatory
	// Date Field Format Check conditionals
	// - Not date in the past for an account not in arrears.
	// - Not date in the future for an account in arrears.
	// - Future Date for Overdrafts, credit cards, trade finance.
	Attribute("InstalmentDueDate", String, "Instalment Due Date", func() {
		Meta("rpc:tag", "38")
	})

	// The guidelines as per CBK on loan classification
	// Options available :
	// A = Normal (0-30)
	// B = Watch (>30-90)
	// C = Substandard (>90-180)
	// D = Doubtful (>180-360)
	// E = Loss (>360)
	//
	// Mandatory with Warning.
	// Lookup Check.
	// Not more than 1 character.
	Attribute("PrudentialRiskClassification", String, "Prudential Risk Classification", func() {
		Meta("rpc:tag", "39")
	})

	// For Loans, give the approved loan amount.
	// For Overdrafts, Credit cards and Trade Finance Products, issue the approved limits.
	//
	// Mandatory Field
	// Amount Field
	// Must not be negative
	// Must be Greater than Zero
	// Not more than 16 characters
	Attribute("OriginalAmount", String, "Original Amount", func() {
		Meta("rpc:tag", "40")
	})

	// The ISO Currency Code for the Currency in which the facility is given.
	// Default is KES for Kenya Shillings.
	// Refer to Appendix B for valid Currency Codes
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 3 characters
	// ISO currency codes
	// Lookup
	Attribute("FacilityCurrency", String, "Currency of Facility", func() {
		Meta("rpc:tag", "41")
	})

	// The Current Balance in Kenya shillings, if the
	// Currency is not Kenya shillings.
	// Same as the Original Amount if the facility was given in Kenya Shillings.
	// Use the exchange rate as at the reporting (end month) date.
	//
	// Mandatory Field
	// The Value should be same as 4.1.43 if field 4.1.41 is in KES
	// Amount Field
	// Not more than 16 characters
	Attribute("KESCurrentBalance", String, "Current Balance in Kenya Shillings", func() {
		Meta("rpc:tag", "42")
	})

	// Balance in the account as at reporting (end month) date
	// For Overdrafts:
	// If the overdraft has been utilized, give the debit (negative) balance due.
	// If the overdraft has not been utilized, or if a previously submitted
	// overdraft has been paid up, give a Zero (0)
	//
	// Mandatory Field
	// Amount Field
	// Must Not be negative
	// Cannot be zero if Overdue Balance > 0
	// Cannot be Zero if DPD > 0
	// Cannot be > 0 if acct status is
	// Closed, Paid Up, Fully Settled, Early Settlement
	// Not more than 16 characters
	Attribute("CurrentBalance", String, "Current Balance", func() {
		Meta("rpc:tag", "43")
	})

	// Arrears amount in a facility.
	// Mandatory if Overdue date is not null, Nr of Days in Arrears
	// & Nr of Instalments in Arrears > 0
	//
	// Mandatory Field
	// Amount Field
	// Cannot be negative
	// Not more than 16 characters
	// Must be filled if 4.1.45
	Attribute("OverdueBalance", String, "Overdue Balance", func() {
		Meta("rpc:tag", "44")
	})

	// If the Account is in arrears, the Date when the Account fell overdue.
	// For Overdrafts, Trade Finance Products and Credit Cards, the date of
	// expiry or date it went over the limit Conditional to Overdue Balance,
	// Nr of Days In Arrears and Number of Instalments In Arrears.
	//
	// Date format
	// Conditional Check based on 4.1.44, 4.1.46 and 4.1.47
	// Cannot be blank if overdue balance
	// Is greater than zero
	// Not in the future
	// Not greater than snapshot date
	Attribute("OverdueDate", String, "Overdue Date", func() {
		Meta("rpc:tag", "45")
	})

	// 000 - If the account is not in arrears to accommodate performing loans
	// Count the number of days account is in arrears (e.g. 120 days in arrears)
	//
	// Number Field
	// Not negative
	// Mandatory Fields
	// If
	// A = Normal (0-30)
	// B = Watch (31-90)
	// C = Substandard (91-180)
	// D = Doubtful (181-360)
	// E = Loss ( >360)
	// The number of days in arrears cannot be more than the period the account has
	// been opened
	// The difference of the number of days in arrears cannot be greater than
	// the difference of the reporting period
	Attribute("", String, "Number of Days In Arrears", func() {
		Meta("rpc:tag", "46")
	})

	// Number of missed Instalments in the Facility.
	// Required if the Account is in arrears.
	// Conditional to Number of Days In Arrears, Overdue Date, Overdue Balance.
	//
	// Number
	// Conditional Check to 4.1.46, 4.1.45 and 4.1.44
	// Not negative
	// Not more instalments in arrears than the total number of instalments of the account.
	Attribute("InstalmentsInArrears", String, "Number of Instalments In Arrears", func() {
		Meta("rpc:tag", "47")
	})

	// A - Closed – No more admin processes running such as instalment
	// demands or interest charges to account, and no further facilities
	// can be offered on this account.
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
	// J – Called Up : The facility has been called up.
	// Once the client has paid up, the status should be updated to Option A,
	// H; or otherwise Option L
	// K – Suspended – The facility has been put on hold for an indefinite period of time
	// L– Client Deceased
	// M – Deferred – This refers to facilities whose payments have been
	// put on hold for a definite period or in moratorium (Grace Period)
	// N – Not Updated – This status is reserved for CRBs (if last record
	// status is not CLOSED)
	// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved
	//
	// Mandatory Field
	// Lookup
	// Not more than 1 character
	// The account Status cannot be Closed, Fully Settled with days in arrears greater than 0
	Attribute("AccountStatus", String, "Account Status", func() {
		Meta("rpc:tag", "48")
	})

	// The date of the status Change.
	//
	// Date Format Field
	// Mandatory Field
	// Cannot be in the future
	// Not greater than snapshot date
	// Not more than 8 characters
	Attribute("AccountStatusDate", String, "Account Status Date", func() {
		Meta("rpc:tag", "49")
	})

	// Reason for account closure.
	//
	// Alphanumeric
	// If Option “A” (closed) on Field 4.1.48 (Account Status) Must be filled
	// Not more than 50 characters
	Attribute("AccountClosureReason", String, "Account Closure Reason", func() {
		Meta("rpc:tag", "50")
	})

	// Repayment Period for the Facility in months.
	// This is the initial contractual period 999 for Credit Cards.
	//
	// Number Field Format
	// Mandatory With Warning
	// Not more than 3 characters
	Attribute("RepaymentPeriod", String, "Repayment Period", func() {
		Meta("rpc:tag", "51")
	})

	// Mandatory if Account Status = Deferred.
	// Deferred Payment Date must be in the future.
	//
	// Date Formats field
	// Date cannot be in the past
	// If account status =M "Deferred" we expect Deferred Payment date
	Attribute("DeferredPaymentDate", String, "Deferred Payment Date", func() {
		Meta("rpc:tag", "52")
	})

	// Conditional to field Account Status, Deferred Payment Date.
	//
	// Amount Field
	// Conditional Fields to 4.1.49, 4.1.52
	// Not more than 16 characters
	Attribute("DeferredPaymentAmount", String, "Deferred Payment Amount", func() {
		Meta("rpc:tag", "53")
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
	// Mandatory
	// Alphanumeric Field
	Attribute("PaymentFrequency", String, "Payment Frequency", func() {
		Meta("rpc:tag", "54")
	})

	// Date of facility drawdown.
	//
	// Mandatory Field
	// Date format
	// Cannot be less than the account open date.
	// Date cannot be in future
	Attribute("DisbursementDate", String, "Disbursement Date", func() {
		Meta("rpc:tag", "55")
	})

	// Instalment amount of loans.
	// Optional for Overdrafts, Credit Cards and Trade Finance facilities.
	//
	// Conditional based on 4.1.36
	// Not negative
	// If the account is not a final status this field is not zero
	// Not more than 16 characters
	Attribute("NextInstalmentAmount", String, "Next Instalment Amount", func() {
		Meta("rpc:tag", "56")
	})

	// Last date when payments were received into the facility.
	// Not mandatory for Trade Finance Products.
	// Null if no payment has been received.
	//
	// Date Format
	// Conditional if field 4.1.58 is populated
	Attribute("LatestPaymentDate", String, "Date of Latest Payment", func() {
		Meta("rpc:tag", "57")
	})

	// Last payment received into the facility.
	// Not mandatory for Trade Finance Products.
	// Mandatory if Date of Latest Payment is not null.
	//
	// Amount Field
	// Conditional if field 4.1.57 is populated
	// Not more than 16 characters
	Attribute("LastPaymentAmount", String, "Last Payment Amount", func() {
		Meta("rpc:tag", "58")
	})

	// If the Facility is Secured.
	// Options Available :
	// U - Unsecured
	// S - Fully Secured
	//
	// Lookup
	// Alphanumeric
	// Mandatory with warning
	// Should check against collateral API if secured
	// Not more than 1 character
	Attribute("SecurityType", String, "Type of Security", func() {
		Meta("rpc:tag", "59")
	})

})
