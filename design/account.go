package design

import (
	. "goa.design/goa/v3/dsl"
)

// G. Account Information
var Account = Type("account", func() {
	Description("")
	TypeName("Account")
	ContentType("application/json")

	// Bank and Branch Code of the Branch as Per the KBA Codes, in a Five-digit code.
	// The first two digits signify the bank code as provided by the KBA Secretariat.
	// The last three codes are the branch codes as provided by the bank to the KBA Secretariat.
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Account Holder Type.
	// Options are :
	// A - Personal
	// B - Corporate
	// C - Joint Account
	// D - SME Account
	// E - NGO Account
	// F - Other Account Type
	Attribute("AccountHolderType", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Kenya Revenue authority Income Tax Pin Number
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Account Holder Nationality
	// Default to Kenyan
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// The Primary Identification document Provided by the customer on opening the Account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	// Company Registration No
	//
	// National Identification is the preferred document but the other are acceptable.
	//Alien Registration Certificates are issued to registered foreign nationals.
	//Service Identification documents are issued to the National forces like Police and Army.
	//The company registration Number is the Registration Number of the Institution holding the Account.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// For Companies – the company registration Certificate Number.
	Attribute("PrimaryIDocumentNo", String, "", func() {
		Meta("rpc:tag", "6")
	})

	// Any other secondary identification document provided on account opening.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	// For companies – The Company VAT Registration
	Attribute("SecondaryID", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// If company then the Company VAT Number
	Attribute("SecondaryIDNumber", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// The Account Product Type.
	// Options are :
	// A - Current Account
	// B - Loan account
	// C - Credit Card
	// D - Line of Credit
	// E - Revolving Credit
	Attribute("AccountProductType", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// The Date the Account was opened
	Attribute("DateAccountOpened", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// If a loan or facility, the date due for payment
	Attribute("DueDate", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Client Reference Number linking Account to Banks Customer system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// The Account Number as per the Bank’s Accounting system
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Industry Code for the Line of business as per the Central Bank Codes. Options are :
	// Agriculture
	// Manufacturing
	// Building/Construction
	// Mining/Quarrying
	// Energy /water
	// Trade
	// Tourism/ Restaurant/Hotels
	// Transport/ Communications
	// Real Estate
	Attribute("GroupCode", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Original Amount or Approved Account Limit
	Attribute("OriginalAmount", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Available Balance
	Attribute("CurrentBalance", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Balance in account or Arrears
	Attribute("OverdueBalance", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// Date Account Fell Overdue
	Attribute("OverdueDate", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// If the Account is in arrears, the number of days the account is in arrears.
	Attribute("NofDaysInArrears", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// The Account status as at the time of reporting.
	// Options Available:
	// A - Closed
	// B - Dormant
	// C - Performing
	// D - Non-Performing
	// E - Write-Off
	// F - Legal
	// G - Collection
	Attribute("AccountStatus", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// The date of the status Change. For Normal Accounts or no change since
	// record creation, then the Date of the Account Opening.
	Attribute("AccountStatusDate", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// If status changed to closed, the reason the account was closed.
	// Mandatory if the Account is closed
	Attribute("AccountClosureReason", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Repayment Period in Months.
	// Where a repayment period is specified in week, then the repayment period
	// should be approximated up to the nearest month.
	Attribute("RepaymentPeriod", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// Date of first instalment made on the account if it is a loan account.
	Attribute("DateOfFirstPayment", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// The amount of the first instalment on the account if it is a loan.
	Attribute("FirstPaymentAmount", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// Date of last instalment made on the account if it is a loan account.
	Attribute("DateOfLastPayment", String, "", func() {
		Meta("rpc:tag", "26")
	})

	// The amount of last instalment.
	Attribute("LastPaymentAmount", String, "", func() {
		Meta("rpc:tag", "27")
	})

	// Whether the account is secured or not.
	// Options available :
	// U - Unsecured
	// S - Secured
	Attribute("SecurityType", String, "", func() {
		Meta("rpc:tag", "28")
	})

	// If a loan the number of instalments on the loan.
	Attribute("NumberOfInstalments", String, "", func() {
		Meta("rpc:tag", "29")
	})

	// The interest charged on the account if a loan.
	// Options Available :
	// Interest-Free
	// Reducing Bal
	// Flat Rate
	Attribute("InterestType", String, "", func() {
		Meta("rpc:tag", "30")
	})

	// The Rate of interest if a loan
	Attribute("InterestRate", String, "", func() {
		Meta("rpc:tag", "31")
	})

	// ISO Currency Code of the facility provided.
	// Default is Kenya Shillings (KES)
	Attribute("Currency", String, "", func() {
		Meta("rpc:tag", "32")
	})

	// Exchange Rate, Default to 1.00 for KES
	Attribute("CurrencyRate", String, "", func() {
		Meta("rpc:tag", "33")
	})

	// If Security was valued the Valuation amount of the security.
	Attribute("ValueOfSecurity", String, "", func() {
		Meta("rpc:tag", "34")
	})

	// Date of Valuation of Security.
	Attribute("ValuationDate", String, "", func() {
		Meta("rpc:tag", "35")
	})
})
