package design

import (
	. "goa.design/goa/v3/dsl"
)

// A. Institution Information
// The Institution refers to the Participating institution (Currently Commercial Banks).
// Most of the information in this table is static information pertaining to the bank and its registration details.

var BankInfo = Type("bankinfo", func() {
	Description("static information pertaining to the bank and its registration details.")
	TypeName("BankInfo")
	ContentType("application/json")

	// The Name of the Institution as registered with the Registrar of companies
	Attribute("RegisteredName", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Business or Trading Name of the institution, if different from the Registered Name.
	// Same as The Registered Name as registered with the Registrar of companies if not provided.
	Attribute("TradingName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Date of Registration with Registrar of Companies.
	// Date Format is DDMMYYYY
	Attribute("RegistrationDate", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Registration Certificate Number issued by the Registrar of Companies.
	Attribute("RegistrationNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Date the Institution was issued with a banking licence by the Central Bank
	// of Kenya or the Date of Renewal of the Banking License.
	// This date Cannot be in the future
	Attribute("LicenseIssueDate", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Applicable Provider Type. Currently Banking Services.
	// Will extend to other industries once we open the CRB submissions to other provider classes.
	Attribute("ProviderType", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// The Kenya Revenue Authority Income Tax PIN Number.
	// May be Mandatory in Future.
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Institution’s Email Address.
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// The Institution’s Web site url.
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// All Credit facilities with debit balances that arise from business transactions.
	// Defined by the Central Bank
	// i.e. Movements in the account are not merely the result of charges levied.
	Attribute("NrOfActiveAccounts", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// The Total Active bank accounts in the institution
	Attribute("ActiveAccountsValue", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// The Total number of Closed accounts in the institution in the period since the last reporting date.
	Attribute("NrOfClosedAccounts", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Total number of New accounts opened in the institution in the period since the last reporting date.
	Attribute("NrOfNewAccounts", String, " ", func() {
		Meta("rpc:tag", "13")
	})

	// Total Value of New Accounts since the last reporting date.
	Attribute("NewAccountsValue", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Total Number of Non-Performing Loan Accounts since the last reporting date
	// As defined by the Central Bank under section 1.4.2 of the prudential guidelines
	// no CBK/PG/04 on risk classification of Assets and provisioning.
	// The guideline provides a definition for both loans and overdrafts.
	Attribute("NrNonPerformingLoans", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Total Value of Non-Performing Loan Accounts since the Last reporting date.
	Attribute("NonPerformingLoansValue", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// The Two-Digit KBA allocated Bank Code
	Attribute("BankCode", String, "", func() {
		Meta("rpc:tag", "17")
	})
	View("default", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("LicenseIssueDate")
		Attribute("ProviderType")
		Attribute("PINNumber")
		Attribute("Email")
		Attribute("Website")
		Attribute("NrOfActiveAccounts")
		Attribute("ActiveAccountsValue")
		Attribute("NrOfClosedAccounts")
		Attribute("NrOfNewAccounts")
		Attribute("NewAccountsValue")
		Attribute("NrNonPerformingLoans")
		Attribute("NonPerformingLoansValue")
		Attribute("BankCode")
	})

	Required("RegisteredName")
})

var StoredInstitutionInfo = ResultType("application/vnd.goa.bankinfo", func() {
	Description("")
	TypeName("StoredInstitutionInfo")
	ContentType("application/json")
	Attributes(func() {
		Extend(BankInfo)
		Required("RegisteredName")
	})
	View("default", func() {
		Attribute("RegisteredName")
		Attribute("TradingName")
		Attribute("RegistrationDate")
		Attribute("RegistrationNumber")
		Attribute("LicenseIssueDate")
		Attribute("ProviderType")
		Attribute("PINNumber")
		Attribute("Email")
		Attribute("Website")
		Attribute("NrOfActiveAccounts")
		Attribute("ActiveAccountsValue")
		Attribute("NrOfClosedAccounts")
		Attribute("NrOfNewAccounts")
		Attribute("NewAccountsValue")
		Attribute("NrNonPerformingLoans")
		Attribute("NonPerformingLoansValue")
		Attribute("BankCode")
	})
})

// B. Institution Branch Information
// The Institution Branches are defined in the following table.
// All the branches of the institution are required to be reported.

var BranchInfo = Type("branchinfo", func() {
	Description("Institution Branches")
	TypeName("BranchInfo")
	ContentType("application/json")

	// The Bank and Branch Code as allocated by the KBA Secretariat.
	// Also Known as the Branch Sort Code
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Branch Name
	Attribute("BranchName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Date the Branch Opened for business.
	// Date should not be in the future.
	Attribute("DateOpened", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Branch email Address
	// Same as Parent Bank email Address if none provided.
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Website link, if different from the Main institution website.
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Number of Active Accounts in the Branch.
	// Definition of active accounts as per the Central Bank.
	Attribute("NrOfActiveAccounts", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Total Value of the Active Accounts in the branch.
	// The Number of closed accounts in the branch since the last report.
	Attribute("ActiveAccountsValue", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Number of closed accounts in the branch since the last report.
	Attribute("NrOfClosedAccounts", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// The Number of New accounts opened in the branch since the last report.
	Attribute("NrOfNewAccounts", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Total Value of the New Accounts opened in the branch since the last report.
	Attribute("NewAccountsValue", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Number of Non-Performing Loan Accounts in the branch since the last report.
	// As defined by the Central Bank under section 1.4.2 of the prudential guidelines
	// no CBK/PG/04 on risk classification of Assets and provisioning.
	// The guideline provides a definition for both loans and overdrafts.
	Attribute("NrNonPerformingLoans", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Total Value of the Non-Performing Loan accounts in the branch since the last report.
	Attribute("NonPerformingLoanValue", String, "", func() {
		Meta("rpc:tag", "12")
	})
})

// C. Individual Consumer Information
// The Individual consumer record refers to the individual customer of the
// institutions and contains a profile of the customer from the account opening
// information and any other information the customer provides to the institution.

var CustomerInfo = Type("customerinfo", func() {
	Description("Individual consumer record")
	TypeName("CustomerInfo")
	ContentType("application/json")

	// The Family Name or Surname
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "1")
	})

	//The First Name
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	Attribute("Forename3", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Options are :
	// Mr., Mrs., Miss, Ms, Dr. , Prof., Hon.
	Attribute("Salutation", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Date of Birth of the Customer.
	// Cannot be in the Future.
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Client Reference Number linking client to Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number linking client to Banking system
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Male/Female ( M/F)
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Kenya Revenue Authority Income Tax PIN Number.
	// May be required at a later Date.
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// The Nationality of the Customer- Defaults to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Options to Select From will be provided
	// Options Available :
	// M – Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Primary Identification document Provided on Opening the Account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	//
	// For Kenya nationals, default should be the National Identification.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Identification document specified above.
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Any Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification document is Provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Any Other Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of Other Identification Document, where provided.
	// Mandatory if Other Identification Document is Provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Any Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification document is Provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// The Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// The Any other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// Client’s Postal Address Line1
	// This is the first line of the Full Client’s Postal address
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Client’s Postal Address Line 2
	// This is the second line of the Full Client’s Postal address
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Town of Client’s Postal Address
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Country of Client’s Postal Address
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// Post Code of Client’s Address
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Client’s residential Address e.g. Street Address, Estate or village
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// House or Apartment number of Client’s Residence
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Plot Land Ref (LR) No of Client’s residence
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Town of Client’s residence
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Country of Client’s residence
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Client’s Email Address if provided
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "34")
	})

})

// D. Individual Consumer Employment Information
// The Institution’s customers could be employed or self-employed.
// Details of the customer’s employment are captured in the employment information record.
// Where a customer has provided employment record details, the institution is required to
// provide one record for the employment details as laid out below

var EmploymentInfo = Type("employmentinfo", func() {
	Description("customer’s employment details")
	TypeName("EmploymentInfo")
	ContentType("application/json")

	// The Family Name or Surname of Customer.
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "1")
	})

	//The First Name of Customer.
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	Attribute("Forename3", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Options are :
	// Mr., Mrs., Miss, Ms, Dr. , Prof., Hon.
	Attribute("Salutation", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Date of Birth of the Customer.
	// Cannot be in the Future.
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Client Reference Number linking client to Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number linking client to Banking system
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Male/Female ( M/F)
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Kenya Revenue Authority Income Tax PIN Number.
	// May be required at a later Date.
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// The Nationality of the Customer- Defaults to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Options to Select From will be provided
	// Options Available :
	// M – Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Primary Identification document Provided on Opening the Account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	//
	// For Kenya nationals, default should be the National Identification.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Identification document specified above.
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification document is Provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Any Other Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of Other Identification Document, where provided.
	// Mandatory if Other Identification Document is Provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Any Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification document is Provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// Client’s Email Address if provided
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Name of the Employer if not self-employed.
	Attribute("EmployerName", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Type Of employment.
	// Options Available :
	// Casual Contract
	// Permanent,
	// Pensionable
	//Self-Employed
	Attribute("EmploymentType", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The Position in Organisation
	Attribute("EmployeePosition", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// The Department Stationed
	Attribute("EmployeeDepartment", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// The band within which the Customer’s Gross Monthly Salary falls.
	// Options given are :
	// A - 0 To 50,000 KES
	// B - 50,000 To 100,000 KES
	// C – 100,000 To 200,000 KES
	// D – 200,000 To 250,000 KES
	// E - Over 250,000 KES
	Attribute("SalaryBand", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Employer’s Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// The Employer’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Any other Employer Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Employer’s Postal Address Line1.
	// First Line of the full employer’s Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Employer’s Postal Address Line 2.
	// Second Line of the full employer’s Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Employer’s Town of Postal Address.
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Employer’s Country of Postal Address.
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Post Code of Employer’s Address.
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Employer’s registered Office Street Address.
	// The Location ( Street Name) of the Employer’s Offices.
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "35")
	})

	// Employer’s Office building and office number.
	// The Location ( Building or Apartment No) of the address Employer’s Office
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "36")
	})

	// Office Plot Land Ref (LR) Number.
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "37")
	})

	// Employers Address Town.
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "38")
	})

	// Employer’s Location Country
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "39")
	})

})

// E. Non Individual Consumer Information
var CompanyInfo = Type("companyinfo", func() {
	Description("Non Individual Consumer Information")
	TypeName("CompanyInfo")
	ContentType("application/json")

	// The Name as Registered with the Registrar of Companies
	Attribute("RegisteredName", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Business or Trading Name
	Attribute("TradingName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Date Registered with the Registrar of Companies
	Attribute("RegistrationDate", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// The Registration Certificate Number
	Attribute("RegistrationNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Country of Registration of the Company, Defaults to Kenyan.
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Client Reference Number linking Company to Core Banking system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// The Account Number Linking Client to Bank’s Accounting System
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// The Type of Organisation.
	// The type of Organisation.
	// Options Available :
	// A - Limited Company
	// B - Sole Proprietor
	// C - NGO
	// D - Other
	Attribute("CompanyType", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Industry Code for the Line of business as per the Central Bank Returns – Options are :
	// Agriculture
	// Manufacturing
	// Building/Construction
	// Mining/Quarrying
	// Energy /water
	// Trade
	// Tourism/ Restaurant/Hotels
	// Transport/ Communications
	// Real Estate
	Attribute("IndustryCode", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// The Income Tax PIN Number.
	// May be Mandatory Later
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Value Added Tax Registration Number
	Attribute("VATNumber", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// 1 is a sole proprietor
	Attribute("NofDirectors", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// 1 if a sole proprietor
	Attribute("NofShareholders", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Business Email Address
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Company Website url
	Attribute("Website", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Status of Trading of the company –
	// Options Available :
	// Dormant
	// Actively Trading
	// Under Management
	// Dissolved
	Attribute("Status", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Options :
	// Actively Trading
	// Under management
	// Dissolved
	Attribute("TradingStatus", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// Date of the Current Status.
	// Provide the Date of the Last status.
	Attribute("StatusDate", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// The consumer’s Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// The consumer’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// The consumer’s other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// Consumer Postal Address Line1.
	// First Line of Company’s Full Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Consumer Postal Address Line 2.
	// Second Line of Company’s Full Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// Town of consumer Postal Address.
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Country of consumer Postal Address.
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Post Code of consumer Address.
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// Consumer Registered Office location building.
	// Location ( e.g. Building or Apartment) housing the company’s Registered Office.
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// Consumer Office location Street.
	// Street Name of Company’s Registered office.
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// Office location Plot Land Ref (LR) No.
	// Land Reference Number where Office is located
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Office Location Town
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Office Location Country
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "31")
	})
})

// F. Stakeholder Information
// Where the Stakeholder is an institution or organisation, the names should be provided
// should be the company registered name only. The Other Names (e.g. surname, first Names)
// should either not be provided or the Company name can be given in all the name fields.
var StakeholderInfo = Type("stakeholderinfo", func() {
	Description("Stakeholder Information")
	TypeName("StakeholderInfo")
	ContentType("application/json")

	// The Family Name of stakeholder, or the registered company name
	// if the stakeholder is a company or institution.
	Attribute("Surname", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The First Name of stakeholder.
	Attribute("Forename1", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// The Given Name of stakeholder.
	Attribute("Forename2", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Other Name or Initials
	Attribute("Forename3", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Options – Mr., Mrs., Miss, Ms, Dr. , Prof., Hon.
	Attribute("Salutation", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// Date of Birth or date of registration for non-individual stakeholders.
	// Cannot be in the Future
	Attribute("DateOfBirth", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Client Reference Number linking stakeholder to Core Banking system.
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Account Number linking stakeholder to Banking system.
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Gender of the stakeholder.
	// Options Available :
	// M – Male
	// F - Female
	// I - Institution or Organisation
	Attribute("Gender", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// The Kenya Revenue Authority Income Tax PIN Number.
	// May be required at a later Date
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Nationality of the Individual stakeholder or the country
	// of registration for non-individual stakeholders.
	// Default - Kenyan
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Applicable to non-individual stakeholders only.
	// Options to Select From :
	// M - Married
	// S - Single
	// D - Divorced
	// W - Widowed
	// U - Unknown
	Attribute("MaritalStatus", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Primary Identification document Provided by stake holder.
	// Options Are:
	// National ID
	// Passport
	// lien Registration
	// Service ID
	// Company Registration Certificate
	//
	// National Identification is the preferred document but the other are acceptable.
	// Alien Registration Certificates are issued to registered foreign nationals.
	// Service Identification documents are issued to the National forces like Police and Army.
	//The company registration Number is the Registration Number of the Institutional stake holder.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// The Number of the Primary Id Provided
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Any Secondary Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	//
	// For Individual Stakeholders
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Secondary Identification Doc Number if provided.
	// Mandatory if Secondary Identification Document Type is Provided.
	Attribute("SecondaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Any Other Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The Number of Other Identification Document, where provided .
	// Mandatory if Other Identification Document Type is Provided.
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Any Additional Identification Document Provided.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("AdditionalIDocument", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Additional Identification document Number, where provided.
	// Mandatory if Additional Identification Document Type is Provided.
	Attribute("AdditionalIDocNumber", String, "", func() {
		Meta("rpc:tag", "20")
	})
	Attribute("Email", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// The Company Registration Certificate Number
	Attribute("CompanyRegistrationNo", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// The Company Income Tax PIN Number
	Attribute("CompanyPINNo", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The Income Tax VAT registration Number.
	Attribute("CompanyVATNo", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// Type of stake held in the Company.
	// Options are :
	// A - Managing Director
	// B - Director
	// C - Share Holder
	// D - Company Secretary
	// E - Senior Management
	// For limited liability borrowers, A Profile Record for all directors
	// must be provided. Shareholders’ Profiles are not required.
	Attribute("StakeholderType", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Number of Shares held in the Company.
	// N/A for limited liability companies
	Attribute("NofShares", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// The Percentage of shares Held in the Company.
	// Percentage of shareholding in the Company for limited liability companies.
	// Not required if no of shares is provided.
	Attribute("PercentageShares", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// The Stake holder’s Primary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber1", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// The Stake holder’s Secondary Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber2", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// Stake holder’s other Telephone contact Number in the Form of :
	// CCC-AAA-NNNNNNN
	// Where :
	// CCC is the Country Code
	// AAA is the Area Code
	// NNNNNNN is the Number
	Attribute("TelephoneNumber3", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// Stake holder’s Postal Address Line1.
	// First Line of Full stakeholder Postal Address.
	Attribute("PostalAddress1", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// Stake holder’s Postal Address Line 2.
	// Second Line of Full stakeholder Postal Address.
	Attribute("PostalAddress2", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Town of Stake holder’s Postal Address
	Attribute("Town", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// Country of Stake holder’s Postal Address
	Attribute("Country", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Post Code of Stake holder’s Address
	Attribute("PostCode", String, "", func() {
		Meta("rpc:tag", "35")
	})

	// Stake holder’s Residence or registered Office if a non-individual.
	// Location of stakeholder ( e.g. Office Address or Building)
	Attribute("PhysicalAddress1", String, "", func() {
		Meta("rpc:tag", "36")
	})

	// Street of Stake holder’s residence or registered office.
	// Location street of Stakeholder’s location.
	Attribute("PhysicalAddress2", String, "", func() {
		Meta("rpc:tag", "37")
	})

	// Location street of Stakeholder’s location.
	Attribute("PhysicalAddress3", String, "", func() {
		Meta("rpc:tag", "38")
	})

	// Plot Land Ref (LR) No of Stake holder’s office/residence
	Attribute("PlotNumber", String, "", func() {
		Meta("rpc:tag", "39")
	})

	// Stake holder’s Town of location.
	Attribute("LocationTown", String, "", func() {
		Meta("rpc:tag", "40")
	})

	// Stake holder’s Country of location.
	Attribute("LocationCountry", String, "", func() {
		Meta("rpc:tag", "41")
	})
})


// G. Account Information
