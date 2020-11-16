package kreditbureau

type IncomeBand string

const (
	KES0To50000       IncomeBand = "A – KES 0/- To 50,000/-"
	KES50001To100000             = "B – KES 50,001/- To 100,000/-"
	KES10001To200000             = "C - KES 100,001/- To 200,000/-"
	KES200001To500000            = "D - KES 200,001/- To 500,000/-"
	OverKES500000                = "E - Over KES 500,000/-"
)
