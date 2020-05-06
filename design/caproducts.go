package design

import (
	. "goa.design/goa/v3/dsl"
)

// Credit Application Product Types

// A – Overdraft
var CAOverdraft = Type("CAOverdraft", func() {
	Description("Product Type A: Overdraft")

	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// B – Credit Cards
var CACreditCards = Type("CACreditCards", func() {
	Description("Product Type B: Credit Cards")

	// Not Mandatory for Credit Cards
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

})

// C – Business Working Capital Loans
var CABusinessWorkingCapitalLoans = Type("CABusinessWorkingCapitalLoans", func() {
	Description("Product Type C: Business Working Capital Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// D – Business Expansion Loans
var CABusinessExpansionLoans = Type("CABusinessExpansionLoans", func() {
	Description("Product Type D: Business Expansion Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")
})

// E – Mortgage Loans
var CAMortgageLoans = Type("CAMortgageLoans", func() {
	Description("Product Type E: Mortgage Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// F – Asset Finance Loans
var CAAssetFinanceLoans = Type("CAAssetFinanceLoans", func() {
	Description("Product Type F: Asset Finance Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// G – Trade Finance Facilities
var CATradeFinance = Type("CATradeFinance", func() {
	Description("Product Type G: Trade Finance Facilities")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// H – Personal Loans
var CAPersonalLoans = Type("CAPersonalLoans", func() {
	Description("Product Type H: Personal Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// I – Mobile Loan
var CAMobileLoan = Type("CAMobileLoan", func() {
	Description("Product Type I: Mobile Loan")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// J – Insurance Premium Financing
var CAInsurancePremiumFinancing = Type("CAInsurancePremiumFinancing", func() {
	Description("Product Type J: Insurance Premium Financing")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// K – Group Loans
var CAGroupLoans = Type("CAGroupLoans", func() {
	Description("Product Type K: Group Loans")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})

// L – Uncleared Effects
var CAUnclearedEffects = Type("CAUnclearedEffects", func() {
	Description("Product Type L: Uncleared Effects")
	Attribute("LoanTerm", Int, func() {
		Description("Term of Loan") // Number of Months Applied for.
		MinLength(1)
		MaxLength(3)
		Minimum(1)
		Meta("rpc:tag", "27")
	})

	Required("LoanTerm")

})
