package design

import (
	. "goa.design/goa/v3/dsl"
)

// Individual Consumer, Employer and Account API
var Individual = Type("individual", func() {
	Description("Individual Consumer, Employer and Account API")
	TypeName("ICEAccount")
	ContentType("application/json")

	Attribute("Surname", String, func() {
		Description("Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
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

	// Name of a business if the individual runs a business as a sole proprietor
	// that is not registered;
	// or if the sole proprietor took a credit facility to finance the business.
	Attribute("TradingAs", String, func() {
		Description("Trading As")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	// Assist in matching for unique identification and detection of possible fraud.
	// The field needs to be mandatory if the Id No cannot be validated or changes
	// in case of Passport renewal & Passport number changes
	//
	// Not a future date
	// Not under 18
	// Not over 100 years
	Attribute("DoB", String, func() {
		Description("Date of Birth")
		Format(FormatDate)
		Meta("rpc:tag", "6")
	})

	// Client reference Number in Core Lenders system.
	// Client Number if Lender is Client Centric Account
	// Number if Lender is not Client Centric.
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "7")
	})

	// Account Number Linking Customer to Core Lenders system
	// Same as Client Number For Account-Centric System.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "8")
	})

	// Contract Unique Identifier used previously to report this facility.
	Attribute("OldAccountNumber", String, func() {
		Description("Old Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "9")
	})

	// Male/Female ( M/F)
	Attribute("Gender", String, func() {
		Description("Gender")
		Enum("Male", "Female")
		Meta("rpc:tag", "10")
	})

	Attribute("Nationality", String, func() {
		Description("ISO Country Code for the Consumer’s Nationality")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "11")
	})

	// M – Married
	// S - Single
	// D - Divorced
	// W - Widowed
	Attribute("MaritalStatus", String, func() {
		Description("Consumer’s Marital status")
		Enum("M", "S", "D", "W")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "12")
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
		Meta("rpc:tag", "13")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocNumber", String, func() {
		Description("Primary Identification Doc Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "14")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("SecondaryIDocumentType", Int, func() {
		Description("Secondary Identification Document Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "15")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("SecondaryIDocumentNumber", String, func() {
			Description("Secondary Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "16")
		})

		Required("SecondaryIDocumentNumber")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("OtherIDocumentType", Int, func() {
		Description("Other Identification Document Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "17")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("OtherIDocumentNumber", String, func() {
			Description("Other Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "18")
		})
		Required("OtherIDocumentNumber")
	})

	Attribute("PassportCountryCode", String, func() {
		Description("Country Code where the Passport was Registered")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "19")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("MobileTelephoneNumber", String, func() {
		Description("Mobile Telephone Number") // Consumers Main Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "20")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("HomeTelephoneNumber", String, func() {
		Description("Home Telephone Number") // Consumer’s Secondary Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "21")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("WorkTelephoneNumber", String, func() {
		Description("Work Telephone Number") // Consumer’s Office Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "22")
	})

	Attribute("PostalAddress1", String, func() {
		Description("Consumer’s Postal Address Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "23")
	})

	Attribute("PostalAddress2", String, func() {
		Description("Address Extra Details")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "24")
	})

	Attribute("PostalLocationTown", String, func() {
		Description("Consumer’s Postal Address Town")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "25")
	})

	Attribute("PostalLocationCountry", String, func() {
		Description("Country of Consumer’s Postal Address")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "26")
	})

	Attribute("PostCode", String, func() {
		Description("Post Code of Address")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "27")
	})

	Attribute("PhysicalAddress1", String, func() {
		Description("Consumer’s Physical ( Residential Address) Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "28")
	})

	Attribute("PhysicalAddress2", String, func() {
		Description("Consumer’s Physical (residential Address) Line 2")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "29")
	})

	Attribute("PlotNumber", String, func() {
		Description("Plot Land Ref (LR) No of Consumer’s residential Address")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "30")
	})

	Attribute("LocationTown", String, func() {
		Description("Town of Consumer’s residential Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "31")
	})

	Attribute("LocationCountry", String, func() {
		Description("ISO Code of the Country of the Consumer’s residential Address")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "32")
	})

	// A – Owned
	// B – Rented
	// C – Mortgaged
	// D – Family Owned
	Attribute("ResidencyType", String, func() {
		Description("Type of housing the borrower lives in")
		Enum("A", "B", "C", "D")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("D")
		Meta("rpc:tag", "33")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, func() {
		Description("Income Tax PIN No")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "34")
	})

	Attribute("ConsumerEmail", String, func() {
		Description("The Consumer Work Email Address") // If employed or available.
		MinLength(1)
		MaxLength(50)
		Format(FormatEmail) // Validate Email Syntax
		Meta("rpc:tag", "35")
	})

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
	Attribute("OccupationalIndustryType", String, func() {
		Description("Occupational Industry Type of the Borrower")
		Enum(001, 002, 003, 004, 005, 006, 007, 8, 9, 010, 011)
		MinLength(1)
		MaxLength(3)
		Minimum(001)
		Maximum(011)
		Meta("rpc:tag", "37")
	})

	// 001=Casual
	// 002= Contract
	// 003= Permanent
	// 004=Pensioner
	// 005= Self-Employed
	Attribute("EmploymentType", ArrayOf(Enum), func() {
		Description("Employment type")
		Enum(CasualEmployment,
			ContractEmployment,
			PermanentEmployment,
			PensionedEmployment,
			SelfEmployment,
		)
		Meta("rpc:tag", "39")
	})

	Attribute("LendersRegisteredName", String, func() {
		Description("Lenders Registered Name") // As registered with the Registrar of companies.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "41")
	})

	Attribute("LendersTradingName", String, func() {
		Description("Lenders Trading Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "42")
	})

	Attribute("LendersBranchName", String, func() {
		Description("Lenders Branch Name")
		MinLength(1)
		MaxLength(50)
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
	// IXXX MUST match the API name
	// XBB (code of the institution making the submission)
	Attribute("LendersBranchCode", String, func() {
		Description("Lenders Branch Code")
		MinLength(1)
		MaxLength(7)
		Meta("rpc:tag", "44")
	})

	// If the Account is operated jointly
	// or by a single individual.
	// J = Joint
	// S = Single
	Attribute("AJSIndicator", String, func() {
		Description("Account Joint/Single Indicator")
		Enum("J", "S")
		MinLength(1)
		MinLength(1)
		Meta("rpc:tag", "45")
	})

	// The Account Product Type :
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
	Attribute("AccountProductType", ArrayOf(Enum), func() {
		Description("Account Product Type")
		Enum(
			IndividualsOverdraft,                   // A – Overdraft
			IndividualsCreditCards,                 // B – Credit Cards
			IndividualsBusinessWorkingCapitalLoans, // C – Business Working Capital Loans
			IndividualsBusinessExpansionLoans,      // D – Business Expansion Loans
			IndividualsMortgageLoans,               // E – Mortgage Loans
			IndividualsAssetFinanceLoans,           // F – Asset Finance Loans
			IndividualsTradeFinance,                // G – Trade Finance Facilities
			IndividualsPersonalLoans,               // H – Personal Loans
			IndividualsMobileLoan,                  // I – Mobile Loan
			IndividualsInsurancePremiumFinancing,   // J – Insurance Premium Financing
			IndividualsGroupLoans,                  // K – Group Loans
			IndividualsUnclearedEffects,            // L – Uncleared Effects
		)
		Meta("rpc:tag", "46")
	})

	Attribute("DateAccountOpened", String, func() {
		Description("Date the Account was opened")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "47")
	})

	// If the account is not in arrears, give the next instalment date.
	// If the account is in arrears, give the overdue date.
	// For Overdrafts, Credit Cards and Trade Finance products, give the expiry date.
	Attribute("InstalmentDueDate", String, func() {
		Description("Instalment Due Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "48")
	})

	// For Loans, give the approved loan amount.
	// For Overdrafts, Credit cards and Trade Finance Products, issue the approved limits.
	Attribute("OriginalAmount", Int64, func() {
		Description("Original Amount")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "49")
	})

	// The ISO Currency Code for the Currency in which the facility is given.
	Attribute("FacilityCurrency", String, func() {
		Description("Currency of Facility")
		MinLength(3)
		MaxLength(3)
		Default("KES") // Default is KES for Kenya Shillings.
		Meta("rpc:tag", "50")
	})

	// The current balance in Kenya shillings,
	// if the Currency is not Kenya shillings.
	// Same as the Original Amount if the facility was given in Kenya Shillings.
	// Use the exchange rate as at the reporting (month end) date.
	Attribute("CurrentBalanceInKES", Int64, func() {
		Description("Current Balance in Kenya Shillings")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "51")
	})

	// Balance in the account as at reporting (end month) date
	// For Overdrafts:
	// If the overdraft has been utilized, give the debit (negative) balance due.
	// If the overdraft has not been utilized, or if a previously submitted
	// overdraft has been paid up, give a Zero (0)
	Attribute("CurrentBalance", Int64, func() {
		Description("Current Balance")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "52")
	})

	// Arrears amount in a facility.
	// Mandatory if Overdue date is not null,
	// Number of Days in Arrears & Number of Instalments in Arrears > 0
	Attribute("OverdueBalance", Int64, func() {
		Description("Overdue Balance")
		Meta("rpc:tag", "53")

		// If the Account is in arrears, the Date when the Account fell overdue.
		// For Overdrafts, Trade Finance Products and Credit Cards, the
		// date of expiry or date it went over the limit Conditional to Overdue Balance,
		// Number of Days In Arrears and Number of Instalments In Arrears.
		Attribute("OverdueDate", String, func() {
			Description("Overdue Date")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "54")
		})
		// 000 - If the account is not in arrears to accommodate performing loans
		// Count the number of days account is in arrears (e.g. 120 days in arrears)
		Attribute("DaysInArrears", Int, func() {
			Description("Number of Days In Arrears")
			MinLength(1)
			MaxLength(5)
			Meta("rpc:tag", "55")
		})

		// Number of missed Instalments in the Facility.
		// Required if the Account is in arrears.
		// Conditional to Number of Days In Arrears, Overdue Date, Overdue Balance
		Attribute("InstalmentsInArrears", Int, func() {
			Description("Number of Instalments In Arrears")
			MinLength(1)
			MaxLength(3)
			Meta("rpc:tag", "56")
		})
		Required("OverdueDate", "DaysInArrears", "InstalmentsInArrears")
	})

	// This is as per CBK guidelines on classification of assets
	// A = Normal (0-29)
	// B = Watch(31-89)
	// C = Substandard (91-179)
	// D = Doubtful( 181-360)
	// E = Loss (>360)
	Attribute("PrudentialRiskClassification", String, func() {
		Description("Prudential Risk Classification")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("B")
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
	Attribute("AccountStatus", ArrayOf(Enum), func() {
		Description("Account Status")
		Enum(ClosedIndividualsAccountStatus, // A - Closed
			WriteOffIndividualsAccountStatus,       // C - Write-Off
			LegalIndividualsAccountStatus,          // D - Legal
			CollectionIndividualsAccountStatus,     // E - Collection
			ActiveIndividualsAccountStatus,         // F – Active
			RescheduledIndividualsAccountStatus,    // G – Facility Rescheduled
			SettledIndividualsAccountStatus,        // H –Settled
			CalledUpIndividualsAccountStatus,       // J – Called Up
			SuspendedIndividualsAccountStatus,      // K – Suspended
			ClientDeceasedIndividualsAccountStatus, // L– Client Deceased
			DeferredIndividualsAccountStatus,       // M – Deferred
			NotUpdatedIndividualsAccountStatus,     // N – Not Updated
			DisputedIndividualsAccountStatus,       // P – Disputed
		)
		Meta("rpc:tag", "58")
	})

	// The repayment Period, in months for the Facility.
	// 999 for Credit Cards
	Attribute("RepaymentPeriod", String, func() {
		Description("Repayment Period")
		MinLength(1)
		MaxLength(3)
		Meta("rpc:tag", "61")
	})

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
		Description("Frequency in which Instalments are to be paid on the facility")
		Enum("W", "M", "Q", "S", "Y", "B", "T", "C", "D", "F")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "64")
	})

	Attribute("DisbursementDate", String, func() {
		Description("Disbursement Date") // Date of facility drawdown.
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "65")
	})

	// U - Unsecured
	// S - Fully Secured
	Attribute("SecurityType", String, func() {
		Description("Indicator of whether the facilities are secured or Not")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "69")
	})

	// The Unique Identifier of the Group the borrower belongs
	// to under which the facility was granted.
	Attribute("GroupID", String, func() {
		Description("Group ID")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "70")
	})

})

