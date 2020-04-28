package statusupdate

// Status Update Reason Code
type (
	StatusUpdate string
)

const (
	A StatusUpdate = "Additional Income provided"
	// (e.g. if client presents proof that an account is paid up or status on credit bureau is out dated)
	B = "Credit Profile Updated "
	C = "Additional deposit provided by client"
)
