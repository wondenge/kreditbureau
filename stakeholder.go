package kreditbureau

type StakeholderType string

const (
	Director         StakeholderType = "A – Director"
	Shareholder                      = "B - Shareholder"
	Partner                          = "C - Partner"
	Proprietor                       = "D - Proprietor"
	TrustBeneficiary                 = "E - Trust Beneficiary"
	PowerOfAttorney                  = "F - Power of Attorney"
	Guarantor                        = "G - Guarantor"
)

type Stakeholder struct {
	CustomerName        CustomerName
	CompanyName         CompanyName
	DateOfBirth         DateOfBirth
	Gender              Gender
	Nationality         Nationality
	ClientNumber        ClientNumber
	AccountNumber       AccountNumber
	UserID              UserID
	EmailAddress        EmailAddress
	CompanyRegistration CompanyRegistration
	CompanyPINNumber    CompanyPINNumber
	CompanyVATNumber    CompanyVATNumber
	StakeholderType     StakeholderType

	// The Shareholding in the company/non-individual entity.
	// The value cannot be a negative value or exceed 100.
	SharePercentage float64
	TelephoneNumber TelephoneNumber
	PostalAddress   PostalAddress
	PhysicalAddress PhysicalAddress
}
