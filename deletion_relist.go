package kreditbureau

type FunctionType string

const (
	Complete            FunctionType = "A-Complete"
	DefaultHistory                   = "B-Default History"
	SpecificMonthUpdate              = "C-Specific Month Update"
	Relist                           = "D- Relist"
)

type DeletionRelist struct {

	// The type of the intended deletion.
	FunctionType FunctionType
	CustomerName CustomerName

	// Primary Identification Document provided on opening the account.
	PrimaryID     PrimaryID
	AccountNumber AccountNumber

	// The date the record is to be effected.
	SubmissionDate SubmissionDate

	// The new number of days the account has been in arrears.
	// This should be rounded to 20days.
	NewDaysInArrears int
}
