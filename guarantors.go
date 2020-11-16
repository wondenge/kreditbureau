package kreditbureau

type GuarantorsRelationship string

const (
	Family                GuarantorsRelationship = "A = Family" // (e.g. Husband, Wife, Daughter, etc.)
	ShareholderOrPartner                         = "B = Shareholder, Partner, etc."
	FriendOrWorkColleague                        = "C = Friend/Work Colleague"
)
