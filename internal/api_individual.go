package internal

// Individual Consumer, Employer and Account File.
type Individual struct {
	CustomerName CustomerName

	// Name of a business if the individual runs a business as a sole proprietor that
	// is not registered; or if the sole proprietor took a credit facility to finance the business.
	TradingAs        string
	DateOfBirth      DateOfBirth
	ClientNumber     ClientNumber
	AccountNumber    AccountNumber
	OldAccountNumber OldAccountNumber
	Gender           Gender
	Nationality      Nationality
	MaritalStatus    MaritalStatus
	UserID           UserID

	// The ISO Country Code where the Passport was registered.
	PassportCountryCode int
	TelephoneNumber     TelephoneNumber
	PostalAddress       PostalAddress
	PhysicalAddress     PhysicalAddress
	ResidencyType       ResidencyType
	PersonalPINNumber   PersonalPINNumber
	EmailAddress        EmailAddress
	Employer            Employer
	IndividualAccount   Account

	// The Unique Identifier of the Group the borrower belongs to under which
	// the facility was granted.
	GroupID string
}
