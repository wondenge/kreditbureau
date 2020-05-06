package design

import (
	. "goa.design/goa/v3/dsl"
)

// Non Individual Consumer and Account API
var Company = Type("company", func() {
	Description("Non Individual Consumer and Account API")
	TypeName("Company")
	ContentType("application/json")

	Attribute("RegisteredName", String, func() {
		Description("Name as Registered with the Registrar of Companies")
		MinLength(1)
		MaxLength(150)
		Meta("rpc:tag", "1")
	})

	Attribute("TradingName", String, "Trading Name", func() {
		Description("Business or Trading Name")
		MinLength(1)
		MaxLength(150)
		Meta("rpc:tag", "2")
	})

	Attribute("RegistrationDate", String, func() {
		Description("Date Registered with the Registrar of Companies")
		Format(FormatDate)
		MaxLength(8)
		Meta("rpc:tag", "3")
	})

	Attribute("RegistrationNumber", String, func() {
		Description("The Registration Certificate Number")
		MinLength(1)
		MaxLength(25)
		Meta("rpc:tag", "4")
	})

	// The Registration Certificate Number
	// - format to conform to Registrar of Companies in Kenya.
	// Required, only if applicable because the company type changes,
	// the history details of the previous type can be matched to the
	// new company type.
	Attribute("PreviousRegistrationNumber", String, func() {
		Description("Previous Registration Number")
		MinLength(1)
		MaxLength(25)
		Meta("rpc:tag", "5")
	})

	// The Country of Registration of the Company,
	// Defaults to Kenyan.
	// Matches COUNTRY ISO CODE
	Attribute("Nationality", String, func() {
		Description("Nationality")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "6")
	})

	// The Client Reference Number linking Company to Core Banking system.
	// Client reference Number in Core Lenders system.
	// Client Number if Lender is Client Centric.
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "7")
	})

	// The Account Number Linking Client to Bank’s Accounting System.
	// Account Number if Lender is not Client Centric.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "8")
	})

	// Old Account Number linking to the previous MIS System
	Attribute("OldAccountNumber", String, func() {
		Description("Old Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "9")
	})

	// The Type of Organisation.
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
	Attribute("CompanyType", String, func() {
		Description("Company Type")
		Enum("A", "B", "C", "D", "E", "F", "G", "H", "J", "L", "M")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("M")
		Meta("rpc:tag", "10")
	})

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
	Attribute("IndustryCode", Int, func() {
		Description("Industry Code for the Borrower’s Line of business")
		Enum(001, 002, 003, 004, 005, 006, 007, 8, 9, 010, 011, 012)
		MinLength(1)
		MaxLength(3)
		Minimum(001)
		Maximum(012)
		Meta("rpc:tag", "11")
	})

	// Use the Central Bank of Kenya range options.
	Attribute("AnnualTurnoverAmount", Int64, func() {
		Description("Borrower’s Annual Gross Turnover Range, in KES")
		Minimum(0)
		Meta("rpc:tag", "12")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINNumber", String, func() {
		Description("Income Tax PIN No")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "13")
	})

	Attribute("VATNumber", String, func() {
		Description("Value Added Tax Registration Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "14")
	})

	// Indicate the number of persons who hold stake/shares in
	// a non-individual entity. 1 for Sole Proprietor.
	Attribute("NofStakeholders", Int, func() {
		Description("Number of Stakeholders")
		Minimum(1)
		Meta("rpc:tag", "15")
	})

	// 001=Dormant
	// 002=Actively Trading
	// 003=Under Receivership/Liquidation
	// 004=Liquidated
	// 005=Under Management
	// 006=Dissolved
	Attribute("TradingStatus", Int, func() {
		Description("Status of Trading of the consumer")
		Enum(001, 002, 003, 004, 005, 006)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(006)
		Meta("rpc:tag", "16")
	})

	// Not a date before the company registration date.
	Attribute("TradingStatusDate", String, func() {
		Description("Date of the Current Trading Status.")
		Format(FormatDate)
		Meta("rpc:tag", "17")
	})

	// Format of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("MainTelephoneNumber", String, func() {
		Description("Primary Telephone contact Number") // Main Telephone Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "18")
	})

	// Format of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("OtherTelephoneNumber", String, func() {
		Description("Other Telephone Number")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "19")
	})

	Attribute("PostalAddress1", String, func() {
		Description("Consumer’s Postal Address Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "20")
	})

	Attribute("PostalAddress2", String, func() {
		Description("Postal Address Extra Details")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "21")
	})

	Attribute("Town", String, func() {
		Description("Postal Town of Postal Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "22")
	})

	Attribute("Country", String, func() {
		Description("Country of Operation for the Consumer")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "23")
	})

	Attribute("PostCode", String, func() {
		Description("Post Code of Consumer’s Postal Address")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "24")
	})

	Attribute("PhysicalAddress1", String, func() {
		Description("Physical Address Line 1 of the Consumer’s business Location ( Office Location)")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "25")
	})

	Attribute("PhysicalAddress2", String, func() {
		Description("Physical Address Line 2 of the Consumer’s business Location ( Office Location)")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "26")
	})

	Attribute("PlotNumber", String, "Plot Number", func() {
		Description("Plot Land Ref (LR) No of the Consumer’s business Location ( Office Location)")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "27")
	})

	Attribute("PhysicalLocationTown", String, func() {
		Description("Town of the Consumer’s business Location ( Office Location)")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "28")
	})

	Attribute("PhysicalLocationCountry", String, func() {
		Description("Country of the Office Location")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "29")
	})

	Attribute("LendersRegisteredName", String, func() {
		Description("Name of Lender as registered with the Registrar of companies")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "30")
	})

	Attribute("LendersTradingName", String, func() {
		Description("The Lenders Trading Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "31")
	})

	Attribute("LendersBranchName", String, func() {
		Description("Lenders Branch Name")
		MinLength(1)
		MaxLength(50)
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
	// IXXX MUST match the API name XBB (code of the institution making the submission)
	Attribute("LendersBranchCode", String, func() {
		Description("Lenders Branch Code")
		MaxLength(7)
		Meta("rpc:tag", "33")
	})

	// J = Joint
	// S = Single
	Attribute("AccountIndicator", String, func() {
		Description("Whether the Account is operated Jointly or By a single individual")
		Enum("J", "S")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "34")
	})

	// A – Corporate
	// B – SME
	// C - NGO/Church/Society
	Attribute("OrganisationClass", String, func() {
		Description("Class of Organisation")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("B")
		Meta("rpc:tag", "35")
	})

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
	// Option B,H, I, K not allowed
	Attribute("AccountProductType", ArrayOf(Enum), func() {
		Description("Account Product Type")
		Enum(CompanyOverdraft,
			CompanyCreditCards,
			CompanyBusinessExpansionLoans,
			CompanyMortgageLoans,
			CompanyAssetFinanceLoans,
			CompanyTradeFinance,
			CompanyInsurancePremiumFinancing,
			CompanyUnclearedEffects,
		)
		Meta("rpc:tag", "36")
	})

	// Date should not be more than the snapshot date
	Attribute("DateAccountOpened", String, func() {
		Description("Date When the Account was opened")
		Format(FormatDate)
		Meta("rpc:tag", "37")
	})

	// If the account is not in arrears, give the next instalment date.
	// If the account is in arrears, give the overdue date
	// For Overdrafts, Credit Cards and Trade Finance products, give the expiry date.
	// - Not date in the past for an account not in arrears.
	// - Not date in the future for an account in arrears.
	// - Future Date for Overdrafts, credit cards, trade finance.
	Attribute("InstalmentDueDate", String, func() {
		Description("Instalment Due Date")
		Format(FormatDate)
		Meta("rpc:tag", "38")
	})

	// A = Normal (0-30)
	// B = Watch (>30-90)
	// C = Substandard (>90-180)
	// D = Doubtful (>180-360)
	// E = Loss (>360)
	Attribute("PrudentialRiskClassification", String, func() {
		Description("Prudential Risk Classification") // The guidelines as per CBK on loan classification
		Enum("A", "B", "C", "D", "E")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("E")
		Meta("rpc:tag", "39")
	})

	// For Loans, give the approved loan amount.
	// For Overdrafts, Credit cards and Trade Finance Products, issue the approved limits.
	Attribute("OriginalAmount", Int64, func() {
		Description("Original Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "40")
	})

	// The ISO Currency Code for the Currency in which the facility is given.
	// Default is KES for Kenya Shillings.
	// Refer to Appendix B for valid Currency Codes
	Attribute("FacilityCurrency", String, func() {
		Description("Currency of Facility")
		MinLength(3)
		MaxLength(3)
		Default("KES")
		Meta("rpc:tag", "41")
	})

	// The Current Balance in Kenya shillings, if the
	// Currency is not Kenya shillings.
	// Same as the Original Amount if the facility was given in Kenya Shillings.
	// Use the exchange rate as at the reporting (end month) date.
	// The Value should be same as 4.1.43 if field 4.1.41 is in KES
	Attribute("KESCurrentBalance", Int64, func() {
		Description("Current Balance in Kenya Shillings")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "42")
	})

	// For Overdrafts:
	// If the overdraft has been utilized, give the debit (negative) balance due.
	// If the overdraft has not been utilized, or if a previously submitted
	// overdraft has been paid up, give a Zero (0)
	//
	// Cannot be zero if Overdue Balance > 0
	// Cannot be Zero if DPD > 0
	// Cannot be > 0 if acct status is
	// Closed, Paid Up, Fully Settled, Early Settlement
	Attribute("CurrentBalance", Int64, func() {
		Description("Balance in the account as at reporting (end month) date")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "43")
	})

	//
	// Mandatory if Overdue date is not null, Nr of Days in Arrears
	// & Nr of Instalments in Arrears > 0
	//
	// Must be filled if 4.1.45
	Attribute("OverdueBalance", Int64, func() {
		Description("Arrears amount in a facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "44")

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
		Attribute("OverdueDate", String, func() {
			Description("Overdue Date")
			Format(FormatDate)
			Meta("rpc:tag", "45")
		})

		// 000 - If the account is not in arrears to accommodate performing loans
		// Count the number of days account is in arrears (e.g. 120 days in arrears)
		//
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
		Attribute("DaysInArrears", Int, func() {
			Description("Number of Days In Arrears")
			Minimum(0)
			Meta("rpc:tag", "46")
		})

		// Number of missed Instalments in the Facility.
		// Required if the Account is in arrears.
		// Conditional to Number of Days In Arrears, Overdue Date, Overdue Balance.
		//
		// Conditional Check to 4.1.46, 4.1.45 and 4.1.44
		// Not more instalments in arrears than the total number of instalments of the account.
		Attribute("InstalmentsInArrears", Int, func() {
			Description("Number of Instalments In Arrears")
			Minimum(0)
			Meta("rpc:tag", "47")
		})

		Required("OverdueDate", "DaysInArrears", "InstalmentsInArrears")
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
	// The account Status cannot be Closed, Fully Settled with days in arrears greater than 0
	Attribute("AccountStatus", ArrayOf(Enum), func() {
		Description("Account Status")
		Enum(ClosedCompanyAccountStatus, // A - Closed
			WriteOffCompanyAccountStatus,       // C - Write-Off
			LegalCompanyAccountStatus,          // D - Legal
			CollectionCompanyAccountStatus,     // E - Collection
			ActiveCompanyAccountStatus,         // F – Active
			RescheduledCompanyAccountStatus,    // G – Facility Rescheduled
			SettledCompanyAccountStatus,        // H –Settled
			CalledUpCompanyAccountStatus,       // J – Called Up
			SuspendedCompanyAccountStatus,      // K – Suspended
			ClientDeceasedCompanyAccountStatus, // L– Client Deceased
			DeferredCompanyAccountStatus,       // M – Deferred
			NotUpdatedCompanyAccountStatus,     // N – Not Updated
			DisputedCompanyAccountStatus,       // P – Disputed
		)
		Meta("rpc:tag", "48")
	})

	// This is the initial contractual period 999 for Credit Cards.
	Attribute("RepaymentPeriod", Int, func() {
		Description("Repayment Period for the Facility in months")
		MinLength(1)
		MaxLength(3)
		Minimum(0)
		Meta("rpc:tag", "51")
	})

	// The Frequency in which Instalments are to be paid on the facility.
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
	Attribute("PaymentFrequency", String, func() {
		Description("Payment Frequency")
		Enum("W", "M", "Q", "S", "Y", "B", "T", "C", "D", "F")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "54")
	})

	// Cannot be less than the account open date.
	Attribute("DisbursementDate", String, func() {
		Description("Disbursement Date") // Date of facility drawdown.
		Format(FormatDate)
		Meta("rpc:tag", "55")
	})

	// U - Unsecured
	// S - Fully Secured
	Attribute("SecurityType", String, func() {
		Description("Type of Security") // If the Facility is Secured.
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "59")
	})

})

