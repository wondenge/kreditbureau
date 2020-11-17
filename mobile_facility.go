package kreditbureau

type MobileFacility struct {
	CustomerName CustomerName

	// Name of a business if the individual runs a business as a sole
	// proprietor that is not registered;
	// or if the sole proprietor took a credit facility to finance the business.
	TradingAs       string
	DateOfBirth     DateOfBirth
	ClientNumber    ClientNumber
	AccountNumber   AccountNumber
	Gender          Gender
	Nationality     Nationality
	PrimaryID       PrimaryID
	TelephoneNumber TelephoneNumber

	AccountProductType

	// If the account is not in arrears, report the next instalment date.
	// If the account is in arrears, report the overdue date
	InstalmentDueDate string

	// The principal Amount issued
	OriginalAmount float64

	// The currency code for the facility's currency
	FacilityCurrency int

	// Current Balance equivalent in Kenya shillings
	CurrentBalanceInKES float64

	// The current balance in the mobile account.
	CurrentBalance float64

	// Overdue Amount in a facility.
	// If account is not in arrears report Zero (0).
	OverdueBalance float64

	// The date when the account fell Overdue.
	OverdueDate string

	// The number of days the account has been in arrears
	//(Difference between reporting date and overdue date, calculated in days).
	// If the account is not in arrears, report Zero(0)
	DaysInArrears int

	// The Number of missed Instalments in the Facility.
	InstalmentsInArrears int

	// The date of the status.
	AccountStatusDate string

	// Repayment Period for the Facility in months.
	// This is the initial contractual period.
	RepaymentPeriod int

	// Date of facility drawdown.
	DisbursementDate string

	// Instalment amount for the loan.
	InstalmentAmount float64

	// The date when payments were last received into the facility.
	LatestPaymentDate string

	// Last Payment Amount received into the facility Report Zero(0) if no payment was received.
	LastPaymentAmount float64
}
