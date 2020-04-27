package accountstatus

type (
	AccountStatus string
)

// 	For Field 4.5.14, Options A, B, C, D, E and F are the only acceptable options
const (
	// No more admin processes running such
	// as instalment demands or interest charges
	// to account, and no further facilities can
	// be offered on this account.
	A AccountStatus = "Closed"

	// No activity for 2 years. This applies for
	// Overdraft/Current Accounts.
	B = "Dormant"

	// For facilities that don’t form part of the
	// outstanding portfolio in the Balance Sheet,
	// but are still outstanding in the books of
	// accounts.
	C = "Write-Off"

	// With legal officer in court
	D = "Legal"

	// With collection bureau
	E = "Collection"

	// For facilities that form part of the outstanding
	// portfolio, and are  reported in the Balance Sheet.
	F = "Active"

	// For Rescheduled/Restructured Facilities
	G = "Facility Rescheduled"

	// The facility has been cleared. This status can
	// only be used for revolving facilities.
	H = "Settled"

	// The facility has been called up.Once the client
	// has paid up, the status should be updated to
	// Option A, H; or otherwise Option L
	J = "Called Up"

	// The facility has been put on hold for an indefinite
	// period of time
	K = "Suspended"

	L = "Client Deceased"

	// This refers to facilities whose payments have been
	// put on hold for a definite period or in moratorium
	//(Grace Period)
	M = "Deferred"

	// This status is reserved for CRBs (if last record
	// status is not CLOSED)
	N = "Not Updated"

	// Refers to a Record that the Client has disputed at
	// the CRB. The account cannot be updated until the
	// dispute is resolved
	P = "Disputed"
)
