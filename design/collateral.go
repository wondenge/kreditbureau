package design

import (
	. "goa.design/goa/v3/dsl"
)

var Collateral = Type("collateral", func() {
	Description("Collateral API")
	TypeName("Collateral")
	ContentType("application/json")

	// The Name of Lender as registered with the Registrar of companies.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersRegisteredName", String, "Lenders Registered Name", func() {
		Meta("rpc:tag", "1")
	})

	// The Lenders Trading Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersTradingName", String, "Lenders Trading Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Lenders Branch Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 characters
	Attribute("LendersBranchName", String, "Lenders Branch Name", func() {
		Meta("rpc:tag", "3")
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
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 7 characters
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "4")
	})

	// Client Reference Number Linking Borrower‘s Account in the Lenders systems.
	// This is for Client-Centric systems.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "5")
	})

	// Account Number Linking Collateral to Borrower’s Account in the Lender’s System.
	// The Account Number Is the Same as Client Number for Account Centric Systems.
	//
	// Mandatory field
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("AccountNumber", String, "Account Number", func() {
		Meta("rpc:tag", "6")
	})

	// The Borrower’s Primary Identification document Type Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 - Company Registration No
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 3 characters
	Attribute("PrimaryIDocument", String, "Primary Identification Document", func() {
		Meta("rpc:tag", "7")
	})

	// The Primary Identification document Number.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "8")
	})

	// The Borrower’s Secondary Identification document Type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 – Company Registration Number
	//
	// Data is Alphanumeric
	// Not more than 3 characters
	Attribute("SecondaryIDocument", String, "Secondary Identification Document", func() {
		Meta("rpc:tag", "9")
	})

	// Mandatory if Secondary Identification Document type is provided.
	// Conditional Field to 4.6.9
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("SecondaryIDocumentNumber", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "10")
	})

	// The Borrower’s Other Identification document type, if Provided to the Lender.
	// Options Are:
	// 001 – National ID
	// 002 – Passport
	// 003 – Alien Registration
	// 004 – Service ID
	// 005 – Company Registration Number
	//
	// Data is Alpha Numeric
	// Not more than 3 characters
	Attribute("OtherIDocument", String, "Other Identification Document", func() {
		Meta("rpc:tag", "11")
	})

	// Mandatory if Other Identification document type is provided.
	//
	// Conditional Field to 4.6.9
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("OtherIDocumentNumber", String, "Other Identification Document Number", func() {
		Meta("rpc:tag", "12")
	})

	// Revenue Authority Personal Income Tax Number.
	//
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, "PIN Number", func() {
		Meta("rpc:tag", "13")
	})

	// Applicable to Corporate Customers.
	//
	// Alphanumeric
	// Not more than 20 characters
	Attribute("CompanyVATNumber", String, "Company VAT Number", func() {
		Meta("rpc:tag", "14")
	})

	// Options should be:
	// A – Land
	// B – Property and Buildings
	// C – Cash
	// D – Treasury Bills and Bonds
	// E – Shares
	// F – Debentures
	// G – Chattels/Charges over movable assets excluding vehicles
	// H – Chattels over Vehicles
	// I – Insurance Policy
	//
	// Alphanumeric
	// Lookup
	// Not more than 1 character
	Attribute("CollateralType", String, "Collateral Type", func() {
		Meta("rpc:tag", "15")
	})

	// To match/trace collateral used on same exposures.
	// Mandatory if Collateral type is Land, Property and Buildings.
	//
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("CollateralReferenceNumber", String, "Collateral Reference Number", func() {
		Meta("rpc:tag", "16")
	})

	// Amount of the collateral as at Last valuation.
	//
	// Amount Field
	// Mandatory Field
	// Not more than 16 Characters
	Attribute("CollateralLastValuationAmount", String, "Collateral Last Valuation Amount", func() {
		Meta("rpc:tag", "17")
	})

	// Currency Code of the Collateral currency.
	// Default is Kenya Shillings.
	//
	// ISO codes
	// Alphanumeric
	// Not more than 3 characters
	Attribute("Collateral Currency", String, "Collateral Currency", func() {
		Meta("rpc:tag", "18")
	})

	// Conditional to field Collateral Type, if Land, Property and Buildings.
	//
	// Currency Field
	// Not more than 16 Characters
	Attribute("CollateralForcedSaleValue", String, "Collateral Forced Sale Value", func() {
		Meta("rpc:tag", "19")
	})

	// Date when collateral is scheduled for valuation.
	//
	// Date Field
	// Can’t be in the past
	// Not more than 8 characters
	Attribute("NextValuationDate", String, "Next Valuation Date", func() {
		Meta("rpc:tag", "20")
	})

	// Required if the Collateral has an expiry Date.
	//
	// Date Field
	// Can’t be in the past
	// Not more than 8 characters
	Attribute("CollateralExpiryDate", String, "Collateral Expiry Date", func() {
		Meta("rpc:tag", "21")
	})

	// Brief description of instruments used to recover.
	//
	// Alphanumeric
	// Not more than 50 characters
	Attribute("InstrumentOfClaimOrRecoveryType", String, "Instrument of Claim/Recovery Type", func() {
		Meta("rpc:tag", "22")
	})

	// The Date the Collateral was last Valued.
	//
	// Mandatory Field
	// Date Field
	// Can’t be in the future
	// Not more than 8 characters
	Attribute("LastValuationDate", String, "Last Valuation Date", func() {
		Meta("rpc:tag", "23")
	})

	// If the Collateral is shared for several Loans
	// Y= If Collateral is shared
	// N= If Collateral is not shared
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	Attribute("SharedCollateral", String, "Shared Collateral", func() {
		Meta("rpc:tag", "24")
	})

	// Conditional if Shared Collateral = Y
	//
	// Number Field
	// Not more than 3 characters
	Attribute("PortionOfCollateralShared", String, "Portion of Collateral Shared", func() {
		Meta("rpc:tag", "25")
	})

	// Y= If multiple collateral
	// N= If no multiple collateral
	//
	// Lookup
	// Alphanumeric
	// Not more than 1 character
	// If yes, then there should be more than one record - Issue a waning
	Attribute("MultipleCollateral", String, "Multiple Collateral", func() {
		Meta("rpc:tag", "26")
	})
})

var StoredCollateral = ResultType("", func() {
	TypeName("StoredCollateral")
	Attributes(func() {
		Extend(Collateral)
		Required()
	})
	View("", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocument")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocument")
		Attribute("OtherIDocumentNumber")
		Attribute("PINumber")
		Attribute("CompanyVATNumber")
		Attribute("CollateralType")
		Attribute("CollateralReferenceNumber")
		Attribute("CollateralLastValuationAmount")
		Attribute("Collateral Currency")
		Attribute("CollateralForcedSaleValue")
		Attribute("NextValuationDate")
		Attribute("CollateralExpiryDate")
		Attribute("InstrumentOfClaimOrRecoveryType")
		Attribute("LastValuationDate")
		Attribute("SharedCollateral")
		Attribute("PortionOfCollateralShared")
		Attribute("MultipleCollateral")
	})
})