var StoredCompany = ResultType("application/vnd.goa.company", func() {
	TypeName("StoredCompany")
	Attributes(func() {
		Extend(Company)
		Required(
			"RegisteredName",
			"TradingName",
			"RegistrationDate",
			"RegistrationNumber",
			"Nationality",
			"ClientNumber",
			"AccountNumber",
			"CompanyType",
			"IndustryCode",
			"AnnualTurnoverAmount",
			"PINNumber",
			"TradingStatus",
			"MainTelephoneNumber",
			"PhysicalAddress1",
			"PhysicalLocationTown",
			"PhysicalLocationCountry",
			"LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"AccountProductType",
			"DateAccountOpened",
			"InstalmentDueDate",
			"PrudentialRiskClassification",
			"OriginalAmount",
			"FacilityCurrency",
			"KESCurrentBalance",
			"CurrentBalance",
			"AccountStatus",
			"RepaymentPeriod",
			"PaymentFrequency",
			"DisbursementDate",
			"SecurityType",
		)
	})

	View("default", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("PreviousRegistrationNumber")
		Attribute("Nationality")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("OldAccountNumber")
		Attribute("CompanyType")
		Attribute("IndustryCode")
		Attribute("AnnualTurnoverAmount")
		Attribute("PINNumber")
		Attribute("VATNumber")
		Attribute("NofStakeholders")
		Attribute("TradingStatus")
		Attribute("TradingStatusDate")
		Attribute("MainTelephoneNumber")
		Attribute("OtherTelephoneNumber")
		Attribute("PostalAddress1")
		Attribute("PostalAddress2")
		Attribute("Town")
		Attribute("Country")
		Attribute("PostCode")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalAddress2")
		Attribute("PlotNumber")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AccountIndicator")
		Attribute("OrganisationClass")
		Attribute("AccountProductType")
		Attribute("DateAccountOpened")
		Attribute("InstalmentDueDate")
		Attribute("PrudentialRiskClassification")
		Attribute("OriginalAmount")
		Attribute("FacilityCurrency")
		Attribute("KESCurrentBalance")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("OverdueDate")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("AccountStatus")
		Attribute("AccountStatusDate")
		Attribute("AccountClosureReason")
		Attribute("RepaymentPeriod")
		Attribute("DeferredPaymentDate")
		Attribute("DeferredPaymentAmount")
		Attribute("PaymentFrequency")
		Attribute("DisbursementDate")
		Attribute("NextInstalmentAmount")
		Attribute("LatestPaymentDate")
		Attribute("LastPaymentAmount")
		Attribute("SecurityType")
	})

	View("mandatory", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("Nationality")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("CompanyType")
		Attribute("IndustryCode")
		Attribute("AnnualTurnoverAmount")
		Attribute("PINNumber")
		Attribute("TradingStatus")
		Attribute("MainTelephoneNumber")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AccountProductType")
		Attribute("DateAccountOpened")
		Attribute("InstalmentDueDate")
		Attribute("PrudentialRiskClassification")
		Attribute("OriginalAmount")
		Attribute("FacilityCurrency")
		Attribute("KESCurrentBalance")
		Attribute("CurrentBalance")
		Attribute("AccountStatus")
		Attribute("RepaymentPeriod")
		Attribute("PaymentFrequency")
		Attribute("DisbursementDate")
		Attribute("SecurityType")
	})
})
