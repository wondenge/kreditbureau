package design

import (
	. "goa.design/goa/v3/dsl"
)

var GroupGuarantee = Type("groupguarantee", func() {
	Description("Group Guarantee API")
	TypeName("GroupGuarantee")
	ContentType("application/json")

	// The Name of Lender Reporting the Loan Application,
	// as registered with the Registrar of companies.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("LendersRegisteredName", String, "Lenders Registered Name", func() {
		Meta("rpc:tag", "1")
	})

	// The Lenders Trading Name.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 Characters
	Attribute("LendersTradingName", String, "Lenders Trading Name", func() {
		Meta("rpc:tag", "2")
	})

	// The Lenders Branch Name, where the Loan Application
	// is reported to ave taken place.
	//
	// Mandatory Field
	// Alphanumeric
	// Not more than 50 Characters
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
	Attribute("LendersBranchCode", String, "Lenders Branch Code", func() {
		Meta("rpc:tag", "4")
	})

	// The Unique Identifier of the Group per Lender.
	// All persons with this Group ID will be considered to be in
	// the same Group.
	//
	// Alphanumeric
	// Mandatory Field
	// Not more than 10 Characters
	Attribute("GroupID", String, "Group ID", func() {
		Meta("rpc:tag", "5")
	})

	// The Name of the Group
	// Alphanumeric
	// Not more than 100 Characters
	// Mandatory Field
	Attribute("GroupName", String, "Group Name", func() {
		Meta("rpc:tag", "6")
	})

	// An identifier (Unique within a Group ID) that indicates that the
	// Group Member belongs to a smaller subset of the group; usually
	// for guarantee purposes. All persons with the same Sub Group ID
	// for the Same Group ID are deemed to co-guarantee each other’s group loans.
	//
	// Alphanumeric
	// Not more than 10 Characters
	Attribute("SubGroupID", String, "Sub Group ID", func() {
		Meta("rpc:tag", "7")
	})

	// The name of a Sub-Group within a Group.
	//
	// Alphanumeric
	// Not more than 100 Characters
	Attribute("SubGroupName", String, "Sub Group Name", func() {
		Meta("rpc:tag", "8")
	})

	// The Family Name or Surname.
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe, Space.
	Attribute("Surname", String, "Surname", func() {
		Meta("rpc:tag", "9")
	})

	// The First Name
	//
	// Alphanumeric
	// More than 1 character
	// Mandatory
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("FirstName", String, "The First Name", func() {
		Meta("rpc:tag", "10")
	})

	// The Given Name
	//
	// Alphanumeric
	// More than 1 character
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe,
	Attribute("GivenName", String, "The Given Name", func() {
		Meta("rpc:tag", "11")
	})

	// Other Name or Initials
	//
	// Alphanumeric
	// Not more than 50 Characters
	// Allow Hyphen, Apostrophe.
	Attribute("OtherName", String, "Other Name or Initials", func() {
		Meta("rpc:tag", "12")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Mandatory Field
	// Data is Alpha Numeric
	// Not more than 3 characters
	Attribute("PrimaryIDocument", String, "Primary Identification Document", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Identification Document provided
	// on Opening of the Account.
	//
	// Mandatory Field
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien : Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("PrimaryIDocNumber", String, "Primary Identification Doc Number", func() {
		Meta("rpc:tag", "14")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	// 004 - Service ID
	// 005 - Company Registration Number
	//
	// Data is Alpha Numeric
	// Not more than 3 characters
	Attribute("SecondaryIDocument", String, "Secondary Identification Document", func() {
		Meta("rpc:tag", "15")
	})

	// The Number of the secondary Identification Document provided
	// on Opening of the Account.
	//
	// For ID: Numeric Characters, Between 1-8 characters
	// For Passport: Alphanumeric
	// For Alien: Numeric characters, Not more than 6 characters
	// For Service ID: Numeric Characters, Not more than 6 characters
	Attribute("SecondaryIDocumentNumber", String, "Secondary Identification Document Number", func() {
		Meta("rpc:tag", "16")
	})

	// Refer to for Acceptable Values
	// 001 - National ID
	// 002 - Passport
	// 003 - Alien Registration
	//  - Service ID
	// 005 - Company Registration Number
	//
	// Data is Alpha Numeric
	// Not more than 3 characters
	Attribute("OtherIDocument", String, "Other Identification Document", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of the other Identification Document provided
	// on Opening of the Account.
	//
	// Data is Alphanumeric
	// Not more than 20 characters
	Attribute("OtherIDocumentNumber", String, "Other Identification Document Number", func() {
		Meta("rpc:tag", "18")
	})

	// Unique identification of a client within the lenders Core System.
	//
	// Data is Alpha Numeric
	// Not more than 20 characters
	Attribute("ClientNumber", String, "Client Number", func() {
		Meta("rpc:tag", "19")
	})

	// Revenue Authority Personal Income Tax Number.
	//
	// Data Format is Alphanumeric
	// Not more than 11 characters
	// - 1 chr is Alphanumeric
	// - Last chr is Alphanumeric
	// - Rest are 9 Numeric characters only
	Attribute("PINumber", String, "PIN Number", func() {
		Meta("rpc:tag", "20")
	})

	// The Status of this Member within the Group.
	// 	A  = "Active Member of the Group"
	//	B  = "Dormant Member of the Group"
	//	C  = "Exited the Group"
	//
	// Alphanumeric
	// Not more than 1 character
	// Mandatory
	Attribute("MemberStatus", String, "Member Status", func() {
		Meta("rpc:tag", "21")
	})

	// The Date as of the Status indicated.
	//
	// Date Field
	// Mandatory Field
	// Can’t be in future
	// Not more than 8 Characters
	Attribute("MemberStatusDate", String, "Member Status Date", func() {
		Meta("rpc:tag", "22")
	})
})

var StoredGroupGuarantee = ResultType("", func() {
	TypeName("StoredGroupGuarantee")
	Attributes(func() {
		Extend(GroupGuarantee)
		Required()
	})
	View("", func() {
		Attribute("LendersRegisteredName")
		Attribute("LendersTradingName")
		Attribute("LendersBranchName")
		Attribute("LendersBranchCode")
		Attribute("GroupID")
		Attribute("GroupName")
		Attribute("SubGroupID")
		Attribute("SubGroupName")
		Attribute("Surname")
		Attribute("FirstName")
		Attribute("GivenName")
		Attribute("OtherName")
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
})
