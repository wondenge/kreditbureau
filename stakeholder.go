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
