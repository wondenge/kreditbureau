package design

import (
	. "goa.design/goa/v3/dsl"
)

var Stakeholder = Type("stakeholder", func() {
	Description("Stakeholder API")
	TypeName("Stakeholder")
	ContentType("application/json")

	// Stakeholder’s Date of Birth (individual) Or
	// Stakeholder’s Date of Registration (non -individual).
	Attribute("DoBOrDoR", String, func() {
		Description("Date of Birth / Date of Registration")
		MaxLength(8)
		Format(FormatDate)
		Meta("rpc:tag", "6")
	})

	// M – Male
	// F – Female
	// I – Institution
	Attribute("Gender", String, ArrayOf(Enum), func() {
		Description("Stakeholder’s Gender")
		Enum(StakeholderMale, StakeholderFemale, StakeholderInstitution)
		Meta("rpc:tag", "7")
	})

	Attribute("Nationality", String, "Nationality", func() {
		Description("ISO Country Code for Stakeholder’s Nationality")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "8")
	})

	// Client reference Number in Core Lenders system for the Account in
	// which this Stakeholder has linkage . For Client-Centric Systems.
	Attribute("ClientNumber", String, func() {
		Description("Client Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "9")
	})

	// Account Number in Core Lenders system for the Account in which
	// this Stakeholder has linkage.
	// Same as Client Number for Account Centric Systems
	Attribute("AccountNumber", String, func() {
		Description("Account Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "10")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// If Gender = I, then stakeholder identification = 005
	Attribute("PrimaryIDocument", Int, func() {
		Description("Stakeholder’s Primary ID")
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
		Attribute("PrimaryIDocNumber", String, func() {
			Description("Stakeholder’s Primary ID Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "12")
		})
		Required("PrimaryIDocNumber")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("SecondaryIDocumentType", Int, func() {
		Description("Stakeholder’s Secondary ID Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "13")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("SecondaryIDocumentNumber", String, func() {
			Description("Stakeholder’s Secondary ID Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "14")
		})
		Required("SecondaryIDocumentNumber")
	})

	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	Attribute("OtherIDocumentType", Int, func() {
		Description("Stakeholder Other ID Type")
		Enum(001, 002, 003, 004, 005)
		MinLength(3)
		MaxLength(3)
		Minimum(001)
		Maximum(005)
		Meta("rpc:tag", "15")

		// For ID: Numeric Characters, Between 1-8 characters
		// For Passport: Alphanumeric
		// For Alien: Numeric characters, Not more than 6 characters
		// For Service ID: Numeric Characters , Not more characters than 6
		Attribute("OtherIDocNumber", String, func() {
			Description("Stakeholder’s Other ID Number")
			MinLength(6)
			MaxLength(8)
			Meta("rpc:tag", "16")
		})

		Required("OtherIDocNumber")
	})

	Attribute("Email", String, func() {
		Description("Stakeholder’s Email Address")
		MinLength(1)
		MaxLength(50)
		Format(FormatEmail) // Validate Email Syntax
		Meta("rpc:tag", "17")
	})

	// A – Director
	// B - Shareholder
	// C - Partner
	// D - Proprietor
	// E - Trust Beneficiary
	// F - Power of Attorney
	// G - Guarantor
	Attribute("StakeholderType", String, func() {
		Description("Stakeholder Type")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("G")
		Meta("rpc:tag", "22")
	})

	// If stakeholder is a shareholder,
	//the Shareholding in the company.
	Attribute("PercentageShares", Int, func() {
		Description("Percentage of Shares in company")
		MinLength(1)
		MaxLength(6)
		Minimum(0)
		Meta("rpc:tag", "23")
	})

	// Form of : CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("MobileTelephoneNumber", String, "Mobile Telephone Number", func() {
		Description("Mobile Telephone Number") // Primary Stakeholder’s Telephone contact Number
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "24")
	})

	// Form of : CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("HomeTelephoneNumber", String, "Home Telephone Number", func() {
		Description("Secondary Telephone contact Number")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "25")
	})

	// Form of : CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("WorkTelephoneNumber", String, func() {
		Description("Work Telephone Number")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "26")
	})

	Attribute("PostalAddress1", String, "Postal Address 1", func() {
		Description("Stakeholder’s Postal Address Line 1")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "27")
	})

	Attribute("PostalAddress2", String, "Postal Address 2", func() {
		Description("Stakeholder’s Postal Address Line 2")
		MinLength(1)
		MaxLength(15)
		Meta("rpc:tag", "28")
	})

	Attribute("Town", String, "Town", func() {
		Description("Town of Stakeholder’s Postal Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "29")
	})

	Attribute("Country", String, func() {
		Description("Country of Stakeholder’s Postal Address")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "30")
	})

	Attribute("PostCode", String, func() {
		Description("Post Code of Stakeholder’s Postal Address")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "31")
	})

	// If stakeholder is an individual;
	// this would be the Residential Address.
	// If the stakeholder is an institution;
	// this would be the registered Office Address.
	Attribute("PhysicalAddress1", String, "Physical Address 1", func() {
		Description("Stakeholder’s Physical Address Line 1")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "32")
	})

	Attribute("PhysicalAddress2", String, func() {
		Description("Stakeholder’s Physical Address Line 2")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "33")
	})

	Attribute("PhysicalAddress3", String, func() {
		Description("Stakeholder’s Physical Address Line 3")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "34")
	})

	Attribute("PlotNumber", String, func() {
		Description("Plot Land Ref (LR) No of Stakeholder’s Physical Address")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "35")
	})

	Attribute("PhysicalLocationTown", String, func() {
		Description("Town of Stakeholder’s Physical Address")
		MinLength(1)
		MaxLength(30)
		Meta("rpc:tag", "36")
	})

	Attribute("PhysicalLocationCountry", String, func() {
		Description("ISO Country Code of the Stakeholder’s Physical Address")
		MinLength(2)
		MaxLength(2)
		Default("KE")
		Meta("rpc:tag", "37")
	})
})

