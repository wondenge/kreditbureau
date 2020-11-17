package internal

type EmploymentType string

const (
	Casual       EmploymentType = "001 - Casual"
	Contract                    = "002 - Contract"
	Permanent                   = "003 - Permanent"
	Pensioner                   = "004 - Pensioner"
	SelfEmployed                = "005 - Self-Employed"
)

type Employer struct {

	// The Occupational Industry Type of the Borrower
	OccupationalIndustryType IndustryCode

	// Type of employment
	EmploymentType EmploymentType

	// The Employer Name.
	EmployerName string

	// The Date Consumer Was employed
	EmploymentDate string

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	IncomeAmount float64
}
