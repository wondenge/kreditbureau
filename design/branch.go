package design

import (
	. "goa.design/goa/v3/dsl"
)

// B. Institution Branch Information
//
// The Institution Branches are defined in the following table.
// All the branches of the institution are required to be reported.

var Branch = Type("branch", func() {
	Description("")
	TypeName("Branch")
	ContentType("application/json")

	// The Bank and Branch Code as allocated by the KBA Secretariat.
	// Also Known as the Branch Sort Code
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Branch Name
	Attribute("BranchName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Date the Branch Opened for business.
	// Date should not be in the future.
	Attribute("DateOpened", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Branch email Address
	// Same as Parent Bank email Address if none provided.
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Website link, if different from the Main institution website.
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Number of Active Accounts in the Branch.
	// Definition of active accounts as per the Central Bank.
	Attribute("NrOfActiveAccounts", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Total Value of the Active Accounts in the branch.
	// The Number of closed accounts in the branch since the last report.
	Attribute("ActiveAccountsValue", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Number of closed accounts in the branch since the last report.
	Attribute("NrOfClosedAccounts", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// The Number of New accounts opened in the branch since the last report.
	Attribute("NrOfNewAccounts", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Total Value of the New Accounts opened in the branch since the last report.
	Attribute("NewAccountsValue", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Number of Non-Performing Loan Accounts in the branch since the last report.
	// As defined by the Central Bank under section 1.4.2 of the prudential guidelines
	// no CBK/PG/04 on risk classification of Assets and provisioning.
	// The guideline provides a definition for both loans and overdrafts.
	Attribute("NrNonPerformingLoans", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Total Value of the Non-Performing Loan accounts in the branch since the last report.
	Attribute("NonPerformingLoanValue", String, "", func() {
		Meta("rpc:tag", "12")
	})
})
