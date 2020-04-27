package accountproductype

type (
	AccountProductType string
)

// Option B,H, I, K not applicable to the Non-Individual Consumer File.
const (
	A AccountProductType = "Overdraft"
	B                    = "Credit Cards"
	C                    = "Business Working Capital Loans"
	D                    = "Business Expansion Loans"
	E                    = "Mortgage Loans"
	F                    = "Asset Finance Loans"
	G                    = "Trade Finance Facilities"
	H                    = "Personal Loans"
	I                    = "Mobile Loan"
	J                    = "Insurance Premium Financing"
	K                    = "Group Loans"
	L                    = "Uncleared Effects"
)
