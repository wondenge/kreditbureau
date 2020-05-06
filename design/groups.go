package design

import (
	. "goa.design/goa/v3/dsl"
)

var GroupGuarantee = Type("groupguarantee", func() {
	Description("Group Guarantee API")
	TypeName("GroupGuarantee")
	ContentType("application/json")

	Attribute("LendersRegisteredName", String, func() {
		Description("Lenders Registered Name") // As registered with the Registrar of companies.
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
		Description("Lenders Branch Name") // Where the Loan Application happened.
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

	// All persons with this Group ID will be considered to be in the same Group.
	Attribute("GroupID", String, "Group ID", func() {
		Description("Unique Identifier of the Group per Lender")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "5")
	})

	Attribute("GroupName", String, func() {
		Description("The Name of the Group")
		MinLength(1)
		MaxLength(100)
		Meta("rpc:tag", "6")
	})

	// An identifier (Unique within a Group ID) that indicates that
	// the Group Member belongs to a smaller subset of the group;
	// usually for guarantee purposes.
	// All persons with the same Sub Group ID for the Same Group ID
	// are deemed to co-guarantee each other’s group loans.
	Attribute("SubGroupID", String, func() {
		Description("Unique Identifier of the Sub Group per Lender")
		MinLength(1)
		MaxLength(10)
		Meta("rpc:tag", "7")
	})

	Attribute("SubGroupName", String, func() {
		Description("Name of a Sub-Group within a Group.")
		MinLength(1)
		MaxLength(100)
		Meta("rpc:tag", "8")
	})

	Attribute("Surname", String, func() {
		Description("The Family Name or Surname") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "9")
	})

	Attribute("Forename1", String, func() {
		Description("The First Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "10")
	})

	Attribute("Forename2", String, func() {
		Description("The Given Name") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "11")
	})

	Attribute("Forename3", String, func() {
		Description("Other Name or Initials") // Allow Hyphen, Apostrophe, Space.
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "12")
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
		Meta("rpc:tag", "13")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("PrimaryIDocNumber", String, func() {
		Description("Primary Identification Doc Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "14")
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
		Meta("rpc:tag", "15")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("SecondaryIDocumentNumber", String, func() {
		Description("Secondary Identification Document Number")
		MinLength(6)
		MaxLength(8)
		Meta("rpc:tag", "16")
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
		Meta("rpc:tag", "17")
	})

	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters , Not more characters than 6
	Attribute("OtherIDocumentNumber", String, func() {
		Description("Other ID Number provided")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "18")
	})

	Attribute("ClientNumber", String, func() {
		Description("Unique identification of a client")
		MinLength(1)
		MaxLength(20)
		Meta("rpc:tag", "19")
	})

	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, func() {
		Description("Revenue Authority Personal Income Tax Number")
		MinLength(11)
		MaxLength(11)
		Meta("rpc:tag", "20")
	})

	// 	A  = "Active Member of the Group"
	//	B  = "Dormant Member of the Group"
	//	C  = "Exited the Group"
	Attribute("MemberStatus", String, func() {
		Description("Status of member within the group")
		Enum("A", "B", "C")
		MinLength(1)
		MaxLength(1)
		Minimum("A")
		Maximum("C")
		Meta("rpc:tag", "21")
	})

	Attribute("MemberStatusDate", String, func() {
		Description("Date as of the Status indicated")
		MaxLength(8)
		Format(FormatDate) // Date Field
		Meta("rpc:tag", "22")
	})
})

var StoredGroupGuarantee = ResultType("application/vnd.goa.groupguarantee", func() {
	TypeName("StoredGroupGuarantee")
	Attributes(func() {
		Extend(GroupGuarantee)
		Required(
			"LendersRegisteredName",
			"LendersTradingName",
			"LendersBranchName",
			"LendersBranchCode",
			"GroupID",
			"GroupName",
			"Surname",
			"Forename1",
			"PrimaryIDocument",
			"PrimaryIDocNumber",
			"MemberStatus",
			"MemberStatusDate",
		)
	})
	View("default", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("GroupID")
		Attribute("GroupName")
		Attribute("SubGroupID")
		Attribute("SubGroupName")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("Forename2")
		Attribute("Forename3")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("SecondaryIDocument")
		Attribute("SecondaryIDocumentNumber")
		Attribute("OtherIDocument")
		Attribute("OtherIDocumentNumber")
		Attribute("ClientNumber")
		Attribute("PINumber")
		Attribute("MemberStatus")
		Attribute("MemberStatusDate")
	})

	View("mandatory", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("GroupID")
		Attribute("GroupName")
		Attribute("Surname")
		Attribute("Forename1")
		Attribute("PrimaryIDocument")
		Attribute("PrimaryIDocNumber")
		Attribute("MemberStatus")
		Attribute("MemberStatusDate")
	})
})