var StoredIndividual = ResultType("application/vnd.goa.individual", func() {
	TypeName("StoredIndividual")
	Attributes(func() {
		Extend(Individual)
		Required(
			"Surname",
			"Forename1",
			"DoB",
			"AccountNumber",
			"Gender",
			"Nationality",
			"PrimaryIDocumentType",
			"PrimaryIDocNumber",
			"PassportCountryCode",
			"LocationCountry",
			"OccupationalIndustryType",
			"EmploymentType",
			"LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"AJSIndicator",
			"AccountProductType",
			"DateAccountOpened",
			"InstalmentDueDate",
			"OriginalAmount",
			"FacilityCurrency",
			"CurrentBalanceInKES",
			"CurrentBalance",
			"OverdueBalance",
			"PrudentialRiskClassification",
			"AccountStatus",
			"RepaymentPeriod",
			"PaymentFrequency",
			"DisbursementDate",
			"SecurityType",
		)
	})
	View("default", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("TradingAs")
		Attribute("DoB")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("OldAccountNumber")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("MaritalStatus")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocumentType")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocumentType")
		Attribute("OtherIDocumentNumber")
		Attribute("PassportCountryCode")
		Attribute("MobileTelephoneNumber")
		Attribute("HomeTelephoneNumber")
		Attribute("WorkTelephoneNumber")
		Attribute("PostalAddress1")
		Attribute("PostalAddress2")
		Attribute("PostalLocationTown")
		Attribute("PostalLocationCountry")
		Attribute("PostCode")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalAddress2")
		Attribute("PlotNumber")
		Attribute("LocationTown")
		Attribute("LocationCountry")
		Attribute("ResidencyType")
		Attribute("PINumber")
		Attribute("ConsumerEmail")
		Attribute("EmployerName")
		Attribute("OccupationalIndustryType")
		Attribute("EmploymentDate")
		Attribute("EmploymentType")
		Attribute("IncomeAmount")
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AJSIndicator")
		Attribute("AccountProductType")
		Attribute("DateAccountOpened")
		Attribute("InstalmentDueDate")
		Attribute("OriginalAmount")
		Attribute("FacilityCurrency")
		Attribute("CurrentBalanceInKES")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("OverdueDate")
		Attribute("DaysInArrears")
		Attribute("InstalmentsInArrears")
		Attribute("PrudentialRiskClassification")
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
		Attribute("GroupID")
	})
	View("mandatory", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("DoB")
		Attribute("AccountNumber")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("PrimaryIDocumentType")
		Attribute("PrimaryIDocNumber")
		Attribute("PassportCountryCode")
		Attribute("LocationCountry")
		Attribute("OccupationalIndustryType")
		Attribute("EmploymentType")
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AJSIndicator")
		Attribute("AccountProductType")
		Attribute("DateAccountOpened")
		Attribute("InstalmentDueDate")
		Attribute("OriginalAmount")
		Attribute("FacilityCurrency")
		Attribute("CurrentBalanceInKES")
		Attribute("CurrentBalance")
		Attribute("OverdueBalance")
		Attribute("PrudentialRiskClassification")
		Attribute("AccountStatus")
		Attribute("RepaymentPeriod")
		Attribute("PaymentFrequency")
		Attribute("DisbursementDate")
		Attribute("SecurityType")
	})
})
