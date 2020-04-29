package design

import (
	. "goa.design/goa/v3/dsl"
)

// L. Fraudulent Activities
var Fraud = Type("fraud", func() {
	Description("Fraudulent Activities")
	TypeName("Fraud")
	ContentType("application/json")

	// The Fraud Status as reported by the Institution.
	// Options Are :
	// A - Attempted
	// B - Suspected
	// C - Alleged
	// D - Proven
	// E - Admitted
	Attribute("FraudStatus", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// Whether Fraudster is an individual or Corporate
	// Options are:
	// I - Individual
	// C – Corporate entity
	//
	// Provide Stakeholder Profile information for Corporate Consumers.
	Attribute("IndividualOrCorporateName", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Name of the defaulter
	Attribute("Name", String, "", func() {
		Meta("rpc:tag", "3")
	})

	// Kenya Revenue Authority Income Tax PIN Number
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "4")
	})

	// Nationality of involved party or Country of Registration if Corporate.
	// Default Kenyan
	Attribute("Nationality", String, "", func() {
		Meta("rpc:tag", "5")
	})

	// The Primary Identification document Provided on Opening the Account.
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
	// The company registration Number is the Registration Number of the company account holder.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "6")
	})

	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "7")
	})

	// Other Secondary identification information on the account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("SecondaryIDocument", String, "", func() {
		Meta("rpc:tag", "8")
	})

	// Mandatory if Secondary Identification document provided.
	Attribute("SecondaryIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Any Other Identification information on the Account.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	Attribute("OtherIDocument", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Mandatory if Other Identification document provided
	Attribute("OtherIDocumentNumber", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// If Company and If info available
	Attribute("CompanyVATNumber", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// Bank and Branch Code where fraud committed
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Reference Number linking Customer account on which incidence involved to Banks’ Core system
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Account Number linking Account involved in the fraud incident to the Bank’s Accounting system.
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// The Account Product Type.
	// Options are:
	// A - Current Account
	// B - Loan Account
	// C - Credit Card
	// D - Line of Credit
	// E - Revolving Credit
	Attribute("ConsumerClassification", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Date the fraud incident took place or was detected
	Attribute("IncidentDate", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// Date the Fraud was reported to BFID or police
	Attribute("ReportDate", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// Amount involved in the fraud
	Attribute("Amount", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Actual loss incurred in the fraud as at reporting date.
	Attribute("LossAmount", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// Currency of the amounts involved default is Kenya shillings - KES
	Attribute("CurrencyCode", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// Short explanation of the nature of the incident
	Attribute("IncidentDetails", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Summary of any information available from the incident for investigation/forensic use
	Attribute("ForensicInformation", String, "", func() {
		Meta("rpc:tag", "23")
	})

})
