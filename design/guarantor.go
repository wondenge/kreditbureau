package design

import (
	. "goa.design/goa/v3/dsl"
)

var Guarantor = Type("guarantor", func() {
	Description("Guarantor API")
	TypeName("Guarantor")
	ContentType("application/json")

	Attribute("LendersRegisteredName", String, func() {
		Description("Lenders Registered Name") // as registered with the Registrar of companies
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

	// Client Reference Number Linking guarantor to Account being
	// guaranteed in the Lenders systems.
	// This is for Client-Centric systems.
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "5")
	})

	// Account Number Linking Guarantor to Account Being Guaranteed
	// to Core Lenders system.
	// The Account Number Is the Same as Client Number for Account
	// Centric Systems.
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "6")
	})

	// Guarantor’s Date of Birth if it is an individual
	// Guarantor’s date of Registration if it is an Institution.
	//
	// Not a future date
	// If an individual, Client Type A,
	// - Not under 18
	// - And Not over 100 years
	// Date Format YYYYMMDD – Issue Warning
	// Registration Date Not in the Future
	Attribute("DOBorDOR", String, func() {
		Description("Date Of Birth / Date of Registration")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "12")
	})

	// M - Male
	// F - Female
	// I - Institutional
	Attribute("Gender", ArrayOf(Enum), func() {
		Description("Guarantor’s Gender")
		Enum(
			MaleGuarantor,          // M - Male
			FemaleGuarantor,        // F - Female
			InstitutionalGuarantor, // I - Institutional
		)
		Meta("rpc:tag", "13")
	})

	Attribute("Nationality", String, func() {
		Description("ISO Code for the Country of Guarantor’s Nationality")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "14")
	})

	// M - Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, func() {
		Description("Guarantor Marital Status")
		Enum("M", "S", "D", "W", "U")
		MinLength(1)
		MaxLength(1)
		Meta("rpc:tag", "15")
	})

	Attribute("GuaranteeType", ArrayOf(Enum), func() {
		Description("Guarantee Type")
		Enum(
			DirectorsGuarantee,              // A - Directors guarantee
			PersonalGuarantee,               // B - Personal guarantee
			CorporateGuarantee,              // C - Corporate guarantee
			BankGuarantee,                   // D - Bank Guarantee
			PersonalGuaranteeWithCollateral, // F – Personal Guarantee with Collateral
		)
		Meta("rpc:tag", "16")
	})

	Attribute("GuaranteeLimit", Int64, func() {
		Description("Amount of Limit of the Guarantee")
		MinLength(1)
		MaxLength(16)
		Minimum(0)
		Meta("rpc:tag", "18")
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
		Meta("rpc:tag", "19")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocNumber", String, func() {
		Description("Primary Identification Doc Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "20")
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
		Meta("rpc:tag", "21")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("SecondaryIDocumentNumber", String, func() {
			Description("Secondary Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "22")
		})
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
		Meta("rpc:tag", "23")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("OtherIDocumentNumber", String, func() {
			Description("Other Identification Document Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "24")
		})
	})

	Attribute("EmployerName", String, func() {
		Description("Guarantor’s Employer Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "25")

		// A - Casual
		// B - Contract
		// C - Permanent
		// D - Pensionable
		// E - Self-Employed
		Attribute("EmploymentType", String, func() {
			Description("Type of employment of the guarantor")
			Enum("A", "B", "C", "D", "E")
			MinLength(1)
			MaxLength(1)
			Minimum("A")
			Maximum("E")
			Meta("rpc:tag", "26")
		})

		// A – KES 0/- To 50,000/-
		// B – KES 50,001/- To 100,000/-
		// C - KES 100,001/- To 200,000/-
		// D - KES 200,001/- To 500,000/-
		// E - Over KES 500,000/-
		Attribute("IncomeBand", String, func() {
			Description("band within which the Customer’s Gross Monthly Salary falls")
			Enum("A", "B", "C", "D", "E")
			MinLength(1)
			MaxLength(1)
			Minimum("A")
			Maximum("E")
			Meta("rpc:tag", "27")
		})

		Required("EmploymentType")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("MobileTelephoneNumber", String, func() {
		Description("Mobile Telephone Number") // Primary Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "28")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("HomeTelephoneNumber", String, func() {
		Description("Home Telephone Number") // Secondary Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "29")
	})

	// In the Form of : CCCAAANNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("WorkTelephoneNumber", String, func() {
		Description("Work Telephone Number") // Employer’s Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "30")
	})

	Attribute("PostalAddress1", String, func() {
		Description("Guarantor’s Postal Address Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "31")
	})

	Attribute("PostalAddress2", String, func() {
		Description("Postal Address Extra details")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "32")
	})

	Attribute("Town", String, func() {
		Description("Town of Postal Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "33")
	})

	Attribute("Country", String, func() {
		Description("Country of Guarantor’s Postal Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "34")
	})

	Attribute("PostCode", String, func() {
		Description("Post Code of Postal Address")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "35")
	})

	// If the Guarantor is an Individual then this would be the
	// Residential Address of the Guarantor.
	// If the Guarantor is an Institution, then this would be
	// the Guarantor’s registered Office Physical Address.
	Attribute("PhysicalAddress1", String, func() {
		Description("Guarantor’s Physical Address Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "36")
	})

	Attribute("PhysicalAddress2", String, func() {
		Description(" Guarantor’s Physical Address Line 2")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "37")
	})

	Attribute("PlotNumber", String, func() {
		Description("Plot Land Ref (LR) No of Guarantor’s Physical Address")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "38")
	})

	Attribute("PhysicalLocationTown", String, func() {
		Description("Town of Guarantor’s Physical Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "39")
	})

	Attribute("PhysicalLocationCountry", String, func() {
		Description("Country of Residence or Registered Office of the Guarantor")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "40")
	})

	Attribute("Email", String, func() {
		Description("Email address of the Guarantor")
		MinLength(1)
		MaxLength(50)
		Format(FormatEmail)
		Meta("rpc:tag", "41")
	})

})

var StoredGuarantor = ResultType("application/vnd.goa.guarantor", func() {
	TypeName("StoredGuarantor")
	Attributes(func() {
		Extend(Guarantor)
		Required("LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"AccountNumber",
			"DOBorDOR",
			"Gender",
			"Nationality",
			"GuaranteeType",
			"GuaranteeLimit",
			"PrimaryIDocument",
			"PrimaryIDocNumber",
			"PostalAddress1",
			"Town",
			"Country",
			"PhysicalLocationTown",
			"PhysicalLocationCountry")
	})
	View("default", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyOrCorporateName")
		Attribute("DOBorDOR")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("MaritalStatus")
		Attribute("GuaranteeType")
		Attribute("GuarantorRelationship")
		Attribute("GuaranteeLimit")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocument")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocument")
		Attribute("OtherIDocumentNumber")
		Attribute("EmployerName")
		Attribute("EmploymentType")
		Attribute("IncomeBand")
		Attribute("MobileTelephoneNumber")
		Attribute("HomeTelephoneNumber")
		Attribute("WorkTelephoneNumber")
		Attribute("PostalAddress1")
		Attribute("PostalAddress2")
		Attribute("Town")
		Attribute("Country")
		Attribute("PostCode")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalAddress2")
		Attribute("PlotNumber")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
		Attribute("Email")
	})
	View("mandatory", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("AccountNumber")
		Attribute("DOBorDOR")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("GuaranteeType")
		Attribute("GuaranteeLimit")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("PostalAddress1")
		Attribute("Town")
		Attribute("Country")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
	})
})

