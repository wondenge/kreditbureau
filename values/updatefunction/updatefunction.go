package updatefunction

type (
	UpdateFunction string
)

const (
	A UpdateFunction = "Merge one Account to Another"
	B                = "Add an Account to a Profile"
	C                = "Remove an Account from a Profile"
	D                = "Merge Identification Documents"
	E                = "Unmerge Identification Documents"
	F                = "Remove a Name from a profile"
	G                = "Expunge an Account"
	H                = "Unlink a Guarantor"
	I                = "Unlink a Stakeholder"
)
