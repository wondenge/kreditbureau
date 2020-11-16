package kreditbureau

type GuaranteeType string
type GroupGuaranteeMemberStatus string

const (
	DirectorsGuarantee              GuaranteeType = "A - Directors guarantee"
	PersonalGuarantee                             = "B - Personal guarantee"
	CorporateGuarantee                            = "C - Corporate guarantee"
	BankGuarantee                                 = "D - Bank Guarantee"
	PersonalGuaranteeWithCollateral               = "F – Personal Guarantee with Collateral"
)

const (
	ActiveGroupMember  GroupGuaranteeMemberStatus = "A – Active Member of the Group"
	DormantGroupMember                            = "B – Dormant Member of the Group"
	ExitedGroup                                   = "C – Exited the Group"
)
