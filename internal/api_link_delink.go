package internal

type LinkDelink string

const (
	Link   LinkDelink = "001 - Link"
	Delink            = "002 – Delink"
)

type LinkDelinkID struct {
	LinkDelinkFunction LinkDelink

	// Primary Identification Document provided on Opening the Account.
	// Primary ID to Link.
	LinkPrimaryID PrimaryID
}
