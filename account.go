package kreditbureau

type AccountIndicator string
type AccountProductType string
type AccountStatus string

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
