package internal

type TelephoneNumber struct {
	MobileNumber string
	HomeNumber   string
	WorkNumber   string
}

type PostalAddress struct {
	PostalAddress1 string
	PostalAddress2 string
	PostalTown     string
	PostalCountry  string
	PostalCode     int
}

type PhysicalAddress struct {
	PhysicalAddress1 string
	PhysicalAddress2 string
	PlotNumber       string
	PhysicalTown     string
	PhysicalCountry  string
}

type EmailAddress string