var StoredStakeholder = ResultType("application/vnd.goa.stakeholder", func() {
	TypeName("StoredStakeholder")
	Attributes(func() {
		Extend(Stakeholder)
		Required(
			"DoBOrDoR",
			"Gender",
			"Nationality",
			"AccountNumber",
			"PrimaryIDocument",
			"StakeholderType",
			"PhysicalAddress1",
			"PhysicalLocationTown",
			"PhysicalLocationCountry",
		)
	})

	View("default", func() {
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("CompanyOrCorporateName")
		Attribute("DoBOrDoR")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("ClientNumber")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocumentType")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocumentType")
		Attribute("OtherIDocNumber")
		Attribute("Email")
		Attribute("CompanyRegistrationNumber")
		Attribute("PreviousRegistrationNumber")
		Attribute("CompanyPINumber")
		Attribute("CompanyVATNumber")
		Attribute("StakeholderType")
		Attribute("PercentageShares")
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
		Attribute("PhysicalAddress3")
		Attribute("PlotNumber")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
	})

	View("mandatory", func() {
		Attribute("DoBOrDoR")
		Attribute("Gender")
		Attribute("Nationality")
		Attribute("AccountNumber")
		Attribute("PrimaryIDocument")
		Attribute("StakeholderType")
		Attribute("PhysicalAddress1")
		Attribute("PhysicalLocationTown")
		Attribute("PhysicalLocationCountry")
	})
})

var StakeholderMale = Type("StakeholderMale", func() {
	Description("Male Stakeholder Name properties")

	// Conditional to 4.3.7
	Attribute("Surname", String, "Surname", func() {
		Description("Stakeholder’s Family Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "1")
	})

	// Conditional to 4.3.7
	Attribute("Forename1", String, func() {
		Description("Stakeholder’s First Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "2")
	})

	Attribute("Forename2", String, func() {
		Description("Stakeholder’s Given Name")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "3")
	})

	Attribute("Forename3", String, func() {
		Description("Stakeholder’s Other Name or Initials")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "4")
	})

	Required("Surname", "Forename1")

})
var StakeholderFemale = Type("StakeholderFemale ", func() {
	Description("Female Stakeholder Name properties")
	Attributes(func() {
		Extend(StakeholderMale)
		Required("Surname", "Forename1")

	})

})

var StakeholderInstitution = Type("", func() {
	Description("Institution Stakeholder properties")

	// Name of the Company/Corporate entity that is a stakeholder
	// to the Subject Account.
	Attribute("CompanyOrCorporateName", String, func() {
		Description("Name of the Company/Corporate entity")
		MinLength(1)
		MaxLength(70)
		Meta("rpc:tag", "5")
	})

	// Should be the same with field 4.3.12 if option 005 is provided in 4.3.11
	Attribute("CompanyRegistrationNumber", String, func() {
		Description("Company Registration Certificate Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "18")
	})

	// The Registration Certificate Number - format to conform to
	// Registrar of Companies in Kenya.
	// Required, only if applicable because the company type changes,
	// the history details of the previous type can be matched to
	// the new company type.
	Attribute("PreviousRegistrationNumber", String, func() {
		Description("Previous Registration Number")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "19")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("CompanyPINumber", String, "Company PIN Number", func() {
		Description("Company Income Tax PIN Number")
		MinLength(1)
		MaxLength(11)
		Meta("rpc:tag", "20")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("CompanyVATNumber", String, func() {
		Description("Company Income Tax VAT Number")
		MinLength(1)
		MaxLength(11)
		Meta("rpc:tag", "21")
	})

	Required("CompanyOrCorporateName")
})
