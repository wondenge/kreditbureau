package design

import (
	. "goa.design/goa/v3/dsl"
)

var Collateral = Type("collateral", func() {
	Description("Collateral API")
	TypeName("Collateral")
	ContentType("application/json")

	// Registered with the Registrar of companies.
	Attribute("LendersRegisteredName", String, func() {
		Description("Lenders Registered Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "1")
	})

	Attribute("LendersTradingName", String, func() {
		Description("Lenders Trading Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("LendersBranchName", String, func() {
		Description("Lenders Branch Name")
		MinLength(1)
		MaxLength(50)
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
	Attribute("LendersBranchCode", String, func() {
		Description("Lenders Branch Code")
		MinLength(1)
		MaxLength(7)
		Meta("rpc:tag", "4")
	})

	// Client Reference Number Linking Borrower‘s
	// Account in the Lenders systems.
	// This is for Client-Centric systems.
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "5")
	})

	// Account Number Linking Collateral to Borrower’s
	// Account in the Lender’s System.
	// The Account Number Is the Same as Client Number
	// for Account Centric Systems.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "6")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("PrimaryIDocument", Int, func() {
		Description("Primary Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "7")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("PrimaryIDocNumber", String, func() {
			Description("Primary Identification Doc Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "8")
		})

		Required("PrimaryIDocNumber")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("SecondaryIDocument", Int, func() {
		Description("Secondary Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "9")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("SecondaryIDocumentNumber", String, func() {
			Description("Secondary Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "10")
		})

		Required("SecondaryIDocumentNumber")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("OtherIDocument", Int, func() {
		Description("Other Identification Document")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "11")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("OtherIDocumentNumber", String, func() {
			Description("Other Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "12")
		})
		Required("OtherIDocumentNumber")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, func() {
		Description("Revenue Authority Personal Income Tax Number")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "13")
	})

	Attribute("CompanyVATNumber", String, "Company VAT Number", func() {
		Description("Company VAT Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "14")
	})

	Attribute("CollateralType", ArrayOf(Enum), func() {
		Description("Collateral Type")
		Enum(CollateralTypeA, // A – Land
			CollateralTypeB, // B – Property and Buildings
			CollateralTypeC, // C – Cash
			CollateralTypeD, // D – Treasury Bills and Bonds
			CollateralTypeE, // E – Shares
			CollateralTypeF, // F – Debentures
			CollateralTypeG, // G – Chattels/Charges over movable assets excluding vehicles
			CollateralTypeH, // H – Chattels over Vehicles
			CollateralTypeI, // I – Insurance Policy
		)
		Meta("rpc:tag", "15")
	})

	// If the Collateral is shared for several Loans
	// Y= If Collateral is shared
	// N= If Collateral is not shared
	Attribute("SharedCollateral", String, func() {
		Description("Collateral is shared for several Loans")
		Enum("Y", "N")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "24")

		// Conditional if Shared Collateral = Y
		Attribute("PortionOfCollateralShared", Int, func() {
			Description("Portion of Collateral Shared")
			MinLength(1)
			MaxLength(3)
			Meta("rpc:tag", "25")
		})
	})

	// Y= If multiple collateral
	// N= If no multiple collateral
	Attribute("MultipleCollateral", String, func() {
		Description("Multiple Collateral")
		Enum("Y", "N")
		MinLength(1)
		MaxLength(1)
		Example("Y")
		Meta("rpc:tag", "26")
	})
})

var StoredCollateral = ResultType("application/vnd.goa.collateral", func() {
	TypeName("StoredCollateral")
	Attributes(func() {
		Extend(Collateral)
		Required("LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"AccountNumber",
			"PrimaryIDocument",
			"PrimaryIDocNumber",
			"CollateralType",
		)
	})

	View("default", func() {
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

	View("mandatory", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("CollateralType")
	})
})

// A – Land
var CollateralTypeA = Type("CollateralTypeA", func() {
	Description("Collateral Type A - Land")

	// To match/trace collateral used on same exposures.
	// Mandatory if Collateral type is Land, Property and Buildings.
	Attribute("CollateralReferenceNumber", String, func() {
		Description("Collateral Reference Number")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "16")
	})

	// Amount of the collateral as at Last valuation.
	//
	// Amount Field
	// Mandatory Field
	// Not more than 16 Characters
	Attribute("CollateralLastValuationAmount", Int, func() {
		Description("Collateral Last Valuation Amount")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "17")
	})

	Attribute("Collateral Currency", String, func() {
		Description("Currency Code of the Collateral currency")
		MinLength(3)
		MaxLength(3)
		Default("KES") // Default is Kenya Shillings.
		Meta("rpc:tag", "18")
	})

	// Conditional to field Collateral Type, if Land, Property and Buildings.
	Attribute("CollateralForcedSaleValue", Int64, func() {
		Description("Collateral Forced Sale Value")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "19")
	})

	// Date when collateral is scheduled for valuation.
	Attribute("NextValuationDate", String, "", func() {
		Description("Next Valuation Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "20")
	})

	Attribute("CollateralExpiryDate", String, "", func() {
		Description("Collateral Expiry Date")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "21")
	})

	Attribute("InstrumentOfClaimOrRecoveryType", String, func() {
		Description("Brief description of instruments used to recover")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "22")
	})

	Attribute("LastValuationDate", String, func() {
		Description("Date the Collateral was last Valued")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "23")
	})

	Required("CollateralReferenceNumber",
		"CollateralForcedSaleValue",
		"CollateralLastValuationAmount",
		"Collateral Currency",
		"LastValuationDate",
	)

})

// B – Property and Buildings
var CollateralTypeB = Type("CollateralTypeB", func() {
	Description("Collateral Type B - Property and Buildings")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralReferenceNumber",
			"CollateralForcedSaleValue",
			"CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// C – Cash
var CollateralTypeC = Type("CollateralTypeC", func() {
	Description("Collateral Type C - Cash")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// D – Treasury Bills and Bonds
var CollateralTypeD = Type("CollateralTypeD", func() {
	Description("Collateral Type D - Treasury Bills and Bonds")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralReferenceNumber",
			"CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// E – Shares
var CollateralTypeE = Type("CollateralTypeE", func() {
	Description("Collateral Type E - Shares")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralReferenceNumber",
			"CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// F – Debentures
var CollateralTypeF = Type("CollateralTypeF", func() {
	Description("Collateral Type F - Debentures")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// G – Chattels/Charges over movable assets excluding vehicles
var CollateralTypeG = Type("CollateralTypeG ", func() {
	Description("Collateral Type G - Chattels/Charges over movable assets excluding vehicles")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// H – Chattels over Vehicles
var CollateralTypeH = Type("CollateralTypeH", func() {
	Description("Collateral Type H - Chattels over Vehicles")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralReferenceNumber",
			"CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})

// I – Insurance Policy
var CollateralTypeI = Type("CollateralTypeI", func() {
	Description("Collateral Type I - Insurance Policy")
	Attributes(func() {
		Extend(CollateralTypeA)
		Required("CollateralReferenceNumber",
			"CollateralLastValuationAmount",
			"Collateral Currency",
			"LastValuationDate",
		)
	})
})
