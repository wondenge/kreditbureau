package design

import (
	. "goa.design/goa/v3/dsl"
)

// K. Collateral Register

var Collateral = Type("collateral", func() {
	Description("")
	TypeName("Collateral")
	ContentType("application/json")

	// The Primary Identification document Provided by Loanee.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	// Company Registration No
	//
	// National Identification is the preferred document for individuals but the others are acceptable.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	// The company registration Number is the Registration Number of the Institution providing the collateral.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Number of the Primary Id Provided
	// For Kenya nationals, the default primary identification is the National Identification Card Number
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification Document Type is provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Any Other Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Number of Other Identification Document, where provided .
	// Mandatory if Other Identification Document Type is provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification Document Type is provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Kenya Revenue Authority PIN Number of collateral Holder
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "9")
	})

	//  Location where collateral is situated, Default is Kenya .
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Applicable to Companies and non-individual consumers
	Attribute("CompanyRegistrationNumber", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Bank and Branch Code accepting the Collateral
	Attribute("BankAndBranchCode", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// Reference Number Linking Collateral to Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Account number Linking Bank’s Accounting System to Account on which collateral is pledged
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Type of Collateral.
	// Options Are:
	// A - Property
	// B - Funds
	// C - Land title
	// D - Log Book
	// E - Debentures/Shares
	// F – Insurance
	// G - Others
	Attribute("CollateralType", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Identification number or Certificate number of the collateral
	Attribute("CollateralReferenceNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Description of the Collateral
	Attribute("CollateralDescription", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Last valuation of the collateral
	Attribute("CollateralLastValuation", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// ISO Currency Code of the Valuation
	Attribute("CollateralCurrency", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Mandatory if Land, vehicle or house
	Attribute("CollateralForcedSaleValue", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// Mandatory if Forced Sale Value is provided
	Attribute("ForcedSaleDate", String, "", func() {
		Meta("rpc:tag", "21")
	})
	Attribute("CollateralExpiryDate", String, "", func() {
		Meta("rpc:tag", "22")
	})
	Attribute("InstrumentOfClaim", String, "", func() {
		Meta("rpc:tag", "23")
	})
	Attribute("LastValuationDate", String, "", func() {
		Meta("rpc:tag", "24")
	})
	Attribute("RecoveryType", String, "", func() {
		Meta("rpc:tag", "25")
	})
})
