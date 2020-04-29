package design

import (
	. "goa.design/goa/v3/dsl"
)

// A. Institution Information
//
// The Institution refers to the Participating institution
// (Currently Commercial Banks). Most of the information in
// this table is static information pertaining to the bank
// and its registration details.
var Institution = Type("institution", func() {
	Description("")
	TypeName("Institution")
	ContentType("application/json")

	// The Name of the Institution as registered with the
	// Registrar of companies
	Attribute("RegisteredName", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Business or Trading Name of the institution,
	// if different from the Registered Name. Same as The
	// Registered Name as registered with the Registrar
	// of companies if not provided.
	Attribute("TradingName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Date of Registration with Registrar of Companies.
	// Date Format is DDMMYYYY
	Attribute("RegistrationDate", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Registration Certificate Number issued by the
	// Registrar of Companies.
	Attribute("RegistrationNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Date the Institution was issued with a banking
	// licence by the Central Bank of Kenya or the Date
	// of Renewal of the Banking License. This date Cannot
	// be in the future
	Attribute("LicenseIssueDate", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Applicable Provider Type.
	// Currently Banking Services. Will extend to other
	// industries once we open the CRB submissions to
	// other provider classes.
	Attribute("ProviderType", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// The Kenya Revenue Authority Income Tax PIN Number.
	// May be Mandatory in Future.
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Institution’s Email Address.
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// The Institution’s Web site url.
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// All Credit facilities with debit balances that arise
	// from business transactions. Defined by the Central Bank
	// i.e. Movements in the account are not merely the
	// result of charges levied.
	Attribute("NrOfActiveAccounts", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// The Total Active bank accounts in the institution
	Attribute("ActiveAccountsValue", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// The Total number of Closed accounts in the institution
	// in the period since the last reporting date.
	Attribute("NrOfClosedAccounts", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Total number of New accounts opened in the institution
	// in the period since the last reporting date.
	Attribute("NrOfNewAccounts", String, " ", func() {
		Meta("rpc:tag", "13")
	})

	// Total Value of New Accounts since the last reporting date.
	Attribute("NewAccountsValue", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Total Number of Non-Performing Loan Accounts since the last
	// reporting date As defined by the Central Bank under
	// section 1.4.2 of the prudential guidelines no CBK/PG/04 on
	// risk classification of Assets and provisioning. The guideline
	// provides a definition for both loans and overdrafts.
	Attribute("NrNonPerformingLoans", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Total Value of Non-Performing Loan Accounts since the Last
	// reporting date.
	Attribute("NonPerformingLoansValue", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// The Two-Digit KBA allocated Bank Code
	Attribute("BankCode", String, "", func() {
		Meta("rpc:tag", "17")
	})
	View("default", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("LicenseIssueDate")
		Attribute("ProviderType")
		Attribute("PINNumber")
		Attribute("Email")
		Attribute("Website")
		Attribute("NrOfActiveAccounts")
		Attribute("ActiveAccountsValue")
		Attribute("NrOfClosedAccounts")
		Attribute("NrOfNewAccounts")
		Attribute("NewAccountsValue")
		Attribute("NrNonPerformingLoans")
		Attribute("NonPerformingLoansValue")
		Attribute("BankCode")
	})

	Required("RegisteredName")
})

var StoredInstitution = ResultType("application/vnd.goa.institution", func() {
	Description("")
	TypeName("StoredInstitution")
	ContentType("application/json")

	Attributes(func() {
		Extend(Institution)
		Required("RegisteredName")
	})
	View("default", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("LicenseIssueDate")
		Attribute("ProviderType")
		Attribute("PINNumber")
		Attribute("Email")
		Attribute("Website")
		Attribute("NrOfActiveAccounts")
		Attribute("ActiveAccountsValue")
		Attribute("NrOfClosedAccounts")
		Attribute("NrOfNewAccounts")
		Attribute("NewAccountsValue")
		Attribute("NrNonPerformingLoans")
		Attribute("NonPerformingLoansValue")
		Attribute("BankCode")
	})
})
