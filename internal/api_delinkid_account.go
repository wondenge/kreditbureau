package internal


type IDToLink struct {
	CustomerName  CustomerName
	AccountNumber AccountNumber

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Link.
	LinkPrimaryID PrimaryID
}

type IDToDelink struct {
	CustomerName  CustomerName
	AccountNumber AccountNumber

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Delink.
	DelinkPrimaryID PrimaryID
}
