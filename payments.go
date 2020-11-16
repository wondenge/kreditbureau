package kreditbureau

type PaymentFrequency string

const (
	Weekly       PaymentFrequency = "W - Weekly"
	Monthly                       = "M - Monthly"
	Quarterly                     = "Q – Quarterly"
	SemiAnnually                  = "S – Semi-Annually"
	Annually                      = "Y – Annually ( Yearly)"
	Bullet                        = "B – Bullet"
	TriAnnual                     = "T – Tri-annual (Three times in a year)"
	Biennial                      = "C – Biennial (Once in two years)"
	Daily                         = "D – Daily"
	Fortnightly                   = "F - Fortnightly"
)
