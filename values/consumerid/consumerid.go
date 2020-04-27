package consumerid

type (
	// For Individual Consumer File
	IDType int
)

const (
	NationalID IDType = iota
	Passport
	AlienRegistration
	ServiceID
)
