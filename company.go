package kreditbureau

type Registrar struct {

	// The Name as Registered with the Registrar of Companies.
	RegisteredName string

	// The Business or Trading Name.
	TradingName string

	CompanyRegistration CompanyRegistration

	// The Type of Organisation.
	CompanyType OrganizationType

	// Industry Code for the Borrower’s Line of business.
	// Use the highest level of Central Bank of Kenya Codes.
	IndustryCode IndustryCode

	// The Borrower’s Annual Gross Turnover, in Kenya Shillings.
	AnnualTurnoverAmount float64

	PINNumber CompanyPINNumber
	VATNumber CompanyVATNumber

	// Indicate the number of persons who hold stake/shares in a non-individual entity.
	// 1 for Sole Proprietor.
	NumberOfStakeholders int

	// Status of Trading of the consumer.
	TradingStatus TradingStatus

	// Date of the Current Trading Status
	TradingStatusDate string
}

type CompanyRegistration struct {

	// Date Registered with the Registrar of Companies.
	RegistrationDate DateOfRegistration

	// The Registration Certificate Number - format to conform
	// to Registrar of Companies in Kenya
	RegistrationNumber string

	// The Registration Certificate Number - format to conform to
	// Registrar of Companies in Kenya.
	// Required, only if applicable because the company type changes,
	// the history details of the previous type can be matched to the new company type
	PreviousRegistrationNumber string
}
