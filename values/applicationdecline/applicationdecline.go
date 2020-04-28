package applicationdecline

// Reason for Application Decline
type (
	ApplicationDecline string
)

const (
	A ApplicationDecline = "Over indebted"
	B                    = "Failed credit criteria"
	C                    = "Failed verification – (e.g.Income/Employer could not be verified)"
	D                    = "Lacks ability to repay"
	E                    = "Weak Credit History"
	F                    = "Insufficient Collateral"
)
