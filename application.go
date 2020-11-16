package kreditbureau

type ApplicationStatus string
type ApplicationDeclineReason string
type StatusUpdateReason string
type UpdateFunction string

const (
	Pending               ApplicationStatus = "A - Pending"
	AwaitingDocs                            = "B - Awaiting Docs"
	SecuritiesPerfection                    = "C - Securities Perfection"
	DeclinedByBank                          = "D - Declined by the Bank"
	Withdrawn                               = "E - Withdrawn"
	Approved                                = "F - Approved"
	PendingDisbursement                     = "G - Pending disbursement"
	FullyDisbursed                          = "H – Fully Disbursed"
	CustomerDeclinesOffer                   = "I – Customer Declines Offer"
)

const (
	OverIndebted           ApplicationDeclineReason = "A= Over indebted"
	FailedCreditCriteria                            = "B = Failed credit criteria"
	FailedVerification                              = "C= Failed verification" // e.g.Income/Employer could not be verified
	LacksAbilityToRepay                             = "D = Lacks ability to repay"
	WeakCreditHistory                               = "E = Weak Credit History"
	InsufficientCollateral                          = "F = Insufficient Collateral"
)

const (
	AdditionalIncomeProvided  StatusUpdateReason = "A = Additional Income provided"
	CreditProfileUpdated                         = "B = Credit Profile Updated " // e.g. if client presents proof that an account is paid up or status on credit bureau is out dated
	AdditionalDepositByClient                    = "C= Additional deposit provided by client"
)

const (
	MergeOneAccountToAnother UpdateFunction = "A – Merge one Account to Another"
	AddAccountToProfile                     = "B – Add an Account to a Profile"
	RemoveAccountFromProfile                = "C – Remove an Account from a Profile"
	MergeIDs                                = "D – Merge Identification Documents"
	UnmergeIDs                              = "E – Unmerge Identification Documents"
	RemoveNameFromProfile                   = "F – Remove a Name from a profile"
	ExpungeAccount                          = "G – Expunge an Account"
	UnlinkGuarantor                         = "H – Unlink a Guarantor"
	UnlinkStakeholder                       = "I – Unlink a Stakeholder"
)
