package design

import (
	. "goa.design/goa/v3/dsl"
)

// J. Credit Application Information

var CreditApplication = Type("creditapplication", func() {

	Description("")
	TypeName("CreditApplication")
	ContentType("application/json")

	// The Primary Identification document Provided by the Applicant.
	// Options Are:
	// National ID
	// Passport
	// Alien Registration
	// Service ID
	// Company Registration No
	//
	// National Identification is the preferred document for individuals but the others are acceptable.
	// Alien Registration Certificates are issued to registered foreign nationals.
	//Service Identification documents are issued to the National forces like Police and Army.
	//The company registration Number is the Registration Number of the Institutional Applicant.
	Attribute("PrimaryIDocument", String, "", func() {
		Meta("rpc:tag", "1")
	})

	// The Number of the Primary Id Provided
	Attribute("PrimaryIDocNumber", String, "", func() {
		Meta("rpc:tag", "2")
	})

	// Secondary Identification Document
	// Provided by applicant.
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

	// Any Other Additional Identification Document Provided by applicant.
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

	// Additional Identification Document Provided by applicant.
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

	// Applicant’s Kenya Revenue authority PIN Number
	Attribute("PINNumber", String, "", func() {
		Meta("rpc:tag", "9")
	})

	// Applicable to Companies and non-individual consumers
	Attribute("CompanyRegistrationNumber", String, "", func() {
		Meta("rpc:tag", "10")
	})

	// Applicable to Companies and non-individual consumers
	Attribute("CompanyVATNumber", String, "", func() {
		Meta("rpc:tag", "11")
	})

	// Bank and Branch where the Application was made.
	Attribute("BranchCode", String, "", func() {
		Meta("rpc:tag", "12")
	})

	// Client Number making the Application, if used.
	Attribute("ClientNumber", String, "", func() {
		Meta("rpc:tag", "13")
	})

	// Account Number if available of applicant
	Attribute("AccountNumber", String, "", func() {
		Meta("rpc:tag", "14")
	})

	// Application reference Number provided by the institution
	Attribute("ApplicationNumber", String, "", func() {
		Meta("rpc:tag", "15")
	})

	// Whether the application is for a secured or unsecured facility
	// Options :
	// U - Unsecured
	// S - Secured
	Attribute("FacilityApplicationType", String, "", func() {
		Meta("rpc:tag", "16")
	})

	// Date Application was made
	Attribute("ApplicationDate", String, "", func() {
		Meta("rpc:tag", "17")
	})

	// The amount of Application for facility
	Attribute("ApplicationAmount", String, "", func() {
		Meta("rpc:tag", "18")
	})

	// ISO Currency Code of Facility. Default is Kenya Shillings ( KES)
	Attribute("ApplicationCurrency", String, "", func() {
		Meta("rpc:tag", "19")
	})

	// Kenya equivalent of the Application.
	// This should be the same as the application amount if the amount is in Kenya shillings.
	Attribute("KESEquivalent", String, "", func() {
		Meta("rpc:tag", "20")
	})

	// The Exchange rate used in the application.
	// If application is in Kenya shillings the Rate should be 1.00
	Attribute("ExchangeRate", String, "", func() {
		Meta("rpc:tag", "21")
	})

	// Product Type.
	// Options are :
	// A – Current Account
	// B - Loan Account
	// C - Credit Card
	// D - Line of Credit
	// E - Revolving
	Attribute("ProductType", String, "", func() {
		Meta("rpc:tag", "22")
	})

	// Repayment period in Number of Months Applied for.
	// Where the repayment is not specified in months, then the period shall
	// be approximate dup to the nearest months.
	Attribute("ApplicationDuration", String, "", func() {
		Meta("rpc:tag", "23")
	})

	// The status of the Application at reporting Time.
	// Options are :
	// A - Pending
	// B - Awaiting documentation
	// C - Securities Perfection
	// D - Declined
	// E - Withdrawn
	// F - Approved
	// G - Pending disbursement
	Attribute("ApplicationStatus", String, "", func() {
		Meta("rpc:tag", "24")
	})

	// The date of the Status above.
	// When the status changes from one state to another as the application is processed,
	// the date of the status change from the processing is reported.
	Attribute("StatusDate", String, "", func() {
		Meta("rpc:tag", "25")
	})

	// The Reason for the status change
	Attribute("StatusUpdateReason", String, "", func() {
		Meta("rpc:tag", "26")
	})
})
