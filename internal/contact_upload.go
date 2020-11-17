package internal

type ContactUpload struct {
	SubmissionDate        SubmissionDate
	ClientType            ClientType
	CustomerName          CustomerName
	CompanyName           CompanyName
	PrimaryID             PrimaryID
	AccountNumber         AccountNumber
	TelephoneNumber       TelephoneNumber
	PostalAddress         PostalAddress
	PhysicalAddress       PhysicalAddress
	DateAtPhysicalAddress DateAtPhysicalAddress
	EmailAddress          EmailAddress
}
