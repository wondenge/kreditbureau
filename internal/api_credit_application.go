package internal

type ApplicationStatus string
type ApplicationDeclineReason string
type ApplicationStatusUpdateReason string
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
	AdditionalIncomeProvided  ApplicationStatusUpdateReason = "A = Additional Income provided"
	CreditProfileUpdated                                    = "B = Credit Profile Updated " // e.g. if client presents proof that an account is paid up or status on credit bureau is out dated
	AdditionalDepositByClient                               = "C= Additional deposit provided by client"
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

type Facility struct {
	Lenders           Lenders
	ClientType        ClientType
	CustomerName      CustomerName
	Registrar         Registrar
	UserID            UserID
	ClientNumber      ClientNumber
	PersonalPINNumber PersonalPINNumber
	AccountNumber     AccountNumber

	// The Date the Application was made to the Lender
	ApplicationDate string

	// Internal Application Reference number
	ApplicationNumber       string
	FacilityApplicationType SecurityType

	// Loan Amount applied for.
	ApplicationAmount float64

	// Currency in which facility is requested for.
	// Default is KES
	ApplicationCurrency CurrencyCode

	// The Account Product Type.
	ProductType AccountProductType

	// Number of Months Applied for.
	TermOfLoan int

	// The Current status of the Application.
	ApplicationStatus        ApplicationStatus
	ApplicationDeclineReason ApplicationDeclineReason

	// Date of status of Update
	ApplicationStatusDate         string
	ApplicationStatusUpdateReason ApplicationStatusUpdateReason
}
