package design

import (
	. "goa.design/goa/v3/dsl"
)

// I. Bounced Cheque Information

var BouncedCheque = Type("bouncedcheque", func() {

	Description("")
	TypeName("BouncedCheque")
	ContentType("application/json")

	// Bank and Branch code on Cheque – the drawer’s Bank and Branch Code
	Attribute("BranchCodeOnCheque", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// Customer Id Code, if used by the bank.
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Drawer’s Bank Account number as reflected on bounced Cheque leaf
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Type of Cheque.
	// Options are :
	// A - Corporate Cheque
	// B - Personal cheque
	// C - Dividend cheque
	Attribute("ChequeAccountType", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Amount on cheque
	Attribute("ChequeAmount", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Cheque Serial Number
	Attribute("ChequeNumber", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// ISO Currency Code of Cheque
	Attribute("ChequeCurrency", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Cheque Issue Date
	Attribute("ChequeDate", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Beneficiary or Payee Name on the Cheque
	Attribute("BeneficiaryNameOrPayee", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Date the cheque was unpaid by the paying Bank.
	Attribute("ChequeBounceDate", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Account status of the Account.
	// Options are:
	// A - Closed
	// B - Dormant
	// C - Performing
	// D - Non-Performing
	// E - Write-Off
	// F - Legal
	// G - Collection
	Attribute("AccountStatus", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// KBA Reason Code for Bouncing as provided by the KBA.
	Attribute("ChequeBounceReason", String, "", func() {
		Meta("rpc:tag", "12")
	})
})
