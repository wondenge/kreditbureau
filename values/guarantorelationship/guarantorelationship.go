package guarantorelationship

type (
	GuarantorRelationship string
)

const (
	A GuarantorRelationship = "Family (e.g. Husband, Wife, Daughter, etc.)"

	// If Guarantee Type = “B”
	B = "Shareholder, Partner, etc."

	// If Guarantee Type = “A” or “C”
	C = "Friend/Work Colleague"
)
