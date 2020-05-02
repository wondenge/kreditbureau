package design

import (
	. "goa.design/goa/v3/dsl"
)

var BouncedCheque = Type("bouncedcheque", func() {

	Description("Bounced Cheque API")
	TypeName("BouncedCheque")
	ContentType("application/json")

	// The Type of Drawer. It has the 2 options.
	// A – Individual Credit Consumer
	// B – Non-Individual Credit Consumer
	//
	// Mandatory Field
	// Lookup
	// Alphanumeric
	// Not more than 1 characters
	Attribute("ClientType", String, "Client Type", func() {
		Meta("rpc:tag", "1")
	})

	// The Family Name or Surname.
	//
	// Conditional to field 4.5.1
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "2")
	})

	// The First Name
	//
	// Conditional to field 4.5.1
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Forename1", String, "The First Name", func() {
		Meta("rpc:tag", "3")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Forename2", String, "The Given Name", func() {
		Meta("rpc:tag", "4")
	})

	// Other Name or Initials.
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space
	Attribute("Forename3", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "5")
	})

	// Name of the Company/Corporate entity that is a Guarantor to the Subject Account.
	//
	// Alphanumeric
	// Not more than 100 Characters
	// Conditional to field 4.5.1
	Attribute("CompanyName", String, "Company Name", func() {
		Meta("rpc:tag", "6")
	})

	// The Type of Primary Identification document Provided on Opening the Account.
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	//
	// Mandatory Field
	// Data is Alphanumeric
	// Not more than 3 characters
	// Conditional to field 4.5.1
	Attribute("PrimaryIDocument Type", String, "Primary Identification Document Type", func() {
		Meta("rpc:tag", "7")
	})

	// The Number of the Primary Identification Document Provided on Opening the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocumentNumber", String, "Primary Identification Document Number", func() {
		Meta("rpc:tag", "8")
	})

	// Bank and Branch code on Bounced Cheque.
	//
	// The format of the branch code is IXXXYYY
	// Where
	// I – the Institution type code
	// B for Banks
	// D For MFBs
	// S For Saccos
	// M for MFIs
	// L for Leasing Companies
	// XXX is the Lenders Institution Code left padded with Zeros
	// e.g. 098 for a bank whose code is 98.
	// YYY is the Lenders Branch Code left padded with Zeros
	// e.g. 009 for a branch whose code is 9.
	Attribute("BranchCodeOncheque", String, "Branch Code on cheque", func() {
		Meta("rpc:tag", "9")
	})

	// Client Reference Code Linking Customer to Banking System
	//
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "10")
	})

	// Bank Account on Bounced Cheque
	//
	// Mandatory
	// Alphanumeric
	// Not more than 20 Characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "11")
	})

	// Type of Cheque. Options Are :
	// A - Corporate
	// B - Personal
	// C - Dividend
	//
	// Alphanumeric
	// Lookup
	// Not more than 1 Character
	Attribute("ChequeAccountType", String, "Cheque Account Type", func() {
		Meta("rpc:tag", "12")
	})

	// Amount on cheque
	// Mandatory
	// Amount Field
	// Not more than 16 characters
	Attribute("ChequeAmount", String, "Cheque Amount", func() {
		Meta("rpc:tag", "13")
	})

	// Cheque Number
	// Mandatory
	// Number (no spaces)
	// Not more than 9 characters
	Attribute("ChequeNumber", String, "Cheque Number", func() {
		Meta("rpc:tag", "14")
	})

	// ISO Currency Code
	//
	// Mandatory
	// Alphanumeric Field
	// Not more than 3 characters
	Attribute("ChequeCurrency", String, "Cheque Currency", func() {
		Meta("rpc:tag", "15")
	})

	// Date on Bounced Cheque
	//
	// Mandatory
	// Date format
	// Cannot be a future date
	Attribute("ChequeDate", String, "Cheque Date", func() {
		Meta("rpc:tag", "16")
	})

	// The Date the Cheque was unpaid by the paying Bank.
	//
	// Mandatory Field
	// Date format
	// Cannot be in the future
	Attribute("ChequeBounceDate", String, "Cheque Bounce Date", func() {
		Meta("rpc:tag", "17")
	})

	// Reason Code for Bouncing cheque Use options below:
	// 30 – Cheque Unpaid Because of Suspected Criminal Activity
	// 31 – Date Expired/Cheque Stale
	// 32 – Cheque Post-dated
	// 33 – Date Irregular
	// 37 – Payee Name Incomplete
	// 38 – Payee Name Irregular
	// 40 – Amounts in Words and Figures Differ
	// 41 – Amount in Words required
	// 42 – Amount in Figures required
	// 43 – Amount in Figures Irregular/incomplete
	// 53 – Signature Differs
	// 54 – Drawer Signature Required
	// 55 – Not signed in accordance with Mandate
	// 56 – Endorsement Irregular
	// 57 – Alteration, requires drawers signature
	// 58 – Payee Name Required
	// 62 – Effects not Cleared
	// 63 – Insufficient Funds – Refer to Drawer
	// 64 – Account Dormant – Refer to Drawer
	// 66 – Cheque Re-presented more than Once
	// 69 – Invalid Account Number
	// 70 – Title of Account Irregular
	// 71 – Originator Reference Not Quoted
	// 72 – Account Transferred
	// 73 – Customer/Drawer Deceased
	// 74 – Account Closed
	// 75 – Title of Account Required
	// 76 – Debit in Excess of Direct Debit Authority
	// 77 – Frozen Account
	// 79 – Payment Stopped by Drawer
	// 80 – Confirmation Awaited
	//
	// Alphanumeric
	// Mandatory Field
	// Lookup Value
	// Not more than 2 characters
	Attribute("ChequeBounceReasonCodes", String, "Cheque Bounce Reason Codes", func() {
		Meta("rpc:tag", "18")
	})

	// The Loan Account to which the Cheque was deposited to.
	//
	// Mandatory
	// Alphanumeric
	// Not more than 20 characters
	Attribute("LoanAccount", String, "Loan Account", func() {
		Meta("rpc:tag", "19")
	})

	// The amount on the Cheque in Kenya Shillings.
	// If the Cheque is drawn in foreign currency, use the Exchange Rate on the Reporting Date.
	// If the Cheque is drawn in Kenya Shillings, report the actual amount on the cheque.
	Attribute("ChequeAmountInKES", String, "Cheque Amount in KES", func() {
		Meta("rpc:tag", "20")
	})

})
