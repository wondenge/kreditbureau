package employmentype

type (
	EmploymentType int
)

const (
	Casual EmploymentType = iota
	Contract
	Permanent
	Pensioner
	SelfEmployed
)
