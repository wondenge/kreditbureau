package internal

type AccountIndicator string
type AccountProductType string
type AccountStatus string
type AccountNumber string
type NewAccountNumber string
type OldAccountNumber string

const (
	JointAccount  AccountIndicator = "J = Joint"
	SingleAccount                  = "S = Single"
)

const (
	Overdraft                   AccountProductType = "A – Overdraft"
	CreditCards                                    = "B – Credit Cards"
	BusinessWorkingCapitalLoans                    = "C – Business Working Capital Loans"
	BusinessExpansionLoans                         = "D – Business Expansion Loans"
	MortgageLoans                                  = "E – Mortgage Loans"
	AssetFinanceLoans                              = "F – Asset Finance Loans"
	TradeFinanceFacilities                         = "G – Trade Finance Facilities"
	PersonalLoans                                  = "H – Personal Loans"
	MobileLoan                                     = "I – Mobile Loan"
	InsurancePremiumFinancing                      = "J – Insurance Premium Financing"
	GroupLoans                                     = "K – Group Loans"
	UnclearedEffects                               = "L – Uncleared Effects"
)

const (
	// No more admin processes running such as instalment demands
	// or interest charges to account, and no further facilities
	// can be offered on this account.
	Closed AccountStatus = "A - Closed"

	// No activity for 2 years.
	// This applies for Overdraft/Current Accounts.
	Dormant = "B - Dormant"

	// For facilities that don’t form part of the outstanding portfolio
	// in the Balance Sheet, but are still outstanding in the books of accounts.
	WriteOff = "C - Write-Off"

	// With legal officer in court.
	Legal = "D - Legal"

	// With collection bureau.
	Collection = "E - Collection"

	// For facilities that form part of the outstanding portfolio, and
	// are reported in the Balance Sheet.
	Active = "F – Active"

	// For Rescheduled/Restructured Facilities
	FacilityRescheduled = "G – Facility Rescheduled"

	// The facility has been cleared.
	// This status can only be used for revolving facilities.
	Settled = "H –Settled"

	// The facility has been called up.
	// Once the client has paid up, the status should be updated to Option A, H;
	// or otherwise Option L.
	CalledUp = "J – Called Up"

	// The facility has been put on hold for an indefinite period of time.
	Suspended = "K – Suspended"

	ClientDeceased = "L– Client Deceased"

	// This refers to facilities whose payments have been put on hold for
	// a definite period or in moratorium (Grace Period).
	Deferred = "M – Deferred"

	// This status is reserved for CRBs (if last record status is not CLOSED).
	NotUpdated = "N – Not Updated"

	// Refers to a Record that the Client has disputed at the CRB.
	// The account cannot be updated until the dispute is resolved.
	Disputed = "P – Disputed"
)

type Account struct {
	Lenders            Lenders
	AccountIndicator   AccountIndicator
	OrganizationClass  OrganizationClass
	AccountProductType AccountProductType

	// Date When the Account was opened
	DateAccountOpened string

	// If the account is not in arrears, give the next instalment date.
	// If the account is in arrears, give the overdue date.
	// For Overdrafts, Credit Cards and Trade Finance products, submit the expiry date.
	InstalmentDueDate string

	// Classification of the Account
	PrudentialRiskClassification AssetClassification

	// For Loans, report the approved loan amount.
	// For Overdrafts, Credit cards and Trade Finance Products, report the approved limits.
	// If an overdraft was issued without a limit, the maximum amount overdrawn
	// should be reported instead.
	// All amounts should be a positive figure greater than Zero.
	// If the field is populated with a Zero (0) or a value less than Zero,
	// the record will be rejected.
	// If no value is supplied, the record will be rejected.
	OriginalAmount float64

	// The ISO Currency Code for the Currency in which the facility is given.
	// Default is KES for Kenya Shillings.
	FacilityCurrency CurrencyCode

	// The Current Balance in Kenya shillings, if the Currency is not Kenya shillings.
	// Use the exchange rate as at the reporting (end month) date.
	CurrentBalanceKES float64

	// Balance in the account as at reporting (end month) date.
	// For Overdrafts:
	// If the overdraft has been utilized, give the debit (negative) balance due.
	// If the overdraft has not been utilized, or if a previously submitted
	// overdraft has been paid up, give a Zero (0).
	// All amounts should be reported as positive values.
	// Negative values will be rejected.
	CurrentBalance float64

	// Arrears amount in a facility.
	// If the account is not in arrears, report Zero (0).
	OverdueBalance string

	// The Date when the Account fell overdue.
	OverdueDate string

	// The number of days the account has been in arrears.
	// (The Difference between the reporting date and the overdue date, calculated in days).
	// If the account is not in arrears, report 0.
	NumberOfDaysInArrears int

	// The Number of missed Instalments in the Facility.
	NumberOfInstalmentsInArrears int

	// The Account status as at the time of reporting.
	AccountStatus AccountStatus

	// The date of the status.
	AccountStatusDate string

	// Reason for account closure.
	AccountClosureReason string

	// Repayment Period for the Facility in months.
	// This is the initial contractual period.
	RepaymentPeriod int

	// The date when the next payment has been deferred to.
	// The Deferred Payment Date MUST be in the future.
	DeferredPaymentDate string

	// The Amount to be paid at the Deferred Payment Date.
	DeferredPaymentAmount float64

	// The Frequency in which Instalments are to be paid on the facility.
	PaymentFrequency PaymentFrequency

	// Date of facility drawdown.
	DisbursementDate string

	// Instalment amount of loans.
	// If the Instalment Amount is not yet due, input the expected instalment
	// amount (or amount to be repaid next) for the facility.
	NextInstalmentAmount float64

	// The date when payments were last received into the facility.
	// Not mandatory for Trade Finance Products.
	// Leave the Field Blank if no payment has been received.
	// Date should not be in the future.
	LatestPaymentDate string

	// Last payment Amount received into the facility.
	// If several payments were made on the date of latest payment,
	// return the sum of payments made that day.
	// If no payment has been received, Report 0.
	LastPaymentAmount float64

	// If the Facility is Secured.
	SecurityType SecurityType


}
