package kreditbureau

type OrganizationType string
type OrganizationClass string

const (
	PublicLimitedCompany  OrganizationType = "A – Public Limited company"
	PrivateLimitedCompany                  = "B - Private Limited company"
	Partnership                            = "C - Partnership"
	SoleProprietor                         = "D - Sole Proprietor"
	Society                                = "E – Society"
	ClubOrAssociation                      = "F - Club/Association"
	InvestmentGroup                        = "G – Investment Group"
	PublicCorporation                      = "H - Public Corporation"
	NGO                                    = "J - NGO"
	Church                                 = "L - Church"
	School                                 = "M - School"
)

const (
	Corporate            OrganizationClass = "A – Corporate"
	SME                                    = "B – SME"
	NGOOrChurchOrSociety                   = "C - NGO/Church/Society"
)
