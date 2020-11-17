package kreditbureau

type IDType string

const (
	NationalID            IDType = "001 – National ID"
	Passport                     = "002 – Passport"
	AlienRegistration            = "003 – Alien Registration"
	ServiceID                    = "004 – Service ID"
	CompanyRegistrationNo        = "005 - Company Registration No"
)

type PrimaryID struct {
	PrimaryIDType   IDType
	PrimaryIDNumber string
}

type SecondaryID struct {
	SecondaryIDType   IDType
	SecondaryIDNumber string
}

type OtherID struct {
	OtherIDType   IDType
	OtherIDNumber string
}

type UserID struct {
	PrimaryID PrimaryID
	SecondaryID SecondaryID
	OtherID OtherID
}
