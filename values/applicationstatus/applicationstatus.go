package applicationstatus

type (
	ApplicationStatus string
)

const (
	A ApplicationStatus = "Pending"
	B                   = "Awaiting Docs"
	C                   = "Securities Perfection"
	D                   = "Declined by the Bank"
	E                   = "Withdrawn"
	F                   = "Approved"
	G                   = "Pending disbursement"
	H                   = "Fully Disbursed"
	I                   = "Customer Declines Offer"
)
