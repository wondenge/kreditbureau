package stakeholderid

type (
	// For Stakeholder, Guarantor, Collateral and Credit Application Files
	IDType int
)

const (
	NationalID IDType = iota
	Passport
	AlienRegistration
	ServiceID
	CompanyRegistrationNo
)
