package internal

type LinkDelink string

const (
	Link   LinkDelink = "001 - Link"
	Delink            = "002 – Delink"
)

type LinkIDFromAccount struct {
	CustomerName  CustomerName
	AccountNumber AccountNumber

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Link.
	LinkPrimaryID PrimaryID
}

type DelinkIDFromAccount struct {
	CustomerName  CustomerName
	AccountNumber AccountNumber

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Delink.
	DelinkPrimaryID PrimaryID
}

type LinkIDs struct {
	LinkDelinkFunction LinkDelink

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Link.
	LinkPrimaryID PrimaryID
}

type DelinkIDs struct {
	LinkDelinkFunction LinkDelink

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Delink.
	DelinkPrimaryID PrimaryID

}
