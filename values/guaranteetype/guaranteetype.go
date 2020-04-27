package guaranteetype

type (
	GuaranteeType string
)

const (
	A GuaranteeType = "Directors guarantee"
	B               = "Personal guarantee"
	C               = "Corporate guarantee"
	D               = "Bank Guarantee"
	F               = "Personal Guarantee with Collateral"
)
