package design

import (
	. "goa.design/goa/v3/dsl"
)

// Individual Product Types

// A – Overdraft
var IndividualsOverdraft = Type("IndividualsOverdraft", func() {
	Description("Product Type A: Overdraft")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// B – Credit Cards
var IndividualsCreditCards = Type("IndividualsCreditCards", func() {
	Description("Product Type B: Credit Cards")

	// Optional for  Credit Cards  facilities.
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("LastPaymentAmount")

})

// C – Business Working Capital Loans
var IndividualsBusinessWorkingCapitalLoans = Type("IndividualsBusinessWorkingCapitalLoans", func() {
	Description("Product Type C: Business Working Capital Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// D – Business Expansion Loans
var IndividualsBusinessExpansionLoans = Type("IndividualsBusinessExpansionLoans", func() {
	Description("Product Type D: Business Expansion Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// E – Mortgage Loans
var IndividualsMortgageLoans = Type("IndividualsMortgageLoans", func() {
	Description("Product Type E: Mortgage Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// F – Asset Finance Loans
var IndividualsAssetFinanceLoans = Type("IndividualsAssetFinanceLoans", func() {
	Description("Product Type F: Asset Finance Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// G – Trade Finance Facilities
var IndividualsTradeFinance = Type("IndividualsTradeFinance", func() {
	Description("Product Type G: Trade Finance Facilities")

	// Optional for Trade Finance facilities.
	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

})

// H – Personal Loans
var IndividualsPersonalLoans = Type("IndividualsPersonalLoans", func() {
	Description("Product Type H: Personal Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// I – Mobile Loan
var IndividualsMobileLoan = Type("IndividualsMobileLoan", func() {
	Description("Product Type I: Mobile Loan")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// J – Insurance Premium Financing
var IndividualsInsurancePremiumFinancing = Type("IndividualsInsurancePremiumFinancing", func() {
	Description("Product Type J: Insurance Premium Financing")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// K – Group Loans
var IndividualsGroupLoans = Type("IndividualsGroupLoans", func() {
	Description("Product Type K: Group Loans")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})

// L – Uncleared Effects
var IndividualsUnclearedEffects = Type("IndividualsUnclearedEffects", func() {
	Description("Product Type L: Uncleared Effects")

	Attribute("NextInstalmentAmount", Int64, func() {
		Description("Instalment amount of loans")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "66")
	})

	Attribute("LastPaymentAmount", Int64, func() {
		Description("Last payment received into the facility")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "68")

		Attribute("LatestPaymentDate", String, func() {
			Description("Last date when payments were received into the facility")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "67")
		})
		Required("LatestPaymentDate")
	})

	Required("NextInstalmentAmount", "LastPaymentAmount")

})
