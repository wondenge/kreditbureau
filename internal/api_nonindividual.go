package internal


// Non-individual Consumer and Account File.
type Company struct {
	Registrar        Registrar
	Nationality      Nationality
	ClientNumber     ClientNumber
	AccountNumber    AccountNumber
	OldAccountNumber OldAccountNumber

	TelephoneNumber TelephoneNumber
	PostalAddress   PostalAddress
	PhysicalAddress PhysicalAddress
	CompanyAccount  Account
}
