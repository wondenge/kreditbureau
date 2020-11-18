package internal

type CompanyName string
type DateOfBirth string
type DateOfRegistration string
type ClientType string
type ClientNumber string

const (
	IndividualCreditConsumer    ClientType = "A – Individual Credit Consumer"
	NonIndividualCreditConsumer            = "B – Non-Individual Credit Consumer"
)

type CustomerName struct {

	// The Family Name or Surname
	Surname string

	// The First Name
	Forename1 string

	// The Given Name
	Forename2 string

	// Other Name or Initials
	Forename3 string
}
