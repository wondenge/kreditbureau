package design

import (
	. "goa.design/goa/v3/dsl"
)

// 001=Casual
var CasualEmployment = Type("CasualEmployment", func() {
	Description("Employment Type 001: Casual")

	// If the consumer is employed, the Employer Name.
	// The field is conditional based on the Employment Type Field.
	// It is not mandatory for Self-Employed and Pensioners.
	//
	// Alphanumeric Field
	// Not more than 50 characters
	Attribute("EmployerName", String, "Employer Name", func() {
		Meta("rpc:tag", "36")
	})

	// If the Consumer is employed, the Date Consumer Was employed.
	//
	// Date Field
	// Should not be in the future
	// Not more than 8 characters
	Attribute("EmploymentDate", String, "Employment Date", func() {
		Meta("rpc:tag", "38")
	})

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	Required("EmployerName", "EmploymentDate", "IncomeAmount")
})

// 002= Contract
var ContractEmployment = Type("ContractEmployment", func() {
	Description("Employment Type 002: Contract ")

	// If the consumer is employed, the Employer Name.
	// The field is conditional based on the Employment Type Field.
	// It is not mandatory for Self-Employed and Pensioners.
	//
	// Alphanumeric Field
	// Not more than 50 characters
	Attribute("EmployerName", String, "Employer Name", func() {
		Meta("rpc:tag", "36")
	})

	// If the Consumer is employed, the Date Consumer Was employed.
	//
	// Date Field
	// Should not be in the future
	// Not more than 8 characters
	Attribute("EmploymentDate", String, "Employment Date", func() {
		Meta("rpc:tag", "38")
	})

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	Required("EmployerName", "EmploymentDate", "IncomeAmount")
})

// 003= Permanent
var PermanentEmployment = Type("PermanentEmployment", func() {
	Description("Employment Type 003: Permanent ")

	// If the consumer is employed, the Employer Name.
	// The field is conditional based on the Employment Type Field.
	// It is not mandatory for Self-Employed and Pensioners.
	//
	// Alphanumeric Field
	// Not more than 50 characters
	Attribute("EmployerName", String, "Employer Name", func() {
		Meta("rpc:tag", "36")
	})

	// If the Consumer is employed, the Date Consumer Was employed.
	//
	// Date Field
	// Should not be in the future
	// Not more than 8 characters
	Attribute("EmploymentDate", String, "Employment Date", func() {
		Meta("rpc:tag", "38")
	})

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	Required("EmployerName", "EmploymentDate", "IncomeAmount")
})

// 004=Pensioner
var PensionedEmployment = Type("PensionedEmployment", func() {
	Description("Employment Type 004: Pensioner")

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	Required("IncomeAmount")
})

// 005= Self-Employed
var SelfEmployment = Type("SelfEmployment", func() {
	Description("Employment Type 005: Self-Employed")

	// The Borrower’s Average Monthly Income, in Kenya Shillings.
	//
	// Mandatory Field
	// Amount Field
	// Not more than 16 characters
	Attribute("IncomeAmount", String, "Income Amount", func() {
		Meta("rpc:tag", "40")
	})

	Required("IncomeAmount")
})
