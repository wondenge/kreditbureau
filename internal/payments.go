package internal

type PaymentFrequency string

const (
	Weekly       PaymentFrequency = "W - Weekly"
	Monthly                       = "M - Monthly"
	Quarterly                     = "Q – Quarterly"
	SemiAnnually                  = "S – Semi-Annually"
	Annually                      = "Y – Annually ( Yearly)"
	Bullet                        = "B – Bullet"
	TriAnnual                     = "T – Tri-annual (Three times in a year)"
	Biennial                      = "C – Biennial (Once in two years)"
	Daily                         = "D – Daily"
	Fortnightly                   = "F - Fortnightly"
)

// Repayments are daily reports payment made into credit facility.
type Repayments struct {
	SnapshotDate     SnapshotDate
	ClientType       ClientType
	CustomerName     CustomerName
	CompanyName      CompanyName
	PrimaryID        PrimaryID
	AccountNumber    AccountNumber
	FacilityCurrency CurrencyCode

	// For Loans, submit the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; submit the approved limits.
	// For overdrafts with no limits, input the maximum amount overdrawn.
	OriginalAmount float64

	// The sum total amount paid into the loan on the Payment Date.
	// The amount must be a positive value, which will result in a reduction of Current Balance.
	PaymentAmount float64

	// The date when a payment was made into the Loan
	PaymentDate string

	// The new Current Balance in the loan account after the payment.
	CurrentBalance float64

	// The total number of instalments outstanding after this payment.
	// If the Loan is not in arrears after the payment, return report 0.
	InstallmentsInArrears int

	// The number of days that the loan is overdue.
	// If the loan is not in arrears after the payment, return 0.
	DaysInArrears int

	AccountStatus AccountStatus
}