// M - Male
var MaleGuarantor = Type("MaleGuarantor", func() {
	Attribute("Surname", String, func() {
		Description("Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "7")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "8")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "9")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "10")
	})

	Required("Surname", "Forename1", "Forename2", "Forename3")

})

// F - Female
var FemaleGuarantor = Type("FemaleGuarantor", func() {
	Attributes(func() {
		Extend(MaleGuarantor)
		Required("Surname", "Forename1", "Forename2", "Forename3")
	})
})

// I - Institutional
var InstitutionalGuarantor = Type("InstitutionalGuarantor", func() {
	Attribute("CompanyOrCorporateName", String, func() {
		Description("Name of the Company/Corporate entity that is a Guarantor")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "11")
	})
	Required("CompanyOrCorporateName")
})

// A - Directors guarantee
var DirectorsGuarantee = Type("DirectorsGuarantee", func() {
	Description("Guarantee Type: A - Directors Guarantee")

	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// B = Shareholder, Partner, etc.
	// C = Other
	Attribute("GuarantorRelationship", String, func() {
		Description("Clients relations with the Guarantor")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Default("B") // B = Shareholder, Partner, etc.
		Meta("rpc:tag", "17")
	})

	Required("GuarantorRelationship")

})

// B - Personal guarantee
var PersonalGuarantee = Type("PersonalGuarantee", func() {
	Description("Guarantee Type: B - Personal Guarantee")

	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// B = Shareholder, Partner, etc.
	// C = Other
	Attribute("GuarantorRelationship", String, func() {
		Description("Clients relations with the Guarantor")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Default("A") // A = Family (e.g. Husband, Wife, Daughter, etc.)
		Meta("rpc:tag", "17")
	})

	Required("GuarantorRelationship")
})

// C - Corporate guarantee
var CorporateGuarantee = Type("CorporateGuarantee", func() {
	Description("Guarantee Type: C - Corporate guarantee")

	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// B = Shareholder, Partner, etc.
	// C = Other
	Attribute("GuarantorRelationship", String, func() {
		Description("Clients relations with the Guarantor")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Default("B") // B = Shareholder, Partner, etc.
		Meta("rpc:tag", "17")
	})

	Required("GuarantorRelationship")
})

// D - Bank Guarantee
var BankGuarantee = Type("BankGuarantee", func() {
	Description("Guarantee Type: D - Bank Guarantee")

	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// B = Shareholder, Partner, etc.
	// C = Other
	Attribute("GuarantorRelationship", String, func() {
		Description("Clients relations with the Guarantor")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Default("C") // C = Other
		Meta("rpc:tag", "17")
	})

	Required("GuarantorRelationship")

})

// F – Personal Guarantee with Collateral
var PersonalGuaranteeWithCollateral = Type("PersonalGuaranteeWithCollateral", func() {
	Description("Guarantee Type: F – Personal Guarantee with Collateral")

	// A = Family (e.g. Husband, Wife, Daughter, etc.)
	// B = Shareholder, Partner, etc.
	// C = Other
	Attribute("GuarantorRelationship", String, func() {
		Description("Clients relations with the Guarantor")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Default("C") // C = Other
		Meta("rpc:tag", "17")
	})

	Required("GuarantorRelationship")
})
