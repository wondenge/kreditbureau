package internal

type GuarantorsRelationship string

const (
	Family                GuarantorsRelationship = "A = Family" // (e.g. Husband, Wife, Daughter, etc.)
	ShareholderOrPartner                         = "B = Shareholder, Partner, etc."
	FriendOrWorkColleague                        = "C = Friend/Work Colleague"
)

type NewGuarantor struct {
	Lenders                Lenders
	ClientNumber           ClientNumber
	AccountNumber          AccountNumber
	CustomerName           CustomerName
	CompanyName            CompanyName
	ClientType             ClientType
	Nationality            Nationality
	MaritalStatus          MaritalStatus
	GuaranteeType          GuaranteeType
	GuarantorsRelationship GuarantorsRelationship
	GuaranteeLimit         GuaranteeLimit
	UserID                 UserID
	Employer               Employer
	TelephoneNumber        TelephoneNumber
	PostalAddress          PostalAddress
	PhysicalAddress        PhysicalAddress
	EmailAddress           EmailAddress
}
