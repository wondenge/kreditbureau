package kreditbureau

type DeletionRelistFunctionType string

const (
	Complete            DeletionRelistFunctionType = "A-Complete"
	DefaultHistory                                 = "B-Default History"
	SpecificMonthUpdate                            = "C-Specific Month Update"
	Relist                                         = "D- Relist"
)
