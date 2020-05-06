package design

import (
	. "goa.design/goa/v3/dsl"
)

// Company Product Types

// A – Overdraft
var CompanyOverdraft = Type("CompanyOverdraft", func() {
	Description("Product Type A: Overdraft")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans") // Optional for Overdrafts.
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
			Format(FormatDate)
		})
		Required("LatestPaymentDate", "LastPaymentAmount")
	})

	Required("LastPaymentAmount")

})

// B – Credit Cards
var CompanyCreditCards = Type("CompanyCreditCards", func() {
	Description("Product Type B: Credit Cards")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans") // Optional for Credit Cards.
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("LastPaymentAmount")
})

// C – Business Working Capital Loans
var CompanyBusinessWorkingCapitalLoans = Type("CompanyBusinessWorkingCapitalLoans", func() {
	Description("Product Type C: Business Working Capital Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// D – Business Expansion Loans
var CompanyBusinessExpansionLoans = Type("CompanyBusinessExpansionLoans", func() {
	Description("Product Type D: Business Expansion Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// E – Mortgage Loans
var CompanyMortgageLoans = Type("CompanyMortgageLoans", func() {
	Description("Product Type E: Mortgage Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// F – Asset Finance Loans
var CompanyAssetFinanceLoans = Type("CompanyAssetFinanceLoans", func() {
	Description("Product Type F: Asset Finance Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// G – Trade Finance Facilities
var CompanyTradeFinance = Type("CompanyTradeFinance", func() {
	Description("Product Type G: Trade Finance Facilities")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans") // Optional Trade Finance facility.
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})
})

// H – Personal Loans
var CompanyPersonalLoans = Type("CompanyPersonalLoans", func() {
	Description("Product Type H: Personal Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// I – Mobile Loan
var CompanyMobileLoan = Type("CompanyMobileLoan", func() {
	Description("Product Type I: Mobile Loan")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// J – Insurance Premium Financing
var CompanyInsurancePremiumFinancing = Type("CompanyInsurancePremiumFinancing", func() {
	Description("Product Type J: Insurance Premium Financing")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// K – Group Loans
var CompanyGroupLoans = Type("CompanyGroupLoans", func() {
	Description("Product Type K: Group Loans")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})

// L – Uncleared Effects
var CompanyUnclearedEffects = Type("CompanyUnclearedEffects", func() {
	Description("Product Type L: Uncleared Effects")
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Minimum(0) // If the account is not a final status this field is not zero
		Meta("rpc:tag", "56")
	})
	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "58")
		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			Format(FormatDate)
			Format(FormatDate)
			Meta("rpc:tag", "57")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")
})
