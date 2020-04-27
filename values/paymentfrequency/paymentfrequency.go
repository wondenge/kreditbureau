package paymentfrequency

type (
	PaymentFrequency string
)

const (
	W PaymentFrequency = "Weekly"
	M                  = "Monthly"
	Q                  = "Quarterly"
	S                  = "Semi - Annually"
	Y                  = "Annually(Yearly)"
	B                  = "Bullet"
	T                  = "Tri - annual(Three times in a year)"
	C                  = "Biennial (Once in two years)"
	D                  = "Daily"
	F                  = "Fortnightly"
)
