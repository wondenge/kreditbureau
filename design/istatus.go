package design

import (
	. "goa.design/goa/v3/dsl"
)

/* Individuals Account Status */

// A - Closed – No more admin processes running such as instalment demands
// or interest charges to account, and no further facilities can be offered
// on this account.
var ClosedIndividualsAccountStatus = Type("ClosedIndividualsAccountStatus", func() {
	Description("Closed: No further facilities can be offered on this account")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Attribute("AccountClosureReason", String, func() {
		Description("Account Closure Reason")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "60")
	})

	Required("Account Status Date", "AccountClosureReason")

})

// C - Write-Off – For facilities that don’t form part of the outstanding portfolio
// in the Balance Sheet, but are still outstanding in the books of accounts.
var WriteOffIndividualsAccountStatus = Type("WriteOffIndividualsAccountStatus", func() {
	Description("WriteOff: Facilities that don’t form part of the outstanding portfolio")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// D - Legal -with legal officer in court
var LegalIndividualsAccountStatus = Type("LegalIndividualsAccountStatus", func() {
	Description("Legal: With legal officer in court ")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// E - Collection- with collection bureau
var CollectionIndividualsAccountStatus = Type("CollectionIndividualsAccountStatus", func() {
	Description("Collection: With collection bureau")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// F – Active - For facilities that form part of the outstanding portfolio, and are
// reported in the Balance Sheet.
var ActiveIndividualsAccountStatus = Type("ActiveIndividualsAccountStatus", func() {
	Description("Active : Facilities that form part of the outstanding portfolio")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
var RescheduledIndividualsAccountStatus = Type("RescheduledIndividualsAccountStatus", func() {
	Description("Facility Rescheduled: Rescheduled/Restructured Facilities")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// H –Settled – The facility has been cleared. This status can only be used for
// revolving facilities
var SettledIndividualsAccountStatus = Type("SettledIndividualsAccountStatus", func() {
	Description("Settled: Facility has been cleared")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// J – Called Up : The facility has been called up. Once the client has paid up,
// the status should be updated to Option A, // H; or otherwise Option L
var CalledUpIndividualsAccountStatus = Type("CalledUpIndividualsAccountStatus", func() {
	Description("Called Up: Facility has been called up ")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// K – Suspended – The facility has been put on hold for an indefinite period of time
var SuspendedIndividualsAccountStatus = Type("SuspendedIndividualsAccountStatus", func() {
	Description("Suspended: Facility put on hold for an indefinite period of time")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// L– Client Deceased
var ClientDeceasedIndividualsAccountStatus = Type("ClientIndividualsAccountStatus", func() {
	Description("Client Deceased")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})
	Required("Account Status Date")

})

// M – Deferred – This refers to facilities whose payments have been put on hold
// for a definite period or in moratorium (Grace Period)
var DeferredIndividualsAccountStatus = Type("DeferredIndividualsAccountStatus", func() {
	Description("Deferred: Facilities whose payments put on hold for a definite period ")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Attribute("DeferredPaymentAmount", Int64, func() {
		Description("Deferred Payment Amount")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "63")

		// Mandatory if Account Status = Deferred.
		// Deferred Payment Date must be in the future
		Attribute("DeferredPaymentDate", String, func() {
			Description("Deferred Payment Date")
			MaxLength(8)
			Format(FormatDate)
			Meta("rpc:tag", "62")
		})

		Required("DeferredPaymentDate")
	})

	Required("Account Status Date", "DeferredPaymentAmount")

})

// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
var NotUpdatedIndividualsAccountStatus = Type("NotUpdatedIndividualsAccountStatus", func() {
	Description("Not Updated: reserved for CRBs (if last record status is not CLOSED)")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})

// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
var DisputedIndividualsAccountStatus = Type("DisputedIndividualsAccountStatus", func() {
	Description("Disputed: Client has disputed at the CRB")

	// The date of the status Change.
	Attribute("AccountStatusDate", String, func() {
		Description("Account Status Date")
		Format(FormatDate)
		Meta("rpc:tag", "59")
	})

	Required("Account Status Date")

})
