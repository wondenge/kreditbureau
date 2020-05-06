package design

import (
	. "goa.design/goa/v3/dsl"
)

var BouncedCheque = Type("bouncedcheque", func() {

	Description("Bounced Cheque API")
	TypeName("BouncedCheque")
	ContentType("application/json")

	// BouncedChequeCTA – Individual Credit Consumer
	// BouncedChequeCTB – Non-Individual Credit Consumer
	Attribute("ClientType", ArrayOf(Enum), func() {
		Description("The Type of Drawer")
		Enum(BouncedChequeCTA, BouncedChequeCTB)
		Meta("rpc:tag", "1")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocument Type", String, func() {
		Description("Primary Identification Document Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "7")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocumentNumber", String, func() {
		Description("Primary Identification Document Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "8")
	})

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
		Description("Bank and Branch code on Bounced Cheque")
		Meta("rpc:tag", "9")
	})

	// Client Reference Code Linking Customer to Banking System
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "10")
	})

	Attribute("AccountNumber", String, "Account Number", func() {
		Description("Bank Account on Bounced Cheque")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "11")
	})

	// A - Corporate
	// B - Personal
	// C - Dividend
	Attribute("ChequeAccountType", String, func() {
		Description("Cheque Account Type")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("B")
		Meta("rpc:tag", "12")
	})

	Attribute("ChequeAmount", Int64, func() {
		Description("Cheque Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "13")
	})

	Attribute("ChequeNumber", Int, func() {
		Description("Cheque Number") // Number (no spaces)
		MinLength(1)
		MaxLength(9)
		Meta("rpc:tag", "14")
	})

	Attribute("ChequeCurrency", String, func() {
		Description("ISO Currency Code")
		MinLength(3)
		MaxLength(3)
		Default("KES")
		Meta("rpc:tag", "15")
	})

	Attribute("ChequeDate", String, "Cheque Date", func() {
		Description(" Date on Bounced Cheque")
		Format(FormatDate)
		Meta("rpc:tag", "16")
	})

	Attribute("ChequeBounceDate", String, func() {
		Description("Date Cheque was unpaid by the paying Bank")
		Format(FormatDate)
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
	Attribute("ChequeBounceReasonCodes", Int, func() {
		Description("Cheque Bounce Reason Codes")
		MinLength(2)
		MaxLength(2)
		Minimum(30)
		Maximum(80)
		Meta("rpc:tag", "18")
	})

	Attribute("LoanAccount", String, func() {
		Description("Loan Account to which Cheque was deposited")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "19")
	})

	// If the Cheque is drawn in foreign currency, use the Exchange
	// Rate on the Reporting Date.
	// If the Cheque is drawn in Kenya Shillings, report the actual
	// amount on the cheque.
	Attribute("ChequeAmountInKES", Int64, func() {
		Description("Cheque Amount in KES")
		MinLength(1)
		Minimum(0)
		Meta("rpc:tag", "20")
	})

})

// A – Individual Credit Consumer
// B – Non-Individual Credit Consumer
var StoredBouncedCheque = ResultType("application/vnd.goa.bouncedcheque", func() {
	TypeName("StoredBouncedCheque")
	Attributes(func() {
		Extend(BouncedCheque)
		Required(
			"ClientType",
			"PrimaryIDocument Type",
			"PrimaryIDocumentNumber",
			"BranchCodeOncheque",
			"AccountNumber",
			"ChequeAccountType",
			"ChequeAmount",
			"ChequeNumber",
			"ChequeCurrency",
			"ChequeDate",
			"ChequeBounceDate",
			"ChequeBounceReasonCodes",
			"LoanAccount",
			"ChequeAmountInKES",
		)
	})

	View("default", func() {
		Attribute("ClientType")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyName")
		Attribute("PrimaryIDocument Type")
		Attribute("PrimaryIDocumentNumber")
		Attribute("BranchCodeOncheque")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("ChequeAccountType")
		Attribute("ChequeAmount")
		Attribute("ChequeNumber")
		Attribute("ChequeCurrency")
		Attribute("ChequeDate")
		Attribute("ChequeBounceDate")
		Attribute("ChequeBounceReasonCodes")
		Attribute("LoanAccount")
		Attribute("ChequeAmountInKES")
	})

	View("mandatory", func() {
		Attribute("ClientType")
		Attribute("PrimaryIDocument Type")
		Attribute("PrimaryIDocumentNumber")
		Attribute("BranchCodeOncheque")
		Attribute("AccountNumber")
		Attribute("ChequeAccountType")
		Attribute("ChequeAmount")
		Attribute("ChequeNumber")
		Attribute("ChequeCurrency")
		Attribute("ChequeDate")
		Attribute("ChequeBounceDate")
		Attribute("ChequeBounceReasonCodes")
		Attribute("LoanAccount")
		Attribute("ChequeAmountInKES")
	})

})

var BouncedChequeCTA = Type("BouncedChequeCTA", func() {
	Description("A – Individual Credit Consumer")

	Attribute("Surname", String, func() {
		Description("Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "5")
	})

	Required("Surname", "Forename1", "Forename2", "Forename3")

})

var BouncedChequeCTB = Type("BouncedChequeCTB", func() {
	Description("B – NonIndividual Credit Consumer")

	// Name of the Company/Corporate entity that is a Guarantor
	// to the Subject Account.
	Attribute("CompanyName", String, func() {
		Description("Name of the Company/Corporate entity")
		MinLength(1)
		Maximum(100)
		Meta("rpc:tag", "6")
	})

	Required("CompanyName")

})
